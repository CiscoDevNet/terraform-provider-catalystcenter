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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcWirelessCleanAirConfiguration(t *testing.T) {
	if os.Getenv("CC315") == "" {
		t.Skip("skipping test, set environment variable CC315")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "design_name", "cleanAir_802_11b_custom"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "radio_band", "2_4GHZ"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "clean_air", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "clean_air_device_reporting", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "persistent_device_propagation", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "ble_beacon", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "bluetooth_paging_inquiry", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "bluetooth_sco_acl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "continuous_transmitter", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "generic_dect", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "generic_tdd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "jammer", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "microwave_oven", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "motorola_canopy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "si_fhss", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "spectrum_80211_fh", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "spectrum_80211_non_standard_channel", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "spectrum_802154", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "spectrum_inverted", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "super_ag", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "video_camera", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "wimax_fixed", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "wimax_mobile", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_cleanair_configuration.test", "xbox", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcWirelessCleanAirConfigurationConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessCleanAirConfigurationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_wireless_cleanair_configuration.test",
		ImportState:  true,
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
func testAccCcWirelessCleanAirConfigurationConfig_minimum() string {
	config := `resource "catalystcenter_wireless_cleanair_configuration" "test" {` + "\n"
	config += `	design_name = "cleanAir_802_11b_custom"` + "\n"
	config += `	radio_band = "2_4GHZ"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcWirelessCleanAirConfigurationConfig_all() string {
	config := `resource "catalystcenter_wireless_cleanair_configuration" "test" {` + "\n"
	config += `	design_name = "cleanAir_802_11b_custom"` + "\n"
	config += `	radio_band = "2_4GHZ"` + "\n"
	config += `	clean_air = true` + "\n"
	config += `	clean_air_device_reporting = true` + "\n"
	config += `	persistent_device_propagation = true` + "\n"
	config += `	ble_beacon = true` + "\n"
	config += `	bluetooth_paging_inquiry = true` + "\n"
	config += `	bluetooth_sco_acl = true` + "\n"
	config += `	continuous_transmitter = true` + "\n"
	config += `	generic_dect = true` + "\n"
	config += `	generic_tdd = true` + "\n"
	config += `	jammer = true` + "\n"
	config += `	microwave_oven = true` + "\n"
	config += `	motorola_canopy = true` + "\n"
	config += `	si_fhss = true` + "\n"
	config += `	spectrum_80211_fh = true` + "\n"
	config += `	spectrum_80211_non_standard_channel = true` + "\n"
	config += `	spectrum_802154 = true` + "\n"
	config += `	spectrum_inverted = true` + "\n"
	config += `	super_ag = true` + "\n"
	config += `	video_camera = true` + "\n"
	config += `	wimax_fixed = true` + "\n"
	config += `	wimax_mobile = true` + "\n"
	config += `	xbox = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
