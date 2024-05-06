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
type Network struct {
	Id                            types.String `tfsdk:"id"`
	SiteId                        types.String `tfsdk:"site_id"`
	DhcpServers                   types.Set    `tfsdk:"dhcp_servers"`
	DomainName                    types.String `tfsdk:"domain_name"`
	PrimaryDnsServer              types.String `tfsdk:"primary_dns_server"`
	SecondaryDnsServer            types.String `tfsdk:"secondary_dns_server"`
	SyslogServers                 types.Set    `tfsdk:"syslog_servers"`
	CatalystCenterAsSyslogServer  types.Bool   `tfsdk:"catalyst_center_as_syslog_server"`
	SnmpServers                   types.Set    `tfsdk:"snmp_servers"`
	CatalystCenterAsSnmpServer    types.Bool   `tfsdk:"catalyst_center_as_snmp_server"`
	NetflowCollector              types.String `tfsdk:"netflow_collector"`
	NetflowCollectorPort          types.Int64  `tfsdk:"netflow_collector_port"`
	NtpServers                    types.Set    `tfsdk:"ntp_servers"`
	Timezone                      types.String `tfsdk:"timezone"`
	NetworkAaaServerType          types.String `tfsdk:"network_aaa_server_type"`
	NetworkAaaServerPrimaryIp     types.String `tfsdk:"network_aaa_server_primary_ip"`
	NetworkAaaServerSecondaryIp   types.String `tfsdk:"network_aaa_server_secondary_ip"`
	NetworkAaaServerProtocol      types.String `tfsdk:"network_aaa_server_protocol"`
	NetworkAaaServerSharedSecret  types.String `tfsdk:"network_aaa_server_shared_secret"`
	EndpointAaaServerType         types.String `tfsdk:"endpoint_aaa_server_type"`
	EndpointAaaServerPrimaryIp    types.String `tfsdk:"endpoint_aaa_server_primary_ip"`
	EndpointAaaServerSecondaryIp  types.String `tfsdk:"endpoint_aaa_server_secondary_ip"`
	EndpointAaaServerProtocol     types.String `tfsdk:"endpoint_aaa_server_protocol"`
	EndpointAaaServerSharedSecret types.String `tfsdk:"endpoint_aaa_server_shared_secret"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Network) getPath() string {
	return "/dna/intent/api/v2/network"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Network) toBody(ctx context.Context, state Network) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DhcpServers.IsNull() {
		var values []string
		data.DhcpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "settings.dhcpServer", values)
	}
	if !data.DomainName.IsNull() {
		body, _ = sjson.Set(body, "settings.dnsServer.domainName", data.DomainName.ValueString())
	}
	if !data.PrimaryDnsServer.IsNull() {
		body, _ = sjson.Set(body, "settings.dnsServer.primaryIpAddress", data.PrimaryDnsServer.ValueString())
	}
	if !data.SecondaryDnsServer.IsNull() {
		body, _ = sjson.Set(body, "settings.dnsServer.secondaryIpAddress", data.SecondaryDnsServer.ValueString())
	}
	if !data.SyslogServers.IsNull() {
		var values []string
		data.SyslogServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "settings.syslogServer.ipAddresses", values)
	}
	if !data.CatalystCenterAsSyslogServer.IsNull() {
		body, _ = sjson.Set(body, "settings.syslogServer.configureDnacIP", data.CatalystCenterAsSyslogServer.ValueBool())
	}
	if !data.SnmpServers.IsNull() {
		var values []string
		data.SnmpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "settings.snmpServer.ipAddresses", values)
	}
	if !data.CatalystCenterAsSnmpServer.IsNull() {
		body, _ = sjson.Set(body, "settings.snmpServer.configureDnacIP", data.CatalystCenterAsSnmpServer.ValueBool())
	}
	if !data.NetflowCollector.IsNull() {
		body, _ = sjson.Set(body, "settings.netflowcollector.ipAddress", data.NetflowCollector.ValueString())
	}
	if !data.NetflowCollectorPort.IsNull() {
		body, _ = sjson.Set(body, "settings.netflowcollector.port", data.NetflowCollectorPort.ValueInt64())
	}
	if !data.NtpServers.IsNull() {
		var values []string
		data.NtpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "settings.ntpServer", values)
	}
	if !data.Timezone.IsNull() {
		body, _ = sjson.Set(body, "settings.timezone", data.Timezone.ValueString())
	}
	if !data.NetworkAaaServerType.IsNull() {
		body, _ = sjson.Set(body, "settings.network_aaa.servers", data.NetworkAaaServerType.ValueString())
	}
	if !data.NetworkAaaServerPrimaryIp.IsNull() {
		body, _ = sjson.Set(body, "settings.network_aaa.network", data.NetworkAaaServerPrimaryIp.ValueString())
	}
	if !data.NetworkAaaServerSecondaryIp.IsNull() {
		body, _ = sjson.Set(body, "settings.network_aaa.ipAddress", data.NetworkAaaServerSecondaryIp.ValueString())
	}
	if !data.NetworkAaaServerProtocol.IsNull() {
		body, _ = sjson.Set(body, "settings.network_aaa.protocol", data.NetworkAaaServerProtocol.ValueString())
	}
	if !data.NetworkAaaServerSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "settings.network_aaa.sharedSecret", data.NetworkAaaServerSharedSecret.ValueString())
	}
	if !data.EndpointAaaServerType.IsNull() {
		body, _ = sjson.Set(body, "settings.clientAndEndpoint_aaa.servers", data.EndpointAaaServerType.ValueString())
	}
	if !data.EndpointAaaServerPrimaryIp.IsNull() {
		body, _ = sjson.Set(body, "settings.clientAndEndpoint_aaa.network", data.EndpointAaaServerPrimaryIp.ValueString())
	}
	if !data.EndpointAaaServerSecondaryIp.IsNull() {
		body, _ = sjson.Set(body, "settings.clientAndEndpoint_aaa.ipAddress", data.EndpointAaaServerSecondaryIp.ValueString())
	}
	if !data.EndpointAaaServerProtocol.IsNull() {
		body, _ = sjson.Set(body, "settings.clientAndEndpoint_aaa.protocol", data.EndpointAaaServerProtocol.ValueString())
	}
	if !data.EndpointAaaServerSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "settings.clientAndEndpoint_aaa.sharedSecret", data.EndpointAaaServerSharedSecret.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Network) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.#(key=\"dhcp.server\").value"); value.Exists() && len(value.Array()) > 0 {
		data.DhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.DhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"dns.server\").value.0.domainName"); value.Exists() {
		data.DomainName = types.StringValue(value.String())
	} else {
		data.DomainName = types.StringNull()
	}
	if value := res.Get("response.#(key=\"dns.server\").value.0.primaryIpAddress"); value.Exists() {
		data.PrimaryDnsServer = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServer = types.StringNull()
	}
	if value := res.Get("response.#(key=\"dns.server\").value.0.secondaryIpAddress"); value.Exists() {
		data.SecondaryDnsServer = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServer = types.StringNull()
	}
	if value := res.Get("response.#(key=\"syslog.server\").value.0.ipAddresses"); value.Exists() && len(value.Array()) > 0 {
		data.SyslogServers = helpers.GetStringSet(value.Array())
	} else {
		data.SyslogServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"syslog.server\").value.0.configureDnacIP"); value.Exists() {
		data.CatalystCenterAsSyslogServer = types.BoolValue(value.Bool())
	} else {
		data.CatalystCenterAsSyslogServer = types.BoolNull()
	}
	if value := res.Get("response.#(key=\"snmp.trap.receiver\").value.0.ipAddresses"); value.Exists() && len(value.Array()) > 0 {
		data.SnmpServers = helpers.GetStringSet(value.Array())
	} else {
		data.SnmpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"snmp.trap.receiver\").value.0.configureDnacIP"); value.Exists() {
		data.CatalystCenterAsSnmpServer = types.BoolValue(value.Bool())
	} else {
		data.CatalystCenterAsSnmpServer = types.BoolNull()
	}
	if value := res.Get("response.#(key=\"netflow.collector\").value.0.ipAddress"); value.Exists() {
		data.NetflowCollector = types.StringValue(value.String())
	} else {
		data.NetflowCollector = types.StringNull()
	}
	if value := res.Get("response.#(key=\"netflow.collector\").value.0.port"); value.Exists() {
		data.NetflowCollectorPort = types.Int64Value(value.Int())
	} else {
		data.NetflowCollectorPort = types.Int64Null()
	}
	if value := res.Get("response.#(key=\"ntp.server\").value"); value.Exists() && len(value.Array()) > 0 {
		data.NtpServers = helpers.GetStringSet(value.Array())
	} else {
		data.NtpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"timezone.site\").value.0"); value.Exists() {
		data.Timezone = types.StringValue(value.String())
	} else {
		data.Timezone = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Network) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.#(key=\"dhcp.server\").value"); value.Exists() && !data.DhcpServers.IsNull() {
		data.DhcpServers = helpers.GetStringSet(value.Array())
	} else {
		data.DhcpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"dns.server\").value.0.domainName"); value.Exists() && !data.DomainName.IsNull() {
		data.DomainName = types.StringValue(value.String())
	} else {
		data.DomainName = types.StringNull()
	}
	if value := res.Get("response.#(key=\"dns.server\").value.0.primaryIpAddress"); value.Exists() && !data.PrimaryDnsServer.IsNull() {
		data.PrimaryDnsServer = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServer = types.StringNull()
	}
	if value := res.Get("response.#(key=\"dns.server\").value.0.secondaryIpAddress"); value.Exists() && !data.SecondaryDnsServer.IsNull() {
		data.SecondaryDnsServer = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServer = types.StringNull()
	}
	if value := res.Get("response.#(key=\"syslog.server\").value.0.ipAddresses"); value.Exists() && !data.SyslogServers.IsNull() {
		data.SyslogServers = helpers.GetStringSet(value.Array())
	} else {
		data.SyslogServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"syslog.server\").value.0.configureDnacIP"); value.Exists() && !data.CatalystCenterAsSyslogServer.IsNull() {
		data.CatalystCenterAsSyslogServer = types.BoolValue(value.Bool())
	} else {
		data.CatalystCenterAsSyslogServer = types.BoolNull()
	}
	if value := res.Get("response.#(key=\"snmp.trap.receiver\").value.0.ipAddresses"); value.Exists() && !data.SnmpServers.IsNull() {
		data.SnmpServers = helpers.GetStringSet(value.Array())
	} else {
		data.SnmpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"snmp.trap.receiver\").value.0.configureDnacIP"); value.Exists() && !data.CatalystCenterAsSnmpServer.IsNull() {
		data.CatalystCenterAsSnmpServer = types.BoolValue(value.Bool())
	} else {
		data.CatalystCenterAsSnmpServer = types.BoolNull()
	}
	if value := res.Get("response.#(key=\"netflow.collector\").value.0.ipAddress"); value.Exists() && !data.NetflowCollector.IsNull() {
		data.NetflowCollector = types.StringValue(value.String())
	} else {
		data.NetflowCollector = types.StringNull()
	}
	if value := res.Get("response.#(key=\"netflow.collector\").value.0.port"); value.Exists() && !data.NetflowCollectorPort.IsNull() {
		data.NetflowCollectorPort = types.Int64Value(value.Int())
	} else {
		data.NetflowCollectorPort = types.Int64Null()
	}
	if value := res.Get("response.#(key=\"ntp.server\").value"); value.Exists() && !data.NtpServers.IsNull() {
		data.NtpServers = helpers.GetStringSet(value.Array())
	} else {
		data.NtpServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.#(key=\"timezone.site\").value.0"); value.Exists() && !data.Timezone.IsNull() {
		data.Timezone = types.StringValue(value.String())
	} else {
		data.Timezone = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Network) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DhcpServers.IsNull() {
		return false
	}
	if !data.DomainName.IsNull() {
		return false
	}
	if !data.PrimaryDnsServer.IsNull() {
		return false
	}
	if !data.SecondaryDnsServer.IsNull() {
		return false
	}
	if !data.SyslogServers.IsNull() {
		return false
	}
	if !data.CatalystCenterAsSyslogServer.IsNull() {
		return false
	}
	if !data.SnmpServers.IsNull() {
		return false
	}
	if !data.CatalystCenterAsSnmpServer.IsNull() {
		return false
	}
	if !data.NetflowCollector.IsNull() {
		return false
	}
	if !data.NetflowCollectorPort.IsNull() {
		return false
	}
	if !data.NtpServers.IsNull() {
		return false
	}
	if !data.Timezone.IsNull() {
		return false
	}
	if !data.NetworkAaaServerType.IsNull() {
		return false
	}
	if !data.NetworkAaaServerPrimaryIp.IsNull() {
		return false
	}
	if !data.NetworkAaaServerSecondaryIp.IsNull() {
		return false
	}
	if !data.NetworkAaaServerProtocol.IsNull() {
		return false
	}
	if !data.NetworkAaaServerSharedSecret.IsNull() {
		return false
	}
	if !data.EndpointAaaServerType.IsNull() {
		return false
	}
	if !data.EndpointAaaServerPrimaryIp.IsNull() {
		return false
	}
	if !data.EndpointAaaServerSecondaryIp.IsNull() {
		return false
	}
	if !data.EndpointAaaServerProtocol.IsNull() {
		return false
	}
	if !data.EndpointAaaServerSharedSecret.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
