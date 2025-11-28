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
func TestAccCcIntegrateISE(t *testing.T) {
	if os.Getenv("SYSTEM") == "" {
		t.Skip("skipping test, set environment variable SYSTEM")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_integrate_ise.test", "ise_instance_id", "0df326df-3609-417d-bbfc-fdffe125a172"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcIntegrateISEPrerequisitesConfig + testAccCcIntegrateISEConfig_all(),
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
const testAccCcIntegrateISEPrerequisitesConfig = `
resource "catalystcenter_authentication_policy_server" "ise" {
  authentication_port      = 1812
  accounting_port          = 1813
  ip_address               = "198.18.133.27"
  pxgrid_enabled           = true
  use_dnac_cert_for_pxgrid = false
  is_ise_enabled           = true
  port                     = 49
  protocol                 = "RADIUS"
  retries                  = 3
  timeout_seconds          = 4
  role                     = "primary"
  shared_secret            = "C1sco12345"
  cisco_ise_dtos = [
    {
      user_name       = "admin"
      password        = "C1sco12345"
      fqdn            = "198.18.133.27"
      ip_address      = "198.18.133.27"
      subscriber_name = "198.18.133.27"
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcIntegrateISEConfig_minimum() string {
	config := `resource "catalystcenter_integrate_ise" "test" {` + "\n"
	config += `	ise_instance_id = "0df326df-3609-417d-bbfc-fdffe125a172"` + "\n"
	config += `	is_cert_accepted_by_user = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcIntegrateISEConfig_all() string {
	config := `resource "catalystcenter_integrate_ise" "test" {` + "\n"
	config += `	ise_instance_id = "0df326df-3609-417d-bbfc-fdffe125a172"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
