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
func TestAccCcWirelessRRMFRAConfiguration(t *testing.T) {
	if os.Getenv("CC315") == "" {
		t.Skip("skipping test, set environment variable CC315")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rrm_fra_configuration.test", "design_name", "Test"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rrm_fra_configuration.test", "radio_band", "2_4GHZ_5GHZ"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rrm_fra_configuration.test", "fra_freeze", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rrm_fra_configuration.test", "fra_status", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rrm_fra_configuration.test", "fra_interval", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rrm_fra_configuration.test", "fra_sensitivity", "MEDIUM"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcWirelessRRMFRAConfigurationConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessRRMFRAConfigurationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_wireless_rrm_fra_configuration.test",
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
func testAccCcWirelessRRMFRAConfigurationConfig_minimum() string {
	config := `resource "catalystcenter_wireless_rrm_fra_configuration" "test" {` + "\n"
	config += `	design_name = "Test"` + "\n"
	config += `	radio_band = "2_4GHZ_5GHZ"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcWirelessRRMFRAConfigurationConfig_all() string {
	config := `resource "catalystcenter_wireless_rrm_fra_configuration" "test" {` + "\n"
	config += `	design_name = "Test"` + "\n"
	config += `	radio_band = "2_4GHZ_5GHZ"` + "\n"
	config += `	fra_freeze = true` + "\n"
	config += `	fra_status = true` + "\n"
	config += `	fra_interval = 3` + "\n"
	config += `	fra_sensitivity = "MEDIUM"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
