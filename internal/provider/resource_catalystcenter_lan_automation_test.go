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
func TestAccCcLANAutomation(t *testing.T) {
	if os.Getenv("LANAUTO") == "" {
		t.Skip("skipping test, set environment variable LANAUTO")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "discovered_device_site_name_hierarchy", "Global/Area1/Area2"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "primary_device_management_ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "peer_device_management_ip_address", "1.2.3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "ip_pools.0.ip_pool_name", "POOL1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "ip_pools.0.ip_pool_role", "MAIN_POOL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "multicast_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "host_name_prefix", "TEST"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "isis_domain_password", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_lan_automation.test", "redistribute_isis_to_bgp", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcLANAutomationConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcLANAutomationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_lan_automation.test",
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
func testAccCcLANAutomationConfig_minimum() string {
	config := `resource "catalystcenter_lan_automation" "test" {` + "\n"
	config += `	discovered_device_site_name_hierarchy = "Global/Area1/Area2"` + "\n"
	config += `	primary_device_management_ip_address = "1.2.3.4"` + "\n"
	config += `	primary_device_interface_names = ["HundredGigE1/0/1"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcLANAutomationConfig_all() string {
	config := `resource "catalystcenter_lan_automation" "test" {` + "\n"
	config += `	discovered_device_site_name_hierarchy = "Global/Area1/Area2"` + "\n"
	config += `	primary_device_management_ip_address = "1.2.3.4"` + "\n"
	config += `	peer_device_management_ip_address = "1.2.3.5"` + "\n"
	config += `	primary_device_interface_names = ["HundredGigE1/0/1"]` + "\n"
	config += `	ip_pools = [{` + "\n"
	config += `	  ip_pool_name = "POOL1"` + "\n"
	config += `	  ip_pool_role = "MAIN_POOL"` + "\n"
	config += `	}]` + "\n"
	config += `	multicast_enabled = true` + "\n"
	config += `	host_name_prefix = "TEST"` + "\n"
	config += `	isis_domain_password = "cisco123"` + "\n"
	config += `	redistribute_isis_to_bgp = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
