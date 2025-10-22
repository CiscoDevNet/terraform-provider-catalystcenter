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
type ProvisionDevices struct {
	Id               types.String                       `tfsdk:"id"`
	SiteId           types.String                       `tfsdk:"site_id"`
	ProvisionDevices []ProvisionDevicesProvisionDevices `tfsdk:"provision_devices"`
}

type ProvisionDevicesProvisionDevices struct {
	Id              types.String `tfsdk:"id"`
	SiteId          types.String `tfsdk:"site_id"`
	NetworkDeviceId types.String `tfsdk:"network_device_id"`
	Reprovision     types.Bool   `tfsdk:"reprovision"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ProvisionDevices) getPath() string {
	return "/dna/intent/api/v1/sda/provisionDevices"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ProvisionDevices) toBody(ctx context.Context, state ProvisionDevices) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if len(data.ProvisionDevices) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.ProvisionDevices {
			itemBody := ""
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.SiteId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "siteId", item.SiteId.ValueString())
			}
			if !item.NetworkDeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkDeviceId", item.NetworkDeviceId.ValueString())
			}
			if !item.Reprovision.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "", item.Reprovision.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ProvisionDevices) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.ProvisionDevices = make([]ProvisionDevicesProvisionDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ProvisionDevicesProvisionDevices{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("siteId"); cValue.Exists() {
				item.SiteId = types.StringValue(cValue.String())
			} else {
				item.SiteId = types.StringNull()
			}
			if cValue := v.Get("networkDeviceId"); cValue.Exists() {
				item.NetworkDeviceId = types.StringValue(cValue.String())
			} else {
				item.NetworkDeviceId = types.StringNull()
			}
			if cValue := v.Get(""); cValue.Exists() {
				item.Reprovision = types.BoolValue(cValue.Bool())
			} else {
				item.Reprovision = types.BoolValue(false)
			}
			data.ProvisionDevices = append(data.ProvisionDevices, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

func (data *ProvisionDevices) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []ProvisionDevicesProvisionDevices

	res = res.Get("response")
	for i := range data.ProvisionDevices {
		keys := [...]string{"networkDeviceId"}
		keyValues := [...]string{data.ProvisionDevices[i].NetworkDeviceId.ValueString()}

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
		if value := r.Get("siteId"); value.Exists() && !data.ProvisionDevices[i].SiteId.IsNull() {
			data.ProvisionDevices[i].SiteId = types.StringValue(value.String())
		} else {
			data.ProvisionDevices[i].SiteId = types.StringNull()
		}
		if value := r.Get("networkDeviceId"); value.Exists() && !data.ProvisionDevices[i].NetworkDeviceId.IsNull() {
			data.ProvisionDevices[i].NetworkDeviceId = types.StringValue(value.String())
		} else {
			data.ProvisionDevices[i].NetworkDeviceId = types.StringNull()
		}
		if value := r.Get(""); value.Exists() && !data.ProvisionDevices[i].Reprovision.IsNull() {
			data.ProvisionDevices[i].Reprovision = types.BoolValue(value.Bool())
		} else {
			data.ProvisionDevices[i].Reprovision = types.BoolValue(false)
		}
		if data.ProvisionDevices[i].Id != types.StringNull() {
			final = append(final, data.ProvisionDevices[i])
		}
	}
	data.ProvisionDevices = final
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ProvisionDevices) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.ProvisionDevices {
		keys := [...]string{"networkDeviceId"}
		keyValues := [...]string{data.ProvisionDevices[i].NetworkDeviceId.ValueString()}

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
		if data.ProvisionDevices[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.ProvisionDevices[i].Id.IsNull() {
				data.ProvisionDevices[i].Id = types.StringValue(value.String())
			} else {
				data.ProvisionDevices[i].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ProvisionDevices) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.ProvisionDevices) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
