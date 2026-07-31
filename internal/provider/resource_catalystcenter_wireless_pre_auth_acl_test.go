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
func TestAccCcWirelessPreAuthACL(t *testing.T) {
	if os.Getenv("CC32") == "" {
		t.Skip("skipping test, set environment variable CC32")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "acl_list_name", "PreAuth_ACL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "acl_name", "PreAuth_ACL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ipv6_acl_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "include_auto_generated_rules", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.source_address", "250.162.252.170"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.source_subnet_mask_or_prefix", "32"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.destination_address", "250.162.252.171"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.destination_subnet_mask_or_prefix", "32"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.source_ports", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.destination_ports", "100-200"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_pre_auth_acl.test", "ip_acl_rules.0.protocol", "IP"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessPreAuthACLConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_wireless_pre_auth_acl.test",
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
func testAccCcWirelessPreAuthACLConfig_minimum() string {
	config := `resource "catalystcenter_wireless_pre_auth_acl" "test" {` + "\n"
	config += `	acl_list_name = "PreAuth_ACL"` + "\n"
	config += `	acl_name = "PreAuth_ACL"` + "\n"
	config += `	ip_acl_rules = [{` + "\n"
	config += `	  source_address = "250.162.252.170"` + "\n"
	config += `	  source_subnet_mask_or_prefix = 32` + "\n"
	config += `	  destination_address = "250.162.252.171"` + "\n"
	config += `	  destination_subnet_mask_or_prefix = 32` + "\n"
	config += `	  protocol = "IP"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcWirelessPreAuthACLConfig_all() string {
	config := `resource "catalystcenter_wireless_pre_auth_acl" "test" {` + "\n"
	config += `	acl_list_name = "PreAuth_ACL"` + "\n"
	config += `	acl_name = "PreAuth_ACL"` + "\n"
	config += `	ipv6_acl_enabled = false` + "\n"
	config += `	include_auto_generated_rules = true` + "\n"
	config += `	ip_acl_rules = [{` + "\n"
	config += `	  source_address = "250.162.252.170"` + "\n"
	config += `	  source_subnet_mask_or_prefix = 32` + "\n"
	config += `	  destination_address = "250.162.252.171"` + "\n"
	config += `	  destination_subnet_mask_or_prefix = 32` + "\n"
	config += `	  source_ports = "100"` + "\n"
	config += `	  destination_ports = "100-200"` + "\n"
	config += `	  protocol = "IP"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
