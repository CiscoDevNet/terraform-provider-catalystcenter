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
func TestAccCcTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "name", "Template1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "device_types.0.product_family", "Switches and Hubs"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "device_types.0.product_series", "Cisco Catalyst 9300 Series Switches"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "device_types.0.product_type", "Cisco Catalyst 9300 Switch"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "language", "JINJA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "software_type", "IOS-XE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "software_variant", "XE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "software_version", "16.12.1a"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_content", "hostname {{hostname}}"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.data_type", "STRING"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.default_value", "ABC"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.description", "My parameter"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.display_name", "Custom hostname"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.instruction_text", "My instructions"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.not_param", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.param_array", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.parameter_name", "hostname"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.required", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_template.test", "template_params.0.selection_type", "SINGLE_SELECT"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcTemplatePrerequisitesConfig + testAccCcTemplateConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcTemplatePrerequisitesConfig + testAccCcTemplateConfig_all(),
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
const testAccCcTemplatePrerequisitesConfig = `
resource "catalystcenter_project" "test" {
  name        = "Project1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcTemplateConfig_minimum() string {
	config := `resource "catalystcenter_template" "test" {` + "\n"
	config += `	project_id = catalystcenter_project.test.id` + "\n"
	config += `	name = "Template1"` + "\n"
	config += `	device_types = [{` + "\n"
	config += `	  product_family = "Switches and Hubs"` + "\n"
	config += `	}]` + "\n"
	config += `	language = "JINJA"` + "\n"
	config += `	software_type = "IOS-XE"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcTemplateConfig_all() string {
	config := `resource "catalystcenter_template" "test" {` + "\n"
	config += `	project_id = catalystcenter_project.test.id` + "\n"
	config += `	name = "Template1"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	device_types = [{` + "\n"
	config += `	  product_family = "Switches and Hubs"` + "\n"
	config += `	  product_series = "Cisco Catalyst 9300 Series Switches"` + "\n"
	config += `	  product_type = "Cisco Catalyst 9300 Switch"` + "\n"
	config += `	}]` + "\n"
	config += `	language = "JINJA"` + "\n"
	config += `	software_type = "IOS-XE"` + "\n"
	config += `	software_variant = "XE"` + "\n"
	config += `	software_version = "16.12.1a"` + "\n"
	config += `	template_content = "hostname {{hostname}}"` + "\n"
	config += `	template_params = [{` + "\n"
	config += `	  data_type = "STRING"` + "\n"
	config += `	  default_value = "ABC"` + "\n"
	config += `	  description = "My parameter"` + "\n"
	config += `	  display_name = "Custom hostname"` + "\n"
	config += `	  instruction_text = "My instructions"` + "\n"
	config += `	  not_param = false` + "\n"
	config += `	  param_array = false` + "\n"
	config += `	  parameter_name = "hostname"` + "\n"
	config += `	  required = false` + "\n"
	config += `	  selection_type = "SINGLE_SELECT"` + "\n"
	config += `	  selection_values = {host1 = "host1"}` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
