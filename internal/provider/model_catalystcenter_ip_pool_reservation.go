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
type IPPoolReservation struct {
	Id               types.String `tfsdk:"id"`
	SiteId           types.String `tfsdk:"site_id"`
	Name             types.String `tfsdk:"name"`
	Type             types.String `tfsdk:"type"`
	Ipv6AddressSpace types.Bool   `tfsdk:"ipv6_address_space"`
	Ipv4GlobalPool   types.String `tfsdk:"ipv4_global_pool"`
	Ipv4Prefix       types.Bool   `tfsdk:"ipv4_prefix"`
	Ipv4PrefixLength types.Int64  `tfsdk:"ipv4_prefix_length"`
	Ipv4Subnet       types.String `tfsdk:"ipv4_subnet"`
	Ipv4Gateway      types.String `tfsdk:"ipv4_gateway"`
	Ipv4DhcpServers  types.Set    `tfsdk:"ipv4_dhcp_servers"`
	Ipv4DnsServers   types.Set    `tfsdk:"ipv4_dns_servers"`
	Ipv6GlobalPool   types.String `tfsdk:"ipv6_global_pool"`
	Ipv6Prefix       types.Bool   `tfsdk:"ipv6_prefix"`
	Ipv6PrefixLength types.Int64  `tfsdk:"ipv6_prefix_length"`
	Ipv6Subnet       types.String `tfsdk:"ipv6_subnet"`
	Ipv6Gateway      types.String `tfsdk:"ipv6_gateway"`
	Ipv6DhcpServers  types.Set    `tfsdk:"ipv6_dhcp_servers"`
	Ipv6DnsServers   types.Set    `tfsdk:"ipv6_dns_servers"`
	Ipv4TotalHost    types.Int64  `tfsdk:"ipv4_total_host"`
	Ipv6TotalHost    types.Int64  `tfsdk:"ipv6_total_host"`
	SlaacSupport     types.Bool   `tfsdk:"slaac_support"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPPoolReservation) getPath() string {
	return "/dna/intent/api/v1/reserve-ip-subpool"
}

// End of section. //template:end getPath

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
	if !data.Type.IsNull() && put == false {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.Ipv6AddressSpace.IsNull() {
		body, _ = sjson.Set(body, "ipv6AddressSpace", data.Ipv6AddressSpace.ValueBool())
	}
	if !data.Ipv4GlobalPool.IsNull() && put == false {
		body, _ = sjson.Set(body, "ipv4GlobalPool", data.Ipv4GlobalPool.ValueString())
	}
	if !data.Ipv4Prefix.IsNull() && put == false {
		body, _ = sjson.Set(body, "ipv4Prefix", data.Ipv4Prefix.ValueBool())
	}
	if !data.Ipv4PrefixLength.IsNull() && put == false {
		body, _ = sjson.Set(body, "ipv4PrefixLength", data.Ipv4PrefixLength.ValueInt64())
	}
	if !data.Ipv4Subnet.IsNull() && put == false {
		body, _ = sjson.Set(body, "ipv4Subnet", data.Ipv4Subnet.ValueString())
	}
	if !data.Ipv4Gateway.IsNull() {
		body, _ = sjson.Set(body, "ipv4GateWay", data.Ipv4Gateway.ValueString())
	}
	if !data.Ipv4DhcpServers.IsNull() {
		var values []string
		data.Ipv4DhcpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipv4DhcpServers", values)
	}
	if !data.Ipv4DnsServers.IsNull() {
		var values []string
		data.Ipv4DnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipv4DnsServers", values)
	}
	if !data.Ipv6GlobalPool.IsNull() {
		body, _ = sjson.Set(body, "ipv6GlobalPool", data.Ipv6GlobalPool.ValueString())
	}
	if !data.Ipv6Prefix.IsNull() {
		body, _ = sjson.Set(body, "ipv6Prefix", data.Ipv6Prefix.ValueBool())
	}
	if !data.Ipv6PrefixLength.IsNull() {
		body, _ = sjson.Set(body, "ipv6PrefixLength", data.Ipv6PrefixLength.ValueInt64())
	}
	if !data.Ipv6Subnet.IsNull() {
		body, _ = sjson.Set(body, "ipv6Subnet", data.Ipv6Subnet.ValueString())
	}
	if !data.Ipv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "ipv6GateWay", data.Ipv6Gateway.ValueString())
	}
	if !data.Ipv6DhcpServers.IsNull() {
		var values []string
		data.Ipv6DhcpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipv6DhcpServers", values)
	}
	if !data.Ipv6DnsServers.IsNull() {
		var values []string
		data.Ipv6DnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipv6DnsServers", values)
	}
	if !data.Ipv4TotalHost.IsNull() {
		body, _ = sjson.Set(body, "ipv4TotalHost", data.Ipv4TotalHost.ValueInt64())
	}
	if !data.Ipv6TotalHost.IsNull() {
		body, _ = sjson.Set(body, "ipv6TotalHost", data.Ipv6TotalHost.ValueInt64())
	}
	if !data.SlaacSupport.IsNull() {
		body, _ = sjson.Set(body, "slaacSupport", data.SlaacSupport.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPPoolReservation) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("groupName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *IPPoolReservation) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("groupName"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *IPPoolReservation) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Ipv6AddressSpace.IsNull() {
		return false
	}
	if !data.Ipv4GlobalPool.IsNull() {
		return false
	}
	if !data.Ipv4Prefix.IsNull() {
		return false
	}
	if !data.Ipv4PrefixLength.IsNull() {
		return false
	}
	if !data.Ipv4Subnet.IsNull() {
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
	if !data.Ipv6GlobalPool.IsNull() {
		return false
	}
	if !data.Ipv6Prefix.IsNull() {
		return false
	}
	if !data.Ipv6PrefixLength.IsNull() {
		return false
	}
	if !data.Ipv6Subnet.IsNull() {
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
	if !data.Ipv4TotalHost.IsNull() {
		return false
	}
	if !data.Ipv6TotalHost.IsNull() {
		return false
	}
	if !data.SlaacSupport.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
