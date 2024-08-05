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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcFabricDevice(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "network_device_id", "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "device_roles.0", "CONTROL_PLANE_NODE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "border_types.0", "LAYER_3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "local_autonomous_system_number", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "default_exit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "import_external_routes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "border_priority", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_device.test", "prepend_autonomous_system_count", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricDevicePrerequisitesConfig + testAccDataSourceCcFabricDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricDevicePrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
  depends_on = [catalystcenter_area.test]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricDeviceConfig() string {
	config := `resource "catalystcenter_fabric_device" "test" {` + "\n"
	config += `	network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	device_roles = ["CONTROL_PLANE_NODE"]` + "\n"
	config += `	border_types = ["LAYER_3"]` + "\n"
	config += `	local_autonomous_system_number = "65000"` + "\n"
	config += `	default_exit = true` + "\n"
	config += `	import_external_routes = false` + "\n"
	config += `	border_priority = 5` + "\n"
	config += `	prepend_autonomous_system_count = 1` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_device" "test" {
			network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
			fabric_id = catalystcenter_fabric_site.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
