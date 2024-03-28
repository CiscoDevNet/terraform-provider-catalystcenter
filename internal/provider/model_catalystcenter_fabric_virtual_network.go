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
type FabricVirtualNetwork struct {
	Id                 types.String `tfsdk:"id"`
	VirtualNetworkName types.String `tfsdk:"virtual_network_name"`
	IsGuest            types.Bool   `tfsdk:"is_guest"`
	SgNames            types.Set    `tfsdk:"sg_names"`
	VmanageVpnId       types.String `tfsdk:"vmanage_vpn_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricVirtualNetwork) getPath() string {
	return "/dna/intent/api/v1/virtual-network"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricVirtualNetwork) toBody(ctx context.Context, state FabricVirtualNetwork) string {
	body := ""
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.IsGuest.IsNull() {
		body, _ = sjson.Set(body, "isGuestVirtualNetwork", data.IsGuest.ValueBool())
	}
	if !data.SgNames.IsNull() {
		var values []string
		data.SgNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "scalableGroupNames", values)
	}
	if !data.VmanageVpnId.IsNull() {
		body, _ = sjson.Set(body, "vManageVpnId", data.VmanageVpnId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricVirtualNetwork) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("isGuestVirtualNetwork"); value.Exists() {
		data.IsGuest = types.BoolValue(value.Bool())
	} else {
		data.IsGuest = types.BoolValue(false)
	}
	if value := res.Get("scalableGroupNames"); value.Exists() {
		data.SgNames = helpers.GetStringSet(value.Array())
	} else {
		data.SgNames = types.SetNull(types.StringType)
	}
	if value := res.Get("vManageVpnId"); value.Exists() {
		data.VmanageVpnId = types.StringValue(value.String())
	} else {
		data.VmanageVpnId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricVirtualNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("isGuestVirtualNetwork"); value.Exists() && !data.IsGuest.IsNull() {
		data.IsGuest = types.BoolValue(value.Bool())
	} else if data.IsGuest.ValueBool() != false {
		data.IsGuest = types.BoolNull()
	}
	if value := res.Get("scalableGroupNames"); value.Exists() && !data.SgNames.IsNull() {
		data.SgNames = helpers.GetStringSet(value.Array())
	} else {
		data.SgNames = types.SetNull(types.StringType)
	}
	if value := res.Get("vManageVpnId"); value.Exists() && !data.VmanageVpnId.IsNull() {
		data.VmanageVpnId = types.StringValue(value.String())
	} else {
		data.VmanageVpnId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricVirtualNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.VirtualNetworkName.IsNull() {
		return false
	}
	if !data.IsGuest.IsNull() {
		return false
	}
	if !data.SgNames.IsNull() {
		return false
	}
	if !data.VmanageVpnId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
