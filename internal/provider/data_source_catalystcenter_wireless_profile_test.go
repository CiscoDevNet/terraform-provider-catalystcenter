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
func TestAccDataSourceCcWirelessProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_profile.test", "wireless_profile_name", "Wireless_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_profile.test", "ssid_details.0.enable_fabric", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_profile.test", "ssid_details.0.enable_flex_connect", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcWirelessProfilePrerequisitesConfig + testAccDataSourceCcWirelessProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcWirelessProfilePrerequisitesConfig = `
resource "catalystcenter_wireless_enterprise_ssid" "test" {
  name                                  = "mySSID1"
  security_level                        = "wpa3_enterprise"
  passphrase                            = "Cisco123"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcWirelessProfileConfig() string {
	config := `resource "catalystcenter_wireless_profile" "test" {` + "\n"
	config += `	wireless_profile_name = "Wireless_Profile_1"` + "\n"
	config += `	ssid_details = [{` + "\n"
	config += `	  ssid_name = catalystcenter_wireless_enterprise_ssid.test.name` + "\n"
	config += `	  enable_fabric = false` + "\n"
	config += `	  enable_flex_connect = false` + "\n"
	config += `	}]` + "\n"
	config += `	ap_zones = [{` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_wireless_profile" "test" {
			wireless_profile_name = "Wireless_Profile_1"
			depends_on = [catalystcenter_wireless_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
