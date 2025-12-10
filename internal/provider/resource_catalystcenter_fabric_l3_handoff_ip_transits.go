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
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &FabricL3HandoffIPTransitsResource{}
var _ resource.ResourceWithImportState = &FabricL3HandoffIPTransitsResource{}

func NewFabricL3HandoffIPTransitsResource() resource.Resource {
	return &FabricL3HandoffIPTransitsResource{}
}

type FabricL3HandoffIPTransitsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *FabricL3HandoffIPTransitsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l3_handoff_ip_transits"
}

func (r *FabricL3HandoffIPTransitsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages Layer 3 handoffs with IP transit on fabric device within a single resource, specifying a list of handoffs as input").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the fabric device").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric this device belongs to").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"l3_handoffs": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of layer 3 handoff ip transits").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the layer 3 handoff ip transit").String,
							Computed:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric this device belongs to").String,
							Required:            true,
						},
						"network_device_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the fabric device").String,
							Required:            true,
						},
						"transit_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the transit network of the layer 3 handoff ip transit").String,
							Required:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name of the layer 3 handoff ip transit").String,
							Optional:            true,
						},
						"external_connectivity_ip_pool_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer").String,
							Optional:            true,
						},
						"virtual_network_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the virtual network associated with this fabric site").String,
							Required:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094)").String,
							Required:            true,
						},
						"tcp_mss_adjustment": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6").AddIntegerRangeDescription(500, 1440).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(500, 1440),
							},
						},
						"local_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name").String,
							Optional:            true,
						},
						"remote_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name").String,
							Optional:            true,
						},
						"local_ipv6_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name").String,
							Optional:            true,
						},
						"remote_ipv6_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *FabricL3HandoffIPTransitsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *FabricL3HandoffIPTransitsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricL3HandoffIPTransits

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FabricL3HandoffIPTransits{})
	var planList []FabricL3HandoffIPTransits
	maxElementsPerShard := 20
	originalList := plan.L3Handoffs
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.L3Handoffs = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, FabricL3HandoffIPTransits{})
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
	plan.Id = types.StringValue(fmt.Sprint(plan.NetworkDeviceId.ValueString()))
	params = ""
	params += "?networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString()) + "&fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
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
func (r *FabricL3HandoffIPTransitsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricL3HandoffIPTransits

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?networkDeviceId=" + url.QueryEscape(state.NetworkDeviceId.ValueString()) + "&fabricId=" + url.QueryEscape(state.FabricId.ValueString())
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
func (r *FabricL3HandoffIPTransitsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricL3HandoffIPTransits

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
	var toDelete = FabricL3HandoffIPTransits{
		L3Handoffs: []FabricL3HandoffIPTransitsL3Handoffs{},
	}
	var toCreate = FabricL3HandoffIPTransits{
		L3Handoffs: []FabricL3HandoffIPTransitsL3Handoffs{},
	}
	var toUpdate = FabricL3HandoffIPTransits{
		L3Handoffs: []FabricL3HandoffIPTransitsL3Handoffs{},
	}
	var toReplace = FabricL3HandoffIPTransits{
		L3Handoffs: []FabricL3HandoffIPTransitsL3Handoffs{},
	}

	planMap := make(map[string]FabricL3HandoffIPTransitsL3Handoffs)
	stateMap := make(map[string]FabricL3HandoffIPTransitsL3Handoffs)

	// Populate state map
	for _, v := range state.L3Handoffs {
		stateMap[strconv.FormatInt(v.VlanId.ValueInt64(), 10)] = v
	}

	// Populate plan map
	for _, v := range plan.L3Handoffs {
		planMap[strconv.FormatInt(v.VlanId.ValueInt64(), 10)] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.L3Handoffs = append(toDelete.L3Handoffs, stateItem)
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
				if planItem.TransitNetworkId != stateItem.TransitNetworkId {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.InterfaceName != stateItem.InterfaceName {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.ExternalConnectivityIpPoolName != stateItem.ExternalConnectivityIpPoolName {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.VirtualNetworkName != stateItem.VirtualNetworkName {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.VlanId != stateItem.VlanId {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.LocalIpAddress != stateItem.LocalIpAddress {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.RemoteIpAddress != stateItem.RemoteIpAddress {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.LocalIpv6Address != stateItem.LocalIpv6Address {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				if planItem.RemoteIpv6Address != stateItem.RemoteIpv6Address {
					toReplace.L3Handoffs = append(toReplace.L3Handoffs, planItem)
					continue
				}
				toUpdate.L3Handoffs = append(toUpdate.L3Handoffs, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.L3Handoffs = append(toCreate.L3Handoffs, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.L3Handoffs) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.L3Handoffs)))
		// Clear IDs before recreating
		var toReplaceNoId = FabricL3HandoffIPTransits{
			L3Handoffs: []FabricL3HandoffIPTransitsL3Handoffs{},
		}
		for _, item := range toReplace.L3Handoffs {
			item.Id = types.StringNull()
			toReplaceNoId.L3Handoffs = append(toReplaceNoId.L3Handoffs, item)
		}

		// Replace is done by delete + create
		toDelete.L3Handoffs = append(toDelete.L3Handoffs, toReplace.L3Handoffs...)
		toCreate.L3Handoffs = append(toCreate.L3Handoffs, toReplaceNoId.L3Handoffs...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.L3Handoffs) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.L3Handoffs)))
		for _, v := range toDelete.L3Handoffs {
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
	if len(toCreate.L3Handoffs) > 0 {

		maxElementsPerShard := 20
		var createList []FabricL3HandoffIPTransits
		for i := 0; i < len(toCreate.L3Handoffs); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.L3Handoffs))
			chunk := toCreate.L3Handoffs[i:end]
			currentPlanForShard := plan
			currentPlanForShard.L3Handoffs = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.L3Handoffs)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, FabricL3HandoffIPTransits{}) // Convert to request body
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
		params += "?networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString()) + "&fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
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
	if len(toUpdate.L3Handoffs) > 0 {

		maxElementsPerShard := 20
		var updateList []FabricL3HandoffIPTransits
		for i := 0; i < len(toUpdate.L3Handoffs); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toUpdate.L3Handoffs))
			chunk := toUpdate.L3Handoffs[i:end]
			currentPlanForShard := plan
			currentPlanForShard.L3Handoffs = chunk
			updateList = append(updateList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.L3Handoffs)))
		planIndexMap := make(map[string]int)

		for _, pl := range updateList {
			for i, v := range plan.L3Handoffs {
				planIndexMap[strconv.FormatInt(v.VlanId.ValueInt64(), 10)] = i
			}
			for _, item := range pl.L3Handoffs {
				toUpdateKey := strconv.FormatInt(item.VlanId.ValueInt64(), 10)
				if updatedItem, exists := planMap[toUpdateKey]; exists {
					if index, found := planIndexMap[toUpdateKey]; found {
						plan.L3Handoffs[index] = updatedItem
					}
				}
			}

			body := pl.toBody(ctx, FabricL3HandoffIPTransits{})
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
func (r *FabricL3HandoffIPTransitsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricL3HandoffIPTransits

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	params := "?networkDeviceId=" + url.QueryEscape(state.NetworkDeviceId.ValueString()) + "&fabricId=" + url.QueryEscape(state.FabricId.ValueString())
	res, err := r.client.Delete(state.getPath()+params, cc.UseMutex)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
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
func (r *FabricL3HandoffIPTransitsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_device_id>,<fabric_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_device_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[1])...)
}

// End of section. //template:end import
