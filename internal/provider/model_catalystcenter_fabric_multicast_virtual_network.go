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
type FabricMulticastVirtualNetwork struct {
	Id                 types.String                                `tfsdk:"id"`
	FabricId           types.String                                `tfsdk:"fabric_id"`
	VirtualNetworkName types.String                                `tfsdk:"virtual_network_name"`
	IpPoolName         types.String                                `tfsdk:"ip_pool_name"`
	Ipv4SsmRanges      types.Set                                   `tfsdk:"ipv4_ssm_ranges"`
	MulticastRps       []FabricMulticastVirtualNetworkMulticastRps `tfsdk:"multicast_rps"`
}

type FabricMulticastVirtualNetworkMulticastRps struct {
	RpDeviceLocation types.String `tfsdk:"rp_device_location"`
	Ipv4Address      types.String `tfsdk:"ipv4_address"`
	Ipv6Address      types.String `tfsdk:"ipv6_address"`
	IsDefaultV4Rp    types.Bool   `tfsdk:"is_default_v4_rp"`
	IsDefaultV6Rp    types.Bool   `tfsdk:"is_default_v6_rp"`
	NetworkDeviceIds types.Set    `tfsdk:"network_device_ids"`
	Ipv4AsmRanges    types.Set    `tfsdk:"ipv4_asm_ranges"`
	Ipv6AsmRanges    types.Set    `tfsdk:"ipv6_asm_ranges"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricMulticastVirtualNetwork) getPath() string {
	return "/dna/intent/api/v1/sda/multicast/virtualNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath
func (data FabricMulticastVirtualNetwork) getFallbackPath() string {
	return ""
}

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricMulticastVirtualNetwork) toBody(ctx context.Context, state FabricMulticastVirtualNetwork) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.IpPoolName.IsNull() {
		body, _ = sjson.Set(body, "0.ipPoolName", data.IpPoolName.ValueString())
	}
	if !data.Ipv4SsmRanges.IsNull() {
		var values []string
		data.Ipv4SsmRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.ipv4SsmRanges", values)
	}
	if len(data.MulticastRps) > 0 {
		body, _ = sjson.Set(body, "0.multicastRPs", []interface{}{})
		for _, item := range data.MulticastRps {
			itemBody := ""
			if !item.RpDeviceLocation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rpDeviceLocation", item.RpDeviceLocation.ValueString())
			}
			if !item.Ipv4Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4Address", item.Ipv4Address.ValueString())
			}
			if !item.Ipv6Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6Address", item.Ipv6Address.ValueString())
			}
			if !item.IsDefaultV4Rp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isDefaultV4RP", item.IsDefaultV4Rp.ValueBool())
			}
			if !item.IsDefaultV6Rp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isDefaultV6RP", item.IsDefaultV6Rp.ValueBool())
			}
			if !item.NetworkDeviceIds.IsNull() {
				var values []string
				item.NetworkDeviceIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "networkDeviceIds", values)
			}
			if !item.Ipv4AsmRanges.IsNull() {
				var values []string
				item.Ipv4AsmRanges.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipv4AsmRanges", values)
			}
			if !item.Ipv6AsmRanges.IsNull() {
				var values []string
				item.Ipv6AsmRanges.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipv6AsmRanges", values)
			}
			body, _ = sjson.SetRaw(body, "0.multicastRPs.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricMulticastVirtualNetwork) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.ipPoolName"); value.Exists() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
	if value := res.Get("response.0.ipv4SsmRanges"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4SsmRanges = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv4SsmRanges = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.multicastRPs"); value.Exists() && len(value.Array()) > 0 {
		data.MulticastRps = make([]FabricMulticastVirtualNetworkMulticastRps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricMulticastVirtualNetworkMulticastRps{}
			if cValue := v.Get("rpDeviceLocation"); cValue.Exists() {
				item.RpDeviceLocation = types.StringValue(cValue.String())
			} else {
				item.RpDeviceLocation = types.StringNull()
			}
			if cValue := v.Get("ipv4Address"); cValue.Exists() {
				item.Ipv4Address = types.StringValue(cValue.String())
			} else {
				item.Ipv4Address = types.StringNull()
			}
			if cValue := v.Get("ipv6Address"); cValue.Exists() {
				item.Ipv6Address = types.StringValue(cValue.String())
			} else {
				item.Ipv6Address = types.StringNull()
			}
			if cValue := v.Get("isDefaultV4RP"); cValue.Exists() {
				item.IsDefaultV4Rp = types.BoolValue(cValue.Bool())
			} else {
				item.IsDefaultV4Rp = types.BoolNull()
			}
			if cValue := v.Get("isDefaultV6RP"); cValue.Exists() {
				item.IsDefaultV6Rp = types.BoolValue(cValue.Bool())
			} else {
				item.IsDefaultV6Rp = types.BoolNull()
			}
			if cValue := v.Get("networkDeviceIds"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.NetworkDeviceIds = helpers.GetStringSet(cValue.Array())
			} else {
				item.NetworkDeviceIds = types.SetNull(types.StringType)
			}
			if cValue := v.Get("ipv4AsmRanges"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv4AsmRanges = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ipv4AsmRanges = types.SetNull(types.StringType)
			}
			if cValue := v.Get("ipv6AsmRanges"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv6AsmRanges = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ipv6AsmRanges = types.SetNull(types.StringType)
			}
			data.MulticastRps = append(data.MulticastRps, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricMulticastVirtualNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.ipPoolName"); value.Exists() && !data.IpPoolName.IsNull() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
	if value := res.Get("response.0.ipv4SsmRanges"); value.Exists() && !data.Ipv4SsmRanges.IsNull() {
		data.Ipv4SsmRanges = helpers.GetStringSet(value.Array())
	} else {
		data.Ipv4SsmRanges = types.SetNull(types.StringType)
	}
	for i := range data.MulticastRps {
		keys := [...]string{"rpDeviceLocation", "ipv4Address", "ipv6Address", "isDefaultV4RP", "isDefaultV6RP"}
		keyValues := [...]string{data.MulticastRps[i].RpDeviceLocation.ValueString(), data.MulticastRps[i].Ipv4Address.ValueString(), data.MulticastRps[i].Ipv6Address.ValueString(), strconv.FormatBool(data.MulticastRps[i].IsDefaultV4Rp.ValueBool()), strconv.FormatBool(data.MulticastRps[i].IsDefaultV6Rp.ValueBool())}

		var r gjson.Result
		res.Get("response.0.multicastRPs").ForEach(
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
		if value := r.Get("rpDeviceLocation"); value.Exists() && !data.MulticastRps[i].RpDeviceLocation.IsNull() {
			data.MulticastRps[i].RpDeviceLocation = types.StringValue(value.String())
		} else {
			data.MulticastRps[i].RpDeviceLocation = types.StringNull()
		}
		if value := r.Get("ipv4Address"); value.Exists() && !data.MulticastRps[i].Ipv4Address.IsNull() {
			data.MulticastRps[i].Ipv4Address = types.StringValue(value.String())
		} else {
			data.MulticastRps[i].Ipv4Address = types.StringNull()
		}
		if value := r.Get("ipv6Address"); value.Exists() && !data.MulticastRps[i].Ipv6Address.IsNull() {
			data.MulticastRps[i].Ipv6Address = types.StringValue(value.String())
		} else {
			data.MulticastRps[i].Ipv6Address = types.StringNull()
		}
		if value := r.Get("isDefaultV4RP"); value.Exists() && !data.MulticastRps[i].IsDefaultV4Rp.IsNull() {
			data.MulticastRps[i].IsDefaultV4Rp = types.BoolValue(value.Bool())
		} else {
			data.MulticastRps[i].IsDefaultV4Rp = types.BoolNull()
		}
		if value := r.Get("isDefaultV6RP"); value.Exists() && !data.MulticastRps[i].IsDefaultV6Rp.IsNull() {
			data.MulticastRps[i].IsDefaultV6Rp = types.BoolValue(value.Bool())
		} else {
			data.MulticastRps[i].IsDefaultV6Rp = types.BoolNull()
		}
		if value := r.Get("networkDeviceIds"); value.Exists() && !data.MulticastRps[i].NetworkDeviceIds.IsNull() {
			data.MulticastRps[i].NetworkDeviceIds = helpers.GetStringSet(value.Array())
		} else {
			data.MulticastRps[i].NetworkDeviceIds = types.SetNull(types.StringType)
		}
		if value := r.Get("ipv4AsmRanges"); value.Exists() && !data.MulticastRps[i].Ipv4AsmRanges.IsNull() {
			data.MulticastRps[i].Ipv4AsmRanges = helpers.GetStringSet(value.Array())
		} else {
			data.MulticastRps[i].Ipv4AsmRanges = types.SetNull(types.StringType)
		}
		if value := r.Get("ipv6AsmRanges"); value.Exists() && !data.MulticastRps[i].Ipv6AsmRanges.IsNull() {
			data.MulticastRps[i].Ipv6AsmRanges = helpers.GetStringSet(value.Array())
		} else {
			data.MulticastRps[i].Ipv6AsmRanges = types.SetNull(types.StringType)
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricMulticastVirtualNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.IpPoolName.IsNull() {
		return false
	}
	if !data.Ipv4SsmRanges.IsNull() {
		return false
	}
	if len(data.MulticastRps) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
