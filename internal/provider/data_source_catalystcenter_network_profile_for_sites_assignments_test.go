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
func TestAccDataSourceCcNetworkProfileForSitesAssignments(t *testing.T) {
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcNetworkProfileForSitesAssignmentsPrerequisitesConfig + testAccDataSourceCcNetworkProfileForSitesAssignmentsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcNetworkProfileForSitesAssignmentsPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_network_profile" "example" {
  name = "Profile1"
  type = "switching"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcNetworkProfileForSitesAssignmentsConfig() string {
	config := `resource "catalystcenter_network_profile_for_sites_assignments" "test" {` + "\n"
	config += `	network_profile_id = catalystcenter_network_profile.example.id` + "\n"
	config += `	items = [{` + "\n"
	config += `	  id = catalystcenter_area.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_network_profile_for_sites_assignments" "test" {
			network_profile_id = catalystcenter_network_profile.example.id
			depends_on = [catalystcenter_network_profile_for_sites_assignments.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
