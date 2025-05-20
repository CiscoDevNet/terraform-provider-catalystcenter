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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcAssignDevicesToTag(t *testing.T) {
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcAssignDevicesToTagPrerequisitesConfig + testAccCcAssignDevicesToTagConfig_all(),
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
const testAccCcAssignDevicesToTagPrerequisitesConfig = `
resource "catalystcenter_tag" "test" {
  name        = "Tag1"
  description = "Tag1 Description"
  system_tag  = false
}
resource "catalystcenter_device" "test" {
  cli_transport           = "ssh"
  compute_device          = false
  enable_password         = "cisco123"
  extended_discovery_info = "DISCOVER_WITH_CANNED_DATA"
  http_password           = "cisco123"
  http_port               = "80"
  http_secure             = true
  http_user_name          = "admin"
  ip_address              = "1.2.3.4"
  meraki_org_ids          = ["12345678901234567890"]
  netconf_port            = "830"
  password                = "cisco123"
  serial_number           = "FOC12345678"
  snmp_auth_passphrase    = "cisco123"
  snmp_auth_protocol      = "sha"
  snmp_mode               = "authPriv"
  snmp_priv_passphrase    = "cisco123"
  snmp_priv_protocol      = "AES128"
  snmp_ro_community       = "com1"
  snmp_rw_community       = "com2"
  snmp_retry              = 3
  snmp_timeout            = 10
  snmp_user_name          = "admin"
  snmp_version            = "v3"
  type                    = "NETWORK_DEVICE"
  user_name               = "admin"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAssignDevicesToTagConfig_minimum() string {
	config := `resource "catalystcenter_assign_devices_to_tag" "test" {` + "\n"
	config += `	tag_id = catalystcenter_tag.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAssignDevicesToTagConfig_all() string {
	config := `resource "catalystcenter_assign_devices_to_tag" "test" {` + "\n"
	config += `	tag_id = catalystcenter_tag.test.id` + "\n"
	config += `	device_ids = [catalystcenter_device.test.id]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
