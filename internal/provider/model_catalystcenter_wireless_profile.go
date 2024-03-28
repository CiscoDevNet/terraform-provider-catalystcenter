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
type WirelessProfile struct {
	Id          types.String                 `tfsdk:"id"`
	Name        types.String                 `tfsdk:"name"`
	SsidDetails []WirelessProfileSsidDetails `tfsdk:"ssid_details"`
}

type WirelessProfileSsidDetails struct {
	Name              types.String `tfsdk:"name"`
	EnableFabric      types.Bool   `tfsdk:"enable_fabric"`
	EnableFlexConnect types.Bool   `tfsdk:"enable_flex_connect"`
	LocalToVlan       types.Int64  `tfsdk:"local_to_vlan"`
	InterfaceName     types.String `tfsdk:"interface_name"`
	WlanProfileName   types.String `tfsdk:"wlan_profile_name"`
	PolicyProfileName types.String `tfsdk:"policy_profile_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessProfile) getPath() string {
	return "/dna/intent/api/v1/wireless/profile"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data WirelessProfile) getPathDelete() string {
	return "/dna/intent/api/v1/wireless-profile"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessProfile) toBody(ctx context.Context, state WirelessProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "profileDetails.name", data.Name.ValueString())
	}
	if len(data.SsidDetails) > 0 {
		body, _ = sjson.Set(body, "profileDetails.ssidDetails", []interface{}{})
		for _, item := range data.SsidDetails {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.EnableFabric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableFabric", item.EnableFabric.ValueBool())
			}
			if !item.EnableFlexConnect.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "flexConnect.enableFlexConnect", item.EnableFlexConnect.ValueBool())
			}
			if !item.LocalToVlan.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "flexConnect.localToVlan", item.LocalToVlan.ValueInt64())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceName", item.InterfaceName.ValueString())
			}
			if !item.WlanProfileName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "wlanProfileName", item.WlanProfileName.ValueString())
			}
			if !item.PolicyProfileName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policyProfileName", item.PolicyProfileName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "profileDetails.ssidDetails.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.profileDetails.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("0.profileDetails.ssidDetails"); value.Exists() {
		data.SsidDetails = make([]WirelessProfileSsidDetails, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := WirelessProfileSsidDetails{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("enableFabric"); cValue.Exists() {
				item.EnableFabric = types.BoolValue(cValue.Bool())
			} else {
				item.EnableFabric = types.BoolNull()
			}
			if cValue := v.Get("flexConnect.enableFlexConnect"); cValue.Exists() {
				item.EnableFlexConnect = types.BoolValue(cValue.Bool())
			} else {
				item.EnableFlexConnect = types.BoolNull()
			}
			if cValue := v.Get("flexConnect.localToVlan"); cValue.Exists() {
				item.LocalToVlan = types.Int64Value(cValue.Int())
			} else {
				item.LocalToVlan = types.Int64Null()
			}
			if cValue := v.Get("interfaceName"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			} else {
				item.InterfaceName = types.StringNull()
			}
			data.SsidDetails = append(data.SsidDetails, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.profileDetails.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := range data.SsidDetails {
		keys := [...]string{"name"}
		keyValues := [...]string{data.SsidDetails[i].Name.ValueString()}

		var r gjson.Result
		res.Get("0.profileDetails.ssidDetails").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.SsidDetails[i].Name.IsNull() {
			data.SsidDetails[i].Name = types.StringValue(value.String())
		} else {
			data.SsidDetails[i].Name = types.StringNull()
		}
		if value := r.Get("enableFabric"); value.Exists() && !data.SsidDetails[i].EnableFabric.IsNull() {
			data.SsidDetails[i].EnableFabric = types.BoolValue(value.Bool())
		} else {
			data.SsidDetails[i].EnableFabric = types.BoolNull()
		}
		if value := r.Get("flexConnect.enableFlexConnect"); value.Exists() && !data.SsidDetails[i].EnableFlexConnect.IsNull() {
			data.SsidDetails[i].EnableFlexConnect = types.BoolValue(value.Bool())
		} else {
			data.SsidDetails[i].EnableFlexConnect = types.BoolNull()
		}
		if value := r.Get("flexConnect.localToVlan"); value.Exists() && !data.SsidDetails[i].LocalToVlan.IsNull() {
			data.SsidDetails[i].LocalToVlan = types.Int64Value(value.Int())
		} else {
			data.SsidDetails[i].LocalToVlan = types.Int64Null()
		}
		if value := r.Get("interfaceName"); value.Exists() && !data.SsidDetails[i].InterfaceName.IsNull() {
			data.SsidDetails[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.SsidDetails[i].InterfaceName = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if len(data.SsidDetails) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
