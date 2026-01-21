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

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

func (data IPPool) getFallbackPath() string {
	return "/dna/intent/api/v1/global-pool"
}

// End of section. //template:end getFallbackPath

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

func (data *IPPool) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else if value := res.Get("ipPoolName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("poolType"); value.Exists() {
		data.PoolType = types.StringValue(value.String())
	} else if value := res.Get("ipPoolType"); value.Exists() {
		// Normalize type case: "generic" -> "Generic", "tunnel" -> "Tunnel"
		poolType := value.String()
		if poolType != "" {
			poolType = strings.ToUpper(poolType[:1]) + strings.ToLower(poolType[1:])
		}
		data.PoolType = types.StringValue(poolType)
	} else {
		data.PoolType = types.StringNull()
	}
	// Handle subnet - parse from CIDR if needed
	if value := res.Get("addressSpace.subnet"); value.Exists() {
		data.AddressSpaceSubnet = types.StringValue(value.String())
	} else if value := res.Get("ipPoolCidr"); value.Exists() {
		// Parse CIDR notation: "192.168.0.0/16" -> "192.168.0.0"
		cidr := value.String()
		if parts := strings.Split(cidr, "/"); len(parts) >= 1 {
			data.AddressSpaceSubnet = types.StringValue(parts[0])
		} else {
			data.AddressSpaceSubnet = types.StringValue(cidr)
		}
	} else {
		data.AddressSpaceSubnet = types.StringNull()
	}
	// Handle prefix length - parse from CIDR if needed
	if value := res.Get("addressSpace.prefixLength"); value.Exists() {
		data.AddressSpacePrefixLength = types.Int64Value(value.Int())
	} else if value := res.Get("ipPoolCidr"); value.Exists() {
		// Parse CIDR notation: "192.168.0.0/16" -> 16
		cidr := value.String()
		if parts := strings.Split(cidr, "/"); len(parts) == 2 {
			if prefix, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
				data.AddressSpacePrefixLength = types.Int64Value(prefix)
			} else {
				data.AddressSpacePrefixLength = types.Int64Null()
			}
		} else {
			data.AddressSpacePrefixLength = types.Int64Null()
		}
	} else {
		data.AddressSpacePrefixLength = types.Int64Null()
	}
	// Handle gateway - extract first element from array if needed
	if value := res.Get("addressSpace.gatewayIpAddress"); value.Exists() {
		data.AddressSpaceGateway = types.StringValue(value.String())
	} else if value := res.Get("gateways"); value.Exists() {
		// Extract first gateway from array
		if value.IsArray() && len(value.Array()) > 0 {
			data.AddressSpaceGateway = types.StringValue(value.Array()[0].String())
		} else {
			data.AddressSpaceGateway = types.StringNull()
		}
	} else {
		data.AddressSpaceGateway = types.StringNull()
	}
	if value := res.Get("addressSpace.dhcpServers"); value.Exists() && len(value.Array()) > 0 {
		data.AddressSpaceDhcpServers = helpers.GetStringSet(value.Array())
	} else if value := res.Get("dhcpServerIps"); value.Exists() && len(value.Array()) > 0 {
		data.AddressSpaceDhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("addressSpace.dnsServers"); value.Exists() && len(value.Array()) > 0 {
		data.AddressSpaceDnsServers = helpers.GetStringSet(value.Array())
	} else if value := res.Get("dnsServerIps"); value.Exists() && len(value.Array()) > 0 {
		data.AddressSpaceDnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDnsServers = types.SetNull(types.StringType)
	}
}

func (data *IPPool) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else if value := res.Get("ipPoolName"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("poolType"); value.Exists() && !data.PoolType.IsNull() {
		data.PoolType = types.StringValue(value.String())
	} else if value := res.Get("ipPoolType"); value.Exists() && !data.PoolType.IsNull() {
		// Normalize type case: "generic" -> "Generic", "tunnel" -> "Tunnel"
		poolType := value.String()
		if poolType != "" {
			poolType = strings.ToUpper(poolType[:1]) + strings.ToLower(poolType[1:])
		}
		data.PoolType = types.StringValue(poolType)
	} else {
		data.PoolType = types.StringNull()
	}
	// Handle subnet - parse from CIDR if needed
	if value := res.Get("addressSpace.subnet"); value.Exists() && !data.AddressSpaceSubnet.IsNull() {
		data.AddressSpaceSubnet = types.StringValue(value.String())
	} else if value := res.Get("ipPoolCidr"); value.Exists() && !data.AddressSpaceSubnet.IsNull() {
		// Parse CIDR notation: "192.168.0.0/16" -> "192.168.0.0"
		cidr := value.String()
		if parts := strings.Split(cidr, "/"); len(parts) >= 1 {
			data.AddressSpaceSubnet = types.StringValue(parts[0])
		} else {
			data.AddressSpaceSubnet = types.StringValue(cidr)
		}
	} else {
		data.AddressSpaceSubnet = types.StringNull()
	}
	// Handle prefix length - parse from CIDR if needed
	if value := res.Get("addressSpace.prefixLength"); value.Exists() && !data.AddressSpacePrefixLength.IsNull() {
		data.AddressSpacePrefixLength = types.Int64Value(value.Int())
	} else if value := res.Get("ipPoolCidr"); value.Exists() && !data.AddressSpacePrefixLength.IsNull() {
		// Parse CIDR notation: "192.168.0.0/16" -> 16
		cidr := value.String()
		if parts := strings.Split(cidr, "/"); len(parts) == 2 {
			if prefix, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
				data.AddressSpacePrefixLength = types.Int64Value(prefix)
			} else {
				data.AddressSpacePrefixLength = types.Int64Null()
			}
		} else {
			data.AddressSpacePrefixLength = types.Int64Null()
		}
	} else {
		data.AddressSpacePrefixLength = types.Int64Null()
	}
	// Handle gateway - extract first element from array if needed
	if value := res.Get("addressSpace.gatewayIpAddress"); value.Exists() && !data.AddressSpaceGateway.IsNull() {
		data.AddressSpaceGateway = types.StringValue(value.String())
	} else if value := res.Get("gateways"); value.Exists() && !data.AddressSpaceGateway.IsNull() {
		// Extract first gateway from array
		if value.IsArray() && len(value.Array()) > 0 {
			data.AddressSpaceGateway = types.StringValue(value.Array()[0].String())
		} else {
			data.AddressSpaceGateway = types.StringNull()
		}
	} else {
		data.AddressSpaceGateway = types.StringNull()
	}
	if value := res.Get("addressSpace.dhcpServers"); value.Exists() && !data.AddressSpaceDhcpServers.IsNull() {
		data.AddressSpaceDhcpServers = helpers.GetStringSet(value.Array())
	} else if value := res.Get("dhcpServerIps"); value.Exists() && !data.AddressSpaceDhcpServers.IsNull() {
		data.AddressSpaceDhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("addressSpace.dnsServers"); value.Exists() && !data.AddressSpaceDnsServers.IsNull() {
		data.AddressSpaceDnsServers = helpers.GetStringSet(value.Array())
	} else if value := res.Get("dnsServerIps"); value.Exists() && !data.AddressSpaceDnsServers.IsNull() {
		data.AddressSpaceDnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.AddressSpaceDnsServers = types.SetNull(types.StringType)
	}
}

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
