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
func TestAccCcFabricVLANToSSID(t *testing.T) {
	if os.Getenv("WIRELESS") == "" {
		t.Skip("skipping test, set environment variable WIRELESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_vlan_to_ssid.test", "fabric_id", "5e8e3e3e-1b6b-4b6b-8b6b-1b6b4b6b8b6b"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_vlan_to_ssid.test", "mappings.0.vlan_name", "VLAN_1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_vlan_to_ssid.test", "mappings.0.ssid_details.0.name", "mySSID1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_vlan_to_ssid.test", "mappings.0.ssid_details.0.security_group_tag", "Auditors"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcFabricVLANToSSIDConfig_all(),
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
func testAccCcFabricVLANToSSIDConfig_minimum() string {
	config := `resource "catalystcenter_fabric_vlan_to_ssid" "test" {` + "\n"
	config += `	fabric_id = "5e8e3e3e-1b6b-4b6b-8b6b-1b6b4b6b8b6b"` + "\n"
	config += `	mappings = [{` + "\n"
	config += `	  vlan_name = "VLAN_1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFabricVLANToSSIDConfig_all() string {
	config := `resource "catalystcenter_fabric_vlan_to_ssid" "test" {` + "\n"
	config += `	fabric_id = "5e8e3e3e-1b6b-4b6b-8b6b-1b6b4b6b8b6b"` + "\n"
	config += `	mappings = [{` + "\n"
	config += `	  vlan_name = "VLAN_1"` + "\n"
	config += `	  ssid_details = [{` + "\n"
	config += `		name = "mySSID1"` + "\n"
	config += `		security_group_tag = "Auditors"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
