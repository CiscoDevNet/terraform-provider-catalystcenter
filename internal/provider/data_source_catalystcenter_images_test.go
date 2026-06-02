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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcImages(t *testing.T) {
	if os.Getenv("IMAGES") == "" {
		t.Skip("skipping test, set environment variable IMAGES")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.id", "faa9c5f7-d093-459a-8164-cc555bbf3b80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.name", "cat9k_iosxe.17.12.01.SPA.bin"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.family", "Cisco Catalyst 9300 Switch"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.vendor", "CISCO"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.version", "17.12.01"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.application_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_images.test", "images.0.image_type", ""))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcImagesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

func testAccDataSourceCcImagesConfig() string {

	config := `
		data "catalystcenter_images" "test" {
		}
	`
	return config
}
