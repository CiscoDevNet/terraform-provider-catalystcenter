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
func TestAccCcIPPool(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool.test", "name", "MyPool1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool.test", "pool_type", "Generic"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool.test", "address_space_subnet", "192.168.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool.test", "address_space_prefix_length", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_ip_pool.test", "address_space_gateway", "192.168.0.1"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcIPPoolConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_ip_pool.test",
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
func testAccCcIPPoolConfig_minimum() string {
	config := `resource "catalystcenter_ip_pool" "test" {` + "\n"
	config += `	name = "MyPool1"` + "\n"
	config += `	pool_type = "Generic"` + "\n"
	config += `	address_space_subnet = "192.168.0.0"` + "\n"
	config += `	address_space_prefix_length = 16` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcIPPoolConfig_all() string {
	config := `resource "catalystcenter_ip_pool" "test" {` + "\n"
	config += `	name = "MyPool1"` + "\n"
	config += `	pool_type = "Generic"` + "\n"
	config += `	address_space_subnet = "192.168.0.0"` + "\n"
	config += `	address_space_prefix_length = 16` + "\n"
	config += `	address_space_gateway = "192.168.0.1"` + "\n"
	config += `	address_space_dhcp_servers = ["192.168.0.10"]` + "\n"
	config += `	address_space_dns_servers = ["192.168.0.53"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
