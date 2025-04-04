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
func TestAccCcFabricPortAssignments(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_port_assignments.test", "port_assignments.0.network_device_id", "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_port_assignments.test", "port_assignments.0.interface_name", "GigabitEthernet1/0/2"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_port_assignments.test", "port_assignments.0.connected_device_type", "USER_DEVICE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_port_assignments.test", "port_assignments.0.data_vlan_name", "DATA_VLAN"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_port_assignments.test", "port_assignments.0.voice_vlan_name", "VOICE_VLAN"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_port_assignments.test", "port_assignments.0.authenticate_template_name", "No Authentication"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcFabricPortAssignmentsPrerequisitesConfig + testAccCcFabricPortAssignmentsConfig_all(),
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
const testAccCcFabricPortAssignmentsPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcFabricPortAssignmentsConfig_minimum() string {
	config := `resource "catalystcenter_fabric_port_assignments" "test" {` + "\n"
	config += `	fabric_id = "e02d9911-b0a7-435b-bb46-079d877d7b3e"` + "\n"
	config += `	network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b"` + "\n"
	config += `	port_assignments = [{` + "\n"
	config += `	  fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	  network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	  interface_name = "GigabitEthernet1/0/2"` + "\n"
	config += `	  connected_device_type = "USER_DEVICE"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFabricPortAssignmentsConfig_all() string {
	config := `resource "catalystcenter_fabric_port_assignments" "test" {` + "\n"
	config += `	fabric_id = "e02d9911-b0a7-435b-bb46-079d877d7b3e"` + "\n"
	config += `	network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b"` + "\n"
	config += `	port_assignments = [{` + "\n"
	config += `	  fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	  network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	  interface_name = "GigabitEthernet1/0/2"` + "\n"
	config += `	  connected_device_type = "USER_DEVICE"` + "\n"
	config += `	  data_vlan_name = "DATA_VLAN"` + "\n"
	config += `	  voice_vlan_name = "VOICE_VLAN"` + "\n"
	config += `	  authenticate_template_name = "No Authentication"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
