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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcTransitPeerNetwork(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_transit_peer_network.test", "transit_peer_network_name", "TRANSIT_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_transit_peer_network.test", "transit_peer_network_type", "ip_transit"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_transit_peer_network.test", "routing_protocol_name", "BGP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_transit_peer_network.test", "autonomous_system_number", "65010"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcTransitPeerNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcTransitPeerNetworkConfig() string {
	config := `resource "catalystcenter_transit_peer_network" "test" {` + "\n"
	config += `	transit_peer_network_name = "TRANSIT_1"` + "\n"
	config += `	transit_peer_network_type = "ip_transit"` + "\n"
	config += `	routing_protocol_name = "BGP"` + "\n"
	config += `	autonomous_system_number = "65010"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_transit_peer_network" "test" {
			transit_peer_network_name = "TRANSIT_1"
			depends_on = [catalystcenter_transit_peer_network.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
