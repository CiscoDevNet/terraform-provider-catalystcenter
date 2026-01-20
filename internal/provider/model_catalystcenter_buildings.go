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

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Buildings struct {
	Id        types.String                  `tfsdk:"id"`
	Scope     types.String                  `tfsdk:"scope"`
	Buildings map[string]BuildingsBuildings `tfsdk:"buildings"`
}

type BuildingsBuildings struct {
	Id                  types.String  `tfsdk:"id"`
	ParentId            types.String  `tfsdk:"parent_id"`
	ParentNameHierarchy types.String  `tfsdk:"parent_name_hierarchy"`
	Name                types.String  `tfsdk:"name"`
	Country             types.String  `tfsdk:"country"`
	Address             types.String  `tfsdk:"address"`
	Latitude            types.Float64 `tfsdk:"latitude"`
	Longitude           types.Float64 `tfsdk:"longitude"`
}

func (data Buildings) getPath() string {
	return "/dna/intent/api/v1/sites/bulk"
}

func (data Buildings) toBody(ctx context.Context, state Buildings) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.Buildings) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.Buildings {
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
			itemBody, _ = sjson.Set(itemBody, "type", "building")
			if !item.Country.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "country", item.Country.ValueString())
			}
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Latitude.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "latitude", item.Latitude.ValueFloat64())
			}
			if !item.Longitude.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "longitude", item.Longitude.ValueFloat64())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

func (data *Buildings) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Buildings = make(map[string]BuildingsBuildings)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BuildingsBuildings{}
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
			if cValue := v.Get("country"); cValue.Exists() {
				item.Country = types.StringValue(cValue.String())
			} else {
				item.Country = types.StringNull()
			}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			} else {
				item.Address = types.StringNull()
			}
			if cValue := v.Get("latitude"); cValue.Exists() {
				item.Latitude = types.Float64Value(cValue.Float())
			} else {
				item.Latitude = types.Float64Null()
			}
			if cValue := v.Get("longitude"); cValue.Exists() {
				item.Longitude = types.Float64Value(cValue.Float())
			} else {
				item.Longitude = types.Float64Null()
			}
			// Use parentNameHierarchy/name as map key
			if !item.ParentNameHierarchy.IsNull() && !item.Name.IsNull() {
				key := item.ParentNameHierarchy.ValueString() + "/" + item.Name.ValueString()

				// Apply scope filtering if scope is specified
				if !data.Scope.IsNull() && data.Scope.ValueString() != "" {
					scope := data.Scope.ValueString()
					// Only include items within the scope (exact match or children)
					if !isWithinScope(key, scope) {
						return true // Skip this item
					}
				}

				data.Buildings[key] = item
			}
			return true
		})
	}
}

func (data *Buildings) updateFromBody(ctx context.Context, res gjson.Result) {
	res = res.Get("response")

	// Build a lookup map with scope filtering
	responseMap := make(map[string]gjson.Result)
	res.ForEach(func(_, v gjson.Result) bool {
		if v.Get("type").String() == "building" {
			nameHierarchy := v.Get("nameHierarchy").String()

			// Apply scope filtering if scope is specified
			if !data.Scope.IsNull() && data.Scope.ValueString() != "" {
				scope := data.Scope.ValueString()
				// Only include items within the scope (exact match or children)
				if !isWithinScope(nameHierarchy, scope) {
					return true // Skip this item
				}
			}

			responseMap[nameHierarchy] = v
		}
		return true
	})

	for key, building := range data.Buildings {
		// Construct full nameHierarchy as parentNameHierarchy + "/" + name
		fullHierarchy := building.ParentNameHierarchy.ValueString() + "/" + building.Name.ValueString()

		r, found := responseMap[fullHierarchy]

		if found {
			if value := r.Get("id"); value.Exists() {
				building.Id = types.StringValue(value.String())
			} else if building.Id.IsNull() {
				building.Id = types.StringNull()
			}
			// Read parentId from response
			if value := r.Get("parentId"); value.Exists() {
				building.ParentId = types.StringValue(value.String())
			} else {
				building.ParentId = types.StringNull()
			}
			// ParentNameHierarchy is not updated from response - it's user-specified and stays as-is
			if value := r.Get("name"); value.Exists() && !building.Name.IsNull() {
				building.Name = types.StringValue(value.String())
			} else {
				building.Name = types.StringNull()
			}
			if value := r.Get("country"); value.Exists() && !building.Country.IsNull() {
				building.Country = types.StringValue(value.String())
			} else {
				building.Country = types.StringNull()
			}
			if value := r.Get("address"); value.Exists() && !building.Address.IsNull() {
				building.Address = types.StringValue(value.String())
			} else {
				building.Address = types.StringNull()
			}
			if value := r.Get("latitude"); value.Exists() && !building.Latitude.IsNull() {
				building.Latitude = types.Float64Value(value.Float())
			} else {
				building.Latitude = types.Float64Null()
			}
			if value := r.Get("longitude"); value.Exists() && !building.Longitude.IsNull() {
				building.Longitude = types.Float64Value(value.Float())
			} else {
				building.Longitude = types.Float64Null()
			}

			// Update map with modified building if it has a valid ID
			if building.Id != types.StringNull() {
				data.Buildings[key] = building
			} else {
				// Remove from map if no valid ID (drift detection)
				delete(data.Buildings, key)
			}
		} else {
			// If not found in API response, remove from map (drift detection)
			delete(data.Buildings, key)
		}
	}
}

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Buildings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")

	// Build a lookup map
	responseMap := make(map[string]gjson.Result)
	res.ForEach(func(_, v gjson.Result) bool {
		if v.Get("type").String() == "building" {
			nameHierarchy := v.Get("nameHierarchy").String()
			responseMap[nameHierarchy] = v
		}
		return true
	})

	for key, building := range data.Buildings {
		// Construct full nameHierarchy as parentNameHierarchy + "/" + name
		fullHierarchy := building.ParentNameHierarchy.ValueString() + "/" + building.Name.ValueString()

		r, found := responseMap[fullHierarchy]

		if found {
			if building.Id.IsUnknown() {
				if value := r.Get("id"); value.Exists() && !building.Id.IsNull() {
					building.Id = types.StringValue(value.String())
				} else {
					building.Id = types.StringNull()
				}
			}
			if building.ParentId.IsUnknown() {
				if value := r.Get("parentId"); value.Exists() {
					building.ParentId = types.StringValue(value.String())
				} else {
					building.ParentId = types.StringNull()
				}
			}
			// Update map with modified building
			data.Buildings[key] = building
		}
	}
}

func (data *Buildings) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Buildings) > 0 {
		return false
	}
	return true
}
