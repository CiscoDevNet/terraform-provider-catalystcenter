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
func TestAccDataSourceCcAppPolicyQueuingProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "name", "branch-queuing"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "description", "Branch WAN queuing profile"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.type", "BANDWIDTH"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.is_common_between_all_interface_speeds", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.interface_speed_bandwidth_clauses.0.interface_speed", "ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.interface_speed_bandwidth_clauses.0.tc_bandwidth_settings.0.traffic_class", "VOIP_TELEPHONY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.interface_speed_bandwidth_clauses.0.tc_bandwidth_settings.0.bandwidth_percentage", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.tc_dscp_settings.0.traffic_class", "VOIP_TELEPHONY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_app_policy_queuing_profile.test", "clauses.0.tc_dscp_settings.0.dscp", "46"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcAppPolicyQueuingProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcAppPolicyQueuingProfileConfig() string {
	config := `resource "catalystcenter_app_policy_queuing_profile" "test" {` + "\n"
	config += `	name = "branch-queuing"` + "\n"
	config += `	description = "Branch WAN queuing profile"` + "\n"
	config += `	clauses = [{` + "\n"
	config += `	  type = "BANDWIDTH"` + "\n"
	config += `	  is_common_between_all_interface_speeds = true` + "\n"
	config += `	  interface_speed_bandwidth_clauses = [{` + "\n"
	config += `		interface_speed = "ALL"` + "\n"
	config += `      tc_bandwidth_settings = [{` + "\n"
	config += `			traffic_class = "VOIP_TELEPHONY"` + "\n"
	config += `			bandwidth_percentage = 10` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	  tc_dscp_settings = [{` + "\n"
	config += `		traffic_class = "VOIP_TELEPHONY"` + "\n"
	config += `		dscp = "46"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_app_policy_queuing_profile" "test" {
			id = catalystcenter_app_policy_queuing_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
