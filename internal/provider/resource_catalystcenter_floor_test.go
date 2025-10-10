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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcFloor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "name", "Floor1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "floor_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "rf_model", "Drywall Office Only"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "width", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "length", "50.5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "height", "3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_floor.test", "units_of_measure", "meters"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcFloorPrerequisitesConfig + testAccCcFloorConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcFloorPrerequisitesConfig + testAccCcFloorConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_floor.test",
		ImportState:  true,
		ImportStateIdFunc: func(s *terraform.State) (string, error) {
			rs, ok := s.RootModule().Resources["catalystcenter_floor.test"]
			if !ok {
				return "", fmt.Errorf("resource not found in state")
			}
			return fmt.Sprintf("%s,%s", rs.Primary.Attributes["units_of_measure"], rs.Primary.ID), nil
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccCcFloorPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcFloorConfig_minimum() string {
	config := `resource "catalystcenter_floor" "test" {` + "\n"
	config += `	parent_id = catalystcenter_building.test.id` + "\n"
	config += `	name = "Floor1"` + "\n"
	config += `	floor_number = 1` + "\n"
	config += `	rf_model = "Drywall Office Only"` + "\n"
	config += `	width = 40` + "\n"
	config += `	length = 50.5` + "\n"
	config += `	height = 3.5` + "\n"
	config += `	units_of_measure = "meters"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFloorConfig_all() string {
	config := `resource "catalystcenter_floor" "test" {` + "\n"
	config += `	parent_id = catalystcenter_building.test.id` + "\n"
	config += `	name = "Floor1"` + "\n"
	config += `	floor_number = 1` + "\n"
	config += `	rf_model = "Drywall Office Only"` + "\n"
	config += `	width = 40` + "\n"
	config += `	length = 50.5` + "\n"
	config += `	height = 3.5` + "\n"
	config += `	units_of_measure = "meters"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
