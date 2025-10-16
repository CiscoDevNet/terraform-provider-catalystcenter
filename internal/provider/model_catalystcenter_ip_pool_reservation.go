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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type IPPoolReservation struct {
	Id                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	PoolType                     types.String `tfsdk:"pool_type"`
	SiteId                       types.String `tfsdk:"site_id"`
	Ipv4Subnet                   types.String `tfsdk:"ipv4_subnet"`
	Ipv4PrefixLength             types.Int64  `tfsdk:"ipv4_prefix_length"`
	Ipv4Gateway                  types.String `tfsdk:"ipv4_gateway"`
	Ipv4DhcpServers              types.Set    `tfsdk:"ipv4_dhcp_servers"`
	Ipv4DnsServers               types.Set    `tfsdk:"ipv4_dns_servers"`
	Ipv4TotalAddresses           types.String `tfsdk:"ipv4_total_addresses"`
	Ipv4UnassignableAddresses    types.String `tfsdk:"ipv4_unassignable_addresses"`
	Ipv4AssignedAddresses        types.String `tfsdk:"ipv4_assigned_addresses"`
	Ipv4DefaultAssignedAddresses types.String `tfsdk:"ipv4_default_assigned_addresses"`
	Ipv4GlobalPoolId             types.String `tfsdk:"ipv4_global_pool_id"`
	Ipv6Subnet                   types.String `tfsdk:"ipv6_subnet"`
	Ipv6PrefixLength             types.Int64  `tfsdk:"ipv6_prefix_length"`
	Ipv6Gateway                  types.String `tfsdk:"ipv6_gateway"`
	Ipv6DhcpServers              types.Set    `tfsdk:"ipv6_dhcp_servers"`
	Ipv6DnsServers               types.Set    `tfsdk:"ipv6_dns_servers"`
	Ipv6TotalAddresses           types.String `tfsdk:"ipv6_total_addresses"`
	Ipv6UnassignableAddresses    types.String `tfsdk:"ipv6_unassignable_addresses"`
	Ipv6AssignedAddresses        types.String `tfsdk:"ipv6_assigned_addresses"`
	Ipv6DefaultAssignedAddresses types.String `tfsdk:"ipv6_default_assigned_addresses"`
	Ipv6SlaacSupport             types.Bool   `tfsdk:"ipv6_slaac_support"`
	Ipv6GlobalPoolId             types.String `tfsdk:"ipv6_global_pool_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPPoolReservation) getPath() string {
	return "/dna/intent/api/v1/ipam/siteIpAddressPools"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IPPoolReservation) toBody(ctx context.Context, state IPPoolReservation) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PoolType.IsNull() {
		body, _ = sjson.Set(body, "poolType", data.PoolType.ValueString())
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if !data.Ipv4Subnet.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.subnet", data.Ipv4Subnet.ValueString())
	}
	if !data.Ipv4PrefixLength.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.prefixLength", data.Ipv4PrefixLength.ValueInt64())
	}
	if !data.Ipv4Gateway.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.gatewayIpAddress", data.Ipv4Gateway.ValueString())
	}
	if !data.Ipv4DhcpServers.IsNull() {
		var values []string
		data.Ipv4DhcpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipV4AddressSpace.dhcpServers", values)
	}
	if !data.Ipv4DnsServers.IsNull() {
		var values []string
		data.Ipv4DnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipV4AddressSpace.dnsServers", values)
	}
	if !data.Ipv4TotalAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.totalAddresses", data.Ipv4TotalAddresses.ValueString())
	}
	if !data.Ipv4UnassignableAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.unassignableAddresses", data.Ipv4UnassignableAddresses.ValueString())
	}
	if !data.Ipv4AssignedAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.assignedAddresses", data.Ipv4AssignedAddresses.ValueString())
	}
	if !data.Ipv4DefaultAssignedAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.defaultAssignedAddresses", data.Ipv4DefaultAssignedAddresses.ValueString())
	}
	if !data.Ipv4GlobalPoolId.IsNull() {
		body, _ = sjson.Set(body, "ipV4AddressSpace.globalPoolId", data.Ipv4GlobalPoolId.ValueString())
	}
	if !data.Ipv6Subnet.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.subnet", data.Ipv6Subnet.ValueString())
	}
	if !data.Ipv6PrefixLength.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.prefixLength", data.Ipv6PrefixLength.ValueInt64())
	}
	if !data.Ipv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.gatewayIpAddress", data.Ipv6Gateway.ValueString())
	}
	if !data.Ipv6DhcpServers.IsNull() {
		var values []string
		data.Ipv6DhcpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipV6AddressSpace.dhcpServers", values)
	}
	if !data.Ipv6DnsServers.IsNull() {
		var values []string
		data.Ipv6DnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipV6AddressSpace.dnsServers", values)
	}
	if !data.Ipv6TotalAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.totalAddresses", data.Ipv6TotalAddresses.ValueString())
	}
	if !data.Ipv6UnassignableAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.unassignableAddresses", data.Ipv6UnassignableAddresses.ValueString())
	}
	if !data.Ipv6AssignedAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.assignedAddresses", data.Ipv6AssignedAddresses.ValueString())
	}
	if !data.Ipv6DefaultAssignedAddresses.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.defaultAssignedAddresses", data.Ipv6DefaultAssignedAddresses.ValueString())
	}
	if !data.Ipv6SlaacSupport.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.slaacSupport", data.Ipv6SlaacSupport.ValueBool())
	}
	if !data.Ipv6GlobalPoolId.IsNull() {
		body, _ = sjson.Set(body, "ipV6AddressSpace.globalPoolId", data.Ipv6GlobalPoolId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPPoolReservation) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("poolType"); value.Exists() {
		data.PoolType = types.StringValue(value.String())
	} else {
		data.PoolType = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.subnet"); value.Exists() {
		data.Ipv4Subnet = types.StringValue(value.String())
	} else {
		data.Ipv4Subnet = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.prefixLength"); value.Exists() {
		data.Ipv4PrefixLength = types.Int64Value(value.Int())
	} else {
		data.Ipv4PrefixLength = types.Int64Null()
	}
	if value := res.Get("ipV4AddressSpace.gatewayIpAddress"); value.Exists() {
		data.Ipv4Gateway = types.StringValue(value.String())
	} else {
		data.Ipv4Gateway = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.dhcpServers"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4DhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv4DhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV4AddressSpace.dnsServers"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv4DnsServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV4AddressSpace.totalAddresses"); value.Exists() {
		data.Ipv4TotalAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4TotalAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.unassignableAddresses"); value.Exists() {
		data.Ipv4UnassignableAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4UnassignableAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.assignedAddresses"); value.Exists() {
		data.Ipv4AssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4AssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.defaultAssignedAddresses"); value.Exists() {
		data.Ipv4DefaultAssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4DefaultAssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.globalPoolId"); value.Exists() {
		data.Ipv4GlobalPoolId = types.StringValue(value.String())
	} else {
		data.Ipv4GlobalPoolId = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.subnet"); value.Exists() {
		data.Ipv6Subnet = types.StringValue(value.String())
	} else {
		data.Ipv6Subnet = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.prefixLength"); value.Exists() {
		data.Ipv6PrefixLength = types.Int64Value(value.Int())
	} else {
		data.Ipv6PrefixLength = types.Int64Null()
	}
	if value := res.Get("ipV6AddressSpace.gatewayIpAddress"); value.Exists() {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.dhcpServers"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6DhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv6DhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV6AddressSpace.dnsServers"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv6DnsServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV6AddressSpace.totalAddresses"); value.Exists() {
		data.Ipv6TotalAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6TotalAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.unassignableAddresses"); value.Exists() {
		data.Ipv6UnassignableAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6UnassignableAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.assignedAddresses"); value.Exists() {
		data.Ipv6AssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6AssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.defaultAssignedAddresses"); value.Exists() {
		data.Ipv6DefaultAssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6DefaultAssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.slaacSupport"); value.Exists() {
		data.Ipv6SlaacSupport = types.BoolValue(value.Bool())
	} else {
		data.Ipv6SlaacSupport = types.BoolNull()
	}
	if value := res.Get("ipV6AddressSpace.globalPoolId"); value.Exists() {
		data.Ipv6GlobalPoolId = types.StringValue(value.String())
	} else {
		data.Ipv6GlobalPoolId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *IPPoolReservation) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("poolType"); value.Exists() && !data.PoolType.IsNull() {
		data.PoolType = types.StringValue(value.String())
	} else {
		data.PoolType = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.subnet"); value.Exists() && !data.Ipv4Subnet.IsNull() {
		data.Ipv4Subnet = types.StringValue(value.String())
	} else {
		data.Ipv4Subnet = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.prefixLength"); value.Exists() && !data.Ipv4PrefixLength.IsNull() {
		data.Ipv4PrefixLength = types.Int64Value(value.Int())
	} else {
		data.Ipv4PrefixLength = types.Int64Null()
	}
	if value := res.Get("ipV4AddressSpace.gatewayIpAddress"); value.Exists() && !data.Ipv4Gateway.IsNull() {
		data.Ipv4Gateway = types.StringValue(value.String())
	} else {
		data.Ipv4Gateway = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.dhcpServers"); value.Exists() && !data.Ipv4DhcpServers.IsNull() {
		data.Ipv4DhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv4DhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV4AddressSpace.dnsServers"); value.Exists() && !data.Ipv4DnsServers.IsNull() {
		data.Ipv4DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv4DnsServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV4AddressSpace.totalAddresses"); value.Exists() && !data.Ipv4TotalAddresses.IsNull() {
		data.Ipv4TotalAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4TotalAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.unassignableAddresses"); value.Exists() && !data.Ipv4UnassignableAddresses.IsNull() {
		data.Ipv4UnassignableAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4UnassignableAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.assignedAddresses"); value.Exists() && !data.Ipv4AssignedAddresses.IsNull() {
		data.Ipv4AssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4AssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.defaultAssignedAddresses"); value.Exists() && !data.Ipv4DefaultAssignedAddresses.IsNull() {
		data.Ipv4DefaultAssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv4DefaultAssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV4AddressSpace.globalPoolId"); value.Exists() && !data.Ipv4GlobalPoolId.IsNull() {
		data.Ipv4GlobalPoolId = types.StringValue(value.String())
	} else {
		data.Ipv4GlobalPoolId = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.subnet"); value.Exists() && !data.Ipv6Subnet.IsNull() {
		data.Ipv6Subnet = types.StringValue(value.String())
	} else {
		data.Ipv6Subnet = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.prefixLength"); value.Exists() && !data.Ipv6PrefixLength.IsNull() {
		data.Ipv6PrefixLength = types.Int64Value(value.Int())
	} else {
		data.Ipv6PrefixLength = types.Int64Null()
	}
	if value := res.Get("ipV6AddressSpace.gatewayIpAddress"); value.Exists() && !data.Ipv6Gateway.IsNull() {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.dhcpServers"); value.Exists() && !data.Ipv6DhcpServers.IsNull() {
		data.Ipv6DhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv6DhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV6AddressSpace.dnsServers"); value.Exists() && !data.Ipv6DnsServers.IsNull() {
		data.Ipv6DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv6DnsServers = types.SetNull(types.StringType)
	}
	if value := res.Get("ipV6AddressSpace.totalAddresses"); value.Exists() && !data.Ipv6TotalAddresses.IsNull() {
		data.Ipv6TotalAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6TotalAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.unassignableAddresses"); value.Exists() && !data.Ipv6UnassignableAddresses.IsNull() {
		data.Ipv6UnassignableAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6UnassignableAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.assignedAddresses"); value.Exists() && !data.Ipv6AssignedAddresses.IsNull() {
		data.Ipv6AssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6AssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.defaultAssignedAddresses"); value.Exists() && !data.Ipv6DefaultAssignedAddresses.IsNull() {
		data.Ipv6DefaultAssignedAddresses = types.StringValue(value.String())
	} else {
		data.Ipv6DefaultAssignedAddresses = types.StringNull()
	}
	if value := res.Get("ipV6AddressSpace.slaacSupport"); value.Exists() && !data.Ipv6SlaacSupport.IsNull() {
		data.Ipv6SlaacSupport = types.BoolValue(value.Bool())
	} else {
		data.Ipv6SlaacSupport = types.BoolNull()
	}
	if value := res.Get("ipV6AddressSpace.globalPoolId"); value.Exists() && !data.Ipv6GlobalPoolId.IsNull() {
		data.Ipv6GlobalPoolId = types.StringValue(value.String())
	} else {
		data.Ipv6GlobalPoolId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *IPPoolReservation) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.PoolType.IsNull() {
		return false
	}
	if !data.SiteId.IsNull() {
		return false
	}
	if !data.Ipv4Subnet.IsNull() {
		return false
	}
	if !data.Ipv4PrefixLength.IsNull() {
		return false
	}
	if !data.Ipv4Gateway.IsNull() {
		return false
	}
	if !data.Ipv4DhcpServers.IsNull() {
		return false
	}
	if !data.Ipv4DnsServers.IsNull() {
		return false
	}
	if !data.Ipv4TotalAddresses.IsNull() {
		return false
	}
	if !data.Ipv4UnassignableAddresses.IsNull() {
		return false
	}
	if !data.Ipv4AssignedAddresses.IsNull() {
		return false
	}
	if !data.Ipv4DefaultAssignedAddresses.IsNull() {
		return false
	}
	if !data.Ipv4GlobalPoolId.IsNull() {
		return false
	}
	if !data.Ipv6Subnet.IsNull() {
		return false
	}
	if !data.Ipv6PrefixLength.IsNull() {
		return false
	}
	if !data.Ipv6Gateway.IsNull() {
		return false
	}
	if !data.Ipv6DhcpServers.IsNull() {
		return false
	}
	if !data.Ipv6DnsServers.IsNull() {
		return false
	}
	if !data.Ipv6TotalAddresses.IsNull() {
		return false
	}
	if !data.Ipv6UnassignableAddresses.IsNull() {
		return false
	}
	if !data.Ipv6AssignedAddresses.IsNull() {
		return false
	}
	if !data.Ipv6DefaultAssignedAddresses.IsNull() {
		return false
	}
	if !data.Ipv6SlaacSupport.IsNull() {
		return false
	}
	if !data.Ipv6GlobalPoolId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
