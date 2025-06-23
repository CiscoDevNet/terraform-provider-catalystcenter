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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessRFProfile struct {
	Id                                                                         types.String `tfsdk:"id"`
	RfProfileName                                                              types.String `tfsdk:"rf_profile_name"`
	DefaultRfProfile                                                           types.Bool   `tfsdk:"default_rf_profile"`
	EnableRadioTypeA                                                           types.Bool   `tfsdk:"enable_radio_type_a"`
	EnableRadioTypeB                                                           types.Bool   `tfsdk:"enable_radio_type_b"`
	EnableRadioType6GHz                                                        types.Bool   `tfsdk:"enable_radio_type6_g_hz"`
	RadioTypeAParentProfile                                                    types.String `tfsdk:"radio_type_a_parent_profile"`
	RadioTypeARadioChannels                                                    types.String `tfsdk:"radio_type_a_radio_channels"`
	RadioTypeADataRates                                                        types.String `tfsdk:"radio_type_a_data_rates"`
	RadioTypeAMandatoryDataRates                                               types.String `tfsdk:"radio_type_a_mandatory_data_rates"`
	RadioTypeAPowerThresholdV1                                                 types.Int64  `tfsdk:"radio_type_a_power_threshold_v1"`
	RadioTypeARxSopThreshold                                                   types.String `tfsdk:"radio_type_a_rx_sop_threshold"`
	RadioTypeAMinPowerLevel                                                    types.Int64  `tfsdk:"radio_type_a_min_power_level"`
	RadioTypeAMaxPowerLevel                                                    types.Int64  `tfsdk:"radio_type_a_max_power_level"`
	RadioTypeAChannelWidth                                                     types.String `tfsdk:"radio_type_a_channel_width"`
	RadioTypeAPreamblePuncture                                                 types.Bool   `tfsdk:"radio_type_a_preamble_puncture"`
	RadioTypeAZeroWaitDfsEnable                                                types.Bool   `tfsdk:"radio_type_a_zero_wait_dfs_enable"`
	RadioTypeACustomRxSopThreshold                                             types.Int64  `tfsdk:"radio_type_a_custom_rx_sop_threshold"`
	RadioTypeAMaxRadioClients                                                  types.Int64  `tfsdk:"radio_type_a_max_radio_clients"`
	RadioTypeAFraPropertiesClientAware                                         types.Bool   `tfsdk:"radio_type_a_fra_properties_client_aware"`
	RadioTypeAFraPropertiesClientSelect                                        types.Int64  `tfsdk:"radio_type_a_fra_properties_client_select"`
	RadioTypeAFraPropertiesClientReset                                         types.Int64  `tfsdk:"radio_type_a_fra_properties_client_reset"`
	RadioTypeACoverageHoleDetectionPropertiesChdClientLevel                    types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_client_level"`
	RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold              types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold"`
	RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold             types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold"`
	RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel                 types.Int64  `tfsdk:"radio_type_a_coverage_hole_detection_properties_chd_exception_level"`
	RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect             types.Bool   `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect"`
	RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold types.Int64  `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold"`
	RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect                types.Bool   `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect"`
	RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold    types.Int64  `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold"`
	RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold    types.Int64  `tfsdk:"radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold"`
	RadioTypeBParentProfile                                                    types.String `tfsdk:"radio_type_b_parent_profile"`
	RadioTypeBRadioChannels                                                    types.String `tfsdk:"radio_type_b_radio_channels"`
	RadioTypeBDataRates                                                        types.String `tfsdk:"radio_type_b_data_rates"`
	RadioTypeBMandatoryDataRates                                               types.String `tfsdk:"radio_type_b_mandatory_data_rates"`
	RadioTypeBPowerThresholdV1                                                 types.Int64  `tfsdk:"radio_type_b_power_threshold_v1"`
	RadioTypeBRxSopThreshold                                                   types.String `tfsdk:"radio_type_b_rx_sop_threshold"`
	RadioTypeBMinPowerLevel                                                    types.Int64  `tfsdk:"radio_type_b_min_power_level"`
	RadioTypeBMaxPowerLevel                                                    types.Int64  `tfsdk:"radio_type_b_max_power_level"`
	RadioTypeBCustomRxSopThreshold                                             types.Int64  `tfsdk:"radio_type_b_custom_rx_sop_threshold"`
	RadioTypeBMaxRadioClients                                                  types.Int64  `tfsdk:"radio_type_b_max_radio_clients"`
	RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel                    types.Int64  `tfsdk:"radio_type_b_coverage_hole_detection_properties_chd_client_level"`
	RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold              types.Int64  `tfsdk:"radio_type_b_coverage_hole_detection_properties_chd_data_rssi_threshold"`
	RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold             types.Int64  `tfsdk:"radio_type_b_coverage_hole_detection_properties_chd_voice_rssi_threshold"`
	RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel                 types.Int64  `tfsdk:"radio_type_b_coverage_hole_detection_properties_chd_exception_level"`
	RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect              types.Bool   `tfsdk:"radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect"`
	RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold  types.Int64  `tfsdk:"radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold"`
	RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect                 types.Bool   `tfsdk:"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect"`
	RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold     types.Int64  `tfsdk:"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold"`
	RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold     types.Int64  `tfsdk:"radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold"`
	RadioTypeCParentProfile                                                    types.String `tfsdk:"radio_type_c_parent_profile"`
	RadioTypeCRadioChannels                                                    types.String `tfsdk:"radio_type_c_radio_channels"`
	RadioTypeCDataRates                                                        types.String `tfsdk:"radio_type_c_data_rates"`
	RadioTypeCMandatoryDataRates                                               types.String `tfsdk:"radio_type_c_mandatory_data_rates"`
	RadioTypeCPowerThresholdV1                                                 types.Int64  `tfsdk:"radio_type_c_power_threshold_v1"`
	RadioTypeCRxSopThreshold                                                   types.String `tfsdk:"radio_type_c_rx_sop_threshold"`
	RadioTypeCMinPowerLevel                                                    types.Int64  `tfsdk:"radio_type_c_min_power_level"`
	RadioTypeCMaxPowerLevel                                                    types.Int64  `tfsdk:"radio_type_c_max_power_level"`
	RadioTypeCEnableStandardPowerService                                       types.Bool   `tfsdk:"radio_type_c_enable_standard_power_service"`
	RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink               types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_down_link"`
	RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink                 types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_up_link"`
	RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink                types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_up_link"`
	RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink              types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_down_link"`
	RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink               types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_down_link"`
	RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink                 types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_up_link"`
	RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink                types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_up_link"`
	RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink              types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_down_link"`
	RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu                types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_multi_ru"`
	RadioTypeCMultiBssidPropertiesTargetWakeTime                               types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_target_wake_time"`
	RadioTypeCMultiBssidPropertiesTwtBroadcastSupport                          types.Bool   `tfsdk:"radio_type_c_multi_bssid_properties_twt_broadcast_support"`
	RadioTypeCPreamblePuncture                                                 types.Bool   `tfsdk:"radio_type_c_preamble_puncture"`
	RadioTypeCMinDbsWidth                                                      types.Int64  `tfsdk:"radio_type_c_min_dbs_width"`
	RadioTypeCMaxDbsWidth                                                      types.Int64  `tfsdk:"radio_type_c_max_dbs_width"`
	RadioTypeCMaxRadioClients                                                  types.Int64  `tfsdk:"radio_type_c_max_radio_clients"`
	RadioTypeCCustomRxSopThreshold                                             types.Int64  `tfsdk:"radio_type_c_custom_rx_sop_threshold"`
	RadioTypeCPscEnforcingEnabled                                              types.Bool   `tfsdk:"radio_type_c_psc_enforcing_enabled"`
	RadioTypeCDiscoveryFrames6ghz                                              types.String `tfsdk:"radio_type_c_discovery_frames_6ghz"`
	RadioTypeCBroadcastProbeResponseInterval                                   types.Int64  `tfsdk:"radio_type_c_broadcast_probe_response_interval"`
	RadioTypeCFraPropertiesClientResetCount                                    types.Int64  `tfsdk:"radio_type_c_fra_properties_client_reset_count"`
	RadioTypeCFraPropertiesClientUtilizationThreshold                          types.Int64  `tfsdk:"radio_type_c_fra_properties_client_utilization_threshold"`
	RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel                    types.Int64  `tfsdk:"radio_type_c_coverage_hole_detection_properties_chd_client_level"`
	RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold              types.Int64  `tfsdk:"radio_type_c_coverage_hole_detection_properties_chd_data_rssi_threshold"`
	RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold             types.Int64  `tfsdk:"radio_type_c_coverage_hole_detection_properties_chd_voice_rssi_threshold"`
	RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel                 types.Int64  `tfsdk:"radio_type_c_coverage_hole_detection_properties_chd_exception_level"`
	RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect              types.Bool   `tfsdk:"radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect"`
	RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold  types.Int64  `tfsdk:"radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold"`
	RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect                 types.Bool   `tfsdk:"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect"`
	RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold     types.Int64  `tfsdk:"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold"`
	RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold     types.Int64  `tfsdk:"radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessRFProfile) getPath() string {
	return "/dna/intent/api/v1/wirelessSettings/rfProfiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessRFProfile) toBody(ctx context.Context, state WirelessRFProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.RfProfileName.IsNull() {
		body, _ = sjson.Set(body, "rfProfileName", data.RfProfileName.ValueString())
	}
	if !data.DefaultRfProfile.IsNull() {
		body, _ = sjson.Set(body, "defaultRfProfile", data.DefaultRfProfile.ValueBool())
	}
	if !data.EnableRadioTypeA.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeA", data.EnableRadioTypeA.ValueBool())
	}
	if !data.EnableRadioTypeB.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeB", data.EnableRadioTypeB.ValueBool())
	}
	if !data.EnableRadioType6GHz.IsNull() {
		body, _ = sjson.Set(body, "enableRadioType6GHz", data.EnableRadioType6GHz.ValueBool())
	}
	if !data.RadioTypeAParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.parentProfile", data.RadioTypeAParentProfile.ValueString())
	}
	if !data.RadioTypeARadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.radioChannels", data.RadioTypeARadioChannels.ValueString())
	}
	if !data.RadioTypeADataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.dataRates", data.RadioTypeADataRates.ValueString())
	}
	if !data.RadioTypeAMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.mandatoryDataRates", data.RadioTypeAMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeAPowerThresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.powerThresholdV1", data.RadioTypeAPowerThresholdV1.ValueInt64())
	}
	if !data.RadioTypeARxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.rxSopThreshold", data.RadioTypeARxSopThreshold.ValueString())
	}
	if !data.RadioTypeAMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.minPowerLevel", data.RadioTypeAMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeAMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.maxPowerLevel", data.RadioTypeAMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeAChannelWidth.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.channelWidth", data.RadioTypeAChannelWidth.ValueString())
	}
	if !data.RadioTypeAPreamblePuncture.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.preamblePuncture", data.RadioTypeAPreamblePuncture.ValueBool())
	}
	if !data.RadioTypeAZeroWaitDfsEnable.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.zeroWaitDfsEnable", data.RadioTypeAZeroWaitDfsEnable.ValueBool())
	}
	if !data.RadioTypeACustomRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.customRxSopThreshold", data.RadioTypeACustomRxSopThreshold.ValueInt64())
	}
	if !data.RadioTypeAMaxRadioClients.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.maxRadioClients", data.RadioTypeAMaxRadioClients.ValueInt64())
	}
	if !data.RadioTypeAFraPropertiesClientAware.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.fraPropertiesA.clientAware", data.RadioTypeAFraPropertiesClientAware.ValueBool())
	}
	if !data.RadioTypeAFraPropertiesClientSelect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.fraPropertiesA.clientSelect", data.RadioTypeAFraPropertiesClientSelect.ValueInt64())
	}
	if !data.RadioTypeAFraPropertiesClientReset.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.fraPropertiesA.clientReset", data.RadioTypeAFraPropertiesClientReset.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdClientLevel", data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdDataRssiThreshold", data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold", data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.coverageHoleDetectionProperties.chdExceptionLevel", data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.ValueInt64())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect", data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold", data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetect", data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold", data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.ValueInt64())
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold", data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeBParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.parentProfile", data.RadioTypeBParentProfile.ValueString())
	}
	if !data.RadioTypeBRadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.radioChannels", data.RadioTypeBRadioChannels.ValueString())
	}
	if !data.RadioTypeBDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.dataRates", data.RadioTypeBDataRates.ValueString())
	}
	if !data.RadioTypeBMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.mandatoryDataRates", data.RadioTypeBMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeBPowerThresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.powerThresholdV1", data.RadioTypeBPowerThresholdV1.ValueInt64())
	}
	if !data.RadioTypeBRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.rxSopThreshold", data.RadioTypeBRxSopThreshold.ValueString())
	}
	if !data.RadioTypeBMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.minPowerLevel", data.RadioTypeBMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeBMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.maxPowerLevel", data.RadioTypeBMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeBCustomRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.customRxSopThreshold", data.RadioTypeBCustomRxSopThreshold.ValueInt64())
	}
	if !data.RadioTypeBMaxRadioClients.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.maxRadioClients", data.RadioTypeBMaxRadioClients.ValueInt64())
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.coverageHoleDetectionProperties.chdClientLevel", data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel.ValueInt64())
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.coverageHoleDetectionProperties.chdDataRssiThreshold", data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold", data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.coverageHoleDetectionProperties.chdExceptionLevel", data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel.ValueInt64())
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect", data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold", data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetect", data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold", data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.ValueInt64())
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold", data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeCParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.parentProfile", data.RadioTypeCParentProfile.ValueString())
	}
	if !data.RadioTypeCRadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.radioChannels", data.RadioTypeCRadioChannels.ValueString())
	}
	if !data.RadioTypeCDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.dataRates", data.RadioTypeCDataRates.ValueString())
	}
	if !data.RadioTypeCMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.mandatoryDataRates", data.RadioTypeCMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeCPowerThresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.powerThresholdV1", data.RadioTypeCPowerThresholdV1.ValueInt64())
	}
	if !data.RadioTypeCRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.rxSopThreshold", data.RadioTypeCRxSopThreshold.ValueString())
	}
	if !data.RadioTypeCMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.minPowerLevel", data.RadioTypeCMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeCMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.maxPowerLevel", data.RadioTypeCMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeCEnableStandardPowerService.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.enableStandardPowerService", data.RadioTypeCEnableStandardPowerService.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11axParameters.ofdmaDownLink", data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11axParameters.ofdmaUpLink", data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11axParameters.muMimoUpLink", data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11axParameters.muMimoDownLink", data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaDownLink", data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaUpLink", data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11beParameters.muMimoUpLink", data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11beParameters.muMimoDownLink", data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaMultiRu", data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesTargetWakeTime.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.targetWakeTime", data.RadioTypeCMultiBssidPropertiesTargetWakeTime.ValueBool())
	}
	if !data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.multiBssidProperties.twtBroadcastSupport", data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport.ValueBool())
	}
	if !data.RadioTypeCPreamblePuncture.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.preamblePuncture", data.RadioTypeCPreamblePuncture.ValueBool())
	}
	if !data.RadioTypeCMinDbsWidth.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.minDbsWidth", data.RadioTypeCMinDbsWidth.ValueInt64())
	}
	if !data.RadioTypeCMaxDbsWidth.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.maxDbsWidth", data.RadioTypeCMaxDbsWidth.ValueInt64())
	}
	if !data.RadioTypeCMaxRadioClients.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.maxRadioClients", data.RadioTypeCMaxRadioClients.ValueInt64())
	}
	if !data.RadioTypeCCustomRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.customRxSopThreshold", data.RadioTypeCCustomRxSopThreshold.ValueInt64())
	}
	if !data.RadioTypeCPscEnforcingEnabled.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.pscEnforcingEnabled", data.RadioTypeCPscEnforcingEnabled.ValueBool())
	}
	if !data.RadioTypeCDiscoveryFrames6ghz.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.discoveryFrames6GHz", data.RadioTypeCDiscoveryFrames6ghz.ValueString())
	}
	if !data.RadioTypeCBroadcastProbeResponseInterval.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.broadcastProbeResponseInterval", data.RadioTypeCBroadcastProbeResponseInterval.ValueInt64())
	}
	if !data.RadioTypeCFraPropertiesClientResetCount.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.fraPropertiesC.clientResetCount", data.RadioTypeCFraPropertiesClientResetCount.ValueInt64())
	}
	if !data.RadioTypeCFraPropertiesClientUtilizationThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.fraPropertiesC.clientUtilizationThreshold", data.RadioTypeCFraPropertiesClientUtilizationThreshold.ValueInt64())
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.coverageHoleDetectionProperties.chdClientLevel", data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel.ValueInt64())
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.coverageHoleDetectionProperties.chdDataRssiThreshold", data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold", data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.ValueInt64())
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.coverageHoleDetectionProperties.chdExceptionLevel", data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel.ValueInt64())
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect", data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold", data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetect", data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect.ValueBool())
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold", data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.ValueInt64())
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold", data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessRFProfile) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.rfProfileName"); value.Exists() {
		data.RfProfileName = types.StringValue(value.String())
	} else {
		data.RfProfileName = types.StringNull()
	}
	if value := res.Get("response.defaultRfProfile"); value.Exists() {
		data.DefaultRfProfile = types.BoolValue(value.Bool())
	} else {
		data.DefaultRfProfile = types.BoolNull()
	}
	if value := res.Get("response.enableRadioTypeA"); value.Exists() {
		data.EnableRadioTypeA = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeA = types.BoolNull()
	}
	if value := res.Get("response.enableRadioTypeB"); value.Exists() {
		data.EnableRadioTypeB = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeB = types.BoolNull()
	}
	if value := res.Get("response.enableRadioType6GHz"); value.Exists() {
		data.EnableRadioType6GHz = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioType6GHz = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.parentProfile"); value.Exists() {
		data.RadioTypeAParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeAParentProfile = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.radioChannels"); value.Exists() {
		data.RadioTypeARadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeARadioChannels = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.dataRates"); value.Exists() {
		data.RadioTypeADataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeADataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeAMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeAMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeAPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeARxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeARxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeAMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeAMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.channelWidth"); value.Exists() {
		data.RadioTypeAChannelWidth = types.StringValue(value.String())
	} else {
		data.RadioTypeAChannelWidth = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.preamblePuncture"); value.Exists() {
		data.RadioTypeAPreamblePuncture = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAPreamblePuncture = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.zeroWaitDfsEnable"); value.Exists() {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.customRxSopThreshold"); value.Exists() {
		data.RadioTypeACustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.maxRadioClients"); value.Exists() {
		data.RadioTypeAMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.fraPropertiesA.clientAware"); value.Exists() {
		data.RadioTypeAFraPropertiesClientAware = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAFraPropertiesClientAware = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.fraPropertiesA.clientSelect"); value.Exists() {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.fraPropertiesA.clientReset"); value.Exists() {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.parentProfile"); value.Exists() {
		data.RadioTypeBParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeBParentProfile = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.radioChannels"); value.Exists() {
		data.RadioTypeBRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeBRadioChannels = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.dataRates"); value.Exists() {
		data.RadioTypeBDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBDataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeBMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeBPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeBRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeBRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeBMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeBMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.customRxSopThreshold"); value.Exists() {
		data.RadioTypeBCustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.maxRadioClients"); value.Exists() {
		data.RadioTypeBMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.parentProfile"); value.Exists() {
		data.RadioTypeCParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeCParentProfile = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.radioChannels"); value.Exists() {
		data.RadioTypeCRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeCRadioChannels = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.dataRates"); value.Exists() {
		data.RadioTypeCDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCDataRates = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeCMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeCPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeCRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeCRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeCMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeCMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.enableStandardPowerService"); value.Exists() {
		data.RadioTypeCEnableStandardPowerService = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCEnableStandardPowerService = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.ofdmaDownLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.ofdmaUpLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.muMimoUpLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.muMimoDownLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaDownLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaUpLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.muMimoUpLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.muMimoDownLink"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaMultiRu"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.targetWakeTime"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesTargetWakeTime = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesTargetWakeTime = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.twtBroadcastSupport"); value.Exists() {
		data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.preamblePuncture"); value.Exists() {
		data.RadioTypeCPreamblePuncture = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCPreamblePuncture = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.minDbsWidth"); value.Exists() {
		data.RadioTypeCMinDbsWidth = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinDbsWidth = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.maxDbsWidth"); value.Exists() {
		data.RadioTypeCMaxDbsWidth = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxDbsWidth = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.maxRadioClients"); value.Exists() {
		data.RadioTypeCMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.customRxSopThreshold"); value.Exists() {
		data.RadioTypeCCustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.pscEnforcingEnabled"); value.Exists() {
		data.RadioTypeCPscEnforcingEnabled = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCPscEnforcingEnabled = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.discoveryFrames6GHz"); value.Exists() {
		data.RadioTypeCDiscoveryFrames6ghz = types.StringValue(value.String())
	} else {
		data.RadioTypeCDiscoveryFrames6ghz = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.broadcastProbeResponseInterval"); value.Exists() {
		data.RadioTypeCBroadcastProbeResponseInterval = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCBroadcastProbeResponseInterval = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.fraPropertiesC.clientResetCount"); value.Exists() {
		data.RadioTypeCFraPropertiesClientResetCount = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCFraPropertiesClientResetCount = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.fraPropertiesC.clientUtilizationThreshold"); value.Exists() {
		data.RadioTypeCFraPropertiesClientUtilizationThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCFraPropertiesClientUtilizationThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessRFProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rfProfileName"); value.Exists() && !data.RfProfileName.IsNull() {
		data.RfProfileName = types.StringValue(value.String())
	} else {
		data.RfProfileName = types.StringNull()
	}
	if value := res.Get("response.defaultRfProfile"); value.Exists() && !data.DefaultRfProfile.IsNull() {
		data.DefaultRfProfile = types.BoolValue(value.Bool())
	} else {
		data.DefaultRfProfile = types.BoolNull()
	}
	if value := res.Get("response.enableRadioTypeA"); value.Exists() && !data.EnableRadioTypeA.IsNull() {
		data.EnableRadioTypeA = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeA = types.BoolNull()
	}
	if value := res.Get("response.enableRadioTypeB"); value.Exists() && !data.EnableRadioTypeB.IsNull() {
		data.EnableRadioTypeB = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeB = types.BoolNull()
	}
	if value := res.Get("response.enableRadioType6GHz"); value.Exists() && !data.EnableRadioType6GHz.IsNull() {
		data.EnableRadioType6GHz = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioType6GHz = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.parentProfile"); value.Exists() && !data.RadioTypeAParentProfile.IsNull() {
		data.RadioTypeAParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeAParentProfile = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.radioChannels"); value.Exists() && !data.RadioTypeARadioChannels.IsNull() {
		data.RadioTypeARadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeARadioChannels = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.dataRates"); value.Exists() && !data.RadioTypeADataRates.IsNull() {
		data.RadioTypeADataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeADataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeAMandatoryDataRates.IsNull() {
		data.RadioTypeAMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeAMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeAPowerThresholdV1.IsNull() {
		data.RadioTypeAPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeARxSopThreshold.IsNull() {
		data.RadioTypeARxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeARxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.minPowerLevel"); value.Exists() && !data.RadioTypeAMinPowerLevel.IsNull() {
		data.RadioTypeAMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeAMaxPowerLevel.IsNull() {
		data.RadioTypeAMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.channelWidth"); value.Exists() && !data.RadioTypeAChannelWidth.IsNull() {
		data.RadioTypeAChannelWidth = types.StringValue(value.String())
	} else {
		data.RadioTypeAChannelWidth = types.StringNull()
	}
	if value := res.Get("response.radioTypeAProperties.preamblePuncture"); value.Exists() && !data.RadioTypeAPreamblePuncture.IsNull() {
		data.RadioTypeAPreamblePuncture = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAPreamblePuncture = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.zeroWaitDfsEnable"); value.Exists() && !data.RadioTypeAZeroWaitDfsEnable.IsNull() {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAZeroWaitDfsEnable = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.customRxSopThreshold"); value.Exists() && !data.RadioTypeACustomRxSopThreshold.IsNull() {
		data.RadioTypeACustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.maxRadioClients"); value.Exists() && !data.RadioTypeAMaxRadioClients.IsNull() {
		data.RadioTypeAMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.fraPropertiesA.clientAware"); value.Exists() && !data.RadioTypeAFraPropertiesClientAware.IsNull() {
		data.RadioTypeAFraPropertiesClientAware = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeAFraPropertiesClientAware = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.fraPropertiesA.clientSelect"); value.Exists() && !data.RadioTypeAFraPropertiesClientSelect.IsNull() {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientSelect = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.fraPropertiesA.clientReset"); value.Exists() && !data.RadioTypeAFraPropertiesClientReset.IsNull() {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAFraPropertiesClientReset = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() && !data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeAProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.parentProfile"); value.Exists() && !data.RadioTypeBParentProfile.IsNull() {
		data.RadioTypeBParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeBParentProfile = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.radioChannels"); value.Exists() && !data.RadioTypeBRadioChannels.IsNull() {
		data.RadioTypeBRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeBRadioChannels = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.dataRates"); value.Exists() && !data.RadioTypeBDataRates.IsNull() {
		data.RadioTypeBDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBDataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeBMandatoryDataRates.IsNull() {
		data.RadioTypeBMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeBPowerThresholdV1.IsNull() {
		data.RadioTypeBPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeBRxSopThreshold.IsNull() {
		data.RadioTypeBRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeBRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.radioTypeBProperties.minPowerLevel"); value.Exists() && !data.RadioTypeBMinPowerLevel.IsNull() {
		data.RadioTypeBMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeBMaxPowerLevel.IsNull() {
		data.RadioTypeBMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.customRxSopThreshold"); value.Exists() && !data.RadioTypeBCustomRxSopThreshold.IsNull() {
		data.RadioTypeBCustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.maxRadioClients"); value.Exists() && !data.RadioTypeBMaxRadioClients.IsNull() {
		data.RadioTypeBMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() && !data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() && !data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() && !data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() && !data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() && !data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() && !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() && !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioTypeBProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.parentProfile"); value.Exists() && !data.RadioTypeCParentProfile.IsNull() {
		data.RadioTypeCParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeCParentProfile = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.radioChannels"); value.Exists() && !data.RadioTypeCRadioChannels.IsNull() {
		data.RadioTypeCRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeCRadioChannels = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.dataRates"); value.Exists() && !data.RadioTypeCDataRates.IsNull() {
		data.RadioTypeCDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCDataRates = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeCMandatoryDataRates.IsNull() {
		data.RadioTypeCMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeCPowerThresholdV1.IsNull() {
		data.RadioTypeCPowerThresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCPowerThresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeCRxSopThreshold.IsNull() {
		data.RadioTypeCRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeCRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.minPowerLevel"); value.Exists() && !data.RadioTypeCMinPowerLevel.IsNull() {
		data.RadioTypeCMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeCMaxPowerLevel.IsNull() {
		data.RadioTypeCMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.enableStandardPowerService"); value.Exists() && !data.RadioTypeCEnableStandardPowerService.IsNull() {
		data.RadioTypeCEnableStandardPowerService = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCEnableStandardPowerService = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.ofdmaDownLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.ofdmaUpLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.muMimoUpLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11axParameters.muMimoDownLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaDownLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaUpLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.muMimoUpLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.muMimoDownLink"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.dot11beParameters.ofdmaMultiRu"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu.IsNull() {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.targetWakeTime"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesTargetWakeTime.IsNull() {
		data.RadioTypeCMultiBssidPropertiesTargetWakeTime = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesTargetWakeTime = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.multiBssidProperties.twtBroadcastSupport"); value.Exists() && !data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport.IsNull() {
		data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.preamblePuncture"); value.Exists() && !data.RadioTypeCPreamblePuncture.IsNull() {
		data.RadioTypeCPreamblePuncture = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCPreamblePuncture = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.minDbsWidth"); value.Exists() && !data.RadioTypeCMinDbsWidth.IsNull() {
		data.RadioTypeCMinDbsWidth = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinDbsWidth = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.maxDbsWidth"); value.Exists() && !data.RadioTypeCMaxDbsWidth.IsNull() {
		data.RadioTypeCMaxDbsWidth = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxDbsWidth = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.maxRadioClients"); value.Exists() && !data.RadioTypeCMaxRadioClients.IsNull() {
		data.RadioTypeCMaxRadioClients = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxRadioClients = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.customRxSopThreshold"); value.Exists() && !data.RadioTypeCCustomRxSopThreshold.IsNull() {
		data.RadioTypeCCustomRxSopThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCustomRxSopThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.pscEnforcingEnabled"); value.Exists() && !data.RadioTypeCPscEnforcingEnabled.IsNull() {
		data.RadioTypeCPscEnforcingEnabled = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCPscEnforcingEnabled = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.discoveryFrames6GHz"); value.Exists() && !data.RadioTypeCDiscoveryFrames6ghz.IsNull() {
		data.RadioTypeCDiscoveryFrames6ghz = types.StringValue(value.String())
	} else {
		data.RadioTypeCDiscoveryFrames6ghz = types.StringNull()
	}
	if value := res.Get("response.radioType6GHzProperties.broadcastProbeResponseInterval"); value.Exists() && !data.RadioTypeCBroadcastProbeResponseInterval.IsNull() {
		data.RadioTypeCBroadcastProbeResponseInterval = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCBroadcastProbeResponseInterval = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.fraPropertiesC.clientResetCount"); value.Exists() && !data.RadioTypeCFraPropertiesClientResetCount.IsNull() {
		data.RadioTypeCFraPropertiesClientResetCount = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCFraPropertiesClientResetCount = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.fraPropertiesC.clientUtilizationThreshold"); value.Exists() && !data.RadioTypeCFraPropertiesClientUtilizationThreshold.IsNull() {
		data.RadioTypeCFraPropertiesClientUtilizationThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCFraPropertiesClientUtilizationThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdClientLevel"); value.Exists() && !data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdDataRssiThreshold"); value.Exists() && !data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdVoiceRssiThreshold"); value.Exists() && !data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.coverageHoleDetectionProperties.chdExceptionLevel"); value.Exists() && !data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetect"); value.Exists() && !data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axNonSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetect"); value.Exists() && !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolValue(value.Bool())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect = types.BoolNull()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMinThreshold"); value.Exists() && !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold = types.Int64Null()
	}
	if value := res.Get("response.radioType6GHzProperties.spatialReuseProperties.dot11axSrgObssPacketDetectMaxThreshold"); value.Exists() && !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold = types.Int64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessRFProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DefaultRfProfile.IsNull() {
		return false
	}
	if !data.EnableRadioTypeA.IsNull() {
		return false
	}
	if !data.EnableRadioTypeB.IsNull() {
		return false
	}
	if !data.EnableRadioType6GHz.IsNull() {
		return false
	}
	if !data.RadioTypeAParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeARadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeADataRates.IsNull() {
		return false
	}
	if !data.RadioTypeAMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeAPowerThresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeARxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeAMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeAMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeAChannelWidth.IsNull() {
		return false
	}
	if !data.RadioTypeAPreamblePuncture.IsNull() {
		return false
	}
	if !data.RadioTypeAZeroWaitDfsEnable.IsNull() {
		return false
	}
	if !data.RadioTypeACustomRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeAMaxRadioClients.IsNull() {
		return false
	}
	if !data.RadioTypeAFraPropertiesClientAware.IsNull() {
		return false
	}
	if !data.RadioTypeAFraPropertiesClientSelect.IsNull() {
		return false
	}
	if !data.RadioTypeAFraPropertiesClientReset.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeACoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeASpartialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeBRadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeBDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeBMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeBPowerThresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeBRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBCustomRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBMaxRadioClients.IsNull() {
		return false
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBCoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeCRadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeCDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeCMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeCPowerThresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeCRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCEnableStandardPowerService.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaDownLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersOfdmaUpLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoUpLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11axParametersMuMimoDownLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaDownLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaUpLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoUpLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersMuMimoDownLink.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesDot11beParametersOfdmaMultiRu.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesTargetWakeTime.IsNull() {
		return false
	}
	if !data.RadioTypeCMultiBssidPropertiesTwtBroadcastSupport.IsNull() {
		return false
	}
	if !data.RadioTypeCPreamblePuncture.IsNull() {
		return false
	}
	if !data.RadioTypeCMinDbsWidth.IsNull() {
		return false
	}
	if !data.RadioTypeCMaxDbsWidth.IsNull() {
		return false
	}
	if !data.RadioTypeCMaxRadioClients.IsNull() {
		return false
	}
	if !data.RadioTypeCCustomRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCPscEnforcingEnabled.IsNull() {
		return false
	}
	if !data.RadioTypeCDiscoveryFrames6ghz.IsNull() {
		return false
	}
	if !data.RadioTypeCBroadcastProbeResponseInterval.IsNull() {
		return false
	}
	if !data.RadioTypeCFraPropertiesClientResetCount.IsNull() {
		return false
	}
	if !data.RadioTypeCFraPropertiesClientUtilizationThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdClientLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdDataRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdVoiceRssiThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCCoverageHoleDetectionPropertiesChdExceptionLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axNonSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetect.IsNull() {
		return false
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMinThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCSpatialReusePropertiesDot11axSrgObssPacketDetectMaxThreshold.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
