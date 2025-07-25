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
type VirtualNetworkToFabricSite struct {
	Id                 types.String `tfsdk:"id"`
	VirtualNetworkName types.String `tfsdk:"virtual_network_name"`
	VirtualNetworkId   types.String `tfsdk:"virtual_network_id"`
	FabricSiteId       types.String `tfsdk:"fabric_site_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data VirtualNetworkToFabricSite) getPath() string {
	return "/dna/intent/api/v1/sda/layer3VirtualNetworks"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

func (data VirtualNetworkToFabricSite) toBody(ctx context.Context, state VirtualNetworkToFabricSite, fabricIds []string) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.VirtualNetworkId.IsNull() {
		body, _ = sjson.Set(body, "0.id", data.VirtualNetworkId.ValueString())
	}
	if len(fabricIds) > 0 {
		body, _ = sjson.Set(body, "0.fabricIds", fabricIds)
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *VirtualNetworkToFabricSite) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.id"); value.Exists() {
		data.VirtualNetworkId = types.StringValue(value.String())
	} else {
		data.VirtualNetworkId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *VirtualNetworkToFabricSite) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.id"); value.Exists() && !data.VirtualNetworkId.IsNull() {
		data.VirtualNetworkId = types.StringValue(value.String())
	} else {
		data.VirtualNetworkId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *VirtualNetworkToFabricSite) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.VirtualNetworkId.IsNull() {
		return false
	}
	if !data.FabricSiteId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
