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
func TestAccCcWirelessSSID(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "ssid", "mySSID1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "auth_type", "WPA3_PERSONAL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "fast_lane", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "mac_filtering", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "ssid_radio_type", "Triple band operation(2.4GHz, 5GHz and 6GHz)"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "broadcast_ssid", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "fast_transition", "ADAPTIVE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "session_timeout_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "session_timeout", "1800"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "client_exclusion", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "client_exclusion_timeout", "1800"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "basic_service_set_max_idle", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "basic_service_set_client_idle_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "directed_multicast_service", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "neighbor_list", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "mft_client_protection", "OPTIONAL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "aaa_override", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "protected_management_frame", "REQUIRED"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "rsn_cipher_suite_ccmp128", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "wlan_type", "Enterprise"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "auth_key_sae_ext", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "ghz24_policy", "dot11-g-only"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "hex", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_ssid.test", "random_mac_filter", "false"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessSSIDPrerequisitesConfig + testAccCcWirelessSSIDConfig_all(),
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
const testAccCcWirelessSSIDPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcWirelessSSIDConfig_minimum() string {
	config := `resource "catalystcenter_wireless_ssid" "test" {` + "\n"
	config += `	site_id = data.catalystcenter_site.test.id` + "\n"
	config += `	ssid = "mySSID1"` + "\n"
	config += `	auth_type = "WPA3_PERSONAL"` + "\n"
	config += `	wlan_type = "Enterprise"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcWirelessSSIDConfig_all() string {
	config := `resource "catalystcenter_wireless_ssid" "test" {` + "\n"
	config += `	site_id = data.catalystcenter_site.test.id` + "\n"
	config += `	ssid = "mySSID1"` + "\n"
	config += `	auth_type = "WPA3_PERSONAL"` + "\n"
	config += `	passphrase = "Cisco123"` + "\n"
	config += `	fast_lane = false` + "\n"
	config += `	mac_filtering = false` + "\n"
	config += `	ssid_radio_type = "Triple band operation(2.4GHz, 5GHz and 6GHz)"` + "\n"
	config += `	broadcast_ssid = true` + "\n"
	config += `	fast_transition = "ADAPTIVE"` + "\n"
	config += `	session_timeout_enable = true` + "\n"
	config += `	session_timeout = 1800` + "\n"
	config += `	client_exclusion = true` + "\n"
	config += `	client_exclusion_timeout = 1800` + "\n"
	config += `	basic_service_set_max_idle = true` + "\n"
	config += `	basic_service_set_client_idle_timeout = 300` + "\n"
	config += `	directed_multicast_service = true` + "\n"
	config += `	neighbor_list = true` + "\n"
	config += `	mft_client_protection = "OPTIONAL"` + "\n"
	config += `	aaa_override = false` + "\n"
	config += `	protected_management_frame = "REQUIRED"` + "\n"
	config += `	rsn_cipher_suite_ccmp128 = true` + "\n"
	config += `	wlan_type = "Enterprise"` + "\n"
	config += `	auth_key_sae_ext = true` + "\n"
	config += `	ghz24_policy = "dot11-g-only"` + "\n"
	config += `	hex = false` + "\n"
	config += `	random_mac_filter = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
