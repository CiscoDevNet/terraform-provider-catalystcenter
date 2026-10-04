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
func TestAccCcApplicationPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.name", "Branch_Office_QoS_Policy_collaboration-apps"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.policy_scope", "Branch_Office_QoS_Policy"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.delete_policy_status", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.advanced_policy_scope_name", "Branch_Office_QoS_Policy"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.clause_type", "BUSINESS_RELEVANCE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.relevance_level", "BUSINESS_RELEVANT"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.device_removal_behavior", "RESTORE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.host_tracking_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.queuing_profile_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application_policy.test", "items.0.application_set_id", "12345678-1234-1234-1234-123456789012"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcApplicationPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcApplicationPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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
func testAccCcApplicationPolicyConfig_minimum() string {
	config := `resource "catalystcenter_application_policy" "test" {` + "\n"
	config += `	policy_scope = "Branch_Office_QoS_Policy"` + "\n"
	config += `	items = [{` + "\n"
	config += `	  name = "Branch_Office_QoS_Policy_collaboration-apps"` + "\n"
	config += `	  policy_scope = "Branch_Office_QoS_Policy"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcApplicationPolicyConfig_all() string {
	config := `resource "catalystcenter_application_policy" "test" {` + "\n"
	config += `	policy_scope = "Branch_Office_QoS_Policy"` + "\n"
	config += `	items = [{` + "\n"
	config += `	  name = "Branch_Office_QoS_Policy_collaboration-apps"` + "\n"
	config += `	  policy_scope = "Branch_Office_QoS_Policy"` + "\n"
	config += `	  priority = "100"` + "\n"
	config += `	  delete_policy_status = "NONE"` + "\n"
	config += `	  advanced_policy_scope_name = "Branch_Office_QoS_Policy"` + "\n"
	config += `	  site_ids = ["12345678-1234-1234-1234-123456789012"]` + "\n"
	config += `	  ssids = ["corp-ssid"]` + "\n"
	config += `	  clause_type = "BUSINESS_RELEVANCE"` + "\n"
	config += `	  relevance_level = "BUSINESS_RELEVANT"` + "\n"
	config += `	  device_removal_behavior = "RESTORE"` + "\n"
	config += `	  host_tracking_enabled = true` + "\n"
	config += `	  queuing_profile_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	  application_set_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
