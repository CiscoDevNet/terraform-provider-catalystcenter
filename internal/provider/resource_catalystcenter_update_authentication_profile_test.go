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
func TestAccCcUpdateAuthenticationProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_authentication_profile.test", "authentication_profile_id", "cc0b30bc-94e7-458a-80e2-c7bbecc829f5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_authentication_profile.test", "authentication_profile_name", "Closed Authentication"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_authentication_profile.test", "authentication_order", "mac"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_authentication_profile.test", "dot1x_to_mab_fallback_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_authentication_profile.test", "wake_on_lan", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_authentication_profile.test", "number_of_hosts", "Unlimited"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcUpdateAuthenticationProfileConfig_all(),
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
func testAccCcUpdateAuthenticationProfileConfig_minimum() string {
	config := `resource "catalystcenter_update_authentication_profile" "test" {` + "\n"
	config += `	authentication_profile_id = "cc0b30bc-94e7-458a-80e2-c7bbecc829f5"` + "\n"
	config += `	authentication_profile_name = "Closed Authentication"` + "\n"
	config += `	authentication_order = "mac"` + "\n"
	config += `	dot1x_to_mab_fallback_timeout = 10` + "\n"
	config += `	wake_on_lan = true` + "\n"
	config += `	number_of_hosts = "Unlimited"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcUpdateAuthenticationProfileConfig_all() string {
	config := `resource "catalystcenter_update_authentication_profile" "test" {` + "\n"
	config += `	authentication_profile_id = "cc0b30bc-94e7-458a-80e2-c7bbecc829f5"` + "\n"
	config += `	authentication_profile_name = "Closed Authentication"` + "\n"
	config += `	authentication_order = "mac"` + "\n"
	config += `	dot1x_to_mab_fallback_timeout = 10` + "\n"
	config += `	wake_on_lan = true` + "\n"
	config += `	number_of_hosts = "Unlimited"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
