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
var _ resource.Resource = &WirelessRFProfileResource{}
var _ resource.ResourceWithImportState = &WirelessRFProfileResource{}

func NewWirelessRFProfileResource() resource.Resource {
	return &WirelessRFProfileResource{}
}

type WirelessRFProfileResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *WirelessRFProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profile"
}

func (r *WirelessRFProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Wireless RF Profile.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"rf_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RF Profile Name").String,
				Required:            true,
			},
			"default_rf_profile": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specifies if the RF Profile is the default. Only one RF Profile can be marked as default at a time.").String,
				Required:            true,
			},
			"enable_radio_type_a": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the 5 GHz radio band in the RF Profile.").String,
				Required:            true,
			},
			"enable_radio_type_b": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the 2.4 GHz radio band in the RF Profile.").String,
				Required:            true,
			},
			"enable_radio_type6_g_hz": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the 6 GHz radio band in the RF Profile.").String,
				Required:            true,
			},
			"radio_type_a_parent_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Parent profile for the 5 GHz radio band.").AddStringEnumDescription("HIGH", "TYPICAL", "LOW", "CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HIGH", "TYPICAL", "LOW", "CUSTOM"),
				},
			},
			"radio_type_a_radio_channels": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DCA channels for the 5 GHz radio band, passed in comma-separated format without spaces.").String,
				Optional:            true,
			},
			"radio_type_a_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Data rates for the 5 GHz radio band, passed in comma-separated format without spaces.").String,
				Optional:            true,
			},
			"radio_type_a_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mandatory data rates for the 5 GHz radio band, passed in comma-separated format, must be a subset of dataRates with a maximum of 2 values.").String,
				Optional:            true,
			},
			"radio_type_a_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Power threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-80, -50).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-80, -50),
				},
			},
			"radio_type_a_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RX-SOP threshold for the 5 GHz radio band.").AddStringEnumDescription("HIGH", "MEDIUM", "LOW", "AUTO", "CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HIGH", "MEDIUM", "LOW", "AUTO", "CUSTOM"),
				},
			},
			"radio_type_a_min_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum power level for the 5 GHz radio band.").AddIntegerRangeDescription(-10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-10, 30),
				},
			},
			"radio_type_a_max_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum power level for the 5 GHz radio band.").AddIntegerRangeDescription(-10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-10, 30),
				},
			},
			"radio_type_a_channel_width": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Channel width for the 5 GHz radio band.").AddStringEnumDescription("20", "40", "80", "160", "best").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("20", "40", "80", "160", "best"),
				},
			},
			"radio_type_a_preamble_puncture": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable preamble puncturing for the 5 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_a_zero_wait_dfs_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Zero Wait DFS for IOS-XE-based Wireless Controllers running version 17.9.1 and above.").String,
				Optional:            true,
			},
			"radio_type_a_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom RX-SOP threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-85, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-85, -60),
				},
			},
			"radio_type_a_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client limit for the 5 GHz radio band.").AddIntegerRangeDescription(0, 500).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 500),
				},
			},
			"radio_type_a_fra_properties_client_aware": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client awareness for the 5 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_a_fra_properties_client_select": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client selection percentage for the 5 GHz radio band.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"radio_type_a_fra_properties_client_reset": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client reset percentage for the 5 GHz radio band.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"radio_type_a_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection client level for the 5 GHz radio band.").AddIntegerRangeDescription(1, 200).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 200),
				},
			},
			"radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection data RSSI threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-90, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-90, -60),
				},
			},
			"radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection voice RSSI threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-90, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-90, -60),
				},
			},
			"radio_type_a_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection exception level (%) for the 5 GHz radio band.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax Non-SRG OBSS PD for the 5 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax Non-SRG OBSS PD max threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD for the 5 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD min threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_a_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD max threshold for the 5 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_b_parent_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Parent profile for the 2.4 GHz radio band.").AddStringEnumDescription("HIGH", "TYPICAL", "LOW", "CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HIGH", "TYPICAL", "LOW", "CUSTOM"),
				},
			},
			"radio_type_b_radio_channels": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DCA channels for the 2.4 GHz radio band, passed in comma-separated format without spaces.").String,
				Optional:            true,
			},
			"radio_type_b_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Data rates for the 2.4 GHz radio band, passed in comma-separated format without spaces.").String,
				Optional:            true,
			},
			"radio_type_b_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mandatory data rates for the 2.4 GHz radio band, passed in comma-separated format, must be a subset of dataRates with a maximum of 2 values.").String,
				Optional:            true,
			},
			"radio_type_b_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Power threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-80, -50).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-80, -50),
				},
			},
			"radio_type_b_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RX-SOP threshold for the 2.4 GHz radio band.").AddStringEnumDescription("HIGH", "MEDIUM", "LOW", "AUTO", "CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HIGH", "MEDIUM", "LOW", "AUTO", "CUSTOM"),
				},
			},
			"radio_type_b_min_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum power level for the 2.4 GHz radio band.").AddIntegerRangeDescription(-10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-10, 30),
				},
			},
			"radio_type_b_max_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum power level for the 2.4 GHz radio band.").AddIntegerRangeDescription(-10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-10, 30),
				},
			},
			"radio_type_b_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom RX-SOP threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-85, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-85, -60),
				},
			},
			"radio_type_b_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client limit for the 2.4 GHz radio band.").AddIntegerRangeDescription(0, 500).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 500),
				},
			},
			"radio_type_b_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection client level for the 2.4 GHz radio band.").AddIntegerRangeDescription(1, 200).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 200),
				},
			},
			"radio_type_b_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection data RSSI threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-90, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-90, -60),
				},
			},
			"radio_type_b_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection voice RSSI threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-90, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-90, -60),
				},
			},
			"radio_type_b_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection exception level (%) for the 2.4 GHz radio band.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax Non-SRG OBSS PD for the 2.4 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax Non-SRG OBSS PD max threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD for the 2.4 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD min threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD max threshold for the 2.4 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_c_parent_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Parent profile for the 6 GHz radio band.").AddStringEnumDescription("CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CUSTOM"),
				},
			},
			"radio_type_c_radio_channels": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DCA channels of 6 GHz radio band, passed in comma-separated format without spaces.").String,
				Optional:            true,
			},
			"radio_type_c_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Data rates of 6 GHz radio band, passed in comma-separated format without spaces.").String,
				Optional:            true,
			},
			"radio_type_c_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mandatory data rates of 6 GHz radio band, passed in comma-separated format without spaces. Must be a subset of dataRates with a maximum of 2 values.").String,
				Optional:            true,
			},
			"radio_type_c_power_threshold_v1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Power threshold of the 6 GHz radio band.").AddIntegerRangeDescription(-80, -50).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-80, -50),
				},
			},
			"radio_type_c_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RX-SOP threshold of the 6 GHz radio band.").AddStringEnumDescription("HIGH", "MEDIUM", "LOW", "AUTO", "CUSTOM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HIGH", "MEDIUM", "LOW", "AUTO", "CUSTOM"),
				},
			},
			"radio_type_c_min_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum power level of the 6 GHz radio band.").AddIntegerRangeDescription(-10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-10, 30),
				},
			},
			"radio_type_c_max_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum power level of the 6 GHz radio band.").AddIntegerRangeDescription(-10, 30).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-10, 30),
				},
			},
			"radio_type_c_enable_standard_power_service": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if Standard Power Service is enabled, else False.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_down_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA Downlink for 802.11ax parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_up_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA Uplink for 802.11ax parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_up_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MU-MIMO Uplink for 802.11ax parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_down_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MU-MIMO Downlink for 802.11ax parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_down_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA dma down link for 802.11be parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_up_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA dma up link for 802.11be parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_up_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA mu mimo up link for 802.11be parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_down_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA mu mimo down link for 802.11be parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_multi_ru": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OFDMA Multi-RU for 802.11be parameters in the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_target_wake_time": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Target Wake Time for the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_multi_bssid_properties_twt_broadcast_support": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("TWT Broadcast Support for the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_preamble_puncture": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or Disable Preamble Puncturing. This WiFi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher.").String,
				Optional:            true,
			},
			"radio_type_c_min_dbs_width": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum DBS Width for the 6 GHz radio band.").AddStringEnumDescription("20", "40", "80", "160", "320").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.OneOf(20, 40, 80, 160, 320),
				},
			},
			"radio_type_c_max_dbs_width": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum DBS Width for the 6 GHz radio band.").AddStringEnumDescription("20", "40", "80", "160", "320").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.OneOf(20, 40, 80, 160, 320),
				},
			},
			"radio_type_c_max_radio_clients": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client limit for the 6 GHz radio band.").AddIntegerRangeDescription(0, 500).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 500),
				},
			},
			"radio_type_c_custom_rx_sop_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom RX-SOP threshold for the 6 GHz radio band.").AddIntegerRangeDescription(-85, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-85, -60),
				},
			},
			"radio_type_c_psc_enforcing_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable PSC enforcement for the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_discovery_frames_6ghz": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Discovery Frames for the 6 GHz radio band.").AddStringEnumDescription("None", "Broadcast Probe Response", "FILS Discovery").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("None", "Broadcast Probe Response", "FILS Discovery"),
				},
			},
			"radio_type_c_broadcast_probe_response_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Broadcast Probe Response Interval for the 6 GHz radio band.").AddIntegerRangeDescription(5, 25).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 25),
				},
			},
			"radio_type_c_fra_properties_client_reset_count": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client Reset Count for the 6 GHz radio band.").AddIntegerRangeDescription(1, 10).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"radio_type_c_fra_properties_client_utilization_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Client Utilization Threshold for the 6 GHz radio band.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"radio_type_c_coverage_hole_detection_properties_chd_client_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection client level for the 6 GHz radio band.").AddIntegerRangeDescription(1, 200).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 200),
				},
			},
			"radio_type_c_coverage_hole_detection_properties_chd_data_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection data RSSI threshold for the 6 GHz radio band.").AddIntegerRangeDescription(-90, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-90, -60),
				},
			},
			"radio_type_c_coverage_hole_detection_properties_chd_voice_rssi_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection voice RSSI threshold for the 6 GHz radio band.").AddIntegerRangeDescription(-90, -60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-90, -60),
				},
			},
			"radio_type_c_coverage_hole_detection_properties_chd_exception_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Coverage Hole Detection exception level (%) for the 6 GHz radio band.").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax Non-SRG OBSS PD for the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax Non-SRG OBSS PD Max Threshold for the 6 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD for the 6 GHz radio band.").String,
				Optional:            true,
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD Min Threshold for the 6 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
			"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dot11ax SRG OBSS PD Max Threshold for the 6 GHz radio band.").AddIntegerRangeDescription(-82, -62).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(-82, -62),
				},
			},
		},
	}
}

func (r *WirelessRFProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *WirelessRFProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessRFProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessRFProfile{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	params += "?rfProfileName=" + url.QueryEscape(plan.RfProfileName.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.0.id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *WirelessRFProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessRFProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?rfProfileName=" + url.QueryEscape(state.RfProfileName.ValueString())
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
func (r *WirelessRFProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessRFProfile

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
func (r *WirelessRFProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessRFProfile

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
func (r *WirelessRFProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <rf_profile_name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("rf_profile_name"), idParts[0])...)
}

// End of section. //template:end import
