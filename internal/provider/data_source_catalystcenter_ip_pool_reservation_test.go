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
func TestAccDataSourceCcIPPoolReservation(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pool_reservation.test", "name", "MyRes1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pool_reservation.test", "ipv4_prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pool_reservation.test", "ipv4_subnet", "172.32.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pool_reservation.test", "ipv4_gateway", "172.32.1.1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcIPPoolReservationPrerequisitesConfig + testAccDataSourceCcIPPoolReservationConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcIPPoolReservationPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_id   = data.catalystcenter_site.test.id
  depends_on  = [catalystcenter_ip_pool.test]
}
resource "catalystcenter_ip_pool" "test" {
  name             = "MyPool1"
  ip_subnet        = "172.32.0.0/16"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcIPPoolReservationConfig() string {
	config := `resource "catalystcenter_ip_pool_reservation" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	name = "MyRes1"` + "\n"
	config += `	type = "Generic"` + "\n"
	config += `	ipv6_address_space = false` + "\n"
	config += `	ipv4_global_pool = "172.32.0.0/16"` + "\n"
	config += `	ipv4_prefix = true` + "\n"
	config += `	ipv4_prefix_length = 24` + "\n"
	config += `	ipv4_subnet = "172.32.1.0"` + "\n"
	config += `	ipv4_gateway = "172.32.1.1"` + "\n"
	config += `	ipv4_dhcp_servers = ["1.2.3.4"]` + "\n"
	config += `	ipv4_dns_servers = ["2.3.4.5"]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_ip_pool_reservation" "test" {
			site_id = catalystcenter_area.test.id
			name = "MyRes1"
			depends_on = [catalystcenter_ip_pool_reservation.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
