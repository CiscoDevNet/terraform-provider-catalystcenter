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
type WirelessEnterpriseSSID struct {
	Id                               types.String                             `tfsdk:"id"`
	Name                             types.String                             `tfsdk:"name"`
	SecurityLevel                    types.String                             `tfsdk:"security_level"`
	Passphrase                       types.String                             `tfsdk:"passphrase"`
	EnableFastLane                   types.Bool                               `tfsdk:"enable_fast_lane"`
	EnableMacFiltering               types.Bool                               `tfsdk:"enable_mac_filtering"`
	TrafficType                      types.String                             `tfsdk:"traffic_type"`
	RadioPolicy                      types.String                             `tfsdk:"radio_policy"`
	EnableBroadcastSsid              types.Bool                               `tfsdk:"enable_broadcast_ssid"`
	FastTransition                   types.String                             `tfsdk:"fast_transition"`
	EnableSessionTimeOut             types.Bool                               `tfsdk:"enable_session_time_out"`
	SessionTimeOut                   types.Int64                              `tfsdk:"session_time_out"`
	EnableClientExclusion            types.Bool                               `tfsdk:"enable_client_exclusion"`
	ClientExclusionTimeout           types.Int64                              `tfsdk:"client_exclusion_timeout"`
	EnableBasicServiceSetMaxIdle     types.Bool                               `tfsdk:"enable_basic_service_set_max_idle"`
	BasicServiceSetClientIdleTimeout types.Int64                              `tfsdk:"basic_service_set_client_idle_timeout"`
	EnableDirectedMulticastService   types.Bool                               `tfsdk:"enable_directed_multicast_service"`
	EnableNeighborList               types.Bool                               `tfsdk:"enable_neighbor_list"`
	MfpClientProtection              types.String                             `tfsdk:"mfp_client_protection"`
	NasOptions                       types.Set                                `tfsdk:"nas_options"`
	ProfileName                      types.String                             `tfsdk:"profile_name"`
	PolicyProfileName                types.String                             `tfsdk:"policy_profile_name"`
	AaaOverride                      types.Bool                               `tfsdk:"aaa_override"`
	CoverageHoleDetectionEnable      types.Bool                               `tfsdk:"coverage_hole_detection_enable"`
	ProtectedManagementFrame         types.String                             `tfsdk:"protected_management_frame"`
	MultiPskSettings                 []WirelessEnterpriseSSIDMultiPskSettings `tfsdk:"multi_psk_settings"`
	ClientRateLimit                  types.Int64                              `tfsdk:"client_rate_limit"`
}

