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
func TestAccCcRoutingProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_routing_profile.test", "name", "Profile1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_routing_profile.test", "profile_attributes.0.key", "nfv.device.name"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_routing_profile.test", "profile_attributes.0.value", "ISR"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_routing_profile.test", "profile_attributes.0.attributes.0.key", "nfv.device.deviceType"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_routing_profile.test", "profile_attributes.0.attributes.0.value", "Cisco Catalyst 8000V Edge Software"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcRoutingProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcRoutingProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_routing_profile.test",
		ImportState:  true,
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
func testAccCcRoutingProfileConfig_minimum() string {
	config := `resource "catalystcenter_routing_profile" "test" {` + "\n"
	config += `	name = "Profile1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcRoutingProfileConfig_all() string {
	config := `resource "catalystcenter_routing_profile" "test" {` + "\n"
	config += `	name = "Profile1"` + "\n"
	config += `	profile_attributes = [{` + "\n"
	config += `	  key = "nfv.device.name"` + "\n"
	config += `	  value = "ISR"` + "\n"
	config += `	  attributes = [{` + "\n"
	config += `		key = "nfv.device.deviceType"` + "\n"
	config += `		value = "Cisco Catalyst 8000V Edge Software"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
