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
func TestAccDataSourceCcUpdateAuthenticationProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_update_authentication_profile.test", "authentication_profile_name", "Open Authentication"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_update_authentication_profile.test", "authentication_order", "mac"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_update_authentication_profile.test", "dot1x_to_mab_fallback_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_update_authentication_profile.test", "wake_on_lan", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_update_authentication_profile.test", "number_of_hosts", "Unlimited"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcUpdateAuthenticationProfilePrerequisitesConfig + testAccDataSourceCcUpdateAuthenticationProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcUpdateAuthenticationProfilePrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  authentication_profile_name = "Open Authentication"
  pub_sub_enabled             = false
  depends_on                  = [catalystcenter_aaa_settings.test]
}
resource "catalystcenter_aaa_settings" "test" {
  site_id                       = catalystcenter_area.test.id
  network_aaa_server_type       = "ISE"
  network_aaa_protocol          = "RADIUS"
  network_aaa_pan               = "10.62.190.86"
  network_aaa_primary_server_ip = "10.62.190.86"
  client_aaa_server_type        = "ISE"
  client_aaa_protocol           = "RADIUS"
  client_aaa_pan                = "10.62.190.86"
  client_aaa_primary_server_ip  = "10.62.190.86"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcUpdateAuthenticationProfileConfig() string {
	config := `resource "catalystcenter_update_authentication_profile" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	authentication_profile_name = "Open Authentication"` + "\n"
	config += `	authentication_order = "mac"` + "\n"
	config += `	dot1x_to_mab_fallback_timeout = 10` + "\n"
	config += `	wake_on_lan = true` + "\n"
	config += `	number_of_hosts = "Unlimited"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_update_authentication_profile" "test" {
			fabric_id = catalystcenter_fabric_site.test.id
			authentication_profile_name = "Open Authentication"
			depends_on = [catalystcenter_update_authentication_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
