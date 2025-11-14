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
func TestAccCcFabricEWLC(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_ewlc.test", "fabric_id", "8e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_ewlc.test", "network_device_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_ewlc.test", "enable_wireless", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_ewlc.test", "enable_rolling_ap_upgrade", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_fabric_ewlc.test", "ap_reboot_percentage", "25"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcFabricEWLCConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcFabricEWLCConfig_all(),
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
func testAccCcFabricEWLCConfig_minimum() string {
	config := `resource "catalystcenter_fabric_ewlc" "test" {` + "\n"
	config += `	fabric_id = "8e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	network_device_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	enable_wireless = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcFabricEWLCConfig_all() string {
	config := `resource "catalystcenter_fabric_ewlc" "test" {` + "\n"
	config += `	fabric_id = "8e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"` + "\n"
	config += `	network_device_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	enable_wireless = true` + "\n"
	config += `	enable_rolling_ap_upgrade = false` + "\n"
	config += `	ap_reboot_percentage = 25` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
