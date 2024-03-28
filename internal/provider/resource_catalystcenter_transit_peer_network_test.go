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
func TestAccCcTransitPeerNetwork(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_transit_peer_network.test", "transit_peer_network_name", "TRANSIT_1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_transit_peer_network.test", "transit_peer_network_type", "ip_transit"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_transit_peer_network.test", "routing_protocol_name", "BGP"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_transit_peer_network.test", "autonomous_system_number", "65010"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcTransitPeerNetworkConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_transit_peer_network.test",
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcTransitPeerNetworkConfig_minimum() string {
	config := `resource "catalystcenter_transit_peer_network" "test" {` + "\n"
	config += `	transit_peer_network_name = "TRANSIT_1"` + "\n"
	config += `	transit_peer_network_type = "ip_transit"` + "\n"
	config += `	routing_protocol_name = "BGP"` + "\n"
	config += `	autonomous_system_number = "65010"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcTransitPeerNetworkConfig_all() string {
	config := `resource "catalystcenter_transit_peer_network" "test" {` + "\n"
	config += `	transit_peer_network_name = "TRANSIT_1"` + "\n"
	config += `	transit_peer_network_type = "ip_transit"` + "\n"
	config += `	routing_protocol_name = "BGP"` + "\n"
	config += `	autonomous_system_number = "65010"` + "\n"
	config += `	transit_control_plane_settings = [{` + "\n"
	config += `	  site_name_hierarchy = "Global/Area1"` + "\n"
	config += `	  device_management_ip_address = "10.0.0.1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
