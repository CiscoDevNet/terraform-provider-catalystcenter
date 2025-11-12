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
func TestAccCcExtranetPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_extranet_policy.test", "extranet_policy_name", "MyExtranetPolicy"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcExtranetPolicyPrerequisitesConfig + testAccCcExtranetPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcExtranetPolicyPrerequisitesConfig + testAccCcExtranetPolicyConfig_all(),
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
const testAccCcExtranetPolicyPrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "Area1"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_fabric_site" "test" {
  site_id                     = catalystcenter_area.test.id
  pub_sub_enabled             = false
  authentication_profile_name = "No Authentication"
}
resource "catalystcenter_fabric_l3_virtual_network" "test" {
  virtual_network_name = "SDA_VN1"
  fabric_ids           = [catalystcenter_fabric_site.test.id]
}
resource "catalystcenter_fabric_l3_virtual_network" "test2" {
  virtual_network_name = "SDA_VN2"
  fabric_ids           = [catalystcenter_fabric_site.test.id]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcExtranetPolicyConfig_minimum() string {
	config := `resource "catalystcenter_extranet_policy" "test" {` + "\n"
	config += `	extranet_policy_name = "MyExtranetPolicy"` + "\n"
	config += `	provider_virtual_network_name = catalystcenter_fabric_l3_virtual_network.test.virtual_network_name` + "\n"
	config += `	subscriber_virtual_network_names = [catalystcenter_fabric_l3_virtual_network.test2.virtual_network_name]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcExtranetPolicyConfig_all() string {
	config := `resource "catalystcenter_extranet_policy" "test" {` + "\n"
	config += `	extranet_policy_name = "MyExtranetPolicy"` + "\n"
	config += `	fabric_ids = [catalystcenter_fabric_site.test.id]` + "\n"
	config += `	provider_virtual_network_name = catalystcenter_fabric_l3_virtual_network.test.virtual_network_name` + "\n"
	config += `	subscriber_virtual_network_names = [catalystcenter_fabric_l3_virtual_network.test2.virtual_network_name]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
