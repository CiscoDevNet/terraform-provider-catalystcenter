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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcWirelessRFProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "rf_profile_name", "RF_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "default_rf_profile", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "enable_radio_type_a", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "enable_radio_type_b", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "enable_radio_type6_g_hz", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_parent_profile", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_radio_channels", "36,40,44,48,52,56,60,64,144,149,153,157,161,165,169,173"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_data_rates", "6,9,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_mandatory_data_rates", "12,24"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_power_threshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_rx_sop_threshold", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_max_power_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_channel_width", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_preamble_puncture", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_zero_wait_dfs_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_custom_rx_sop_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_max_radio_clients", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_fra_properties_client_aware", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_fra_properties_client_select", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_fra_properties_client_reset", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_coverage_hole_detection_properties_chd_client_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold", "-65"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_coverage_hole_detection_properties_chd_exception_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_parent_profile", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_radio_channels", "1,2,3,4,5,6,7,8,9,10,11,12,13,14"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_data_rates", "1,2,5.5,6,9,11,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_mandatory_data_rates", "1,2"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_power_threshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_rx_sop_threshold", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_max_power_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_custom_rx_sop_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_max_radio_clients", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_coverage_hole_detection_properties_chd_client_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_coverage_hole_detection_properties_chd_data_rssi_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_coverage_hole_detection_properties_chd_voice_rssi_threshold", "-65"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_coverage_hole_detection_properties_chd_exception_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_parent_profile", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_radio_channels", "1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229,233"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_data_rates", "6,9,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_mandatory_data_rates", "6,9"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_power_threshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_rx_sop_threshold", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_max_power_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_enable_standard_power_service", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_down_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_up_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_up_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_down_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_down_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_up_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_up_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_down_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_multi_ru", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_target_wake_time", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_multi_bssid_properties_twt_broadcast_support", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_preamble_puncture", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_min_dbs_width", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_max_dbs_width", "160"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_max_radio_clients", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_custom_rx_sop_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_psc_enforcing_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_discovery_frames_6ghz", "Broadcast Probe Response"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_broadcast_probe_response_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_fra_properties_client_reset_count", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_fra_properties_client_utilization_threshold", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_coverage_hole_detection_properties_chd_client_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_coverage_hole_detection_properties_chd_data_rssi_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_coverage_hole_detection_properties_chd_voice_rssi_threshold", "-65"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_coverage_hole_detection_properties_chd_exception_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold", "-70"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessRFProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcWirelessRFProfileConfig_minimum() string {
	config := `resource "catalystcenter_wireless_rf_profile" "test" {` + "\n"
	config += `	rf_profile_name = "RF_Profile_1"` + "\n"
	config += `	default_rf_profile = false` + "\n"
	config += `	enable_radio_type_a = true` + "\n"
	config += `	enable_radio_type_b = true` + "\n"
	config += `	enable_radio_type6_g_hz = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcWirelessRFProfileConfig_all() string {
	config := `resource "catalystcenter_wireless_rf_profile" "test" {` + "\n"
	config += `	rf_profile_name = "RF_Profile_1"` + "\n"
	config += `	default_rf_profile = false` + "\n"
	config += `	enable_radio_type_a = true` + "\n"
	config += `	enable_radio_type_b = true` + "\n"
	config += `	enable_radio_type6_g_hz = false` + "\n"
	config += `	radio_type_a_parent_profile = "CUSTOM"` + "\n"
	config += `	radio_type_a_radio_channels = "36,40,44,48,52,56,60,64,144,149,153,157,161,165,169,173"` + "\n"
	config += `	radio_type_a_data_rates = "6,9,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_a_mandatory_data_rates = "12,24"` + "\n"
	config += `	radio_type_a_power_threshold_v1 = -60` + "\n"
	config += `	radio_type_a_rx_sop_threshold = "CUSTOM"` + "\n"
	config += `	radio_type_a_min_power_level = 8` + "\n"
	config += `	radio_type_a_max_power_level = 20` + "\n"
	config += `	radio_type_a_channel_width = "20"` + "\n"
	config += `	radio_type_a_preamble_puncture = false` + "\n"
	config += `	radio_type_a_zero_wait_dfs_enable = false` + "\n"
	config += `	radio_type_a_custom_rx_sop_threshold = -70` + "\n"
	config += `	radio_type_a_max_radio_clients = 200` + "\n"
	config += `	radio_type_a_fra_properties_client_aware = false` + "\n"
	config += `	radio_type_a_fra_properties_client_select = 50` + "\n"
	config += `	radio_type_a_fra_properties_client_reset = 5` + "\n"
	config += `	radio_type_a_coverage_hole_detection_properties_chd_client_level = 10` + "\n"
	config += `	radio_type_a_coverage_hole_detection_properties_chd_data_rssi_threshold = -70` + "\n"
	config += `	radio_type_a_coverage_hole_detection_properties_chd_voice_rssi_threshold = -65` + "\n"
	config += `	radio_type_a_coverage_hole_detection_properties_chd_exception_level = 20` + "\n"
	config += `	radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect = false` + "\n"
	config += `	radio_type_a_spartial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold = -70` + "\n"
	config += `	radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect = false` + "\n"
	config += `	radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold = -70` + "\n"
	config += `	radio_type_a_spartial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold = -70` + "\n"
	config += `	radio_type_b_parent_profile = "CUSTOM"` + "\n"
	config += `	radio_type_b_radio_channels = "1,2,3,4,5,6,7,8,9,10,11,12,13,14"` + "\n"
	config += `	radio_type_b_data_rates = "1,2,5.5,6,9,11,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_b_mandatory_data_rates = "1,2"` + "\n"
	config += `	radio_type_b_power_threshold_v1 = -60` + "\n"
	config += `	radio_type_b_rx_sop_threshold = "CUSTOM"` + "\n"
	config += `	radio_type_b_min_power_level = 8` + "\n"
	config += `	radio_type_b_max_power_level = 20` + "\n"
	config += `	radio_type_b_custom_rx_sop_threshold = -70` + "\n"
	config += `	radio_type_b_max_radio_clients = 200` + "\n"
	config += `	radio_type_b_coverage_hole_detection_properties_chd_client_level = 10` + "\n"
	config += `	radio_type_b_coverage_hole_detection_properties_chd_data_rssi_threshold = -70` + "\n"
	config += `	radio_type_b_coverage_hole_detection_properties_chd_voice_rssi_threshold = -65` + "\n"
	config += `	radio_type_b_coverage_hole_detection_properties_chd_exception_level = 20` + "\n"
	config += `	radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect = false` + "\n"
	config += `	radio_type_b_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold = -70` + "\n"
	config += `	radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect = false` + "\n"
	config += `	radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold = -70` + "\n"
	config += `	radio_type_b_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold = -70` + "\n"
	config += `	radio_type_c_parent_profile = "CUSTOM"` + "\n"
	config += `	radio_type_c_radio_channels = "1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125,129,133,137,141,145,149,153,157,161,165,169,173,177,181,185,189,193,197,201,205,209,213,217,221,225,229,233"` + "\n"
	config += `	radio_type_c_data_rates = "6,9,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_c_mandatory_data_rates = "6,9"` + "\n"
	config += `	radio_type_c_power_threshold_v1 = -60` + "\n"
	config += `	radio_type_c_rx_sop_threshold = "CUSTOM"` + "\n"
	config += `	radio_type_c_min_power_level = 8` + "\n"
	config += `	radio_type_c_max_power_level = 20` + "\n"
	config += `	radio_type_c_enable_standard_power_service = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_down_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11ax_parameters_ofdma_up_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_up_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11ax_parameters_mu_mimo_down_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_down_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_up_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_up_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11be_parameters_mu_mimo_down_link = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_dot11be_parameters_ofdma_multi_ru = true` + "\n"
	config += `	radio_type_c_multi_bssid_properties_target_wake_time = false` + "\n"
	config += `	radio_type_c_multi_bssid_properties_twt_broadcast_support = false` + "\n"
	config += `	radio_type_c_preamble_puncture = false` + "\n"
	config += `	radio_type_c_min_dbs_width = 20` + "\n"
	config += `	radio_type_c_max_dbs_width = 160` + "\n"
	config += `	radio_type_c_max_radio_clients = 200` + "\n"
	config += `	radio_type_c_custom_rx_sop_threshold = -70` + "\n"
	config += `	radio_type_c_psc_enforcing_enabled = false` + "\n"
	config += `	radio_type_c_discovery_frames_6ghz = "Broadcast Probe Response"` + "\n"
	config += `	radio_type_c_broadcast_probe_response_interval = 10` + "\n"
	config += `	radio_type_c_fra_properties_client_reset_count = 2` + "\n"
	config += `	radio_type_c_fra_properties_client_utilization_threshold = 80` + "\n"
	config += `	radio_type_c_coverage_hole_detection_properties_chd_client_level = 10` + "\n"
	config += `	radio_type_c_coverage_hole_detection_properties_chd_data_rssi_threshold = -70` + "\n"
	config += `	radio_type_c_coverage_hole_detection_properties_chd_voice_rssi_threshold = -65` + "\n"
	config += `	radio_type_c_coverage_hole_detection_properties_chd_exception_level = 20` + "\n"
	config += `	radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect = false` + "\n"
	config += `	radio_type_c_spatial_reuse_properties_dot11ax_non_srg_obss_packet_detect_max_threshold = -70` + "\n"
	config += `	radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect = false` + "\n"
	config += `	radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_min_threshold = -70` + "\n"
	config += `	radio_type_c_spatial_reuse_properties_dot11ax_srg_obss_packet_detect_max_threshold = -70` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
