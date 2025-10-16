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
type IPPool struct {
	Id                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"`
	PoolType                 types.String `tfsdk:"pool_type"`
	AddressSpaceSubnet       types.String `tfsdk:"address_space_subnet"`
	AddressSpacePrefixLength types.Int64  `tfsdk:"address_space_prefix_length"`
	AddressSpaceGateway      types.String `tfsdk:"address_space_gateway"`
	AddressSpaceDhcpServers  types.Set    `tfsdk:"address_space_dhcp_servers"`
	AddressSpaceDnsServers   types.Set    `tfsdk:"address_space_dns_servers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPPool) getPath() string {
	return "/dna/intent/api/v1/ipam/globalIpAddressPools"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IPPool) toBody(ctx context.Context, state IPPool) string {
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
	if !data.AddressSpaceSubnet.IsNull() {
		body, _ = sjson.Set(body, "addressSpace.subnet", data.AddressSpaceSubnet.ValueString())
	}
	if !data.AddressSpacePrefixLength.IsNull() {
		body, _ = sjson.Set(body, "addressSpace.prefixLength", data.AddressSpacePrefixLength.ValueInt64())
	}
	if !data.AddressSpaceGateway.IsNull() {
		body, _ = sjson.Set(body, "addressSpace.gatewayIpAddress", data.AddressSpaceGateway.ValueString())
	}
	if !data.AddressSpaceDhcpServers.IsNull() {
		var values []string
		data.AddressSpaceDhcpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "addressSpace.dhcpServers", values)
	}
	if !data.AddressSpaceDnsServers.IsNull() {
		var values []string
		data.AddressSpaceDnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "addressSpace.dnsServers", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPPool) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("addressSpace.subnet"); value.Exists() {
		data.AddressSpaceSubnet = types.StringValue(value.String())
	} else {
		data.AddressSpaceSubnet = types.StringNull()
	}
	if value := res.Get("addressSpace.prefixLength"); value.Exists() {
		data.AddressSpacePrefixLength = types.Int64Value(value.Int())
	} else {
		data.AddressSpacePrefixLength = types.Int64Null()
	}
	if value := res.Get("addressSpace.gatewayIpAddress"); value.Exists() {
		data.AddressSpaceGateway = types.StringValue(value.String())
	} else {
		data.AddressSpaceGateway = types.StringNull()
	}
	if value := res.Get("addressSpace.dhcpServers"); value.Exists() && len(value.Array()) > 0 {
		data.AddressSpaceDhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("addressSpace.dnsServers"); value.Exists() && len(value.Array()) > 0 {
		data.AddressSpaceDnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDnsServers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *IPPool) updateFromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("addressSpace.subnet"); value.Exists() && !data.AddressSpaceSubnet.IsNull() {
		data.AddressSpaceSubnet = types.StringValue(value.String())
	} else {
		data.AddressSpaceSubnet = types.StringNull()
	}
	if value := res.Get("addressSpace.prefixLength"); value.Exists() && !data.AddressSpacePrefixLength.IsNull() {
		data.AddressSpacePrefixLength = types.Int64Value(value.Int())
	} else {
		data.AddressSpacePrefixLength = types.Int64Null()
	}
	if value := res.Get("addressSpace.gatewayIpAddress"); value.Exists() && !data.AddressSpaceGateway.IsNull() {
		data.AddressSpaceGateway = types.StringValue(value.String())
	} else {
		data.AddressSpaceGateway = types.StringNull()
	}
	if value := res.Get("addressSpace.dhcpServers"); value.Exists() && !data.AddressSpaceDhcpServers.IsNull() {
		data.AddressSpaceDhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("addressSpace.dnsServers"); value.Exists() && !data.AddressSpaceDnsServers.IsNull() {
		data.AddressSpaceDnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDnsServers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *IPPool) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.PoolType.IsNull() {
		return false
	}
	if !data.AddressSpaceSubnet.IsNull() {
		return false
	}
	if !data.AddressSpacePrefixLength.IsNull() {
		return false
	}
	if !data.AddressSpaceGateway.IsNull() {
		return false
	}
	if !data.AddressSpaceDhcpServers.IsNull() {
		return false
	}
	if !data.AddressSpaceDnsServers.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
