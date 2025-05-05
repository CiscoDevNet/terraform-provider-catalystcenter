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
func TestAccDataSourceCcFabricL3HandoffSDATransit(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_sda_transit.test", "network_device_id", "a144a086-750c-4af1-ac57-feab33a69851"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_sda_transit.test", "affinity_id_prime", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_sda_transit.test", "affinity_id_decider", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_sda_transit.test", "connected_to_internet", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_l3_handoff_sda_transit.test", "is_multicast_over_transit_enabled", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricL3HandoffSDATransitPrerequisitesConfig + testAccDataSourceCcFabricL3HandoffSDATransitConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricL3HandoffSDATransitPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = true
  authentication_profile_name = "No Authentication"
  depends_on = [catalystcenter_area.test]
}
resource "catalystcenter_transit_network" "test" {
  name = "SDA_TRANSIT_1"
  type = "SDA_LISP_BGP_TRANSIT"
  control_plane_network_device_ids = ["12345678-1234-1234-1234-123456789012"]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricL3HandoffSDATransitConfig() string {
	config := `resource "catalystcenter_fabric_l3_handoff_sda_transit" "test" {` + "\n"
	config += `	network_device_id = "a144a086-750c-4af1-ac57-feab33a69851"` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	transit_network_id = catalystcenter_transit_network.test.id` + "\n"
	config += `	affinity_id_prime = 100` + "\n"
	config += `	affinity_id_decider = 100` + "\n"
	config += `	connected_to_internet = false` + "\n"
	config += `	is_multicast_over_transit_enabled = false` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_l3_handoff_sda_transit" "test" {
			network_device_id = "a144a086-750c-4af1-ac57-feab33a69851"
			fabric_id = catalystcenter_fabric_site.test.id
			depends_on = [catalystcenter_fabric_l3_handoff_sda_transit.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
