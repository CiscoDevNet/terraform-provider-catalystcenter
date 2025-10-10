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
type AnycastGateways struct {
	Id              types.String                     `tfsdk:"id"`
	FabricId        types.String                     `tfsdk:"fabric_id"`
	AnycastGateways []AnycastGatewaysAnycastGateways `tfsdk:"anycast_gateways"`
}

type AnycastGatewaysAnycastGateways struct {
	Id                                    types.String `tfsdk:"id"`
	FabricId                              types.String `tfsdk:"fabric_id"`
	VirtualNetworkName                    types.String `tfsdk:"virtual_network_name"`
	IpPoolName                            types.String `tfsdk:"ip_pool_name"`
	TcpMssAdjustment                      types.Int64  `tfsdk:"tcp_mss_adjustment"`
	VlanName                              types.String `tfsdk:"vlan_name"`
	VlanId                                types.Int64  `tfsdk:"vlan_id"`
	TrafficType                           types.String `tfsdk:"traffic_type"`
	PoolType                              types.String `tfsdk:"pool_type"`
	SecurityGroupName                     types.String `tfsdk:"security_group_name"`
	CriticalPool                          types.Bool   `tfsdk:"critical_pool"`
	L2FloodingEnabled                     types.Bool   `tfsdk:"l2_flooding_enabled"`
	WirelessPool                          types.Bool   `tfsdk:"wireless_pool"`
	IpDirectedBroadcast                   types.Bool   `tfsdk:"ip_directed_broadcast"`
	IntraSubnetRoutingEnabled             types.Bool   `tfsdk:"intra_subnet_routing_enabled"`
	MultipleIpToMacAddresses              types.Bool   `tfsdk:"multiple_ip_to_mac_addresses"`
	SupplicantBasedExtendedNodeOnboarding types.Bool   `tfsdk:"supplicant_based_extended_node_onboarding"`
	GroupBasedPolicyEnforcementEnabled    types.Bool   `tfsdk:"group_based_policy_enforcement_enabled"`
	AutoGenerateVlanName                  types.Bool   `tfsdk:"auto_generate_vlan_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AnycastGateways) getPath() string {
	return "/dna/intent/api/v1/sda/anycastGateways"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AnycastGateways) toBody(ctx context.Context, state AnycastGateways) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "fabricId", data.FabricId.ValueString())
	}
	if len(data.AnycastGateways) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.AnycastGateways {
			itemBody := ""
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.FabricId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fabricId", item.FabricId.ValueString())
			}
			if !item.VirtualNetworkName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "virtualNetworkName", item.VirtualNetworkName.ValueString())
			}
			if !item.IpPoolName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipPoolName", item.IpPoolName.ValueString())
			}
			if !item.TcpMssAdjustment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tcpMssAdjustment", item.TcpMssAdjustment.ValueInt64())
			}
			if item.VlanName.ValueString() != "" && !item.VlanName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanName", item.VlanName.ValueString())
			}
			if item.VlanId.ValueInt64() != 0 && !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueInt64())
			}
			if !item.TrafficType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "trafficType", item.TrafficType.ValueString())
			}
			if !item.PoolType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "poolType", item.PoolType.ValueString())
			}
			if !item.SecurityGroupName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "securityGroupName", item.SecurityGroupName.ValueString())
			}
			if !item.CriticalPool.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isCriticalPool", item.CriticalPool.ValueBool())
			}
			if !item.L2FloodingEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isLayer2FloodingEnabled", item.L2FloodingEnabled.ValueBool())
			}
			if !item.WirelessPool.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isWirelessPool", item.WirelessPool.ValueBool())
			}
			if !item.IpDirectedBroadcast.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isIpDirectedBroadcast", item.IpDirectedBroadcast.ValueBool())
			}
			if !item.IntraSubnetRoutingEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isIntraSubnetRoutingEnabled", item.IntraSubnetRoutingEnabled.ValueBool())
			}
			if !item.MultipleIpToMacAddresses.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isMultipleIpToMacAddresses", item.MultipleIpToMacAddresses.ValueBool())
			}
			if !item.SupplicantBasedExtendedNodeOnboarding.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isSupplicantBasedExtendedNodeOnboarding", item.SupplicantBasedExtendedNodeOnboarding.ValueBool())
			}
			if !item.GroupBasedPolicyEnforcementEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isGroupBasedPolicyEnforcementEnabled", item.GroupBasedPolicyEnforcementEnabled.ValueBool())
			}
			if !item.AutoGenerateVlanName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "autoGenerateVlanName", item.AutoGenerateVlanName.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AnycastGateways) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.AnycastGateways = make([]AnycastGatewaysAnycastGateways, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AnycastGatewaysAnycastGateways{}
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
			if cValue := v.Get("virtualNetworkName"); cValue.Exists() {
				item.VirtualNetworkName = types.StringValue(cValue.String())
			} else {
				item.VirtualNetworkName = types.StringNull()
			}
			if cValue := v.Get("ipPoolName"); cValue.Exists() {
				item.IpPoolName = types.StringValue(cValue.String())
			} else {
				item.IpPoolName = types.StringNull()
			}
			if cValue := v.Get("tcpMssAdjustment"); cValue.Exists() {
				item.TcpMssAdjustment = types.Int64Value(cValue.Int())
			} else {
				item.TcpMssAdjustment = types.Int64Null()
			}
			if cValue := v.Get("vlanName"); cValue.Exists() {
				item.VlanName = types.StringValue(cValue.String())
			} else {
				item.VlanName = types.StringNull()
			}
			if cValue := v.Get("vlanId"); cValue.Exists() {
				item.VlanId = types.Int64Value(cValue.Int())
			} else {
				item.VlanId = types.Int64Null()
			}
			if cValue := v.Get("trafficType"); cValue.Exists() {
				item.TrafficType = types.StringValue(cValue.String())
			} else {
				item.TrafficType = types.StringNull()
			}
			if cValue := v.Get("poolType"); cValue.Exists() {
				item.PoolType = types.StringValue(cValue.String())
			} else {
				item.PoolType = types.StringNull()
			}
			if cValue := v.Get("securityGroupName"); cValue.Exists() {
				item.SecurityGroupName = types.StringValue(cValue.String())
			} else {
				item.SecurityGroupName = types.StringNull()
			}
			if cValue := v.Get("isCriticalPool"); cValue.Exists() {
				item.CriticalPool = types.BoolValue(cValue.Bool())
			} else {
				item.CriticalPool = types.BoolNull()
			}
			if cValue := v.Get("isLayer2FloodingEnabled"); cValue.Exists() {
				item.L2FloodingEnabled = types.BoolValue(cValue.Bool())
			} else {
				item.L2FloodingEnabled = types.BoolNull()
			}
			if cValue := v.Get("isWirelessPool"); cValue.Exists() {
				item.WirelessPool = types.BoolValue(cValue.Bool())
			} else {
				item.WirelessPool = types.BoolNull()
			}
			if cValue := v.Get("isIpDirectedBroadcast"); cValue.Exists() {
				item.IpDirectedBroadcast = types.BoolValue(cValue.Bool())
			} else {
				item.IpDirectedBroadcast = types.BoolNull()
			}
			if cValue := v.Get("isIntraSubnetRoutingEnabled"); cValue.Exists() {
				item.IntraSubnetRoutingEnabled = types.BoolValue(cValue.Bool())
			} else {
				item.IntraSubnetRoutingEnabled = types.BoolNull()
			}
			if cValue := v.Get("isMultipleIpToMacAddresses"); cValue.Exists() {
				item.MultipleIpToMacAddresses = types.BoolValue(cValue.Bool())
			} else {
				item.MultipleIpToMacAddresses = types.BoolNull()
			}
			if cValue := v.Get("isSupplicantBasedExtendedNodeOnboarding"); cValue.Exists() {
				item.SupplicantBasedExtendedNodeOnboarding = types.BoolValue(cValue.Bool())
			} else {
				item.SupplicantBasedExtendedNodeOnboarding = types.BoolNull()
			}
			if cValue := v.Get("isGroupBasedPolicyEnforcementEnabled"); cValue.Exists() {
				item.GroupBasedPolicyEnforcementEnabled = types.BoolValue(cValue.Bool())
			} else {
				item.GroupBasedPolicyEnforcementEnabled = types.BoolNull()
			}
			item.AutoGenerateVlanName = types.BoolValue(false)
			data.AnycastGateways = append(data.AnycastGateways, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AnycastGateways) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []AnycastGatewaysAnycastGateways

	res = res.Get("response")
	for i := range data.AnycastGateways {
		keys := [...]string{"ipPoolName"}
		keyValues := [...]string{data.AnycastGateways[i].IpPoolName.ValueString()}

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
		if value := r.Get("id"); value.Exists() && !data.AnycastGateways[i].Id.IsNull() {
			data.AnycastGateways[i].Id = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].Id = types.StringNull()
		}
		if value := r.Get("fabricId"); value.Exists() && !data.AnycastGateways[i].FabricId.IsNull() {
			data.AnycastGateways[i].FabricId = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].FabricId = types.StringNull()
		}
		if value := r.Get("virtualNetworkName"); value.Exists() && !data.AnycastGateways[i].VirtualNetworkName.IsNull() {
			data.AnycastGateways[i].VirtualNetworkName = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].VirtualNetworkName = types.StringNull()
		}
		if value := r.Get("ipPoolName"); value.Exists() && !data.AnycastGateways[i].IpPoolName.IsNull() {
			data.AnycastGateways[i].IpPoolName = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].IpPoolName = types.StringNull()
		}
		if value := r.Get("tcpMssAdjustment"); value.Exists() && !data.AnycastGateways[i].TcpMssAdjustment.IsNull() {
			data.AnycastGateways[i].TcpMssAdjustment = types.Int64Value(value.Int())
		} else {
			data.AnycastGateways[i].TcpMssAdjustment = types.Int64Null()
		}
		if value := r.Get("vlanName"); value.Exists() && !data.AnycastGateways[i].VlanName.IsNull() {
			data.AnycastGateways[i].VlanName = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].VlanName = types.StringNull()
		}
		if value := r.Get("vlanId"); value.Exists() && !data.AnycastGateways[i].VlanId.IsNull() {
			data.AnycastGateways[i].VlanId = types.Int64Value(value.Int())
		} else {
			data.AnycastGateways[i].VlanId = types.Int64Null()
		}
		if value := r.Get("trafficType"); value.Exists() && !data.AnycastGateways[i].TrafficType.IsNull() {
			data.AnycastGateways[i].TrafficType = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].TrafficType = types.StringNull()
		}
		if value := r.Get("poolType"); value.Exists() && !data.AnycastGateways[i].PoolType.IsNull() {
			data.AnycastGateways[i].PoolType = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].PoolType = types.StringNull()
		}
		if value := r.Get("securityGroupName"); value.Exists() && !data.AnycastGateways[i].SecurityGroupName.IsNull() {
			data.AnycastGateways[i].SecurityGroupName = types.StringValue(value.String())
		} else {
			data.AnycastGateways[i].SecurityGroupName = types.StringNull()
		}
		if value := r.Get("isCriticalPool"); value.Exists() && !data.AnycastGateways[i].CriticalPool.IsNull() {
			data.AnycastGateways[i].CriticalPool = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].CriticalPool = types.BoolNull()
		}
		if value := r.Get("isLayer2FloodingEnabled"); value.Exists() && !data.AnycastGateways[i].L2FloodingEnabled.IsNull() {
			data.AnycastGateways[i].L2FloodingEnabled = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].L2FloodingEnabled = types.BoolNull()
		}
		if value := r.Get("isWirelessPool"); value.Exists() && !data.AnycastGateways[i].WirelessPool.IsNull() {
			data.AnycastGateways[i].WirelessPool = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].WirelessPool = types.BoolNull()
		}
		if value := r.Get("isIpDirectedBroadcast"); value.Exists() && !data.AnycastGateways[i].IpDirectedBroadcast.IsNull() {
			data.AnycastGateways[i].IpDirectedBroadcast = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].IpDirectedBroadcast = types.BoolNull()
		}
		if value := r.Get("isIntraSubnetRoutingEnabled"); value.Exists() && !data.AnycastGateways[i].IntraSubnetRoutingEnabled.IsNull() {
			data.AnycastGateways[i].IntraSubnetRoutingEnabled = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].IntraSubnetRoutingEnabled = types.BoolNull()
		}
		if value := r.Get("isMultipleIpToMacAddresses"); value.Exists() && !data.AnycastGateways[i].MultipleIpToMacAddresses.IsNull() {
			data.AnycastGateways[i].MultipleIpToMacAddresses = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].MultipleIpToMacAddresses = types.BoolNull()
		}
		if value := r.Get("isSupplicantBasedExtendedNodeOnboarding"); value.Exists() && !data.AnycastGateways[i].SupplicantBasedExtendedNodeOnboarding.IsNull() {
			data.AnycastGateways[i].SupplicantBasedExtendedNodeOnboarding = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].SupplicantBasedExtendedNodeOnboarding = types.BoolNull()
		}
		if value := r.Get("isGroupBasedPolicyEnforcementEnabled"); value.Exists() && !data.AnycastGateways[i].GroupBasedPolicyEnforcementEnabled.IsNull() {
			data.AnycastGateways[i].GroupBasedPolicyEnforcementEnabled = types.BoolValue(value.Bool())
		} else {
			data.AnycastGateways[i].GroupBasedPolicyEnforcementEnabled = types.BoolNull()
		}
		if data.AnycastGateways[i].Id != types.StringNull() {
			final = append(final, data.AnycastGateways[i])
		}
	}
	data.AnycastGateways = final
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AnycastGateways) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.AnycastGateways {
		keys := [...]string{"ipPoolName"}
		keyValues := [...]string{data.AnycastGateways[i].IpPoolName.ValueString()}

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
		if data.AnycastGateways[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.AnycastGateways[i].Id.IsNull() {
				data.AnycastGateways[i].Id = types.StringValue(value.String())
			} else {
				data.AnycastGateways[i].Id = types.StringNull()
			}
		}
		if data.AnycastGateways[i].VlanName.IsUnknown() {
			if value := r.Get("vlanName"); value.Exists() && !data.AnycastGateways[i].VlanName.IsNull() {
				data.AnycastGateways[i].VlanName = types.StringValue(value.String())
			} else {
				data.AnycastGateways[i].VlanName = types.StringNull()
			}
		}
		if data.AnycastGateways[i].VlanId.IsUnknown() {
			if value := r.Get("vlanId"); value.Exists() && !data.AnycastGateways[i].VlanId.IsNull() {
				data.AnycastGateways[i].VlanId = types.Int64Value(value.Int())
			} else {
				data.AnycastGateways[i].VlanId = types.Int64Null()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AnycastGateways) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.AnycastGateways) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
