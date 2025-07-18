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
func TestAccDataSourceCcVirtualNetworkToFabricSite(t *testing.T) {
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcVirtualNetworkToFabricSitePrerequisitesConfig + testAccDataSourceCcVirtualNetworkToFabricSiteConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcVirtualNetworkToFabricSitePrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  authentication_profile_name = "No Authentication"
  pub_sub_enabled             = false

  depends_on = [catalystcenter_area.test]
}
resource "catalystcenter_fabric_virtual_network" "test" {
  virtual_network_name = "SDA_VN1"
  is_guest             = false
  sg_names             = ["Employees"]

  depends_on = [catalystcenter_fabric_site.test]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcVirtualNetworkToFabricSiteConfig() string {
	config := `resource "catalystcenter_virtual_network_to_fabric_site" "test" {` + "\n"
	config += `	virtual_network_name = catalystcenter_fabric_virtual_network.test.id` + "\n"
	config += `	site_name_hierarchy = "${catalystcenter_area.test.parent_name}/${catalystcenter_area.test.name}"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_virtual_network_to_fabric_site" "test" {
			virtual_network_name = catalystcenter_fabric_virtual_network.test.id
			site_name_hierarchy = "${catalystcenter_area.test.parent_name}/${catalystcenter_area.test.name}"
			depends_on = [catalystcenter_virtual_network_to_fabric_site.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
