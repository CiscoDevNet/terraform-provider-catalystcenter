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

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceCcNetworkDevices(t *testing.T) {
	if os.Getenv("INVENTORY") == "" {
		t.Skip("skipping test, set environment variable INVENTORY")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network_devices.test", "response.0.id", "e0ba1a00-b69b-45aa-8c13-4cdfb59afe65"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network_devices.test", "response.0.hostname", "sw2.ciscotest.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network_devices.test", "response.0.management_ip_address", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network_devices.test", "response.0.platform_id", "C9KV-UADP-8P"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network_devices.test", "response.0.role", "CORE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_network_devices.test", "response.0.software_type", "IOS-XE"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcNetworkDevicesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceCcNetworkDevicesConfig() string {
	config := `
		data "catalystcenter_network_devices" "test" {
		}
	`
	return config
}
