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
func TestAccDataSourceCcDiscovery(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "ip_address_list", "192.168.0.1-192.168.0.99"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "netconf_port", "830"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "protocol_order", "SSH"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "retry", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "timeout_seconds", "5"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcDiscoveryPrerequisitesConfig + testAccDataSourceCcDiscoveryConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcDiscoveryPrerequisitesConfig = `
resource "catalystcenter_credentials_cli" "test" {
  description = "My CLI credentials"
  username    = "user1"
  password    = "password1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcDiscoveryConfig() string {
	config := `resource "catalystcenter_discovery" "test" {` + "\n"
	config += `	discovery_type = "Range"` + "\n"
	config += `	global_credential_id_list = [catalystcenter_credentials_cli.test.id]` + "\n"
	config += `	ip_address_list = "192.168.0.1-192.168.0.99"` + "\n"
	config += `	name = "testing terraform provider"` + "\n"
	config += `	netconf_port = "830"` + "\n"
	config += `	protocol_order = "SSH"` + "\n"
	config += `	retry = 3` + "\n"
	config += `	snmp_auth_passphrase = "8 chars+"` + "\n"
	config += `	snmp_auth_protocol = "SHA"` + "\n"
	config += `	snmp_mode = "AuthPriv"` + "\n"
	config += `	snmp_priv_passphrase = "8 chars+"` + "\n"
	config += `	snmp_priv_protocol = "AES128"` + "\n"
	config += `	snmp_user_name = "testuser"` + "\n"
	config += `	snmp_version = "v3"` + "\n"
	config += `	timeout_seconds = 5` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_discovery" "test" {
			id = catalystcenter_discovery.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
