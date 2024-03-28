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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

func TestAccDataSourceCcWirelessRFProfile(t *testing.T) {
	if os.Getenv("WIRELESS_RF_PROFILE") == "" {
		t.Skip("skipping test, set environment variable WIRELESS_RF_PROFILE")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "name", "RF_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "default_rf_profile", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_radio_type_a", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_radio_type_b", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_radio_type_c", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "channel_width", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_custom", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_brown_field", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_parent_profile", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_radio_channels", "36,40,44,48,52,56,60,64,144,149,153,157,161,165,169,173"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_data_rates", "6,9,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_mandatory_data_rates", "12,24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_power_treshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_rx_sop_threshold", "LOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_max_power_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_parent_profile", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_radio_channels", "1,6,11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_data_rates", "9,11,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_mandatory_data_rates", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_power_treshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_rx_sop_threshold", "LOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_max_power_level", "20"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcWirelessRFProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcWirelessRFProfileConfig() string {
	config := `resource "catalystcenter_wireless_rf_profile" "test" {` + "\n"
	config += `	name = "RF_Profile_1"` + "\n"
	config += `	default_rf_profile = false` + "\n"
	config += `	enable_radio_type_a = true` + "\n"
	config += `	enable_radio_type_b = true` + "\n"
	config += `	enable_radio_type_c = false` + "\n"
	config += `	channel_width = "20"` + "\n"
	config += `	enable_custom = true` + "\n"
	config += `	enable_brown_field = false` + "\n"
	config += `	radio_type_a_parent_profile = "CUSTOM"` + "\n"
	config += `	radio_type_a_radio_channels = "36,40,44,48,52,56,60,64,144,149,153,157,161,165,169,173"` + "\n"
	config += `	radio_type_a_data_rates = "6,9,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_a_mandatory_data_rates = "12,24"` + "\n"
	config += `	radio_type_a_power_treshold_v1 = -60` + "\n"
	config += `	radio_type_a_rx_sop_threshold = "LOW"` + "\n"
	config += `	radio_type_a_min_power_level = 8` + "\n"
	config += `	radio_type_a_max_power_level = 20` + "\n"
	config += `	radio_type_b_parent_profile = "CUSTOM"` + "\n"
	config += `	radio_type_b_radio_channels = "1,6,11"` + "\n"
	config += `	radio_type_b_data_rates = "9,11,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_b_mandatory_data_rates = "12"` + "\n"
	config += `	radio_type_b_power_treshold_v1 = -60` + "\n"
	config += `	radio_type_b_rx_sop_threshold = "LOW"` + "\n"
	config += `	radio_type_b_min_power_level = 8` + "\n"
	config += `	radio_type_b_max_power_level = 20` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_wireless_rf_profile" "test" {
			id = catalystcenter_wireless_rf_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
