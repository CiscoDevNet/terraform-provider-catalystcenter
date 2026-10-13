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
func TestAccCcApplication(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "name", "my-custom-app"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.dscp", "18"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.server_name", "app.example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.url", "example.com/path"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.traffic_class", "TRANSACTIONAL_DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.category_id", "de3f7e1e-bdda-4a2c-87f8-0ff49b463997"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.rank", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_applications.0.engine_id", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_identity.0.protocol", "TCP"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_identity.0.ports", "65529"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_identity.0.lower_port", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_application.test", "network_identity.0.upper_port", "0"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcApplicationPrerequisitesConfig + testAccCcApplicationConfig_all(),
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
const testAccCcApplicationPrerequisitesConfig = `
resource "catalystcenter_application_set" "test" {
  name                       = "tf-acc-application-set"
  default_business_relevance = "BUSINESS_RELEVANT"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcApplicationConfig_minimum() string {
	config := `resource "catalystcenter_application" "test" {` + "\n"
	config += `	name = "my-custom-app"` + "\n"
	config += `	network_applications = [{` + "\n"
	config += `	  help_string = "This is a custom application"` + "\n"
	config += `	  application_type = "CUSTOM"` + "\n"
	config += `	  server_type = "_server-ip"` + "\n"
	config += `	  traffic_class = "TRANSACTIONAL_DATA"` + "\n"
	config += `	  category_id = "de3f7e1e-bdda-4a2c-87f8-0ff49b463997"` + "\n"
	config += `	  ignore_conflict = false` + "\n"
	config += `	  rank = 1` + "\n"
	config += `	  engine_id = 6` + "\n"
	config += `	}]` + "\n"
	config += `	application_set_id = catalystcenter_application_set.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcApplicationConfig_all() string {
	config := `resource "catalystcenter_application" "test" {` + "\n"
	config += `	name = "my-custom-app"` + "\n"
	config += `	network_applications = [{` + "\n"
	config += `	  help_string = "This is a custom application"` + "\n"
	config += `	  application_type = "CUSTOM"` + "\n"
	config += `	  server_type = "_server-ip"` + "\n"
	config += `	  dscp = "18"` + "\n"
	config += `	  app_protocol = "TCP"` + "\n"
	config += `	  server_name = "app.example.com"` + "\n"
	config += `	  url = "example.com/path"` + "\n"
	config += `	  traffic_class = "TRANSACTIONAL_DATA"` + "\n"
	config += `	  category_id = "de3f7e1e-bdda-4a2c-87f8-0ff49b463997"` + "\n"
	config += `	  ignore_conflict = false` + "\n"
	config += `	  rank = 1` + "\n"
	config += `	  engine_id = 6` + "\n"
	config += `	}]` + "\n"
	config += `	application_set_id = catalystcenter_application_set.test.id` + "\n"
	config += `	network_identity = [{` + "\n"
	config += `	  protocol = "TCP"` + "\n"
	config += `	  ports = "65529"` + "\n"
	config += `	  lower_port = 0` + "\n"
	config += `	  upper_port = 0` + "\n"
	config += `	  ipv4_subnet = ["10.0.0.0/24"]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
