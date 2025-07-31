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
func TestAccCcAssignManagedAPLocations(t *testing.T) {
	if os.Getenv("WIRELESS") == "" {
		t.Skip("skipping test, set environment variable WIRELESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_assign_managed_ap_locations.test", "device_id", "75b0f85a-8157-4db3-ae2d-9807c893319a"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcAssignManagedAPLocationsPrerequisitesConfig + testAccCcAssignManagedAPLocationsConfig_all(),
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
const testAccCcAssignManagedAPLocationsPrerequisitesConfig = `
data "catalystcenter_area" "test" {
  name = "Global"
}
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_building" "test" {
  name        = "Building1"
  parent_name = "Global/Area1"
  country     = "United States"
  address     = "150 W Tasman Dr, San Jose"
  latitude    = 37.338
  longitude   = -121.832
  depends_on = [catalystcenter_area.test]
}
resource "catalystcenter_wireless_ssid" "test" {
  site_id                               = catalystcenter_area.test.id
  ssid                                  = "mySSID1"
  auth_type                             = "WPA3_PERSONAL"
  passphrase                            = "Cisco123"
  fast_lane                             = false
  mac_filtering                         = false
  ssid_radio_type                       = "Triple band operation(2.4GHz, 5GHz and 6GHz)"
  broadcast_ssid                        = true
  fast_transition                       = "ADAPTIVE"
  session_timeout_enable                = true
  session_timeout                       = 1800
  client_exclusion                      = true
  client_exclusion_timeout              = 1800
  basic_service_set_max_idle            = true
  basic_service_set_client_idle_timeout = 300
  directed_multicast_service            = true
  neighbor_list                         = true
  mft_client_protection                 = "OPTIONAL"
  aaa_override                          = false
  protected_management_frame            = "REQUIRED"
  rsn_cipher_suite_ccmp128              = true
  wlan_type                             = "Enterprise"
  auth_key_sae_ext                      = true
  ghz24_policy                          = "dot11-g-only"
  hex                                   = false
  random_mac_filter                     = false
}
resource "catalystcenter_wireless_profile" "test" {
  wireless_profile_name = "Wireless_Profile_1"
  ssid_details = [
    {
      ssid_name           = "mySSID1"
      enable_fabric       = false
      enable_flex_connect = false
    }
  ]
}
resource "catalystcenter_associate_site_to_network_profile" "test" {
  network_profile_id = catalystcenter_wireless_profile.test.id
  site_id            = catalystcenter_building.test.id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAssignManagedAPLocationsConfig_minimum() string {
	config := `resource "catalystcenter_assign_managed_ap_locations" "test" {` + "\n"
	config += `	device_id = "75b0f85a-8157-4db3-ae2d-9807c893319a"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAssignManagedAPLocationsConfig_all() string {
	config := `resource "catalystcenter_assign_managed_ap_locations" "test" {` + "\n"
	config += `	primary_managed_ap_locations_site_ids = [catalystcenter_building.test.id]` + "\n"
	config += `	device_id = "75b0f85a-8157-4db3-ae2d-9807c893319a"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
