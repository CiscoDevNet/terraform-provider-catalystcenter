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
func TestAccDataSourceCcFloors(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.parent_name_hierarchy", "Global/USA/Building1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.name", "Floor1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.floor_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.rf_model", "Drywall Office Only"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.width", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.length", "50.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.height", "3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_floors.test", "floors.0.units_of_measure", "meters"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFloorsPrerequisitesConfig + testAccDataSourceCcFloorsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFloorsPrerequisitesConfig = `
resource "catalystcenter_areas" "test" {
  areas = [
    {
      parent_name_hierarchy = "Global"
      name                  = "USA"
    }
  ]
}
resource "catalystcenter_buildings" "test" {
  buildings = [
    {
      parent_name_hierarchy = "Global/USA"
      name                  = "Building1"
      country               = "United States"
      address               = "150 W Tasman Dr, San Jose"
      latitude              = 37.338
      longitude             = -121.832
    }
  ]
  depends_on = [catalystcenter_areas.test]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFloorsConfig() string {
	config := `resource "catalystcenter_floors" "test" {` + "\n"
	config += `	floors = [{` + "\n"
	config += `	  parent_name_hierarchy = "Global/USA/Building1"` + "\n"
	config += `	  name = "Floor1"` + "\n"
	config += `	  floor_number = 1` + "\n"
	config += `	  rf_model = "Drywall Office Only"` + "\n"
	config += `	  width = 40` + "\n"
	config += `	  length = 50.5` + "\n"
	config += `	  height = 3.5` + "\n"
	config += `	  units_of_measure = "meters"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_floors" "test" {
			id = catalystcenter_floors.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
