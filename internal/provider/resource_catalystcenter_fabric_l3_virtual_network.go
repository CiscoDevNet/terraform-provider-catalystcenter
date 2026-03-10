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
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &FabricL3VirtualNetworkResource{}
var _ resource.ResourceWithImportState = &FabricL3VirtualNetworkResource{}

func NewFabricL3VirtualNetworkResource() resource.Resource {
	return &FabricL3VirtualNetworkResource{}
}

type FabricL3VirtualNetworkResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *FabricL3VirtualNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l3_virtual_network"
}

func (r *FabricL3VirtualNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Fabric L3 Virtual Network.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the layer 3 virtual network. If `INFRA_VN` or `DEFAULT_VN` is used, those layer 3 virtual networks will be updated instead of created.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fabric_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IDs of the fabrics this layer 3 virtual network is to be assigned to.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"anchored_site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fabric ID of the fabric site this layer 3 virtual network is to be anchored at.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"additive_fabric_ids": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to `true`, the `fabric_ids` declared in this resource are merged with the existing fabric associations on Catalyst Center (additive on create/update, subtractive on delete), rather than replacing the entire set. Use this in environments where multiple resources or external processes manage fabric associations for the same L3 Virtual Network.").String,
				Optional:            true,
			},
		},
	}
}

