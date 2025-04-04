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
	"reflect"
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &FabricPortAssignmentsResource{}
var _ resource.ResourceWithImportState = &FabricPortAssignmentsResource{}

func NewFabricPortAssignmentsResource() resource.Resource {
	return &FabricPortAssignmentsResource{}
}

type FabricPortAssignmentsResource struct {
	client *cc.Client
}

func (r *FabricPortAssignmentsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_port_assignments"
}

func (r *FabricPortAssignmentsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages port assignments in SD-Access fabric.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric the device is assigned to").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the port assignment").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"port_assignments": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of port assignments in SD-Access fabric").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the port assignment").String,
							Computed:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric the device is assigned to").String,
							Required:            true,
						},
						"network_device_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the port assignment").String,
							Required:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name of the port assignment").String,
							Required:            true,
						},
						"connected_device_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Connected device type of the port assignment").AddStringEnumDescription("USER_DEVICE", "ACCESS_POINT", "TRUNKING_DEVICE", "AUTHENTICATOR_SWITCH", "SUPPLICANT_BASED_EXTENDED_NODE").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("USER_DEVICE", "ACCESS_POINT", "TRUNKING_DEVICE", "AUTHENTICATOR_SWITCH", "SUPPLICANT_BASED_EXTENDED_NODE"),
							},
						},
						"data_vlan_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Data VLAN name of the port assignment").String,
							Optional:            true,
						},
						"voice_vlan_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Voice VLAN name of the port assignment").String,
							Optional:            true,
						},
						"authenticate_template_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authenticate template name of the port assignment").AddStringEnumDescription("No Authentication", "Open Authentication", "Closed Authentication", "Low Impact").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("No Authentication", "Open Authentication", "Closed Authentication", "Low Impact"),
							},
						},
						"security_group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Security group name of the port assignment").String,
							Optional:            true,
						},
						"interface_description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface description of the port assignment").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *FabricPortAssignmentsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *FabricPortAssignmentsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricPortAssignments

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FabricPortAssignments{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCDP10000" {
			// Log a warning and continue execution when device is unreachable
			failureReason := res.Get("response.failureReason").String()
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			return
		}
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.NetworkDeviceId.ValueString()))
	params = ""
	params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *FabricPortAssignmentsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricPortAssignments

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(state.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(state.NetworkDeviceId.ValueString())
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

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *FabricPortAssignmentsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricPortAssignments

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

	// Initialize toDelete, toCreate, and toUpdate with empty slices
	var toDelete = FabricPortAssignments{
		PortAssignments: []FabricPortAssignmentsPortAssignments{},
	}
	var toCreate = FabricPortAssignments{
		PortAssignments: []FabricPortAssignmentsPortAssignments{},
	}
	var toUpdate = FabricPortAssignments{
		PortAssignments: []FabricPortAssignmentsPortAssignments{},
	}

	planMap := make(map[string]FabricPortAssignmentsPortAssignments)
	stateMap := make(map[string]FabricPortAssignmentsPortAssignments)

	// Populate state map
	for _, v := range state.PortAssignments {
		stateMap[v.InterfaceName.ValueString()] = v
	}

	// Populate plan map
	for _, v := range plan.PortAssignments {
		planMap[v.InterfaceName.ValueString()] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.PortAssignments = append(toDelete.PortAssignments, stateItem)
		}
	}

	// Find items to create and update
	for planKey, planItem := range planMap {
		if stateItem, exists := stateMap[planKey]; exists {
			// Exists in both, check if different
			if !reflect.DeepEqual(planItem, stateItem) {
				// Update planItem but ensure ID comes from stateItem
				planItem.Id = stateItem.Id
				planMap[planKey] = planItem // Store back in planMap
				toUpdate.PortAssignments = append(toUpdate.PortAssignments, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.PortAssignments = append(toCreate.PortAssignments, planItem)
		}
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.PortAssignments) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.PortAssignments)))
		for _, v := range toDelete.PortAssignments {
			res, err := r.client.Delete(plan.getPath() + "/" + url.QueryEscape(v.Id.ValueString()))
			if err != nil {
				errorCode := res.Get("response.errorCode").String()
				if errorCode == "NCDP10000" {
					// Log a warning and continue execution when device is unreachable
					failureReason := res.Get("response.failureReason").String()
					resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
				} else {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "DELETE", err, res.String()))
					return
				}
			}
		}
	}

	// CREATE
	// If there are objects marked for create
	if len(toCreate.PortAssignments) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.PortAssignments)))
		body := toCreate.toBody(ctx, FabricPortAssignments{}) // Convert to request body
		params := ""
		res, err := r.client.Post(plan.getPath()+params, body)
		if err != nil {
			errorCode := res.Get("response.errorCode").String()
			if errorCode == "NCDP10000" {
				// Log a warning and continue execution when device is unreachable
				failureReason := res.Get("response.failureReason").String()
				resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
				return
			}
		}
		params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString())
		res, err = r.client.Get(plan.getPath() + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		// Populate missing IDs using fromBodyUnknowns
		plan.fromBodyUnknowns(ctx, res)
	}

	// UPDATE
	// Update objects (objects that have different definition in plan and state)
	if len(toUpdate.PortAssignments) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.PortAssignments)))
		planIndexMap := make(map[string]int)
		for i, v := range plan.PortAssignments {
			planIndexMap[v.InterfaceName.ValueString()] = i
		}
		for _, item := range toUpdate.PortAssignments {
			toUpdateKey := item.InterfaceName.ValueString()
			if updatedItem, exists := planMap[toUpdateKey]; exists {
				if index, found := planIndexMap[toUpdateKey]; found {
					plan.PortAssignments[index] = updatedItem
				}
			}
		}

		body := toUpdate.toBody(ctx, FabricPortAssignments{})
		params := ""
		res, err := r.client.Put(plan.getPath()+params, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *FabricPortAssignmentsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricPortAssignments

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	params := "?fabricId=" + url.QueryEscape(state.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(state.NetworkDeviceId.ValueString())
	res, err := r.client.Delete(state.getPath() + params)
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCDP10000" {
			// Log a warning and continue execution when device is unreachable
			failureReason := res.Get("response.failureReason").String()
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "DELETE", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *FabricPortAssignmentsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <fabric_id>,<network_device_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_device_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

// End of section. //template:end import
