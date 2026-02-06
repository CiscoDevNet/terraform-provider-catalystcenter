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
func TestAccDataSourceCcAPProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "ap_profile_name", "AP_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "description", "My AP Profile Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "remote_worker_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "awips_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "awips_forensic_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "pmf_denial_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_ap_profile.test", "mesh_enabled", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcAPProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcAPProfileConfig() string {
	config := `resource "catalystcenter_ap_profile" "test" {` + "\n"
	config += `	ap_profile_name = "AP_Profile_1"` + "\n"
	config += `	description = "My AP Profile Description"` + "\n"
	config += `	remote_worker_enabled = false` + "\n"
	config += `	awips_enabled = false` + "\n"
	config += `	awips_forensic_enabled = false` + "\n"
	config += `	pmf_denial_enabled = false` + "\n"
	config += `	mesh_enabled = false` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_ap_profile" "test" {
			ap_profile_name = "AP_Profile_1"
			depends_on = [catalystcenter_ap_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
