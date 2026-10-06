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
func TestAccCcPnPNetworkDeviceClaim(t *testing.T) {
	if os.Getenv("PNP") == "" {
		t.Skip("skipping test, set environment variable PNP")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "device_id", "683413ec67f7d77edb0fd605"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "device_type", "SVL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "remove_inactive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "template_id", "1e3b9f1b-acd9-462a-8a2e-13b527934549"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "template_parameters.0.name", "param5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "template_parameters.0.value", "value5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "domain", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "svl_members.0.serial_number", "FJC28041KKA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "svl_members.0.role", "ACTIVE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "svl_members.0.svl_links.0.local_interface", "TwentyFiveGigE1/0/47"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_network_device_claim.test", "svl_members.0.svl_links.0.remote_interface", "TwentyFiveGigE2/0/47"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcPnPNetworkDeviceClaimConfig_all(),
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcPnPNetworkDeviceClaimConfig_minimum() string {
	config := `resource "catalystcenter_pnp_network_device_claim" "test" {` + "\n"
	config += `	device_id = "683413ec67f7d77edb0fd605"` + "\n"
	config += `	device_type = "SVL"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcPnPNetworkDeviceClaimConfig_all() string {
	config := `resource "catalystcenter_pnp_network_device_claim" "test" {` + "\n"
	config += `	device_id = "683413ec67f7d77edb0fd605"` + "\n"
	config += `	device_type = "SVL"` + "\n"
	config += `	remove_inactive = false` + "\n"
	config += `	template_id = "1e3b9f1b-acd9-462a-8a2e-13b527934549"` + "\n"
	config += `	template_parameters = [{` + "\n"
	config += `	  name = "param5"` + "\n"
	config += `	  value = "value5"` + "\n"
	config += `	}]` + "\n"
	config += `	domain = 2` + "\n"
	config += `	svl_members = [{` + "\n"
	config += `	  serial_number = "FJC28041KKA"` + "\n"
	config += `	  role = "ACTIVE"` + "\n"
	config += `	  svl_links = [{` + "\n"
	config += `		local_interface = "TwentyFiveGigE1/0/47"` + "\n"
	config += `		remote_interface = "TwentyFiveGigE2/0/47"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
