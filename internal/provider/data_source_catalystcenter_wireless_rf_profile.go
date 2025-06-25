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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessRFProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessRFProfileDataSource{}
)

func NewWirelessRFProfileDataSource() datasource.DataSource {
	return &WirelessRFProfileDataSource{}
}

type WirelessRFProfileDataSource struct {
	client *cc.Client
}

func (d *WirelessRFProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profile"
}

func (d *WirelessRFProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless RF Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"rf_profile_name": schema.StringAttribute{
				MarkdownDescription: "RF Profile Name",
				Required:            true,
			},
			"default_rf_profile": schema.BoolAttribute{
				MarkdownDescription: "Specifies if the RF Profile is the default. Only one RF Profile can be marked as default at a time.",
				Computed:            true,
			},
			"enable_radio_type_a": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable the 5 GHz radio band in the RF Profile.",
				Computed:            true,
			},
			"enable_radio_type_b": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable the 2.4 GHz radio band in the RF Profile.",
				Computed:            true,
			},
			"enable_radio_type6_g_hz": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable the 6 GHz radio band in the RF Profile.",
				Computed:            true,
			},
			"radio_type_a_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Parent profile for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_radio_channels": schema.StringAttribute{
				MarkdownDescription: "DCA channels for the 5 GHz radio band, passed in comma-separated format without spaces.",
				Computed:            true,
			},
			"radio_type_a_data_rates": schema.StringAttribute{
				MarkdownDescription: "Data rates for the 5 GHz radio band, passed in comma-separated format without spaces.",
				Computed:            true,
			},
			"radio_type_a_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Mandatory data rates for the 5 GHz radio band, passed in comma-separated format, must be a subset of dataRates with a maximum of 2 values.",
				Computed:            true,
			},
			"radio_type_a_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Power threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "RX-SOP threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Minimum power level for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Maximum power level for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_channel_width": schema.StringAttribute{
				MarkdownDescription: "Channel width for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_preamble_puncture": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable preamble puncturing for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_zero_wait_dfs_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable Zero Wait DFS for IOS-XE-based Wireless Controllers running version 17.9.1 and above.",
				Computed:            true,
			},
			"radio_type_a_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: "Custom RX-SOP threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: "Client limit for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_fra_properties_client_aware": schema.BoolAttribute{
				MarkdownDescription: "Client awareness for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_fra_properties_client_select": schema.Int64Attribute{
				MarkdownDescription: "Client selection percentage for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_fra_properties_client_reset": schema.Int64Attribute{
				MarkdownDescription: "Client reset percentage for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection client level for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection data RSSI threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection voice RSSI threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection exception level (%) for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Dot11ax Non-SRG OBSS PD for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax Non-SRG OBSS PD max threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD min threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD max threshold for the 5 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Parent profile for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_radio_channels": schema.StringAttribute{
				MarkdownDescription: "DCA channels for the 2.4 GHz radio band, passed in comma-separated format without spaces.",
				Computed:            true,
			},
			"radio_type_b_data_rates": schema.StringAttribute{
				MarkdownDescription: "Data rates for the 2.4 GHz radio band, passed in comma-separated format without spaces.",
				Computed:            true,
			},
			"radio_type_b_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Mandatory data rates for the 2.4 GHz radio band, passed in comma-separated format, must be a subset of dataRates with a maximum of 2 values.",
				Computed:            true,
			},
			"radio_type_b_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Power threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "RX-SOP threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Minimum power level for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Maximum power level for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: "Custom RX-SOP threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: "Client limit for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection client level for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection data RSSI threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection voice RSSI threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection exception level (%) for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Dot11ax Non-SRG OBSS PD for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax Non-SRG OBSS PD max threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD min threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD max threshold for the 2.4 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Parent profile for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_radio_channels": schema.StringAttribute{
				MarkdownDescription: "DCA channels of 6 GHz radio band, passed in comma-separated format without spaces.",
				Computed:            true,
			},
			"radio_type_c_data_rates": schema.StringAttribute{
				MarkdownDescription: "Data rates of 6 GHz radio band, passed in comma-separated format without spaces.",
				Computed:            true,
			},
			"radio_type_c_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Mandatory data rates of 6 GHz radio band, passed in comma-separated format without spaces. Must be a subset of dataRates with a maximum of 2 values.",
				Computed:            true,
			},
			"radio_type_c_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Power threshold of the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "RX-SOP threshold of the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Minimum power level of the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Maximum power level of the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_enable_standard_power_service": schema.BoolAttribute{
				MarkdownDescription: "True if Standard Power Service is enabled, else False.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_down_link": schema.BoolAttribute{
				MarkdownDescription: "OFDMA Downlink for 802.11ax parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_up_link": schema.BoolAttribute{
				MarkdownDescription: "OFDMA Uplink for 802.11ax parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_up_link": schema.BoolAttribute{
				MarkdownDescription: "MU-MIMO Uplink for 802.11ax parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_down_link": schema.BoolAttribute{
				MarkdownDescription: "MU-MIMO Downlink for 802.11ax parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_down_link": schema.BoolAttribute{
				MarkdownDescription: "OFDMA dma down link for 802.11be parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_up_link": schema.BoolAttribute{
				MarkdownDescription: "OFDMA dma up link for 802.11be parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_up_link": schema.BoolAttribute{
				MarkdownDescription: "OFDMA mu mimo up link for 802.11be parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_down_link": schema.BoolAttribute{
				MarkdownDescription: "OFDMA mu mimo down link for 802.11be parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_multi_ru": schema.BoolAttribute{
				MarkdownDescription: "OFDMA Multi-RU for 802.11be parameters in the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_target_wake_time": schema.BoolAttribute{
				MarkdownDescription: "Target Wake Time for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_multi_bssid_properties_twt_broadcast_support": schema.BoolAttribute{
				MarkdownDescription: "TWT Broadcast Support for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_preamble_puncture": schema.BoolAttribute{
				MarkdownDescription: "Enable or Disable Preamble Puncturing. This WiFi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher.",
				Computed:            true,
			},
			"radio_type_c_min_dbs_width": schema.Int64Attribute{
				MarkdownDescription: "Minimum DBS Width for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_max_dbs_width": schema.Int64Attribute{
				MarkdownDescription: "Maximum DBS Width for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: "Client limit for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: "Custom RX-SOP threshold for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_psc_enforcing_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable PSC enforcement for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_discovery_frames_6ghz": schema.StringAttribute{
				MarkdownDescription: "Discovery Frames for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_broadcast_probe_response_interval": schema.Int64Attribute{
				MarkdownDescription: "Broadcast Probe Response Interval for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_fra_properties_client_reset_count": schema.Int64Attribute{
				MarkdownDescription: "Client Reset Count for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_fra_properties_client_utilization_threshold": schema.Int64Attribute{
				MarkdownDescription: "Client Utilization Threshold for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection client level for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection data RSSI threshold for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection voice RSSI threshold for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: "Coverage Hole Detection exception level (%) for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Dot11ax Non-SRG OBSS PD for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax Non-SRG OBSS PD Max Threshold for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD Min Threshold for the 6 GHz radio band.",
				Computed:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: "Dot11ax SRG OBSS PD Max Threshold for the 6 GHz radio band.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessRFProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessRFProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessRFProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?rfProfileName=" + url.QueryEscape(config.RfProfileName.ValueString())
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
