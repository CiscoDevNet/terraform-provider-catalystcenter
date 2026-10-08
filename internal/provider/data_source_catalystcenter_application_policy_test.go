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
func TestAccDataSourceCcApplicationPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcApplicationPolicyPrerequisitesConfig + testAccDataSourceCcApplicationPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceCcApplicationPolicyPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}

resource "catalystcenter_area" "test" {
  name      = "AppPolicyTestArea"
  parent_id = data.catalystcenter_site.test.id
}

data "catalystcenter_application_set" "test" {
  name = "collaboration-apps"
}

data "catalystcenter_application_policy_queuing_profile" "test" {
  name = "CVD_QUEUING_PROFILE"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcApplicationPolicyConfig() string {
	config := `resource "catalystcenter_application_policy" "test" {` + "\n"
	config += `	policy_scope = "Branch_Office_QoS_Policy"` + "\n"
	config += `	site_ids = [catalystcenter_area.test.id]` + "\n"
	config += `	items = [{ name = "Branch_Office_QoS_Policy_collaboration-apps", priority = "100", delete_policy_status = "NONE", clause_type = "BUSINESS_RELEVANCE", relevance_level = "BUSINESS_RELEVANT", application_set_id = data.catalystcenter_application_set.test.id }, { name = "Branch_Office_QoS_Policy_queuing_customization", priority = "100", delete_policy_status = "NONE", queuing_profile_id = data.catalystcenter_application_policy_queuing_profile.test.id }]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_application_policy" "test" {
			id = catalystcenter_application_policy.test.id
			policy_scope = "Branch_Office_QoS_Policy"
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
