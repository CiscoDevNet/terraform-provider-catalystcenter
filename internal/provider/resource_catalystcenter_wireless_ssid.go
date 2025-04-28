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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &WirelessSSIDResource{}
var _ resource.ResourceWithImportState = &WirelessSSIDResource{}

func NewWirelessSSIDResource() resource.Resource {
	return &WirelessSSIDResource{}
}

type WirelessSSIDResource struct {
	client *cc.Client
}

func (r *WirelessSSIDResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid"
}

func (r *WirelessSSIDResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Wireless SSID.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Site ID - only site level Global is supported").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ssid": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the SSID").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"auth_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("L2 Authentication Type. If authType is not open, then at least one RSN Cipher Suite and corresponding valid AKM must be enabled.").AddStringEnumDescription("WPA2_ENTERPRISE", "WPA2_PERSONAL", "OPEN", "WPA3_ENTERPRISE", "WPA3_PERSONAL", "WPA2_WPA3_PERSONAL", "WPA2_WPA3_ENTERPRISE", "OPEN_SECURED").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("WPA2_ENTERPRISE", "WPA2_PERSONAL", "OPEN", "WPA3_ENTERPRISE", "WPA3_PERSONAL", "WPA2_WPA3_PERSONAL", "WPA2_WPA3_ENTERPRISE", "OPEN_SECURED"),
				},
			},
			"passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters").String,
				Optional:            true,
			},
			"fast_lane": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if FastLane is enabled, else False").String,
				Optional:            true,
			},
			"mac_filtering": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device").String,
				Optional:            true,
			},
			"ssid_radio_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio Policy").AddStringEnumDescription("Triple band operation(2.4GHz, 5GHz and 6GHz)", "5GHz only", "2.4GHz only", "6GHz only", "2.4 and 5 GHz", "2.4 and 6 GHz", "5 and 6 GHz").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Triple band operation(2.4GHz, 5GHz and 6GHz)", "5GHz only", "2.4GHz only", "6GHz only", "2.4 and 5 GHz", "2.4 and 6 GHz", "5 and 6 GHz"),
				},
			},
			"broadcast_ssid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks").String,
				Optional:            true,
			},
			"fast_transition": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Fast Transition").AddStringEnumDescription("ADAPTIVE", "ENABLE", "DISABLE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ADAPTIVE", "ENABLE", "DISABLE"),
				},
			},
			"session_timeout_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Turn on the feature that imposes a time limit on user sessions").String,
				Optional:            true,
			},
			"session_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity").AddIntegerRangeDescription(1, 86400).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
			},
			"client_exclusion": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Activate the feature that allows for the exclusion of clients").String,
				Optional:            true,
			},
			"client_exclusion_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts").AddIntegerRangeDescription(0, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"basic_service_set_max_idle": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Activate the maximum idle feature for the Basic Service Set").String,
				Optional:            true,
			},
			"basic_service_set_client_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out").AddIntegerRangeDescription(15, 100000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(15, 100000),
				},
			},
			"directed_multicast_service": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The Directed Multicast Service feature becomes operational when it is set to true").String,
				Optional:            true,
			},
			"neighbor_list": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The Neighbor List feature is enabled when it is set to true").String,
				Optional:            true,
			},
			"mft_client_protection": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management Frame Protection Client").AddStringEnumDescription("OPTIONAL", "DISABLED", "REQUIRED").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("OPTIONAL", "DISABLED", "REQUIRED"),
				},
			},
			"nas_options": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Nas Options").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName").String,
				Optional:            true,
			},
			"aaa_override": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Activate the AAA Override feature when set to true").String,
				Optional:            true,
			},
			"coverage_hole_detection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection Enable").String,
				Optional:            true,
			},
			"protected_management_frame": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("(REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)").AddStringEnumDescription("OPTIONAL", "DISABLED", "REQUIRED").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("OPTIONAL", "DISABLED", "REQUIRED"),
				},
			},
			"multi_psk_settings": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Multi PSK Settings (Only applicable for SSID with PERSONAL auth type and PSK)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"priority": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority").String,
							Required:            true,
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
				MarkdownDescription: helpers.NewAttributeDescription("This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve").AddIntegerRangeDescription(8000, 100000000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8000, 100000000000),
				},
			},
			"rsn_cipher_suite_gcmp256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated").String,
				Optional:            true,
			},
			"rsn_cipher_suite_ccmp256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated").String,
				Optional:            true,
			},
			"rsn_cipher_suite_gcmp128": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated").String,
				Optional:            true,
			},
			"rsn_cipher_suite_ccmp128": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated").String,
				Optional:            true,
			},
			"ghz6_policy_client_steering": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if 6 GHz Policy Client Steering is enabled, else False").String,
				Optional:            true,
			},
			"auth_key8021x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the 802.1X authentication key is in use").String,
				Optional:            true,
			},
			"auth_key8021x_plus_ft": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the 802.1X-Plus-FT authentication key is in use").String,
				Optional:            true,
			},
			"auth_key8021x_sha256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on").String,
				Optional:            true,
			},
			"auth_key_sae": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated").String,
				Optional:            true,
			},
			"auth_key_sae_plus_ft": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated").String,
				Optional:            true,
			},
			"auth_key_psk": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Pre-shared Key (PSK) authentication feature is enabled").String,
				Optional:            true,
			},
			"auth_key_psk_plus_ft": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated").String,
				Optional:            true,
			},
			"auth_key_owe": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on").String,
				Optional:            true,
			},
			"auth_key_easy_psk": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated").String,
				Optional:            true,
			},
			"auth_key_easy_psk_sha256": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true").String,
				Optional:            true,
			},
			"open_ssid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID").String,
				Optional:            true,
			},
			"wlan_band_select": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band").String,
				Optional:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set SSID's admin status as 'Enabled' when set to true").String,
				Optional:            true,
			},
			"auth_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of Authentication/Authorization server IpAddresses").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"acct_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of Accounting server IpAddresses").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"egress_qos": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Egress QOS").AddStringEnumDescription("PLATINUM", "SILVER", "GOLD", "BRONZE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("PLATINUM", "SILVER", "GOLD", "BRONZE"),
				},
			},
			"ingress_qos": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ingress QOS").AddStringEnumDescription("PLATINUM-UP", "SILVER-UP", "GOLD-UP", "BRONZE-UP").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("PLATINUM-UP", "SILVER-UP", "GOLD-UP", "BRONZE-UP"),
				},
			},
			"wlan_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wlan Type").AddStringEnumDescription("Enterprise", "Guest").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enterprise", "Guest"),
				},
			},
			"l3_auth_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("L3 Authentication Type").AddStringEnumDescription("open", "web_auth").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("open", "web_auth"),
				},
			},
			"auth_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth").AddStringEnumDescription("auth_ise", "auth_external", "auth_internal").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auth_ise", "auth_external", "auth_internal"),
				},
			},
			"external_auth_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)").String,
				Optional:            true,
			},
			"web_passthrough": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements").String,
				Optional:            true,
			},
			"sleeping_client": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, this will activate the timeout settings that apply to clients in sleep mode").String,
				Optional:            true,
			},
			"sleeping_client_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network").AddIntegerRangeDescription(10, 43200).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 43200),
				},
			},
			"acl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Pre-Auth Access Control List (ACL) Name").String,
				Optional:            true,
			},
			"posturing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.").String,
				Optional:            true,
			},
			"auth_key_suite_b1x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When activated by setting it to true, the SuiteB-1x authentication key feature is engaged").String,
				Optional:            true,
			},
			"auth_key_suite_b1921x": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the SuiteB192-1x authentication key feature is enabled").String,
				Optional:            true,
			},
			"auth_key_sae_ext": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on").String,
				Optional:            true,
			},
			"auth_key_sae_ext_plus_ft": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled").String,
				Optional:            true,
			},
			"ap_beacon_protection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network").String,
				Optional:            true,
			},
			"ghz24_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType").AddStringEnumDescription("dot11-g-only", "dot11-bg-only").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dot11-g-only", "dot11-bg-only"),
				},
			},
			"cckm_tsf_tolerance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cckm TImestamp Tolerance(in milliseconds)").AddIntegerRangeDescription(1000, 5000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 5000),
				},
			},
			"cckm": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if CCKM is enabled, else False").String,
				Optional:            true,
			},
			"hex": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if passphrase is in Hex format, else False").String,
				Optional:            true,
			},
			"random_mac_filter": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Deny clients using randomized MAC addresses when set to true").String,
				Optional:            true,
			},
			"fast_transition_over_the_distributed_system": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Fast Transition over the Distributed System when set to true").String,
				Optional:            true,
			},
		},
	}
}

func (r *WirelessSSIDResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *WirelessSSIDResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessSSID

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessSSID{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.#(ssid==\"" + plan.Ssid.ValueString() + "\").id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *WirelessSSIDResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessSSID

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	res = res.Get("response.#(id==\"" + state.Id.ValueString() + "\")")

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
func (r *WirelessSSIDResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessSSID

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
func (r *WirelessSSIDResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessSSID

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
func (r *WirelessSSIDResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <site_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

// End of section. //template:end import
