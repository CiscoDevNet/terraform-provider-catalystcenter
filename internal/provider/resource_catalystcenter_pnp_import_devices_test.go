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
func TestAccCcPnPImportDevices(t *testing.T) {
	if os.Getenv("PNP") == "" {
		t.Skip("skipping test, set environment variable PNP")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_import_devices.test", "devices.0.serial_number", "FOC12345678"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_import_devices.test", "devices.0.stack", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_import_devices.test", "devices.0.pid", "C9300-24P"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_import_devices.test", "devices.0.hostname", "switch1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcPnPImportDevicesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcPnPImportDevicesConfig_all(),
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
func testAccCcPnPImportDevicesConfig_minimum() string {
	config := `resource "catalystcenter_pnp_import_devices" "test" {` + "\n"
	config += `	devices = [{` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcPnPImportDevicesConfig_all() string {
	config := `resource "catalystcenter_pnp_import_devices" "test" {` + "\n"
	config += `	devices = [{` + "\n"
	config += `	  serial_number = "FOC12345678"` + "\n"
	config += `	  stack = false` + "\n"
	config += `	  pid = "C9300-24P"` + "\n"
	config += `	  hostname = "switch1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
