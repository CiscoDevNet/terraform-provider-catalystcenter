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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcWirelessProfileSiteTag(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_profile_site_tag.test", "site_tag_name", "SiteTag1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcWirelessProfileSiteTagPrerequisitesConfig + testAccDataSourceCcWirelessProfileSiteTagConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcWirelessProfileSiteTagPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test_area" {
  name      = "SiteTagTestArea"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_building" "test_building" {
  name      = "SiteTagTestBuilding"
  parent_id = catalystcenter_area.test_area.id
  country   = "United States"
  address   = "123 Test Street, San Jose, CA"
  latitude  = 37.3382
  longitude = -121.8863
}
resource "catalystcenter_wireless_ssid" "test" {
  site_id                               = data.catalystcenter_site.test.id
  ssid                                  = "SiteTagTestSSID"
  auth_type                             = "WPA3_PERSONAL"
  passphrase                            = "Cisco12345"
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
  wireless_profile_name = "SiteTagTestProfile"
  ssid_details = [{
    ssid_name           = catalystcenter_wireless_ssid.test.ssid
    enable_fabric       = false
    enable_flex_connect = false
  }]

  depends_on = [catalystcenter_wireless_ssid.test, catalystcenter_ap_profile.test]
}
resource "catalystcenter_ap_profile" "test" {
  ap_profile_name = "SiteTagTestAPProfile"
}
resource "catalystcenter_network_profile_for_sites_assignments" "test" {
  network_profile_id = catalystcenter_wireless_profile.test.id
  items = [{
    id = catalystcenter_building.test_building.id
  }]

  depends_on = [catalystcenter_building.test_building]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcWirelessProfileSiteTagConfig() string {
	config := `resource "catalystcenter_wireless_profile_site_tag" "test" {` + "\n"
	config += `	wireless_profile_id = catalystcenter_network_profile_for_sites_assignments.test.network_profile_id` + "\n"
	config += `	site_tag_name = "SiteTag1"` + "\n"
	config += `	ap_profile_name = catalystcenter_ap_profile.test.ap_profile_name` + "\n"
	config += `	site_ids = [catalystcenter_building.test_building.id]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_wireless_profile_site_tag" "test" {
			id = catalystcenter_wireless_profile_site_tag.test.id
			wireless_profile_id = catalystcenter_network_profile_for_sites_assignments.test.network_profile_id
			site_tag_id = catalystcenter_wireless_profile_site_tag.test.site_tag_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
