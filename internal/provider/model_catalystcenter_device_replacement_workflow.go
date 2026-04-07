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
type DeviceReplacementWorkflow struct {
	Id                            types.String `tfsdk:"id"`
	FaultyDeviceSerialNumber      types.String `tfsdk:"faulty_device_serial_number"`
	ReplacementDeviceSerialNumber types.String `tfsdk:"replacement_device_serial_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DeviceReplacementWorkflow) getPath() string {
	return "/dna/intent/api/v1/device-replacement/workflow"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DeviceReplacementWorkflow) toBody(ctx context.Context, state DeviceReplacementWorkflow) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.FaultyDeviceSerialNumber.IsNull() {
		body, _ = sjson.Set(body, "faultyDeviceSerialNumber", data.FaultyDeviceSerialNumber.ValueString())
	}
	if !data.ReplacementDeviceSerialNumber.IsNull() {
		body, _ = sjson.Set(body, "replacementDeviceSerialNumber", data.ReplacementDeviceSerialNumber.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DeviceReplacementWorkflow) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("faultyDeviceSerialNumber"); value.Exists() {
		data.FaultyDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.FaultyDeviceSerialNumber = types.StringNull()
	}
	if value := res.Get("replacementDeviceSerialNumber"); value.Exists() {
		data.ReplacementDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.ReplacementDeviceSerialNumber = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DeviceReplacementWorkflow) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("faultyDeviceSerialNumber"); value.Exists() && !data.FaultyDeviceSerialNumber.IsNull() {
		data.FaultyDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.FaultyDeviceSerialNumber = types.StringNull()
	}
	if value := res.Get("replacementDeviceSerialNumber"); value.Exists() && !data.ReplacementDeviceSerialNumber.IsNull() {
		data.ReplacementDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.ReplacementDeviceSerialNumber = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DeviceReplacementWorkflow) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.ReplacementDeviceSerialNumber.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
