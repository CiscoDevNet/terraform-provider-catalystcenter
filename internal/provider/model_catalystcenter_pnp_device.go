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
type PnPDevice struct {
	Id           types.String `tfsdk:"id"`
	SerialNumber types.String `tfsdk:"serial_number"`
	Stack        types.Bool   `tfsdk:"stack"`
	Pid          types.String `tfsdk:"pid"`
	Hostname     types.String `tfsdk:"hostname"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PnPDevice) getPath() string {
	return "/dna/intent/api/v1/onboarding/pnp-device"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PnPDevice) toBody(ctx context.Context, state PnPDevice) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SerialNumber.IsNull() {
		body, _ = sjson.Set(body, "deviceInfo.serialNumber", data.SerialNumber.ValueString())
	}
	if !data.Stack.IsNull() {
		body, _ = sjson.Set(body, "deviceInfo.stack", data.Stack.ValueBool())
	}
	if !data.Pid.IsNull() {
		body, _ = sjson.Set(body, "deviceInfo.pid", data.Pid.ValueString())
	}
	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, "deviceInfo.hostname", data.Hostname.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PnPDevice) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("0.deviceInfo.serialNumber"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("0.deviceInfo.stack"); value.Exists() {
		data.Stack = types.BoolValue(value.Bool())
	} else {
		data.Stack = types.BoolNull()
	}
	if value := res.Get("0.deviceInfo.pid"); value.Exists() {
		data.Pid = types.StringValue(value.String())
	} else {
		data.Pid = types.StringNull()
	}
	if value := res.Get("0.deviceInfo.hostname"); value.Exists() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PnPDevice) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceInfo.serialNumber"); value.Exists() && !data.SerialNumber.IsNull() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get("0.deviceInfo.stack"); value.Exists() && !data.Stack.IsNull() {
		data.Stack = types.BoolValue(value.Bool())
	} else {
		data.Stack = types.BoolNull()
	}
	if value := res.Get("0.deviceInfo.pid"); value.Exists() && !data.Pid.IsNull() {
		data.Pid = types.StringValue(value.String())
	} else {
		data.Pid = types.StringNull()
	}
	if value := res.Get("0.deviceInfo.hostname"); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PnPDevice) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Stack.IsNull() {
		return false
	}
	if !data.Pid.IsNull() {
		return false
	}
	if !data.Hostname.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