func (r *FabricL3VirtualNetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

func (r *FabricL3VirtualNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricL3VirtualNetwork

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Check if VirtualNetworkName is "INFRA_VN" or "DEFAULT_VN"
	if plan.VirtualNetworkName.ValueString() == "INFRA_VN" || plan.VirtualNetworkName.ValueString() == "DEFAULT_VN" {
		// Perform GET request to retrieve the ID (and existing fabricIds for additive mode)
		params := "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
		res, err := r.client.Get(plan.getPath() + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		// Extract the ID from the GET response
		id := res.Get("response.0.id").String()
		if id == "" {
			resp.Diagnostics.AddError("Invalid Response", "The GET request did not return a valid ID.")
			return
		}

		plan.Id = types.StringValue(id)
		tflog.Debug(ctx, fmt.Sprintf("Updated plan.Id: %s", plan.Id.ValueString()))

		if plan.AdditiveFabricIds.ValueBool() {
			// Additive mode: merge existing fabricIds with plan fabricIds
			var planFabricIds []string
			plan.FabricIds.ElementsAs(ctx, &planFabricIds, false)

			existingFabricIds := []string{}
			for _, fid := range res.Get("response.0.fabricIds").Array() {
				existingFabricIds = append(existingFabricIds, fid.String())
			}

			// Compute union: existing ∪ plan (deduplicated)
			mergedFabricIds := make([]string, len(existingFabricIds))
			copy(mergedFabricIds, existingFabricIds)
			for _, fid := range planFabricIds {
				if !slices.Contains(mergedFabricIds, fid) {
					mergedFabricIds = append(mergedFabricIds, fid)
				}
			}

			MAX_RETRIES := 3
			for try := 0; try <= MAX_RETRIES; try++ {
				mergedPlan := plan
				mergedPlan.FabricIds, _ = types.SetValueFrom(ctx, types.StringType, mergedFabricIds)
				body := mergedPlan.toBody(ctx, FabricL3VirtualNetwork{})
				res, err = r.client.Put(plan.getPath(), body, cc.UseMutex)
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
					return
				}

				params = "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
				res, err = r.client.Get(plan.getPath() + params)
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
					return
				}

				getFabricIds := []string{}
				for _, fid := range res.Get("response.0.fabricIds").Array() {
					getFabricIds = append(getFabricIds, fid.String())
				}
				sortedMerged := make([]string, len(mergedFabricIds))
				copy(sortedMerged, mergedFabricIds)
				slices.Sort(sortedMerged)
				slices.Sort(getFabricIds)
				if slices.Equal(getFabricIds, sortedMerged) {
					break
				}
				if try != MAX_RETRIES {
					tflog.Warn(ctx, fmt.Sprintf("%s: Assigned fabric ids does not match. Retrying for %d time", plan.Id.ValueString(), try))
				} else {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Assigned fabric ids does not match. Please try again later, %s", res.String()))
					return
				}
			}
		} else {
			// Normal mode: PUT with plan fabricIds (overwrites)
			body := plan.toBody(ctx, FabricL3VirtualNetwork{})
			res, err = r.client.Put(plan.getPath(), body, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		// Create object
		body := plan.toBody(ctx, FabricL3VirtualNetwork{})

		params := ""
		res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
		if err != nil {
			if r.AllowExistingOnCreate {
				tflog.Info(ctx, fmt.Sprintf("Failed to configure object (%s), got error: %s, %s. allow_existing_on_create is true, beginning update", "POST", err, res.String()))
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
				return
			}
		}

		var fabricIds []string
		plan.FabricIds.ElementsAs(ctx, &fabricIds, false)
		params = ""
		params += "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
		if len(fabricIds) > 0 {
			params += "&fabricIds=" + url.QueryEscape(fabricIds[0])
		}
		if !plan.AnchoredSiteId.IsNull() && plan.AnchoredSiteId.ValueString() != "" {
			params += "&anchoredSiteId=" + url.QueryEscape(plan.AnchoredSiteId.ValueString())
		}
		res, err = r.client.Get(plan.getPath() + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		plan.Id = types.StringValue(res.Get("response.0.id").String())
		if !r.AllowExistingOnCreate {
			tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))
		} else {
			if plan.AdditiveFabricIds.ValueBool() {
				// Additive mode under AllowExistingOnCreate: merge existing fabricIds with plan
				var planFabricIds []string
				plan.FabricIds.ElementsAs(ctx, &planFabricIds, false)

				MAX_RETRIES := 3
				for try := 0; try <= MAX_RETRIES; try++ {
					// GET existing fabricIds
					params = "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
					res, err = r.client.Get(plan.getPath() + params)
					if err != nil {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
						return
					}

					existingFabricIds := []string{}
					for _, fid := range res.Get("response.0.fabricIds").Array() {
						existingFabricIds = append(existingFabricIds, fid.String())
					}

					// Compute union: existing ∪ plan (deduplicated)
					mergedFabricIds := make([]string, len(existingFabricIds))
					copy(mergedFabricIds, existingFabricIds)
					for _, fid := range planFabricIds {
						if !slices.Contains(mergedFabricIds, fid) {
							mergedFabricIds = append(mergedFabricIds, fid)
						}
					}

					mergedPlan := plan
					mergedPlan.FabricIds, _ = types.SetValueFrom(ctx, types.StringType, mergedFabricIds)
					body = mergedPlan.toBody(ctx, FabricL3VirtualNetwork{Id: plan.Id})
					res, err = r.client.Put(plan.getPath(), body, cc.UseMutex)
					if err != nil {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
						return
					}

					params = "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
					res, err = r.client.Get(plan.getPath() + params)
					if err != nil {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
						return
					}

					getFabricIds := []string{}
					for _, fid := range res.Get("response.0.fabricIds").Array() {
						getFabricIds = append(getFabricIds, fid.String())
					}
					sortedMerged := make([]string, len(mergedFabricIds))
					copy(sortedMerged, mergedFabricIds)
					slices.Sort(sortedMerged)
					slices.Sort(getFabricIds)
					if slices.Equal(getFabricIds, sortedMerged) {
						break
					}
					if try != MAX_RETRIES {
						tflog.Warn(ctx, fmt.Sprintf("%s: Assigned fabric ids does not match. Retrying for %d time", plan.Id.ValueString(), try))
					} else {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Assigned fabric ids does not match. Please try again later, %s", res.String()))
						return
					}
				}
				tflog.Debug(ctx, fmt.Sprintf("%s: Fallback to additive update existing resource finished successfully", plan.Id.ValueString()))
			} else {
				// Normal mode under AllowExistingOnCreate: plain PUT with plan fabricIds
				params = ""
				body = plan.toBody(ctx, FabricL3VirtualNetwork{Id: plan.Id})
				res, err = r.client.Put(plan.getPath()+params, body, cc.UseMutex)
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
					return
				}
				tflog.Debug(ctx, fmt.Sprintf("%s: Fallback to update existing resource finished successfully", plan.Id.ValueString()))
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FabricL3VirtualNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricL3VirtualNetwork

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	var fabricIds []string
	state.FabricIds.ElementsAs(ctx, &fabricIds, false)
	params := ""
	params += "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
	if len(fabricIds) > 0 {
		params += "&fabricIds=" + url.QueryEscape(fabricIds[0])
	}
	if !state.AnchoredSiteId.IsNull() && state.AnchoredSiteId.ValueString() != "" {
		params += "&anchoredSiteId=" + url.QueryEscape(state.AnchoredSiteId.ValueString())
	}
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// In additive mode, save declared fabric_ids before updateFromBody overwrites them
	// with the full API response (which includes externally-managed IDs). We restore
	// them afterwards so Terraform only tracks what was declared, preventing spurious
	// plan diffs caused by fabric IDs managed outside this resource.
	savedFabricIds := state.FabricIds

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	if state.AdditiveFabricIds.ValueBool() {
		state.FabricIds = savedFabricIds
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *FabricL3VirtualNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricL3VirtualNetwork

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// If only additive_fabric_ids changed (fabric_ids is unchanged), no API call is needed.
	// The flag is a Terraform-only behaviour switch — flipping it alone does not require
	// any modification on Catalyst Center.
	if plan.FabricIds.Equal(state.FabricIds) && plan.AnchoredSiteId.Equal(state.AnchoredSiteId) {
		tflog.Debug(ctx, fmt.Sprintf("%s: Only additive_fabric_ids changed, skipping API call", plan.Id.ValueString()))
		diags = resp.State.Set(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	if plan.AdditiveFabricIds.ValueBool() {
		// Additive mode: merge (existing - state) ∪ plan fabric IDs
		var stateFabricIds []string
		state.FabricIds.ElementsAs(ctx, &stateFabricIds, false)
		var planFabricIds []string
		plan.FabricIds.ElementsAs(ctx, &planFabricIds, false)

		// If transitioning from normal mode to additive mode, state.FabricIds may contain
		// externally-managed IDs that were synced by the previous Read. In that case we
		// must not subtract state from existing, otherwise we would remove IDs we don't own.
		// Instead, treat this first apply as a pure union: existing ∪ plan.
		transitioningToAdditive := !state.AdditiveFabricIds.ValueBool()

		MAX_RETRIES := 3
		for try := 0; try <= MAX_RETRIES; try++ {
			// GET existing fabricIds from CatC
			params := "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
			res, err := r.client.Get(plan.getPath()+params, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
				return
			}

			existingFabricIds := []string{}
			for _, fid := range res.Get("response.0.fabricIds").Array() {
				existingFabricIds = append(existingFabricIds, fid.String())
			}

			// Compute new fabric IDs:
			// - Steady-state additive: (existingFabricIds - stateFabricIds) ∪ planFabricIds
			//   This correctly handles adds and removes relative to what this resource declared.
			// - Transition to additive: existing ∪ plan (no subtraction — state may contain
			//   externally-managed IDs from the previous normal-mode Read, subtracting them
			//   would incorrectly remove IDs we don't own).
			newFabricIds := []string{}
			if transitioningToAdditive {
				newFabricIds = append(newFabricIds, existingFabricIds...)
				for _, fid := range planFabricIds {
					if !slices.Contains(newFabricIds, fid) {
						newFabricIds = append(newFabricIds, fid)
					}
				}
			} else {
				for _, fid := range existingFabricIds {
					if !slices.Contains(stateFabricIds, fid) {
						newFabricIds = append(newFabricIds, fid)
					}
				}
				for _, fid := range planFabricIds {
					if !slices.Contains(newFabricIds, fid) {
						newFabricIds = append(newFabricIds, fid)
					}
				}
			}

			mergedPlan := plan
			mergedPlan.FabricIds, _ = types.SetValueFrom(ctx, types.StringType, newFabricIds)
			body := mergedPlan.toBody(ctx, state)
			res, err = r.client.Put(plan.getPath(), body, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}

			// Verify
			params = "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
			res, err = r.client.Get(plan.getPath()+params, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
				return
			}

			getFabricIds := []string{}
			for _, fid := range res.Get("response.0.fabricIds").Array() {
				getFabricIds = append(getFabricIds, fid.String())
			}
			sortedNew := make([]string, len(newFabricIds))
			copy(sortedNew, newFabricIds)
			slices.Sort(sortedNew)
			slices.Sort(getFabricIds)
			if slices.Equal(getFabricIds, sortedNew) {
				break
			}
			if try != MAX_RETRIES {
				tflog.Warn(ctx, fmt.Sprintf("%s: Assigned fabric ids does not match. Retrying for %d time", plan.Id.ValueString(), try))
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Assigned fabric ids does not match. Please try again later, %s", res.String()))
				return
			}
		}
	} else {
		// Normal mode: plain PUT with plan fabricIds
		body := plan.toBody(ctx, state)
		params := ""
		res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
		_ = res
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FabricL3VirtualNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricL3VirtualNetwork

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	if state.AdditiveFabricIds.ValueBool() {
		// Additive mode: subtract only the fabric IDs this resource owns; do not delete the VN itself
		var stateFabricIds []string
		state.FabricIds.ElementsAs(ctx, &stateFabricIds, false)

		MAX_RETRIES := 3
		for try := 0; try <= MAX_RETRIES; try++ {
			params := "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
			res, err := r.client.Get(state.getPath()+params, cc.UseMutex)
			if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
				resp.State.RemoveResource(ctx)
				return
			} else if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
				return
			}

			existingFabricIds := []string{}
			for _, fid := range res.Get("response.0.fabricIds").Array() {
				existingFabricIds = append(existingFabricIds, fid.String())
			}

			// newFabricIds = existingFabricIds - stateFabricIds (subtract only what this resource owns)
			newFabricIds := []string{}
			for _, fid := range existingFabricIds {
				if !slices.Contains(stateFabricIds, fid) {
					newFabricIds = append(newFabricIds, fid)
				}
			}

			mergedState := state
			mergedState.FabricIds, _ = types.SetValueFrom(ctx, types.StringType, newFabricIds)
			body := mergedState.toBody(ctx, FabricL3VirtualNetwork{})
			res, err = r.client.Put(state.getPath(), body, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}

			// Verify
			params = "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
			res, err = r.client.Get(state.getPath()+params, cc.UseMutex)
			if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
				resp.State.RemoveResource(ctx)
				return
			} else if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
				return
			}

			getFabricIds := []string{}
			for _, fid := range res.Get("response.0.fabricIds").Array() {
				getFabricIds = append(getFabricIds, fid.String())
			}
			sortedNew := make([]string, len(newFabricIds))
			copy(sortedNew, newFabricIds)
			slices.Sort(sortedNew)
			slices.Sort(getFabricIds)
			if slices.Equal(getFabricIds, sortedNew) {
				break
			}
			if try != MAX_RETRIES {
				tflog.Warn(ctx, fmt.Sprintf("%s: Assigned fabric ids does not match. Retrying for %d time", state.Id.ValueString(), try))
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Assigned fabric ids does not match. Please try again later, %s", res.String()))
				return
			}
		}
	} else if state.VirtualNetworkName.ValueString() != "INFRA_VN" && state.VirtualNetworkName.ValueString() != "DEFAULT_VN" {
		// Normal mode, normal VN: DELETE the virtual network
		params := ""
		params += "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
		res, err := r.client.Delete(state.getPath()+params, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		// Normal mode, INFRA_VN / DEFAULT_VN: clear all fabric site associations via PUT with null fabricIds
		state.FabricIds = types.SetNull(types.StringType)
		body := state.toBody(ctx, FabricL3VirtualNetwork{})
		res, err := r.client.Put(state.getPath(), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
		_ = res
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *FabricL3VirtualNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <virtual_network_name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("virtual_network_name"), idParts[0])...)
}

// End of section. //template:end import
