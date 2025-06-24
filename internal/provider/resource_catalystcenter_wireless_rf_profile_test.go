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
func TestAccCcWirelessRFProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "rf_profile_name", "RF_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "default_rf_profile", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "enable_radio_type_a", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "enable_radio_type_b", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_rf_profile.test", "enable_radio_type6_g_hz", "false"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessRFProfileConfig_all(),
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
func testAccCcWirelessRFProfileConfig_minimum() string {
	config := `resource "catalystcenter_wireless_rf_profile" "test" {` + "\n"
	config += `	rf_profile_name = "RF_Profile_1"` + "\n"
	config += `	default_rf_profile = false` + "\n"
	config += `	enable_radio_type_a = true` + "\n"
	config += `	enable_radio_type_b = false` + "\n"
	config += `	enable_radio_type6_g_hz = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcWirelessRFProfileConfig_all() string {
	config := `resource "catalystcenter_wireless_rf_profile" "test" {` + "\n"
	config += `	rf_profile_name = "RF_Profile_1"` + "\n"
	config += `	default_rf_profile = false` + "\n"
	config += `	enable_radio_type_a = true` + "\n"
	config += `	enable_radio_type_b = false` + "\n"
	config += `	enable_radio_type6_g_hz = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
