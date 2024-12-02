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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessSSIDDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDDataSource{}
)

func NewWirelessSSIDDataSource() datasource.DataSource {
	return &WirelessSSIDDataSource{}
}

type WirelessSSIDDataSource struct {
	client *cc.Client
}

func (d *WirelessSSIDDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid"
}

func (d *WirelessSSIDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless SSID.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "Site ID - only site level Global is supported",
				Required:            true,
			},
			"ssid": schema.StringAttribute{
				MarkdownDescription: "Name of the SSID",
				Computed:            true,
			},
			"auth_type": schema.StringAttribute{
				MarkdownDescription: "L2 Authentication Type. If authType is not open, then at least one RSN Cipher Suite and corresponding valid AKM must be enabled.",
				Computed:            true,
			},
			"passphrase": schema.StringAttribute{
				MarkdownDescription: "Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters",
				Computed:            true,
			},
			"fast_lane": schema.BoolAttribute{
				MarkdownDescription: "True if FastLane is enabled, else False",
				Computed:            true,
			},
			"mac_filtering": schema.BoolAttribute{
				MarkdownDescription: "When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device",
				Computed:            true,
			},
			"ssid_radio_type": schema.StringAttribute{
				MarkdownDescription: "Radio Policy",
				Computed:            true,
			},
			"broadcast_ssid": schema.BoolAttribute{
				MarkdownDescription: "When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks",
				Computed:            true,
			},
			"fast_transition": schema.StringAttribute{
				MarkdownDescription: "Fast Transition",
				Computed:            true,
			},
			"session_timeout_enable": schema.BoolAttribute{
				MarkdownDescription: "Turn on the feature that imposes a time limit on user sessions",
				Computed:            true,
			},
			"session_timeout": schema.Int64Attribute{
				MarkdownDescription: "This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity",
				Computed:            true,
			},
			"client_exclusion": schema.BoolAttribute{
				MarkdownDescription: "Activate the feature that allows for the exclusion of clients",
				Computed:            true,
			},
			"client_exclusion_timeout": schema.Int64Attribute{
				MarkdownDescription: "This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts",
				Computed:            true,
			},
			"basic_service_set_max_idle": schema.BoolAttribute{
				MarkdownDescription: "Activate the maximum idle feature for the Basic Service Set",
				Computed:            true,
			},
			"basic_service_set_client_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out",
				Computed:            true,
			},
			"directed_multicast_service": schema.BoolAttribute{
				MarkdownDescription: "The Directed Multicast Service feature becomes operational when it is set to true",
				Computed:            true,
			},
			"neighbor_list": schema.BoolAttribute{
				MarkdownDescription: "The Neighbor List feature is enabled when it is set to true",
				Computed:            true,
			},
			"mft_client_protection": schema.StringAttribute{
				MarkdownDescription: "Management Frame Protection Client",
				Computed:            true,
			},
			"nas_options": schema.SetAttribute{
				MarkdownDescription: "Nas Options",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: "WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName",
				Computed:            true,
			},
			"aaa_override": schema.BoolAttribute{
				MarkdownDescription: "Activate the AAA Override feature when set to true",
				Computed:            true,
			},
			"coverage_hole_detection": schema.BoolAttribute{
				MarkdownDescription: "Coverage Hole Detection Enable",
				Computed:            true,
			},
			"protected_management_frame": schema.StringAttribute{
				MarkdownDescription: "(REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)",
				Computed:            true,
			},
			"multi_psk_settings": schema.ListNestedAttribute{
				MarkdownDescription: "Multi PSK Settings (Only applicable for SSID with PERSONAL auth type and PSK)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"priority": schema.StringAttribute{
							MarkdownDescription: "Priority",
							Computed:            true,
						},
						"passphrase_type": schema.StringAttribute{
							MarkdownDescription: "Passphrase Type",
							Computed:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: "Passphrase",
							Computed:            true,
						},
					},
				},
			},
			"client_rate_limit": schema.Int64Attribute{
				MarkdownDescription: "This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve",
				Computed:            true,
			},
			"rsn_cipher_suite_gcmp256": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated",
				Computed:            true,
			},
			"rsn_cipher_suite_ccmp256": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated",
				Computed:            true,
			},
			"rsn_cipher_suite_gcmp128": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated",
				Computed:            true,
			},
			"rsn_cipher_suite_ccmp128": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated",
				Computed:            true,
			},
			"ghz6_policy_client_steering": schema.BoolAttribute{
				MarkdownDescription: "True if 6 GHz Policy Client Steering is enabled, else False",
				Computed:            true,
			},
			"auth_key8021x": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the 802.1X authentication key is in use",
				Computed:            true,
			},
			"auth_key8021x_plus_tf": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the 802.1X-Plus-FT authentication key is in use",
				Computed:            true,
			},
			"auth_key8021x_sha256": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on",
				Computed:            true,
			},
			"auth_key_sae": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated",
				Computed:            true,
			},
			"auth_key_sae_plus_ft": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated",
				Computed:            true,
			},
			"auth_key_psk": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Pre-shared Key (PSK) authentication feature is enabled",
				Computed:            true,
			},
			"auth_key_psk_plus_ft": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated",
				Computed:            true,
			},
			"auth_key_owe": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on",
				Computed:            true,
			},
			"auth_key_easy_psk": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated",
				Computed:            true,
			},
			"auth_key_easy_psk_sha256": schema.BoolAttribute{
				MarkdownDescription: "The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true",
				Computed:            true,
			},
			"open_ssid": schema.BoolAttribute{
				MarkdownDescription: "Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID",
				Computed:            true,
			},
			"wlan_band_select": schema.BoolAttribute{
				MarkdownDescription: "Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Set SSID's admin status as 'Enabled' when set to true",
				Computed:            true,
			},
			"auth_servers": schema.SetAttribute{
				MarkdownDescription: "List of Authentication/Authorization server IpAddresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"acct_servers": schema.SetAttribute{
				MarkdownDescription: "List of Accounting server IpAddresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"egress_qos": schema.StringAttribute{
				MarkdownDescription: "Egress QOS",
				Computed:            true,
			},
			"ingress_qos": schema.StringAttribute{
				MarkdownDescription: "Ingress QOS",
				Computed:            true,
			},
			"wlan_type": schema.StringAttribute{
				MarkdownDescription: "Wlan Type",
				Computed:            true,
			},
			"l3_auth_type": schema.StringAttribute{
				MarkdownDescription: "L3 Authentication Type",
				Computed:            true,
			},
			"auth_server": schema.StringAttribute{
				MarkdownDescription: "Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth",
				Computed:            true,
			},
			"external_auth_ip_address": schema.StringAttribute{
				MarkdownDescription: "External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)",
				Computed:            true,
			},
			"web_passthrough": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements",
				Computed:            true,
			},
			"sleeping_client": schema.BoolAttribute{
				MarkdownDescription: "When set to true, this will activate the timeout settings that apply to clients in sleep mode",
				Computed:            true,
			},
			"sleeping_client_timeout": schema.Int64Attribute{
				MarkdownDescription: "This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network",
				Computed:            true,
			},
			"acl_name": schema.StringAttribute{
				MarkdownDescription: "Pre-Auth Access Control List (ACL) Name",
				Computed:            true,
			},
			"posturing": schema.BoolAttribute{
				MarkdownDescription: "Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.",
				Computed:            true,
			},
			"auth_key_suite_b1x": schema.BoolAttribute{
				MarkdownDescription: "When activated by setting it to true, the SuiteB-1x authentication key feature is engaged",
				Computed:            true,
			},
			"auth_key_suite_b1921x": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the SuiteB192-1x authentication key feature is enabled",
				Computed:            true,
			},
			"auth_key_sae_ext": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on",
				Computed:            true,
			},
			"auth_key_sae_ext_plus_tf": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled",
				Computed:            true,
			},
			"ap_beacon_protection": schema.BoolAttribute{
				MarkdownDescription: "When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network",
				Computed:            true,
			},
			"ghz24_policy": schema.StringAttribute{
				MarkdownDescription: "2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType",
				Computed:            true,
			},
			"cckm_tsf_tolerance": schema.Int64Attribute{
				MarkdownDescription: "Cckm TImestamp Tolerance(in milliseconds)",
				Computed:            true,
			},
			"cckm": schema.BoolAttribute{
				MarkdownDescription: "True if CCKM is enabled, else False",
				Computed:            true,
			},
			"hex": schema.BoolAttribute{
				MarkdownDescription: "True if passphrase is in Hex format, else False",
				Computed:            true,
			},
			"random_mac_filter": schema.BoolAttribute{
				MarkdownDescription: "Deny clients using randomized MAC addresses when set to true",
				Computed:            true,
			},
			"fast_transition_over_the_distributed_system": schema.BoolAttribute{
				MarkdownDescription: "Enable Fast Transition over the Distributed System when set to true",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessSSIDDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessSSIDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSID

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	res = res.Get("response.#(id==\"" + config.Id.ValueString() + "\")")

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
