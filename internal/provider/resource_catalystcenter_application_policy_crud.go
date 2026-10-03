// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var _ resource.ResourceWithModifyPlan = &ApplicationPolicyResource{}

// The Application Policy API has a single write endpoint,
// POST /dna/intent/api/v1/app-policy-intent, whose body carries three buckets:
// createList, updateList and deleteList. Reads come from a different endpoint,
// GET /dna/intent/api/v1/app-policy?policyScope=<scope>, which returns the sibling
// policies flat.
//
// The generator can express Create (createList) and Read, but not Update or Delete:
//   - Update must use the updateList key and must echo back opaque ids the
//     controller assigned (the sibling id plus ids nested under advancedPolicyScope,
//     exclusiveContract and producer). Sending the same payload under createList is
//     rejected with NCAS10239 "Policy name already exists".
//   - Delete has no DELETE verb; it is a POST carrying deleteList with the ids of
//     every sibling policy.
//
// Both are therefore implemented here, outside the //template markers, so they
// survive `go generate`.

// opaqueIDPaths are the controller-assigned identifiers that an updateList entry
// must carry over from the currently stored sibling policy.
var opaqueIDPaths = []string{
	"id",
	"advancedPolicyScope.id",
	"advancedPolicyScope.advancedPolicyScopeElement.0.id",
	"exclusiveContract.id",
	"exclusiveContract.clause.0.id",
	"producer.id",
	"consumer.id",
}

// getSiblings returns the sibling policies currently stored for a policy scope,
// keyed by their name.
func (r *ApplicationPolicyResource) getSiblings(ctx context.Context, scope string) (map[string]gjson.Result, error) {
	res, err := r.client.Get("/dna/intent/api/v1/app-policy?policyScope=" + url.QueryEscape(scope))
	if err != nil {
		return nil, err
	}
	siblings := make(map[string]gjson.Result)
	res.Get("response").ForEach(func(_, v gjson.Result) bool {
		siblings[v.Get("name").String()] = v
		return true
	})
	return siblings, nil
}

// carryOverOpaqueIDs copies the controller-assigned identifiers from a stored
// sibling policy into a planned entry, so the entry updates in place instead of
// being treated as a new object.
func carryOverOpaqueIDs(item string, existing gjson.Result) string {
	for _, p := range opaqueIDPaths {
		if v := existing.Get(p); v.Exists() {
			item, _ = sjson.Set(item, p, v.Value())
		}
	}
	return item
}

// ModifyPlan checks, at plan time, that each sibling policy belongs to the policy
// it is declared under.
//
// The controller groups sibling policies by their `policyScope` field, and
// GET /app-policy?policyScope=<scope> filters on that field alone. Two different
// checks follow from this, with different severity:
//
//   - A wrong `policy_scope` is an error. The sibling is stored under another
//     scope, so this resource never reads it back and never deletes it. It is
//     left behind on the controller.
//   - A `name` without the scope prefix is a warning only. Verified against the
//     controller: the API accepts such a name, and GET still returns the sibling
//     because the filter uses `policyScope`, not the name. Only the naming
//     convention that Catalyst Center uses to group rows in the GUI is broken.
func (r *ApplicationPolicyResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan ApplicationPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	scope := plan.PolicyScope
	if scope.IsNull() || scope.IsUnknown() {
		return
	}
	want := scope.ValueString()
	prefix := want + "_"

	for i, item := range plan.Items {
		itemPath := path.Root("items").AtListIndex(i)

		if !item.Name.IsNull() && !item.Name.IsUnknown() {
			if name := item.Name.ValueString(); !strings.HasPrefix(name, prefix) {
				resp.Diagnostics.AddAttributeWarning(
					itemPath.AtName("name"),
					"Sibling policy name does not follow the naming convention",
					fmt.Sprintf("Name %q does not start with %q. Catalyst Center names sibling policies "+
						"<policy_scope>_<application set>, <policy_scope>_queuing_customization or "+
						"<policy_scope>_global_policy_configuration. The controller accepts this name and "+
						"Terraform still manages the sibling, but the Application Policy GUI groups rows by "+
						"this convention.", name, prefix),
				)
			}
		}

		if !item.PolicyScope.IsNull() && !item.PolicyScope.IsUnknown() {
			if got := item.PolicyScope.ValueString(); got != want {
				resp.Diagnostics.AddAttributeError(
					itemPath.AtName("policy_scope"),
					"Mismatched policy scope",
					fmt.Sprintf("Item policy_scope %q must equal the resource's policy_scope %q. "+
						"The controller stores the sibling under %q, so this resource never reads it back "+
						"and never deletes it, and the sibling is left behind on the controller.",
						got, want, got),
				)
			}
		}
	}
}

func (r *ApplicationPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplicationPolicy

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	scope := plan.PolicyScope.ValueString()
	existing, err := r.getSiblings(ctx, scope)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read current policy (GET), got error: %s", err))
		return
	}

	// plan.toBody produces {"createList":[...]}; redistribute those entries into the
	// create/update buckets depending on whether the sibling already exists.
	planned := gjson.Get(plan.toBody(ctx, state), "createList")

	body := ""
	createCount, updateCount := 0, 0
	keep := make(map[string]bool)

	planned.ForEach(func(_, item gjson.Result) bool {
		name := item.Get("name").String()
		keep[name] = true
		if cur, ok := existing[name]; ok {
			body, _ = sjson.SetRaw(body, "updateList.-1", carryOverOpaqueIDs(item.Raw, cur))
			updateCount++
		} else {
			body, _ = sjson.SetRaw(body, "createList.-1", item.Raw)
			createCount++
		}
		return true
	})

	// Siblings present on the controller but no longer in the plan are removed.
	deleteCount := 0
	for name, cur := range existing {
		if keep[name] {
			continue
		}
		if id := cur.Get("id"); id.Exists() {
			body, _ = sjson.Set(body, "deleteList.-1", id.String())
			deleteCount++
		}
	}

	if createCount == 0 && updateCount == 0 && deleteCount == 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Update had nothing to do", plan.Id.ValueString()))
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update create=%d update=%d delete=%d", plan.Id.ValueString(), createCount, updateCount, deleteCount))

	res, err := r.client.Post("/dna/intent/api/v1/app-policy-intent", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplicationPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplicationPolicy

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	existing, err := r.getSiblings(ctx, state.PolicyScope.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read current policy (GET), got error: %s", err))
		return
	}

	body := ""
	count := 0
	for _, cur := range existing {
		if id := cur.Get("id"); id.Exists() {
			body, _ = sjson.Set(body, "deleteList.-1", id.String())
			count++
		}
	}

	if count > 0 {
		res, err := r.client.Post("/dna/intent/api/v1/app-policy-intent", body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (POST), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully (%d sibling policies)", state.Id.ValueString(), count))

	resp.State.RemoveResource(ctx)
}
