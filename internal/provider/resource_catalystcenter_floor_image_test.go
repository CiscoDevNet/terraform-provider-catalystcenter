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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcFloorImage(t *testing.T) {
	if os.Getenv("FLOOR_IMAGE") == "" {
		t.Skip("skipping test, set environment variable FLOOR_IMAGE")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcFloorImagePrerequisitesConfig + testAccCcFloorImageConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccCcFloorImagePrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_building" "test" {
  name      = "FloorImageTestBldg"
  parent_id = data.catalystcenter_site.test.id
  country   = "United States"
  address   = "150 W Tasman Dr, San Jose"
  latitude  = 37.338
  longitude = -121.832
}
resource "catalystcenter_floor" "test" {
  name             = "FloorImageTestFloor"
  parent_id        = catalystcenter_building.test.id
  floor_number     = 1
  rf_model         = "Drywall Office Only"
  width            = 100
  length           = 100
  height           = 3
  units_of_measure = "feet"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcFloorImageConfig_minimum() string {
	config := `resource "catalystcenter_floor_image" "test" {` + "\n"
	config += `	floor_id = catalystcenter_floor.test.id` + "\n"
	config += `	source_path = "testdata/floor_plans/test_floor_plan.png"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFloorImageConfig_all() string {
	config := `resource "catalystcenter_floor_image" "test" {` + "\n"
	config += `	floor_id = catalystcenter_floor.test.id` + "\n"
	config += `	source_path = "testdata/floor_plans/test_floor_plan.png"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
