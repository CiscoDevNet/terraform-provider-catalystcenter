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
type TransitNetwork struct {
	Id                            types.String `tfsdk:"id"`
	Name                          types.String `tfsdk:"name"`
	Type                          types.String `tfsdk:"type"`
	RoutingProtocolName           types.String `tfsdk:"routing_protocol_name"`
	AutonomousSystemNumber        types.String `tfsdk:"autonomous_system_number"`
	ControlPlaneNetworkDeviceIds  types.Set    `tfsdk:"control_plane_network_device_ids"`
	IsMulticastOverTransitEnabled types.Bool   `tfsdk:"is_multicast_over_transit_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransitNetwork) getPath() string {
	return "/dna/intent/api/v1/sda/transitNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransitNetwork) toBody(ctx context.Context, state TransitNetwork) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "0.name", data.Name.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "0.type", data.Type.ValueString())
	}
	if !data.RoutingProtocolName.IsNull() {
		body, _ = sjson.Set(body, "0.ipTransitSettings.routingProtocolName", data.RoutingProtocolName.ValueString())
	}
	if !data.AutonomousSystemNumber.IsNull() {
		body, _ = sjson.Set(body, "0.ipTransitSettings.autonomousSystemNumber", data.AutonomousSystemNumber.ValueString())
	}
	if !data.ControlPlaneNetworkDeviceIds.IsNull() {
		var values []string
		data.ControlPlaneNetworkDeviceIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.sdaTransitSettings.controlPlaneNetworkDeviceIds", values)
	}
	if !data.IsMulticastOverTransitEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.sdaTransitSettings.isMulticastOverTransitEnabled", data.IsMulticastOverTransitEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransitNetwork) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("response.0.ipTransitSettings.routingProtocolName"); value.Exists() {
		data.RoutingProtocolName = types.StringValue(value.String())
	} else {
		data.RoutingProtocolName = types.StringNull()
	}
	if value := res.Get("response.0.ipTransitSettings.autonomousSystemNumber"); value.Exists() {
		data.AutonomousSystemNumber = types.StringValue(value.String())
	} else {
		data.AutonomousSystemNumber = types.StringNull()
	}
	if value := res.Get("response.0.sdaTransitSettings.controlPlaneNetworkDeviceIds"); value.Exists() && len(value.Array()) > 0 {
		data.ControlPlaneNetworkDeviceIds = helpers.GetStringSet(value.Array())
	} else {
		data.ControlPlaneNetworkDeviceIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.sdaTransitSettings.isMulticastOverTransitEnabled"); value.Exists() {
		data.IsMulticastOverTransitEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsMulticastOverTransitEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransitNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("response.0.ipTransitSettings.routingProtocolName"); value.Exists() && !data.RoutingProtocolName.IsNull() {
		data.RoutingProtocolName = types.StringValue(value.String())
	} else {
		data.RoutingProtocolName = types.StringNull()
	}
	if value := res.Get("response.0.ipTransitSettings.autonomousSystemNumber"); value.Exists() && !data.AutonomousSystemNumber.IsNull() {
		data.AutonomousSystemNumber = types.StringValue(value.String())
	} else {
		data.AutonomousSystemNumber = types.StringNull()
	}
	if value := res.Get("response.0.sdaTransitSettings.controlPlaneNetworkDeviceIds"); value.Exists() && !data.ControlPlaneNetworkDeviceIds.IsNull() {
		data.ControlPlaneNetworkDeviceIds = helpers.GetStringSet(value.Array())
	} else {
		data.ControlPlaneNetworkDeviceIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.sdaTransitSettings.isMulticastOverTransitEnabled"); value.Exists() && !data.IsMulticastOverTransitEnabled.IsNull() {
		data.IsMulticastOverTransitEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsMulticastOverTransitEnabled = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransitNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Type.IsNull() {
		return false
	}
	if !data.RoutingProtocolName.IsNull() {
		return false
	}
	if !data.AutonomousSystemNumber.IsNull() {
		return false
	}
	if !data.ControlPlaneNetworkDeviceIds.IsNull() {
		return false
	}
	if !data.IsMulticastOverTransitEnabled.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
