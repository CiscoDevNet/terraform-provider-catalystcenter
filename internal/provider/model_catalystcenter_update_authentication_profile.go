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
type UpdateAuthenticationProfile struct {
	Id                        types.String                                           `tfsdk:"id"`
	FabricId                  types.String                                           `tfsdk:"fabric_id"`
	AuthenticationProfileName types.String                                           `tfsdk:"authentication_profile_name"`
	AuthenticationOrder       types.String                                           `tfsdk:"authentication_order"`
	Dot1xToMabFallbackTimeout types.Int64                                            `tfsdk:"dot1x_to_mab_fallback_timeout"`
	WakeOnLan                 types.Bool                                             `tfsdk:"wake_on_lan"`
	NumberOfHosts             types.String                                           `tfsdk:"number_of_hosts"`
	IsBpduGuardEnabled        types.Bool                                             `tfsdk:"is_bpdu_guard_enabled"`
	PreAuthAclEnabled         types.Bool                                             `tfsdk:"pre_auth_acl_enabled"`
	PreAuthAclImplicitAction  types.String                                           `tfsdk:"pre_auth_acl_implicit_action"`
	PreAuthAclDescription     types.String                                           `tfsdk:"pre_auth_acl_description"`
	PreAuthAclAccessContracts []UpdateAuthenticationProfilePreAuthAclAccessContracts `tfsdk:"pre_auth_acl_access_contracts"`
}

