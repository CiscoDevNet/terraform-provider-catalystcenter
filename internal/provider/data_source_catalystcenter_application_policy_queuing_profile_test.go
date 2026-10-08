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
func TestAccDataSourceCcApplicationPolicyQueuingProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_application_policy_queuing_profile.test", "name", "branch-queuing"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_application_policy_queuing_profile.test", "description", "Branch WAN queuing profile"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcApplicationPolicyQueuingProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcApplicationPolicyQueuingProfileConfig() string {
	config := `resource "catalystcenter_application_policy_queuing_profile" "test" {` + "\n"
	config += `	name = "branch-queuing"` + "\n"
	config += `	description = "Branch WAN queuing profile"` + "\n"
	config += `	clauses = [{ type = "BANDWIDTH", is_common_between_all_interface_speeds = false, interface_speed_bandwidth_clauses = [{ interface_speed = "HUNDRED_GBPS", tc_bandwidth_settings = [{ traffic_class = "VOIP_TELEPHONY", bandwidth_percentage = 10 }, { traffic_class = "BROADCAST_VIDEO", bandwidth_percentage = 10 }, { traffic_class = "REAL_TIME_INTERACTIVE", bandwidth_percentage = 13 }, { traffic_class = "MULTIMEDIA_CONFERENCING", bandwidth_percentage = 10 }, { traffic_class = "MULTIMEDIA_STREAMING", bandwidth_percentage = 10 }, { traffic_class = "NETWORK_CONTROL", bandwidth_percentage = 3 }, { traffic_class = "SIGNALING", bandwidth_percentage = 2 }, { traffic_class = "OPS_ADMIN_MGMT", bandwidth_percentage = 5 }, { traffic_class = "TRANSACTIONAL_DATA", bandwidth_percentage = 10 }, { traffic_class = "BULK_DATA", bandwidth_percentage = 4 }, { traffic_class = "SCAVENGER", bandwidth_percentage = 1 }, { traffic_class = "BEST_EFFORT", bandwidth_percentage = 22 }] }, { interface_speed = "TEN_GBPS", tc_bandwidth_settings = [{ traffic_class = "VOIP_TELEPHONY", bandwidth_percentage = 10 }, { traffic_class = "BROADCAST_VIDEO", bandwidth_percentage = 10 }, { traffic_class = "REAL_TIME_INTERACTIVE", bandwidth_percentage = 13 }, { traffic_class = "MULTIMEDIA_CONFERENCING", bandwidth_percentage = 10 }, { traffic_class = "MULTIMEDIA_STREAMING", bandwidth_percentage = 10 }, { traffic_class = "NETWORK_CONTROL", bandwidth_percentage = 3 }, { traffic_class = "SIGNALING", bandwidth_percentage = 2 }, { traffic_class = "OPS_ADMIN_MGMT", bandwidth_percentage = 5 }, { traffic_class = "TRANSACTIONAL_DATA", bandwidth_percentage = 10 }, { traffic_class = "BULK_DATA", bandwidth_percentage = 4 }, { traffic_class = "SCAVENGER", bandwidth_percentage = 1 }, { traffic_class = "BEST_EFFORT", bandwidth_percentage = 22 }] }, { interface_speed = "ONE_GBPS", tc_bandwidth_settings = [{ traffic_class = "VOIP_TELEPHONY", bandwidth_percentage = 10 }, { traffic_class = "BROADCAST_VIDEO", bandwidth_percentage = 10 }, { traffic_class = "REAL_TIME_INTERACTIVE", bandwidth_percentage = 13 }, { traffic_class = "MULTIMEDIA_CONFERENCING", bandwidth_percentage = 10 }, { traffic_class = "MULTIMEDIA_STREAMING", bandwidth_percentage = 10 }, { traffic_class = "NETWORK_CONTROL", bandwidth_percentage = 3 }, { traffic_class = "SIGNALING", bandwidth_percentage = 2 }, { traffic_class = "OPS_ADMIN_MGMT", bandwidth_percentage = 5 }, { traffic_class = "TRANSACTIONAL_DATA", bandwidth_percentage = 10 }, { traffic_class = "BULK_DATA", bandwidth_percentage = 4 }, { traffic_class = "SCAVENGER", bandwidth_percentage = 1 }, { traffic_class = "BEST_EFFORT", bandwidth_percentage = 22 }] }, { interface_speed = "HUNDRED_MBPS", tc_bandwidth_settings = [{ traffic_class = "VOIP_TELEPHONY", bandwidth_percentage = 10 }, { traffic_class = "BROADCAST_VIDEO", bandwidth_percentage = 10 }, { traffic_class = "REAL_TIME_INTERACTIVE", bandwidth_percentage = 13 }, { traffic_class = "MULTIMEDIA_CONFERENCING", bandwidth_percentage = 10 }, { traffic_class = "MULTIMEDIA_STREAMING", bandwidth_percentage = 10 }, { traffic_class = "NETWORK_CONTROL", bandwidth_percentage = 3 }, { traffic_class = "SIGNALING", bandwidth_percentage = 2 }, { traffic_class = "OPS_ADMIN_MGMT", bandwidth_percentage = 5 }, { traffic_class = "TRANSACTIONAL_DATA", bandwidth_percentage = 10 }, { traffic_class = "BULK_DATA", bandwidth_percentage = 4 }, { traffic_class = "SCAVENGER", bandwidth_percentage = 1 }, { traffic_class = "BEST_EFFORT", bandwidth_percentage = 22 }] }, { interface_speed = "TEN_MBPS", tc_bandwidth_settings = [{ traffic_class = "VOIP_TELEPHONY", bandwidth_percentage = 10 }, { traffic_class = "BROADCAST_VIDEO", bandwidth_percentage = 10 }, { traffic_class = "REAL_TIME_INTERACTIVE", bandwidth_percentage = 13 }, { traffic_class = "MULTIMEDIA_CONFERENCING", bandwidth_percentage = 10 }, { traffic_class = "MULTIMEDIA_STREAMING", bandwidth_percentage = 10 }, { traffic_class = "NETWORK_CONTROL", bandwidth_percentage = 3 }, { traffic_class = "SIGNALING", bandwidth_percentage = 2 }, { traffic_class = "OPS_ADMIN_MGMT", bandwidth_percentage = 5 }, { traffic_class = "TRANSACTIONAL_DATA", bandwidth_percentage = 10 }, { traffic_class = "BULK_DATA", bandwidth_percentage = 4 }, { traffic_class = "SCAVENGER", bandwidth_percentage = 1 }, { traffic_class = "BEST_EFFORT", bandwidth_percentage = 22 }] }, { interface_speed = "ONE_MBPS", tc_bandwidth_settings = [{ traffic_class = "VOIP_TELEPHONY", bandwidth_percentage = 10 }, { traffic_class = "BROADCAST_VIDEO", bandwidth_percentage = 10 }, { traffic_class = "REAL_TIME_INTERACTIVE", bandwidth_percentage = 13 }, { traffic_class = "MULTIMEDIA_CONFERENCING", bandwidth_percentage = 10 }, { traffic_class = "MULTIMEDIA_STREAMING", bandwidth_percentage = 10 }, { traffic_class = "NETWORK_CONTROL", bandwidth_percentage = 3 }, { traffic_class = "SIGNALING", bandwidth_percentage = 2 }, { traffic_class = "OPS_ADMIN_MGMT", bandwidth_percentage = 5 }, { traffic_class = "TRANSACTIONAL_DATA", bandwidth_percentage = 10 }, { traffic_class = "BULK_DATA", bandwidth_percentage = 4 }, { traffic_class = "SCAVENGER", bandwidth_percentage = 1 }, { traffic_class = "BEST_EFFORT", bandwidth_percentage = 22 }] }] }, { type = "DSCP_CUSTOMIZATION", tc_dscp_settings = [{ traffic_class = "VOIP_TELEPHONY", dscp = "46" }, { traffic_class = "BROADCAST_VIDEO", dscp = "40" }, { traffic_class = "REAL_TIME_INTERACTIVE", dscp = "32" }, { traffic_class = "MULTIMEDIA_CONFERENCING", dscp = "34" }, { traffic_class = "MULTIMEDIA_STREAMING", dscp = "26" }, { traffic_class = "NETWORK_CONTROL", dscp = "48" }, { traffic_class = "SIGNALING", dscp = "24" }, { traffic_class = "OPS_ADMIN_MGMT", dscp = "16" }, { traffic_class = "TRANSACTIONAL_DATA", dscp = "18" }, { traffic_class = "BULK_DATA", dscp = "10" }, { traffic_class = "SCAVENGER", dscp = "8" }, { traffic_class = "BEST_EFFORT", dscp = "0" }] }]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_application_policy_queuing_profile" "test" {
			id = catalystcenter_application_policy_queuing_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
