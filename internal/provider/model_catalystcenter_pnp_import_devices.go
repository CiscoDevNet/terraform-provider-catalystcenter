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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PnPImportDevices struct {
	Id      types.String              `tfsdk:"id"`
	Devices []PnPImportDevicesDevices `tfsdk:"devices"`
}

type PnPImportDevicesDevices struct {
	SerialNumber types.String `tfsdk:"serial_number"`
	Stack        types.Bool   `tfsdk:"stack"`
	Pid          types.String `tfsdk:"pid"`
	Hostname     types.String `tfsdk:"hostname"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PnPImportDevices) getPath() string {
	return "/dna/intent/api/v1/onboarding/pnp-device/import"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PnPImportDevices) toBody(ctx context.Context, state PnPImportDevices) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.Devices) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.SerialNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInfo.serialNumber", item.SerialNumber.ValueString())
			}
			if !item.Stack.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInfo.stack", item.Stack.ValueBool())
			}
			if !item.Pid.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInfo.pid", item.Pid.ValueString())
			}
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInfo.hostname", item.Hostname.ValueString())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PnPImportDevices) fromBody(ctx context.Context, res gjson.Result) {

	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]PnPImportDevicesDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PnPImportDevicesDevices{}
			if cValue := v.Get("deviceInfo.serialNumber"); cValue.Exists() {
				item.SerialNumber = types.StringValue(cValue.String())
			} else {
				item.SerialNumber = types.StringNull()
			}
			if cValue := v.Get("deviceInfo.stack"); cValue.Exists() {
				item.Stack = types.BoolValue(cValue.Bool())
			} else {
				item.Stack = types.BoolNull()
			}
			if cValue := v.Get("deviceInfo.pid"); cValue.Exists() {
				item.Pid = types.StringValue(cValue.String())
			} else {
				item.Pid = types.StringNull()
			}
			if cValue := v.Get("deviceInfo.hostname"); cValue.Exists() {
				item.Hostname = types.StringValue(cValue.String())
			} else {
				item.Hostname = types.StringNull()
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PnPImportDevices) updateFromBody(ctx context.Context, res gjson.Result) {

	for i := range data.Devices {
		keys := [...]string{"deviceInfo.serialNumber", "deviceInfo.stack", "deviceInfo.pid", "deviceInfo.hostname"}
		keyValues := [...]string{data.Devices[i].SerialNumber.ValueString(), strconv.FormatBool(data.Devices[i].Stack.ValueBool()), data.Devices[i].Pid.ValueString(), data.Devices[i].Hostname.ValueString()}

		var r gjson.Result
		res.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("deviceInfo.serialNumber"); value.Exists() && !data.Devices[i].SerialNumber.IsNull() {
			data.Devices[i].SerialNumber = types.StringValue(value.String())
		} else {
			data.Devices[i].SerialNumber = types.StringNull()
		}
		if value := r.Get("deviceInfo.stack"); value.Exists() && !data.Devices[i].Stack.IsNull() {
			data.Devices[i].Stack = types.BoolValue(value.Bool())
		} else {
			data.Devices[i].Stack = types.BoolNull()
		}
		if value := r.Get("deviceInfo.pid"); value.Exists() && !data.Devices[i].Pid.IsNull() {
			data.Devices[i].Pid = types.StringValue(value.String())
		} else {
			data.Devices[i].Pid = types.StringNull()
		}
		if value := r.Get("deviceInfo.hostname"); value.Exists() && !data.Devices[i].Hostname.IsNull() {
			data.Devices[i].Hostname = types.StringValue(value.String())
		} else {
			data.Devices[i].Hostname = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PnPImportDevices) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Devices) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
