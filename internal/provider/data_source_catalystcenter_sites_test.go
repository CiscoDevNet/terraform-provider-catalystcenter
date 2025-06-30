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
func TestAccDataSourceCcSites(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "type", "building"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.name", "Building1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.name_hierarchy", "Global/Building1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.type", "building"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.latitude", "37.338"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.longitude", "-121.832"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.address", "150 W Tasman Dr, San Jose"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_sites.test", "sites.0.country", "United States"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcSitesPrerequisitesConfig + testAccDataSourceCcSitesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcSitesPrerequisitesConfig = `
resource "catalystcenter_building" "test" {
  name        = "Building1"
  parent_name = "Global"
  country     = "United States"
  address     = "150 W Tasman Dr, San Jose"
  latitude    = 37.338
  longitude   = -121.832
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcSitesConfig() string {

	config := `
		data "catalystcenter_sites" "test" {
			type = "building"
 			depends_on = [catalystcenter_building.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
