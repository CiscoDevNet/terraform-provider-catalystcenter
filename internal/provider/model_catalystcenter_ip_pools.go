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
)

// End of section. //template:end imports

type IPPools struct {
	Pools []IPPoolsPools `tfsdk:"pools"`
}

type IPPoolsPools struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	PoolType         types.String `tfsdk:"pool_type"`
	GatewayIpAddress types.String `tfsdk:"gateway_ip_address"`
	Subnet           types.String `tfsdk:"subnet"`
	PrefixLength     types.String `tfsdk:"prefix_length"`
	DhcpServers      types.Set    `tfsdk:"dhcp_servers"`
	DnsServers       types.Set    `tfsdk:"dns_servers"`
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPPools) getPath() string {
	return "/dna/intent/api/v1/ipam/globalIpAddressPools"
}

// End of section. //template:end getPath

func (data *IPPools) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Pools = make([]IPPoolsPools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IPPoolsPools{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("poolType"); cValue.Exists() {
				item.PoolType = types.StringValue(cValue.String())
			} else {
				item.PoolType = types.StringValue("generic")
			}
			if cValue := v.Get("addressSpace.gatewayIpAddress"); cValue.Exists() {
				item.GatewayIpAddress = types.StringValue(cValue.String())
			} else {
				item.GatewayIpAddress = types.StringNull()
			}
			if cValue := v.Get("addressSpace.subnet"); cValue.Exists() {
				item.Subnet = types.StringValue(cValue.String())
			} else {
				item.Subnet = types.StringNull()
			}
			if cValue := v.Get("addressSpace.prefixLength"); cValue.Exists() {
				item.PrefixLength = types.StringValue(cValue.String())
			} else {
				item.PrefixLength = types.StringNull()
			}
			if cValue := v.Get("addressSpace.dhcpServers"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.DhcpServers = helpers.GetStringSet(cValue.Array())
			} else {
				item.DhcpServers = types.SetNull(types.StringType)
			}
			if cValue := v.Get("addressSpace.dnsServers"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.DnsServers = helpers.GetStringSet(cValue.Array())
			} else {
				item.DnsServers = types.SetNull(types.StringType)
			}
			data.Pools = append(data.Pools, item)
			return true
		})
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *IPPools) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Pools) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
