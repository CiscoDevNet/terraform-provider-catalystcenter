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
func TestAccDataSourceCcApplyPendingFabricEvents(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_apply_pending_fabric_events.test", "fabric_id", "2a3b4c5d-6789-4abc-90de-f1234567890a"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_apply_pending_fabric_events.test", "event_id", ""))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcApplyPendingFabricEventsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcApplyPendingFabricEventsConfig() string {
	config := `resource "catalystcenter_apply_pending_fabric_events" "test" {` + "\n"
	config += `	fabric_id = "2a3b4c5d-6789-4abc-90de-f1234567890a"` + "\n"
	config += `	event_id = ""` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_apply_pending_fabric_events" "test" {
			id = catalystcenter_apply_pending_fabric_events.test.id
			fabric_id = "2a3b4c5d-6789-4abc-90de-f1234567890a"
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
