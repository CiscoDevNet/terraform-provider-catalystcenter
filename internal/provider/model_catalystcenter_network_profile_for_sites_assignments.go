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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type NetworkProfileForSitesAssignments struct {
	Id               types.String                             `tfsdk:"id"`
	NetworkProfileId types.String                             `tfsdk:"network_profile_id"`
	Items            []NetworkProfileForSitesAssignmentsItems `tfsdk:"items"`
}

type NetworkProfileForSitesAssignmentsItems struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NetworkProfileForSitesAssignments) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/networkProfilesForSites/%v/siteAssignments", url.QueryEscape(data.NetworkProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data NetworkProfileForSitesAssignments) toBody(ctx context.Context, state NetworkProfileForSitesAssignments) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *NetworkProfileForSitesAssignments) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	data.Id = types.StringValue(fmt.Sprint(data.NetworkProfileId.ValueString()))
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Items = make([]NetworkProfileForSitesAssignmentsItems, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkProfileForSitesAssignmentsItems{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			data.Items = append(data.Items, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *NetworkProfileForSitesAssignments) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Items[i].Id.ValueString()}

		var r gjson.Result
		res.Get("response").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("id"); value.Exists() && !data.Items[i].Id.IsNull() {
			data.Items[i].Id = types.StringValue(value.String())
		} else {
			data.Items[i].Id = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *NetworkProfileForSitesAssignments) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Items) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
