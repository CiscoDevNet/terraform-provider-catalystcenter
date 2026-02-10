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
func TestAccDataSourceCcDot11beProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_dot11be_profile.test", "profile_name", "dot11be_profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_dot11be_profile.test", "ofdma_down_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_dot11be_profile.test", "ofdma_up_link", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_dot11be_profile.test", "mu_mimo_down_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_dot11be_profile.test", "mu_mimo_up_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_dot11be_profile.test", "ofdma_multi_ru", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcDot11beProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcDot11beProfileConfig() string {
	config := `resource "catalystcenter_dot11be_profile" "test" {` + "\n"
	config += `	profile_name = "dot11be_profile_1"` + "\n"
	config += `	ofdma_down_link = true` + "\n"
	config += `	ofdma_up_link = true` + "\n"
	config += `	mu_mimo_down_link = false` + "\n"
	config += `	mu_mimo_up_link = false` + "\n"
	config += `	ofdma_multi_ru = false` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_dot11be_profile" "test" {
			profile_name = "dot11be_profile_1"
			depends_on = [catalystcenter_dot11be_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
