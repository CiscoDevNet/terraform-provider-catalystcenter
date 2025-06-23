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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type FabricL3HandoffSDATransit struct {
	Id                            types.String `tfsdk:"id"`
	NetworkDeviceId               types.String `tfsdk:"network_device_id"`
	FabricId                      types.String `tfsdk:"fabric_id"`
	TransitNetworkId              types.String `tfsdk:"transit_network_id"`
	AffinityIdPrime               types.Int64  `tfsdk:"affinity_id_prime"`
	AffinityIdDecider             types.Int64  `tfsdk:"affinity_id_decider"`
	ConnectedToInternet           types.Bool   `tfsdk:"connected_to_internet"`
	IsMulticastOverTransitEnabled types.Bool   `tfsdk:"is_multicast_over_transit_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricL3HandoffSDATransit) getPath() string {
	return "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricL3HandoffSDATransit) toBody(ctx context.Context, state FabricL3HandoffSDATransit) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "0.networkDeviceId", data.NetworkDeviceId.ValueString())
	}
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.TransitNetworkId.IsNull() {
		body, _ = sjson.Set(body, "0.transitNetworkId", data.TransitNetworkId.ValueString())
	}
	if !data.AffinityIdPrime.IsNull() {
		body, _ = sjson.Set(body, "0.affinityIdPrime", data.AffinityIdPrime.ValueInt64())
	}
	if !data.AffinityIdDecider.IsNull() {
		body, _ = sjson.Set(body, "0.affinityIdDecider", data.AffinityIdDecider.ValueInt64())
	}
	if !data.ConnectedToInternet.IsNull() {
		body, _ = sjson.Set(body, "0.connectedToInternet", data.ConnectedToInternet.ValueBool())
	}
	if !data.IsMulticastOverTransitEnabled.IsNull() {
		body, _ = sjson.Set(body, "0.isMulticastOverTransitEnabled", data.IsMulticastOverTransitEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricL3HandoffSDATransit) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	data.Id = types.StringValue(fmt.Sprint(data.NetworkDeviceId.ValueString()))
	if value := res.Get("response.0.networkDeviceId"); value.Exists() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.transitNetworkId"); value.Exists() {
		data.TransitNetworkId = types.StringValue(value.String())
	} else {
		data.TransitNetworkId = types.StringNull()
	}
	if value := res.Get("response.0.affinityIdPrime"); value.Exists() {
		data.AffinityIdPrime = types.Int64Value(value.Int())
	} else {
		data.AffinityIdPrime = types.Int64Null()
	}
	if value := res.Get("response.0.affinityIdDecider"); value.Exists() {
		data.AffinityIdDecider = types.Int64Value(value.Int())
	} else {
		data.AffinityIdDecider = types.Int64Null()
	}
	if value := res.Get("response.0.connectedToInternet"); value.Exists() {
		data.ConnectedToInternet = types.BoolValue(value.Bool())
	} else {
		data.ConnectedToInternet = types.BoolNull()
	}
	if value := res.Get("response.0.isMulticastOverTransitEnabled"); value.Exists() {
		data.IsMulticastOverTransitEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsMulticastOverTransitEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricL3HandoffSDATransit) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.networkDeviceId"); value.Exists() && !data.NetworkDeviceId.IsNull() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.transitNetworkId"); value.Exists() && !data.TransitNetworkId.IsNull() {
		data.TransitNetworkId = types.StringValue(value.String())
	} else {
		data.TransitNetworkId = types.StringNull()
	}
	if value := res.Get("response.0.affinityIdPrime"); value.Exists() && !data.AffinityIdPrime.IsNull() {
		data.AffinityIdPrime = types.Int64Value(value.Int())
	} else {
		data.AffinityIdPrime = types.Int64Null()
	}
	if value := res.Get("response.0.affinityIdDecider"); value.Exists() && !data.AffinityIdDecider.IsNull() {
		data.AffinityIdDecider = types.Int64Value(value.Int())
	} else {
		data.AffinityIdDecider = types.Int64Null()
	}
	if value := res.Get("response.0.connectedToInternet"); value.Exists() && !data.ConnectedToInternet.IsNull() {
		data.ConnectedToInternet = types.BoolValue(value.Bool())
	} else {
		data.ConnectedToInternet = types.BoolNull()
	}
	if value := res.Get("response.0.isMulticastOverTransitEnabled"); value.Exists() && !data.IsMulticastOverTransitEnabled.IsNull() {
		data.IsMulticastOverTransitEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsMulticastOverTransitEnabled = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricL3HandoffSDATransit) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.TransitNetworkId.IsNull() {
		return false
	}
	if !data.AffinityIdPrime.IsNull() {
		return false
	}
	if !data.AffinityIdDecider.IsNull() {
		return false
	}
	if !data.ConnectedToInternet.IsNull() {
		return false
	}
	if !data.IsMulticastOverTransitEnabled.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
