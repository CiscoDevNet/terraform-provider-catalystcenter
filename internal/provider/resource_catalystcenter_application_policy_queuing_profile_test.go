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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcApplicationPolicyQueuingProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "name", "branch-queuing"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "description", "Branch WAN queuing profile"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.type", "BANDWIDTH"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.is_common_between_all_interface_speeds", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.interface_speed_bandwidth_clauses.0.interface_speed", "ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.interface_speed_bandwidth_clauses.0.tc_bandwidth_settings.0.traffic_class", "VOIP_TELEPHONY"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.interface_speed_bandwidth_clauses.0.tc_bandwidth_settings.0.bandwidth_percentage", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.tc_dscp_settings.0.traffic_class", "VOIP_TELEPHONY"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy_queuing_profile.test", "clauses.0.tc_dscp_settings.0.dscp", "46"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcApplicationPolicyQueuingProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcApplicationPolicyQueuingProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_application_policy_queuing_profile.test",
		ImportState:  true,
		ImportStateIdFunc: func(s *terraform.State) (string, error) {
			rs, ok := s.RootModule().Resources["catalystcenter_application_policy_queuing_profile.test"]
			if !ok {
				return "", fmt.Errorf("resource not found in state")
			}
			return fmt.Sprintf("%s,%s", rs.Primary.Attributes["name"], rs.Primary.ID), nil
		},
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
func testAccCcApplicationPolicyQueuingProfileConfig_minimum() string {
	config := `resource "catalystcenter_application_policy_queuing_profile" "test" {` + "\n"
	config += `	name = "branch-queuing"` + "\n"
	config += `	clauses = [{` + "\n"
	config += `	  type = "BANDWIDTH"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcApplicationPolicyQueuingProfileConfig_all() string {
	config := `resource "catalystcenter_application_policy_queuing_profile" "test" {` + "\n"
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
	return config
}

// End of section. //template:end testAccConfigAll
