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
func TestAccCcImage(t *testing.T) {
	if os.Getenv("TEST_IMAGE_FROM_FILE") == "" {
		t.Skip("skipping test, set environment variable TEST_IMAGE_FROM_FILE")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image.test", "third_party_application_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image.test", "family", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image.test", "name", "basename(\"../software.bin\")"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image.test", "third_party_vendor", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcImageConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcImageConfig_all(),
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
func testAccCcImageConfig_minimum() string {
	config := `resource "catalystcenter_image" "test" {` + "\n"
	config += `	source_path = "../software.bin"` + "\n"
	config += `	name = "basename(\"../software.bin\")"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcImageConfig_all() string {
	config := `resource "catalystcenter_image" "test" {` + "\n"
	config += `	third_party_application_type = ""` + "\n"
	config += `	family = ""` + "\n"
	config += `	source_path = "../software.bin"` + "\n"
	config += `	name = "basename(\"../software.bin\")"` + "\n"
	config += `	third_party_vendor = ""` + "\n"
	config += `	is_third_party = ` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
