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
		},
	}
}

func (r *FabricL3VirtualNetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
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
		// Perform GET request to retrieve the ID
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

		// Perform PUT request with the retrieved ID
		body := plan.toBody(ctx, FabricL3VirtualNetwork{})
		res, err = r.client.Put(plan.getPath(), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
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

// Section below is generated&owned by "gen/generator.go". //template:begin update
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

	body := plan.toBody(ctx, state)
	params := ""
	res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

func (r *FabricL3VirtualNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricL3VirtualNetwork

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	if state.VirtualNetworkName.ValueString() != "INFRA_VN" && state.VirtualNetworkName.ValueString() != "DEFAULT_VN" {
		params := ""
		params += "?virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
		res, err := r.client.Delete(state.getPath()+params, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		// For INFRA_VN and DEFAULT_VN, we do not delete them, we just remove fabric site associations
		// Perform PUT request with empty fabricIds
		state.FabricIds = types.SetNull(types.StringType)
		body := state.toBody(ctx, FabricL3VirtualNetwork{})
		res, err := r.client.Put(state.getPath(), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
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
