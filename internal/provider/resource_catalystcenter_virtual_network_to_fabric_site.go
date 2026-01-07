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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &VirtualNetworkToFabricSiteResource{}
var _ resource.ResourceWithImportState = &VirtualNetworkToFabricSiteResource{}

func NewVirtualNetworkToFabricSiteResource() resource.Resource {
	return &VirtualNetworkToFabricSiteResource{}
}

type VirtualNetworkToFabricSiteResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *VirtualNetworkToFabricSiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_network_to_fabric_site"
}

func (r *VirtualNetworkToFabricSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manage assignment of a virtual network (VN) to an SD-Access Fabric Site.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Virtual Network Name").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"virtual_network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the Virtual Network").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fabric_site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric this Layer 3 Virtual Network is to be assigned to.").String,
				Required:            true,
			},
		},
	}
}

func (r *VirtualNetworkToFabricSiteResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

func (r *VirtualNetworkToFabricSiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VirtualNetworkToFabricSite

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	MAX_RETRIES := 3
	for try := 0; try <= MAX_RETRIES; try++ {
		params := ""
		params += "?virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString())
		res, err := r.client.Get(plan.getPath() + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

		// Extract existing fabricIds list
		existingFabricIds := []string{}
		for _, id := range res.Get("response.0.fabricIds").Array() {
			existingFabricIds = append(existingFabricIds, id.String())
		}

		// Append the new FabricSiteId if it’s not already in the list
		alreadyExists := slices.Contains(existingFabricIds, plan.FabricSiteId.ValueString())

		if !alreadyExists {
			existingFabricIds = append(existingFabricIds, plan.FabricSiteId.ValueString())
		}

		// Create object
		body := plan.toBody(ctx, VirtualNetworkToFabricSite{}, existingFabricIds)

		params = ""
		res, err = r.client.Put(plan.getPath()+params, body, cc.UseMutex)
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

		newFabricIds := []string{}
		for _, id := range res.Get("response.0.fabricIds").Array() {
			newFabricIds = append(newFabricIds, id.String())
		}
		slices.Sort(newFabricIds)
		slices.Sort(existingFabricIds)
		if slices.Equal(newFabricIds, existingFabricIds) {
			break
		}
		if try != MAX_RETRIES {
			tflog.Warn(ctx, fmt.Sprintf("%s: Assigned fabric ids does not match. Retrying for %d time", plan.Id.ValueString(), try))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Assigned fabric ids does not match. Please try again later, %s", res.String()))
			return
		}
	}

	compositeId := plan.FabricSiteId.ValueString() + "--" + plan.VirtualNetworkName.ValueString()
	plan.Id = basetypes.NewStringPointerValue(&compositeId)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VirtualNetworkToFabricSiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VirtualNetworkToFabricSite

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
	params += "&fabricId=" + url.QueryEscape(state.FabricSiteId.ValueString())
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	// Verify that the requested fabric site ID is in the response and set the composite ID
	fabricIds := res.Get("response.0.fabricIds").Array()
	found := false
	for _, id := range fabricIds {
		if id.String() == state.FabricSiteId.ValueString() {
			found = true
			break
		}
	}

	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Ensure the composite ID is set correctly
	compositeId := state.FabricSiteId.ValueString() + "--" + state.VirtualNetworkName.ValueString()
	state.Id = basetypes.NewStringPointerValue(&compositeId)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *VirtualNetworkToFabricSiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VirtualNetworkToFabricSite

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

func (r *VirtualNetworkToFabricSiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VirtualNetworkToFabricSite

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	MAX_RETRIES := 3
	for try := 1; try <= MAX_RETRIES; try++ {

		params := ""
		params += "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
		res, err := r.client.Get(state.getPath()+params, cc.UseMutex)
		if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		existingFabricIds := []string{}
		for _, id := range res.Get("response.0.fabricIds").Array() {
			existingFabricIds = append(existingFabricIds, id.String())
		}

		tflog.Debug(ctx, fmt.Sprintf("Existing fabric IDs before removal: %v", existingFabricIds))

		// Extract existing fabricIds list
		newFabricIds := []string{}
		for _, id := range existingFabricIds {
			if id != state.FabricSiteId.ValueString() {
				newFabricIds = append(newFabricIds, id)
			}
		}

		body := state.toBody(ctx, VirtualNetworkToFabricSite{}, newFabricIds)

		res, err = r.client.Put(state.getPath(), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
			return
		}

		params = "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
		res, err = r.client.Get(state.getPath()+params, cc.UseMutex)
		if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
			resp.State.RemoveResource(ctx)
			return
		}

		getFabricIds := []string{}
		for _, id := range res.Get("response.0.fabricIds").Array() {
			getFabricIds = append(getFabricIds, id.String())
		}
		slices.Sort(newFabricIds)
		slices.Sort(getFabricIds)
		if slices.Equal(newFabricIds, getFabricIds) {
			break
		}
		if try != MAX_RETRIES {
			tflog.Warn(ctx, fmt.Sprintf("Assigned fabric ids does not match. Retrying for %d time(s). Response: %s", try, res.String()))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Assigned fabric ids does not match. Please try again later, %s", res.String()))
			return
		}

	}
	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *VirtualNetworkToFabricSiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <virtual_network_name>,<fabric_site_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("virtual_network_name"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_site_id"), idParts[1])...)
}

// End of section. //template:end import
