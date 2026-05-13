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
func TestAccCcPnPDeviceClaimSite(t *testing.T) {
	if os.Getenv("PNP") == "" {
		t.Skip("skipping test, set environment variable PNP")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "device_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "site_id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "type", "Default"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "image_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "image_skip", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "config_id", "template1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "config_parameters.0.name", "HOSTNAME"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_pnp_device_claim_site.test", "config_parameters.0.value", "switch1"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcPnPDeviceClaimSitePrerequisitesConfig + testAccCcPnPDeviceClaimSiteConfig_all(),
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
const testAccCcPnPDeviceClaimSitePrerequisitesConfig = `
data "catalystcenter_site" "test" {
  name_hierarchy = "Global"
}
resource "catalystcenter_area" "test" {
  name      = "Area1"
  parent_id = data.catalystcenter_site.test.id
}
resource "catalystcenter_pnp_device" "test" {
  serial_number = "FOC2724Y8RE"
  pid           = "C9300X-48HX"
  hostname      = "PNP-SW"
  stack         = false
  lifecycle {
    ignore_changes = [hostname, stack, pid, serial_number]
  }
}
resource "catalystcenter_project" "test" {
  name = "Project1"
}
resource "catalystcenter_template" "test" {
  project_id  = catalystcenter_project.test.id
  name        = "Template1"
  description = "My description"
  device_types = [
    {
      product_family = "Switches and Hubs"
      product_series = "Cisco Catalyst 9300 Series Switches"
      product_type   = "Cisco Catalyst 9300 Switch"
    }
  ]
  language         = "JINJA"
  software_type    = "IOS-XE"
  software_variant = "XE"
  software_version = "17.12.4"
  template_content = "hostname {{HOSTNAME}}"
}
resource "catalystcenter_template_version" "test" {
  template_id = catalystcenter_template.test.id
  comments    = "Initial version"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcPnPDeviceClaimSiteConfig_minimum() string {
	config := `resource "catalystcenter_pnp_device_claim_site" "test" {` + "\n"
	config += `	device_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	site_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	type = "Default"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcPnPDeviceClaimSiteConfig_all() string {
	config := `resource "catalystcenter_pnp_device_claim_site" "test" {` + "\n"
	config += `	device_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	site_id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	type = "Default"` + "\n"
	config += `	image_id = ""` + "\n"
	config += `	image_skip = true` + "\n"
	config += `	config_id = "template1"` + "\n"
	config += `	config_parameters = [{` + "\n"
	config += `	  name = "HOSTNAME"` + "\n"
	config += `	  value = "switch1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
