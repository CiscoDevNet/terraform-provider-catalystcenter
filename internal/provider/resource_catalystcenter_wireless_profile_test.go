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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAcc
func TestAccCcWirelessProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_profile.test", "name", "Wireless_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_profile.test", "ssid_details.0.name", "mySSID1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_profile.test", "ssid_details.0.enable_fabric", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_wireless_profile.test", "ssid_details.0.enable_flex_connect", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcWirelessProfilePrerequisitesConfig + testAccCcWirelessProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcWirelessProfilePrerequisitesConfig + testAccCcWirelessProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

//template:end testAcc

//template:begin testPrerequisites
const testAccCcWirelessProfilePrerequisitesConfig = `
resource "catalystcenter_wireless_enterprise_ssid" "example" {
  name                                  = "mySSID1"
  security_level                        = "wpa3_enterprise"
  passphrase                            = "Cisco123"
}

`

//template:end testPrerequisites

//template:begin testAccConfigMinimal
func testAccCcWirelessProfileConfig_minimum() string {
	config := `resource "catalystcenter_wireless_profile" "test" {` + "\n"
	config += `	name = "Wireless_Profile_1"` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccCcWirelessProfileConfig_all() string {
	config := `resource "catalystcenter_wireless_profile" "test" {` + "\n"
	config += `	name = "Wireless_Profile_1"` + "\n"
	config += `	ssid_details = [{` + "\n"
	config += `	  name = "mySSID1"` + "\n"
	config += `	  enable_fabric = true` + "\n"
	config += `	  enable_flex_connect = false` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll