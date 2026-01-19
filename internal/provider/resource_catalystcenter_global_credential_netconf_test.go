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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcGlobalCredentialNETCONF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_global_credential_netconf.test", "netconf_port", "830"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_global_credential_netconf.test", "comments", "My NETCONF Credential"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_global_credential_netconf.test", "credential_type", "GLOBAL"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_global_credential_netconf.test", "description", "NETCONF credential for network devices"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcGlobalCredentialNETCONFConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_global_credential_netconf.test",
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
func testAccCcGlobalCredentialNETCONFConfig_minimum() string {
	config := `resource "catalystcenter_global_credential_netconf" "test" {` + "\n"
	config += `	netconf_port = "830"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcGlobalCredentialNETCONFConfig_all() string {
	config := `resource "catalystcenter_global_credential_netconf" "test" {` + "\n"
	config += `	netconf_port = "830"` + "\n"
	config += `	comments = "My NETCONF Credential"` + "\n"
	config += `	credential_type = "GLOBAL"` + "\n"
	config += `	description = "NETCONF credential for network devices"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
