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
type ProvisionAccessPoints struct {
	Id             types.String                          `tfsdk:"id"`
	NetworkDevices []ProvisionAccessPointsNetworkDevices `tfsdk:"network_devices"`
	RfProfileName  types.String                          `tfsdk:"rf_profile_name"`
	ApZoneName     types.String                          `tfsdk:"ap_zone_name"`
	SiteId         types.String                          `tfsdk:"site_id"`
}

type ProvisionAccessPointsNetworkDevices struct {
	Id          types.String `tfsdk:"id"`
	DeviceId    types.String `tfsdk:"device_id"`
	MeshRole    types.String `tfsdk:"mesh_role"`
	Reprovision types.Bool   `tfsdk:"reprovision"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ProvisionAccessPoints) getPath() string {
	return "/dna/intent/api/v1/wirelessAccessPoints/provision"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ProvisionAccessPoints) toBody(ctx context.Context, state ProvisionAccessPoints) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.NetworkDevices) > 0 {
		body, _ = sjson.Set(body, "networkDevices", []interface{}{})
		for _, item := range data.NetworkDevices {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.DeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceId", item.DeviceId.ValueString())
			}
			if !item.MeshRole.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "meshRole", item.MeshRole.ValueString())
			}
			if !item.Reprovision.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "", item.Reprovision.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "networkDevices.-1", itemBody)
		}
	}
	if !data.RfProfileName.IsNull() {
		body, _ = sjson.Set(body, "rfProfileName", data.RfProfileName.ValueString())
	}
	if !data.ApZoneName.IsNull() {
		body, _ = sjson.Set(body, "apZoneName", data.ApZoneName.ValueString())
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ProvisionAccessPoints) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("networkDevices"); value.Exists() && len(value.Array()) > 0 {
		data.NetworkDevices = make([]ProvisionAccessPointsNetworkDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ProvisionAccessPointsNetworkDevices{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("deviceId"); cValue.Exists() {
				item.DeviceId = types.StringValue(cValue.String())
			} else {
				item.DeviceId = types.StringNull()
			}
			if cValue := v.Get("meshRole"); cValue.Exists() {
				item.MeshRole = types.StringValue(cValue.String())
			} else {
				item.MeshRole = types.StringNull()
			}
			if cValue := v.Get(""); cValue.Exists() {
				item.Reprovision = types.BoolValue(cValue.Bool())
			} else {
				item.Reprovision = types.BoolNull()
			}
			data.NetworkDevices = append(data.NetworkDevices, item)
			return true
		})
	}
	if value := res.Get("rfProfileName"); value.Exists() {
		data.RfProfileName = types.StringValue(value.String())
	} else {
		data.RfProfileName = types.StringNull()
	}
	if value := res.Get("apZoneName"); value.Exists() {
		data.ApZoneName = types.StringValue(value.String())
	} else {
		data.ApZoneName = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ProvisionAccessPoints) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.NetworkDevices {
		keys := [...]string{"id", "deviceId", "meshRole", ""}
		keyValues := [...]string{data.NetworkDevices[i].Id.ValueString(), data.NetworkDevices[i].DeviceId.ValueString(), data.NetworkDevices[i].MeshRole.ValueString(), strconv.FormatBool(data.NetworkDevices[i].Reprovision.ValueBool())}

		var r gjson.Result
		res.Get("networkDevices").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.NetworkDevices[i].Id.IsNull() {
			data.NetworkDevices[i].Id = types.StringValue(value.String())
		} else {
			data.NetworkDevices[i].Id = types.StringNull()
		}
		if value := r.Get("deviceId"); value.Exists() && !data.NetworkDevices[i].DeviceId.IsNull() {
			data.NetworkDevices[i].DeviceId = types.StringValue(value.String())
		} else {
			data.NetworkDevices[i].DeviceId = types.StringNull()
		}
		if value := r.Get("meshRole"); value.Exists() && !data.NetworkDevices[i].MeshRole.IsNull() {
			data.NetworkDevices[i].MeshRole = types.StringValue(value.String())
		} else {
			data.NetworkDevices[i].MeshRole = types.StringNull()
		}
		if value := r.Get(""); value.Exists() && !data.NetworkDevices[i].Reprovision.IsNull() {
			data.NetworkDevices[i].Reprovision = types.BoolValue(value.Bool())
		} else {
			data.NetworkDevices[i].Reprovision = types.BoolNull()
		}
	}
	if value := res.Get("rfProfileName"); value.Exists() && !data.RfProfileName.IsNull() {
		data.RfProfileName = types.StringValue(value.String())
	} else {
		data.RfProfileName = types.StringNull()
	}
	if value := res.Get("apZoneName"); value.Exists() && !data.ApZoneName.IsNull() {
		data.ApZoneName = types.StringValue(value.String())
	} else {
		data.ApZoneName = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ProvisionAccessPoints) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.NetworkDevices) > 0 {
		return false
	}
	if !data.RfProfileName.IsNull() {
		return false
	}
	if !data.ApZoneName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
