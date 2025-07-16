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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessDeviceProvision struct {
	Id                                  types.String                        `tfsdk:"id"`
	NetworkDeviceId                     types.String                        `tfsdk:"network_device_id"`
	ApRebootPercentage                  types.Int64                         `tfsdk:"ap_reboot_percentage"`
	EnableRollingApUpgrade              types.Bool                          `tfsdk:"enable_rolling_ap_upgrade"`
	SkipApProvision                     types.Bool                          `tfsdk:"skip_ap_provision"`
	ApAuthorizationListName             types.String                        `tfsdk:"ap_authorization_list_name"`
	AuthorizeMeshAndNonMeshAccessPoints types.Bool                          `tfsdk:"authorize_mesh_and_non_mesh_access_points"`
	Reprovision                         types.Bool                          `tfsdk:"reprovision"`
	Interfaces                          []WirelessDeviceProvisionInterfaces `tfsdk:"interfaces"`
}

type WirelessDeviceProvisionInterfaces struct {
	InterfaceName      types.String `tfsdk:"interface_name"`
	VlanId             types.Int64  `tfsdk:"vlan_id"`
	InterfaceIpAddress types.String `tfsdk:"interface_ip_address"`
	InterfaceNetmask   types.Int64  `tfsdk:"interface_netmask"`
	InterfaceGateway   types.String `tfsdk:"interface_gateway"`
	LagOrPortNumber    types.Int64  `tfsdk:"lag_or_port_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessDeviceProvision) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessControllers/%v/provision", url.QueryEscape(data.NetworkDeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data WirelessDeviceProvision) getPathDelete() string {
	return "/dna/intent/api/v1/sda/provisionDevices"
}

// End of section. //template:end getPathDelete

func (data WirelessDeviceProvision) getPathUpdate() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessControllers/%v/provision", url.QueryEscape(data.NetworkDeviceId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessDeviceProvision) toBody(ctx context.Context, state WirelessDeviceProvision) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.ApRebootPercentage.IsNull() {
		body, _ = sjson.Set(body, "rollingApUpgrade.apRebootPercentage", data.ApRebootPercentage.ValueInt64())
	}
	if !data.EnableRollingApUpgrade.IsNull() {
		body, _ = sjson.Set(body, "rollingApUpgrade.enableRollingApUpgrade", data.EnableRollingApUpgrade.ValueBool())
	}
	if !data.SkipApProvision.IsNull() {
		body, _ = sjson.Set(body, "skipApProvision", data.SkipApProvision.ValueBool())
	}
	if !data.ApAuthorizationListName.IsNull() {
		body, _ = sjson.Set(body, "apAuthorizationListName", data.ApAuthorizationListName.ValueString())
	}
	if !data.AuthorizeMeshAndNonMeshAccessPoints.IsNull() {
		body, _ = sjson.Set(body, "authorizeMeshAndNonMeshAccessPoints", data.AuthorizeMeshAndNonMeshAccessPoints.ValueBool())
	}
	if !data.Reprovision.IsNull() {
		body, _ = sjson.Set(body, "", data.Reprovision.ValueBool())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceName", item.InterfaceName.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueInt64())
			}
			if !item.InterfaceIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceIPAddress", item.InterfaceIpAddress.ValueString())
			}
			if !item.InterfaceNetmask.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceNetmaskInCIDR", item.InterfaceNetmask.ValueInt64())
			}
			if !item.InterfaceGateway.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceGateway", item.InterfaceGateway.ValueString())
			}
			if !item.LagOrPortNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lagOrPortNumber", item.LagOrPortNumber.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "interfaces.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessDeviceProvision) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("rollingApUpgrade.apRebootPercentage"); value.Exists() {
		data.ApRebootPercentage = types.Int64Value(value.Int())
	} else {
		data.ApRebootPercentage = types.Int64Null()
	}
	if value := res.Get("rollingApUpgrade.enableRollingApUpgrade"); value.Exists() {
		data.EnableRollingApUpgrade = types.BoolValue(value.Bool())
	} else {
		data.EnableRollingApUpgrade = types.BoolNull()
	}
	if value := res.Get("skipApProvision"); value.Exists() {
		data.SkipApProvision = types.BoolValue(value.Bool())
	} else {
		data.SkipApProvision = types.BoolNull()
	}
	if value := res.Get("apAuthorizationListName"); value.Exists() {
		data.ApAuthorizationListName = types.StringValue(value.String())
	} else {
		data.ApAuthorizationListName = types.StringNull()
	}
	if value := res.Get("authorizeMeshAndNonMeshAccessPoints"); value.Exists() {
		data.AuthorizeMeshAndNonMeshAccessPoints = types.BoolValue(value.Bool())
	} else {
		data.AuthorizeMeshAndNonMeshAccessPoints = types.BoolNull()
	}
	if value := res.Get(""); value.Exists() {
		data.Reprovision = types.BoolValue(value.Bool())
	} else {
		data.Reprovision = types.BoolNull()
	}
	if value := res.Get("interfaces"); value.Exists() && len(value.Array()) > 0 {
		data.Interfaces = make([]WirelessDeviceProvisionInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := WirelessDeviceProvisionInterfaces{}
			if cValue := v.Get("interfaceName"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			} else {
				item.InterfaceName = types.StringNull()
			}
			if cValue := v.Get("vlanId"); cValue.Exists() {
				item.VlanId = types.Int64Value(cValue.Int())
			} else {
				item.VlanId = types.Int64Null()
			}
			if cValue := v.Get("interfaceIPAddress"); cValue.Exists() {
				item.InterfaceIpAddress = types.StringValue(cValue.String())
			} else {
				item.InterfaceIpAddress = types.StringNull()
			}
			if cValue := v.Get("interfaceNetmaskInCIDR"); cValue.Exists() {
				item.InterfaceNetmask = types.Int64Value(cValue.Int())
			} else {
				item.InterfaceNetmask = types.Int64Null()
			}
			if cValue := v.Get("interfaceGateway"); cValue.Exists() {
				item.InterfaceGateway = types.StringValue(cValue.String())
			} else {
				item.InterfaceGateway = types.StringNull()
			}
			if cValue := v.Get("lagOrPortNumber"); cValue.Exists() {
				item.LagOrPortNumber = types.Int64Value(cValue.Int())
			} else {
				item.LagOrPortNumber = types.Int64Null()
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessDeviceProvision) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("rollingApUpgrade.apRebootPercentage"); value.Exists() && !data.ApRebootPercentage.IsNull() {
		data.ApRebootPercentage = types.Int64Value(value.Int())
	} else {
		data.ApRebootPercentage = types.Int64Null()
	}
	if value := res.Get("rollingApUpgrade.enableRollingApUpgrade"); value.Exists() && !data.EnableRollingApUpgrade.IsNull() {
		data.EnableRollingApUpgrade = types.BoolValue(value.Bool())
	} else {
		data.EnableRollingApUpgrade = types.BoolNull()
	}
	if value := res.Get("skipApProvision"); value.Exists() && !data.SkipApProvision.IsNull() {
		data.SkipApProvision = types.BoolValue(value.Bool())
	} else {
		data.SkipApProvision = types.BoolNull()
	}
	if value := res.Get("apAuthorizationListName"); value.Exists() && !data.ApAuthorizationListName.IsNull() {
		data.ApAuthorizationListName = types.StringValue(value.String())
	} else {
		data.ApAuthorizationListName = types.StringNull()
	}
	if value := res.Get("authorizeMeshAndNonMeshAccessPoints"); value.Exists() && !data.AuthorizeMeshAndNonMeshAccessPoints.IsNull() {
		data.AuthorizeMeshAndNonMeshAccessPoints = types.BoolValue(value.Bool())
	} else {
		data.AuthorizeMeshAndNonMeshAccessPoints = types.BoolNull()
	}
	if value := res.Get(""); value.Exists() && !data.Reprovision.IsNull() {
		data.Reprovision = types.BoolValue(value.Bool())
	} else {
		data.Reprovision = types.BoolNull()
	}
	for i := range data.Interfaces {
		keys := [...]string{"interfaceName", "vlanId", "interfaceIPAddress", "interfaceNetmaskInCIDR", "interfaceGateway", "lagOrPortNumber"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString(), strconv.FormatInt(data.Interfaces[i].VlanId.ValueInt64(), 10), data.Interfaces[i].InterfaceIpAddress.ValueString(), strconv.FormatInt(data.Interfaces[i].InterfaceNetmask.ValueInt64(), 10), data.Interfaces[i].InterfaceGateway.ValueString(), strconv.FormatInt(data.Interfaces[i].LagOrPortNumber.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("interfaces").ForEach(
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
		if value := r.Get("interfaceName"); value.Exists() && !data.Interfaces[i].InterfaceName.IsNull() {
			data.Interfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.Interfaces[i].InterfaceName = types.StringNull()
		}
		if value := r.Get("vlanId"); value.Exists() && !data.Interfaces[i].VlanId.IsNull() {
			data.Interfaces[i].VlanId = types.Int64Value(value.Int())
		} else {
			data.Interfaces[i].VlanId = types.Int64Null()
		}
		if value := r.Get("interfaceIPAddress"); value.Exists() && !data.Interfaces[i].InterfaceIpAddress.IsNull() {
			data.Interfaces[i].InterfaceIpAddress = types.StringValue(value.String())
		} else {
			data.Interfaces[i].InterfaceIpAddress = types.StringNull()
		}
		if value := r.Get("interfaceNetmaskInCIDR"); value.Exists() && !data.Interfaces[i].InterfaceNetmask.IsNull() {
			data.Interfaces[i].InterfaceNetmask = types.Int64Value(value.Int())
		} else {
			data.Interfaces[i].InterfaceNetmask = types.Int64Null()
		}
		if value := r.Get("interfaceGateway"); value.Exists() && !data.Interfaces[i].InterfaceGateway.IsNull() {
			data.Interfaces[i].InterfaceGateway = types.StringValue(value.String())
		} else {
			data.Interfaces[i].InterfaceGateway = types.StringNull()
		}
		if value := r.Get("lagOrPortNumber"); value.Exists() && !data.Interfaces[i].LagOrPortNumber.IsNull() {
			data.Interfaces[i].LagOrPortNumber = types.Int64Value(value.Int())
		} else {
			data.Interfaces[i].LagOrPortNumber = types.Int64Null()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessDeviceProvision) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.ApRebootPercentage.IsNull() {
		return false
	}
	if !data.EnableRollingApUpgrade.IsNull() {
		return false
	}
	if !data.SkipApProvision.IsNull() {
		return false
	}
	if !data.ApAuthorizationListName.IsNull() {
		return false
	}
	if !data.AuthorizeMeshAndNonMeshAccessPoints.IsNull() {
		return false
	}
	if !data.Reprovision.IsNull() {
		return false
	}
	if len(data.Interfaces) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
