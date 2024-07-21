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
type FabricL3HandoffIPTransit struct {
	Id                             types.String `tfsdk:"id"`
	NetworkDeviceId                types.String `tfsdk:"network_device_id"`
	FabricId                       types.String `tfsdk:"fabric_id"`
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
func (data FabricL3HandoffIPTransit) getPath() string {
	return "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricL3HandoffIPTransit) toBody(ctx context.Context, state FabricL3HandoffIPTransit) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "0.networkDeviceId", data.NetworkDeviceId.ValueString())
	}
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.TransitNetworkId.IsNull() {
		body, _ = sjson.Set(body, "0.transitNetworkId", data.TransitNetworkId.ValueString())
	}
	if !data.InterfaceName.IsNull() {
		body, _ = sjson.Set(body, "0.interfaceName", data.InterfaceName.ValueString())
	}
	if !data.ExternalConnectivityIpPoolName.IsNull() {
		body, _ = sjson.Set(body, "0.externalConnectivityIpPoolName", data.ExternalConnectivityIpPoolName.ValueString())
	}
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "0.vlanId", data.VlanId.ValueInt64())
	}
	if !data.TcpMssAdjustment.IsNull() {
		body, _ = sjson.Set(body, "0.tcpMssAdjustment", data.TcpMssAdjustment.ValueInt64())
	}
	if !data.LocalIpAddress.IsNull() {
		body, _ = sjson.Set(body, "0.localIpAddress", data.LocalIpAddress.ValueString())
	}
	if !data.RemoteIpAddress.IsNull() {
		body, _ = sjson.Set(body, "0.remoteIpAddress", data.RemoteIpAddress.ValueString())
	}
	if !data.LocalIpv6Address.IsNull() {
		body, _ = sjson.Set(body, "0.localIpv6Address", data.LocalIpv6Address.ValueString())
	}
	if !data.RemoteIpv6Address.IsNull() {
		body, _ = sjson.Set(body, "0.remoteIpv6Address", data.RemoteIpv6Address.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricL3HandoffIPTransit) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.networkDeviceId"); value.Exists() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.transitNetworkId"); value.Exists() {
		data.TransitNetworkId = types.StringValue(value.String())
	} else {
		data.TransitNetworkId = types.StringNull()
	}
	if value := res.Get("response.0.interfaceName"); value.Exists() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("response.0.externalConnectivityIpPoolName"); value.Exists() {
		data.ExternalConnectivityIpPoolName = types.StringValue(value.String())
	} else {
		data.ExternalConnectivityIpPoolName = types.StringNull()
	}
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.vlanId"); value.Exists() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("response.0.tcpMssAdjustment"); value.Exists() {
		data.TcpMssAdjustment = types.Int64Value(value.Int())
	} else {
		data.TcpMssAdjustment = types.Int64Null()
	}
	if value := res.Get("response.0.localIpAddress"); value.Exists() {
		data.LocalIpAddress = types.StringValue(value.String())
	} else {
		data.LocalIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.remoteIpAddress"); value.Exists() {
		data.RemoteIpAddress = types.StringValue(value.String())
	} else {
		data.RemoteIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.localIpv6Address"); value.Exists() {
		data.LocalIpv6Address = types.StringValue(value.String())
	} else {
		data.LocalIpv6Address = types.StringNull()
	}
	if value := res.Get("response.0.remoteIpv6Address"); value.Exists() {
		data.RemoteIpv6Address = types.StringValue(value.String())
	} else {
		data.RemoteIpv6Address = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricL3HandoffIPTransit) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.networkDeviceId"); value.Exists() && !data.NetworkDeviceId.IsNull() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.transitNetworkId"); value.Exists() && !data.TransitNetworkId.IsNull() {
		data.TransitNetworkId = types.StringValue(value.String())
	} else {
		data.TransitNetworkId = types.StringNull()
	}
	if value := res.Get("response.0.interfaceName"); value.Exists() && !data.InterfaceName.IsNull() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("response.0.externalConnectivityIpPoolName"); value.Exists() && !data.ExternalConnectivityIpPoolName.IsNull() {
		data.ExternalConnectivityIpPoolName = types.StringValue(value.String())
	} else {
		data.ExternalConnectivityIpPoolName = types.StringNull()
	}
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("response.0.tcpMssAdjustment"); value.Exists() && !data.TcpMssAdjustment.IsNull() {
		data.TcpMssAdjustment = types.Int64Value(value.Int())
	} else {
		data.TcpMssAdjustment = types.Int64Null()
	}
	if value := res.Get("response.0.localIpAddress"); value.Exists() && !data.LocalIpAddress.IsNull() {
		data.LocalIpAddress = types.StringValue(value.String())
	} else {
		data.LocalIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.remoteIpAddress"); value.Exists() && !data.RemoteIpAddress.IsNull() {
		data.RemoteIpAddress = types.StringValue(value.String())
	} else {
		data.RemoteIpAddress = types.StringNull()
	}
	if value := res.Get("response.0.localIpv6Address"); value.Exists() && !data.LocalIpv6Address.IsNull() {
		data.LocalIpv6Address = types.StringValue(value.String())
	} else {
		data.LocalIpv6Address = types.StringNull()
	}
	if value := res.Get("response.0.remoteIpv6Address"); value.Exists() && !data.RemoteIpv6Address.IsNull() {
		data.RemoteIpv6Address = types.StringValue(value.String())
	} else {
		data.RemoteIpv6Address = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricL3HandoffIPTransit) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.TransitNetworkId.IsNull() {
		return false
	}
	if !data.InterfaceName.IsNull() {
		return false
	}
	if !data.ExternalConnectivityIpPoolName.IsNull() {
		return false
	}
	if !data.VirtualNetworkName.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.TcpMssAdjustment.IsNull() {
		return false
	}
	if !data.LocalIpAddress.IsNull() {
		return false
	}
	if !data.RemoteIpAddress.IsNull() {
		return false
	}
	if !data.LocalIpv6Address.IsNull() {
		return false
	}
	if !data.RemoteIpv6Address.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
