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
func TestAccCcBuildings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.0.parent_name_hierarchy", "Global/USA/San Jose"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.0.name", "Building1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.0.country", "United States"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.0.address", "150 W Tasman Dr, San Jose"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.0.latitude", "37.338"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.0.longitude", "-121.832"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcBuildingsPrerequisitesConfig + testAccCcBuildingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcBuildingsPrerequisitesConfig + testAccCcBuildingsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_buildings.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccCcBuildingsPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name       = "USA"
  parent_id  = data.catalystcenter_site.test.id
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcBuildingsConfig_minimum() string {
	config := `resource "catalystcenter_buildings" "test" {` + "\n"
	config += `	buildings = [{` + "\n"
	config += `	  parent_name_hierarchy = "Global/USA/San Jose"` + "\n"
	config += `	  name = "Building1"` + "\n"
	config += `	  country = "United States"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcBuildingsConfig_all() string {
	config := `resource "catalystcenter_buildings" "test" {` + "\n"
	config += `	buildings = [{` + "\n"
	config += `	  parent_name_hierarchy = "Global/USA/San Jose"` + "\n"
	config += `	  name = "Building1"` + "\n"
	config += `	  country = "United States"` + "\n"
	config += `	  address = "150 W Tasman Dr, San Jose"` + "\n"
	config += `	  latitude = 37.338` + "\n"
	config += `	  longitude = -121.832` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
