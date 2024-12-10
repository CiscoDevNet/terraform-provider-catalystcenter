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
var _ resource.Resource = &FabricPortAssignmentResource{}

func NewFabricPortAssignmentResource() resource.Resource {
	return &FabricPortAssignmentResource{}
}

type FabricPortAssignmentResource struct {
	client *cc.Client
}

func (r *FabricPortAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_port_assignment"
}

func (r *FabricPortAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Optional:            true,
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
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
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

func (r *FabricPortAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (r *FabricPortAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricPortAssignment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FabricPortAssignment{})

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
	params = ""
	params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// Map response to portAssignments
	responseArray := res.Get("response").Array()
	portAssignments := make([]FabricPortAssignmentPortAssignments, len(plan.PortAssignments))
	copy(portAssignments, plan.PortAssignments) // Make a copy to update values
	for i, port := range portAssignments {
		for _, respPort := range responseArray {
			if respPort.Get("interfaceName").String() == port.InterfaceName.ValueString() {
				port.Id = types.StringValue(respPort.Get("id").String())
				break
			}
		}
		portAssignments[i] = port
	}

	// Update the plan with mapped portAssignments
	plan.PortAssignments = portAssignments

	plan.Id = types.StringValue(fmt.Sprint(plan.NetworkDeviceId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *FabricPortAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricPortAssignment

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

func (r *FabricPortAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricPortAssignment

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

	// Identify removed interfaces
	stateAssignments := make(map[string]FabricPortAssignmentPortAssignments)
	for _, sa := range state.PortAssignments {
		stateAssignments[sa.InterfaceName.ValueString()] = sa
	}

	planAssignments := make(map[string]FabricPortAssignmentPortAssignments)
	for _, pa := range plan.PortAssignments {
		planAssignments[pa.InterfaceName.ValueString()] = pa
	}

	// Find interfaces to delete
	for interfaceName, stateAssignment := range stateAssignments {
		if _, exists := planAssignments[interfaceName]; !exists {
			// Interface exists in state but not in plan, delete it
			_, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(stateAssignment.Id.ValueString()))
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete port assignment for interface %s, got error: %s", interfaceName, err))
			} else {
				tflog.Debug(ctx, fmt.Sprintf("Deleted port assignment for interface: %s", interfaceName))
			}
		}
	}

	// Find interfaces to create
	for interfaceName, planAssignment := range planAssignments {
		if _, exists := stateAssignments[interfaceName]; !exists {
			// Create a temporary FabricPortAssignment for the current interface
			singlePortAssignment := FabricPortAssignment{
				FabricId:        plan.FabricId,
				NetworkDeviceId: plan.NetworkDeviceId,
				PortAssignments: []FabricPortAssignmentPortAssignments{planAssignment}, // Filtered for the current interface
			}

			// Generate the body using the existing toBody function
			body := singlePortAssignment.toBody(ctx, state)

			// Send the body as the POST request
			_, err := r.client.Post(state.getPath(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create port assignment for interface %s, got error: %s", interfaceName, err))
			} else {
				tflog.Debug(ctx, fmt.Sprintf("Created port assignment for interface: %s", interfaceName))

				// Get the existing ID for the given interface
				params := "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString())

				res, err := r.client.Get(plan.getPath() + params)
				if err != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
					return
				}

				// Log the full response for debugging purposes
				tflog.Debug(ctx, fmt.Sprintf("API response: %s", res.String()))

				// Check if the ID is present in the response
				planAssignment.Id = types.StringValue(res.Get("response.#(interfaceName==\"" + planAssignment.InterfaceName.ValueString() + "\").id").String())

				// Update the assignment in the plan
				for i := range plan.PortAssignments {
					if plan.PortAssignments[i].InterfaceName.ValueString() == planAssignment.InterfaceName.ValueString() {
						plan.PortAssignments[i] = planAssignment
						break
					}
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Plan Port Assignments: %+v", plan.PortAssignments))

	body := plan.toBody(ctx, state)
	params := ""
	res, err := r.client.Put(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *FabricPortAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricPortAssignment

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
// End of section. //template:end import
