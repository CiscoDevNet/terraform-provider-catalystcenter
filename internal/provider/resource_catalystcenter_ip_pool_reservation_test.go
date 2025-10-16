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
func TestAccCcIPPoolReservation(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "name", "MyRes1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "pool_type", "Generic"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv4_subnet", "192.168.10.0"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv4_prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv4_gateway", "192.168.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv6_subnet", "2001:db8:8000::"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv6_prefix_length", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv6_gateway", "2001:db8:8000::1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool_reservation.test", "ipv6_slaac_support", "true"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcIPPoolReservationPrerequisitesConfig + testAccCcIPPoolReservationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_ip_pool_reservation.test",
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
const testAccCcIPPoolReservationPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "Area1"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_ip_pool" "test" {
  name                        = "MyPool1"
  pool_type                   = "Generic"
  address_space_subnet        = "192.168.0.0"
  address_space_prefix_length = 16
  address_space_gateway       = "192.168.0.1"
  address_space_dhcp_servers  = ["192.168.0.10"]
  address_space_dns_servers   = ["192.168.0.53"]
}
resource "catalystcenter_ip_pool" "test_v6" {
  name                        = "MyPool1V6"
  pool_type                   = "Generic"
  address_space_subnet        = "2001:db8:8000::"
  address_space_prefix_length = 36
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcIPPoolReservationConfig_minimum() string {
	config := `resource "catalystcenter_ip_pool_reservation" "test" {` + "\n"
	config += `	name = "MyRes1"` + "\n"
	config += `	pool_type = "Generic"` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	ipv4_subnet = "192.168.10.0"` + "\n"
	config += `	ipv4_prefix_length = 24` + "\n"
	config += `	ipv4_global_pool_id = catalystcenter_ip_pool.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcIPPoolReservationConfig_all() string {
	config := `resource "catalystcenter_ip_pool_reservation" "test" {` + "\n"
	config += `	name = "MyRes1"` + "\n"
	config += `	pool_type = "Generic"` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	ipv4_subnet = "192.168.10.0"` + "\n"
	config += `	ipv4_prefix_length = 24` + "\n"
	config += `	ipv4_gateway = "192.168.10.1"` + "\n"
	config += `	ipv4_dhcp_servers = ["1.2.3.4"]` + "\n"
	config += `	ipv4_dns_servers = ["2.3.4.5"]` + "\n"
	config += `	ipv4_global_pool_id = catalystcenter_ip_pool.test.id` + "\n"
	config += `	ipv6_subnet = "2001:db8:8000::"` + "\n"
	config += `	ipv6_prefix_length = 64` + "\n"
	config += `	ipv6_gateway = "2001:db8:8000::1"` + "\n"
	config += `	ipv6_dhcp_servers = ["2001:db8::1234"]` + "\n"
	config += `	ipv6_dns_servers = ["2001:db8::1234"]` + "\n"
	config += `	ipv6_slaac_support = true` + "\n"
	config += `	ipv6_global_pool_id = catalystcenter_ip_pool.test_v6.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
