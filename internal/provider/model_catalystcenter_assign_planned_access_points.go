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
type AssignPlannedAccessPoints struct {
	Id          types.String                           `tfsdk:"id"`
	FloorId     types.String                           `tfsdk:"floor_id"`
	Assignments []AssignPlannedAccessPointsAssignments `tfsdk:"assignments"`
}

type AssignPlannedAccessPointsAssignments struct {
	PlannedAccessPointName types.String `tfsdk:"planned_access_point_name"`
	AccessPointId          types.String `tfsdk:"access_point_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssignPlannedAccessPoints) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v2/floors/%v/plannedAccessPointPositions/assignAccessPointPositions", url.QueryEscape(data.FloorId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

func (data AssignPlannedAccessPoints) getPathGet() string {
	return fmt.Sprintf("/dna/intent/api/v2/floors/%v/accessPointPositions", url.QueryEscape(data.FloorId.ValueString()))
}

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Custom toBody: the API takes planned Access Point IDs, which are resolved
// from planned_access_point_name at apply time.
func assignPlannedAccessPointsBody(items []AssignPlannedAccessPointsAssignments, plannedIds map[string]string) string {
	body := "[]"
	for _, item := range items {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "plannedAccessPointId", plannedIds[item.PlannedAccessPointName.ValueString()])
		itemBody, _ = sjson.Set(itemBody, "accessPointId", item.AccessPointId.ValueString())
		body, _ = sjson.SetRaw(body, "-1", itemBody)
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AssignPlannedAccessPoints) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Assignments = make([]AssignPlannedAccessPointsAssignments, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AssignPlannedAccessPointsAssignments{}
			item.PlannedAccessPointName = types.StringValue("")
			if cValue := v.Get("id"); cValue.Exists() {
				item.AccessPointId = types.StringValue(cValue.String())
			} else {
				item.AccessPointId = types.StringNull()
			}
			data.Assignments = append(data.Assignments, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AssignPlannedAccessPoints) updateFromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.Assignments {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Assignments[i].AccessPointId.ValueString()}

		var r gjson.Result
		res.ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.Assignments[i].AccessPointId.IsNull() {
			data.Assignments[i].AccessPointId = types.StringValue(value.String())
		} else {
			data.Assignments[i].AccessPointId = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssignPlannedAccessPoints) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Assignments) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
