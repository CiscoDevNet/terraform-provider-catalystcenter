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
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type FabricDevices struct {
	Id            types.String                 `tfsdk:"id"`
	FabricId      types.String                 `tfsdk:"fabric_id"`
	FabricDevices []FabricDevicesFabricDevices `tfsdk:"fabric_devices"`
}

type FabricDevicesFabricDevices struct {
	Id                           types.String `tfsdk:"id"`
	NetworkDeviceId              types.String `tfsdk:"network_device_id"`
	FabricId                     types.String `tfsdk:"fabric_id"`
	DeviceRoles                  types.Set    `tfsdk:"device_roles"`
	BorderTypes                  types.Set    `tfsdk:"border_types"`
	LocalAutonomousSystemNumber  types.String `tfsdk:"local_autonomous_system_number"`
	DefaultExit                  types.Bool   `tfsdk:"default_exit"`
	ImportExternalRoutes         types.Bool   `tfsdk:"import_external_routes"`
	BorderPriority               types.Int64  `tfsdk:"border_priority"`
	PrependAutonomousSystemCount types.Int64  `tfsdk:"prepend_autonomous_system_count"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricDevices) getPath() string {
	return "/dna/intent/api/v1/sda/fabricDevices"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricDevices) toBody(ctx context.Context, state FabricDevices) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "", data.FabricId.ValueString())
	}
	if len(data.FabricDevices) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.FabricDevices {
			itemBody := ""
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.NetworkDeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkDeviceId", item.NetworkDeviceId.ValueString())
			}
			if !item.FabricId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fabricId", item.FabricId.ValueString())
			}
			if !item.DeviceRoles.IsNull() {
				var values []string
				item.DeviceRoles.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "deviceRoles", values)
			}
			if !item.BorderTypes.IsNull() {
				var values []string
				item.BorderTypes.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "borderDeviceSettings.borderTypes", values)
			}
			if !item.LocalAutonomousSystemNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderDeviceSettings.layer3Settings.localAutonomousSystemNumber", item.LocalAutonomousSystemNumber.ValueString())
			}
			if !item.DefaultExit.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderDeviceSettings.layer3Settings.isDefaultExit", item.DefaultExit.ValueBool())
			}
			if !item.ImportExternalRoutes.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderDeviceSettings.layer3Settings.importExternalRoutes", item.ImportExternalRoutes.ValueBool())
			}
			if !item.BorderPriority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderDeviceSettings.layer3Settings.borderPriority", item.BorderPriority.ValueInt64())
			}
			if !item.PrependAutonomousSystemCount.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderDeviceSettings.layer3Settings.prependAutonomousSystemCount", item.PrependAutonomousSystemCount.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricDevices) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	data.Id = types.StringValue(fmt.Sprint(data.FabricId.ValueString()))
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.FabricDevices = make([]FabricDevicesFabricDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricDevicesFabricDevices{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("networkDeviceId"); cValue.Exists() {
				item.NetworkDeviceId = types.StringValue(cValue.String())
			} else {
				item.NetworkDeviceId = types.StringNull()
			}
			if cValue := v.Get("fabricId"); cValue.Exists() {
				item.FabricId = types.StringValue(cValue.String())
			} else {
				item.FabricId = types.StringNull()
			}
			if cValue := v.Get("deviceRoles"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.DeviceRoles = helpers.GetStringSet(cValue.Array())
			} else {
				item.DeviceRoles = types.SetNull(types.StringType)
			}
			if cValue := v.Get("borderDeviceSettings.borderTypes"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.BorderTypes = helpers.GetStringSet(cValue.Array())
			} else {
				item.BorderTypes = types.SetNull(types.StringType)
			}
			if cValue := v.Get("borderDeviceSettings.layer3Settings.localAutonomousSystemNumber"); cValue.Exists() {
				item.LocalAutonomousSystemNumber = types.StringValue(cValue.String())
			} else {
				item.LocalAutonomousSystemNumber = types.StringNull()
			}
			if cValue := v.Get("borderDeviceSettings.layer3Settings.isDefaultExit"); cValue.Exists() {
				item.DefaultExit = types.BoolValue(cValue.Bool())
			} else {
				item.DefaultExit = types.BoolNull()
			}
			if cValue := v.Get("borderDeviceSettings.layer3Settings.importExternalRoutes"); cValue.Exists() {
				item.ImportExternalRoutes = types.BoolValue(cValue.Bool())
			} else {
				item.ImportExternalRoutes = types.BoolNull()
			}
			if cValue := v.Get("borderDeviceSettings.layer3Settings.borderPriority"); cValue.Exists() {
				item.BorderPriority = types.Int64Value(cValue.Int())
			} else {
				item.BorderPriority = types.Int64Null()
			}
			if cValue := v.Get("borderDeviceSettings.layer3Settings.prependAutonomousSystemCount"); cValue.Exists() {
				item.PrependAutonomousSystemCount = types.Int64Value(cValue.Int())
			} else {
				item.PrependAutonomousSystemCount = types.Int64Null()
			}
			data.FabricDevices = append(data.FabricDevices, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricDevices) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []FabricDevicesFabricDevices

	res = res.Get("response")
	for i := range data.FabricDevices {
		keys := [...]string{"networkDeviceId"}
		keyValues := [...]string{data.FabricDevices[i].NetworkDeviceId.ValueString()}

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
		if value := r.Get("id"); value.Exists() && !data.FabricDevices[i].Id.IsNull() {
			data.FabricDevices[i].Id = types.StringValue(value.String())
		} else {
			data.FabricDevices[i].Id = types.StringNull()
		}
		if value := r.Get("networkDeviceId"); value.Exists() && !data.FabricDevices[i].NetworkDeviceId.IsNull() {
			data.FabricDevices[i].NetworkDeviceId = types.StringValue(value.String())
		} else {
			data.FabricDevices[i].NetworkDeviceId = types.StringNull()
		}
		if value := r.Get("fabricId"); value.Exists() && !data.FabricDevices[i].FabricId.IsNull() {
			data.FabricDevices[i].FabricId = types.StringValue(value.String())
		} else {
			data.FabricDevices[i].FabricId = types.StringNull()
		}
		if value := r.Get("deviceRoles"); value.Exists() && !data.FabricDevices[i].DeviceRoles.IsNull() {
			data.FabricDevices[i].DeviceRoles = helpers.GetStringSet(value.Array())
		} else {
			data.FabricDevices[i].DeviceRoles = types.SetNull(types.StringType)
		}
		if value := r.Get("borderDeviceSettings.borderTypes"); value.Exists() && !data.FabricDevices[i].BorderTypes.IsNull() {
			data.FabricDevices[i].BorderTypes = helpers.GetStringSet(value.Array())
		} else {
			data.FabricDevices[i].BorderTypes = types.SetNull(types.StringType)
		}
		if value := r.Get("borderDeviceSettings.layer3Settings.localAutonomousSystemNumber"); value.Exists() && !data.FabricDevices[i].LocalAutonomousSystemNumber.IsNull() {
			data.FabricDevices[i].LocalAutonomousSystemNumber = types.StringValue(value.String())
		} else {
			data.FabricDevices[i].LocalAutonomousSystemNumber = types.StringNull()
		}
		if value := r.Get("borderDeviceSettings.layer3Settings.isDefaultExit"); value.Exists() && !data.FabricDevices[i].DefaultExit.IsNull() {
			data.FabricDevices[i].DefaultExit = types.BoolValue(value.Bool())
		} else {
			data.FabricDevices[i].DefaultExit = types.BoolNull()
		}
		if value := r.Get("borderDeviceSettings.layer3Settings.importExternalRoutes"); value.Exists() && !data.FabricDevices[i].ImportExternalRoutes.IsNull() {
			data.FabricDevices[i].ImportExternalRoutes = types.BoolValue(value.Bool())
		} else {
			data.FabricDevices[i].ImportExternalRoutes = types.BoolNull()
		}
		if value := r.Get("borderDeviceSettings.layer3Settings.borderPriority"); value.Exists() && !data.FabricDevices[i].BorderPriority.IsNull() {
			data.FabricDevices[i].BorderPriority = types.Int64Value(value.Int())
		} else {
			data.FabricDevices[i].BorderPriority = types.Int64Null()
		}
		if value := r.Get("borderDeviceSettings.layer3Settings.prependAutonomousSystemCount"); value.Exists() && !data.FabricDevices[i].PrependAutonomousSystemCount.IsNull() {
			data.FabricDevices[i].PrependAutonomousSystemCount = types.Int64Value(value.Int())
		} else {
			data.FabricDevices[i].PrependAutonomousSystemCount = types.Int64Null()
		}
		if data.FabricDevices[i].Id != types.StringNull() {
			final = append(final, data.FabricDevices[i])
		}
	}
	data.FabricDevices = final
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FabricDevices) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.FabricDevices {
		keys := [...]string{"networkDeviceId"}
		keyValues := [...]string{data.FabricDevices[i].NetworkDeviceId.ValueString()}

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
		if data.FabricDevices[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.FabricDevices[i].Id.IsNull() {
				data.FabricDevices[i].Id = types.StringValue(value.String())
			} else {
				data.FabricDevices[i].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricDevices) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.FabricDevices) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
