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
type AnycastGateway struct {
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
	AutoGenerateVlanName                  types.Bool   `tfsdk:"auto_generate_vlan_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AnycastGateway) getPath() string {
	return "/dna/intent/api/v1/sda/anycastGateways"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AnycastGateway) toBody(ctx context.Context, state AnycastGateway) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.IpPoolName.IsNull() {
		body, _ = sjson.Set(body, "0.ipPoolName", data.IpPoolName.ValueString())
	}
	if !data.TcpMssAdjustment.IsNull() {
		body, _ = sjson.Set(body, "0.tcpMssAdjustment", data.TcpMssAdjustment.ValueInt64())
	}
	if !data.VlanName.IsNull() {
		body, _ = sjson.Set(body, "0.vlanName", data.VlanName.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "0.vlanId", data.VlanId.ValueInt64())
	}
	if !data.TrafficType.IsNull() {
		body, _ = sjson.Set(body, "0.trafficType", data.TrafficType.ValueString())
	}
	if !data.PoolType.IsNull() {
		body, _ = sjson.Set(body, "0.poolType", data.PoolType.ValueString())
	}
	if !data.SecurityGroupName.IsNull() {
		body, _ = sjson.Set(body, "0.securityGroupName", data.SecurityGroupName.ValueString())
	}
	if !data.CriticalPool.IsNull() {
		body, _ = sjson.Set(body, "0.isCriticalPool", data.CriticalPool.ValueBool())
	}
	if !data.L2FloodingEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.isLayer2FloodingEnabled", data.L2FloodingEnabled.ValueBool())
	}
	if !data.WirelessPool.IsNull() {
		body, _ = sjson.Set(body, "0.isWirelessPool", data.WirelessPool.ValueBool())
	}
	if !data.IpDirectedBroadcast.IsNull() {
		body, _ = sjson.Set(body, "0.isIpDirectedBroadcast", data.IpDirectedBroadcast.ValueBool())
	}
	if !data.IntraSubnetRoutingEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.isIntraSubnetRoutingEnabled", data.IntraSubnetRoutingEnabled.ValueBool())
	}
	if !data.MultipleIpToMacAddresses.IsNull() {
		body, _ = sjson.Set(body, "0.isMultipleIpToMacAddresses", data.MultipleIpToMacAddresses.ValueBool())
	}
	if !data.SupplicantBasedExtendedNodeOnboarding.IsNull() {
		body, _ = sjson.Set(body, "0.isSupplicantBasedExtendedNodeOnboarding", data.SupplicantBasedExtendedNodeOnboarding.ValueBool())
	}
	if !data.AutoGenerateVlanName.IsNull() {
		body, _ = sjson.Set(body, "0.autoGenerateVlanName", data.AutoGenerateVlanName.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AnycastGateway) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.ipPoolName"); value.Exists() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
	if value := res.Get("response.0.tcpMssAdjustment"); value.Exists() {
		data.TcpMssAdjustment = types.Int64Value(value.Int())
	} else {
		data.TcpMssAdjustment = types.Int64Null()
	}
	if value := res.Get("response.0.vlanName"); value.Exists() {
		data.VlanName = types.StringValue(value.String())
	} else {
		data.VlanName = types.StringNull()
	}
	if value := res.Get("response.0.vlanId"); value.Exists() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("response.0.trafficType"); value.Exists() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("response.0.poolType"); value.Exists() {
		data.PoolType = types.StringValue(value.String())
	} else {
		data.PoolType = types.StringNull()
	}
	if value := res.Get("response.0.securityGroupNames"); value.Exists() {
		data.SecurityGroupName = types.StringValue(value.String())
	} else {
		data.SecurityGroupName = types.StringNull()
	}
	if value := res.Get("response.0.isCriticalPool"); value.Exists() {
		data.CriticalPool = types.BoolValue(value.Bool())
	} else {
		data.CriticalPool = types.BoolNull()
	}
	if value := res.Get("response.0.isLayer2FloodingEnabled"); value.Exists() {
		data.L2FloodingEnabled = types.BoolValue(value.Bool())
	} else {
		data.L2FloodingEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.isWirelessPool"); value.Exists() {
		data.WirelessPool = types.BoolValue(value.Bool())
	} else {
		data.WirelessPool = types.BoolNull()
	}
	if value := res.Get("response.0.isIpDirectedBroadcast"); value.Exists() {
		data.IpDirectedBroadcast = types.BoolValue(value.Bool())
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
	}
	if value := res.Get("response.0.isIntraSubnetRoutingEnabled"); value.Exists() {
		data.IntraSubnetRoutingEnabled = types.BoolValue(value.Bool())
	} else {
		data.IntraSubnetRoutingEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.isMultipleIpToMacAddresses"); value.Exists() {
		data.MultipleIpToMacAddresses = types.BoolValue(value.Bool())
	} else {
		data.MultipleIpToMacAddresses = types.BoolNull()
	}
	if value := res.Get("response.0.isSupplicantBasedExtendedNodeOnboarding"); value.Exists() {
		data.SupplicantBasedExtendedNodeOnboarding = types.BoolValue(value.Bool())
	} else {
		data.SupplicantBasedExtendedNodeOnboarding = types.BoolNull()
	}
	if value := res.Get("response.0.autoGenerateVlanName"); value.Exists() {
		data.AutoGenerateVlanName = types.BoolValue(value.Bool())
	} else {
		data.AutoGenerateVlanName = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AnycastGateway) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.ipPoolName"); value.Exists() && !data.IpPoolName.IsNull() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
	if value := res.Get("response.0.tcpMssAdjustment"); value.Exists() && !data.TcpMssAdjustment.IsNull() {
		data.TcpMssAdjustment = types.Int64Value(value.Int())
	} else {
		data.TcpMssAdjustment = types.Int64Null()
	}
	if value := res.Get("response.0.vlanName"); value.Exists() && !data.VlanName.IsNull() {
		data.VlanName = types.StringValue(value.String())
	} else {
		data.VlanName = types.StringNull()
	}
	if value := res.Get("response.0.vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("response.0.trafficType"); value.Exists() && !data.TrafficType.IsNull() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("response.0.poolType"); value.Exists() && !data.PoolType.IsNull() {
		data.PoolType = types.StringValue(value.String())
	} else {
		data.PoolType = types.StringNull()
	}
	if value := res.Get("response.0.securityGroupNames"); value.Exists() && !data.SecurityGroupName.IsNull() {
		data.SecurityGroupName = types.StringValue(value.String())
	} else {
		data.SecurityGroupName = types.StringNull()
	}
	if value := res.Get("response.0.isCriticalPool"); value.Exists() && !data.CriticalPool.IsNull() {
		data.CriticalPool = types.BoolValue(value.Bool())
	} else {
		data.CriticalPool = types.BoolNull()
	}
	if value := res.Get("response.0.isLayer2FloodingEnabled"); value.Exists() && !data.L2FloodingEnabled.IsNull() {
		data.L2FloodingEnabled = types.BoolValue(value.Bool())
	} else {
		data.L2FloodingEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.isWirelessPool"); value.Exists() && !data.WirelessPool.IsNull() {
		data.WirelessPool = types.BoolValue(value.Bool())
	} else {
		data.WirelessPool = types.BoolNull()
	}
	if value := res.Get("response.0.isIpDirectedBroadcast"); value.Exists() && !data.IpDirectedBroadcast.IsNull() {
		data.IpDirectedBroadcast = types.BoolValue(value.Bool())
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
	}
	if value := res.Get("response.0.isIntraSubnetRoutingEnabled"); value.Exists() && !data.IntraSubnetRoutingEnabled.IsNull() {
		data.IntraSubnetRoutingEnabled = types.BoolValue(value.Bool())
	} else {
		data.IntraSubnetRoutingEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.isMultipleIpToMacAddresses"); value.Exists() && !data.MultipleIpToMacAddresses.IsNull() {
		data.MultipleIpToMacAddresses = types.BoolValue(value.Bool())
	} else {
		data.MultipleIpToMacAddresses = types.BoolNull()
	}
	if value := res.Get("response.0.isSupplicantBasedExtendedNodeOnboarding"); value.Exists() && !data.SupplicantBasedExtendedNodeOnboarding.IsNull() {
		data.SupplicantBasedExtendedNodeOnboarding = types.BoolValue(value.Bool())
	} else {
		data.SupplicantBasedExtendedNodeOnboarding = types.BoolNull()
	}
	if value := res.Get("response.0.autoGenerateVlanName"); value.Exists() && !data.AutoGenerateVlanName.IsNull() {
		data.AutoGenerateVlanName = types.BoolValue(value.Bool())
	} else {
		data.AutoGenerateVlanName = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AnycastGateway) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FabricId.IsNull() {
		return false
	}
	if !data.VirtualNetworkName.IsNull() {
		return false
	}
	if !data.IpPoolName.IsNull() {
		return false
	}
	if !data.TcpMssAdjustment.IsNull() {
		return false
	}
	if !data.VlanName.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.TrafficType.IsNull() {
		return false
	}
	if !data.PoolType.IsNull() {
		return false
	}
	if !data.SecurityGroupName.IsNull() {
		return false
	}
	if !data.CriticalPool.IsNull() {
		return false
	}
	if !data.L2FloodingEnabled.IsNull() {
		return false
	}
	if !data.WirelessPool.IsNull() {
		return false
	}
	if !data.IpDirectedBroadcast.IsNull() {
		return false
	}
	if !data.IntraSubnetRoutingEnabled.IsNull() {
		return false
	}
	if !data.MultipleIpToMacAddresses.IsNull() {
		return false
	}
	if !data.SupplicantBasedExtendedNodeOnboarding.IsNull() {
		return false
	}
	if !data.AutoGenerateVlanName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
