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

type NetworkDevices struct {
	Devices []NetworkDevicesDevices `tfsdk:"devices"`
}

type NetworkDevicesDevices struct {
	Id                  types.String `tfsdk:"id"`
	Hostname            types.String `tfsdk:"hostname"`
	ManagementIpAddress types.String `tfsdk:"management_ip_address"`
	ManagementState     types.String `tfsdk:"management_state"`
	PlatformId          types.String `tfsdk:"platform_id"`
	Role                types.String `tfsdk:"role"`
	SoftwareType        types.String `tfsdk:"software_type"`
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NetworkDevices) getPath() string {
	return "/dna/intent/api/v1/network-device"
}

// End of section. //template:end getPath

func (data NetworkDevices) toBody(ctx context.Context, state NetworkDevices) string {
	body := ""
	if len(data.Devices) > 0 {
		body, _ = sjson.Set(body, "response", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hostname", item.Hostname.ValueString())
			}
			if !item.ManagementIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "managementIpAddress", item.ManagementIpAddress.ValueString())
			}
			if !item.ManagementState.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "managementState", item.ManagementState.ValueString())
			}
			if !item.PlatformId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "platformId", item.PlatformId.ValueString())
			}
			if !item.Role.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "role", item.Role.ValueString())
			}
			if !item.SoftwareType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "softwareType", item.SoftwareType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "response.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *NetworkDevices) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]NetworkDevicesDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkDevicesDevices{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("hostname"); cValue.Exists() {
				item.Hostname = types.StringValue(cValue.String())
			} else {
				item.Hostname = types.StringNull()
			}
			if cValue := v.Get("managementIpAddress"); cValue.Exists() {
				item.ManagementIpAddress = types.StringValue(cValue.String())
			} else {
				item.ManagementIpAddress = types.StringNull()
			}
			if cValue := v.Get("managementState"); cValue.Exists() {
				item.ManagementState = types.StringValue(cValue.String())
			} else {
				item.ManagementState = types.StringNull()
			}
			if cValue := v.Get("platformId"); cValue.Exists() {
				item.PlatformId = types.StringValue(cValue.String())
			} else {
				item.PlatformId = types.StringNull()
			}
			if cValue := v.Get("role"); cValue.Exists() {
				item.Role = types.StringValue(cValue.String())
			} else {
				item.Role = types.StringNull()
			}
			if cValue := v.Get("softwareType"); cValue.Exists() {
				item.SoftwareType = types.StringValue(cValue.String())
			} else {
				item.SoftwareType = types.StringNull()
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *NetworkDevices) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Devices {
		keys := [...]string{"id", "hostname", "managementIpAddress", "managementState", "platformId", "role", "softwareType"}
		keyValues := [...]string{data.Devices[i].Id.ValueString(), data.Devices[i].Hostname.ValueString(), data.Devices[i].ManagementIpAddress.ValueString(), data.Devices[i].ManagementState.ValueString(), data.Devices[i].PlatformId.ValueString(), data.Devices[i].Role.ValueString(), data.Devices[i].SoftwareType.ValueString()}

		var r gjson.Result
		res.Get("response").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.Devices[i].Id.IsNull() {
			data.Devices[i].Id = types.StringValue(value.String())
		} else {
			data.Devices[i].Id = types.StringNull()
		}
		if value := r.Get("hostname"); value.Exists() && !data.Devices[i].Hostname.IsNull() {
			data.Devices[i].Hostname = types.StringValue(value.String())
		} else {
			data.Devices[i].Hostname = types.StringNull()
		}
		if value := r.Get("managementIpAddress"); value.Exists() && !data.Devices[i].ManagementIpAddress.IsNull() {
			data.Devices[i].ManagementIpAddress = types.StringValue(value.String())
		} else {
			data.Devices[i].ManagementIpAddress = types.StringNull()
		}
		if value := r.Get("managementState"); value.Exists() && !data.Devices[i].ManagementState.IsNull() {
			data.Devices[i].ManagementState = types.StringValue(value.String())
		} else {
			data.Devices[i].ManagementState = types.StringNull()
		}
		if value := r.Get("platformId"); value.Exists() && !data.Devices[i].PlatformId.IsNull() {
			data.Devices[i].PlatformId = types.StringValue(value.String())
		} else {
			data.Devices[i].PlatformId = types.StringNull()
		}
		if value := r.Get("role"); value.Exists() && !data.Devices[i].Role.IsNull() {
			data.Devices[i].Role = types.StringValue(value.String())
		} else {
			data.Devices[i].Role = types.StringNull()
		}
		if value := r.Get("softwareType"); value.Exists() && !data.Devices[i].SoftwareType.IsNull() {
			data.Devices[i].SoftwareType = types.StringValue(value.String())
		} else {
			data.Devices[i].SoftwareType = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *NetworkDevices) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Devices) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
