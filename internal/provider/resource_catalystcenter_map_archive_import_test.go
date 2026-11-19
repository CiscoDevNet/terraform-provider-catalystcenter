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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcMapArchiveImport(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_map_archive_import.test", "accept_data_loss", "true"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcMapArchiveImportPrerequisitesConfig + testAccCcMapArchiveImportConfig_all(),
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
const testAccCcMapArchiveImportPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "NAC_MapImport"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_building" "test" {
  name      = "NAC_MapImport_Bldg"
  parent_id = catalystcenter_area.test.id
  country   = "United States"
  address   = "150 W Tasman Dr, San Jose"
  latitude  = 37.338
  longitude = -121.832
}
resource "catalystcenter_floor" "test" {
  name             = "NAC_MapImport_F1"
  parent_id        = catalystcenter_building.test.id
  floor_number     = 1
  rf_model         = "Drywall Office Only"
  width            = 30
  length           = 20
  height           = 3
  units_of_measure = "meters"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcMapArchiveImportConfig_minimum() string {
	config := `resource "catalystcenter_map_archive_import" "test" {` + "\n"
	config += `	source_path = "testdata/map_archives/nac_map_import.tar.gz"` + "\n"
	config += `	accept_data_loss = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcMapArchiveImportConfig_all() string {
	config := `resource "catalystcenter_map_archive_import" "test" {` + "\n"
	config += `	source_path = "testdata/map_archives/nac_map_import.tar.gz"` + "\n"
	config += `	accept_data_loss = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
