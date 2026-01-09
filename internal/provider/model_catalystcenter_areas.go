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
type Areas struct {
	Id    types.String `tfsdk:"id"`
	Areas []AreasAreas `tfsdk:"areas"`
}

type AreasAreas struct {
	Id                  types.String `tfsdk:"id"`
	ParentId            types.String `tfsdk:"parent_id"`
	ParentNameHierarchy types.String `tfsdk:"parent_name_hierarchy"`
	Name                types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Areas) getPath() string {
	return "/dna/intent/api/v1/sites/bulk"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Areas) toBody(ctx context.Context, state Areas) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.Areas) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.Areas {
			itemBody := ""
			// Only include id for PUT operations
			if put && item.Id.ValueString() != "" && !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			// Always include parentId
			if !item.ParentId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parentId", item.ParentId.ValueString())
			}
			if !item.ParentNameHierarchy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parentNameHierarchy", item.ParentNameHierarchy.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "area")
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Areas) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Areas = make([]AreasAreas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AreasAreas{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("parentId"); cValue.Exists() {
				item.ParentId = types.StringValue(cValue.String())
			} else {
				item.ParentId = types.StringNull()
			}
			// Extract parent_name_hierarchy from nameHierarchy (strip last part)
			if cValue := v.Get("nameHierarchy"); cValue.Exists() {
				nameHierarchy := cValue.String()
				// Find last slash and take everything before it
				lastSlash := -1
				for idx := len(nameHierarchy) - 1; idx >= 0; idx-- {
					if nameHierarchy[idx] == '/' {
						lastSlash = idx
						break
					}
				}
				if lastSlash > 0 {
					item.ParentNameHierarchy = types.StringValue(nameHierarchy[:lastSlash])
				} else {
					item.ParentNameHierarchy = types.StringNull()
				}
			} else {
				item.ParentNameHierarchy = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			data.Areas = append(data.Areas, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Areas) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []AreasAreas

	res = res.Get("response")
	for i := range data.Areas {
		keys := [...]string{"nameHierarchy", "type"}
		// Construct full nameHierarchy as parentNameHierarchy + "/" + name
		fullHierarchy := data.Areas[i].ParentNameHierarchy.ValueString() + "/" + data.Areas[i].Name.ValueString()
		keyValues := [...]string{fullHierarchy, "area"}

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
		if value := r.Get("id"); value.Exists() {
			data.Areas[i].Id = types.StringValue(value.String())
		} else if data.Areas[i].Id.IsNull() {
			data.Areas[i].Id = types.StringNull()
		}
		// Read parentId from response
		if value := r.Get("parentId"); value.Exists() {
			data.Areas[i].ParentId = types.StringValue(value.String())
		} else {
			data.Areas[i].ParentId = types.StringNull()
		}
		// ParentNameHierarchy is not updated from response - it's user-specified and stays as-is
		if value := r.Get("name"); value.Exists() && !data.Areas[i].Name.IsNull() {
			data.Areas[i].Name = types.StringValue(value.String())
		} else {
			data.Areas[i].Name = types.StringNull()
		}
		if data.Areas[i].Id != types.StringNull() {
			final = append(final, data.Areas[i])
		}
	}
	data.Areas = final
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Areas) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.Areas {
		var r gjson.Result
		res.ForEach(
			func(_, v gjson.Result) bool {
				// Match by name and type
				if v.Get("name").String() != data.Areas[i].Name.ValueString() {
					return true
				}
				if v.Get("type").String() != "area" {
					return true
				}
				// Match by parent hierarchy - strip last part from nameHierarchy
				nameHierarchy := v.Get("nameHierarchy").String()
				// Find last slash and take everything before it
				lastSlash := -1
				for idx := len(nameHierarchy) - 1; idx >= 0; idx-- {
					if nameHierarchy[idx] == '/' {
						lastSlash = idx
						break
					}
				}
				parentHierarchy := ""
				if lastSlash > 0 {
					parentHierarchy = nameHierarchy[:lastSlash]
				}
				if parentHierarchy != data.Areas[i].ParentNameHierarchy.ValueString() {
					return true
				}
				// Found match
				r = v
				return false
			},
		)
		if data.Areas[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.Areas[i].Id.IsNull() {
				data.Areas[i].Id = types.StringValue(value.String())
			} else {
				data.Areas[i].Id = types.StringNull()
			}
		}
		if data.Areas[i].ParentId.IsUnknown() {
			if value := r.Get("parentId"); value.Exists() {
				data.Areas[i].ParentId = types.StringValue(value.String())
			} else {
				data.Areas[i].ParentId = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Areas) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Areas) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
