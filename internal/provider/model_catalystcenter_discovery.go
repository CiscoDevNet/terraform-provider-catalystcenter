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
type Discovery struct {
	Id                     types.String `tfsdk:"id"`
	CdpLevel               types.Int64  `tfsdk:"cdp_level"`
	DiscoveryType          types.String `tfsdk:"discovery_type"`
	EnablePasswordList     types.Set    `tfsdk:"enable_password_list"`
	GlobalCredentialIdList types.Set    `tfsdk:"global_credential_id_list"`
	HttpReadCredential     types.String `tfsdk:"http_read_credential"`
	HttpWriteCredential    types.String `tfsdk:"http_write_credential"`
	IpAddressList          types.String `tfsdk:"ip_address_list"`
	IpFilterList           types.Set    `tfsdk:"ip_filter_list"`
	LldpLevel              types.Int64  `tfsdk:"lldp_level"`
	Name                   types.String `tfsdk:"name"`
	NetconfPort            types.String `tfsdk:"netconf_port"`
	PasswordList           types.Set    `tfsdk:"password_list"`
	PreferredIpMethod      types.String `tfsdk:"preferred_ip_method"`
	ProtocolOrder          types.String `tfsdk:"protocol_order"`
	Retry                  types.Int64  `tfsdk:"retry"`
	SnmpAuthPassphrase     types.String `tfsdk:"snmp_auth_passphrase"`
	SnmpAuthProtocol       types.String `tfsdk:"snmp_auth_protocol"`
	SnmpMode               types.String `tfsdk:"snmp_mode"`
	SnmpPrivPassphrase     types.String `tfsdk:"snmp_priv_passphrase"`
	SnmpPrivProtocol       types.String `tfsdk:"snmp_priv_protocol"`
	SnmpRoCommunity        types.String `tfsdk:"snmp_ro_community"`
	SnmpRoCommunityDesc    types.String `tfsdk:"snmp_ro_community_desc"`
	SnmpRwCommunity        types.String `tfsdk:"snmp_rw_community"`
	SnmpRwCommunityDesc    types.String `tfsdk:"snmp_rw_community_desc"`
	SnmpUserName           types.String `tfsdk:"snmp_user_name"`
	SnmpVersion            types.String `tfsdk:"snmp_version"`
	TimeoutSeconds         types.Int64  `tfsdk:"timeout_seconds"`
	UserNameList           types.Set    `tfsdk:"user_name_list"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Discovery) getPath() string {
	return "/dna/intent/api/v1/discovery"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Discovery) toBody(ctx context.Context, state Discovery) string {
	body := ""
	if !data.CdpLevel.IsNull() {
		body, _ = sjson.Set(body, "cdpLevel", data.CdpLevel.ValueInt64())
	}
	if !data.DiscoveryType.IsNull() {
		body, _ = sjson.Set(body, "discoveryType", data.DiscoveryType.ValueString())
	}
	if !data.EnablePasswordList.IsNull() {
		var values []string
		data.EnablePasswordList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "enablePasswordList", values)
	}
	if !data.GlobalCredentialIdList.IsNull() {
		var values []string
		data.GlobalCredentialIdList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "globalCredentialIdList", values)
	}
	if !data.HttpReadCredential.IsNull() {
		body, _ = sjson.Set(body, "httpReadCredential", data.HttpReadCredential.ValueString())
	}
	if !data.HttpWriteCredential.IsNull() {
		body, _ = sjson.Set(body, "httpWriteCredential", data.HttpWriteCredential.ValueString())
	}
	if !data.IpAddressList.IsNull() {
		body, _ = sjson.Set(body, "ipAddressList", data.IpAddressList.ValueString())
	}
	if !data.IpFilterList.IsNull() {
		var values []string
		data.IpFilterList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ipFilterList", values)
	}
	if !data.LldpLevel.IsNull() {
		body, _ = sjson.Set(body, "lldpLevel", data.LldpLevel.ValueInt64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.NetconfPort.IsNull() {
		body, _ = sjson.Set(body, "netconfPort", data.NetconfPort.ValueString())
	}
	if !data.PasswordList.IsNull() {
		var values []string
		data.PasswordList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "passwordList", values)
	}
	if !data.PreferredIpMethod.IsNull() {
		body, _ = sjson.Set(body, "preferredIpMethod", data.PreferredIpMethod.ValueString())
	}
	if !data.ProtocolOrder.IsNull() {
		body, _ = sjson.Set(body, "protocolOrder", data.ProtocolOrder.ValueString())
	}
	if !data.Retry.IsNull() {
		body, _ = sjson.Set(body, "retry", data.Retry.ValueInt64())
	}
	if !data.SnmpAuthPassphrase.IsNull() {
		body, _ = sjson.Set(body, "snmpAuthPassphrase", data.SnmpAuthPassphrase.ValueString())
	}
	if !data.SnmpAuthProtocol.IsNull() {
		body, _ = sjson.Set(body, "snmpAuthProtocol", data.SnmpAuthProtocol.ValueString())
	}
	if !data.SnmpMode.IsNull() {
		body, _ = sjson.Set(body, "snmpMode", data.SnmpMode.ValueString())
	}
	if !data.SnmpPrivPassphrase.IsNull() {
		body, _ = sjson.Set(body, "snmpPrivPassphrase", data.SnmpPrivPassphrase.ValueString())
	}
	if !data.SnmpPrivProtocol.IsNull() {
		body, _ = sjson.Set(body, "snmpPrivProtocol", data.SnmpPrivProtocol.ValueString())
	}
	if !data.SnmpRoCommunity.IsNull() {
		body, _ = sjson.Set(body, "snmpRoCommunity", data.SnmpRoCommunity.ValueString())
	}
	if !data.SnmpRoCommunityDesc.IsNull() {
		body, _ = sjson.Set(body, "snmpRoCommunityDesc", data.SnmpRoCommunityDesc.ValueString())
	}
	if !data.SnmpRwCommunity.IsNull() {
		body, _ = sjson.Set(body, "snmpRwCommunity", data.SnmpRwCommunity.ValueString())
	}
	if !data.SnmpRwCommunityDesc.IsNull() {
		body, _ = sjson.Set(body, "snmpRwCommunityDesc", data.SnmpRwCommunityDesc.ValueString())
	}
	if !data.SnmpUserName.IsNull() {
		body, _ = sjson.Set(body, "snmpUserName", data.SnmpUserName.ValueString())
	}
	if !data.SnmpVersion.IsNull() {
		body, _ = sjson.Set(body, "snmpVersion", data.SnmpVersion.ValueString())
	}
	if !data.TimeoutSeconds.IsNull() {
		body, _ = sjson.Set(body, "timeout", data.TimeoutSeconds.ValueInt64())
	}
	if !data.UserNameList.IsNull() {
		var values []string
		data.UserNameList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "userNameList", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Discovery) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("cdpLevel"); value.Exists() {
		data.CdpLevel = types.Int64Value(value.Int())
	} else {
		data.CdpLevel = types.Int64Null()
	}
	if value := res.Get("discoveryType"); value.Exists() {
		data.DiscoveryType = types.StringValue(value.String())
	} else {
		data.DiscoveryType = types.StringNull()
	}
	if value := res.Get("enablePasswordList"); value.Exists() && len(value.Array()) > 0 {
		data.EnablePasswordList = helpers.GetStringSet(value.Array())
	} else {
		data.EnablePasswordList = types.SetNull(types.StringType)
	}
	if value := res.Get("globalCredentialIdList"); value.Exists() && len(value.Array()) > 0 {
		data.GlobalCredentialIdList = helpers.GetStringSet(value.Array())
	} else {
		data.GlobalCredentialIdList = types.SetNull(types.StringType)
	}
	if value := res.Get("httpReadCredential"); value.Exists() {
		data.HttpReadCredential = types.StringValue(value.String())
	} else {
		data.HttpReadCredential = types.StringNull()
	}
	if value := res.Get("httpWriteCredential"); value.Exists() {
		data.HttpWriteCredential = types.StringValue(value.String())
	} else {
		data.HttpWriteCredential = types.StringNull()
	}
	if value := res.Get("ipAddressList"); value.Exists() {
		data.IpAddressList = types.StringValue(value.String())
	} else {
		data.IpAddressList = types.StringNull()
	}
	if value := res.Get("ipFilterList"); value.Exists() && len(value.Array()) > 0 {
		data.IpFilterList = helpers.GetStringSet(value.Array())
	} else {
		data.IpFilterList = types.SetNull(types.StringType)
	}
	if value := res.Get("lldpLevel"); value.Exists() {
		data.LldpLevel = types.Int64Value(value.Int())
	} else {
		data.LldpLevel = types.Int64Null()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("netconfPort"); value.Exists() {
		data.NetconfPort = types.StringValue(value.String())
	} else {
		data.NetconfPort = types.StringNull()
	}
	if value := res.Get("passwordList"); value.Exists() && len(value.Array()) > 0 {
		data.PasswordList = helpers.GetStringSet(value.Array())
	} else {
		data.PasswordList = types.SetNull(types.StringType)
	}
	if value := res.Get("preferredMgmtIPMethod"); value.Exists() {
		data.PreferredIpMethod = types.StringValue(value.String())
	} else {
		data.PreferredIpMethod = types.StringValue("None")
	}
	if value := res.Get("protocolOrder"); value.Exists() {
		data.ProtocolOrder = types.StringValue(value.String())
	} else {
		data.ProtocolOrder = types.StringNull()
	}
	if value := res.Get("retryCount"); value.Exists() {
		data.Retry = types.Int64Value(value.Int())
	} else {
		data.Retry = types.Int64Null()
	}
	if value := res.Get("snmpAuthProtocol"); value.Exists() {
		data.SnmpAuthProtocol = types.StringValue(value.String())
	} else {
		data.SnmpAuthProtocol = types.StringNull()
	}
	if value := res.Get("snmpMode"); value.Exists() {
		data.SnmpMode = types.StringValue(value.String())
	} else {
		data.SnmpMode = types.StringNull()
	}
	if value := res.Get("snmpPrivProtocol"); value.Exists() {
		data.SnmpPrivProtocol = types.StringValue(value.String())
	} else {
		data.SnmpPrivProtocol = types.StringNull()
	}
	if value := res.Get("snmpRoCommunity"); value.Exists() {
		data.SnmpRoCommunity = types.StringValue(value.String())
	} else {
		data.SnmpRoCommunity = types.StringNull()
	}
	if value := res.Get("snmpRoCommunityDesc"); value.Exists() {
		data.SnmpRoCommunityDesc = types.StringValue(value.String())
	} else {
		data.SnmpRoCommunityDesc = types.StringNull()
	}
	if value := res.Get("snmpRwCommunity"); value.Exists() {
		data.SnmpRwCommunity = types.StringValue(value.String())
	} else {
		data.SnmpRwCommunity = types.StringNull()
	}
	if value := res.Get("snmpRwCommunityDesc"); value.Exists() {
		data.SnmpRwCommunityDesc = types.StringValue(value.String())
	} else {
		data.SnmpRwCommunityDesc = types.StringNull()
	}
	if value := res.Get("snmpUserName"); value.Exists() {
		data.SnmpUserName = types.StringValue(value.String())
	} else {
		data.SnmpUserName = types.StringNull()
	}
	if value := res.Get("userNameList"); value.Exists() && len(value.Array()) > 0 {
		data.UserNameList = helpers.GetStringSet(value.Array())
	} else {
		data.UserNameList = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Discovery) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("cdpLevel"); value.Exists() && !data.CdpLevel.IsNull() {
		data.CdpLevel = types.Int64Value(value.Int())
	} else {
		data.CdpLevel = types.Int64Null()
	}
	if value := res.Get("discoveryType"); value.Exists() && !data.DiscoveryType.IsNull() {
		data.DiscoveryType = types.StringValue(value.String())
	} else {
		data.DiscoveryType = types.StringNull()
	}
	if value := res.Get("enablePasswordList"); value.Exists() && !data.EnablePasswordList.IsNull() {
		data.EnablePasswordList = helpers.GetStringSet(value.Array())
	} else {
		data.EnablePasswordList = types.SetNull(types.StringType)
	}
	if value := res.Get("globalCredentialIdList"); value.Exists() && !data.GlobalCredentialIdList.IsNull() {
		data.GlobalCredentialIdList = helpers.GetStringSet(value.Array())
	} else {
		data.GlobalCredentialIdList = types.SetNull(types.StringType)
	}
	if value := res.Get("httpReadCredential"); value.Exists() && !data.HttpReadCredential.IsNull() {
		data.HttpReadCredential = types.StringValue(value.String())
	} else {
		data.HttpReadCredential = types.StringNull()
	}
	if value := res.Get("httpWriteCredential"); value.Exists() && !data.HttpWriteCredential.IsNull() {
		data.HttpWriteCredential = types.StringValue(value.String())
	} else {
		data.HttpWriteCredential = types.StringNull()
	}
	if value := res.Get("ipAddressList"); value.Exists() && !data.IpAddressList.IsNull() {
		data.IpAddressList = types.StringValue(value.String())
	} else {
		data.IpAddressList = types.StringNull()
	}
	if value := res.Get("ipFilterList"); value.Exists() && !data.IpFilterList.IsNull() {
		data.IpFilterList = helpers.GetStringSet(value.Array())
	} else {
		data.IpFilterList = types.SetNull(types.StringType)
	}
	if value := res.Get("lldpLevel"); value.Exists() && !data.LldpLevel.IsNull() {
		data.LldpLevel = types.Int64Value(value.Int())
	} else {
		data.LldpLevel = types.Int64Null()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("netconfPort"); value.Exists() && !data.NetconfPort.IsNull() {
		data.NetconfPort = types.StringValue(value.String())
	} else {
		data.NetconfPort = types.StringNull()
	}
	if value := res.Get("passwordList"); value.Exists() && !data.PasswordList.IsNull() {
		data.PasswordList = helpers.GetStringSet(value.Array())
	} else {
		data.PasswordList = types.SetNull(types.StringType)
	}
	if value := res.Get("preferredMgmtIPMethod"); value.Exists() && !data.PreferredIpMethod.IsNull() {
		data.PreferredIpMethod = types.StringValue(value.String())
	} else if data.PreferredIpMethod.ValueString() != "None" {
		data.PreferredIpMethod = types.StringNull()
	}
	if value := res.Get("protocolOrder"); value.Exists() && !data.ProtocolOrder.IsNull() {
		data.ProtocolOrder = types.StringValue(value.String())
	} else {
		data.ProtocolOrder = types.StringNull()
	}
	if value := res.Get("retryCount"); value.Exists() && !data.Retry.IsNull() {
		data.Retry = types.Int64Value(value.Int())
	} else {
		data.Retry = types.Int64Null()
	}
	if value := res.Get("snmpAuthProtocol"); value.Exists() && !data.SnmpAuthProtocol.IsNull() {
		data.SnmpAuthProtocol = types.StringValue(value.String())
	} else {
		data.SnmpAuthProtocol = types.StringNull()
	}
	if value := res.Get("snmpMode"); value.Exists() && !data.SnmpMode.IsNull() {
		data.SnmpMode = types.StringValue(value.String())
	} else {
		data.SnmpMode = types.StringNull()
	}
	if value := res.Get("snmpPrivProtocol"); value.Exists() && !data.SnmpPrivProtocol.IsNull() {
		data.SnmpPrivProtocol = types.StringValue(value.String())
	} else {
		data.SnmpPrivProtocol = types.StringNull()
	}
	if value := res.Get("snmpRoCommunity"); value.Exists() && !data.SnmpRoCommunity.IsNull() {
		data.SnmpRoCommunity = types.StringValue(value.String())
	} else {
		data.SnmpRoCommunity = types.StringNull()
	}
	if value := res.Get("snmpRoCommunityDesc"); value.Exists() && !data.SnmpRoCommunityDesc.IsNull() {
		data.SnmpRoCommunityDesc = types.StringValue(value.String())
	} else {
		data.SnmpRoCommunityDesc = types.StringNull()
	}
	if value := res.Get("snmpRwCommunity"); value.Exists() && !data.SnmpRwCommunity.IsNull() {
		data.SnmpRwCommunity = types.StringValue(value.String())
	} else {
		data.SnmpRwCommunity = types.StringNull()
	}
	if value := res.Get("snmpRwCommunityDesc"); value.Exists() && !data.SnmpRwCommunityDesc.IsNull() {
		data.SnmpRwCommunityDesc = types.StringValue(value.String())
	} else {
		data.SnmpRwCommunityDesc = types.StringNull()
	}
	if value := res.Get("snmpUserName"); value.Exists() && !data.SnmpUserName.IsNull() {
		data.SnmpUserName = types.StringValue(value.String())
	} else {
		data.SnmpUserName = types.StringNull()
	}
	if value := res.Get("userNameList"); value.Exists() && !data.UserNameList.IsNull() {
		data.UserNameList = helpers.GetStringSet(value.Array())
	} else {
		data.UserNameList = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Discovery) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.CdpLevel.IsNull() {
		return false
	}
	if !data.DiscoveryType.IsNull() {
		return false
	}
	if !data.EnablePasswordList.IsNull() {
		return false
	}
	if !data.GlobalCredentialIdList.IsNull() {
		return false
	}
	if !data.HttpReadCredential.IsNull() {
		return false
	}
	if !data.HttpWriteCredential.IsNull() {
		return false
	}
	if !data.IpAddressList.IsNull() {
		return false
	}
	if !data.IpFilterList.IsNull() {
		return false
	}
	if !data.LldpLevel.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.NetconfPort.IsNull() {
		return false
	}
	if !data.PasswordList.IsNull() {
		return false
	}
	if !data.PreferredIpMethod.IsNull() {
		return false
	}
	if !data.ProtocolOrder.IsNull() {
		return false
	}
	if !data.Retry.IsNull() {
		return false
	}
	if !data.SnmpAuthPassphrase.IsNull() {
		return false
	}
	if !data.SnmpAuthProtocol.IsNull() {
		return false
	}
	if !data.SnmpMode.IsNull() {
		return false
	}
	if !data.SnmpPrivPassphrase.IsNull() {
		return false
	}
	if !data.SnmpPrivProtocol.IsNull() {
		return false
	}
	if !data.SnmpRoCommunity.IsNull() {
		return false
	}
	if !data.SnmpRoCommunityDesc.IsNull() {
		return false
	}
	if !data.SnmpRwCommunity.IsNull() {
		return false
	}
	if !data.SnmpRwCommunityDesc.IsNull() {
		return false
	}
	if !data.SnmpUserName.IsNull() {
		return false
	}
	if !data.SnmpVersion.IsNull() {
		return false
	}
	if !data.TimeoutSeconds.IsNull() {
		return false
	}
	if !data.UserNameList.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
