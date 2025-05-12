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
type UpdateAuthenticationProfile struct {
	Id                        types.String `tfsdk:"id"`
	AuthenticationProfileId   types.String `tfsdk:"authentication_profile_id"`
	FabricId                  types.String `tfsdk:"fabric_id"`
	AuthenticationProfileName types.String `tfsdk:"authentication_profile_name"`
	AuthenticationOrder       types.String `tfsdk:"authentication_order"`
	Dot1xToMabFallbackTimeout types.Int64  `tfsdk:"dot1x_to_mab_fallback_timeout"`
	WakeOnLan                 types.Bool   `tfsdk:"wake_on_lan"`
	NumberOfHosts             types.String `tfsdk:"number_of_hosts"`
	IsBpduGuardEnabled        types.Bool   `tfsdk:"is_bpdu_guard_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data UpdateAuthenticationProfile) getPath() string {
	return "/dna/intent/api/v1/sda/authenticationProfiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data UpdateAuthenticationProfile) toBody(ctx context.Context, state UpdateAuthenticationProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.AuthenticationProfileId.IsNull() {
		body, _ = sjson.Set(body, "0.id", data.AuthenticationProfileId.ValueString())
	}
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.AuthenticationProfileName.IsNull() {
		body, _ = sjson.Set(body, "0.authenticationProfileName", data.AuthenticationProfileName.ValueString())
	}
	if !data.AuthenticationOrder.IsNull() {
		body, _ = sjson.Set(body, "0.authenticationOrder", data.AuthenticationOrder.ValueString())
	}
	if !data.Dot1xToMabFallbackTimeout.IsNull() {
		body, _ = sjson.Set(body, "0.dot1xToMabFallbackTimeout", data.Dot1xToMabFallbackTimeout.ValueInt64())
	}
	if !data.WakeOnLan.IsNull() {
		body, _ = sjson.Set(body, "0.wakeOnLan", data.WakeOnLan.ValueBool())
	}
	if !data.NumberOfHosts.IsNull() {
		body, _ = sjson.Set(body, "0.numberOfHosts", data.NumberOfHosts.ValueString())
	}
	if !data.IsBpduGuardEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.isBpduGuardEnabled", data.IsBpduGuardEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *UpdateAuthenticationProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.id"); value.Exists() {
		data.AuthenticationProfileId = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.authenticationProfileName"); value.Exists() {
		data.AuthenticationProfileName = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileName = types.StringNull()
	}
	if value := res.Get("response.0.authenticationOrder"); value.Exists() {
		data.AuthenticationOrder = types.StringValue(value.String())
	} else {
		data.AuthenticationOrder = types.StringNull()
	}
	if value := res.Get("response.0.dot1xToMabFallbackTimeout"); value.Exists() {
		data.Dot1xToMabFallbackTimeout = types.Int64Value(value.Int())
	} else {
		data.Dot1xToMabFallbackTimeout = types.Int64Null()
	}
	if value := res.Get("response.0.wakeOnLan"); value.Exists() {
		data.WakeOnLan = types.BoolValue(value.Bool())
	} else {
		data.WakeOnLan = types.BoolNull()
	}
	if value := res.Get("response.0.numberOfHosts"); value.Exists() {
		data.NumberOfHosts = types.StringValue(value.String())
	} else {
		data.NumberOfHosts = types.StringNull()
	}
	if value := res.Get("response.0.isBpduGuardEnabled"); value.Exists() {
		data.IsBpduGuardEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsBpduGuardEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *UpdateAuthenticationProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.id"); value.Exists() && !data.AuthenticationProfileId.IsNull() {
		data.AuthenticationProfileId = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.authenticationProfileName"); value.Exists() && !data.AuthenticationProfileName.IsNull() {
		data.AuthenticationProfileName = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileName = types.StringNull()
	}
	if value := res.Get("response.0.authenticationOrder"); value.Exists() && !data.AuthenticationOrder.IsNull() {
		data.AuthenticationOrder = types.StringValue(value.String())
	} else {
		data.AuthenticationOrder = types.StringNull()
	}
	if value := res.Get("response.0.dot1xToMabFallbackTimeout"); value.Exists() && !data.Dot1xToMabFallbackTimeout.IsNull() {
		data.Dot1xToMabFallbackTimeout = types.Int64Value(value.Int())
	} else {
		data.Dot1xToMabFallbackTimeout = types.Int64Null()
	}
	if value := res.Get("response.0.wakeOnLan"); value.Exists() && !data.WakeOnLan.IsNull() {
		data.WakeOnLan = types.BoolValue(value.Bool())
	} else {
		data.WakeOnLan = types.BoolNull()
	}
	if value := res.Get("response.0.numberOfHosts"); value.Exists() && !data.NumberOfHosts.IsNull() {
		data.NumberOfHosts = types.StringValue(value.String())
	} else {
		data.NumberOfHosts = types.StringNull()
	}
	if value := res.Get("response.0.isBpduGuardEnabled"); value.Exists() && !data.IsBpduGuardEnabled.IsNull() {
		data.IsBpduGuardEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsBpduGuardEnabled = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *UpdateAuthenticationProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FabricId.IsNull() {
		return false
	}
	if !data.AuthenticationOrder.IsNull() {
		return false
	}
	if !data.Dot1xToMabFallbackTimeout.IsNull() {
		return false
	}
	if !data.WakeOnLan.IsNull() {
		return false
	}
	if !data.NumberOfHosts.IsNull() {
		return false
	}
	if !data.IsBpduGuardEnabled.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
