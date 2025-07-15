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
func TestAccDataSourceCcIPPools(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pools.test", "pools.0.name", "MyPool1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pools.test", "pools.0.pool_type", "Generic"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pools.test", "pools.0.gateway_ip_address", "21.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pools.test", "pools.0.subnet", "21.1.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ip_pools.test", "pools.0.prefix_length", "24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcIPPoolsPrerequisitesConfig + testAccDataSourceCcIPPoolsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcIPPoolsPrerequisitesConfig = `
resource "catalystcenter_ip_pool" "test" {
  name             = "MyPool1"
  ip_address_space = "IPv4"
  type             = "generic"
  ip_subnet        = "21.1.1.0/24"
  gateway          = ["21.1.1.1"]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcIPPoolsConfig() string {

	config := `
		data "catalystcenter_ip_pools" "test" {
			depends_on = [catalystcenter_ip_pool.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
