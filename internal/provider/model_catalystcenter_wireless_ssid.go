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
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessSSID struct {
	Id                                     types.String                   `tfsdk:"id"`
	SiteId                                 types.String                   `tfsdk:"site_id"`
	Ssid                                   types.String                   `tfsdk:"ssid"`
	AuthType                               types.String                   `tfsdk:"auth_type"`
	Passphrase                             types.String                   `tfsdk:"passphrase"`
	FastLane                               types.Bool                     `tfsdk:"fast_lane"`
	MacFiltering                           types.Bool                     `tfsdk:"mac_filtering"`
	SsidRadioType                          types.String                   `tfsdk:"ssid_radio_type"`
	BroadcastSsid                          types.Bool                     `tfsdk:"broadcast_ssid"`
	FastTransition                         types.String                   `tfsdk:"fast_transition"`
	SessionTimeoutEnable                   types.Bool                     `tfsdk:"session_timeout_enable"`
	SessionTimeout                         types.Int64                    `tfsdk:"session_timeout"`
	ClientExclusion                        types.Bool                     `tfsdk:"client_exclusion"`
	ClientExclusionTimeout                 types.Int64                    `tfsdk:"client_exclusion_timeout"`
	BasicServiceSetMaxIdle                 types.Bool                     `tfsdk:"basic_service_set_max_idle"`
	BasicServiceSetClientIdleTimeout       types.Int64                    `tfsdk:"basic_service_set_client_idle_timeout"`
	DirectedMulticastService               types.Bool                     `tfsdk:"directed_multicast_service"`
	NeighborList                           types.Bool                     `tfsdk:"neighbor_list"`
	MftClientProtection                    types.String                   `tfsdk:"mft_client_protection"`
	NasOptions                             types.Set                      `tfsdk:"nas_options"`
	ProfileName                            types.String                   `tfsdk:"profile_name"`
	AaaOverride                            types.Bool                     `tfsdk:"aaa_override"`
	CoverageHoleDetection                  types.Bool                     `tfsdk:"coverage_hole_detection"`
	ProtectedManagementFrame               types.String                   `tfsdk:"protected_management_frame"`
	MultiPskSettings                       []WirelessSSIDMultiPskSettings `tfsdk:"multi_psk_settings"`
	ClientRateLimit                        types.Int64                    `tfsdk:"client_rate_limit"`
	RsnCipherSuiteGcmp256                  types.Bool                     `tfsdk:"rsn_cipher_suite_gcmp256"`
	RsnCipherSuiteCcmp256                  types.Bool                     `tfsdk:"rsn_cipher_suite_ccmp256"`
	RsnCipherSuiteGcmp128                  types.Bool                     `tfsdk:"rsn_cipher_suite_gcmp128"`
	RsnCipherSuiteCcmp128                  types.Bool                     `tfsdk:"rsn_cipher_suite_ccmp128"`
	Ghz6PolicyClientSteering               types.Bool                     `tfsdk:"ghz6_policy_client_steering"`
	AuthKey8021x                           types.Bool                     `tfsdk:"auth_key8021x"`
	AuthKey8021xPlusFt                     types.Bool                     `tfsdk:"auth_key8021x_plus_ft"`
	AuthKey8021xSha256                     types.Bool                     `tfsdk:"auth_key8021x_sha256"`
	AuthKeySae                             types.Bool                     `tfsdk:"auth_key_sae"`
	AuthKeySaePlusFt                       types.Bool                     `tfsdk:"auth_key_sae_plus_ft"`
	AuthKeyPsk                             types.Bool                     `tfsdk:"auth_key_psk"`
	AuthKeyPskPlusFt                       types.Bool                     `tfsdk:"auth_key_psk_plus_ft"`
	AuthKeyOwe                             types.Bool                     `tfsdk:"auth_key_owe"`
	AuthKeyEasyPsk                         types.Bool                     `tfsdk:"auth_key_easy_psk"`
	AuthKeyEasyPskSha256                   types.Bool                     `tfsdk:"auth_key_easy_psk_sha256"`
	OpenSsid                               types.Bool                     `tfsdk:"open_ssid"`
	WlanBandSelect                         types.Bool                     `tfsdk:"wlan_band_select"`
	Enabled                                types.Bool                     `tfsdk:"enabled"`
	AuthServers                            types.Set                      `tfsdk:"auth_servers"`
	AcctServers                            types.Set                      `tfsdk:"acct_servers"`
	EgressQos                              types.String                   `tfsdk:"egress_qos"`
	IngressQos                             types.String                   `tfsdk:"ingress_qos"`
	WlanType                               types.String                   `tfsdk:"wlan_type"`
	L3AuthType                             types.String                   `tfsdk:"l3_auth_type"`
	AuthServer                             types.String                   `tfsdk:"auth_server"`
	ExternalAuthIpAddress                  types.String                   `tfsdk:"external_auth_ip_address"`
	WebPassthrough                         types.Bool                     `tfsdk:"web_passthrough"`
	SleepingClient                         types.Bool                     `tfsdk:"sleeping_client"`
	SleepingClientTimeout                  types.Int64                    `tfsdk:"sleeping_client_timeout"`
	AclName                                types.String                   `tfsdk:"acl_name"`
	Posturing                              types.Bool                     `tfsdk:"posturing"`
	AuthKeySuiteB1x                        types.Bool                     `tfsdk:"auth_key_suite_b1x"`
	AuthKeySuiteB1921x                     types.Bool                     `tfsdk:"auth_key_suite_b1921x"`
	AuthKeySaeExt                          types.Bool                     `tfsdk:"auth_key_sae_ext"`
	AuthKeySaeExtPlusFt                    types.Bool                     `tfsdk:"auth_key_sae_ext_plus_ft"`
	ApBeaconProtection                     types.Bool                     `tfsdk:"ap_beacon_protection"`
	Ghz24Policy                            types.String                   `tfsdk:"ghz24_policy"`
	CckmTsfTolerance                       types.Int64                    `tfsdk:"cckm_tsf_tolerance"`
	Cckm                                   types.Bool                     `tfsdk:"cckm"`
	Hex                                    types.Bool                     `tfsdk:"hex"`
	RandomMacFilter                        types.Bool                     `tfsdk:"random_mac_filter"`
	FastTransitionOverTheDistributedSystem types.Bool                     `tfsdk:"fast_transition_over_the_distributed_system"`
}

