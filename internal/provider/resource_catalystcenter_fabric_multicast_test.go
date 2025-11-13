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
func TestAccCcFabricMulticast(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_multicast.test", "ipv4_ssm_ranges.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_multicast.test", "multicast_r_ps.0.rp_device_location", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_multicast.test", "multicast_r_ps.0.network_device_ids.0", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcFabricMulticastConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcFabricMulticastConfig_all(),
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
func testAccCcFabricMulticastConfig_minimum() string {
	config := `resource "catalystcenter_fabric_multicast" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	virtual_network_name = catalystcenter_virtual_network_to_fabric_site.test.virtual_network_name` + "\n"
	config += `	ip_pool_name = catalystcenter_ip_pool_reservation.test.name` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFabricMulticastConfig_all() string {
	config := `resource "catalystcenter_fabric_multicast" "test" {` + "\n"
	config += `	fabric_id = catalystcenter_fabric_site.test.id` + "\n"
	config += `	virtual_network_name = catalystcenter_virtual_network_to_fabric_site.test.virtual_network_name` + "\n"
	config += `	ip_pool_name = catalystcenter_ip_pool_reservation.test.name` + "\n"
	config += `	ipv4_ssm_ranges = [""]` + "\n"
	config += `	multicast_r_ps = [{` + "\n"
	config += `	  rp_device_location = ""` + "\n"
	config += `	  network_device_ids = [""]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
