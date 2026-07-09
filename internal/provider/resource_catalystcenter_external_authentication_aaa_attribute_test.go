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
func TestAccCcExternalAuthenticationAAAAttribute(t *testing.T) {
	if os.Getenv("SYSTEM") == "" {
		t.Skip("skipping test, set environment variable SYSTEM")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_external_authentication_aaa_attribute.test", "name", "memberOf"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcExternalAuthenticationAAAAttributePrerequisitesConfig + testAccCcExternalAuthenticationAAAAttributeConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_external_authentication_aaa_attribute.test",
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
const testAccCcExternalAuthenticationAAAAttributePrerequisitesConfig = `
resource "catalystcenter_authentication_policy_server" "aaa" {
  authentication_port = 1812
  accounting_port     = 1813
  ip_address          = "198.18.133.111"
  port                = 49
  protocol            = "RADIUS_TACACS"
  retries             = 3
  timeout_seconds     = 4
  role                = "primary"
  shared_secret       = "C1sco12345"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcExternalAuthenticationAAAAttributeConfig_minimum() string {
	config := `resource "catalystcenter_external_authentication_aaa_attribute" "test" {` + "\n"
	config += `	name = "memberOf"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcExternalAuthenticationAAAAttributeConfig_all() string {
	config := `resource "catalystcenter_external_authentication_aaa_attribute" "test" {` + "\n"
	config += `	name = "memberOf"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
