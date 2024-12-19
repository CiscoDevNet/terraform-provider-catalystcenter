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
type FabricL3VirtualNetwork struct {
	Id                 types.String `tfsdk:"id"`
	VirtualNetworkName types.String `tfsdk:"virtual_network_name"`
	FabricIds          types.Set    `tfsdk:"fabric_ids"`
	AnchoredSiteId     types.String `tfsdk:"anchored_site_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricL3VirtualNetwork) getPath() string {
	return "/dna/intent/api/v1/sda/layer3VirtualNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricL3VirtualNetwork) toBody(ctx context.Context, state FabricL3VirtualNetwork) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.FabricIds.IsNull() {
		var values []string
		data.FabricIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.fabricIds", values)
	}
	if !data.AnchoredSiteId.IsNull() {
		body, _ = sjson.Set(body, "0.anchoredSiteId", data.AnchoredSiteId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricL3VirtualNetwork) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.fabricIds"); value.Exists() && len(value.Array()) > 0 {
		data.FabricIds = helpers.GetStringSet(value.Array())
	} else {
		data.FabricIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.anchoredSiteId"); value.Exists() {
		data.AnchoredSiteId = types.StringValue(value.String())
	} else {
		data.AnchoredSiteId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricL3VirtualNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.fabricIds"); value.Exists() && !data.FabricIds.IsNull() {
		data.FabricIds = helpers.GetStringSet(value.Array())
	} else {
		data.FabricIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.anchoredSiteId"); value.Exists() && !data.AnchoredSiteId.IsNull() {
		data.AnchoredSiteId = types.StringValue(value.String())
	} else {
		data.AnchoredSiteId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricL3VirtualNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	return true
}

// End of section. //template:end isNull
