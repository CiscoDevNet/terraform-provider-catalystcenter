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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Application struct {
	Id                        types.String                           `tfsdk:"id"`
	Name                      types.String                           `tfsdk:"name"`
	NetworkApplications       []ApplicationNetworkApplications       `tfsdk:"network_applications"`
	ApplicationSetId          types.String                           `tfsdk:"application_set_id"`
	NetworkIdentity           []ApplicationNetworkIdentity           `tfsdk:"network_identity"`
	IndicativeNetworkIdentity []ApplicationIndicativeNetworkIdentity `tfsdk:"indicative_network_identity"`
}

type ApplicationNetworkApplications struct {
	HelpString         types.String `tfsdk:"help_string"`
	ApplicationType    types.String `tfsdk:"application_type"`
	Type               types.String `tfsdk:"type"`
	Dscp               types.String `tfsdk:"dscp"`
	AppProtocol        types.String `tfsdk:"app_protocol"`
	ServerName         types.String `tfsdk:"server_name"`
	Url                types.String `tfsdk:"url"`
	TrafficClass       types.String `tfsdk:"traffic_class"`
	CategoryId         types.String `tfsdk:"category_id"`
	IgnoreConflict     types.Bool   `tfsdk:"ignore_conflict"`
	Rank               types.Int64  `tfsdk:"rank"`
	EngineId           types.Int64  `tfsdk:"engine_id"`
	Id                 types.String `tfsdk:"id"`
	ApplicationSubType types.String `tfsdk:"application_sub_type"`
	DisplayName        types.String `tfsdk:"display_name"`
	Name               types.String `tfsdk:"name"`
}

