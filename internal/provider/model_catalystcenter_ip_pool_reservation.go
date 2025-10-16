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
	"strings"

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

func (data *IPPoolReservation) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.groupName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	// Added support for Dual Stack Pool Reservations
	if pools := res.Get("response.0.ipPools"); pools.IsArray() {
		ipPools := pools.Array()
		for _, pool := range ipPools {
			if pool.Get("ipv6").Bool() { // Check if the pool is IPv6
				if value := pool.Get("gateways.0"); value.Exists() {
					data.Ipv6Gateway = types.StringValue(value.String())
				} else {
					data.Ipv6Gateway = types.StringNull()
				}
				if value := pool.Get("dhcpServerIps"); value.Exists() && len(value.Array()) > 0 {
					data.Ipv6DhcpServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv6DhcpServers = types.SetNull(types.StringType)
				}
				if value := pool.Get("dnsServerIps"); value.Exists() && len(value.Array()) > 0 {
					data.Ipv6DnsServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv6DnsServers = types.SetNull(types.StringType)
				}
				// Extract and assign Ipv6Subnet and Ipv6Prefix from ipPoolCidr
				if cidr := pool.Get("ipPoolCidr"); cidr.Exists() && cidr.String() != "" {
					cidrParts := strings.Split(cidr.String(), "/")
					if len(cidrParts) == 2 {
						data.Ipv6Subnet = types.StringValue(cidrParts[0])
						if prefixLength, err := strconv.ParseInt(cidrParts[1], 10, 64); err == nil {
							data.Ipv6PrefixLength = types.Int64Value(prefixLength)
						} else {
							data.Ipv6PrefixLength = types.Int64Null()
						}
					} else {
						data.Ipv6Subnet = types.StringNull()
						data.Ipv6PrefixLength = types.Int64Null()
					}
				} else {
					data.Ipv6Subnet = types.StringNull()
					data.Ipv6PrefixLength = types.Int64Null()
				}
			} else { // Otherwise it's IPv4
				if value := pool.Get("gateways.0"); value.Exists() {
					data.Ipv4Gateway = types.StringValue(value.String())
				} else {
					data.Ipv4Gateway = types.StringNull()
				}
				if value := pool.Get("dhcpServerIps"); value.Exists() && len(value.Array()) > 0 {
					data.Ipv4DhcpServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv4DhcpServers = types.SetNull(types.StringType)
				}
				if value := pool.Get("dnsServerIps"); value.Exists() && len(value.Array()) > 0 {
					data.Ipv4DnsServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv4DnsServers = types.SetNull(types.StringType)
				}
				// Extract and assign Ipv4Subnet and Ipv4Prefix from ipPoolCidr
				if cidr := pool.Get("ipPoolCidr"); cidr.Exists() && cidr.String() != "" {
					cidrParts := strings.Split(cidr.String(), "/")
					if len(cidrParts) == 2 {
						data.Ipv4Subnet = types.StringValue(cidrParts[0])
						if prefixLength, err := strconv.ParseInt(cidrParts[1], 10, 64); err == nil {
							data.Ipv4PrefixLength = types.Int64Value(prefixLength)
						} else {
							data.Ipv4PrefixLength = types.Int64Null()
						}
					} else {
						data.Ipv4Subnet = types.StringNull()
						data.Ipv4PrefixLength = types.Int64Null()
					}
				} else {
					data.Ipv4Subnet = types.StringNull()
					data.Ipv4PrefixLength = types.Int64Null()
				}
			}
		}
	}
}

func (data *IPPoolReservation) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.groupName"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	// Added support for Dual Stack Pool Reservations
	if pools := res.Get("response.0.ipPools"); pools.IsArray() {
		ipPools := pools.Array()
		for _, pool := range ipPools {
			if pool.Get("ipv6").Bool() { // Check if the pool is IPv6
				if value := pool.Get("gateways.0"); value.Exists() && !data.Ipv6Gateway.IsNull() {
					data.Ipv6Gateway = types.StringValue(value.String())
				} else {
					data.Ipv6Gateway = types.StringNull()
				}
				if value := pool.Get("dhcpServerIps"); value.Exists() && !data.Ipv6DhcpServers.IsNull() {
					data.Ipv6DhcpServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv6DhcpServers = types.SetNull(types.StringType)
				}
				if value := pool.Get("dnsServerIps"); value.Exists() && !data.Ipv6DnsServers.IsNull() {
					data.Ipv6DnsServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv6DnsServers = types.SetNull(types.StringType)
				}
				// Extract and assign Ipv6Subnet and Ipv6Prefix from ipPoolCidr
				if cidr := pool.Get("ipPoolCidr"); cidr.Exists() && cidr.String() != "" {
					cidrParts := strings.Split(cidr.String(), "/")
					if len(cidrParts) == 2 {
						if data.Ipv6Subnet.IsNull() {
							data.Ipv6Subnet = types.StringValue(cidrParts[0])
						}
						if prefixLength, err := strconv.ParseInt(cidrParts[1], 10, 64); err == nil && !data.Ipv6PrefixLength.IsNull() {
							data.Ipv6PrefixLength = types.Int64Value(prefixLength)
						} else {
							data.Ipv6PrefixLength = types.Int64Null()
						}
					} else {
						data.Ipv6Subnet = types.StringNull()
						data.Ipv6PrefixLength = types.Int64Null()
					}
				} else {
					data.Ipv6Subnet = types.StringNull()
					data.Ipv6PrefixLength = types.Int64Null()
				}
			} else { // Otherwise it's IPv4
				if value := pool.Get("gateways.0"); value.Exists() && !data.Ipv4Gateway.IsNull() {
					data.Ipv4Gateway = types.StringValue(value.String())
				} else {
					data.Ipv4Gateway = types.StringNull()
				}
				if value := pool.Get("dhcpServerIps"); value.Exists() && !data.Ipv4DhcpServers.IsNull() {
					data.Ipv4DhcpServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv4DhcpServers = types.SetNull(types.StringType)
				}
				if value := pool.Get("dnsServerIps"); value.Exists() && !data.Ipv4DnsServers.IsNull() {
					data.Ipv4DnsServers = helpers.GetStringSet(value.Array())
				} else {
					data.Ipv4DnsServers = types.SetNull(types.StringType)
				}
				// Extract and assign Ipv4Subnet and Ipv4Prefix from ipPoolCidr
				if cidr := pool.Get("ipPoolCidr"); cidr.Exists() && cidr.String() != "" {
					cidrParts := strings.Split(cidr.String(), "/")
					if len(cidrParts) == 2 {
						if data.Ipv4Subnet.IsNull() {
							data.Ipv4Subnet = types.StringValue(cidrParts[0])
						}
						if prefixLength, err := strconv.ParseInt(cidrParts[1], 10, 64); err == nil && !data.Ipv4PrefixLength.IsNull() {
							data.Ipv4PrefixLength = types.Int64Value(prefixLength)
						} else {
							data.Ipv4PrefixLength = types.Int64Null()
						}
					} else {
						data.Ipv4Subnet = types.StringNull()
						data.Ipv4PrefixLength = types.Int64Null()
					}
				} else {
					data.Ipv4Subnet = types.StringNull()
					data.Ipv4PrefixLength = types.Int64Null()
				}
			}
		}
	}
}

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
