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
func TestAccCcVirtualNetworkIPPool(t *testing.T) {
	if os.Getenv("VIRTUAL_NETWORK_IP_POOL") == "" {
		t.Skip("skipping test, set environment variable VIRTUAL_NETWORK_IP_POOL")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "layer2_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "vlan_id", "401"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "vlan_name", "VLAN401"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "traffic_type", "DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "l2_flooding_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "wireless_pool", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "ip_directed_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "common_pool", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_virtual_network_ip_pool.test", "bridge_mode_vm", "false"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcVirtualNetworkIPPoolPrerequisitesConfig + testAccCcVirtualNetworkIPPoolConfig_all(),
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
const testAccCcVirtualNetworkIPPoolPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
  depends_on  = [catalystcenter_ip_pool.test]
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
  authentication_profile_name = "No Authentication"
  pub_sub_enabled             = false
}
resource "catalystcenter_fabric_virtual_network" "test" {
  virtual_network_name = "SDA_VN1"
  is_guest             = false
  sg_names             = ["Employees"]
}
resource "catalystcenter_virtual_network_to_fabric_site" "test" {
  virtual_network_name = "SDA_VN1"
  site_name_hierarchy  = "Global/Area1"
  depends_on = [catalystcenter_fabric_virtual_network.test, catalystcenter_fabric_site.test]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcVirtualNetworkIPPoolConfig_minimum() string {
	config := `resource "catalystcenter_virtual_network_ip_pool" "test" {` + "\n"
	config += `	virtual_network_name = catalystcenter_virtual_network_to_fabric_site.test.virtual_network_name` + "\n"
	config += `	site_name_hierarchy = catalystcenter_virtual_network_to_fabric_site.test.site_name_hierarchy` + "\n"
	config += `	vlan_name = "VLAN401"` + "\n"
	config += `	traffic_type = "DATA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcVirtualNetworkIPPoolConfig_all() string {
	config := `resource "catalystcenter_virtual_network_ip_pool" "test" {` + "\n"
	config += `	virtual_network_name = catalystcenter_virtual_network_to_fabric_site.test.virtual_network_name` + "\n"
	config += `	site_name_hierarchy = catalystcenter_virtual_network_to_fabric_site.test.site_name_hierarchy` + "\n"
	config += `	layer2_only = false` + "\n"
	config += `	ip_pool_name = catalystcenter_ip_pool_reservation.test.name` + "\n"
	config += `	vlan_id = "401"` + "\n"
	config += `	vlan_name = "VLAN401"` + "\n"
	config += `	traffic_type = "DATA"` + "\n"
	config += `	l2_flooding_enabled = false` + "\n"
	config += `	critical_pool = false` + "\n"
	config += `	wireless_pool = false` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `	common_pool = false` + "\n"
	config += `	bridge_mode_vm = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
