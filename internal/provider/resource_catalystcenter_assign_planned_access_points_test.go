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
func TestAccCcAssignPlannedAccessPoints(t *testing.T) {
	if os.Getenv("INVENTORY") == "" {
		t.Skip("skipping test, set environment variable INVENTORY")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_assign_planned_access_points.test", "assignments.0.access_point_id", "b054deab-f9fc-48c9-88d0-ab82f029e6aa"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcAssignPlannedAccessPointsPrerequisitesConfig + testAccCcAssignPlannedAccessPointsConfig_all(),
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
const testAccCcAssignPlannedAccessPointsPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "NAC_PlannedAPAssign"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_building" "test" {
  name      = "NAC_PlannedAPAssign_Bldg"
  parent_id = catalystcenter_area.test.id
  country   = "United States"
  address   = "150 W Tasman Dr, San Jose"
  latitude  = 37.338
  longitude = -121.832
}
resource "catalystcenter_floor" "test" {
  name             = "NAC_PlannedAPAssign_F1"
  parent_id        = catalystcenter_building.test.id
  floor_number     = 1
  rf_model         = "Drywall Office Only"
  width            = 30
  length           = 20
  height           = 3
  units_of_measure = "meters"
}
resource "catalystcenter_map_archive_import" "test" {
  source_path      = "testdata/map_archives/nac_planned_ap_assign.tar.gz"
  accept_data_loss = true
  archive_sha256   = filesha256("testdata/map_archives/nac_planned_ap_assign.tar.gz")
  depends_on       = [catalystcenter_floor.test]
}
data "catalystcenter_site" "floor" {
  name_hierarchy = "Global/NAC_PlannedAPAssign/NAC_PlannedAPAssign_Bldg/NAC_PlannedAPAssign_F1"
  depends_on     = [catalystcenter_map_archive_import.test]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAssignPlannedAccessPointsConfig_minimum() string {
	config := `resource "catalystcenter_assign_planned_access_points" "test" {` + "\n"
	config += `	floor_id = data.catalystcenter_site.floor.id` + "\n"
	config += `	assignments = [{` + "\n"
	config += `	  planned_access_point_name = "NAC-PLANNED-AP-01"` + "\n"
	config += `	  access_point_id = "b054deab-f9fc-48c9-88d0-ab82f029e6aa"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAssignPlannedAccessPointsConfig_all() string {
	config := `resource "catalystcenter_assign_planned_access_points" "test" {` + "\n"
	config += `	floor_id = data.catalystcenter_site.floor.id` + "\n"
	config += `	assignments = [{` + "\n"
	config += `	  planned_access_point_name = "NAC-PLANNED-AP-01"` + "\n"
	config += `	  access_point_id = "b054deab-f9fc-48c9-88d0-ab82f029e6aa"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
