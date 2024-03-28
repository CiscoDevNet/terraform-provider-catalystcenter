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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcDeviceDetail(t *testing.T) {
	if os.Getenv("INVENTORY") == "" {
		t.Skip("skipping test, set environment variable INVENTORY")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "policy_tag_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "nw_device_role", "CORE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "serial_number", "FXS2000Q2T0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "nw_device_name", "device1.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "device_group_hierarchy_id", "/360b1804-969f-4eab-a4ba-9832ea3f1731/26d4bbe6-f908-4b83-86f1-49bfbb1d05f6/"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "cpu", "1.25"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "nw_device_id", "e0ba1a00-b69b-45aa-8c13-4cdfb59afe65"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "site_hierarchy_graph_id", "/2b0a78ee-482e-4b4d-ae85-df289873cbbb/76b11b6a-d94a-431b-8bab-fd16b09f5d40/adad4a8a-17ec-4be3-a4b5-b6549b840afe/e8ce8788-9b13-46ec-86c8-740f7ea443c1/"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "nw_device_family", "Switches and Hubs"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "mac_address", "5C:71:0D:40:50:60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "device_series", "Cisco Catalyst 9400 Series Switches"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "collection_status", "SUCCESS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "maintenance_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "software_version", "17.12.20230427:145516"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "overall_health", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "management_ip_address", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "memory", "39.186811767835785"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "communication_state", "REACHABLE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "nw_device_type", "Cisco Catalyst 9407R Switch"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "platform_id", "C9KV-UADP-8P"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "location", "Global/USA/New York/DC01"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "ha_status", "Non-redundant"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_detail.test", "os_type", "IOS-XE"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcDeviceDetailConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcDeviceDetailConfig() string {
	config := `resource "catalystcenter_device_detail" "test" {` + "\n"
	config += `	policy_tag_name = ""` + "\n"
	config += `	nw_device_role = "CORE"` + "\n"
	config += `	serial_number = "FXS2000Q2T0"` + "\n"
	config += `	nw_device_name = "device1.example.com"` + "\n"
	config += `	device_group_hierarchy_id = "/360b1804-969f-4eab-a4ba-9832ea3f1731/26d4bbe6-f908-4b83-86f1-49bfbb1d05f6/"` + "\n"
	config += `	cpu = "1.25"` + "\n"
	config += `	nw_device_id = "e0ba1a00-b69b-45aa-8c13-4cdfb59afe65"` + "\n"
	config += `	site_hierarchy_graph_id = "/2b0a78ee-482e-4b4d-ae85-df289873cbbb/76b11b6a-d94a-431b-8bab-fd16b09f5d40/adad4a8a-17ec-4be3-a4b5-b6549b840afe/e8ce8788-9b13-46ec-86c8-740f7ea443c1/"` + "\n"
	config += `	nw_device_family = "Switches and Hubs"` + "\n"
	config += `	mac_address = "5C:71:0D:40:50:60"` + "\n"
	config += `	device_series = "Cisco Catalyst 9400 Series Switches"` + "\n"
	config += `	collection_status = "SUCCESS"` + "\n"
	config += `	maintenance_mode = false` + "\n"
	config += `	software_version = "17.12.20230427:145516"` + "\n"
	config += `	tag_id_list = [""]` + "\n"
	config += `	overall_health = 10` + "\n"
	config += `	management_ip_address = "10.10.10.1"` + "\n"
	config += `	memory = "39.186811767835785"` + "\n"
	config += `	communication_state = "REACHABLE"` + "\n"
	config += `	nw_device_type = "Cisco Catalyst 9407R Switch"` + "\n"
	config += `	platform_id = "C9KV-UADP-8P"` + "\n"
	config += `	location = "Global/USA/New York/DC01"` + "\n"
	config += `	ha_status = "Non-redundant"` + "\n"
	config += `	os_type = "IOS-XE"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_device_detail" "test" {
			id = catalystcenter_device_detail.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
