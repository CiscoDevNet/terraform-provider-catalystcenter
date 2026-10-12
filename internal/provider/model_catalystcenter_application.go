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
type Application struct {
	Id                            types.String                 `tfsdk:"id"`
	Name                          types.String                 `tfsdk:"name"`
	ApplicationSetId              types.String                 `tfsdk:"application_set_id"`
	CategoryId                    types.String                 `tfsdk:"category_id"`
	TrafficClass                  types.String                 `tfsdk:"traffic_class"`
	HelpString                    types.String                 `tfsdk:"help_string"`
	Dscp                          types.String                 `tfsdk:"dscp"`
	Rank                          types.Int64                  `tfsdk:"rank"`
	AppProtocol                   types.String                 `tfsdk:"app_protocol"`
	ServerType                    types.String                 `tfsdk:"server_type"`
	ServerName                    types.String                 `tfsdk:"server_name"`
	Url                           types.String                 `tfsdk:"url"`
	NetworkIdentity               []ApplicationNetworkIdentity `tfsdk:"network_identity"`
	InstanceId                    types.Int64                  `tfsdk:"instance_id"`
	DisplayName                   types.String                 `tfsdk:"display_name"`
	InstanceVersion               types.Int64                  `tfsdk:"instance_version"`
	Namespace                     types.String                 `tfsdk:"namespace"`
	Qualifier                     types.String                 `tfsdk:"qualifier"`
	ScalableGroupExternalHandle   types.String                 `tfsdk:"scalable_group_external_handle"`
	NetworkApplicationId          types.String                 `tfsdk:"network_application_id"`
	NetworkApplicationName        types.String                 `tfsdk:"network_application_name"`
	ApplicationSubType            types.String                 `tfsdk:"application_sub_type"`
	NetworkApplicationDisplayName types.String                 `tfsdk:"network_application_display_name"`
	EngineId                      types.String                 `tfsdk:"engine_id"`
	Popularity                    types.Int64                  `tfsdk:"popularity"`
	SelectorId                    types.String                 `tfsdk:"selector_id"`
}

