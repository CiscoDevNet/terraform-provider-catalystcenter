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
type FabricL2Handoff struct {
	Id              types.String `tfsdk:"id"`
	NetworkDeviceId types.String `tfsdk:"network_device_id"`
	FabricId        types.String `tfsdk:"fabric_id"`
	InterfaceName   types.String `tfsdk:"interface_name"`
	InternalVlanId  types.Int64  `tfsdk:"internal_vlan_id"`
	ExternalVlanId  types.Int64  `tfsdk:"external_vlan_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricL2Handoff) getPath() string {
	return "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricL2Handoff) toBody(ctx context.Context, state FabricL2Handoff) string {
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
	if !data.InterfaceName.IsNull() {
		body, _ = sjson.Set(body, "0.interfaceName", data.InterfaceName.ValueString())
	}
	if !data.InternalVlanId.IsNull() {
		body, _ = sjson.Set(body, "0.internalVlanId", data.InternalVlanId.ValueInt64())
	}
	if !data.ExternalVlanId.IsNull() {
		body, _ = sjson.Set(body, "0.externalVlanId", data.ExternalVlanId.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricL2Handoff) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("interfaceName"); value.Exists() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("internalVlanId"); value.Exists() {
		data.InternalVlanId = types.Int64Value(value.Int())
	} else {
		data.InternalVlanId = types.Int64Null()
	}
	if value := res.Get("externalVlanId"); value.Exists() {
		data.ExternalVlanId = types.Int64Value(value.Int())
	} else {
		data.ExternalVlanId = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricL2Handoff) updateFromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("interfaceName"); value.Exists() && !data.InterfaceName.IsNull() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("internalVlanId"); value.Exists() && !data.InternalVlanId.IsNull() {
		data.InternalVlanId = types.Int64Value(value.Int())
	} else {
		data.InternalVlanId = types.Int64Null()
	}
	if value := res.Get("externalVlanId"); value.Exists() && !data.ExternalVlanId.IsNull() {
		data.ExternalVlanId = types.Int64Value(value.Int())
	} else {
		data.ExternalVlanId = types.Int64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricL2Handoff) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.InterfaceName.IsNull() {
		return false
	}
	if !data.InternalVlanId.IsNull() {
		return false
	}
	if !data.ExternalVlanId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