type UpdateAuthenticationProfilePreAuthAclAccessContracts struct {
	Action   types.String `tfsdk:"action"`
	Protocol types.String `tfsdk:"protocol"`
	Port     types.String `tfsdk:"port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data UpdateAuthenticationProfile) getPath() string {
	return "/dna/intent/api/v1/sda/authenticationProfiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

func (data UpdateAuthenticationProfile) toBody(ctx context.Context, state UpdateAuthenticationProfile) string {
	body := ""
	put := false
	if !data.Id.IsNull() {
		put = true
		body, _ = sjson.Set(body, "0.id", data.Id.ValueString())
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.AuthenticationProfileName.IsNull() {
		body, _ = sjson.Set(body, "0.authenticationProfileName", data.AuthenticationProfileName.ValueString())
	}
	if !data.AuthenticationOrder.IsNull() {
		body, _ = sjson.Set(body, "0.authenticationOrder", data.AuthenticationOrder.ValueString())
	}
	if !data.Dot1xToMabFallbackTimeout.IsNull() {
		body, _ = sjson.Set(body, "0.dot1xToMabFallbackTimeout", data.Dot1xToMabFallbackTimeout.ValueInt64())
	}
	if !data.WakeOnLan.IsNull() {
		body, _ = sjson.Set(body, "0.wakeOnLan", data.WakeOnLan.ValueBool())
	}
	if !data.NumberOfHosts.IsNull() {
		body, _ = sjson.Set(body, "0.numberOfHosts", data.NumberOfHosts.ValueString())
	}
	if !data.IsBpduGuardEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.isBpduGuardEnabled", data.IsBpduGuardEnabled.ValueBool())
	}
	if !data.PreAuthAclEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.preAuthAcl.enabled", data.PreAuthAclEnabled.ValueBool())
	}
	if !data.PreAuthAclImplicitAction.IsNull() {
		body, _ = sjson.Set(body, "0.preAuthAcl.implicitAction", data.PreAuthAclImplicitAction.ValueString())
	}
	if !data.PreAuthAclDescription.IsNull() {
		body, _ = sjson.Set(body, "0.preAuthAcl.description", data.PreAuthAclDescription.ValueString())
	}
	if len(data.PreAuthAclAccessContracts) > 0 {
		body, _ = sjson.Set(body, "0.preAuthAcl.accessContracts", []interface{}{})
		for _, item := range data.PreAuthAclAccessContracts {
			itemBody := ""
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueString())
			}
			body, _ = sjson.SetRaw(body, "0.preAuthAcl.accessContracts.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *UpdateAuthenticationProfile) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.authenticationProfileName"); value.Exists() {
		data.AuthenticationProfileName = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileName = types.StringNull()
	}
	if value := res.Get("response.0.authenticationOrder"); value.Exists() {
		data.AuthenticationOrder = types.StringValue(value.String())
	} else {
		data.AuthenticationOrder = types.StringNull()
	}
	if value := res.Get("response.0.dot1xToMabFallbackTimeout"); value.Exists() {
		data.Dot1xToMabFallbackTimeout = types.Int64Value(value.Int())
	} else {
		data.Dot1xToMabFallbackTimeout = types.Int64Null()
	}
	if value := res.Get("response.0.wakeOnLan"); value.Exists() {
		data.WakeOnLan = types.BoolValue(value.Bool())
	} else {
		data.WakeOnLan = types.BoolNull()
	}
	if value := res.Get("response.0.numberOfHosts"); value.Exists() {
		data.NumberOfHosts = types.StringValue(value.String())
	} else {
		data.NumberOfHosts = types.StringNull()
	}
	if value := res.Get("response.0.isBpduGuardEnabled"); value.Exists() {
		data.IsBpduGuardEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsBpduGuardEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.preAuthAcl.enabled"); value.Exists() {
		data.PreAuthAclEnabled = types.BoolValue(value.Bool())
	} else {
		data.PreAuthAclEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.preAuthAcl.implicitAction"); value.Exists() {
		data.PreAuthAclImplicitAction = types.StringValue(value.String())
	} else {
		data.PreAuthAclImplicitAction = types.StringNull()
	}
	if value := res.Get("response.0.preAuthAcl.description"); value.Exists() {
		data.PreAuthAclDescription = types.StringValue(value.String())
	} else {
		data.PreAuthAclDescription = types.StringNull()
	}
	if value := res.Get("response.0.preAuthAcl.accessContracts"); value.Exists() && len(value.Array()) > 0 {
		data.PreAuthAclAccessContracts = make([]UpdateAuthenticationProfilePreAuthAclAccessContracts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := UpdateAuthenticationProfilePreAuthAclAccessContracts{}
			if cValue := v.Get("action"); cValue.Exists() {
				item.Action = types.StringValue(cValue.String())
			} else {
				item.Action = types.StringNull()
			}
			if cValue := v.Get("protocol"); cValue.Exists() {
				item.Protocol = types.StringValue(cValue.String())
			} else {
				item.Protocol = types.StringNull()
			}
			if cValue := v.Get("port"); cValue.Exists() {
				item.Port = types.StringValue(cValue.String())
			} else {
				item.Port = types.StringNull()
			}
			data.PreAuthAclAccessContracts = append(data.PreAuthAclAccessContracts, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *UpdateAuthenticationProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.authenticationProfileName"); value.Exists() && !data.AuthenticationProfileName.IsNull() {
		data.AuthenticationProfileName = types.StringValue(value.String())
	} else {
		data.AuthenticationProfileName = types.StringNull()
	}
	if value := res.Get("response.0.authenticationOrder"); value.Exists() && !data.AuthenticationOrder.IsNull() {
		data.AuthenticationOrder = types.StringValue(value.String())
	} else {
		data.AuthenticationOrder = types.StringNull()
	}
	if value := res.Get("response.0.dot1xToMabFallbackTimeout"); value.Exists() && !data.Dot1xToMabFallbackTimeout.IsNull() {
		data.Dot1xToMabFallbackTimeout = types.Int64Value(value.Int())
	} else {
		data.Dot1xToMabFallbackTimeout = types.Int64Null()
	}
	if value := res.Get("response.0.wakeOnLan"); value.Exists() && !data.WakeOnLan.IsNull() {
		data.WakeOnLan = types.BoolValue(value.Bool())
	} else {
		data.WakeOnLan = types.BoolNull()
	}
	if value := res.Get("response.0.numberOfHosts"); value.Exists() && !data.NumberOfHosts.IsNull() {
		data.NumberOfHosts = types.StringValue(value.String())
	} else {
		data.NumberOfHosts = types.StringNull()
	}
	if value := res.Get("response.0.isBpduGuardEnabled"); value.Exists() && !data.IsBpduGuardEnabled.IsNull() {
		data.IsBpduGuardEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsBpduGuardEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.preAuthAcl.enabled"); value.Exists() && !data.PreAuthAclEnabled.IsNull() {
		data.PreAuthAclEnabled = types.BoolValue(value.Bool())
	} else {
		data.PreAuthAclEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.preAuthAcl.implicitAction"); value.Exists() && !data.PreAuthAclImplicitAction.IsNull() {
		data.PreAuthAclImplicitAction = types.StringValue(value.String())
	} else {
		data.PreAuthAclImplicitAction = types.StringNull()
	}
	if value := res.Get("response.0.preAuthAcl.description"); value.Exists() && !data.PreAuthAclDescription.IsNull() {
		data.PreAuthAclDescription = types.StringValue(value.String())
	} else {
		data.PreAuthAclDescription = types.StringNull()
	}
	for i := range data.PreAuthAclAccessContracts {
		keys := [...]string{"action", "protocol", "port"}
		keyValues := [...]string{data.PreAuthAclAccessContracts[i].Action.ValueString(), data.PreAuthAclAccessContracts[i].Protocol.ValueString(), data.PreAuthAclAccessContracts[i].Port.ValueString()}

		var r gjson.Result
		res.Get("response.0.preAuthAcl.accessContracts").ForEach(
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
		if value := r.Get("action"); value.Exists() && !data.PreAuthAclAccessContracts[i].Action.IsNull() {
			data.PreAuthAclAccessContracts[i].Action = types.StringValue(value.String())
		} else {
			data.PreAuthAclAccessContracts[i].Action = types.StringNull()
		}
		if value := r.Get("protocol"); value.Exists() && !data.PreAuthAclAccessContracts[i].Protocol.IsNull() {
			data.PreAuthAclAccessContracts[i].Protocol = types.StringValue(value.String())
		} else {
			data.PreAuthAclAccessContracts[i].Protocol = types.StringNull()
		}
		if value := r.Get("port"); value.Exists() && !data.PreAuthAclAccessContracts[i].Port.IsNull() {
			data.PreAuthAclAccessContracts[i].Port = types.StringValue(value.String())
		} else {
			data.PreAuthAclAccessContracts[i].Port = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *UpdateAuthenticationProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.AuthenticationOrder.IsNull() {
		return false
	}
	if !data.Dot1xToMabFallbackTimeout.IsNull() {
		return false
	}
	if !data.WakeOnLan.IsNull() {
		return false
	}
	if !data.NumberOfHosts.IsNull() {
		return false
	}
	if !data.IsBpduGuardEnabled.IsNull() {
		return false
	}
	if !data.PreAuthAclEnabled.IsNull() {
		return false
	}
	if !data.PreAuthAclImplicitAction.IsNull() {
		return false
	}
	if !data.PreAuthAclDescription.IsNull() {
		return false
	}
	if len(data.PreAuthAclAccessContracts) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
