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
func TestAccCcAssignCredentials(t *testing.T) {
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcAssignCredentialsPrerequisitesConfig + testAccCcAssignCredentialsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcAssignCredentialsPrerequisitesConfig + testAccCcAssignCredentialsConfig_all(),
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
const testAccCcAssignCredentialsPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
  depends_on = [
    catalystcenter_credentials_cli.test,
    catalystcenter_credentials_https_read.test,
  ]
}

resource "catalystcenter_credentials_cli" "test" {
  description = "TestCli1"
  username    = "user1"
  password    = "password1"
}

resource "catalystcenter_credentials_https_read" "test" {
  description = "TestHttpsRead1"
  username    = "user1"
  password    = "password1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAssignCredentialsConfig_minimum() string {
	config := `resource "catalystcenter_assign_credentials" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	cli_id = catalystcenter_credentials_cli.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAssignCredentialsConfig_all() string {
	config := `resource "catalystcenter_assign_credentials" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	cli_id = catalystcenter_credentials_cli.test.id` + "\n"
	config += `	https_read_id = catalystcenter_credentials_https_read.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
