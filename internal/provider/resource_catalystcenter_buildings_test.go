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

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCcBuildings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.Global/USA/Building1.parent_name_hierarchy", "Global/USA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.Global/USA/Building1.name", "Building1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.Global/USA/Building1.country", "United States"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.Global/USA/Building1.address", "150 W Tasman Dr, San Jose"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.Global/USA/Building1.latitude", "37.338"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_buildings.test", "buildings.Global/USA/Building1.longitude", "-121.832"))

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

const testAccCcBuildingsPrerequisitesConfig = `
resource "catalystcenter_areas" "test" {
  areas = {
    "Global/USA" = {
      parent_name_hierarchy = "Global"
      name                  = "USA"
    }
  }
}

`

func testAccCcBuildingsConfig_minimum() string {
	config := `resource "catalystcenter_buildings" "test" {` + "\n"
	config += `	buildings = {` + "\n"
	config += `	  "Global/USA/Building1" = {` + "\n"
	config += `	    parent_name_hierarchy = "Global/USA"` + "\n"
	config += `	    name = "Building1"` + "\n"
	config += `	    country = "United States"` + "\n"
	config += `	    address = "150 W Tasman Dr, San Jose"` + "\n"
	config += `	  }` + "\n"
	config += `	}` + "\n"
	config += `	depends_on = [catalystcenter_areas.test]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccCcBuildingsConfig_all() string {
	config := `resource "catalystcenter_buildings" "test" {` + "\n"
	config += `	buildings = {` + "\n"
	config += `	  "Global/USA/Building1" = {` + "\n"
	config += `	    parent_name_hierarchy = "Global/USA"` + "\n"
	config += `	    name = "Building1"` + "\n"
	config += `	    country = "United States"` + "\n"
	config += `	    address = "150 W Tasman Dr, San Jose"` + "\n"
	config += `	    latitude = 37.338` + "\n"
	config += `	    longitude = -121.832` + "\n"
	config += `	  }` + "\n"
	config += `	}` + "\n"
	config += `	depends_on = [catalystcenter_areas.test]` + "\n"
	config += `}` + "\n"
	return config
}