type ApplicationNetworkIdentity struct {
	Protocol   types.String `tfsdk:"protocol"`
	Ports      types.String `tfsdk:"ports"`
	LowerPort  types.Int64  `tfsdk:"lower_port"`
	UpperPort  types.Int64  `tfsdk:"upper_port"`
	Ipv4Subnet types.Set    `tfsdk:"ipv4_subnet"`
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

func (data Application) getPathDelete() string {
	return "/dna/intent/api/v1/applications"
}

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
	if !data.ApplicationSetId.IsNull() {
		body, _ = sjson.Set(body, "0.parentScalableGroup.idRef", data.ApplicationSetId.ValueString())
	}
	if !data.CategoryId.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.categoryId", data.CategoryId.ValueString())
	}
	if !data.TrafficClass.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.trafficClass", data.TrafficClass.ValueString())
	}
	if !data.HelpString.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.helpString", data.HelpString.ValueString())
	}
	if !data.Dscp.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.dscp", data.Dscp.ValueString())
	}
	if !data.Rank.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.rank", data.Rank.ValueInt64())
	}
	if !data.AppProtocol.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.appProtocol", data.AppProtocol.ValueString())
	}
	if !data.ServerType.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.type", data.ServerType.ValueString())
	}
	if !data.ServerName.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.serverName", data.ServerName.ValueString())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.url", data.Url.ValueString())
	}
	body, _ = sjson.Set(body, "0.networkApplications.0.applicationType", "CUSTOM")
	body, _ = sjson.Set(body, "0.scalableGroupType", "APPLICATION")
	body, _ = sjson.Set(body, "0.type", "scalablegroup")
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
			body, _ = sjson.SetRaw(body, "0.networkIdentity.-1", itemBody)
		}
	}
	if data.InstanceId.ValueInt64() != 0 && !data.InstanceId.IsNull() {
		body, _ = sjson.Set(body, "0.instanceId", data.InstanceId.ValueInt64())
	}
	if data.DisplayName.ValueString() != "" && !data.DisplayName.IsNull() {
		body, _ = sjson.Set(body, "0.displayName", data.DisplayName.ValueString())
	}
	if data.InstanceVersion.ValueInt64() != 0 && !data.InstanceVersion.IsNull() {
		body, _ = sjson.Set(body, "0.instanceVersion", data.InstanceVersion.ValueInt64())
	}
	if data.Namespace.ValueString() != "" && !data.Namespace.IsNull() {
		body, _ = sjson.Set(body, "0.namespace", data.Namespace.ValueString())
	}
	if data.Qualifier.ValueString() != "" && !data.Qualifier.IsNull() {
		body, _ = sjson.Set(body, "0.qualifier", data.Qualifier.ValueString())
	}
	if data.ScalableGroupExternalHandle.ValueString() != "" && !data.ScalableGroupExternalHandle.IsNull() {
		body, _ = sjson.Set(body, "0.scalableGroupExternalHandle", data.ScalableGroupExternalHandle.ValueString())
	}
	if data.NetworkApplicationId.ValueString() != "" && !data.NetworkApplicationId.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.id", data.NetworkApplicationId.ValueString())
	}
	if data.NetworkApplicationName.ValueString() != "" && !data.NetworkApplicationName.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.name", data.NetworkApplicationName.ValueString())
	}
	if data.ApplicationSubType.ValueString() != "" && !data.ApplicationSubType.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.applicationSubType", data.ApplicationSubType.ValueString())
	}
	if data.NetworkApplicationDisplayName.ValueString() != "" && !data.NetworkApplicationDisplayName.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.displayName", data.NetworkApplicationDisplayName.ValueString())
	}
	if !data.EngineId.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.engineId", data.EngineId.ValueString())
	}
	if data.Popularity.ValueInt64() != 0 && !data.Popularity.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.popularity", data.Popularity.ValueInt64())
	}
	if data.SelectorId.ValueString() != "" && !data.SelectorId.IsNull() {
		body, _ = sjson.Set(body, "0.networkApplications.0.selectorId", data.SelectorId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Application) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	data.Id = types.StringValue(fmt.Sprint(data.Name.ValueString()))
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.parentScalableGroup.idRef"); value.Exists() {
		data.ApplicationSetId = types.StringValue(value.String())
	} else {
		data.ApplicationSetId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.categoryId"); value.Exists() {
		data.CategoryId = types.StringValue(value.String())
	} else {
		data.CategoryId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.trafficClass"); value.Exists() {
		data.TrafficClass = types.StringValue(value.String())
	} else {
		data.TrafficClass = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.helpString"); value.Exists() {
		data.HelpString = types.StringValue(value.String())
	} else {
		data.HelpString = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.dscp"); value.Exists() {
		data.Dscp = types.StringValue(value.String())
	} else {
		data.Dscp = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.rank"); value.Exists() {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
	if value := res.Get("response.0.networkApplications.0.appProtocol"); value.Exists() {
		data.AppProtocol = types.StringValue(value.String())
	} else {
		data.AppProtocol = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.serverName"); value.Exists() {
		data.ServerName = types.StringValue(value.String())
	} else {
		data.ServerName = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.url"); value.Exists() {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
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
			data.NetworkIdentity = append(data.NetworkIdentity, item)
			return true
		})
	}
	if value := res.Get("response.0.instanceId"); value.Exists() {
		data.InstanceId = types.Int64Value(value.Int())
	} else {
		data.InstanceId = types.Int64Null()
	}
	if value := res.Get("response.0.displayName"); value.Exists() {
		data.DisplayName = types.StringValue(value.String())
	} else {
		data.DisplayName = types.StringNull()
	}
	if value := res.Get("response.0.instanceVersion"); value.Exists() {
		data.InstanceVersion = types.Int64Value(value.Int())
	} else {
		data.InstanceVersion = types.Int64Null()
	}
	if value := res.Get("response.0.namespace"); value.Exists() {
		data.Namespace = types.StringValue(value.String())
	} else {
		data.Namespace = types.StringNull()
	}
	if value := res.Get("response.0.qualifier"); value.Exists() {
		data.Qualifier = types.StringValue(value.String())
	} else {
		data.Qualifier = types.StringNull()
	}
	if value := res.Get("response.0.scalableGroupExternalHandle"); value.Exists() {
		data.ScalableGroupExternalHandle = types.StringValue(value.String())
	} else {
		data.ScalableGroupExternalHandle = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.id"); value.Exists() {
		data.NetworkApplicationId = types.StringValue(value.String())
	} else {
		data.NetworkApplicationId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.name"); value.Exists() {
		data.NetworkApplicationName = types.StringValue(value.String())
	} else {
		data.NetworkApplicationName = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.applicationSubType"); value.Exists() {
		data.ApplicationSubType = types.StringValue(value.String())
	} else {
		data.ApplicationSubType = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.displayName"); value.Exists() {
		data.NetworkApplicationDisplayName = types.StringValue(value.String())
	} else {
		data.NetworkApplicationDisplayName = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.engineId"); value.Exists() {
		data.EngineId = types.StringValue(value.String())
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.popularity"); value.Exists() {
		data.Popularity = types.Int64Value(value.Int())
	} else {
		data.Popularity = types.Int64Null()
	}
	if value := res.Get("response.0.networkApplications.0.selectorId"); value.Exists() {
		data.SelectorId = types.StringValue(value.String())
	} else {
		data.SelectorId = types.StringNull()
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
	if value := res.Get("response.0.parentScalableGroup.idRef"); value.Exists() && !data.ApplicationSetId.IsNull() {
		data.ApplicationSetId = types.StringValue(value.String())
	} else {
		data.ApplicationSetId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.categoryId"); value.Exists() && !data.CategoryId.IsNull() {
		data.CategoryId = types.StringValue(value.String())
	} else {
		data.CategoryId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.trafficClass"); value.Exists() && !data.TrafficClass.IsNull() {
		data.TrafficClass = types.StringValue(value.String())
	} else {
		data.TrafficClass = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.helpString"); value.Exists() && !data.HelpString.IsNull() {
		data.HelpString = types.StringValue(value.String())
	} else {
		data.HelpString = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.dscp"); value.Exists() && !data.Dscp.IsNull() {
		data.Dscp = types.StringValue(value.String())
	} else {
		data.Dscp = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.rank"); value.Exists() && !data.Rank.IsNull() {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
	if value := res.Get("response.0.networkApplications.0.appProtocol"); value.Exists() && !data.AppProtocol.IsNull() {
		data.AppProtocol = types.StringValue(value.String())
	} else {
		data.AppProtocol = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.serverName"); value.Exists() && !data.ServerName.IsNull() {
		data.ServerName = types.StringValue(value.String())
	} else {
		data.ServerName = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.url"); value.Exists() && !data.Url.IsNull() {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
	}
	for i := range data.NetworkIdentity {
		keys := [...]string{"protocol"}
		keyValues := [...]string{data.NetworkIdentity[i].Protocol.ValueString()}

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
	}
	if value := res.Get("response.0.instanceId"); value.Exists() && !data.InstanceId.IsNull() {
		data.InstanceId = types.Int64Value(value.Int())
	} else {
		data.InstanceId = types.Int64Null()
	}
	if value := res.Get("response.0.displayName"); value.Exists() && !data.DisplayName.IsNull() {
		data.DisplayName = types.StringValue(value.String())
	} else {
		data.DisplayName = types.StringNull()
	}
	if value := res.Get("response.0.instanceVersion"); value.Exists() && !data.InstanceVersion.IsNull() {
		data.InstanceVersion = types.Int64Value(value.Int())
	} else {
		data.InstanceVersion = types.Int64Null()
	}
	if value := res.Get("response.0.namespace"); value.Exists() && !data.Namespace.IsNull() {
		data.Namespace = types.StringValue(value.String())
	} else {
		data.Namespace = types.StringNull()
	}
	if value := res.Get("response.0.qualifier"); value.Exists() && !data.Qualifier.IsNull() {
		data.Qualifier = types.StringValue(value.String())
	} else {
		data.Qualifier = types.StringNull()
	}
	if value := res.Get("response.0.scalableGroupExternalHandle"); value.Exists() && !data.ScalableGroupExternalHandle.IsNull() {
		data.ScalableGroupExternalHandle = types.StringValue(value.String())
	} else {
		data.ScalableGroupExternalHandle = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.id"); value.Exists() && !data.NetworkApplicationId.IsNull() {
		data.NetworkApplicationId = types.StringValue(value.String())
	} else {
		data.NetworkApplicationId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.name"); value.Exists() && !data.NetworkApplicationName.IsNull() {
		data.NetworkApplicationName = types.StringValue(value.String())
	} else {
		data.NetworkApplicationName = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.applicationSubType"); value.Exists() && !data.ApplicationSubType.IsNull() {
		data.ApplicationSubType = types.StringValue(value.String())
	} else {
		data.ApplicationSubType = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.displayName"); value.Exists() && !data.NetworkApplicationDisplayName.IsNull() {
		data.NetworkApplicationDisplayName = types.StringValue(value.String())
	} else {
		data.NetworkApplicationDisplayName = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.engineId"); value.Exists() && !data.EngineId.IsNull() {
		data.EngineId = types.StringValue(value.String())
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get("response.0.networkApplications.0.popularity"); value.Exists() && !data.Popularity.IsNull() {
		data.Popularity = types.Int64Value(value.Int())
	} else {
		data.Popularity = types.Int64Null()
	}
	if value := res.Get("response.0.networkApplications.0.selectorId"); value.Exists() && !data.SelectorId.IsNull() {
		data.SelectorId = types.StringValue(value.String())
	} else {
		data.SelectorId = types.StringNull()
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
	if data.ApplicationSetId.IsUnknown() {
		if value := res.Get("response.0.parentScalableGroup.idRef"); value.Exists() && !data.ApplicationSetId.IsNull() {
			data.ApplicationSetId = types.StringValue(value.String())
		} else {
			data.ApplicationSetId = types.StringNull()
		}
	}
	if data.CategoryId.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.categoryId"); value.Exists() && !data.CategoryId.IsNull() {
			data.CategoryId = types.StringValue(value.String())
		} else {
			data.CategoryId = types.StringNull()
		}
	}
	if data.TrafficClass.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.trafficClass"); value.Exists() && !data.TrafficClass.IsNull() {
			data.TrafficClass = types.StringValue(value.String())
		} else {
			data.TrafficClass = types.StringNull()
		}
	}
	if data.HelpString.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.helpString"); value.Exists() && !data.HelpString.IsNull() {
			data.HelpString = types.StringValue(value.String())
		} else {
			data.HelpString = types.StringNull()
		}
	}
	if data.Dscp.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.dscp"); value.Exists() && !data.Dscp.IsNull() {
			data.Dscp = types.StringValue(value.String())
		} else {
			data.Dscp = types.StringNull()
		}
	}
	if data.Rank.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.rank"); value.Exists() && !data.Rank.IsNull() {
			data.Rank = types.Int64Value(value.Int())
		} else {
			data.Rank = types.Int64Null()
		}
	}
	if data.AppProtocol.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.appProtocol"); value.Exists() && !data.AppProtocol.IsNull() {
			data.AppProtocol = types.StringValue(value.String())
		} else {
			data.AppProtocol = types.StringNull()
		}
	}
	if data.ServerName.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.serverName"); value.Exists() && !data.ServerName.IsNull() {
			data.ServerName = types.StringValue(value.String())
		} else {
			data.ServerName = types.StringNull()
		}
	}
	if data.Url.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.url"); value.Exists() && !data.Url.IsNull() {
			data.Url = types.StringValue(value.String())
		} else {
			data.Url = types.StringNull()
		}
	}
	if data.InstanceId.IsUnknown() {
		if value := res.Get("response.0.instanceId"); value.Exists() && !data.InstanceId.IsNull() {
			data.InstanceId = types.Int64Value(value.Int())
		} else {
			data.InstanceId = types.Int64Null()
		}
	}
	if data.DisplayName.IsUnknown() {
		if value := res.Get("response.0.displayName"); value.Exists() && !data.DisplayName.IsNull() {
			data.DisplayName = types.StringValue(value.String())
		} else {
			data.DisplayName = types.StringNull()
		}
	}
	if data.InstanceVersion.IsUnknown() {
		if value := res.Get("response.0.instanceVersion"); value.Exists() && !data.InstanceVersion.IsNull() {
			data.InstanceVersion = types.Int64Value(value.Int())
		} else {
			data.InstanceVersion = types.Int64Null()
		}
	}
	if data.Namespace.IsUnknown() {
		if value := res.Get("response.0.namespace"); value.Exists() && !data.Namespace.IsNull() {
			data.Namespace = types.StringValue(value.String())
		} else {
			data.Namespace = types.StringNull()
		}
	}
	if data.Qualifier.IsUnknown() {
		if value := res.Get("response.0.qualifier"); value.Exists() && !data.Qualifier.IsNull() {
			data.Qualifier = types.StringValue(value.String())
		} else {
			data.Qualifier = types.StringNull()
		}
	}
	if data.ScalableGroupExternalHandle.IsUnknown() {
		if value := res.Get("response.0.scalableGroupExternalHandle"); value.Exists() && !data.ScalableGroupExternalHandle.IsNull() {
			data.ScalableGroupExternalHandle = types.StringValue(value.String())
		} else {
			data.ScalableGroupExternalHandle = types.StringNull()
		}
	}
	if data.NetworkApplicationId.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.id"); value.Exists() && !data.NetworkApplicationId.IsNull() {
			data.NetworkApplicationId = types.StringValue(value.String())
		} else {
			data.NetworkApplicationId = types.StringNull()
		}
	}
	if data.NetworkApplicationName.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.name"); value.Exists() && !data.NetworkApplicationName.IsNull() {
			data.NetworkApplicationName = types.StringValue(value.String())
		} else {
			data.NetworkApplicationName = types.StringNull()
		}
	}
	if data.ApplicationSubType.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.applicationSubType"); value.Exists() && !data.ApplicationSubType.IsNull() {
			data.ApplicationSubType = types.StringValue(value.String())
		} else {
			data.ApplicationSubType = types.StringNull()
		}
	}
	if data.NetworkApplicationDisplayName.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.displayName"); value.Exists() && !data.NetworkApplicationDisplayName.IsNull() {
			data.NetworkApplicationDisplayName = types.StringValue(value.String())
		} else {
			data.NetworkApplicationDisplayName = types.StringNull()
		}
	}
	if data.EngineId.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.engineId"); value.Exists() && !data.EngineId.IsNull() {
			data.EngineId = types.StringValue(value.String())
		} else {
			data.EngineId = types.StringNull()
		}
	}
	if data.Popularity.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.popularity"); value.Exists() && !data.Popularity.IsNull() {
			data.Popularity = types.Int64Value(value.Int())
		} else {
			data.Popularity = types.Int64Null()
		}
	}
	if data.SelectorId.IsUnknown() {
		if value := res.Get("response.0.networkApplications.0.selectorId"); value.Exists() && !data.SelectorId.IsNull() {
			data.SelectorId = types.StringValue(value.String())
		} else {
			data.SelectorId = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Application) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.ApplicationSetId.IsNull() {
		return false
	}
	if !data.CategoryId.IsNull() {
		return false
	}
	if !data.TrafficClass.IsNull() {
		return false
	}
	if !data.HelpString.IsNull() {
		return false
	}
	if !data.Dscp.IsNull() {
		return false
	}
	if !data.Rank.IsNull() {
		return false
	}
	if !data.AppProtocol.IsNull() {
		return false
	}
	if !data.ServerType.IsNull() {
		return false
	}
	if !data.ServerName.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if len(data.NetworkIdentity) > 0 {
		return false
	}
	if !data.InstanceId.IsNull() {
		return false
	}
	if !data.DisplayName.IsNull() {
		return false
	}
	if !data.InstanceVersion.IsNull() {
		return false
	}
	if !data.Namespace.IsNull() {
		return false
	}
	if !data.Qualifier.IsNull() {
		return false
	}
	if !data.ScalableGroupExternalHandle.IsNull() {
		return false
	}
	if !data.NetworkApplicationId.IsNull() {
		return false
	}
	if !data.NetworkApplicationName.IsNull() {
		return false
	}
	if !data.ApplicationSubType.IsNull() {
		return false
	}
	if !data.NetworkApplicationDisplayName.IsNull() {
		return false
	}
	if !data.EngineId.IsNull() {
		return false
	}
	if !data.Popularity.IsNull() {
		return false
	}
	if !data.SelectorId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
