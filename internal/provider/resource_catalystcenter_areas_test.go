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

func TestAccCcAreas(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_areas.test", "areas.0.parent_name_hierarchy", "Global"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_areas.test", "areas.0.name", "Area1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcAreasPrerequisitesConfig + testAccCcAreasConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcAreasPrerequisitesConfig + testAccCcAreasConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_areas.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccCcAreasPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}

`

func testAccCcAreasConfig_minimum() string {
	config := `resource "catalystcenter_areas" "test" {` + "\n"
	config += `	areas = [{` + "\n"
	config += `	  parent_name_hierarchy = "Global"` + "\n"
	config += `	  name = "Area1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccCcAreasConfig_all() string {
	config := `resource "catalystcenter_areas" "test" {` + "\n"
	config += `	areas = [{` + "\n"
	config += `	  parent_name_hierarchy = "Global"` + "\n"
	config += `	  name = "Area1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
