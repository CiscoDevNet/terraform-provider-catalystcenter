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
type FabricMulticastVirtualNetworks struct {
	Id              types.String                                    `tfsdk:"id"`
	FabricId        types.String                                    `tfsdk:"fabric_id"`
	VirtualNetworks []FabricMulticastVirtualNetworksVirtualNetworks `tfsdk:"virtual_networks"`
}

type FabricMulticastVirtualNetworksVirtualNetworks struct {
	Id                 types.String                                                `tfsdk:"id"`
	FabricId           types.String                                                `tfsdk:"fabric_id"`
	VirtualNetworkName types.String                                                `tfsdk:"virtual_network_name"`
	IpPoolName         types.String                                                `tfsdk:"ip_pool_name"`
	Ipv4SsmRanges      types.Set                                                   `tfsdk:"ipv4_ssm_ranges"`
	MulticastRps       []FabricMulticastVirtualNetworksVirtualNetworksMulticastRps `tfsdk:"multicast_rps"`
}

type FabricMulticastVirtualNetworksVirtualNetworksMulticastRps struct {
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
func (data FabricMulticastVirtualNetworks) getPath() string {
	return "/dna/intent/api/v1/sda/multicast/virtualNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath
func (data FabricMulticastVirtualNetworks) getFallbackPath() string {
	return ""
}

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricMulticastVirtualNetworks) toBody(ctx context.Context, state FabricMulticastVirtualNetworks) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "fabricId", data.FabricId.ValueString())
	}
	if len(data.VirtualNetworks) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.VirtualNetworks {
			itemBody := ""
			if item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.FabricId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fabricId", item.FabricId.ValueString())
			}
			if !item.VirtualNetworkName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "virtualNetworkName", item.VirtualNetworkName.ValueString())
			}
			if !item.IpPoolName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipPoolName", item.IpPoolName.ValueString())
			}
			if !item.Ipv4SsmRanges.IsNull() {
				var values []string
				item.Ipv4SsmRanges.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipv4SsmRanges", values)
			}
			if len(item.MulticastRps) > 0 {
				itemBody, _ = sjson.Set(itemBody, "multicastRPs", []interface{}{})
				for _, childItem := range item.MulticastRps {
					itemChildBody := ""
					if !childItem.RpDeviceLocation.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "rpDeviceLocation", childItem.RpDeviceLocation.ValueString())
					}
					if !childItem.Ipv4Address.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ipv4Address", childItem.Ipv4Address.ValueString())
					}
					if !childItem.Ipv6Address.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ipv6Address", childItem.Ipv6Address.ValueString())
					}
					if !childItem.IsDefaultV4Rp.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "isDefaultV4RP", childItem.IsDefaultV4Rp.ValueBool())
					}
					if !childItem.IsDefaultV6Rp.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "isDefaultV6RP", childItem.IsDefaultV6Rp.ValueBool())
					}
					if !childItem.NetworkDeviceIds.IsNull() {
						var values []string
						childItem.NetworkDeviceIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "networkDeviceIds", values)
					}
					if !childItem.Ipv4AsmRanges.IsNull() {
						var values []string
						childItem.Ipv4AsmRanges.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "ipv4AsmRanges", values)
					}
					if !childItem.Ipv6AsmRanges.IsNull() {
						var values []string
						childItem.Ipv6AsmRanges.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "ipv6AsmRanges", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "multicastRPs.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricMulticastVirtualNetworks) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.VirtualNetworks = make([]FabricMulticastVirtualNetworksVirtualNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricMulticastVirtualNetworksVirtualNetworks{}
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
			if cValue := v.Get("virtualNetworkName"); cValue.Exists() {
				item.VirtualNetworkName = types.StringValue(cValue.String())
			} else {
				item.VirtualNetworkName = types.StringNull()
			}
			if cValue := v.Get("ipPoolName"); cValue.Exists() {
				item.IpPoolName = types.StringValue(cValue.String())
			} else {
				item.IpPoolName = types.StringNull()
			}
			if cValue := v.Get("ipv4SsmRanges"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv4SsmRanges = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ipv4SsmRanges = types.SetNull(types.StringType)
			}
			if cValue := v.Get("multicastRPs"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MulticastRps = make([]FabricMulticastVirtualNetworksVirtualNetworksMulticastRps, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := FabricMulticastVirtualNetworksVirtualNetworksMulticastRps{}
					if ccValue := cv.Get("rpDeviceLocation"); ccValue.Exists() {
						cItem.RpDeviceLocation = types.StringValue(ccValue.String())
					} else {
						cItem.RpDeviceLocation = types.StringNull()
					}
					if ccValue := cv.Get("ipv4Address"); ccValue.Exists() {
						cItem.Ipv4Address = types.StringValue(ccValue.String())
					} else {
						cItem.Ipv4Address = types.StringNull()
					}
					if ccValue := cv.Get("ipv6Address"); ccValue.Exists() {
						cItem.Ipv6Address = types.StringValue(ccValue.String())
					} else {
						cItem.Ipv6Address = types.StringNull()
					}
					if ccValue := cv.Get("isDefaultV4RP"); ccValue.Exists() {
						cItem.IsDefaultV4Rp = types.BoolValue(ccValue.Bool())
					} else {
						cItem.IsDefaultV4Rp = types.BoolNull()
					}
					if ccValue := cv.Get("isDefaultV6RP"); ccValue.Exists() {
						cItem.IsDefaultV6Rp = types.BoolValue(ccValue.Bool())
					} else {
						cItem.IsDefaultV6Rp = types.BoolNull()
					}
					if ccValue := cv.Get("networkDeviceIds"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.NetworkDeviceIds = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.NetworkDeviceIds = types.SetNull(types.StringType)
					}
					if ccValue := cv.Get("ipv4AsmRanges"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.Ipv4AsmRanges = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.Ipv4AsmRanges = types.SetNull(types.StringType)
					}
					if ccValue := cv.Get("ipv6AsmRanges"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.Ipv6AsmRanges = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.Ipv6AsmRanges = types.SetNull(types.StringType)
					}
					item.MulticastRps = append(item.MulticastRps, cItem)
					return true
				})
			}
			data.VirtualNetworks = append(data.VirtualNetworks, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricMulticastVirtualNetworks) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []FabricMulticastVirtualNetworksVirtualNetworks

	res = res.Get("response")
	for i := range data.VirtualNetworks {
		keys := [...]string{"virtualNetworkName"}
		keyValues := [...]string{data.VirtualNetworks[i].VirtualNetworkName.ValueString()}

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
		if value := r.Get("id"); value.Exists() && !data.VirtualNetworks[i].Id.IsNull() {
			data.VirtualNetworks[i].Id = types.StringValue(value.String())
		} else {
			data.VirtualNetworks[i].Id = types.StringNull()
		}
		if value := r.Get("fabricId"); value.Exists() && !data.VirtualNetworks[i].FabricId.IsNull() {
			data.VirtualNetworks[i].FabricId = types.StringValue(value.String())
		} else {
			data.VirtualNetworks[i].FabricId = types.StringNull()
		}
		if value := r.Get("virtualNetworkName"); value.Exists() && !data.VirtualNetworks[i].VirtualNetworkName.IsNull() {
			data.VirtualNetworks[i].VirtualNetworkName = types.StringValue(value.String())
		} else {
			data.VirtualNetworks[i].VirtualNetworkName = types.StringNull()
		}
		if value := r.Get("ipPoolName"); value.Exists() && !data.VirtualNetworks[i].IpPoolName.IsNull() {
			data.VirtualNetworks[i].IpPoolName = types.StringValue(value.String())
		} else {
			data.VirtualNetworks[i].IpPoolName = types.StringNull()
		}
		if value := r.Get("ipv4SsmRanges"); value.Exists() && !data.VirtualNetworks[i].Ipv4SsmRanges.IsNull() {
			data.VirtualNetworks[i].Ipv4SsmRanges = helpers.GetStringSet(value.Array())
		} else {
			data.VirtualNetworks[i].Ipv4SsmRanges = types.SetNull(types.StringType)
		}
		for ci := range data.VirtualNetworks[i].MulticastRps {
			keys := [...]string{"rpDeviceLocation", "ipv4Address", "ipv6Address", "isDefaultV4RP", "isDefaultV6RP"}
			keyValues := [...]string{data.VirtualNetworks[i].MulticastRps[ci].RpDeviceLocation.ValueString(), data.VirtualNetworks[i].MulticastRps[ci].Ipv4Address.ValueString(), data.VirtualNetworks[i].MulticastRps[ci].Ipv6Address.ValueString(), strconv.FormatBool(data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV4Rp.ValueBool()), strconv.FormatBool(data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV6Rp.ValueBool())}

			var cr gjson.Result
			r.Get("multicastRPs").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("rpDeviceLocation"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].RpDeviceLocation.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].RpDeviceLocation = types.StringValue(value.String())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].RpDeviceLocation = types.StringNull()
			}
			if value := cr.Get("ipv4Address"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].Ipv4Address.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv4Address = types.StringValue(value.String())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv4Address = types.StringNull()
			}
			if value := cr.Get("ipv6Address"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].Ipv6Address.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv6Address = types.StringValue(value.String())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv6Address = types.StringNull()
			}
			if value := cr.Get("isDefaultV4RP"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV4Rp.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV4Rp = types.BoolValue(value.Bool())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV4Rp = types.BoolNull()
			}
			if value := cr.Get("isDefaultV6RP"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV6Rp.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV6Rp = types.BoolValue(value.Bool())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].IsDefaultV6Rp = types.BoolNull()
			}
			if value := cr.Get("networkDeviceIds"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].NetworkDeviceIds.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].NetworkDeviceIds = helpers.GetStringSet(value.Array())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].NetworkDeviceIds = types.SetNull(types.StringType)
			}
			if value := cr.Get("ipv4AsmRanges"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].Ipv4AsmRanges.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv4AsmRanges = helpers.GetStringSet(value.Array())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv4AsmRanges = types.SetNull(types.StringType)
			}
			if value := cr.Get("ipv6AsmRanges"); value.Exists() && !data.VirtualNetworks[i].MulticastRps[ci].Ipv6AsmRanges.IsNull() {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv6AsmRanges = helpers.GetStringSet(value.Array())
			} else {
				data.VirtualNetworks[i].MulticastRps[ci].Ipv6AsmRanges = types.SetNull(types.StringType)
			}
		}
		if data.VirtualNetworks[i].Id != types.StringNull() {
			final = append(final, data.VirtualNetworks[i])
		}
	}
	data.VirtualNetworks = final
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FabricMulticastVirtualNetworks) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.VirtualNetworks {
		keys := [...]string{"virtualNetworkName"}
		keyValues := [...]string{data.VirtualNetworks[i].VirtualNetworkName.ValueString()}

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
		if data.VirtualNetworks[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.VirtualNetworks[i].Id.IsNull() {
				data.VirtualNetworks[i].Id = types.StringValue(value.String())
			} else {
				data.VirtualNetworks[i].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricMulticastVirtualNetworks) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.VirtualNetworks) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
