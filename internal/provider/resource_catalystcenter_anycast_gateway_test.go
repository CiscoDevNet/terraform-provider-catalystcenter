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
func TestAccCcAnycastGateway(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "tcp_mss_adjustment", "1400"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "vlan_name", "VLAN401"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "traffic_type", "DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "critical_pool", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "l2_flooding_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "wireless_pool", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "ip_directed_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "intra_subnet_routing_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anycast_gateway.test", "multiple_ip_to_mac_addresses", "false"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcAnycastGatewayPrerequisitesConfig + testAccCcAnycastGatewayConfig_all(),
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
const testAccCcAnycastGatewayPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_ip_pool" "test" {
  name             = "MyPool1"
  ip_subnet        = "172.32.0.0/16"
}
resource "catalystcenter_ip_pool_reservation" "test" {
  site_id            = catalystcenter_area.test.id
  name               = "MyRes1"
  type               = "Generic"
  ipv4_global_pool   = catalystcenter_ip_pool.test.ip_subnet
  ipv4_prefix        = true
  ipv4_prefix_length = 24
  ipv4_subnet        = "172.32.1.0"
  depends_on = [catalystcenter_ip_pool.test]
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
}
resource "catalystcenter_fabric_virtual_network" "test" {
  virtual_network_name = "SDA_VN1"
  is_guest             = false
  sg_names             = ["Employees"]
}
resource "catalystcenter_virtual_network_to_fabric_site" "test" {
  virtual_network_name = catalystcenter_fabric_virtual_network.test.id
  site_name_hierarchy  = "Global/Area1"
  depends_on = [catalystcenter_fabric_site.test]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAnycastGatewayConfig_minimum() string {
	config := `resource "catalystcenter_anycast_gateway" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	virtual_network_name = catalystcenter_virtual_network_to_fabric_site.test.virtual_network_name` + "\n"
	config += `	ip_pool_name = catalystcenter_ip_pool_reservation.test.name` + "\n"
	config += `	traffic_type = "DATA"` + "\n"
	config += `	auto_generate_vlan_name = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAnycastGatewayConfig_all() string {
	config := `resource "catalystcenter_anycast_gateway" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	virtual_network_name = catalystcenter_virtual_network_to_fabric_site.test.virtual_network_name` + "\n"
	config += `	ip_pool_name = catalystcenter_ip_pool_reservation.test.name` + "\n"
	config += `	tcp_mss_adjustment = 1400` + "\n"
	config += `	vlan_name = "VLAN401"` + "\n"
	config += `	traffic_type = "DATA"` + "\n"
	config += `	critical_pool = false` + "\n"
	config += `	l2_flooding_enabled = false` + "\n"
	config += `	wireless_pool = false` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `	intra_subnet_routing_enabled = false` + "\n"
	config += `	multiple_ip_to_mac_addresses = false` + "\n"
	config += `	auto_generate_vlan_name = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
