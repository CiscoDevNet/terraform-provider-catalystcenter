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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WirelessEnterpriseSSIDResource{}
var _ resource.ResourceWithImportState = &WirelessEnterpriseSSIDResource{}

func NewWirelessEnterpriseSSIDResource() resource.Resource {
	return &WirelessEnterpriseSSIDResource{}
}

type WirelessEnterpriseSSIDResource struct {
	client *cc.Client
}

func (r *WirelessEnterpriseSSIDResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_enterprise_ssid"
}

func (r *WirelessEnterpriseSSIDResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Wireless Enterprise SSID.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SSID Name").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"security_level": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Security Level").AddStringEnumDescription("wpa2_enterprise", "wpa2_personal", "open", "wpa3_enterprise", "wpa3_personal", "wpa2_wpa3_personal", "wpa2_wpa3_enterprise").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("wpa2_enterprise", "wpa2_personal", "open", "wpa3_enterprise", "wpa3_personal", "wpa2_wpa3_personal", "wpa2_wpa3_enterprise"),
				},
			},
			"passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Passphrase").String,
				Optional:            true,
			},
			"enable_fast_lane": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable FastLane").String,
				Optional:            true,
			},
			"enable_mac_filtering": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable MAC Filtering").String,
				Optional:            true,
			},
			"traffic_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Traffic Type").AddStringEnumDescription("voicedata", "data").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("voicedata", "data"),
				},
			},
			"radio_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio Policy").AddStringEnumDescription("Triple band operation(2.4GHz, 5GHz and 6GHz)", "Triple band operation with band select", "5GHz only", "2.4GHz only", "6GHz only").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Triple band operation(2.4GHz, 5GHz and 6GHz)", "Triple band operation with band select", "5GHz only", "2.4GHz only", "6GHz only"),
				},
			},
			"enable_broadcast_ssid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Broadcast SSID").String,
				Optional:            true,
			},
			"fast_transition": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fast Transition").AddStringEnumDescription("Adaptive", "Enable", "Disable").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Adaptive", "Enable", "Disable"),
				},
			},
			"enable_session_time_out": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Session Timeout").String,
				Optional:            true,
			},
			"session_time_out": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Session Time Out").String,
				Optional:            true,
			},
			"enable_client_exclusion": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Client Exclusion").String,
				Optional:            true,
			},
			"client_exclusion_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client Exclusion Timeout").String,
				Optional:            true,
			},
			"enable_basic_service_set_max_idle": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Basic Service Set Max Idle").String,
				Optional:            true,
			},
			"basic_service_set_client_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Basic Service Set Client Idle Timeout").String,
				Optional:            true,
			},
			"enable_directed_multicast_service": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Directed Multicast Service").String,
				Optional:            true,
			},
			"enable_neighbor_list": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Neighbor List").String,
				Optional:            true,
			},
			"mfp_client_protection": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mfp Client Protection").AddStringEnumDescription("Optional", "Disabled", "Required").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Optional", "Disabled", "Required"),
				},
			},
			"nas_options": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Nas Options").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Profile Name").String,
				Optional:            true,
			},
			"policy_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy Profile Name").String,
				Optional:            true,
			},
			"aaa_override": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AAA Override").String,
				Optional:            true,
			},
			"coverage_hole_detection_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection Enable").String,
				Optional:            true,
			},
			"protected_management_frame": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protected Management Frame").AddStringEnumDescription("Optional", "Disabled", "Required").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Optional", "Disabled", "Required"),
				},
			},
			"multi_psk_settings": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Multi PSK Settings (Only applicable for SSID with PERSONAL auth type and PSK)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"priority": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority").String,
							Optional:            true,
						},
						"passphrase_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Passphrase Type").AddStringEnumDescription("ASCII", "HEX").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ASCII", "HEX"),
							},
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Passphrase").String,
							Optional:            true,
						},
					},
				},
			},
			"client_rate_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client Rate Limit (in bits per second)").String,
				Optional:            true,
			},
			"auth_key_mgmt": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"rsn_cipher_suite_gcmp256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rsn Cipher Suite Gcmp256").String,
				Optional:            true,
			},
			"rsn_cipher_suite_ccmp256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rsn Cipher Suite Ccmp256").String,
				Optional:            true,
			},
			"rsn_cipher_suite_gcmp128": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rsn Cipher Suite Gcmp 128").String,
				Optional:            true,
			},
			"ghz6_policy_client_steering": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ghz6 Policy Client Steering").String,
				Optional:            true,
			},
			"ghz24_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ghz24 Policy").AddStringEnumDescription("dot11-g-only", "dot11-bg-only").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dot11-g-only", "dot11-bg-only"),
				},
			},
		},
	}
}

func (r *WirelessEnterpriseSSIDResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *WirelessEnterpriseSSIDResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessEnterpriseSSID

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessEnterpriseSSID{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		if !globalAllowExistingOnCreate {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			return
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Placeholder for updating existing resource"))
		}
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.Name.ValueString()))

	if !globalAllowExistingOnCreate {
		tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully, but allow_existing_on_create is set to true, so the existing resource was updated instead of created", plan.Id.ValueString()))
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *WirelessEnterpriseSSIDResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessEnterpriseSSID

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?ssidName=" + url.QueryEscape(state.Id.ValueString())
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
func (r *WirelessEnterpriseSSIDResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessEnterpriseSSID

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *WirelessEnterpriseSSIDResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessEnterpriseSSID

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *WirelessEnterpriseSSIDResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import
