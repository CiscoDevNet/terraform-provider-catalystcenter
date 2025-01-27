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
func TestAccDataSourceCcTelemetrySettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_telemetry_settings.test", "enable_wired_data_collection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_telemetry_settings.test", "enable_wireless_telemetry", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_telemetry_settings.test", "use_builtin_trap_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_telemetry_settings.test", "use_builtin_syslog_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_telemetry_settings.test", "netflow_collector", "Builtin"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_telemetry_settings.test", "enable_netflow_collector_on_devices", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcTelemetrySettingsPrerequisitesConfig + testAccDataSourceCcTelemetrySettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcTelemetrySettingsPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcTelemetrySettingsConfig() string {
	config := `resource "catalystcenter_telemetry_settings" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	enable_wired_data_collection = true` + "\n"
	config += `	enable_wireless_telemetry = true` + "\n"
	config += `	use_builtin_trap_server = true` + "\n"
	config += `	external_trap_servers = ["10.0.0.1"]` + "\n"
	config += `	use_builtin_syslog_server = true` + "\n"
	config += `	external_syslog_servers = ["10.0.0.1"]` + "\n"
	config += `	netflow_collector = "Builtin"` + "\n"
	config += `	enable_netflow_collector_on_devices = false` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_telemetry_settings" "test" {
			site_id = catalystcenter_area.test.id
			depends_on = [catalystcenter_telemetry_settings.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
