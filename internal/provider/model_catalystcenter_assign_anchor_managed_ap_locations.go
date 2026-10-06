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
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AssignAnchorManagedAPLocations struct {
	Id                              types.String `tfsdk:"id"`
	AnchorManagedApLocationsSiteIds types.Set    `tfsdk:"anchor_managed_ap_locations_site_ids"`
	DeviceId                        types.String `tfsdk:"device_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssignAnchorManagedAPLocations) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessSettings/%v/assignAnchorManagedApLocations", url.QueryEscape(data.DeviceId.ValueString()))
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
func (data AssignAnchorManagedAPLocations) toBody(ctx context.Context, state AssignAnchorManagedAPLocations) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.AnchorManagedApLocationsSiteIds.IsNull() {
		var values []string
		data.AnchorManagedApLocationsSiteIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "anchorManagedAPLocationsSiteIds", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AssignAnchorManagedAPLocations) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("anchorManagedAPLocationsSiteIds"); value.Exists() && len(value.Array()) > 0 {
		data.AnchorManagedApLocationsSiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.AnchorManagedApLocationsSiteIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AssignAnchorManagedAPLocations) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("anchorManagedAPLocationsSiteIds"); value.Exists() && !data.AnchorManagedApLocationsSiteIds.IsNull() {
		data.AnchorManagedApLocationsSiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.AnchorManagedApLocationsSiteIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssignAnchorManagedAPLocations) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.AnchorManagedApLocationsSiteIds.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
