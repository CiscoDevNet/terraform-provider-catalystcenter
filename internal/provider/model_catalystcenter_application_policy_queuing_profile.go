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
type ApplicationPolicyQueuingProfile struct {
	Id          types.String                            `tfsdk:"id"`
	Name        types.String                            `tfsdk:"name"`
	Description types.String                            `tfsdk:"description"`
	Clause      []ApplicationPolicyQueuingProfileClause `tfsdk:"clause"`
}

type ApplicationPolicyQueuingProfileClause struct {
	Type                              types.String                                                          `tfsdk:"type"`
	IsCommonBetweenAllInterfaceSpeeds types.Bool                                                            `tfsdk:"is_common_between_all_interface_speeds"`
	InterfaceSpeedBandwidthClauses    []ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses `tfsdk:"interface_speed_bandwidth_clauses"`
	TcDscpSettings                    []ApplicationPolicyQueuingProfileClauseTcDscpSettings                 `tfsdk:"tc_dscp_settings"`
}

type ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses struct {
	InterfaceSpeed      types.String                                                                             `tfsdk:"interface_speed"`
	TcBandwidthSettings []ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `tfsdk:"tc_bandwidth_settings"`
}
type ApplicationPolicyQueuingProfileClauseTcDscpSettings struct {
	TrafficClass types.String `tfsdk:"traffic_class"`
	Dscp         types.String `tfsdk:"dscp"`
}

type ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	TrafficClass        types.String `tfsdk:"traffic_class"`
	BandwidthPercentage types.Int64  `tfsdk:"bandwidth_percentage"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ApplicationPolicyQueuingProfile) getPath() string {
	return "/dna/intent/api/v1/app-policy-queuing-profile"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ApplicationPolicyQueuingProfile) toBody(ctx context.Context, state ApplicationPolicyQueuingProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "0.name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "0.description", data.Description.ValueString())
	}
	if len(data.Clause) > 0 {
		body, _ = sjson.Set(body, "0.clause", []interface{}{})
		for _, item := range data.Clause {
			itemBody := ""
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.IsCommonBetweenAllInterfaceSpeeds.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isCommonBetweenAllInterfaceSpeeds", item.IsCommonBetweenAllInterfaceSpeeds.ValueBool())
			}
			if len(item.InterfaceSpeedBandwidthClauses) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaceSpeedBandwidthClauses", []interface{}{})
				for _, childItem := range item.InterfaceSpeedBandwidthClauses {
					itemChildBody := ""
					if !childItem.InterfaceSpeed.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "interfaceSpeed", childItem.InterfaceSpeed.ValueString())
					}
					if len(childItem.TcBandwidthSettings) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "tcBandwidthSettings", []interface{}{})
						for _, childChildItem := range childItem.TcBandwidthSettings {
							itemChildChildBody := ""
							if !childChildItem.TrafficClass.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "trafficClass", childChildItem.TrafficClass.ValueString())
							}
							if !childChildItem.BandwidthPercentage.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "bandwidthPercentage", childChildItem.BandwidthPercentage.ValueInt64())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "tcBandwidthSettings.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaceSpeedBandwidthClauses.-1", itemChildBody)
				}
			}
			if len(item.TcDscpSettings) > 0 {
				itemBody, _ = sjson.Set(itemBody, "tcDscpSettings", []interface{}{})
				for _, childItem := range item.TcDscpSettings {
					itemChildBody := ""
					if !childItem.TrafficClass.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "trafficClass", childItem.TrafficClass.ValueString())
					}
					if !childItem.Dscp.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "dscp", childItem.Dscp.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "tcDscpSettings.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "0.clause.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ApplicationPolicyQueuingProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("clause"); value.Exists() && len(value.Array()) > 0 {
		data.Clause = make([]ApplicationPolicyQueuingProfileClause, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationPolicyQueuingProfileClause{}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("isCommonBetweenAllInterfaceSpeeds"); cValue.Exists() {
				item.IsCommonBetweenAllInterfaceSpeeds = types.BoolValue(cValue.Bool())
			} else {
				item.IsCommonBetweenAllInterfaceSpeeds = types.BoolNull()
			}
			if cValue := v.Get("interfaceSpeedBandwidthClauses"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.InterfaceSpeedBandwidthClauses = make([]ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
					if ccValue := cv.Get("interfaceSpeed"); ccValue.Exists() {
						cItem.InterfaceSpeed = types.StringValue(ccValue.String())
					} else {
						cItem.InterfaceSpeed = types.StringNull()
					}
					if ccValue := cv.Get("tcBandwidthSettings"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.TcBandwidthSettings = make([]ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
							if cccValue := ccv.Get("trafficClass"); cccValue.Exists() {
								ccItem.TrafficClass = types.StringValue(cccValue.String())
							} else {
								ccItem.TrafficClass = types.StringNull()
							}
							if cccValue := ccv.Get("bandwidthPercentage"); cccValue.Exists() {
								ccItem.BandwidthPercentage = types.Int64Value(cccValue.Int())
							} else {
								ccItem.BandwidthPercentage = types.Int64Null()
							}
							cItem.TcBandwidthSettings = append(cItem.TcBandwidthSettings, ccItem)
							return true
						})
					}
					item.InterfaceSpeedBandwidthClauses = append(item.InterfaceSpeedBandwidthClauses, cItem)
					return true
				})
			}
			if cValue := v.Get("tcDscpSettings"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.TcDscpSettings = make([]ApplicationPolicyQueuingProfileClauseTcDscpSettings, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationPolicyQueuingProfileClauseTcDscpSettings{}
					if ccValue := cv.Get("trafficClass"); ccValue.Exists() {
						cItem.TrafficClass = types.StringValue(ccValue.String())
					} else {
						cItem.TrafficClass = types.StringNull()
					}
					if ccValue := cv.Get("dscp"); ccValue.Exists() {
						cItem.Dscp = types.StringValue(ccValue.String())
					} else {
						cItem.Dscp = types.StringNull()
					}
					item.TcDscpSettings = append(item.TcDscpSettings, cItem)
					return true
				})
			}
			data.Clause = append(data.Clause, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ApplicationPolicyQueuingProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.Clause {
		keys := [...]string{"type"}
		keyValues := [...]string{data.Clause[i].Type.ValueString()}

		var r gjson.Result
		res.Get("clause").ForEach(
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
		if value := r.Get("type"); value.Exists() && !data.Clause[i].Type.IsNull() {
			data.Clause[i].Type = types.StringValue(value.String())
		} else {
			data.Clause[i].Type = types.StringNull()
		}
		if value := r.Get("isCommonBetweenAllInterfaceSpeeds"); value.Exists() && !data.Clause[i].IsCommonBetweenAllInterfaceSpeeds.IsNull() {
			data.Clause[i].IsCommonBetweenAllInterfaceSpeeds = types.BoolValue(value.Bool())
		} else {
			data.Clause[i].IsCommonBetweenAllInterfaceSpeeds = types.BoolNull()
		}
		for ci := range data.Clause[i].InterfaceSpeedBandwidthClauses {
			keys := [...]string{"interfaceSpeed"}
			keyValues := [...]string{data.Clause[i].InterfaceSpeedBandwidthClauses[ci].InterfaceSpeed.ValueString()}

			var cr gjson.Result
			r.Get("interfaceSpeedBandwidthClauses").ForEach(
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
			if value := cr.Get("interfaceSpeed"); value.Exists() && !data.Clause[i].InterfaceSpeedBandwidthClauses[ci].InterfaceSpeed.IsNull() {
				data.Clause[i].InterfaceSpeedBandwidthClauses[ci].InterfaceSpeed = types.StringValue(value.String())
			} else {
				data.Clause[i].InterfaceSpeedBandwidthClauses[ci].InterfaceSpeed = types.StringNull()
			}
			for cci := range data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings {
				keys := [...]string{"trafficClass"}
				keyValues := [...]string{data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].TrafficClass.ValueString()}

				var ccr gjson.Result
				cr.Get("tcBandwidthSettings").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				if value := ccr.Get("trafficClass"); value.Exists() && !data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].TrafficClass.IsNull() {
					data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].TrafficClass = types.StringValue(value.String())
				} else {
					data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].TrafficClass = types.StringNull()
				}
				if value := ccr.Get("bandwidthPercentage"); value.Exists() && !data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].BandwidthPercentage.IsNull() {
					data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].BandwidthPercentage = types.Int64Value(value.Int())
				} else {
					data.Clause[i].InterfaceSpeedBandwidthClauses[ci].TcBandwidthSettings[cci].BandwidthPercentage = types.Int64Null()
				}
			}
		}
		for ci := range data.Clause[i].TcDscpSettings {
			keys := [...]string{"trafficClass"}
			keyValues := [...]string{data.Clause[i].TcDscpSettings[ci].TrafficClass.ValueString()}

			var cr gjson.Result
			r.Get("tcDscpSettings").ForEach(
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
			if value := cr.Get("trafficClass"); value.Exists() && !data.Clause[i].TcDscpSettings[ci].TrafficClass.IsNull() {
				data.Clause[i].TcDscpSettings[ci].TrafficClass = types.StringValue(value.String())
			} else {
				data.Clause[i].TcDscpSettings[ci].TrafficClass = types.StringNull()
			}
			if value := cr.Get("dscp"); value.Exists() && !data.Clause[i].TcDscpSettings[ci].Dscp.IsNull() {
				data.Clause[i].TcDscpSettings[ci].Dscp = types.StringValue(value.String())
			} else {
				data.Clause[i].TcDscpSettings[ci].Dscp = types.StringNull()
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ApplicationPolicyQueuingProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Description.IsNull() {
		return false
	}
	if len(data.Clause) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
