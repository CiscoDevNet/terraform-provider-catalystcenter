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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
var _ resource.Resource = &AnycastGatewaysResource{}
var _ resource.ResourceWithImportState = &AnycastGatewaysResource{}

func NewAnycastGatewaysResource() resource.Resource {
	return &AnycastGatewaysResource{}
}

type AnycastGatewaysResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *AnycastGatewaysResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_anycast_gateways"
}

func (r *AnycastGatewaysResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages Anycast Gateways within a single resource, specifying a list of anycast gateways as input").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric to contain this anycast gateway").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"anycast_gateways": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of anycast gateways").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the anycast gateway").String,
							Computed:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric to contain this anycast gateway").String,
							Required:            true,
						},
						"virtual_network_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it").String,
							Required:            true,
						},
						"ip_pool_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the IP pool associated with the anycast gateway").String,
							Required:            true,
						},
						"tcp_mss_adjustment": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("TCP maximum segment size adjustment").AddIntegerRangeDescription(500, 1440).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(500, 1440),
							},
						},
						"vlan_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the VLAN of the anycast gateway").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the VLAN of the anycast gateway. allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, 2046, and 4094. if deploying an anycast gateway on a fabric zone, this vlanId must match the vlanId of the corresponding anycast gateway on the fabric site").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"traffic_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of traffic the anycast gateway serves").AddStringEnumDescription("DATA", "VOICE").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("DATA", "VOICE"),
							},
						},
						"pool_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The pool type of the anycast gateway (required for & applicable only to INFRA_VN)").AddStringEnumDescription("EXTENDED_NODE", "FABRIC_AP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("EXTENDED_NODE", "FABRIC_AP"),
							},
						},
						"security_group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the associated Security Group (not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"critical_pool": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable critical VLAN. if true, autoGenerateVlanName must also be true. (isCriticalPool is not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"l2_flooding_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable layer 2 flooding (not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"wireless_pool": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable fabric-enabled wireless (not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"ip_directed_broadcast": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable IP-directed broadcast (not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"intra_subnet_routing_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"multiple_ip_to_mac_addresses": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN)").String,
							Optional:            true,
						},
						"supplicant_based_extended_node_onboarding": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN)").String,
							Optional:            true,
						},
						"group_based_policy_enforcement_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN; defaults to false)").String,
							Optional:            true,
						},
						"auto_generate_vlan_name": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This field cannot be true when vlanName is provided. the vlanName will be generated as ipPoolGroupV4Cidr-virtualNetworkName for non-critical VLANs. for critical VLANs with DATA trafficType, vlanName will be CRITICAL_VLAN. for critical VLANs with VOICE trafficType, vlanName will be VOICE_VLAN").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AnycastGatewaysResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AnycastGatewaysResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AnycastGateways

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AnycastGateways{})
	var planList []AnycastGateways
	maxElementsPerShard := 20
	originalList := plan.AnycastGateways
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.AnycastGateways = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, AnycastGateways{})
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
func (r *AnycastGatewaysResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AnycastGateways

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
func (r *AnycastGatewaysResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AnycastGateways

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
	var toDelete = AnycastGateways{
		AnycastGateways: []AnycastGatewaysAnycastGateways{},
	}
	var toCreate = AnycastGateways{
		AnycastGateways: []AnycastGatewaysAnycastGateways{},
	}
	var toUpdate = AnycastGateways{
		AnycastGateways: []AnycastGatewaysAnycastGateways{},
	}
	var toReplace = AnycastGateways{
		AnycastGateways: []AnycastGatewaysAnycastGateways{},
	}

	planMap := make(map[string]AnycastGatewaysAnycastGateways)
	stateMap := make(map[string]AnycastGatewaysAnycastGateways)

	// Populate state map
	for _, v := range state.AnycastGateways {
		stateMap[v.IpPoolName.ValueString()] = v
	}

	// Populate plan map
	for _, v := range plan.AnycastGateways {
		planMap[v.IpPoolName.ValueString()] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.AnycastGateways = append(toDelete.AnycastGateways, stateItem)
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
					toReplace.AnycastGateways = append(toReplace.AnycastGateways, planItem)
					continue
				}
				if planItem.IpPoolName != stateItem.IpPoolName {
					toReplace.AnycastGateways = append(toReplace.AnycastGateways, planItem)
					continue
				}
				if planItem.CriticalPool != stateItem.CriticalPool {
					toReplace.AnycastGateways = append(toReplace.AnycastGateways, planItem)
					continue
				}
				if planItem.IntraSubnetRoutingEnabled != stateItem.IntraSubnetRoutingEnabled {
					toReplace.AnycastGateways = append(toReplace.AnycastGateways, planItem)
					continue
				}
				toUpdate.AnycastGateways = append(toUpdate.AnycastGateways, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.AnycastGateways = append(toCreate.AnycastGateways, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.AnycastGateways) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.AnycastGateways)))
		// Clear IDs before recreating
		var toReplaceNoId = AnycastGateways{
			AnycastGateways: []AnycastGatewaysAnycastGateways{},
		}
		for _, item := range toReplace.AnycastGateways {
			item.Id = types.StringNull()

			if !item.AutoGenerateVlanName.IsNull() && item.AutoGenerateVlanName.ValueBool() {
				item.VlanName = types.StringNull()
			}
			item.VlanId = types.Int64Null()
			toReplaceNoId.AnycastGateways = append(toReplaceNoId.AnycastGateways, item)
		}

		// Replace is done by delete + create
		toDelete.AnycastGateways = append(toDelete.AnycastGateways, toReplace.AnycastGateways...)
		toCreate.AnycastGateways = append(toCreate.AnycastGateways, toReplaceNoId.AnycastGateways...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.AnycastGateways) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.AnycastGateways)))
		for _, v := range toDelete.AnycastGateways {
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
	if len(toCreate.AnycastGateways) > 0 {

		maxElementsPerShard := 20
		var createList []AnycastGateways
		for i := 0; i < len(toCreate.AnycastGateways); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.AnycastGateways))
			chunk := toCreate.AnycastGateways[i:end]
			currentPlanForShard := plan
			currentPlanForShard.AnycastGateways = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.AnycastGateways)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, AnycastGateways{}) // Convert to request body
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
	if len(toUpdate.AnycastGateways) > 0 {

		maxElementsPerShard := 20
		var updateList []AnycastGateways
		for i := 0; i < len(toUpdate.AnycastGateways); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toUpdate.AnycastGateways))
			chunk := toUpdate.AnycastGateways[i:end]
			currentPlanForShard := plan
			currentPlanForShard.AnycastGateways = chunk
			updateList = append(updateList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.AnycastGateways)))
		planIndexMap := make(map[string]int)

		var getState AnycastGateways
		getParams := ""
		getParams += "?fabricId=" + url.QueryEscape(state.FabricId.ValueString())
		getRes, err := r.client.Get(state.getPath() + getParams)
		if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, getRes.String()))
			return
		}

		if getState.isNull(ctx, getRes) {
			getState.fromBody(ctx, getRes)
		} else {
			getState.updateFromBody(ctx, getRes)
		}

		existingNoPut := make(map[string]types.Int64)
		for _, item := range getState.AnycastGateways {
			updateKey := item.IpPoolName.ValueString()
			existingNoPut[updateKey] = item.VlanId
		}

		for _, pl := range updateList {
			for i, v := range plan.AnycastGateways {
				planIndexMap[v.IpPoolName.ValueString()] = i
			}
			for itemInd, item := range pl.AnycastGateways {
				toUpdateKey := item.IpPoolName.ValueString()
				if updatedItem, exists := planMap[toUpdateKey]; exists {
					if index, found := planIndexMap[toUpdateKey]; found {
						pl.AnycastGateways[itemInd].VlanId = existingNoPut[toUpdateKey]
						plan.AnycastGateways[index] = updatedItem
					}
				}
			}

			body := pl.toBody(ctx, AnycastGateways{})
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
func (r *AnycastGatewaysResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AnycastGateways

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	for _, v := range state.AnycastGateways {
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
func (r *AnycastGatewaysResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
