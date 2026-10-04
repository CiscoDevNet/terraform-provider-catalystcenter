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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ApplicationPolicyResource{}
var _ resource.ResourceWithImportState = &ApplicationPolicyResource{}

func NewApplicationPolicyResource() resource.Resource {
	return &ApplicationPolicyResource{}
}

type ApplicationPolicyResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *ApplicationPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_policy"
}

func (r *ApplicationPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages a Catalyst Center Application QoS Policy. A single logical policy is not one API object: the controller stores it as a set of sibling group-based policies that all share the same `policy_scope`. Each entry in `items` is one such sibling — typically one per application set (carrying a `BUSINESS_RELEVANCE` clause), plus a `<policy>_queuing_customization` entry carrying the queuing profile reference, and optionally a `<policy>_global_policy_configuration` entry carrying `APPLICATION_POLICY_KNOBS`. A site may only be used by one wired policy (`NCAS10157`). `terraform import` expects the policy scope.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_scope": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the policy. On the controller every sibling object carries this as its `policyScope`, and their names are prefixed with it.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"undeploy_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("What to do with the devices when this policy is destroyed. `DELETED` removes the policy from the devices, `RESTORED` returns them to their original configuration. Defaults to `DELETED`, matching the GUI. The controller requires the policy to be undeployed before its objects can be removed, so destroy always performs that step; this only selects which action it uses. Never sent on create or update.").AddStringEnumDescription("DELETED", "RESTORED").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DELETED", "RESTORED"),
				},
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The sibling group-based policies making up this policy. Modelled as a set because the controller does not preserve ordering.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the sibling policy, conventionally `<policy_scope>_<application set>`, `<policy_scope>_queuing_customization` or `<policy_scope>_global_policy_configuration`.").String,
							Required:            true,
						},
						"policy_scope": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Must equal the resource's `policy_scope`").String,
							Required:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority of the sibling policy. `100` normally, `4095` when the producer refers to an application scalable group.").String,
							Optional:            true,
						},
						"delete_policy_status": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Deployment state of the sibling policy").AddStringEnumDescription("NONE", "DELETED", "RESTORED").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NONE", "DELETED", "RESTORED"),
							},
						},
						"advanced_policy_scope_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name carried by the advanced policy scope, normally the policy scope").String,
							Optional:            true,
						},
						"site_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Site IDs this sibling policy is deployed to").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ssids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSIDs this sibling policy applies to").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"clause_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Kind of exclusive contract clause carried by this sibling policy").AddStringEnumDescription("BUSINESS_RELEVANCE", "APPLICATION_POLICY_KNOBS").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BUSINESS_RELEVANCE", "APPLICATION_POLICY_KNOBS"),
							},
						},
						"relevance_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Business relevance, when `clause_type` is `BUSINESS_RELEVANCE`").AddStringEnumDescription("BUSINESS_RELEVANT", "BUSINESS_IRRELEVANT", "DEFAULT").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BUSINESS_RELEVANT", "BUSINESS_IRRELEVANT", "DEFAULT"),
							},
						},
						"device_removal_behavior": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Device removal behaviour, when `clause_type` is `APPLICATION_POLICY_KNOBS`").AddStringEnumDescription("DELETE", "RESTORE", "IGNORE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("DELETE", "RESTORE", "IGNORE"),
							},
						},
						"host_tracking_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether host tracking is enabled, when `clause_type` is `APPLICATION_POLICY_KNOBS`").String,
							Optional:            true,
						},
						"queuing_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the queuing profile, set on the `<policy_scope>_queuing_customization` sibling policy only.").String,
							Optional:            true,
						},
						"application_set_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the application set this sibling policy applies to, set on `BUSINESS_RELEVANCE` sibling policies.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplicationPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ApplicationPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplicationPolicy{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.PolicyScope.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ApplicationPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?policyScope=" + url.QueryEscape(state.PolicyScope.ValueString())
	res, err := r.client.Get("/dna/intent/api/v1/app-policy" + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	if data := res.Get("response"); !data.Exists() || len(data.Array()) == 0 {
		resp.State.RemoveResource(ctx)
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Update and Delete are hand-written in resource_catalystcenter_application_policy_crud.go
// because this API uses a single POST endpoint with createList/updateList/deleteList
// buckets, which the generator cannot express.

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplicationPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <policy_scope>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_scope"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *ApplicationPolicyResource) ReadCache(ctx context.Context, req resource.ReadRequest, state ApplicationPolicy, params string) (cc.Res, error) {
	var err error
	cacheKey := "ApplicationPolicy::"

	_, cacheSuffix, found := strings.Cut(params, "?")
	queryPart, err := url.ParseQuery(cacheSuffix)
	if err == nil {
		delete(queryPart, "id")
		newQuery := queryPart.Encode()
		cacheSuffix = "?" + newQuery
		cacheKey += cacheSuffix
	}

	cachedValue, found := r.cache.Get(cacheKey)
	if found {
		tflog.Debug(ctx, fmt.Sprintf("hit cache for %s", cacheKey))
		ccRes, ok := cachedValue.(cc.Res)
		if ok {
			return ccRes, nil
		}
		tflog.Info(ctx, fmt.Sprintf("Invalid cache entry type for %s", cacheKey))
		r.cache.Delete(cacheKey)
	}
	res, err := r.client.Get("/dna/intent/api/v1/app-policy" + params)
	singleRes := res
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache

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

	// The controller will not release a deployed policy's objects until the policy
	// has been withdrawn from its devices, so destroy is two POSTs rather than one.
	// The first carries updateList with deletePolicyStatus set, which pushes the
	// withdrawal; the second carries deleteList, which removes the objects. Doing
	// only the second leaves the QoS configuration behind on the devices. This
	// mirrors the GUI, which greys out Delete until Undeploy has run.
	undeployAction := "DELETED"
	if !state.UndeployAction.IsNull() && state.UndeployAction.ValueString() != "" {
		undeployAction = state.UndeployAction.ValueString()
	}

	undeployBody := ""
	deleteBody := ""
	count := 0
	// Build the withdrawal entries the same way Update does: from toBody, not from
	// the raw GET. Echoing a GET response back is rejected with NCSP11104 because
	// it carries read-only fields the intent endpoint will not accept.
	stateBody := gjson.Get(state.toBody(ctx, state), "createList")
	stateBody.ForEach(func(_, item gjson.Result) bool {
		cur, ok := existing[item.Get("name").String()]
		if !ok {
			return true
		}
		id := cur.Get("id")
		if !id.Exists() {
			return true
		}
		entry := carryOverOpaqueIDs(item.Raw, cur)
		entry, _ = sjson.Set(entry, "deletePolicyStatus", undeployAction)
		undeployBody, _ = sjson.SetRaw(undeployBody, "updateList.-1", entry)
		deleteBody, _ = sjson.Set(deleteBody, "deleteList.-1", id.String())
		count++
		return true
	})

	if count > 0 {
		res, err := r.client.Post("/dna/intent/api/v1/app-policy-intent", undeployBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to undeploy object before delete (POST %s), got error: %s, %s", undeployAction, err, res.String()))
			return
		}

		res, err = r.client.Post("/dna/intent/api/v1/app-policy-intent", deleteBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (POST), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully (%s, %d sibling policies)", state.Id.ValueString(), undeployAction, count))

	resp.State.RemoveResource(ctx)
}
