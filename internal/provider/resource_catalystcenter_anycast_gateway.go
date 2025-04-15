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
	"sync"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
var _ resource.Resource = &AnycastGatewayResource{}
var _ resource.ResourceWithImportState = &AnycastGatewayResource{}

func NewAnycastGatewayResource() resource.Resource {
	return &AnycastGatewayResource{}
}

type AnycastGatewayResource struct {
	client              *cc.Client
	anycastGatewayMutex *sync.Mutex
}

func (r *AnycastGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_anycast_gateway"
}

func (r *AnycastGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages Anycast Gateways").String,

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
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ip_pool_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the IP pool associated with the anycast gateway").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
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
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
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
	}
}

func (r *AnycastGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.anycastGatewayMutex = req.ProviderData.(*CcProviderData).AnycastGatewayMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AnycastGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AnycastGateway

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	r.anycastGatewayMutex.Lock()
	defer r.anycastGatewayMutex.Unlock()

	// Create object
	body := plan.toBody(ctx, AnycastGateway{})

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
	params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString()) + "&virtualNetworkName=" + url.QueryEscape(plan.VirtualNetworkName.ValueString()) + "&ipPoolName=" + url.QueryEscape(plan.IpPoolName.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.0.id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AnycastGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AnycastGateway

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(state.FabricId.ValueString()) + "&virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString()) + "&ipPoolName=" + url.QueryEscape(state.IpPoolName.ValueString())
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
func (r *AnycastGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AnycastGateway

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
	r.anycastGatewayMutex.Lock()
	defer r.anycastGatewayMutex.Unlock()

	body := plan.toBody(ctx, state)
	params := ""
	res, err := r.client.Put(plan.getPath()+params, body)
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCDP10000" {
			// Log a warning and continue execution when device is unreachable
			failureReason := res.Get("response.failureReason").String()
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (%s), got error: %s, %s", "PUT", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *AnycastGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AnycastGateway

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	r.anycastGatewayMutex.Lock()
	defer r.anycastGatewayMutex.Unlock()
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
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
func (r *AnycastGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <fabric_id>,<virtual_network_name>,<ip_pool_name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("virtual_network_name"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ip_pool_name"), idParts[2])...)
}

// End of section. //template:end import
