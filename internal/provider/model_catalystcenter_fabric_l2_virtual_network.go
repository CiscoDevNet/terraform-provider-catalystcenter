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
type FabricL2VirtualNetwork struct {
	Id                             types.String `tfsdk:"id"`
	FabricId                       types.String `tfsdk:"fabric_id"`
	VlanName                       types.String `tfsdk:"vlan_name"`
	VlanId                         types.Int64  `tfsdk:"vlan_id"`
	TrafficType                    types.String `tfsdk:"traffic_type"`
	FabricEnabledWireless          types.Bool   `tfsdk:"fabric_enabled_wireless"`
	AssociatedL3VirtualNetworkName types.String `tfsdk:"associated_l3_virtual_network_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricL2VirtualNetwork) getPath() string {
	return "/dna/intent/api/v1/sda/layer2VirtualNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricL2VirtualNetwork) toBody(ctx context.Context, state FabricL2VirtualNetwork) string {
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
	if !data.VlanName.IsNull() {
		body, _ = sjson.Set(body, "0.vlanName", data.VlanName.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "0.vlanId", data.VlanId.ValueInt64())
	}
	if !data.TrafficType.IsNull() {
		body, _ = sjson.Set(body, "0.trafficType", data.TrafficType.ValueString())
	}
	if !data.FabricEnabledWireless.IsNull() {
		body, _ = sjson.Set(body, "0.isFabricEnabledWireless", data.FabricEnabledWireless.ValueBool())
	}
	if !data.AssociatedL3VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.associatedLayer3VirtualNetworkName", data.AssociatedL3VirtualNetworkName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricL2VirtualNetwork) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("response.0.vlanName"); value.Exists() {
		data.VlanName = types.StringValue(value.String())
	} else {
		data.VlanName = types.StringNull()
	}
	if value := res.Get("response.0.vlanId"); value.Exists() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("response.0.trafficType"); value.Exists() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("response.0.isFabricEnabledWireless"); value.Exists() {
		data.FabricEnabledWireless = types.BoolValue(value.Bool())
	} else {
		data.FabricEnabledWireless = types.BoolNull()
	}
	if value := res.Get("response.0.associatedLayer3VirtualNetworkName"); value.Exists() {
		data.AssociatedL3VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.AssociatedL3VirtualNetworkName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricL2VirtualNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.vlanName"); value.Exists() && !data.VlanName.IsNull() {
		data.VlanName = types.StringValue(value.String())
	} else {
		data.VlanName = types.StringNull()
	}
	if value := res.Get("response.0.vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("response.0.trafficType"); value.Exists() && !data.TrafficType.IsNull() {
		data.TrafficType = types.StringValue(value.String())
	} else {
		data.TrafficType = types.StringNull()
	}
	if value := res.Get("response.0.isFabricEnabledWireless"); value.Exists() && !data.FabricEnabledWireless.IsNull() {
		data.FabricEnabledWireless = types.BoolValue(value.Bool())
	} else {
		data.FabricEnabledWireless = types.BoolNull()
	}
	if value := res.Get("response.0.associatedLayer3VirtualNetworkName"); value.Exists() && !data.AssociatedL3VirtualNetworkName.IsNull() {
		data.AssociatedL3VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.AssociatedL3VirtualNetworkName = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricL2VirtualNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.TrafficType.IsNull() {
		return false
	}
	if !data.FabricEnabledWireless.IsNull() {
		return false
	}
	if !data.AssociatedL3VirtualNetworkName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
