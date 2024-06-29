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
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	IpAddressSpace types.String `tfsdk:"ip_address_space"`
	Type           types.String `tfsdk:"type"`
	IpSubnet       types.String `tfsdk:"ip_subnet"`
	Gateway        types.Set    `tfsdk:"gateway"`
	DhcpServerIps  types.Set    `tfsdk:"dhcp_server_ips"`
	DnsServerIps   types.Set    `tfsdk:"dns_server_ips"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPPool) getPath() string {
	return "/api/v2/ippool"
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
		body, _ = sjson.Set(body, "ipPoolName", data.Name.ValueString())
	}
	if !data.IpAddressSpace.IsNull() {
		body, _ = sjson.Set(body, "IpAddressSpace", data.IpAddressSpace.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.IpSubnet.IsNull() {
		body, _ = sjson.Set(body, "ipPoolCidr", data.IpSubnet.ValueString())
	}
	if !data.Gateway.IsNull() {
		var values []string
		data.Gateway.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "gateways", values)
	}
	if !data.DhcpServerIps.IsNull() {
		var values []string
		data.DhcpServerIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dhcpServerIps", values)
	}
	if !data.DnsServerIps.IsNull() {
		var values []string
		data.DnsServerIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dnsServerIps", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPPool) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.ipPoolName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("IpAddressSpace"); value.Exists() {
		data.IpAddressSpace = types.StringValue(value.String())
	} else {
		data.IpAddressSpace = types.StringValue("IPv4")
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringValue("generic")
	}
	if value := res.Get("response.ipPoolCidr"); value.Exists() {
		data.IpSubnet = types.StringValue(value.String())
	} else {
		data.IpSubnet = types.StringNull()
	}
	if value := res.Get("response.gateways"); value.Exists() && len(value.Array()) > 0 {
		data.Gateway = helpers.GetStringSet(value.Array())
	} else {
		data.Gateway = types.SetNull(types.StringType)
	}
	if value := res.Get("response.dhcpServerIps"); value.Exists() && len(value.Array()) > 0 {
		data.DhcpServerIps = helpers.GetStringSet(value.Array())
	} else {
		data.DhcpServerIps = types.SetNull(types.StringType)
	}
	if value := res.Get("response.dnsServerIps"); value.Exists() && len(value.Array()) > 0 {
		data.DnsServerIps = helpers.GetStringSet(value.Array())
	} else {
		data.DnsServerIps = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *IPPool) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.ipPoolName"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("IpAddressSpace"); value.Exists() && !data.IpAddressSpace.IsNull() {
		data.IpAddressSpace = types.StringValue(value.String())
	} else if data.IpAddressSpace.ValueString() != "IPv4" {
		data.IpAddressSpace = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else if data.Type.ValueString() != "generic" {
		data.Type = types.StringNull()
	}
	if value := res.Get("response.ipPoolCidr"); value.Exists() && !data.IpSubnet.IsNull() {
		data.IpSubnet = types.StringValue(value.String())
	} else {
		data.IpSubnet = types.StringNull()
	}
	if value := res.Get("response.gateways"); value.Exists() && !data.Gateway.IsNull() {
		data.Gateway = helpers.GetStringSet(value.Array())
	} else {
		data.Gateway = types.SetNull(types.StringType)
	}
	if value := res.Get("response.dhcpServerIps"); value.Exists() && !data.DhcpServerIps.IsNull() {
		data.DhcpServerIps = helpers.GetStringSet(value.Array())
	} else {
		data.DhcpServerIps = types.SetNull(types.StringType)
	}
	if value := res.Get("response.dnsServerIps"); value.Exists() && !data.DnsServerIps.IsNull() {
		data.DnsServerIps = helpers.GetStringSet(value.Array())
	} else {
		data.DnsServerIps = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *IPPool) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.IpAddressSpace.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.IpSubnet.IsNull() {
		return false
	}
	if !data.Gateway.IsNull() {
		return false
	}
	if !data.DhcpServerIps.IsNull() {
		return false
	}
	if !data.DnsServerIps.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
