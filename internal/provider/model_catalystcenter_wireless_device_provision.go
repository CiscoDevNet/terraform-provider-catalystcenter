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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessDeviceProvision struct {
	Id                 types.String                               `tfsdk:"id"`
	DeviceName         types.String                               `tfsdk:"device_name"`
	NetworkDeviceId    types.String                               `tfsdk:"network_device_id"`
	Site               types.String                               `tfsdk:"site"`
	ManagedApLocations types.Set                                  `tfsdk:"managed_ap_locations"`
	DynamicInterfaces  []WirelessDeviceProvisionDynamicInterfaces `tfsdk:"dynamic_interfaces"`
}

type WirelessDeviceProvisionDynamicInterfaces struct {
	InterfaceIpAddress types.String `tfsdk:"interface_ip_address"`
	InterfaceNetmask   types.Int64  `tfsdk:"interface_netmask"`
	InterfaceGateway   types.String `tfsdk:"interface_gateway"`
	LagOrPortNumber    types.String `tfsdk:"lag_or_port_number"`
	VlanId             types.Int64  `tfsdk:"vlan_id"`
	InterfaceName      types.String `tfsdk:"interface_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessDeviceProvision) getPath() string {
	return "/dna/intent/api/v1/wireless/provision"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data WirelessDeviceProvision) getPathDelete() string {
	return "/dna/intent/api/v1/sda/provisionDevices"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessDeviceProvision) toBody(ctx context.Context, state WirelessDeviceProvision) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceName.IsNull() {
		body, _ = sjson.Set(body, "0.deviceName", data.DeviceName.ValueString())
	}
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "", data.NetworkDeviceId.ValueString())
	}
	if !data.Site.IsNull() {
		body, _ = sjson.Set(body, "0.site", data.Site.ValueString())
	}
	if !data.ManagedApLocations.IsNull() {
		var values []string
		data.ManagedApLocations.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.managedAPLocations", values)
	}
	if len(data.DynamicInterfaces) > 0 {
		body, _ = sjson.Set(body, "0.dynamicInterfaces", []interface{}{})
		for _, item := range data.DynamicInterfaces {
			itemBody := ""
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
				itemBody, _ = sjson.Set(itemBody, "lagOrPortNumber", item.LagOrPortNumber.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueInt64())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceName", item.InterfaceName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "0.dynamicInterfaces.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessDeviceProvision) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceName"); value.Exists() {
		data.DeviceName = types.StringValue(value.String())
	} else {
		data.DeviceName = types.StringNull()
	}
	if value := res.Get(""); value.Exists() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("0.site"); value.Exists() {
		data.Site = types.StringValue(value.String())
	} else {
		data.Site = types.StringNull()
	}
	if value := res.Get("0.managedAPLocations"); value.Exists() && len(value.Array()) > 0 {
		data.ManagedApLocations = helpers.GetStringSet(value.Array())
	} else {
		data.ManagedApLocations = types.SetNull(types.StringType)
	}
	if value := res.Get("0.dynamicInterfaces"); value.Exists() && len(value.Array()) > 0 {
		data.DynamicInterfaces = make([]WirelessDeviceProvisionDynamicInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := WirelessDeviceProvisionDynamicInterfaces{}
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
				item.LagOrPortNumber = types.StringValue(cValue.String())
			} else {
				item.LagOrPortNumber = types.StringNull()
			}
			if cValue := v.Get("vlanId"); cValue.Exists() {
				item.VlanId = types.Int64Value(cValue.Int())
			} else {
				item.VlanId = types.Int64Null()
			}
			if cValue := v.Get("interfaceName"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			} else {
				item.InterfaceName = types.StringNull()
			}
			data.DynamicInterfaces = append(data.DynamicInterfaces, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessDeviceProvision) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceName"); value.Exists() && !data.DeviceName.IsNull() {
		data.DeviceName = types.StringValue(value.String())
	} else {
		data.DeviceName = types.StringNull()
	}
	if value := res.Get(""); value.Exists() && !data.NetworkDeviceId.IsNull() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("0.site"); value.Exists() && !data.Site.IsNull() {
		data.Site = types.StringValue(value.String())
	} else {
		data.Site = types.StringNull()
	}
	if value := res.Get("0.managedAPLocations"); value.Exists() && !data.ManagedApLocations.IsNull() {
		data.ManagedApLocations = helpers.GetStringSet(value.Array())
	} else {
		data.ManagedApLocations = types.SetNull(types.StringType)
	}
	for i := range data.DynamicInterfaces {
		keys := [...]string{"interfaceIPAddress", "interfaceNetmaskInCIDR", "interfaceGateway", "lagOrPortNumber", "vlanId", "interfaceName"}
		keyValues := [...]string{data.DynamicInterfaces[i].InterfaceIpAddress.ValueString(), strconv.FormatInt(data.DynamicInterfaces[i].InterfaceNetmask.ValueInt64(), 10), data.DynamicInterfaces[i].InterfaceGateway.ValueString(), data.DynamicInterfaces[i].LagOrPortNumber.ValueString(), strconv.FormatInt(data.DynamicInterfaces[i].VlanId.ValueInt64(), 10), data.DynamicInterfaces[i].InterfaceName.ValueString()}

		var r gjson.Result
		res.Get("0.dynamicInterfaces").ForEach(
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
		if value := r.Get("interfaceIPAddress"); value.Exists() && !data.DynamicInterfaces[i].InterfaceIpAddress.IsNull() {
			data.DynamicInterfaces[i].InterfaceIpAddress = types.StringValue(value.String())
		} else {
			data.DynamicInterfaces[i].InterfaceIpAddress = types.StringNull()
		}
		if value := r.Get("interfaceNetmaskInCIDR"); value.Exists() && !data.DynamicInterfaces[i].InterfaceNetmask.IsNull() {
			data.DynamicInterfaces[i].InterfaceNetmask = types.Int64Value(value.Int())
		} else {
			data.DynamicInterfaces[i].InterfaceNetmask = types.Int64Null()
		}
		if value := r.Get("interfaceGateway"); value.Exists() && !data.DynamicInterfaces[i].InterfaceGateway.IsNull() {
			data.DynamicInterfaces[i].InterfaceGateway = types.StringValue(value.String())
		} else {
			data.DynamicInterfaces[i].InterfaceGateway = types.StringNull()
		}
		if value := r.Get("lagOrPortNumber"); value.Exists() && !data.DynamicInterfaces[i].LagOrPortNumber.IsNull() {
			data.DynamicInterfaces[i].LagOrPortNumber = types.StringValue(value.String())
		} else {
			data.DynamicInterfaces[i].LagOrPortNumber = types.StringNull()
		}
		if value := r.Get("vlanId"); value.Exists() && !data.DynamicInterfaces[i].VlanId.IsNull() {
			data.DynamicInterfaces[i].VlanId = types.Int64Value(value.Int())
		} else {
			data.DynamicInterfaces[i].VlanId = types.Int64Null()
		}
		if value := r.Get("interfaceName"); value.Exists() && !data.DynamicInterfaces[i].InterfaceName.IsNull() {
			data.DynamicInterfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.DynamicInterfaces[i].InterfaceName = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessDeviceProvision) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceName.IsNull() {
		return false
	}
	if !data.Site.IsNull() {
		return false
	}
	if !data.ManagedApLocations.IsNull() {
		return false
	}
	if len(data.DynamicInterfaces) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