type WirelessEnterpriseSSIDMultiPskSettings struct {
	Priority       types.String `tfsdk:"priority"`
	PassphraseType types.String `tfsdk:"passphrase_type"`
	Passphrase     types.String `tfsdk:"passphrase"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessEnterpriseSSID) getPath() string {
	return "/dna/intent/api/v1/enterprise-ssid"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessEnterpriseSSID) toBody(ctx context.Context, state WirelessEnterpriseSSID) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.SecurityLevel.IsNull() {
		body, _ = sjson.Set(body, "securityLevel", data.SecurityLevel.ValueString())
	}
	if !data.Passphrase.IsNull() {
		body, _ = sjson.Set(body, "passphrase", data.Passphrase.ValueString())
	}
	if !data.EnableFastLane.IsNull() {
		body, _ = sjson.Set(body, "enableFastLane", data.EnableFastLane.ValueBool())
	}
	if !data.EnableMacFiltering.IsNull() {
		body, _ = sjson.Set(body, "enableMACFiltering", data.EnableMacFiltering.ValueBool())
	}
	if !data.TrafficType.IsNull() {
		body, _ = sjson.Set(body, "trafficType", data.TrafficType.ValueString())
	}
	if !data.RadioPolicy.IsNull() {
		body, _ = sjson.Set(body, "radioPolicy", data.RadioPolicy.ValueString())
	}
	if !data.EnableBroadcastSsid.IsNull() {
		body, _ = sjson.Set(body, "enableBroadcastSSID", data.EnableBroadcastSsid.ValueBool())
	}
	if !data.FastTransition.IsNull() {
		body, _ = sjson.Set(body, "fastTransition", data.FastTransition.ValueString())
	}
	if !data.EnableSessionTimeOut.IsNull() {
		body, _ = sjson.Set(body, "enableSessionTimeOut", data.EnableSessionTimeOut.ValueBool())
	}
	if !data.SessionTimeOut.IsNull() {
		body, _ = sjson.Set(body, "sessionTimeOut", data.SessionTimeOut.ValueInt64())
	}
	if !data.EnableClientExclusion.IsNull() {
		body, _ = sjson.Set(body, "enableClientExclusion", data.EnableClientExclusion.ValueBool())
	}
	if !data.ClientExclusionTimeout.IsNull() {
		body, _ = sjson.Set(body, "clientExclusionTimeout", data.ClientExclusionTimeout.ValueInt64())
	}
	if !data.EnableBasicServiceSetMaxIdle.IsNull() {
		body, _ = sjson.Set(body, "enableBasicServiceSetMaxIdle", data.EnableBasicServiceSetMaxIdle.ValueBool())
	}
	if !data.BasicServiceSetClientIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, "basicServiceSetClientIdleTimeout", data.BasicServiceSetClientIdleTimeout.ValueInt64())
	}
	if !data.EnableDirectedMulticastService.IsNull() {
		body, _ = sjson.Set(body, "enableDirectedMulticastService", data.EnableDirectedMulticastService.ValueBool())
	}
	if !data.EnableNeighborList.IsNull() {
		body, _ = sjson.Set(body, "enableNeighborList", data.EnableNeighborList.ValueBool())
	}
	if !data.MfpClientProtection.IsNull() {
		body, _ = sjson.Set(body, "mfpClientProtection", data.MfpClientProtection.ValueString())
	}
	if !data.NasOptions.IsNull() {
		var values []string
		data.NasOptions.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "nasOptions", values)
	}
	if !data.ProfileName.IsNull() {
		body, _ = sjson.Set(body, "profileName", data.ProfileName.ValueString())
	}
	if !data.PolicyProfileName.IsNull() {
		body, _ = sjson.Set(body, "policyProfileName", data.PolicyProfileName.ValueString())
	}
	if !data.AaaOverride.IsNull() {
		body, _ = sjson.Set(body, "aaaOverride", data.AaaOverride.ValueBool())
	}
	if !data.CoverageHoleDetectionEnable.IsNull() {
		body, _ = sjson.Set(body, "coverageHoleDetectionEnable", data.CoverageHoleDetectionEnable.ValueBool())
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessEnterpriseSSID) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.ssidDetails.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.securityLevel"); value.Exists() {
		data.SecurityLevel = types.StringValue(value.String())
	} else {
		data.SecurityLevel = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.enableFastLane"); value.Exists() {
		data.EnableFastLane = types.BoolValue(value.Bool())
	} else {
		data.EnableFastLane = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.enableMACFiltering"); value.Exists() {
		data.EnableMacFiltering = types.BoolValue(value.Bool())
	} else {
		data.EnableMacFiltering = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.trafficType"); value.Exists() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.enableBroadcastSSID"); value.Exists() {
		data.EnableBroadcastSsid = types.BoolValue(value.Bool())
	} else {
		data.EnableBroadcastSsid = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.fastTransition"); value.Exists() {
		data.FastTransition = types.StringValue(value.String())
	} else {
		data.FastTransition = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.enableSessionTimeOut"); value.Exists() {
		data.EnableSessionTimeOut = types.BoolValue(value.Bool())
	} else {
		data.EnableSessionTimeOut = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.sessionTimeOut"); value.Exists() {
		data.SessionTimeOut = types.Int64Value(value.Int())
	} else {
		data.SessionTimeOut = types.Int64Null()
	}
	if value := res.Get("0.ssidDetails.0.enableClientExclusion"); value.Exists() {
		data.EnableClientExclusion = types.BoolValue(value.Bool())
	} else {
		data.EnableClientExclusion = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.clientExclusionTimeout"); value.Exists() {
		data.ClientExclusionTimeout = types.Int64Value(value.Int())
	} else {
		data.ClientExclusionTimeout = types.Int64Null()
	}
	if value := res.Get("0.ssidDetails.0.enableBasicServiceSetMaxIdle"); value.Exists() {
		data.EnableBasicServiceSetMaxIdle = types.BoolValue(value.Bool())
	} else {
		data.EnableBasicServiceSetMaxIdle = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.basicServiceSetClientIdleTimeout"); value.Exists() {
		data.BasicServiceSetClientIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.BasicServiceSetClientIdleTimeout = types.Int64Null()
	}
	if value := res.Get("0.ssidDetails.0.enableDirectedMulticastService"); value.Exists() {
		data.EnableDirectedMulticastService = types.BoolValue(value.Bool())
	} else {
		data.EnableDirectedMulticastService = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.enableNeighborList"); value.Exists() {
		data.EnableNeighborList = types.BoolValue(value.Bool())
	} else {
		data.EnableNeighborList = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.mfpClientProtection"); value.Exists() {
		data.MfpClientProtection = types.StringValue(value.String())
	} else {
		data.MfpClientProtection = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.nasOptions"); value.Exists() && len(value.Array()) > 0 {
		data.NasOptions = helpers.GetStringSet(value.Array())
	} else {
		data.NasOptions = types.SetNull(types.StringType)
	}
	if value := res.Get("0.ssidDetails.0.profileName"); value.Exists() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.policyProfileName"); value.Exists() {
		data.PolicyProfileName = types.StringValue(value.String())
	} else {
		data.PolicyProfileName = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.aaaOverride"); value.Exists() {
		data.AaaOverride = types.BoolValue(value.Bool())
	} else {
		data.AaaOverride = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.coverageHoleDetectionEnable"); value.Exists() {
		data.CoverageHoleDetectionEnable = types.BoolValue(value.Bool())
	} else {
		data.CoverageHoleDetectionEnable = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.protectedManagementFrame"); value.Exists() {
		data.ProtectedManagementFrame = types.StringValue(value.String())
	} else {
		data.ProtectedManagementFrame = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.multiPSKSettings"); value.Exists() && len(value.Array()) > 0 {
		data.MultiPskSettings = make([]WirelessEnterpriseSSIDMultiPskSettings, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := WirelessEnterpriseSSIDMultiPskSettings{}
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
	if value := res.Get("0.ssidDetails.0.clientRateLimit"); value.Exists() {
		data.ClientRateLimit = types.Int64Value(value.Int())
	} else {
		data.ClientRateLimit = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessEnterpriseSSID) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.ssidDetails.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.securityLevel"); value.Exists() && !data.SecurityLevel.IsNull() {
		data.SecurityLevel = types.StringValue(value.String())
	} else {
		data.SecurityLevel = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.enableFastLane"); value.Exists() && !data.EnableFastLane.IsNull() {
		data.EnableFastLane = types.BoolValue(value.Bool())
	} else {
		data.EnableFastLane = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.enableMACFiltering"); value.Exists() && !data.EnableMacFiltering.IsNull() {
		data.EnableMacFiltering = types.BoolValue(value.Bool())
	} else {
		data.EnableMacFiltering = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.trafficType"); value.Exists() && !data.TrafficType.IsNull() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.enableBroadcastSSID"); value.Exists() && !data.EnableBroadcastSsid.IsNull() {
		data.EnableBroadcastSsid = types.BoolValue(value.Bool())
	} else {
		data.EnableBroadcastSsid = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.fastTransition"); value.Exists() && !data.FastTransition.IsNull() {
		data.FastTransition = types.StringValue(value.String())
	} else {
		data.FastTransition = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.enableSessionTimeOut"); value.Exists() && !data.EnableSessionTimeOut.IsNull() {
		data.EnableSessionTimeOut = types.BoolValue(value.Bool())
	} else {
		data.EnableSessionTimeOut = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.sessionTimeOut"); value.Exists() && !data.SessionTimeOut.IsNull() {
		data.SessionTimeOut = types.Int64Value(value.Int())
	} else {
		data.SessionTimeOut = types.Int64Null()
	}
	if value := res.Get("0.ssidDetails.0.enableClientExclusion"); value.Exists() && !data.EnableClientExclusion.IsNull() {
		data.EnableClientExclusion = types.BoolValue(value.Bool())
	} else {
		data.EnableClientExclusion = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.clientExclusionTimeout"); value.Exists() && !data.ClientExclusionTimeout.IsNull() {
		data.ClientExclusionTimeout = types.Int64Value(value.Int())
	} else {
		data.ClientExclusionTimeout = types.Int64Null()
	}
	if value := res.Get("0.ssidDetails.0.enableBasicServiceSetMaxIdle"); value.Exists() && !data.EnableBasicServiceSetMaxIdle.IsNull() {
		data.EnableBasicServiceSetMaxIdle = types.BoolValue(value.Bool())
	} else {
		data.EnableBasicServiceSetMaxIdle = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.basicServiceSetClientIdleTimeout"); value.Exists() && !data.BasicServiceSetClientIdleTimeout.IsNull() {
		data.BasicServiceSetClientIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.BasicServiceSetClientIdleTimeout = types.Int64Null()
	}
	if value := res.Get("0.ssidDetails.0.enableDirectedMulticastService"); value.Exists() && !data.EnableDirectedMulticastService.IsNull() {
		data.EnableDirectedMulticastService = types.BoolValue(value.Bool())
	} else {
		data.EnableDirectedMulticastService = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.enableNeighborList"); value.Exists() && !data.EnableNeighborList.IsNull() {
		data.EnableNeighborList = types.BoolValue(value.Bool())
	} else {
		data.EnableNeighborList = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.mfpClientProtection"); value.Exists() && !data.MfpClientProtection.IsNull() {
		data.MfpClientProtection = types.StringValue(value.String())
	} else {
		data.MfpClientProtection = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.nasOptions"); value.Exists() && !data.NasOptions.IsNull() {
		data.NasOptions = helpers.GetStringSet(value.Array())
	} else {
		data.NasOptions = types.SetNull(types.StringType)
	}
	if value := res.Get("0.ssidDetails.0.profileName"); value.Exists() && !data.ProfileName.IsNull() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.policyProfileName"); value.Exists() && !data.PolicyProfileName.IsNull() {
		data.PolicyProfileName = types.StringValue(value.String())
	} else {
		data.PolicyProfileName = types.StringNull()
	}
	if value := res.Get("0.ssidDetails.0.aaaOverride"); value.Exists() && !data.AaaOverride.IsNull() {
		data.AaaOverride = types.BoolValue(value.Bool())
	} else {
		data.AaaOverride = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.coverageHoleDetectionEnable"); value.Exists() && !data.CoverageHoleDetectionEnable.IsNull() {
		data.CoverageHoleDetectionEnable = types.BoolValue(value.Bool())
	} else {
		data.CoverageHoleDetectionEnable = types.BoolNull()
	}
	if value := res.Get("0.ssidDetails.0.protectedManagementFrame"); value.Exists() && !data.ProtectedManagementFrame.IsNull() {
		data.ProtectedManagementFrame = types.StringValue(value.String())
	} else {
		data.ProtectedManagementFrame = types.StringNull()
	}
	for i := range data.MultiPskSettings {
		keys := [...]string{"priority", "passphraseType", "passphrase"}
		keyValues := [...]string{data.MultiPskSettings[i].Priority.ValueString(), data.MultiPskSettings[i].PassphraseType.ValueString(), data.MultiPskSettings[i].Passphrase.ValueString()}

		var r gjson.Result
		res.Get("0.ssidDetails.0.multiPSKSettings").ForEach(
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
	if value := res.Get("0.ssidDetails.0.clientRateLimit"); value.Exists() && !data.ClientRateLimit.IsNull() {
		data.ClientRateLimit = types.Int64Value(value.Int())
	} else {
		data.ClientRateLimit = types.Int64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessEnterpriseSSID) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.SecurityLevel.IsNull() {
		return false
	}
	if !data.Passphrase.IsNull() {
		return false
	}
	if !data.EnableFastLane.IsNull() {
		return false
	}
	if !data.EnableMacFiltering.IsNull() {
		return false
	}
	if !data.TrafficType.IsNull() {
		return false
	}
	if !data.RadioPolicy.IsNull() {
		return false
	}
	if !data.EnableBroadcastSsid.IsNull() {
		return false
	}
	if !data.FastTransition.IsNull() {
		return false
	}
	if !data.EnableSessionTimeOut.IsNull() {
		return false
	}
	if !data.SessionTimeOut.IsNull() {
		return false
	}
	if !data.EnableClientExclusion.IsNull() {
		return false
	}
	if !data.ClientExclusionTimeout.IsNull() {
		return false
	}
	if !data.EnableBasicServiceSetMaxIdle.IsNull() {
		return false
	}
	if !data.BasicServiceSetClientIdleTimeout.IsNull() {
		return false
	}
	if !data.EnableDirectedMulticastService.IsNull() {
		return false
	}
	if !data.EnableNeighborList.IsNull() {
		return false
	}
	if !data.MfpClientProtection.IsNull() {
		return false
	}
	if !data.NasOptions.IsNull() {
		return false
	}
	if !data.ProfileName.IsNull() {
		return false
	}
	if !data.PolicyProfileName.IsNull() {
		return false
	}
	if !data.AaaOverride.IsNull() {
		return false
	}
	if !data.CoverageHoleDetectionEnable.IsNull() {
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
	return true
}

// End of section. //template:end isNull
