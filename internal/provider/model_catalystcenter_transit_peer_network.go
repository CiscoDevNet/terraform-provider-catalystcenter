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
type TransitPeerNetwork struct {
	Id                          types.String                                    `tfsdk:"id"`
	TransitPeerNetworkName      types.String                                    `tfsdk:"transit_peer_network_name"`
	TransitPeerNetworkType      types.String                                    `tfsdk:"transit_peer_network_type"`
	RoutingProtocolName         types.String                                    `tfsdk:"routing_protocol_name"`
	AutonomousSystemNumber      types.String                                    `tfsdk:"autonomous_system_number"`
	TransitControlPlaneSettings []TransitPeerNetworkTransitControlPlaneSettings `tfsdk:"transit_control_plane_settings"`
}

type TransitPeerNetworkTransitControlPlaneSettings struct {
	SiteNameHierarchy         types.String `tfsdk:"site_name_hierarchy"`
	DeviceManagementIpAddress types.String `tfsdk:"device_management_ip_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransitPeerNetwork) getPath() string {
	return "/dna/intent/api/v1/business/sda/transit-peer-network"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransitPeerNetwork) toBody(ctx context.Context, state TransitPeerNetwork) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.TransitPeerNetworkName.IsNull() {
		body, _ = sjson.Set(body, "transitPeerNetworkName", data.TransitPeerNetworkName.ValueString())
	}
	if !data.TransitPeerNetworkType.IsNull() {
		body, _ = sjson.Set(body, "transitPeerNetworkType", data.TransitPeerNetworkType.ValueString())
	}
	if !data.RoutingProtocolName.IsNull() {
		body, _ = sjson.Set(body, "ipTransitSettings.routingProtocolName", data.RoutingProtocolName.ValueString())
	}
	if !data.AutonomousSystemNumber.IsNull() {
		body, _ = sjson.Set(body, "ipTransitSettings.autonomousSystemNumber", data.AutonomousSystemNumber.ValueString())
	}
	if len(data.TransitControlPlaneSettings) > 0 {
		body, _ = sjson.Set(body, "sdaTransitSettings.transitControlPlaneSettings", []interface{}{})
		for _, item := range data.TransitControlPlaneSettings {
			itemBody := ""
			if !item.SiteNameHierarchy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "siteNameHierarchy", item.SiteNameHierarchy.ValueString())
			}
			if !item.DeviceManagementIpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceManagementIpAddress", item.DeviceManagementIpAddress.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sdaTransitSettings.transitControlPlaneSettings.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransitPeerNetwork) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("transitPeerNetworkId"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("transitPeerNetworkName"); value.Exists() {
		data.TransitPeerNetworkName = types.StringValue(value.String())
	} else {
		data.TransitPeerNetworkName = types.StringNull()
	}
	if value := res.Get("transitPeerNetworkType"); value.Exists() {
		data.TransitPeerNetworkType = types.StringValue(value.String())
	} else {
		data.TransitPeerNetworkType = types.StringNull()
	}
	if value := res.Get("ipTransitSettings.routingProtocolName"); value.Exists() {
		data.RoutingProtocolName = types.StringValue(value.String())
	} else {
		data.RoutingProtocolName = types.StringNull()
	}
	if value := res.Get("ipTransitSettings.autonomousSystemNumber"); value.Exists() {
		data.AutonomousSystemNumber = types.StringValue(value.String())
	} else {
		data.AutonomousSystemNumber = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransitPeerNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("transitPeerNetworkName"); value.Exists() && !data.TransitPeerNetworkName.IsNull() {
		data.TransitPeerNetworkName = types.StringValue(value.String())
	} else {
		data.TransitPeerNetworkName = types.StringNull()
	}
	if value := res.Get("transitPeerNetworkType"); value.Exists() && !data.TransitPeerNetworkType.IsNull() {
		data.TransitPeerNetworkType = types.StringValue(value.String())
	} else {
		data.TransitPeerNetworkType = types.StringNull()
	}
	if value := res.Get("ipTransitSettings.routingProtocolName"); value.Exists() && !data.RoutingProtocolName.IsNull() {
		data.RoutingProtocolName = types.StringValue(value.String())
	} else {
		data.RoutingProtocolName = types.StringNull()
	}
	if value := res.Get("ipTransitSettings.autonomousSystemNumber"); value.Exists() && !data.AutonomousSystemNumber.IsNull() {
		data.AutonomousSystemNumber = types.StringValue(value.String())
	} else {
		data.AutonomousSystemNumber = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransitPeerNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.TransitPeerNetworkType.IsNull() {
		return false
	}
	if !data.RoutingProtocolName.IsNull() {
		return false
	}
	if !data.AutonomousSystemNumber.IsNull() {
		return false
	}
	if len(data.TransitControlPlaneSettings) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
