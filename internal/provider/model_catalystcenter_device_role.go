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
type DeviceRole struct {
	Id         types.String `tfsdk:"id"`
	DeviceId   types.String `tfsdk:"device_id"`
	Role       types.String `tfsdk:"role"`
	RoleSource types.String `tfsdk:"role_source"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DeviceRole) getPath() string {
	return "/dna/intent/api/v1/network-device/brief"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DeviceRole) toBody(ctx context.Context, state DeviceRole) string {
	body := ""
	if !data.DeviceId.IsNull() {
		body, _ = sjson.Set(body, "id", data.DeviceId.ValueString())
	}
	if !data.Role.IsNull() {
		body, _ = sjson.Set(body, "role", data.Role.ValueString())
	}
	if !data.RoleSource.IsNull() {
		body, _ = sjson.Set(body, "roleSource", data.RoleSource.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DeviceRole) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("role"); value.Exists() {
		data.Role = types.StringValue(value.String())
	} else {
		data.Role = types.StringNull()
	}
	if value := res.Get("roleSource"); value.Exists() {
		data.RoleSource = types.StringValue(value.String())
	} else {
		data.RoleSource = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DeviceRole) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() && !data.DeviceId.IsNull() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("role"); value.Exists() && !data.Role.IsNull() {
		data.Role = types.StringValue(value.String())
	} else {
		data.Role = types.StringNull()
	}
	if value := res.Get("roleSource"); value.Exists() && !data.RoleSource.IsNull() {
		data.RoleSource = types.StringValue(value.String())
	} else {
		data.RoleSource = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DeviceRole) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Role.IsNull() {
		return false
	}
	if !data.RoleSource.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
