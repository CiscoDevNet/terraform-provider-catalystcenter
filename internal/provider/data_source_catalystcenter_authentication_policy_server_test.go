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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcAuthenticationPolicyServer(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "authentication_port", "1812"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "accounting_port", "1813"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "ip_address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "pxgrid_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "use_dnac_cert_for_pxgrid", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "is_ise_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "port", "49"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "protocol", "RADIUS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "retries", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "role", "secondary"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_authentication_policy_server.test", "timeout_seconds", "2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcAuthenticationPolicyServerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcAuthenticationPolicyServerConfig() string {
	config := `resource "catalystcenter_authentication_policy_server" "test" {` + "\n"
	config += `	authentication_port = 1812` + "\n"
	config += `	accounting_port = 1813` + "\n"
	config += `	ip_address = "10.0.0.1"` + "\n"
	config += `	pxgrid_enabled = true` + "\n"
	config += `	use_dnac_cert_for_pxgrid = false` + "\n"
	config += `	is_ise_enabled = false` + "\n"
	config += `	port = 49` + "\n"
	config += `	protocol = "RADIUS"` + "\n"
	config += `	retries = 2` + "\n"
	config += `	role = "secondary"` + "\n"
	config += `	shared_secret = "Cisco123"` + "\n"
	config += `	timeout_seconds = 2` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_authentication_policy_server" "test" {
			id = catalystcenter_authentication_policy_server.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
