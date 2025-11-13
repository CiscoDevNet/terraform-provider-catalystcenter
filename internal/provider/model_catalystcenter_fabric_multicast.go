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
type FabricMulticast struct {
	Id                 types.String                  `tfsdk:"id"`
	FabricId           types.String                  `tfsdk:"fabric_id"`
	VirtualNetworkName types.String                  `tfsdk:"virtual_network_name"`
	IpPoolName         types.String                  `tfsdk:"ip_pool_name"`
	Ipv4SsmRanges      types.List                    `tfsdk:"ipv4_ssm_ranges"`
	MulticastRPs       []FabricMulticastMulticastRPs `tfsdk:"multicast_r_ps"`
}

type FabricMulticastMulticastRPs struct {
	RpDeviceLocation types.String `tfsdk:"rp_device_location"`
	NetworkDeviceIds types.List   `tfsdk:"network_device_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricMulticast) getPath() string {
	return "/dna/intent/api/v1/sda/multicast/virtualNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath
func (data FabricMulticast) getFallbackPath() string {
	return ""
}

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricMulticast) toBody(ctx context.Context, state FabricMulticast) string {
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
	if len(data.MulticastRPs) > 0 {
		body, _ = sjson.Set(body, "multicastRPs", []interface{}{})
		for _, item := range data.MulticastRPs {
			itemBody := ""
			if !item.RpDeviceLocation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rpDeviceLocation", item.RpDeviceLocation.ValueString())
			}
			if !item.NetworkDeviceIds.IsNull() {
				var values []string
				item.NetworkDeviceIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "networkDeviceIds", values)
			}
			body, _ = sjson.SetRaw(body, "multicastRPs.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricMulticast) fromBody(ctx context.Context, res gjson.Result) {
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
		data.Ipv4SsmRanges = helpers.GetStringList(value.Array())
	} else {
		data.Ipv4SsmRanges = types.ListNull(types.StringType)
	}
	if value := res.Get("multicastRPs"); value.Exists() && len(value.Array()) > 0 {
		data.MulticastRPs = make([]FabricMulticastMulticastRPs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricMulticastMulticastRPs{}
			if cValue := v.Get("rpDeviceLocation"); cValue.Exists() {
				item.RpDeviceLocation = types.StringValue(cValue.String())
			} else {
				item.RpDeviceLocation = types.StringNull()
			}
			if cValue := v.Get("networkDeviceIds"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.NetworkDeviceIds = helpers.GetStringList(cValue.Array())
			} else {
				item.NetworkDeviceIds = types.ListNull(types.StringType)
			}
			data.MulticastRPs = append(data.MulticastRPs, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricMulticast) updateFromBody(ctx context.Context, res gjson.Result) {
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
		data.Ipv4SsmRanges = helpers.GetStringList(value.Array())
	} else {
		data.Ipv4SsmRanges = types.ListNull(types.StringType)
	}
	for i := range data.MulticastRPs {
		keys := [...]string{"rpDeviceLocation"}
		keyValues := [...]string{data.MulticastRPs[i].RpDeviceLocation.ValueString()}

		var r gjson.Result
		res.Get("multicastRPs").ForEach(
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
		if value := r.Get("rpDeviceLocation"); value.Exists() && !data.MulticastRPs[i].RpDeviceLocation.IsNull() {
			data.MulticastRPs[i].RpDeviceLocation = types.StringValue(value.String())
		} else {
			data.MulticastRPs[i].RpDeviceLocation = types.StringNull()
		}
		if value := r.Get("networkDeviceIds"); value.Exists() && !data.MulticastRPs[i].NetworkDeviceIds.IsNull() {
			data.MulticastRPs[i].NetworkDeviceIds = helpers.GetStringList(value.Array())
		} else {
			data.MulticastRPs[i].NetworkDeviceIds = types.ListNull(types.StringType)
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricMulticast) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.IpPoolName.IsNull() {
		return false
	}
	if !data.Ipv4SsmRanges.IsNull() {
		return false
	}
	if len(data.MulticastRPs) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
