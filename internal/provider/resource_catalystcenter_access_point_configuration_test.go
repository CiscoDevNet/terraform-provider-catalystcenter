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
func TestAccCcAccessPointConfiguration(t *testing.T) {
	if os.Getenv("WIRELESS") == "" {
		t.Skip("skipping test, set environment variable WIRELESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_access_point_configuration.test", "ap_list.0.mac_address", "00:11:22:33:44:55"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_access_point_configuration.test", "ap_list.0.ap_name", "AP-old-name"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_access_point_configuration.test", "ap_list.0.ap_name_new", "AP-new-name"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcAccessPointConfigurationConfig_all(),
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
func testAccCcAccessPointConfigurationConfig_minimum() string {
	config := `resource "catalystcenter_access_point_configuration" "test" {` + "\n"
	config += `	ap_list = [{` + "\n"
	config += `	  mac_address = "00:11:22:33:44:55"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAccessPointConfigurationConfig_all() string {
	config := `resource "catalystcenter_access_point_configuration" "test" {` + "\n"
	config += `	ap_list = [{` + "\n"
	config += `	  mac_address = "00:11:22:33:44:55"` + "\n"
	config += `	  ap_name = "AP-old-name"` + "\n"
	config += `	  ap_name_new = "AP-new-name"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
