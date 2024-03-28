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
func TestAccCcDevice(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "cli_transport", "ssh"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "compute_device", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "enable_password", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "extended_discovery_info", "DISCOVER_WITH_CANNED_DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "http_password", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "http_port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "http_secure", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "http_user_name", "admin"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "netconf_port", "830"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "password", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "serial_number", "FOC12345678"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_auth_passphrase", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_auth_protocol", "sha"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_mode", "authPriv"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_priv_passphrase", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_priv_protocol", "AES128"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_ro_community", "com1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_rw_community", "com2"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_retry", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_user_name", "admin"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "snmp_version", "v3"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "type", "NETWORK_DEVICE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_device.test", "user_name", "admin"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcDeviceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcDeviceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_device.test",
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcDeviceConfig_minimum() string {
	config := `resource "catalystcenter_device" "test" {` + "\n"
	config += `	cli_transport = "ssh"` + "\n"
	config += `	enable_password = "cisco123"` + "\n"
	config += `	ip_address = "1.2.3.4"` + "\n"
	config += `	password = "cisco123"` + "\n"
	config += `	snmp_auth_passphrase = "cisco123"` + "\n"
	config += `	snmp_auth_protocol = "sha"` + "\n"
	config += `	snmp_mode = "authPriv"` + "\n"
	config += `	snmp_priv_passphrase = "cisco123"` + "\n"
	config += `	snmp_priv_protocol = "AES128"` + "\n"
	config += `	snmp_ro_community = "com1"` + "\n"
	config += `	snmp_rw_community = "com2"` + "\n"
	config += `	snmp_retry = 3` + "\n"
	config += `	snmp_timeout = 10` + "\n"
	config += `	snmp_user_name = "admin"` + "\n"
	config += `	snmp_version = "v3"` + "\n"
	config += `	type = "NETWORK_DEVICE"` + "\n"
	config += `	user_name = "admin"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcDeviceConfig_all() string {
	config := `resource "catalystcenter_device" "test" {` + "\n"
	config += `	cli_transport = "ssh"` + "\n"
	config += `	compute_device = false` + "\n"
	config += `	enable_password = "cisco123"` + "\n"
	config += `	extended_discovery_info = "DISCOVER_WITH_CANNED_DATA"` + "\n"
	config += `	http_password = "cisco123"` + "\n"
	config += `	http_port = "80"` + "\n"
	config += `	http_secure = true` + "\n"
	config += `	http_user_name = "admin"` + "\n"
	config += `	ip_address = "1.2.3.4"` + "\n"
	config += `	meraki_org_ids = ["12345678901234567890"]` + "\n"
	config += `	netconf_port = "830"` + "\n"
	config += `	password = "cisco123"` + "\n"
	config += `	serial_number = "FOC12345678"` + "\n"
	config += `	snmp_auth_passphrase = "cisco123"` + "\n"
	config += `	snmp_auth_protocol = "sha"` + "\n"
	config += `	snmp_mode = "authPriv"` + "\n"
	config += `	snmp_priv_passphrase = "cisco123"` + "\n"
	config += `	snmp_priv_protocol = "AES128"` + "\n"
	config += `	snmp_ro_community = "com1"` + "\n"
	config += `	snmp_rw_community = "com2"` + "\n"
	config += `	snmp_retry = 3` + "\n"
	config += `	snmp_timeout = 10` + "\n"
	config += `	snmp_user_name = "admin"` + "\n"
	config += `	snmp_version = "v3"` + "\n"
	config += `	type = "NETWORK_DEVICE"` + "\n"
	config += `	user_name = "admin"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
