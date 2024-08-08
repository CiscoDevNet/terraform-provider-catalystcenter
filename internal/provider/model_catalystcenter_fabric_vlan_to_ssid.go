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
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type FabricVLANToSSID struct {
	Id       types.String               `tfsdk:"id"`
	FabricId types.String               `tfsdk:"fabric_id"`
	Mappings []FabricVLANToSSIDMappings `tfsdk:"mappings"`
}

type FabricVLANToSSIDMappings struct {
	VlanName    types.String                          `tfsdk:"vlan_name"`
	SsidDetails []FabricVLANToSSIDMappingsSsidDetails `tfsdk:"ssid_details"`
}

type FabricVLANToSSIDMappingsSsidDetails struct {
	Name             types.String `tfsdk:"name"`
	SecurityGroupTag types.String `tfsdk:"security_group_tag"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricVLANToSSID) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sda/fabrics/%v/vlanToSsids", url.QueryEscape(data.FabricId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricVLANToSSID) toBody(ctx context.Context, state FabricVLANToSSID) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.Mappings) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.Mappings {
			itemBody := ""
			if !item.VlanName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanName", item.VlanName.ValueString())
			}
			if len(item.SsidDetails) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ssidDetails", []interface{}{})
				for _, childItem := range item.SsidDetails {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.SecurityGroupTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "securityGroupTag", childItem.SecurityGroupTag.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ssidDetails.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricVLANToSSID) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Mappings = make([]FabricVLANToSSIDMappings, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricVLANToSSIDMappings{}
			if cValue := v.Get("vlanName"); cValue.Exists() {
				item.VlanName = types.StringValue(cValue.String())
			} else {
				item.VlanName = types.StringNull()
			}
			if cValue := v.Get("ssidDetails"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.SsidDetails = make([]FabricVLANToSSIDMappingsSsidDetails, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := FabricVLANToSSIDMappingsSsidDetails{}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("securityGroupTag"); ccValue.Exists() {
						cItem.SecurityGroupTag = types.StringValue(ccValue.String())
					} else {
						cItem.SecurityGroupTag = types.StringNull()
					}
					item.SsidDetails = append(item.SsidDetails, cItem)
					return true
				})
			}
			data.Mappings = append(data.Mappings, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricVLANToSSID) updateFromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.Mappings {
		keys := [...]string{"vlanName"}
		keyValues := [...]string{data.Mappings[i].VlanName.ValueString()}

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
		if value := r.Get("vlanName"); value.Exists() && !data.Mappings[i].VlanName.IsNull() {
			data.Mappings[i].VlanName = types.StringValue(value.String())
		} else {
			data.Mappings[i].VlanName = types.StringNull()
		}
		for ci := range data.Mappings[i].SsidDetails {
			keys := [...]string{"name", "securityGroupTag"}
			keyValues := [...]string{data.Mappings[i].SsidDetails[ci].Name.ValueString(), data.Mappings[i].SsidDetails[ci].SecurityGroupTag.ValueString()}

			var cr gjson.Result
			r.Get("ssidDetails").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("name"); value.Exists() && !data.Mappings[i].SsidDetails[ci].Name.IsNull() {
				data.Mappings[i].SsidDetails[ci].Name = types.StringValue(value.String())
			} else {
				data.Mappings[i].SsidDetails[ci].Name = types.StringNull()
			}
			if value := cr.Get("securityGroupTag"); value.Exists() && !data.Mappings[i].SsidDetails[ci].SecurityGroupTag.IsNull() {
				data.Mappings[i].SsidDetails[ci].SecurityGroupTag = types.StringValue(value.String())
			} else {
				data.Mappings[i].SsidDetails[ci].SecurityGroupTag = types.StringNull()
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricVLANToSSID) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Mappings) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
