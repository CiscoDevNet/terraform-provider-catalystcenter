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
type FabricPortChannel struct {
	Id                  types.String `tfsdk:"id"`
	NetworkDeviceId     types.String `tfsdk:"network_device_id"`
	FabricId            types.String `tfsdk:"fabric_id"`
	PortChannelName     types.String `tfsdk:"port_channel_name"`
	InterfaceNames      types.Set    `tfsdk:"interface_names"`
	ConnectedDeviceType types.String `tfsdk:"connected_device_type"`
	Protocol            types.String `tfsdk:"protocol"`
	Description         types.String `tfsdk:"description"`
	NativeVlanId        types.Int64  `tfsdk:"native_vlan_id"`
	AllowedVlanRanges   types.String `tfsdk:"allowed_vlan_ranges"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricPortChannel) getPath() string {
	return "/dna/intent/api/v1/sda/portChannels"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricPortChannel) toBody(ctx context.Context, state FabricPortChannel) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "0.networkDeviceId", data.NetworkDeviceId.ValueString())
	}
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if data.PortChannelName.ValueString() != "" && !data.PortChannelName.IsNull() {
		body, _ = sjson.Set(body, "0.portChannelName", data.PortChannelName.ValueString())
	}
	if !data.InterfaceNames.IsNull() {
		var values []string
		data.InterfaceNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.interfaceNames", values)
	}
	if !data.ConnectedDeviceType.IsNull() {
		body, _ = sjson.Set(body, "0.connectedDeviceType", data.ConnectedDeviceType.ValueString())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "0.protocol", data.Protocol.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "0.description", data.Description.ValueString())
	}
	if !data.NativeVlanId.IsNull() {
		body, _ = sjson.Set(body, "0.nativeVlanId", data.NativeVlanId.ValueInt64())
	}
	if !data.AllowedVlanRanges.IsNull() {
		body, _ = sjson.Set(body, "0.allowedVlanRanges", data.AllowedVlanRanges.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricPortChannel) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("networkDeviceId"); value.Exists() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("portChannelName"); value.Exists() {
		data.PortChannelName = types.StringValue(value.String())
	} else {
		data.PortChannelName = types.StringNull()
	}
	if value := res.Get("interfaceNames"); value.Exists() && len(value.Array()) > 0 {
		data.InterfaceNames = helpers.GetStringSet(value.Array())
	} else {
		data.InterfaceNames = types.SetNull(types.StringType)
	}
	if value := res.Get("connectedDeviceType"); value.Exists() {
		data.ConnectedDeviceType = types.StringValue(value.String())
	} else {
		data.ConnectedDeviceType = types.StringNull()
	}
	if value := res.Get("protocol"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("nativeVlanId"); value.Exists() {
		data.NativeVlanId = types.Int64Value(value.Int())
	} else {
		data.NativeVlanId = types.Int64Null()
	}
	if value := res.Get("allowedVlanRanges"); value.Exists() {
		data.AllowedVlanRanges = types.StringValue(value.String())
	} else {
		data.AllowedVlanRanges = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricPortChannel) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("networkDeviceId"); value.Exists() && !data.NetworkDeviceId.IsNull() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("portChannelName"); value.Exists() && !data.PortChannelName.IsNull() {
		data.PortChannelName = types.StringValue(value.String())
	} else {
		data.PortChannelName = types.StringNull()
	}
	if value := res.Get("interfaceNames"); value.Exists() && !data.InterfaceNames.IsNull() {
		data.InterfaceNames = helpers.GetStringSet(value.Array())
	} else {
		data.InterfaceNames = types.SetNull(types.StringType)
	}
	if value := res.Get("connectedDeviceType"); value.Exists() && !data.ConnectedDeviceType.IsNull() {
		data.ConnectedDeviceType = types.StringValue(value.String())
	} else {
		data.ConnectedDeviceType = types.StringNull()
	}
	if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && value.String() != "" && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("nativeVlanId"); value.Exists() && !data.NativeVlanId.IsNull() {
		data.NativeVlanId = types.Int64Value(value.Int())
	} else {
		data.NativeVlanId = types.Int64Null()
	}
	if value := res.Get("allowedVlanRanges"); value.Exists() && !data.AllowedVlanRanges.IsNull() {
		data.AllowedVlanRanges = types.StringValue(value.String())
	} else {
		data.AllowedVlanRanges = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricPortChannel) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.PortChannelName.IsNull() {
		return false
	}
	if !data.InterfaceNames.IsNull() {
		return false
	}
	if !data.ConnectedDeviceType.IsNull() {
		return false
	}
	if !data.Protocol.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.NativeVlanId.IsNull() {
		return false
	}
	if !data.AllowedVlanRanges.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
