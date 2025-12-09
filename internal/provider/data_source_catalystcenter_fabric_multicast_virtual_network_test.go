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
func TestAccDataSourceCcFabricMulticastVirtualNetwork(t *testing.T) {
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricMulticastVirtualNetworkPrerequisitesConfig + testAccDataSourceCcFabricMulticastVirtualNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcFabricMulticastVirtualNetworkPrerequisitesConfig = `
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
resource "catalystcenter_ip_pool_reservation" "test" {
  name                = "MyRes1"
  pool_type           = "Generic"
  site_id             = catalystcenter_area.test.id
  ipv4_subnet         = "192.168.10.0"
  ipv4_prefix_length  = 24
  ipv4_gateway        = "192.168.10.1"
  ipv4_dhcp_servers   = ["1.2.3.4"]
  ipv4_dns_servers    = ["2.3.4.5"]
  ipv4_global_pool_id = catalystcenter_ip_pool.test.id
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
  depends_on = [catalystcenter_ip_pool_reservation.test]
}
resource "catalystcenter_fabric_l3_virtual_network" "test" {
  virtual_network_name = "SDA_VN1"
  fabric_ids           = [catalystcenter_fabric_site.test.id]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricMulticastVirtualNetworkConfig() string {
	config := `resource "catalystcenter_fabric_multicast_virtual_network" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	virtual_network_name = catalystcenter_fabric_l3_virtual_network.test.virtual_network_name` + "\n"
	config += `	ip_pool_name = catalystcenter_ip_pool_reservation.test.name` + "\n"
	config += `	ipv4_ssm_ranges = ["225.0.0.0/8"]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_multicast_virtual_network" "test" {
			fabric_id = catalystcenter_fabric_site.test.id
			virtual_network_name = catalystcenter_fabric_l3_virtual_network.test.virtual_network_name
			depends_on = [catalystcenter_fabric_multicast_virtual_network.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