type ApplicationNetworkIdentity struct {
	Protocol    types.String `tfsdk:"protocol"`
	Ports       types.String `tfsdk:"ports"`
	LowerPort   types.Int64  `tfsdk:"lower_port"`
	UpperPort   types.Int64  `tfsdk:"upper_port"`
	Ipv4Subnet  types.Set    `tfsdk:"ipv4_subnet"`
	Id          types.String `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
}

type ApplicationIndicativeNetworkIdentity struct {
	Protocol    types.String `tfsdk:"protocol"`
	Ports       types.String `tfsdk:"ports"`
	Ipv4Subnet  types.Set    `tfsdk:"ipv4_subnet"`
	Ipv6Subnet  types.Set    `tfsdk:"ipv6_subnet"`
	LowerPort   types.Int64  `tfsdk:"lower_port"`
	UpperPort   types.Int64  `tfsdk:"upper_port"`
	Id          types.String `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Application) getPath() string {
	return "/dna/intent/api/v2/applications"
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
func (data Application) toBody(ctx context.Context, state Application) string {
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
	if len(data.NetworkApplications) > 0 {
		body, _ = sjson.Set(body, "0.networkApplications", []interface{}{})
		for _, item := range data.NetworkApplications {
			itemBody := ""
			if !item.HelpString.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "helpString", item.HelpString.ValueString())
			}
			if !item.ApplicationType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "applicationType", item.ApplicationType.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Dscp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscp", item.Dscp.ValueString())
			}
			if !item.AppProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "appProtocol", item.AppProtocol.ValueString())
			}
			if !item.ServerName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serverName", item.ServerName.ValueString())
			}
			if !item.Url.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "url", item.Url.ValueString())
			}
			if !item.TrafficClass.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "trafficClass", item.TrafficClass.ValueString())
			}
			if !item.CategoryId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "categoryId", item.CategoryId.ValueString())
			}
			if !item.IgnoreConflict.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ignoreConflict", item.IgnoreConflict.ValueBool())
			}
			if !item.Rank.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rank", item.Rank.ValueInt64())
			}
			if !item.EngineId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "engineId", item.EngineId.ValueInt64())
			}
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if item.ApplicationSubType.ValueString() != "" && !item.ApplicationSubType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "applicationSubType", item.ApplicationSubType.ValueString())
			}
			if item.DisplayName.ValueString() != "" && !item.DisplayName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "displayName", item.DisplayName.ValueString())
			}
			if item.Name.ValueString() != "" && !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "0.networkApplications.-1", itemBody)
		}
	}
	if !data.ApplicationSetId.IsNull() {
		body, _ = sjson.Set(body, "0.parentScalableGroup.idRef", data.ApplicationSetId.ValueString())
	}
	if len(data.NetworkIdentity) > 0 {
		body, _ = sjson.Set(body, "0.networkIdentity", []interface{}{})
		for _, item := range data.NetworkIdentity {
			itemBody := ""
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.Ports.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ports", item.Ports.ValueString())
			}
			if !item.LowerPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lowerPort", item.LowerPort.ValueInt64())
			}
			if !item.UpperPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "upperPort", item.UpperPort.ValueInt64())
			}
			if !item.Ipv4Subnet.IsNull() {
				var values []string
				item.Ipv4Subnet.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipv4Subnet", values)
			}
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if item.DisplayName.ValueString() != "" && !item.DisplayName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "displayName", item.DisplayName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "0.networkIdentity.-1", itemBody)
		}
	}
	if len(data.IndicativeNetworkIdentity) > 0 {
		body, _ = sjson.Set(body, "0.indicativeNetworkIdentity", []interface{}{})
		for _, item := range data.IndicativeNetworkIdentity {
			itemBody := ""
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.Ports.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ports", item.Ports.ValueString())
			}
			if !item.Ipv4Subnet.IsNull() {
				var values []string
				item.Ipv4Subnet.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipv4Subnet", values)
			}
			if !item.Ipv6Subnet.IsNull() {
				var values []string
				item.Ipv6Subnet.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipv6Subnet", values)
			}
			if !item.LowerPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lowerPort", item.LowerPort.ValueInt64())
			}
			if !item.UpperPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "upperPort", item.UpperPort.ValueInt64())
			}
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if item.DisplayName.ValueString() != "" && !item.DisplayName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "displayName", item.DisplayName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "0.indicativeNetworkIdentity.-1", itemBody)
		}
	}
	body, _ = sjson.Set(body, "0.scalableGroupType", "APPLICATION")
	body, _ = sjson.Set(body, "0.type", "scalablegroup")
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Application) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.[id,0.id].0"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications"); value.Exists() && len(value.Array()) > 0 {
		data.NetworkApplications = make([]ApplicationNetworkApplications, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationNetworkApplications{}
			item.HelpString = types.StringValue("")
			item.ApplicationType = types.StringValue("")
			item.Type = types.StringValue("")
			if cValue := v.Get("dscp"); cValue.Exists() {
				item.Dscp = types.StringValue(cValue.String())
			} else {
				item.Dscp = types.StringNull()
			}
			if cValue := v.Get("serverName"); cValue.Exists() {
				item.ServerName = types.StringValue(cValue.String())
			} else {
				item.ServerName = types.StringNull()
			}
			if cValue := v.Get("url"); cValue.Exists() {
				item.Url = types.StringValue(cValue.String())
			} else {
				item.Url = types.StringNull()
			}
			if cValue := v.Get("trafficClass"); cValue.Exists() {
				item.TrafficClass = types.StringValue(cValue.String())
			} else {
				item.TrafficClass = types.StringNull()
			}
			if cValue := v.Get("categoryId"); cValue.Exists() {
				item.CategoryId = types.StringValue(cValue.String())
			} else {
				item.CategoryId = types.StringNull()
			}
			item.IgnoreConflict = types.BoolValue(false)
			if cValue := v.Get("rank"); cValue.Exists() {
				item.Rank = types.Int64Value(cValue.Int())
			} else {
				item.Rank = types.Int64Null()
			}
			if cValue := v.Get("engineId"); cValue.Exists() {
				item.EngineId = types.Int64Value(cValue.Int())
			} else {
				item.EngineId = types.Int64Null()
			}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("applicationSubType"); cValue.Exists() {
				item.ApplicationSubType = types.StringValue(cValue.String())
			} else {
				item.ApplicationSubType = types.StringNull()
			}
			if cValue := v.Get("displayName"); cValue.Exists() {
				item.DisplayName = types.StringValue(cValue.String())
			} else {
				item.DisplayName = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			data.NetworkApplications = append(data.NetworkApplications, item)
			return true
		})
	}
	if value := res.Get("response.0.parentScalableGroup.idRef"); value.Exists() {
		data.ApplicationSetId = types.StringValue(value.String())
	} else {
		data.ApplicationSetId = types.StringNull()
	}
	if value := res.Get("response.0.networkIdentity"); value.Exists() && len(value.Array()) > 0 {
		data.NetworkIdentity = make([]ApplicationNetworkIdentity, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationNetworkIdentity{}
			if cValue := v.Get("protocol"); cValue.Exists() {
				item.Protocol = types.StringValue(cValue.String())
			} else {
				item.Protocol = types.StringNull()
			}
			if cValue := v.Get("ports"); cValue.Exists() {
				item.Ports = types.StringValue(cValue.String())
			} else {
				item.Ports = types.StringNull()
			}
			if cValue := v.Get("lowerPort"); cValue.Exists() {
				item.LowerPort = types.Int64Value(cValue.Int())
			} else {
				item.LowerPort = types.Int64Null()
			}
			if cValue := v.Get("upperPort"); cValue.Exists() {
				item.UpperPort = types.Int64Value(cValue.Int())
			} else {
				item.UpperPort = types.Int64Null()
			}
			if cValue := v.Get("ipv4Subnet"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv4Subnet = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ipv4Subnet = types.SetNull(types.StringType)
			}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("displayName"); cValue.Exists() {
				item.DisplayName = types.StringValue(cValue.String())
			} else {
				item.DisplayName = types.StringNull()
			}
			data.NetworkIdentity = append(data.NetworkIdentity, item)
			return true
		})
	}
	if value := res.Get("response.0.indicativeNetworkIdentity"); value.Exists() && len(value.Array()) > 0 {
		data.IndicativeNetworkIdentity = make([]ApplicationIndicativeNetworkIdentity, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationIndicativeNetworkIdentity{}
			if cValue := v.Get("protocol"); cValue.Exists() {
				item.Protocol = types.StringValue(cValue.String())
			} else {
				item.Protocol = types.StringNull()
			}
			if cValue := v.Get("ports"); cValue.Exists() {
				item.Ports = types.StringValue(cValue.String())
			} else {
				item.Ports = types.StringNull()
			}
			if cValue := v.Get("ipv4Subnet"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv4Subnet = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ipv4Subnet = types.SetNull(types.StringType)
			}
			if cValue := v.Get("ipv6Subnet"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv6Subnet = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ipv6Subnet = types.SetNull(types.StringType)
			}
			if cValue := v.Get("lowerPort"); cValue.Exists() {
				item.LowerPort = types.Int64Value(cValue.Int())
			} else {
				item.LowerPort = types.Int64Null()
			}
			if cValue := v.Get("upperPort"); cValue.Exists() {
				item.UpperPort = types.Int64Value(cValue.Int())
			} else {
				item.UpperPort = types.Int64Null()
			}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("displayName"); cValue.Exists() {
				item.DisplayName = types.StringValue(cValue.String())
			} else {
				item.DisplayName = types.StringNull()
			}
			data.IndicativeNetworkIdentity = append(data.IndicativeNetworkIdentity, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Application) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := range data.NetworkApplications {
		values := res.Get("response.0.networkApplications").Array()
		if len(values) != 1 {
			data.NetworkApplications = nil
			break
		}
		r := values[0]
		if value := r.Get("dscp"); value.Exists() && !data.NetworkApplications[i].Dscp.IsNull() {
			data.NetworkApplications[i].Dscp = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].Dscp = types.StringNull()
		}
		if value := r.Get("serverName"); value.Exists() && !data.NetworkApplications[i].ServerName.IsNull() {
			data.NetworkApplications[i].ServerName = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].ServerName = types.StringNull()
		}
		if value := r.Get("url"); value.Exists() && !data.NetworkApplications[i].Url.IsNull() {
			data.NetworkApplications[i].Url = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].Url = types.StringNull()
		}
		if value := r.Get("trafficClass"); value.Exists() && !data.NetworkApplications[i].TrafficClass.IsNull() {
			data.NetworkApplications[i].TrafficClass = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].TrafficClass = types.StringNull()
		}
		if value := r.Get("categoryId"); value.Exists() && !data.NetworkApplications[i].CategoryId.IsNull() {
			data.NetworkApplications[i].CategoryId = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].CategoryId = types.StringNull()
		}
		if value := r.Get("rank"); value.Exists() && !data.NetworkApplications[i].Rank.IsNull() {
			data.NetworkApplications[i].Rank = types.Int64Value(value.Int())
		} else {
			data.NetworkApplications[i].Rank = types.Int64Null()
		}
		if value := r.Get("engineId"); value.Exists() && !data.NetworkApplications[i].EngineId.IsNull() {
			data.NetworkApplications[i].EngineId = types.Int64Value(value.Int())
		} else {
			data.NetworkApplications[i].EngineId = types.Int64Null()
		}
		if value := r.Get("id"); value.Exists() && !data.NetworkApplications[i].Id.IsNull() {
			data.NetworkApplications[i].Id = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].Id = types.StringNull()
		}
		if value := r.Get("applicationSubType"); value.Exists() && !data.NetworkApplications[i].ApplicationSubType.IsNull() {
			data.NetworkApplications[i].ApplicationSubType = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].ApplicationSubType = types.StringNull()
		}
		if value := r.Get("displayName"); value.Exists() && !data.NetworkApplications[i].DisplayName.IsNull() {
			data.NetworkApplications[i].DisplayName = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].DisplayName = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.NetworkApplications[i].Name.IsNull() {
			data.NetworkApplications[i].Name = types.StringValue(value.String())
		} else {
			data.NetworkApplications[i].Name = types.StringNull()
		}
	}
	if value := res.Get("response.0.parentScalableGroup.idRef"); value.Exists() && !data.ApplicationSetId.IsNull() {
		data.ApplicationSetId = types.StringValue(value.String())
	} else {
		data.ApplicationSetId = types.StringNull()
	}
	for i := range data.NetworkIdentity {
		keys := [...]string{"protocol", "ports", "lowerPort", "upperPort"}
		keyValues := [...]string{data.NetworkIdentity[i].Protocol.ValueString(), data.NetworkIdentity[i].Ports.ValueString(), strconv.FormatInt(data.NetworkIdentity[i].LowerPort.ValueInt64(), 10), strconv.FormatInt(data.NetworkIdentity[i].UpperPort.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("response.0.networkIdentity").ForEach(
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
		if value := r.Get("protocol"); value.Exists() && !data.NetworkIdentity[i].Protocol.IsNull() {
			data.NetworkIdentity[i].Protocol = types.StringValue(value.String())
		} else {
			data.NetworkIdentity[i].Protocol = types.StringNull()
		}
		if value := r.Get("ports"); value.Exists() && !data.NetworkIdentity[i].Ports.IsNull() {
			data.NetworkIdentity[i].Ports = types.StringValue(value.String())
		} else {
			data.NetworkIdentity[i].Ports = types.StringNull()
		}
		if value := r.Get("lowerPort"); value.Exists() && !data.NetworkIdentity[i].LowerPort.IsNull() {
			data.NetworkIdentity[i].LowerPort = types.Int64Value(value.Int())
		} else {
			data.NetworkIdentity[i].LowerPort = types.Int64Null()
		}
		if value := r.Get("upperPort"); value.Exists() && !data.NetworkIdentity[i].UpperPort.IsNull() {
			data.NetworkIdentity[i].UpperPort = types.Int64Value(value.Int())
		} else {
			data.NetworkIdentity[i].UpperPort = types.Int64Null()
		}
		if value := r.Get("ipv4Subnet"); value.Exists() && !data.NetworkIdentity[i].Ipv4Subnet.IsNull() {
			data.NetworkIdentity[i].Ipv4Subnet = helpers.GetStringSet(value.Array())
		} else {
			data.NetworkIdentity[i].Ipv4Subnet = types.SetNull(types.StringType)
		}
		if value := r.Get("id"); value.Exists() && !data.NetworkIdentity[i].Id.IsNull() {
			data.NetworkIdentity[i].Id = types.StringValue(value.String())
		} else {
			data.NetworkIdentity[i].Id = types.StringNull()
		}
		if value := r.Get("displayName"); value.Exists() && !data.NetworkIdentity[i].DisplayName.IsNull() {
			data.NetworkIdentity[i].DisplayName = types.StringValue(value.String())
		} else {
			data.NetworkIdentity[i].DisplayName = types.StringNull()
		}
	}
	for i := range data.IndicativeNetworkIdentity {
		keys := [...]string{"protocol", "ports", "lowerPort", "upperPort"}
		keyValues := [...]string{data.IndicativeNetworkIdentity[i].Protocol.ValueString(), data.IndicativeNetworkIdentity[i].Ports.ValueString(), strconv.FormatInt(data.IndicativeNetworkIdentity[i].LowerPort.ValueInt64(), 10), strconv.FormatInt(data.IndicativeNetworkIdentity[i].UpperPort.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("response.0.indicativeNetworkIdentity").ForEach(
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
		if value := r.Get("protocol"); value.Exists() && !data.IndicativeNetworkIdentity[i].Protocol.IsNull() {
			data.IndicativeNetworkIdentity[i].Protocol = types.StringValue(value.String())
		} else {
			data.IndicativeNetworkIdentity[i].Protocol = types.StringNull()
		}
		if value := r.Get("ports"); value.Exists() && !data.IndicativeNetworkIdentity[i].Ports.IsNull() {
			data.IndicativeNetworkIdentity[i].Ports = types.StringValue(value.String())
		} else {
			data.IndicativeNetworkIdentity[i].Ports = types.StringNull()
		}
		if value := r.Get("ipv4Subnet"); value.Exists() && !data.IndicativeNetworkIdentity[i].Ipv4Subnet.IsNull() {
			data.IndicativeNetworkIdentity[i].Ipv4Subnet = helpers.GetStringSet(value.Array())
		} else {
			data.IndicativeNetworkIdentity[i].Ipv4Subnet = types.SetNull(types.StringType)
		}
		if value := r.Get("ipv6Subnet"); value.Exists() && !data.IndicativeNetworkIdentity[i].Ipv6Subnet.IsNull() {
			data.IndicativeNetworkIdentity[i].Ipv6Subnet = helpers.GetStringSet(value.Array())
		} else {
			data.IndicativeNetworkIdentity[i].Ipv6Subnet = types.SetNull(types.StringType)
		}
		if value := r.Get("lowerPort"); value.Exists() && !data.IndicativeNetworkIdentity[i].LowerPort.IsNull() {
			data.IndicativeNetworkIdentity[i].LowerPort = types.Int64Value(value.Int())
		} else {
			data.IndicativeNetworkIdentity[i].LowerPort = types.Int64Null()
		}
		if value := r.Get("upperPort"); value.Exists() && !data.IndicativeNetworkIdentity[i].UpperPort.IsNull() {
			data.IndicativeNetworkIdentity[i].UpperPort = types.Int64Value(value.Int())
		} else {
			data.IndicativeNetworkIdentity[i].UpperPort = types.Int64Null()
		}
		if value := r.Get("id"); value.Exists() && !data.IndicativeNetworkIdentity[i].Id.IsNull() {
			data.IndicativeNetworkIdentity[i].Id = types.StringValue(value.String())
		} else {
			data.IndicativeNetworkIdentity[i].Id = types.StringNull()
		}
		if value := r.Get("displayName"); value.Exists() && !data.IndicativeNetworkIdentity[i].DisplayName.IsNull() {
			data.IndicativeNetworkIdentity[i].DisplayName = types.StringValue(value.String())
		} else {
			data.IndicativeNetworkIdentity[i].DisplayName = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Application) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Name.IsUnknown() {
		if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
	for i := range data.NetworkApplications {
		values := res.Get("response.0.networkApplications").Array()
		if len(values) != 1 {
			data.NetworkApplications = nil
			break
		}
		r := values[0]
		if data.NetworkApplications[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.NetworkApplications[i].Id.IsNull() {
				data.NetworkApplications[i].Id = types.StringValue(value.String())
			} else {
				data.NetworkApplications[i].Id = types.StringNull()
				continue
			}
		}
		if data.NetworkApplications[i].ApplicationSubType.IsUnknown() {
			if value := r.Get("applicationSubType"); value.Exists() && !data.NetworkApplications[i].ApplicationSubType.IsNull() {
				data.NetworkApplications[i].ApplicationSubType = types.StringValue(value.String())
			} else {
				data.NetworkApplications[i].ApplicationSubType = types.StringNull()
			}
		}
		if data.NetworkApplications[i].DisplayName.IsUnknown() {
			if value := r.Get("displayName"); value.Exists() && !data.NetworkApplications[i].DisplayName.IsNull() {
				data.NetworkApplications[i].DisplayName = types.StringValue(value.String())
			} else {
				data.NetworkApplications[i].DisplayName = types.StringNull()
			}
		}
		if data.NetworkApplications[i].Name.IsUnknown() {
			if value := r.Get("name"); value.Exists() && !data.NetworkApplications[i].Name.IsNull() {
				data.NetworkApplications[i].Name = types.StringValue(value.String())
			} else {
				data.NetworkApplications[i].Name = types.StringNull()
			}
		}
	}
	if data.ApplicationSetId.IsUnknown() {
		if value := res.Get("response.0.parentScalableGroup.idRef"); value.Exists() && !data.ApplicationSetId.IsNull() {
			data.ApplicationSetId = types.StringValue(value.String())
		} else {
			data.ApplicationSetId = types.StringNull()
		}
	}
	for i := range data.NetworkIdentity {
		keys := [...]string{"protocol", "ports", "lowerPort", "upperPort"}
		keyValues := [...]string{data.NetworkIdentity[i].Protocol.ValueString(), data.NetworkIdentity[i].Ports.ValueString(), strconv.FormatInt(data.NetworkIdentity[i].LowerPort.ValueInt64(), 10), strconv.FormatInt(data.NetworkIdentity[i].UpperPort.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("response.0.networkIdentity").ForEach(
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
		if data.NetworkIdentity[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.NetworkIdentity[i].Id.IsNull() {
				data.NetworkIdentity[i].Id = types.StringValue(value.String())
			} else {
				data.NetworkIdentity[i].Id = types.StringNull()
				continue
			}
		}
		if data.NetworkIdentity[i].DisplayName.IsUnknown() {
			if value := r.Get("displayName"); value.Exists() && !data.NetworkIdentity[i].DisplayName.IsNull() {
				data.NetworkIdentity[i].DisplayName = types.StringValue(value.String())
			} else {
				data.NetworkIdentity[i].DisplayName = types.StringNull()
			}
		}
	}
	for i := range data.IndicativeNetworkIdentity {
		keys := [...]string{"protocol", "ports", "lowerPort", "upperPort"}
		keyValues := [...]string{data.IndicativeNetworkIdentity[i].Protocol.ValueString(), data.IndicativeNetworkIdentity[i].Ports.ValueString(), strconv.FormatInt(data.IndicativeNetworkIdentity[i].LowerPort.ValueInt64(), 10), strconv.FormatInt(data.IndicativeNetworkIdentity[i].UpperPort.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("response.0.indicativeNetworkIdentity").ForEach(
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
		if data.IndicativeNetworkIdentity[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.IndicativeNetworkIdentity[i].Id.IsNull() {
				data.IndicativeNetworkIdentity[i].Id = types.StringValue(value.String())
			} else {
				data.IndicativeNetworkIdentity[i].Id = types.StringNull()
				continue
			}
		}
		if data.IndicativeNetworkIdentity[i].DisplayName.IsUnknown() {
			if value := r.Get("displayName"); value.Exists() && !data.IndicativeNetworkIdentity[i].DisplayName.IsNull() {
				data.IndicativeNetworkIdentity[i].DisplayName = types.StringValue(value.String())
			} else {
				data.IndicativeNetworkIdentity[i].DisplayName = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Application) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.NetworkApplications) > 0 {
		return false
	}
	if !data.ApplicationSetId.IsNull() {
		return false
	}
	if len(data.NetworkIdentity) > 0 {
		return false
	}
	if len(data.IndicativeNetworkIdentity) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
