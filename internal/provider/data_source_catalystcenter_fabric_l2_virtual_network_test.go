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
func TestAccDataSourceCcFabricL2VirtualNetwork(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_virtual_network.test", "vlan_name", "VLAN401"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_virtual_network.test", "vlan_id", "401"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_virtual_network.test", "traffic_type", "DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l2_virtual_network.test", "fabric_enabled_wireless", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricL2VirtualNetworkPrerequisitesConfig + testAccDataSourceCcFabricL2VirtualNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricL2VirtualNetworkPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricL2VirtualNetworkConfig() string {
	config := `resource "catalystcenter_fabric_l2_virtual_network" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	vlan_name = "VLAN401"` + "\n"
	config += `	vlan_id = 401` + "\n"
	config += `	traffic_type = "DATA"` + "\n"
	config += `	fabric_enabled_wireless = false` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_l2_virtual_network" "test" {
			fabric_id = catalystcenter_fabric_site.test.id
			vlan_name = "VLAN401"
			depends_on = [catalystcenter_fabric_l2_virtual_network.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
