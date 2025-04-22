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
func TestAccCcUpdateDeviceManagementAddress(t *testing.T) {
	if os.Getenv("INVENTORY") == "" {
		t.Skip("skipping test, set environment variable INVENTORY")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_device_management_address.test", "device_id", "65e422375b4b6e40ef3423f8"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_update_device_management_address.test", "new_ip", "192.168.10.1"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcUpdateDeviceManagementAddressConfig_all(),
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
func testAccCcUpdateDeviceManagementAddressConfig_minimum() string {
	config := `resource "catalystcenter_update_device_management_address" "test" {` + "\n"
	config += `	device_id = "65e422375b4b6e40ef3423f8"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcUpdateDeviceManagementAddressConfig_all() string {
	config := `resource "catalystcenter_update_device_management_address" "test" {` + "\n"
	config += `	device_id = "65e422375b4b6e40ef3423f8"` + "\n"
	config += `	new_ip = "192.168.10.1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
