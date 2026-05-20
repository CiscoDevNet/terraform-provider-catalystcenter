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
func TestAccDataSourceCcPlannedAccessPointPosition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "name", "AP-1001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "ap_type", "AP9166I"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "position_x", "50.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "position_y", "30.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "position_z", "3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "radios.0.channel", "36"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "radios.0.tx_power", "17"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "radios.0.antenna_name", "Internal-CW9166I-x-5GHz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "radios.0.antenna_azimuth", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_planned_access_point_position.test", "radios.0.antenna_elevation", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcPlannedAccessPointPositionPrerequisitesConfig + testAccDataSourceCcPlannedAccessPointPositionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcPlannedAccessPointPositionPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "PlannedAPTestArea"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_building" "test" {
  name      = "PlannedAPTestBuilding"
  parent_id = catalystcenter_area.test.id
  country   = "United States"
  address   = "150 W Tasman Dr, San Jose"
  latitude  = 37.338
  longitude = -121.832
}
resource "catalystcenter_floor" "test" {
  parent_id        = catalystcenter_building.test.id
  name             = "PlannedAPTestFloor"
  floor_number     = 1
  rf_model         = "Cubes And Walled Offices"
  width            = 100
  length           = 100
  height           = 10
  units_of_measure = "feet"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcPlannedAccessPointPositionConfig() string {
	config := `resource "catalystcenter_planned_access_point_position" "test" {` + "\n"
	config += `	floor_id = catalystcenter_floor.test.id` + "\n"
	config += `	name = "AP-1001"` + "\n"
	config += `	ap_type = "AP9166I"` + "\n"
	config += `	position_x = 50.5` + "\n"
	config += `	position_y = 30.2` + "\n"
	config += `	position_z = 3.5` + "\n"
	config += `	radios = [{` + "\n"
	config += `	  bands = ["5"]` + "\n"
	config += `	  channel = 36` + "\n"
	config += `	  tx_power = 17` + "\n"
	config += `	  antenna_name = "Internal-CW9166I-x-5GHz"` + "\n"
	config += `	  antenna_azimuth = 0` + "\n"
	config += `	  antenna_elevation = 0` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_planned_access_point_position" "test" {
			floor_id = catalystcenter_floor.test.id
			name = "AP-1001"
			depends_on = [catalystcenter_planned_access_point_position.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
