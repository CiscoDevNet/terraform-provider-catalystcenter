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
func TestAccDataSourceCcFabricZone(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_zone.test", "authentication_profile_name", "No Authentication"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricZonePrerequisitesConfig + testAccDataSourceCcFabricZoneConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricZonePrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "Area1"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
}
resource "catalystcenter_building" "test" {
  name        = "Building1"
  parent_id   = catalystcenter_area.test.id
  country     = "United States"
  address     = "150 W Tasman Dr, San Jose"
  latitude    = 37.338
  longitude   = -121.832
  
  depends_on = [catalystcenter_fabric_site.test]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricZoneConfig() string {
	config := `resource "catalystcenter_fabric_zone" "test" {` + "\n"
	config += `	site_id = catalystcenter_building.test.id` + "\n"
	config += `	authentication_profile_name = "No Authentication"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_zone" "test" {
			site_id = catalystcenter_building.test.id
			depends_on = [catalystcenter_fabric_zone.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
