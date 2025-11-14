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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &FabricMulticastVirtualNetworksResource{}
var _ resource.ResourceWithImportState = &FabricMulticastVirtualNetworksResource{}

func NewFabricMulticastVirtualNetworksResource() resource.Resource {
	return &FabricMulticastVirtualNetworksResource{}
}

type FabricMulticastVirtualNetworksResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *FabricMulticastVirtualNetworksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_multicast_virtual_networks"
}

func (r *FabricMulticastVirtualNetworksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages Fabric Site VN Multicast").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric site this multicast configuration is associated with").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"virtual_networks": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of multicast virtual networks").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the multicast configuration").String,
							Computed:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric site this multicast configuration is associated with").String,
							Required:            true,
						},
						"virtual_network_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the virtual network associated with the fabric site").String,
							Required:            true,
						},
						"ip_pool_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the IP Pool associated with the fabric site").String,
							Required:            true,
						},
						"ipv4_ssm_ranges": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Source Specific Multicast (SSM) ranges").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"multicast_rps": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Multicast Rendezvous Points (RP). Required for Any Source Multicast (ASM) scenario").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"rp_device_location": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Device location of the RP").AddStringEnumDescription("EXTERNAL", "FABRIC").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("EXTERNAL", "FABRIC"),
										},
									},
									"ipv4_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.").String,
										Optional:            true,
									},
									"ipv6_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.").String,
										Optional:            true,
									},
									"is_default_v4_rp": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Specifies whether it is a default IPv4 RP").String,
										Optional:            true,
									},
									"is_default_v6_rp": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Specifies whether it is a default IPv6 RP").String,
										Optional:            true,
									},
									"network_device_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IDs of the network devices. This is a required field for fabric RPs").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"ipv4_asm_ranges": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"ipv6_asm_ranges": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *FabricMulticastVirtualNetworksResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *FabricMulticastVirtualNetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricMulticastVirtualNetworks

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FabricMulticastVirtualNetworks{})
	var planList []FabricMulticastVirtualNetworks
	maxElementsPerShard := 20
	originalList := plan.VirtualNetworks
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.VirtualNetworks = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, FabricMulticastVirtualNetworks{})
		res, err = r.client.Post(plan.getPath()+params, body, cc.UseMutex)
		if err != nil {
			errorCode := res.Get("response.errorCode").String()
			if errorCode == "NCDP10000" {
				// Log a warning and continue execution when device is unreachable
				failureReason := res.Get("response.failureReason").String()
				resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
			} else {
				resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
				break
			}
		}
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.FabricId.ValueString()))
	params = ""
	params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
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
func (r *FabricMulticastVirtualNetworksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricMulticastVirtualNetworks

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(state.FabricId.ValueString())
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *FabricMulticastVirtualNetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricMulticastVirtualNetworks

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

	// Initialize toDelete, toCreate, toReplace, and toUpdate with empty slices
	var toDelete = FabricMulticastVirtualNetworks{
		VirtualNetworks: []FabricMulticastVirtualNetworksVirtualNetworks{},
	}
	var toCreate = FabricMulticastVirtualNetworks{
		VirtualNetworks: []FabricMulticastVirtualNetworksVirtualNetworks{},
	}
	var toUpdate = FabricMulticastVirtualNetworks{
		VirtualNetworks: []FabricMulticastVirtualNetworksVirtualNetworks{},
	}
	var toReplace = FabricMulticastVirtualNetworks{
		VirtualNetworks: []FabricMulticastVirtualNetworksVirtualNetworks{},
	}

	planMap := make(map[string]FabricMulticastVirtualNetworksVirtualNetworks)
	stateMap := make(map[string]FabricMulticastVirtualNetworksVirtualNetworks)

	// Populate state map
	for _, v := range state.VirtualNetworks {
		stateMap[v.VirtualNetworkName.ValueString()] = v
	}

	// Populate plan map
	for _, v := range plan.VirtualNetworks {
		planMap[v.VirtualNetworkName.ValueString()] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.VirtualNetworks = append(toDelete.VirtualNetworks, stateItem)
		}
	}

	// Find items to create update and replace
	for planKey, planItem := range planMap {
		if stateItem, exists := stateMap[planKey]; exists {
			// Exists in both, check if different
			if !reflect.DeepEqual(planItem, stateItem) {
				// Update planItem but ensure ID comes from stateItem
				planItem.Id = stateItem.Id
				planMap[planKey] = planItem // Store back in planMap
				// Check if any field marked as requires_replace differs
				if planItem.VirtualNetworkName != stateItem.VirtualNetworkName {
					toReplace.VirtualNetworks = append(toReplace.VirtualNetworks, planItem)
					continue
				}
				if planItem.IpPoolName != stateItem.IpPoolName {
					toReplace.VirtualNetworks = append(toReplace.VirtualNetworks, planItem)
					continue
				}
				toUpdate.VirtualNetworks = append(toUpdate.VirtualNetworks, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.VirtualNetworks = append(toCreate.VirtualNetworks, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.VirtualNetworks) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.VirtualNetworks)))
		// Clear IDs before recreating
		var toReplaceNoId = FabricMulticastVirtualNetworks{
			VirtualNetworks: []FabricMulticastVirtualNetworksVirtualNetworks{},
		}
		for _, item := range toReplace.VirtualNetworks {
			item.Id = types.StringNull()
			toReplaceNoId.VirtualNetworks = append(toReplaceNoId.VirtualNetworks, item)
		}

		// Replace is done by delete + create
		toDelete.VirtualNetworks = append(toDelete.VirtualNetworks, toReplace.VirtualNetworks...)
		toCreate.VirtualNetworks = append(toCreate.VirtualNetworks, toReplaceNoId.VirtualNetworks...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.VirtualNetworks) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.VirtualNetworks)))
		for _, v := range toDelete.VirtualNetworks {
			if v.Id.IsNull() {
				continue // Skip if id is null
			}
			res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(v.Id.ValueString()), cc.UseMutex)
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
	if len(toCreate.VirtualNetworks) > 0 {

		maxElementsPerShard := 20
		var createList []FabricMulticastVirtualNetworks
		for i := 0; i < len(toCreate.VirtualNetworks); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.VirtualNetworks))
			chunk := toCreate.VirtualNetworks[i:end]
			currentPlanForShard := plan
			currentPlanForShard.VirtualNetworks = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.VirtualNetworks)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, FabricMulticastVirtualNetworks{}) // Convert to request body
			res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
			if err != nil {
				errorCode := res.Get("response.errorCode").String()
				if errorCode == "NCDP10000" {
					// Log a warning and continue execution when device is unreachable
					failureReason := res.Get("response.failureReason").String()
					resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
				} else {
					resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
					break
				}
			}
		}
		params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
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
	if len(toUpdate.VirtualNetworks) > 0 {

		maxElementsPerShard := 20
		var updateList []FabricMulticastVirtualNetworks
		for i := 0; i < len(toUpdate.VirtualNetworks); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toUpdate.VirtualNetworks))
			chunk := toUpdate.VirtualNetworks[i:end]
			currentPlanForShard := plan
			currentPlanForShard.VirtualNetworks = chunk
			updateList = append(updateList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.VirtualNetworks)))
		planIndexMap := make(map[string]int)

		for _, pl := range updateList {
			for i, v := range plan.VirtualNetworks {
				planIndexMap[v.VirtualNetworkName.ValueString()] = i
			}
			for _, item := range pl.VirtualNetworks {
				toUpdateKey := item.VirtualNetworkName.ValueString()
				if updatedItem, exists := planMap[toUpdateKey]; exists {
					if index, found := planIndexMap[toUpdateKey]; found {
						plan.VirtualNetworks[index] = updatedItem
					}
				}
			}

			body := pl.toBody(ctx, FabricMulticastVirtualNetworks{})
			params := ""
			res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				break
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *FabricMulticastVirtualNetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricMulticastVirtualNetworks

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	for _, v := range state.VirtualNetworks {
		res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(v.Id.ValueString()), cc.UseMutex)
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *FabricMulticastVirtualNetworksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <fabric_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import
