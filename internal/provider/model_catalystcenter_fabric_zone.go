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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type FabricZone struct {
	Id                        types.String `tfsdk:"id"`
	SiteId                    types.String `tfsdk:"site_id"`
	AuthenticationProfileName types.String `tfsdk:"authentication_profile_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricZone) getPath() string {
	return "/dna/intent/api/v1/sda/fabricZones"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricZone) toBody(ctx context.Context, state FabricZone) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "0.siteId", data.SiteId.ValueString())
	}
	if !data.AuthenticationProfileName.IsNull() {
		body, _ = sjson.Set(body, "0.authenticationProfileName", data.AuthenticationProfileName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricZone) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("response.0.authenticationProfileName"); value.Exists() {
		data.AuthenticationProfileName = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricZone) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("response.0.authenticationProfileName"); value.Exists() && !data.AuthenticationProfileName.IsNull() {
		data.AuthenticationProfileName = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileName = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricZone) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.AuthenticationProfileName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
