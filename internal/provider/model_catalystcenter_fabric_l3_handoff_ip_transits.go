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
type FabricL3HandoffIPTransits struct {
	Id              types.String                          `tfsdk:"id"`
	NetworkDeviceId types.String                          `tfsdk:"network_device_id"`
	FabricId        types.String                          `tfsdk:"fabric_id"`
	L3Handoffs      []FabricL3HandoffIPTransitsL3Handoffs `tfsdk:"l3_handoffs"`
}

type FabricL3HandoffIPTransitsL3Handoffs struct {
	Id                             types.String `tfsdk:"id"`
	FabricId                       types.String `tfsdk:"fabric_id"`
	NetworkDeviceId                types.String `tfsdk:"network_device_id"`
	TransitNetworkId               types.String `tfsdk:"transit_network_id"`
	InterfaceName                  types.String `tfsdk:"interface_name"`
	ExternalConnectivityIpPoolName types.String `tfsdk:"external_connectivity_ip_pool_name"`
	VirtualNetworkName             types.String `tfsdk:"virtual_network_name"`
	VlanId                         types.Int64  `tfsdk:"vlan_id"`
	TcpMssAdjustment               types.Int64  `tfsdk:"tcp_mss_adjustment"`
	LocalIpAddress                 types.String `tfsdk:"local_ip_address"`
	RemoteIpAddress                types.String `tfsdk:"remote_ip_address"`
	LocalIpv6Address               types.String `tfsdk:"local_ipv6_address"`
	RemoteIpv6Address              types.String `tfsdk:"remote_ipv6_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricL3HandoffIPTransits) getPath() string {
	return "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricL3HandoffIPTransits) toBody(ctx context.Context, state FabricL3HandoffIPTransits) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "", data.NetworkDeviceId.ValueString())
	}
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "", data.FabricId.ValueString())
	}
	if len(data.L3Handoffs) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.L3Handoffs {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.FabricId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fabricId", item.FabricId.ValueString())
			}
			if !item.NetworkDeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkDeviceId", item.NetworkDeviceId.ValueString())
			}
			if !item.TransitNetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "transitNetworkId", item.TransitNetworkId.ValueString())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceName", item.InterfaceName.ValueString())
			}
			if !item.ExternalConnectivityIpPoolName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "externalConnectivityIpPoolName", item.ExternalConnectivityIpPoolName.ValueString())
			}
			if !item.VirtualNetworkName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "virtualNetworkName", item.VirtualNetworkName.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueInt64())
			}
			if !item.TcpMssAdjustment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tcpMssAdjustment", item.TcpMssAdjustment.ValueInt64())
			}
			if !item.LocalIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localIpAddress", item.LocalIpAddress.ValueString())
			}
			if !item.RemoteIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteIpAddress", item.RemoteIpAddress.ValueString())
			}
			if !item.LocalIpv6Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localIpv6Address", item.LocalIpv6Address.ValueString())
			}
			if !item.RemoteIpv6Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteIpv6Address", item.RemoteIpv6Address.ValueString())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricL3HandoffIPTransits) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.L3Handoffs = make([]FabricL3HandoffIPTransitsL3Handoffs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricL3HandoffIPTransitsL3Handoffs{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("fabricId"); cValue.Exists() {
				item.FabricId = types.StringValue(cValue.String())
			} else {
				item.FabricId = types.StringNull()
			}
			if cValue := v.Get("networkDeviceId"); cValue.Exists() {
				item.NetworkDeviceId = types.StringValue(cValue.String())
			} else {
				item.NetworkDeviceId = types.StringNull()
			}
			if cValue := v.Get("transitNetworkId"); cValue.Exists() {
				item.TransitNetworkId = types.StringValue(cValue.String())
			} else {
				item.TransitNetworkId = types.StringNull()
			}
			if cValue := v.Get("interfaceName"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			} else {
				item.InterfaceName = types.StringNull()
			}
			if cValue := v.Get("externalConnectivityIpPoolName"); cValue.Exists() {
				item.ExternalConnectivityIpPoolName = types.StringValue(cValue.String())
			} else {
				item.ExternalConnectivityIpPoolName = types.StringNull()
			}
			if cValue := v.Get("virtualNetworkName"); cValue.Exists() {
				item.VirtualNetworkName = types.StringValue(cValue.String())
			} else {
				item.VirtualNetworkName = types.StringNull()
			}
			if cValue := v.Get("vlanId"); cValue.Exists() {
				item.VlanId = types.Int64Value(cValue.Int())
			} else {
				item.VlanId = types.Int64Null()
			}
			if cValue := v.Get("localIpAddress"); cValue.Exists() {
				item.LocalIpAddress = types.StringValue(cValue.String())
			} else {
				item.LocalIpAddress = types.StringNull()
			}
			if cValue := v.Get("remoteIpAddress"); cValue.Exists() {
				item.RemoteIpAddress = types.StringValue(cValue.String())
			} else {
				item.RemoteIpAddress = types.StringNull()
			}
			if cValue := v.Get("localIpv6Address"); cValue.Exists() {
				item.LocalIpv6Address = types.StringValue(cValue.String())
			} else {
				item.LocalIpv6Address = types.StringNull()
			}
			if cValue := v.Get("remoteIpv6Address"); cValue.Exists() {
				item.RemoteIpv6Address = types.StringValue(cValue.String())
			} else {
				item.RemoteIpv6Address = types.StringNull()
			}
			data.L3Handoffs = append(data.L3Handoffs, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricL3HandoffIPTransits) updateFromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.L3Handoffs {
		keys := [...]string{"vlanId"}
		keyValues := [...]string{strconv.FormatInt(data.L3Handoffs[i].VlanId.ValueInt64(), 10)}

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
		if value := r.Get("id"); value.Exists() && !data.L3Handoffs[i].Id.IsNull() {
			data.L3Handoffs[i].Id = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].Id = types.StringNull()
		}
		if value := r.Get("fabricId"); value.Exists() && !data.L3Handoffs[i].FabricId.IsNull() {
			data.L3Handoffs[i].FabricId = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].FabricId = types.StringNull()
		}
		if value := r.Get("networkDeviceId"); value.Exists() && !data.L3Handoffs[i].NetworkDeviceId.IsNull() {
			data.L3Handoffs[i].NetworkDeviceId = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].NetworkDeviceId = types.StringNull()
		}
		if value := r.Get("transitNetworkId"); value.Exists() && !data.L3Handoffs[i].TransitNetworkId.IsNull() {
			data.L3Handoffs[i].TransitNetworkId = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].TransitNetworkId = types.StringNull()
		}
		if value := r.Get("interfaceName"); value.Exists() && !data.L3Handoffs[i].InterfaceName.IsNull() {
			data.L3Handoffs[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("externalConnectivityIpPoolName"); value.Exists() && !data.L3Handoffs[i].ExternalConnectivityIpPoolName.IsNull() {
			data.L3Handoffs[i].ExternalConnectivityIpPoolName = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].ExternalConnectivityIpPoolName = types.StringNull()
		}
		if value := r.Get("virtualNetworkName"); value.Exists() && !data.L3Handoffs[i].VirtualNetworkName.IsNull() {
			data.L3Handoffs[i].VirtualNetworkName = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].VirtualNetworkName = types.StringNull()
		}
		if value := r.Get("vlanId"); value.Exists() && !data.L3Handoffs[i].VlanId.IsNull() {
			data.L3Handoffs[i].VlanId = types.Int64Value(value.Int())
		} else {
			data.L3Handoffs[i].VlanId = types.Int64Null()
		}
		if value := r.Get("localIpAddress"); value.Exists() && !data.L3Handoffs[i].LocalIpAddress.IsNull() {
			data.L3Handoffs[i].LocalIpAddress = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].LocalIpAddress = types.StringNull()
		}
		if value := r.Get("remoteIpAddress"); value.Exists() && !data.L3Handoffs[i].RemoteIpAddress.IsNull() {
			data.L3Handoffs[i].RemoteIpAddress = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].RemoteIpAddress = types.StringNull()
		}
		if value := r.Get("localIpv6Address"); value.Exists() && !data.L3Handoffs[i].LocalIpv6Address.IsNull() {
			data.L3Handoffs[i].LocalIpv6Address = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].LocalIpv6Address = types.StringNull()
		}
		if value := r.Get("remoteIpv6Address"); value.Exists() && !data.L3Handoffs[i].RemoteIpv6Address.IsNull() {
			data.L3Handoffs[i].RemoteIpv6Address = types.StringValue(value.String())
		} else {
			data.L3Handoffs[i].RemoteIpv6Address = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FabricL3HandoffIPTransits) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.L3Handoffs {
		keys := [...]string{"vlanId"}
		keyValues := [...]string{strconv.FormatInt(data.L3Handoffs[i].VlanId.ValueInt64(), 10)}

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
		if data.L3Handoffs[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.L3Handoffs[i].Id.IsNull() {
				data.L3Handoffs[i].Id = types.StringValue(value.String())
			} else {
				data.L3Handoffs[i].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricL3HandoffIPTransits) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.L3Handoffs) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
