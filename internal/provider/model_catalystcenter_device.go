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
type Device struct {
	Id                    types.String                  `tfsdk:"id"`
	CliTransport          types.String                  `tfsdk:"cli_transport"`
	ComputeDevice         types.Bool                    `tfsdk:"compute_device"`
	EnablePassword        types.String                  `tfsdk:"enable_password"`
	ExtendedDiscoveryInfo types.String                  `tfsdk:"extended_discovery_info"`
	HttpPassword          types.String                  `tfsdk:"http_password"`
	HttpPort              types.String                  `tfsdk:"http_port"`
	HttpSecure            types.Bool                    `tfsdk:"http_secure"`
	HttpUserName          types.String                  `tfsdk:"http_user_name"`
	IpAddress             types.String                  `tfsdk:"ip_address"`
	MerakiOrgIds          types.Set                     `tfsdk:"meraki_org_ids"`
	NetconfPort           types.String                  `tfsdk:"netconf_port"`
	Password              types.String                  `tfsdk:"password"`
	SerialNumber          types.String                  `tfsdk:"serial_number"`
	SnmpAuthPassphrase    types.String                  `tfsdk:"snmp_auth_passphrase"`
	SnmpAuthProtocol      types.String                  `tfsdk:"snmp_auth_protocol"`
	SnmpMode              types.String                  `tfsdk:"snmp_mode"`
	SnmpPrivPassphrase    types.String                  `tfsdk:"snmp_priv_passphrase"`
	SnmpPrivProtocol      types.String                  `tfsdk:"snmp_priv_protocol"`
	SnmpRoCommunity       types.String                  `tfsdk:"snmp_ro_community"`
	SnmpRwCommunity       types.String                  `tfsdk:"snmp_rw_community"`
	SnmpRetry             types.Int64                   `tfsdk:"snmp_retry"`
	SnmpTimeout           types.Int64                   `tfsdk:"snmp_timeout"`
	SnmpUserName          types.String                  `tfsdk:"snmp_user_name"`
	SnmpVersion           types.String                  `tfsdk:"snmp_version"`
	Type                  types.String                  `tfsdk:"type"`
	UpdateMgmtIpAddresses []DeviceUpdateMgmtIpAddresses `tfsdk:"update_mgmt_ip_addresses"`
	UserName              types.String                  `tfsdk:"user_name"`
}

type DeviceUpdateMgmtIpAddresses struct {
	ExistMgmtIpAddress types.String `tfsdk:"exist_mgmt_ip_address"`
	NewMgmtIpAddress   types.String `tfsdk:"new_mgmt_ip_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Device) getPath() string {
	return "/dna/intent/api/v1/network-device"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Device) toBody(ctx context.Context, state Device) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.CliTransport.IsNull() {
		body, _ = sjson.Set(body, "cliTransport", data.CliTransport.ValueString())
	}
	if !data.ComputeDevice.IsNull() {
		body, _ = sjson.Set(body, "computeDevice", data.ComputeDevice.ValueBool())
	}
	if !data.EnablePassword.IsNull() {
		body, _ = sjson.Set(body, "enablePassword", data.EnablePassword.ValueString())
	}
	if !data.ExtendedDiscoveryInfo.IsNull() {
		body, _ = sjson.Set(body, "extendedDiscoveryInfo", data.ExtendedDiscoveryInfo.ValueString())
	}
	if !data.HttpPassword.IsNull() {
		body, _ = sjson.Set(body, "httpPassword", data.HttpPassword.ValueString())
	}
	if !data.HttpPort.IsNull() {
		body, _ = sjson.Set(body, "httpPort", data.HttpPort.ValueString())
	}
	if !data.HttpSecure.IsNull() {
		body, _ = sjson.Set(body, "httpSecure", data.HttpSecure.ValueBool())
	}
	if !data.HttpUserName.IsNull() {
		body, _ = sjson.Set(body, "httpUserName", data.HttpUserName.ValueString())
	}
	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, "ipAddress.0", data.IpAddress.ValueString())
	}
	if !data.MerakiOrgIds.IsNull() {
		var values []string
		data.MerakiOrgIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "merakiOrgId", values)
	}
	if !data.NetconfPort.IsNull() {
		body, _ = sjson.Set(body, "netconfPort", data.NetconfPort.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, "password", data.Password.ValueString())
	}
	if !data.SerialNumber.IsNull() {
		body, _ = sjson.Set(body, "serialNumber", data.SerialNumber.ValueString())
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
		body, _ = sjson.Set(body, "snmpROCommunity", data.SnmpRoCommunity.ValueString())
	}
	if !data.SnmpRwCommunity.IsNull() {
		body, _ = sjson.Set(body, "snmpRWCommunity", data.SnmpRwCommunity.ValueString())
	}
	if !data.SnmpRetry.IsNull() {
		body, _ = sjson.Set(body, "snmpRetry", data.SnmpRetry.ValueInt64())
	}
	if !data.SnmpTimeout.IsNull() {
		body, _ = sjson.Set(body, "snmpTimeout", data.SnmpTimeout.ValueInt64())
	}
	if !data.SnmpUserName.IsNull() {
		body, _ = sjson.Set(body, "snmpUserName", data.SnmpUserName.ValueString())
	}
	if !data.SnmpVersion.IsNull() {
		body, _ = sjson.Set(body, "snmpVersion", data.SnmpVersion.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if len(data.UpdateMgmtIpAddresses) > 0 {
		body, _ = sjson.Set(body, "updateMgmtIPaddressList", []interface{}{})
		for _, item := range data.UpdateMgmtIpAddresses {
			itemBody := ""
			if !item.ExistMgmtIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "existMgmtIpAddress", item.ExistMgmtIpAddress.ValueString())
			}
			if !item.NewMgmtIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "newMgmtIpAddress", item.NewMgmtIpAddress.ValueString())
			}
			body, _ = sjson.SetRaw(body, "updateMgmtIPaddressList.-1", itemBody)
		}
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, "userName", data.UserName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Device) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.managementIpAddress"); value.Exists() {
		data.IpAddress = types.StringValue(value.String())
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get("response.serialNumber"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Device) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.managementIpAddress"); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(value.String())
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get("response.serialNumber"); value.Exists() && !data.SerialNumber.IsNull() {
		data.SerialNumber = types.StringValue(value.String())
	} else {
		data.SerialNumber = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Device) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.CliTransport.IsNull() {
		return false
	}
	if !data.ComputeDevice.IsNull() {
		return false
	}
	if !data.EnablePassword.IsNull() {
		return false
	}
	if !data.ExtendedDiscoveryInfo.IsNull() {
		return false
	}
	if !data.HttpPassword.IsNull() {
		return false
	}
	if !data.HttpPort.IsNull() {
		return false
	}
	if !data.HttpSecure.IsNull() {
		return false
	}
	if !data.HttpUserName.IsNull() {
		return false
	}
	if !data.IpAddress.IsNull() {
		return false
	}
	if !data.MerakiOrgIds.IsNull() {
		return false
	}
	if !data.NetconfPort.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.SerialNumber.IsNull() {
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
	if !data.SnmpRwCommunity.IsNull() {
		return false
	}
	if !data.SnmpRetry.IsNull() {
		return false
	}
	if !data.SnmpTimeout.IsNull() {
		return false
	}
	if !data.SnmpUserName.IsNull() {
		return false
	}
	if !data.SnmpVersion.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if len(data.UpdateMgmtIpAddresses) > 0 {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
