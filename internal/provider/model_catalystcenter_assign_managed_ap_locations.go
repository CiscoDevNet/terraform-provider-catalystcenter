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
type AssignManagedAPLocations struct {
	Id                                 types.String `tfsdk:"id"`
	PrimaryManagedApLocationsSiteIds   types.Set    `tfsdk:"primary_managed_ap_locations_site_ids"`
	SecondaryManagedApLocationsSiteIds types.Set    `tfsdk:"secondary_managed_ap_locations_site_ids"`
	DeviceId                           types.String `tfsdk:"device_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssignManagedAPLocations) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessControllers/%v/assignManagedApLocations", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AssignManagedAPLocations) toBody(ctx context.Context, state AssignManagedAPLocations) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.PrimaryManagedApLocationsSiteIds.IsNull() {
		var values []string
		data.PrimaryManagedApLocationsSiteIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "primaryManagedAPLocationsSiteIds", values)
	}
	if !data.SecondaryManagedApLocationsSiteIds.IsNull() {
		var values []string
		data.SecondaryManagedApLocationsSiteIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "secondaryManagedAPLocationsSiteIds", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AssignManagedAPLocations) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("primaryManagedAPLocationsSiteIds"); value.Exists() && len(value.Array()) > 0 {
		data.PrimaryManagedApLocationsSiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.PrimaryManagedApLocationsSiteIds = types.SetNull(types.StringType)
	}
	if value := res.Get("secondaryManagedAPLocationsSiteIds"); value.Exists() && len(value.Array()) > 0 {
		data.SecondaryManagedApLocationsSiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.SecondaryManagedApLocationsSiteIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AssignManagedAPLocations) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("primaryManagedAPLocationsSiteIds"); value.Exists() && !data.PrimaryManagedApLocationsSiteIds.IsNull() {
		data.PrimaryManagedApLocationsSiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.PrimaryManagedApLocationsSiteIds = types.SetNull(types.StringType)
	}
	if value := res.Get("secondaryManagedAPLocationsSiteIds"); value.Exists() && !data.SecondaryManagedApLocationsSiteIds.IsNull() {
		data.SecondaryManagedApLocationsSiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.SecondaryManagedApLocationsSiteIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssignManagedAPLocations) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.PrimaryManagedApLocationsSiteIds.IsNull() {
		return false
	}
	if !data.SecondaryManagedApLocationsSiteIds.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
