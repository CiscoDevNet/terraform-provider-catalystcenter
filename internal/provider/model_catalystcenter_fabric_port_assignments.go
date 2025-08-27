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
type FabricPortAssignments struct {
	Id              types.String                           `tfsdk:"id"`
	FabricId        types.String                           `tfsdk:"fabric_id"`
	NetworkDeviceId types.String                           `tfsdk:"network_device_id"`
	PortAssignments []FabricPortAssignmentsPortAssignments `tfsdk:"port_assignments"`
}

type FabricPortAssignmentsPortAssignments struct {
	Id                       types.String `tfsdk:"id"`
	FabricId                 types.String `tfsdk:"fabric_id"`
	NetworkDeviceId          types.String `tfsdk:"network_device_id"`
	InterfaceName            types.String `tfsdk:"interface_name"`
	ConnectedDeviceType      types.String `tfsdk:"connected_device_type"`
	DataVlanName             types.String `tfsdk:"data_vlan_name"`
	VoiceVlanName            types.String `tfsdk:"voice_vlan_name"`
	AuthenticateTemplateName types.String `tfsdk:"authenticate_template_name"`
	SecurityGroupName        types.String `tfsdk:"security_group_name"`
	InterfaceDescription     types.String `tfsdk:"interface_description"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricPortAssignments) getPath() string {
	return "/dna/intent/api/v1/sda/portAssignments"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricPortAssignments) toBody(ctx context.Context, state FabricPortAssignments) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "", data.FabricId.ValueString())
	}
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "", data.NetworkDeviceId.ValueString())
	}
	if len(data.PortAssignments) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.PortAssignments {
			itemBody := ""
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.FabricId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fabricId", item.FabricId.ValueString())
			}
			if !item.NetworkDeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkDeviceId", item.NetworkDeviceId.ValueString())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceName", item.InterfaceName.ValueString())
			}
			if !item.ConnectedDeviceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "connectedDeviceType", item.ConnectedDeviceType.ValueString())
			}
			if !item.DataVlanName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dataVlanName", item.DataVlanName.ValueString())
			}
			if !item.VoiceVlanName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "voiceVlanName", item.VoiceVlanName.ValueString())
			}
			if !item.AuthenticateTemplateName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authenticateTemplateName", item.AuthenticateTemplateName.ValueString())
			}
			if !item.SecurityGroupName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "securityGroupName", item.SecurityGroupName.ValueString())
			}
			if !item.InterfaceDescription.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceDescription", item.InterfaceDescription.ValueString())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricPortAssignments) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.PortAssignments = make([]FabricPortAssignmentsPortAssignments, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricPortAssignmentsPortAssignments{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("fabricId"); cValue.Exists() {
				item.FabricId = types.StringValue(cValue.String())
			} else {
				item.FabricId = types.StringNull()
			}
			if cValue := v.Get("networkDeviceId"); cValue.Exists() {
				item.NetworkDeviceId = types.StringValue(cValue.String())
			} else {
				item.NetworkDeviceId = types.StringNull()
			}
			if cValue := v.Get("interfaceName"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			} else {
				item.InterfaceName = types.StringNull()
			}
			if cValue := v.Get("connectedDeviceType"); cValue.Exists() {
				item.ConnectedDeviceType = types.StringValue(cValue.String())
			} else {
				item.ConnectedDeviceType = types.StringNull()
			}
			if cValue := v.Get("dataVlanName"); cValue.Exists() {
				item.DataVlanName = types.StringValue(cValue.String())
			} else {
				item.DataVlanName = types.StringNull()
			}
			if cValue := v.Get("voiceVlanName"); cValue.Exists() {
				item.VoiceVlanName = types.StringValue(cValue.String())
			} else {
				item.VoiceVlanName = types.StringNull()
			}
			if cValue := v.Get("authenticateTemplateName"); cValue.Exists() {
				item.AuthenticateTemplateName = types.StringValue(cValue.String())
			} else {
				item.AuthenticateTemplateName = types.StringNull()
			}
			if cValue := v.Get("securityGroupName"); cValue.Exists() {
				item.SecurityGroupName = types.StringValue(cValue.String())
			} else {
				item.SecurityGroupName = types.StringNull()
			}
			if cValue := v.Get("interfaceDescription"); cValue.Exists() {
				item.InterfaceDescription = types.StringValue(cValue.String())
			} else {
				item.InterfaceDescription = types.StringNull()
			}
			data.PortAssignments = append(data.PortAssignments, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricPortAssignments) updateFromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.PortAssignments {
		keys := [...]string{"interfaceName"}
		keyValues := [...]string{data.PortAssignments[i].InterfaceName.ValueString()}

		var r gjson.Result
		res.ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.PortAssignments[i].Id.IsNull() {
			data.PortAssignments[i].Id = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].Id = types.StringNull()
		}
		if value := r.Get("fabricId"); value.Exists() && !data.PortAssignments[i].FabricId.IsNull() {
			data.PortAssignments[i].FabricId = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].FabricId = types.StringNull()
		}
		if value := r.Get("networkDeviceId"); value.Exists() && !data.PortAssignments[i].NetworkDeviceId.IsNull() {
			data.PortAssignments[i].NetworkDeviceId = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].NetworkDeviceId = types.StringNull()
		}
		if value := r.Get("interfaceName"); value.Exists() && !data.PortAssignments[i].InterfaceName.IsNull() {
			data.PortAssignments[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("connectedDeviceType"); value.Exists() && !data.PortAssignments[i].ConnectedDeviceType.IsNull() {
			data.PortAssignments[i].ConnectedDeviceType = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].ConnectedDeviceType = types.StringNull()
		}
		if value := r.Get("dataVlanName"); value.Exists() && !data.PortAssignments[i].DataVlanName.IsNull() {
			data.PortAssignments[i].DataVlanName = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].DataVlanName = types.StringNull()
		}
		if value := r.Get("voiceVlanName"); value.Exists() && !data.PortAssignments[i].VoiceVlanName.IsNull() {
			data.PortAssignments[i].VoiceVlanName = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].VoiceVlanName = types.StringNull()
		}
		if value := r.Get("authenticateTemplateName"); value.Exists() && !data.PortAssignments[i].AuthenticateTemplateName.IsNull() {
			data.PortAssignments[i].AuthenticateTemplateName = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].AuthenticateTemplateName = types.StringNull()
		}
		if value := r.Get("securityGroupName"); value.Exists() && !data.PortAssignments[i].SecurityGroupName.IsNull() {
			data.PortAssignments[i].SecurityGroupName = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].SecurityGroupName = types.StringNull()
		}
		if value := r.Get("interfaceDescription"); value.Exists() && !data.PortAssignments[i].InterfaceDescription.IsNull() {
			data.PortAssignments[i].InterfaceDescription = types.StringValue(value.String())
		} else {
			data.PortAssignments[i].InterfaceDescription = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FabricPortAssignments) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.PortAssignments {
		keys := [...]string{"interfaceName"}
		keyValues := [...]string{data.PortAssignments[i].InterfaceName.ValueString()}

		var r gjson.Result
		res.ForEach(
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
		if data.PortAssignments[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.PortAssignments[i].Id.IsNull() {
				data.PortAssignments[i].Id = types.StringValue(value.String())
			} else {
				data.PortAssignments[i].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricPortAssignments) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.PortAssignments) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