type WirelessSSIDMultiPskSettings struct {
	Priority       types.String `tfsdk:"priority"`
	PassphraseType types.String `tfsdk:"passphrase_type"`
	Passphrase     types.String `tfsdk:"passphrase"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessSSID) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/wirelessSettings/ssids", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessSSID) toBody(ctx context.Context, state WirelessSSID) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Ssid.IsNull() {
		body, _ = sjson.Set(body, "ssid", data.Ssid.ValueString())
	}
	if !data.AuthType.IsNull() {
		body, _ = sjson.Set(body, "authType", data.AuthType.ValueString())
	}
	if !data.Passphrase.IsNull() {
		body, _ = sjson.Set(body, "passphrase", data.Passphrase.ValueString())
	}
	if !data.FastLane.IsNull() {
		body, _ = sjson.Set(body, "isFastLaneEnabled", data.FastLane.ValueBool())
	}
	if !data.MacFiltering.IsNull() {
		body, _ = sjson.Set(body, "isMacFilteringEnabled", data.MacFiltering.ValueBool())
	}
	if !data.SsidRadioType.IsNull() {
		body, _ = sjson.Set(body, "ssidRadioType", data.SsidRadioType.ValueString())
	}
	if !data.BroadcastSsid.IsNull() {
		body, _ = sjson.Set(body, "isBroadcastSSID", data.BroadcastSsid.ValueBool())
	}
	if !data.FastTransition.IsNull() {
		body, _ = sjson.Set(body, "fastTransition", data.FastTransition.ValueString())
	}
	if !data.SessionTimeoutEnable.IsNull() {
		body, _ = sjson.Set(body, "sessionTimeOutEnable", data.SessionTimeoutEnable.ValueBool())
	}
	if !data.SessionTimeout.IsNull() {
		body, _ = sjson.Set(body, "sessionTimeOut", data.SessionTimeout.ValueInt64())
	}
	if !data.ClientExclusion.IsNull() {
		body, _ = sjson.Set(body, "clientExclusionEnable", data.ClientExclusion.ValueBool())
	}
	if !data.ClientExclusionTimeout.IsNull() {
		body, _ = sjson.Set(body, "clientExclusionTimeout", data.ClientExclusionTimeout.ValueInt64())
	}
	if !data.BasicServiceSetMaxIdle.IsNull() {
		body, _ = sjson.Set(body, "basicServiceSetMaxIdleEnable", data.BasicServiceSetMaxIdle.ValueBool())
	}
	if !data.BasicServiceSetClientIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, "basicServiceSetClientIdleTimeout", data.BasicServiceSetClientIdleTimeout.ValueInt64())
	}
	if !data.DirectedMulticastService.IsNull() {
		body, _ = sjson.Set(body, "directedMulticastServiceEnable", data.DirectedMulticastService.ValueBool())
	}
	if !data.NeighborList.IsNull() {
		body, _ = sjson.Set(body, "neighborListEnable", data.NeighborList.ValueBool())
	}
	if !data.MftClientProtection.IsNull() {
		body, _ = sjson.Set(body, "managementFrameProtectionClientprotection", data.MftClientProtection.ValueString())
	}
	if !data.NasOptions.IsNull() {
		var values []string
		data.NasOptions.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "nasOptions", values)
	}
	if !data.ProfileName.IsNull() {
		body, _ = sjson.Set(body, "profileName", data.ProfileName.ValueString())
	}
	if !data.AaaOverride.IsNull() {
		body, _ = sjson.Set(body, "aaaOverride", data.AaaOverride.ValueBool())
	}
	if !data.CoverageHoleDetection.IsNull() {
		body, _ = sjson.Set(body, "coverageHoleDetectionEnable", data.CoverageHoleDetection.ValueBool())
	}
	if !data.ProtectedManagementFrame.IsNull() {
		body, _ = sjson.Set(body, "protectedManagementFrame", data.ProtectedManagementFrame.ValueString())
	}
	if len(data.MultiPskSettings) > 0 {
		body, _ = sjson.Set(body, "multiPSKSettings", []interface{}{})
		for _, item := range data.MultiPskSettings {
			itemBody := ""
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueString())
			}
			if !item.PassphraseType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "passphraseType", item.PassphraseType.ValueString())
			}
			if !item.Passphrase.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "passphrase", item.Passphrase.ValueString())
			}
			body, _ = sjson.SetRaw(body, "multiPSKSettings.-1", itemBody)
		}
	}
	if !data.ClientRateLimit.IsNull() {
		body, _ = sjson.Set(body, "clientRateLimit", data.ClientRateLimit.ValueInt64())
	}
	if !data.RsnCipherSuiteGcmp256.IsNull() {
		body, _ = sjson.Set(body, "rsnCipherSuiteGcmp256", data.RsnCipherSuiteGcmp256.ValueBool())
	}
	if !data.RsnCipherSuiteCcmp256.IsNull() {
		body, _ = sjson.Set(body, "rsnCipherSuiteCcmp256", data.RsnCipherSuiteCcmp256.ValueBool())
	}
	if !data.RsnCipherSuiteGcmp128.IsNull() {
		body, _ = sjson.Set(body, "rsnCipherSuiteGcmp128", data.RsnCipherSuiteGcmp128.ValueBool())
	}
	if !data.RsnCipherSuiteCcmp128.IsNull() {
		body, _ = sjson.Set(body, "rsnCipherSuiteCcmp128", data.RsnCipherSuiteCcmp128.ValueBool())
	}
	if !data.Ghz6PolicyClientSteering.IsNull() {
		body, _ = sjson.Set(body, "ghz6PolicyClientSteering", data.Ghz6PolicyClientSteering.ValueBool())
	}
	if !data.AuthKey8021x.IsNull() {
		body, _ = sjson.Set(body, "isAuthKey8021x", data.AuthKey8021x.ValueBool())
	}
	if !data.AuthKey8021xPlusFt.IsNull() {
		body, _ = sjson.Set(body, "isAuthKey8021xPlusFT", data.AuthKey8021xPlusFt.ValueBool())
	}
	if !data.AuthKey8021xSha256.IsNull() {
		body, _ = sjson.Set(body, "isAuthKey8021x_SHA256", data.AuthKey8021xSha256.ValueBool())
	}
	if !data.AuthKeySae.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeySae", data.AuthKeySae.ValueBool())
	}
	if !data.AuthKeySaePlusFt.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeySaePlusFT", data.AuthKeySaePlusFt.ValueBool())
	}
	if !data.AuthKeyPsk.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeyPSK", data.AuthKeyPsk.ValueBool())
	}
	if !data.AuthKeyPskPlusFt.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeyPSKPlusFT", data.AuthKeyPskPlusFt.ValueBool())
	}
	if !data.AuthKeyOwe.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeyOWE", data.AuthKeyOwe.ValueBool())
	}
	if !data.AuthKeyEasyPsk.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeyEasyPSK", data.AuthKeyEasyPsk.ValueBool())
	}
	if !data.AuthKeyEasyPskSha256.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeyPSKSHA256", data.AuthKeyEasyPskSha256.ValueBool())
	}
	if !data.OpenSsid.IsNull() {
		body, _ = sjson.Set(body, "openSsid", data.OpenSsid.ValueBool())
	}
	if !data.WlanBandSelect.IsNull() {
		body, _ = sjson.Set(body, "wlanBandSelectEnable", data.WlanBandSelect.ValueBool())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "isEnabled", data.Enabled.ValueBool())
	}
	if !data.AuthServers.IsNull() {
		var values []string
		data.AuthServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "authServers", values)
	}
	if !data.AcctServers.IsNull() {
		var values []string
		data.AcctServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "acctServers", values)
	}
	if !data.EgressQos.IsNull() {
		body, _ = sjson.Set(body, "egressQos", data.EgressQos.ValueString())
	}
	if !data.IngressQos.IsNull() {
		body, _ = sjson.Set(body, "ingressQos", data.IngressQos.ValueString())
	}
	if !data.WlanType.IsNull() {
		body, _ = sjson.Set(body, "wlanType", data.WlanType.ValueString())
	}
	if !data.L3AuthType.IsNull() {
		body, _ = sjson.Set(body, "l3AuthType", data.L3AuthType.ValueString())
	}
	if !data.AuthServer.IsNull() {
		body, _ = sjson.Set(body, "authServer", data.AuthServer.ValueString())
	}
	if !data.ExternalAuthIpAddress.IsNull() {
		body, _ = sjson.Set(body, "externalAuthIpAddress", data.ExternalAuthIpAddress.ValueString())
	}
	if !data.WebPassthrough.IsNull() {
		body, _ = sjson.Set(body, "webPassthrough", data.WebPassthrough.ValueBool())
	}
	if !data.SleepingClient.IsNull() {
		body, _ = sjson.Set(body, "sleepingClientEnable", data.SleepingClient.ValueBool())
	}
	if !data.SleepingClientTimeout.IsNull() {
		body, _ = sjson.Set(body, "sleepingClientTimeout", data.SleepingClientTimeout.ValueInt64())
	}
	if !data.AclName.IsNull() {
		body, _ = sjson.Set(body, "aclName", data.AclName.ValueString())
	}
	if !data.Posturing.IsNull() {
		body, _ = sjson.Set(body, "isPosturingEnabled", data.Posturing.ValueBool())
	}
	if !data.AuthKeySuiteB1x.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeySuiteB1x", data.AuthKeySuiteB1x.ValueBool())
	}
	if !data.AuthKeySuiteB1921x.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeySuiteB1921x", data.AuthKeySuiteB1921x.ValueBool())
	}
	if !data.AuthKeySaeExt.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeySaeExt", data.AuthKeySaeExt.ValueBool())
	}
	if !data.AuthKeySaeExtPlusFt.IsNull() {
		body, _ = sjson.Set(body, "isAuthKeySaeExtPlusFT", data.AuthKeySaeExtPlusFt.ValueBool())
	}
	if !data.ApBeaconProtection.IsNull() {
		body, _ = sjson.Set(body, "isApBeaconProtectionEnabled", data.ApBeaconProtection.ValueBool())
	}
	if !data.Ghz24Policy.IsNull() {
		body, _ = sjson.Set(body, "ghz24Policy", data.Ghz24Policy.ValueString())
	}
	if !data.CckmTsfTolerance.IsNull() {
		body, _ = sjson.Set(body, "cckmTsfTolerance", data.CckmTsfTolerance.ValueInt64())
	}
	if !data.Cckm.IsNull() {
		body, _ = sjson.Set(body, "isCckmEnabled", data.Cckm.ValueBool())
	}
	if !data.Hex.IsNull() {
		body, _ = sjson.Set(body, "isHex", data.Hex.ValueBool())
	}
	if !data.RandomMacFilter.IsNull() {
		body, _ = sjson.Set(body, "isRandomMacFilterEnabled", data.RandomMacFilter.ValueBool())
	}
	if !data.FastTransitionOverTheDistributedSystem.IsNull() {
		body, _ = sjson.Set(body, "fastTransitionOverTheDistributedSystemEnable", data.FastTransitionOverTheDistributedSystem.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessSSID) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("ssid"); value.Exists() {
		data.Ssid = types.StringValue(value.String())
	} else {
		data.Ssid = types.StringNull()
	}
	if value := res.Get("authType"); value.Exists() {
		data.AuthType = types.StringValue(value.String())
	} else {
		data.AuthType = types.StringNull()
	}
	if value := res.Get("isFastLaneEnabled"); value.Exists() {
		data.FastLane = types.BoolValue(value.Bool())
	} else {
		data.FastLane = types.BoolNull()
	}
	if value := res.Get("isMacFilteringEnabled"); value.Exists() {
		data.MacFiltering = types.BoolValue(value.Bool())
	} else {
		data.MacFiltering = types.BoolNull()
	}
	if value := res.Get("ssidRadioType"); value.Exists() {
		data.SsidRadioType = types.StringValue(value.String())
	} else {
		data.SsidRadioType = types.StringNull()
	}
	if value := res.Get("isBroadcastSSID"); value.Exists() {
		data.BroadcastSsid = types.BoolValue(value.Bool())
	} else {
		data.BroadcastSsid = types.BoolNull()
	}
	if value := res.Get("fastTransition"); value.Exists() {
		data.FastTransition = types.StringValue(value.String())
	} else {
		data.FastTransition = types.StringNull()
	}
	if value := res.Get("sessionTimeOutEnable"); value.Exists() {
		data.SessionTimeoutEnable = types.BoolValue(value.Bool())
	} else {
		data.SessionTimeoutEnable = types.BoolNull()
	}
	if value := res.Get("sessionTimeOut"); value.Exists() {
		data.SessionTimeout = types.Int64Value(value.Int())
	} else {
		data.SessionTimeout = types.Int64Null()
	}
	if value := res.Get("clientExclusionEnable"); value.Exists() {
		data.ClientExclusion = types.BoolValue(value.Bool())
	} else {
		data.ClientExclusion = types.BoolNull()
	}
	if value := res.Get("clientExclusionTimeout"); value.Exists() {
		data.ClientExclusionTimeout = types.Int64Value(value.Int())
	} else {
		data.ClientExclusionTimeout = types.Int64Null()
	}
	if value := res.Get("basicServiceSetMaxIdleEnable"); value.Exists() {
		data.BasicServiceSetMaxIdle = types.BoolValue(value.Bool())
	} else {
		data.BasicServiceSetMaxIdle = types.BoolNull()
	}
	if value := res.Get("basicServiceSetClientIdleTimeout"); value.Exists() {
		data.BasicServiceSetClientIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.BasicServiceSetClientIdleTimeout = types.Int64Null()
	}
	if value := res.Get("directedMulticastServiceEnable"); value.Exists() {
		data.DirectedMulticastService = types.BoolValue(value.Bool())
	} else {
		data.DirectedMulticastService = types.BoolNull()
	}
	if value := res.Get("neighborListEnable"); value.Exists() {
		data.NeighborList = types.BoolValue(value.Bool())
	} else {
		data.NeighborList = types.BoolNull()
	}
	if value := res.Get("managementFrameProtectionClientprotection"); value.Exists() {
		data.MftClientProtection = types.StringValue(value.String())
	} else {
		data.MftClientProtection = types.StringNull()
	}
	if value := res.Get("nasOptions"); value.Exists() && len(value.Array()) > 0 {
		data.NasOptions = helpers.GetStringSet(value.Array())
	} else {
		data.NasOptions = types.SetNull(types.StringType)
	}
	if value := res.Get("profileName"); value.Exists() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("aaaOverride"); value.Exists() {
		data.AaaOverride = types.BoolValue(value.Bool())
	} else {
		data.AaaOverride = types.BoolNull()
	}
	if value := res.Get("coverageHoleDetectionEnable"); value.Exists() {
		data.CoverageHoleDetection = types.BoolValue(value.Bool())
	} else {
		data.CoverageHoleDetection = types.BoolNull()
	}
	if value := res.Get("protectedManagementFrame"); value.Exists() {
		data.ProtectedManagementFrame = types.StringValue(value.String())
	} else {
		data.ProtectedManagementFrame = types.StringNull()
	}
	if value := res.Get("multiPSKSettings"); value.Exists() && len(value.Array()) > 0 {
		data.MultiPskSettings = make([]WirelessSSIDMultiPskSettings, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := WirelessSSIDMultiPskSettings{}
			if cValue := v.Get("priority"); cValue.Exists() {
				item.Priority = types.StringValue(cValue.String())
			} else {
				item.Priority = types.StringNull()
			}
			if cValue := v.Get("passphraseType"); cValue.Exists() {
				item.PassphraseType = types.StringValue(cValue.String())
			} else {
				item.PassphraseType = types.StringNull()
			}
			if cValue := v.Get("passphrase"); cValue.Exists() {
				item.Passphrase = types.StringValue(cValue.String())
			} else {
				item.Passphrase = types.StringNull()
			}
			data.MultiPskSettings = append(data.MultiPskSettings, item)
			return true
		})
	}
	if value := res.Get("clientRateLimit"); value.Exists() {
		data.ClientRateLimit = types.Int64Value(value.Int())
	} else {
		data.ClientRateLimit = types.Int64Null()
	}
	if value := res.Get("rsnCipherSuiteGcmp256"); value.Exists() {
		data.RsnCipherSuiteGcmp256 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteGcmp256 = types.BoolNull()
	}
	if value := res.Get("rsnCipherSuiteCcmp256"); value.Exists() {
		data.RsnCipherSuiteCcmp256 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteCcmp256 = types.BoolNull()
	}
	if value := res.Get("rsnCipherSuiteGcmp128"); value.Exists() {
		data.RsnCipherSuiteGcmp128 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteGcmp128 = types.BoolNull()
	}
	if value := res.Get("rsnCipherSuiteCcmp128"); value.Exists() {
		data.RsnCipherSuiteCcmp128 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteCcmp128 = types.BoolNull()
	}
	if value := res.Get("ghz6PolicyClientSteering"); value.Exists() {
		data.Ghz6PolicyClientSteering = types.BoolValue(value.Bool())
	} else {
		data.Ghz6PolicyClientSteering = types.BoolNull()
	}
	if value := res.Get("isAuthKey8021x"); value.Exists() {
		data.AuthKey8021x = types.BoolValue(value.Bool())
	} else {
		data.AuthKey8021x = types.BoolNull()
	}
	if value := res.Get("isAuthKey8021xPlusFT"); value.Exists() {
		data.AuthKey8021xPlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKey8021xPlusFt = types.BoolNull()
	}
	if value := res.Get("isAuthKey8021x_SHA256"); value.Exists() {
		data.AuthKey8021xSha256 = types.BoolValue(value.Bool())
	} else {
		data.AuthKey8021xSha256 = types.BoolNull()
	}
	if value := res.Get("isAuthKeySae"); value.Exists() {
		data.AuthKeySae = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySae = types.BoolNull()
	}
	if value := res.Get("isAuthKeySaePlusFT"); value.Exists() {
		data.AuthKeySaePlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySaePlusFt = types.BoolNull()
	}
	if value := res.Get("isAuthKeyPSK"); value.Exists() {
		data.AuthKeyPsk = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyPsk = types.BoolNull()
	}
	if value := res.Get("isAuthKeyPSKPlusFT"); value.Exists() {
		data.AuthKeyPskPlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyPskPlusFt = types.BoolNull()
	}
	if value := res.Get("isAuthKeyOWE"); value.Exists() {
		data.AuthKeyOwe = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyOwe = types.BoolNull()
	}
	if value := res.Get("isAuthKeyEasyPSK"); value.Exists() {
		data.AuthKeyEasyPsk = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyEasyPsk = types.BoolNull()
	}
	if value := res.Get("isAuthKeyPSKSHA256"); value.Exists() {
		data.AuthKeyEasyPskSha256 = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyEasyPskSha256 = types.BoolNull()
	}
	if value := res.Get("openSsid"); value.Exists() {
		data.OpenSsid = types.BoolValue(value.Bool())
	} else {
		data.OpenSsid = types.BoolNull()
	}
	if value := res.Get("wlanBandSelectEnable"); value.Exists() {
		data.WlanBandSelect = types.BoolValue(value.Bool())
	} else {
		data.WlanBandSelect = types.BoolNull()
	}
	if value := res.Get("isEnabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("authServers"); value.Exists() && len(value.Array()) > 0 {
		data.AuthServers = helpers.GetStringSet(value.Array())
	} else {
		data.AuthServers = types.SetNull(types.StringType)
	}
	if value := res.Get("acctServers"); value.Exists() && len(value.Array()) > 0 {
		data.AcctServers = helpers.GetStringSet(value.Array())
	} else {
		data.AcctServers = types.SetNull(types.StringType)
	}
	if value := res.Get("egressQos"); value.Exists() {
		data.EgressQos = types.StringValue(value.String())
	} else {
		data.EgressQos = types.StringNull()
	}
	if value := res.Get("ingressQos"); value.Exists() {
		data.IngressQos = types.StringValue(value.String())
	} else {
		data.IngressQos = types.StringNull()
	}
	if value := res.Get("wlanType"); value.Exists() {
		data.WlanType = types.StringValue(value.String())
	} else {
		data.WlanType = types.StringNull()
	}
	if value := res.Get("l3AuthType"); value.Exists() {
		data.L3AuthType = types.StringValue(value.String())
	} else {
		data.L3AuthType = types.StringNull()
	}
	if value := res.Get("authServer"); value.Exists() {
		data.AuthServer = types.StringValue(value.String())
	} else {
		data.AuthServer = types.StringNull()
	}
	if value := res.Get("externalAuthIpAddress"); value.Exists() {
		data.ExternalAuthIpAddress = types.StringValue(value.String())
	} else {
		data.ExternalAuthIpAddress = types.StringNull()
	}
	if value := res.Get("webPassthrough"); value.Exists() {
		data.WebPassthrough = types.BoolValue(value.Bool())
	} else {
		data.WebPassthrough = types.BoolNull()
	}
	if value := res.Get("sleepingClientEnable"); value.Exists() {
		data.SleepingClient = types.BoolValue(value.Bool())
	} else {
		data.SleepingClient = types.BoolNull()
	}
	if value := res.Get("sleepingClientTimeout"); value.Exists() {
		data.SleepingClientTimeout = types.Int64Value(value.Int())
	} else {
		data.SleepingClientTimeout = types.Int64Null()
	}
	if value := res.Get("aclName"); value.Exists() {
		data.AclName = types.StringValue(value.String())
	} else {
		data.AclName = types.StringNull()
	}
	if value := res.Get("isPosturingEnabled"); value.Exists() {
		data.Posturing = types.BoolValue(value.Bool())
	} else {
		data.Posturing = types.BoolNull()
	}
	if value := res.Get("isAuthKeySuiteB1x"); value.Exists() {
		data.AuthKeySuiteB1x = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySuiteB1x = types.BoolNull()
	}
	if value := res.Get("isAuthKeySuiteB1921x"); value.Exists() {
		data.AuthKeySuiteB1921x = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySuiteB1921x = types.BoolNull()
	}
	if value := res.Get("isAuthKeySaeExt"); value.Exists() {
		data.AuthKeySaeExt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySaeExt = types.BoolNull()
	}
	if value := res.Get("isAuthKeySaeExtPlusFT"); value.Exists() {
		data.AuthKeySaeExtPlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySaeExtPlusFt = types.BoolNull()
	}
	if value := res.Get("isApBeaconProtectionEnabled"); value.Exists() {
		data.ApBeaconProtection = types.BoolValue(value.Bool())
	} else {
		data.ApBeaconProtection = types.BoolNull()
	}
	if value := res.Get("ghz24Policy"); value.Exists() {
		data.Ghz24Policy = types.StringValue(value.String())
	} else {
		data.Ghz24Policy = types.StringNull()
	}
	if value := res.Get("cckmTsfTolerance"); value.Exists() {
		data.CckmTsfTolerance = types.Int64Value(value.Int())
	} else {
		data.CckmTsfTolerance = types.Int64Null()
	}
	if value := res.Get("isCckmEnabled"); value.Exists() {
		data.Cckm = types.BoolValue(value.Bool())
	} else {
		data.Cckm = types.BoolNull()
	}
	if value := res.Get("isHex"); value.Exists() {
		data.Hex = types.BoolValue(value.Bool())
	} else {
		data.Hex = types.BoolNull()
	}
	if value := res.Get("isRandomMacFilterEnabled"); value.Exists() {
		data.RandomMacFilter = types.BoolValue(value.Bool())
	} else {
		data.RandomMacFilter = types.BoolNull()
	}
	if value := res.Get("fastTransitionOverTheDistributedSystemEnable"); value.Exists() {
		data.FastTransitionOverTheDistributedSystem = types.BoolValue(value.Bool())
	} else {
		data.FastTransitionOverTheDistributedSystem = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessSSID) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("ssid"); value.Exists() && !data.Ssid.IsNull() {
		data.Ssid = types.StringValue(value.String())
	} else {
		data.Ssid = types.StringNull()
	}
	if value := res.Get("authType"); value.Exists() && !data.AuthType.IsNull() {
		data.AuthType = types.StringValue(value.String())
	} else {
		data.AuthType = types.StringNull()
	}
	if value := res.Get("isFastLaneEnabled"); value.Exists() && !data.FastLane.IsNull() {
		data.FastLane = types.BoolValue(value.Bool())
	} else {
		data.FastLane = types.BoolNull()
	}
	if value := res.Get("isMacFilteringEnabled"); value.Exists() && !data.MacFiltering.IsNull() {
		data.MacFiltering = types.BoolValue(value.Bool())
	} else {
		data.MacFiltering = types.BoolNull()
	}
	if value := res.Get("ssidRadioType"); value.Exists() && !data.SsidRadioType.IsNull() {
		data.SsidRadioType = types.StringValue(value.String())
	} else {
		data.SsidRadioType = types.StringNull()
	}
	if value := res.Get("isBroadcastSSID"); value.Exists() && !data.BroadcastSsid.IsNull() {
		data.BroadcastSsid = types.BoolValue(value.Bool())
	} else {
		data.BroadcastSsid = types.BoolNull()
	}
	if value := res.Get("fastTransition"); value.Exists() && !data.FastTransition.IsNull() {
		data.FastTransition = types.StringValue(value.String())
	} else {
		data.FastTransition = types.StringNull()
	}
	if value := res.Get("sessionTimeOutEnable"); value.Exists() && !data.SessionTimeoutEnable.IsNull() {
		data.SessionTimeoutEnable = types.BoolValue(value.Bool())
	} else {
		data.SessionTimeoutEnable = types.BoolNull()
	}
	if value := res.Get("sessionTimeOut"); value.Exists() && !data.SessionTimeout.IsNull() {
		data.SessionTimeout = types.Int64Value(value.Int())
	} else {
		data.SessionTimeout = types.Int64Null()
	}
	if value := res.Get("clientExclusionEnable"); value.Exists() && !data.ClientExclusion.IsNull() {
		data.ClientExclusion = types.BoolValue(value.Bool())
	} else {
		data.ClientExclusion = types.BoolNull()
	}
	if value := res.Get("clientExclusionTimeout"); value.Exists() && !data.ClientExclusionTimeout.IsNull() {
		data.ClientExclusionTimeout = types.Int64Value(value.Int())
	} else {
		data.ClientExclusionTimeout = types.Int64Null()
	}
	if value := res.Get("basicServiceSetMaxIdleEnable"); value.Exists() && !data.BasicServiceSetMaxIdle.IsNull() {
		data.BasicServiceSetMaxIdle = types.BoolValue(value.Bool())
	} else {
		data.BasicServiceSetMaxIdle = types.BoolNull()
	}
	if value := res.Get("basicServiceSetClientIdleTimeout"); value.Exists() && !data.BasicServiceSetClientIdleTimeout.IsNull() {
		data.BasicServiceSetClientIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.BasicServiceSetClientIdleTimeout = types.Int64Null()
	}
	if value := res.Get("directedMulticastServiceEnable"); value.Exists() && !data.DirectedMulticastService.IsNull() {
		data.DirectedMulticastService = types.BoolValue(value.Bool())
	} else {
		data.DirectedMulticastService = types.BoolNull()
	}
	if value := res.Get("neighborListEnable"); value.Exists() && !data.NeighborList.IsNull() {
		data.NeighborList = types.BoolValue(value.Bool())
	} else {
		data.NeighborList = types.BoolNull()
	}
	if value := res.Get("managementFrameProtectionClientprotection"); value.Exists() && !data.MftClientProtection.IsNull() {
		data.MftClientProtection = types.StringValue(value.String())
	} else {
		data.MftClientProtection = types.StringNull()
	}
	if value := res.Get("nasOptions"); value.Exists() && !data.NasOptions.IsNull() {
		data.NasOptions = helpers.GetStringSet(value.Array())
	} else {
		data.NasOptions = types.SetNull(types.StringType)
	}
	if value := res.Get("profileName"); value.Exists() && !data.ProfileName.IsNull() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("aaaOverride"); value.Exists() && !data.AaaOverride.IsNull() {
		data.AaaOverride = types.BoolValue(value.Bool())
	} else {
		data.AaaOverride = types.BoolNull()
	}
	if value := res.Get("coverageHoleDetectionEnable"); value.Exists() && !data.CoverageHoleDetection.IsNull() {
		data.CoverageHoleDetection = types.BoolValue(value.Bool())
	} else {
		data.CoverageHoleDetection = types.BoolNull()
	}
	if value := res.Get("protectedManagementFrame"); value.Exists() && !data.ProtectedManagementFrame.IsNull() {
		data.ProtectedManagementFrame = types.StringValue(value.String())
	} else {
		data.ProtectedManagementFrame = types.StringNull()
	}
	for i := range data.MultiPskSettings {
		keys := [...]string{"priority", "passphraseType", "passphrase"}
		keyValues := [...]string{data.MultiPskSettings[i].Priority.ValueString(), data.MultiPskSettings[i].PassphraseType.ValueString(), data.MultiPskSettings[i].Passphrase.ValueString()}

		var r gjson.Result
		res.Get("multiPSKSettings").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("priority"); value.Exists() && !data.MultiPskSettings[i].Priority.IsNull() {
			data.MultiPskSettings[i].Priority = types.StringValue(value.String())
		} else {
			data.MultiPskSettings[i].Priority = types.StringNull()
		}
		if value := r.Get("passphraseType"); value.Exists() && !data.MultiPskSettings[i].PassphraseType.IsNull() {
			data.MultiPskSettings[i].PassphraseType = types.StringValue(value.String())
		} else {
			data.MultiPskSettings[i].PassphraseType = types.StringNull()
		}
		if value := r.Get("passphrase"); value.Exists() && !data.MultiPskSettings[i].Passphrase.IsNull() {
			data.MultiPskSettings[i].Passphrase = types.StringValue(value.String())
		} else {
			data.MultiPskSettings[i].Passphrase = types.StringNull()
		}
	}
	if value := res.Get("clientRateLimit"); value.Exists() && !data.ClientRateLimit.IsNull() {
		data.ClientRateLimit = types.Int64Value(value.Int())
	} else {
		data.ClientRateLimit = types.Int64Null()
	}
	if value := res.Get("rsnCipherSuiteGcmp256"); value.Exists() && !data.RsnCipherSuiteGcmp256.IsNull() {
		data.RsnCipherSuiteGcmp256 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteGcmp256 = types.BoolNull()
	}
	if value := res.Get("rsnCipherSuiteCcmp256"); value.Exists() && !data.RsnCipherSuiteCcmp256.IsNull() {
		data.RsnCipherSuiteCcmp256 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteCcmp256 = types.BoolNull()
	}
	if value := res.Get("rsnCipherSuiteGcmp128"); value.Exists() && !data.RsnCipherSuiteGcmp128.IsNull() {
		data.RsnCipherSuiteGcmp128 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteGcmp128 = types.BoolNull()
	}
	if value := res.Get("rsnCipherSuiteCcmp128"); value.Exists() && !data.RsnCipherSuiteCcmp128.IsNull() {
		data.RsnCipherSuiteCcmp128 = types.BoolValue(value.Bool())
	} else {
		data.RsnCipherSuiteCcmp128 = types.BoolNull()
	}
	if value := res.Get("ghz6PolicyClientSteering"); value.Exists() && !data.Ghz6PolicyClientSteering.IsNull() {
		data.Ghz6PolicyClientSteering = types.BoolValue(value.Bool())
	} else {
		data.Ghz6PolicyClientSteering = types.BoolNull()
	}
	if value := res.Get("isAuthKey8021x"); value.Exists() && !data.AuthKey8021x.IsNull() {
		data.AuthKey8021x = types.BoolValue(value.Bool())
	} else {
		data.AuthKey8021x = types.BoolNull()
	}
	if value := res.Get("isAuthKey8021xPlusFT"); value.Exists() && !data.AuthKey8021xPlusFt.IsNull() {
		data.AuthKey8021xPlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKey8021xPlusFt = types.BoolNull()
	}
	if value := res.Get("isAuthKey8021x_SHA256"); value.Exists() && !data.AuthKey8021xSha256.IsNull() {
		data.AuthKey8021xSha256 = types.BoolValue(value.Bool())
	} else {
		data.AuthKey8021xSha256 = types.BoolNull()
	}
	if value := res.Get("isAuthKeySae"); value.Exists() && !data.AuthKeySae.IsNull() {
		data.AuthKeySae = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySae = types.BoolNull()
	}
	if value := res.Get("isAuthKeySaePlusFT"); value.Exists() && !data.AuthKeySaePlusFt.IsNull() {
		data.AuthKeySaePlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySaePlusFt = types.BoolNull()
	}
	if value := res.Get("isAuthKeyPSK"); value.Exists() && !data.AuthKeyPsk.IsNull() {
		data.AuthKeyPsk = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyPsk = types.BoolNull()
	}
	if value := res.Get("isAuthKeyPSKPlusFT"); value.Exists() && !data.AuthKeyPskPlusFt.IsNull() {
		data.AuthKeyPskPlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyPskPlusFt = types.BoolNull()
	}
	if value := res.Get("isAuthKeyOWE"); value.Exists() && !data.AuthKeyOwe.IsNull() {
		data.AuthKeyOwe = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyOwe = types.BoolNull()
	}
	if value := res.Get("isAuthKeyEasyPSK"); value.Exists() && !data.AuthKeyEasyPsk.IsNull() {
		data.AuthKeyEasyPsk = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyEasyPsk = types.BoolNull()
	}
	if value := res.Get("isAuthKeyPSKSHA256"); value.Exists() && !data.AuthKeyEasyPskSha256.IsNull() {
		data.AuthKeyEasyPskSha256 = types.BoolValue(value.Bool())
	} else {
		data.AuthKeyEasyPskSha256 = types.BoolNull()
	}
	if value := res.Get("openSsid"); value.Exists() && !data.OpenSsid.IsNull() {
		data.OpenSsid = types.BoolValue(value.Bool())
	} else {
		data.OpenSsid = types.BoolNull()
	}
	if value := res.Get("wlanBandSelectEnable"); value.Exists() && !data.WlanBandSelect.IsNull() {
		data.WlanBandSelect = types.BoolValue(value.Bool())
	} else {
		data.WlanBandSelect = types.BoolNull()
	}
	if value := res.Get("isEnabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("authServers"); value.Exists() && !data.AuthServers.IsNull() {
		data.AuthServers = helpers.GetStringSet(value.Array())
	} else {
		data.AuthServers = types.SetNull(types.StringType)
	}
	if value := res.Get("acctServers"); value.Exists() && !data.AcctServers.IsNull() {
		data.AcctServers = helpers.GetStringSet(value.Array())
	} else {
		data.AcctServers = types.SetNull(types.StringType)
	}
	if value := res.Get("egressQos"); value.Exists() && !data.EgressQos.IsNull() {
		data.EgressQos = types.StringValue(value.String())
	} else {
		data.EgressQos = types.StringNull()
	}
	if value := res.Get("ingressQos"); value.Exists() && !data.IngressQos.IsNull() {
		data.IngressQos = types.StringValue(value.String())
	} else {
		data.IngressQos = types.StringNull()
	}
	if value := res.Get("wlanType"); value.Exists() && !data.WlanType.IsNull() {
		data.WlanType = types.StringValue(value.String())
	} else {
		data.WlanType = types.StringNull()
	}
	if value := res.Get("l3AuthType"); value.Exists() && !data.L3AuthType.IsNull() {
		data.L3AuthType = types.StringValue(value.String())
	} else {
		data.L3AuthType = types.StringNull()
	}
	if value := res.Get("authServer"); value.Exists() && !data.AuthServer.IsNull() {
		data.AuthServer = types.StringValue(value.String())
	} else {
		data.AuthServer = types.StringNull()
	}
	if value := res.Get("externalAuthIpAddress"); value.Exists() && !data.ExternalAuthIpAddress.IsNull() {
		data.ExternalAuthIpAddress = types.StringValue(value.String())
	} else {
		data.ExternalAuthIpAddress = types.StringNull()
	}
	if value := res.Get("webPassthrough"); value.Exists() && !data.WebPassthrough.IsNull() {
		data.WebPassthrough = types.BoolValue(value.Bool())
	} else {
		data.WebPassthrough = types.BoolNull()
	}
	if value := res.Get("sleepingClientEnable"); value.Exists() && !data.SleepingClient.IsNull() {
		data.SleepingClient = types.BoolValue(value.Bool())
	} else {
		data.SleepingClient = types.BoolNull()
	}
	if value := res.Get("sleepingClientTimeout"); value.Exists() && !data.SleepingClientTimeout.IsNull() {
		data.SleepingClientTimeout = types.Int64Value(value.Int())
	} else {
		data.SleepingClientTimeout = types.Int64Null()
	}
	if value := res.Get("aclName"); value.Exists() && !data.AclName.IsNull() {
		data.AclName = types.StringValue(value.String())
	} else {
		data.AclName = types.StringNull()
	}
	if value := res.Get("isPosturingEnabled"); value.Exists() && !data.Posturing.IsNull() {
		data.Posturing = types.BoolValue(value.Bool())
	} else {
		data.Posturing = types.BoolNull()
	}
	if value := res.Get("isAuthKeySuiteB1x"); value.Exists() && !data.AuthKeySuiteB1x.IsNull() {
		data.AuthKeySuiteB1x = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySuiteB1x = types.BoolNull()
	}
	if value := res.Get("isAuthKeySuiteB1921x"); value.Exists() && !data.AuthKeySuiteB1921x.IsNull() {
		data.AuthKeySuiteB1921x = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySuiteB1921x = types.BoolNull()
	}
	if value := res.Get("isAuthKeySaeExt"); value.Exists() && !data.AuthKeySaeExt.IsNull() {
		data.AuthKeySaeExt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySaeExt = types.BoolNull()
	}
	if value := res.Get("isAuthKeySaeExtPlusFT"); value.Exists() && !data.AuthKeySaeExtPlusFt.IsNull() {
		data.AuthKeySaeExtPlusFt = types.BoolValue(value.Bool())
	} else {
		data.AuthKeySaeExtPlusFt = types.BoolNull()
	}
	if value := res.Get("isApBeaconProtectionEnabled"); value.Exists() && !data.ApBeaconProtection.IsNull() {
		data.ApBeaconProtection = types.BoolValue(value.Bool())
	} else {
		data.ApBeaconProtection = types.BoolNull()
	}
	if value := res.Get("ghz24Policy"); value.Exists() && !data.Ghz24Policy.IsNull() {
		data.Ghz24Policy = types.StringValue(value.String())
	} else {
		data.Ghz24Policy = types.StringNull()
	}
	if value := res.Get("cckmTsfTolerance"); value.Exists() && !data.CckmTsfTolerance.IsNull() {
		data.CckmTsfTolerance = types.Int64Value(value.Int())
	} else {
		data.CckmTsfTolerance = types.Int64Null()
	}
	if value := res.Get("isCckmEnabled"); value.Exists() && !data.Cckm.IsNull() {
		data.Cckm = types.BoolValue(value.Bool())
	} else {
		data.Cckm = types.BoolNull()
	}
	if value := res.Get("isHex"); value.Exists() && !data.Hex.IsNull() {
		data.Hex = types.BoolValue(value.Bool())
	} else {
		data.Hex = types.BoolNull()
	}
	if value := res.Get("isRandomMacFilterEnabled"); value.Exists() && !data.RandomMacFilter.IsNull() {
		data.RandomMacFilter = types.BoolValue(value.Bool())
	} else {
		data.RandomMacFilter = types.BoolNull()
	}
	if value := res.Get("fastTransitionOverTheDistributedSystemEnable"); value.Exists() && !data.FastTransitionOverTheDistributedSystem.IsNull() {
		data.FastTransitionOverTheDistributedSystem = types.BoolValue(value.Bool())
	} else {
		data.FastTransitionOverTheDistributedSystem = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessSSID) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Ssid.IsNull() {
		return false
	}
	if !data.AuthType.IsNull() {
		return false
	}
	if !data.Passphrase.IsNull() {
		return false
	}
	if !data.FastLane.IsNull() {
		return false
	}
	if !data.MacFiltering.IsNull() {
		return false
	}
	if !data.SsidRadioType.IsNull() {
		return false
	}
	if !data.BroadcastSsid.IsNull() {
		return false
	}
	if !data.FastTransition.IsNull() {
		return false
	}
	if !data.SessionTimeoutEnable.IsNull() {
		return false
	}
	if !data.SessionTimeout.IsNull() {
		return false
	}
	if !data.ClientExclusion.IsNull() {
		return false
	}
	if !data.ClientExclusionTimeout.IsNull() {
		return false
	}
	if !data.BasicServiceSetMaxIdle.IsNull() {
		return false
	}
	if !data.BasicServiceSetClientIdleTimeout.IsNull() {
		return false
	}
	if !data.DirectedMulticastService.IsNull() {
		return false
	}
	if !data.NeighborList.IsNull() {
		return false
	}
	if !data.MftClientProtection.IsNull() {
		return false
	}
	if !data.NasOptions.IsNull() {
		return false
	}
	if !data.ProfileName.IsNull() {
		return false
	}
	if !data.AaaOverride.IsNull() {
		return false
	}
	if !data.CoverageHoleDetection.IsNull() {
		return false
	}
	if !data.ProtectedManagementFrame.IsNull() {
		return false
	}
	if len(data.MultiPskSettings) > 0 {
		return false
	}
	if !data.ClientRateLimit.IsNull() {
		return false
	}
	if !data.RsnCipherSuiteGcmp256.IsNull() {
		return false
	}
	if !data.RsnCipherSuiteCcmp256.IsNull() {
		return false
	}
	if !data.RsnCipherSuiteGcmp128.IsNull() {
		return false
	}
	if !data.RsnCipherSuiteCcmp128.IsNull() {
		return false
	}
	if !data.Ghz6PolicyClientSteering.IsNull() {
		return false
	}
	if !data.AuthKey8021x.IsNull() {
		return false
	}
	if !data.AuthKey8021xPlusFt.IsNull() {
		return false
	}
	if !data.AuthKey8021xSha256.IsNull() {
		return false
	}
	if !data.AuthKeySae.IsNull() {
		return false
	}
	if !data.AuthKeySaePlusFt.IsNull() {
		return false
	}
	if !data.AuthKeyPsk.IsNull() {
		return false
	}
	if !data.AuthKeyPskPlusFt.IsNull() {
		return false
	}
	if !data.AuthKeyOwe.IsNull() {
		return false
	}
	if !data.AuthKeyEasyPsk.IsNull() {
		return false
	}
	if !data.AuthKeyEasyPskSha256.IsNull() {
		return false
	}
	if !data.OpenSsid.IsNull() {
		return false
	}
	if !data.WlanBandSelect.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.AuthServers.IsNull() {
		return false
	}
	if !data.AcctServers.IsNull() {
		return false
	}
	if !data.EgressQos.IsNull() {
		return false
	}
	if !data.IngressQos.IsNull() {
		return false
	}
	if !data.WlanType.IsNull() {
		return false
	}
	if !data.L3AuthType.IsNull() {
		return false
	}
	if !data.AuthServer.IsNull() {
		return false
	}
	if !data.ExternalAuthIpAddress.IsNull() {
		return false
	}
	if !data.WebPassthrough.IsNull() {
		return false
	}
	if !data.SleepingClient.IsNull() {
		return false
	}
	if !data.SleepingClientTimeout.IsNull() {
		return false
	}
	if !data.AclName.IsNull() {
		return false
	}
	if !data.Posturing.IsNull() {
		return false
	}
	if !data.AuthKeySuiteB1x.IsNull() {
		return false
	}
	if !data.AuthKeySuiteB1921x.IsNull() {
		return false
	}
	if !data.AuthKeySaeExt.IsNull() {
		return false
	}
	if !data.AuthKeySaeExtPlusFt.IsNull() {
		return false
	}
	if !data.ApBeaconProtection.IsNull() {
		return false
	}
	if !data.Ghz24Policy.IsNull() {
		return false
	}
	if !data.CckmTsfTolerance.IsNull() {
		return false
	}
	if !data.Cckm.IsNull() {
		return false
	}
	if !data.Hex.IsNull() {
		return false
	}
	if !data.RandomMacFilter.IsNull() {
		return false
	}
	if !data.FastTransitionOverTheDistributedSystem.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
