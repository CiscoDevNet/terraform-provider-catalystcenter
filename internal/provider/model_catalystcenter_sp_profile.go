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
type SPProfile struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Model       types.String `tfsdk:"model"`
	WanProvider types.String `tfsdk:"wan_provider"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SPProfile) getPath() string {
	return "/dna/intent/api/v2/service-provider"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data SPProfile) getPathDelete() string {
	return "/dna/intent/api/v2/sp-profile"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SPProfile) toBody(ctx context.Context, state SPProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "settings.qos.0.profileName", data.Name.ValueString())
	}
	if !data.Model.IsNull() {
		body, _ = sjson.Set(body, "settings.qos.0.model", data.Model.ValueString())
	}
	if !data.WanProvider.IsNull() {
		body, _ = sjson.Set(body, "settings.qos.0.wanProvider", data.WanProvider.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SPProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("spProfileName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("slaProfileName"); value.Exists() {
		data.Model = types.StringValue(value.String())
	} else {
		data.Model = types.StringNull()
	}
	if value := res.Get("wanProvider"); value.Exists() {
		data.WanProvider = types.StringValue(value.String())
	} else {
		data.WanProvider = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SPProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("spProfileName"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("slaProfileName"); value.Exists() && !data.Model.IsNull() {
		data.Model = types.StringValue(value.String())
	} else {
		data.Model = types.StringNull()
	}
	if value := res.Get("wanProvider"); value.Exists() && !data.WanProvider.IsNull() {
		data.WanProvider = types.StringValue(value.String())
	} else {
		data.WanProvider = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SPProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Model.IsNull() {
		return false
	}
	if !data.WanProvider.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
