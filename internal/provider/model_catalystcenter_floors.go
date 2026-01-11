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

type Floors struct {
	Id     types.String   `tfsdk:"id"`
	Floors []FloorsFloors `tfsdk:"floors"`
}

type FloorsFloors struct {
	Id                  types.String  `tfsdk:"id"`
	ParentId            types.String  `tfsdk:"parent_id"`
	ParentNameHierarchy types.String  `tfsdk:"parent_name_hierarchy"`
	Name                types.String  `tfsdk:"name"`
	FloorNumber         types.Int64   `tfsdk:"floor_number"`
	RfModel             types.String  `tfsdk:"rf_model"`
	Width               types.Float64 `tfsdk:"width"`
	Length              types.Float64 `tfsdk:"length"`
	Height              types.Float64 `tfsdk:"height"`
	UnitsOfMeasure      types.String  `tfsdk:"units_of_measure"`
}

func (data Floors) getPath() string {
	return "/dna/intent/api/v1/sites/bulk"
}

func (data Floors) toBody(ctx context.Context, state Floors) string {
	body := "[]"
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.Floors) > 0 {
		body, _ = sjson.Set(body, "", []interface{}{})
		for _, item := range data.Floors {
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
			itemBody, _ = sjson.Set(itemBody, "type", "floor")
			if !item.FloorNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "floorNumber", item.FloorNumber.ValueInt64())
			}
			if !item.RfModel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rfModel", item.RfModel.ValueString())
			}
			if !item.Width.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "width", item.Width.ValueFloat64())
			}
			if !item.Length.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "length", item.Length.ValueFloat64())
			}
			if !item.Height.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "height", item.Height.ValueFloat64())
			}
			if !item.UnitsOfMeasure.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "unitsOfMeasure", item.UnitsOfMeasure.ValueString())
			}
			body, _ = sjson.SetRaw(body, "-1", itemBody)
		}
	}
	return body
}

func (data *Floors) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Floors = make([]FloorsFloors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FloorsFloors{}
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
			if cValue := v.Get("floorNumber"); cValue.Exists() {
				item.FloorNumber = types.Int64Value(cValue.Int())
			} else {
				item.FloorNumber = types.Int64Null()
			}
			if cValue := v.Get("rfModel"); cValue.Exists() {
				item.RfModel = types.StringValue(cValue.String())
			} else {
				item.RfModel = types.StringNull()
			}
			if cValue := v.Get("width"); cValue.Exists() {
				item.Width = types.Float64Value(cValue.Float())
			} else {
				item.Width = types.Float64Null()
			}
			if cValue := v.Get("length"); cValue.Exists() {
				item.Length = types.Float64Value(cValue.Float())
			} else {
				item.Length = types.Float64Null()
			}
			if cValue := v.Get("height"); cValue.Exists() {
				item.Height = types.Float64Value(cValue.Float())
			} else {
				item.Height = types.Float64Null()
			}
			if cValue := v.Get("unitsOfMeasure"); cValue.Exists() {
				item.UnitsOfMeasure = types.StringValue(cValue.String())
			} else {
				item.UnitsOfMeasure = types.StringNull()
			}
			data.Floors = append(data.Floors, item)
			return true
		})
	}
}

func (data *Floors) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []FloorsFloors

	res = res.Get("response")

	// Build a lookup map
	responseMap := make(map[string]gjson.Result)
	res.ForEach(func(_, v gjson.Result) bool {
		if v.Get("type").String() == "floor" {
			nameHierarchy := v.Get("nameHierarchy").String()
			responseMap[nameHierarchy] = v
		}
		return true
	})

	for i := range data.Floors {
		// Construct full nameHierarchy as parentNameHierarchy + "/" + name
		fullHierarchy := data.Floors[i].ParentNameHierarchy.ValueString() + "/" + data.Floors[i].Name.ValueString()

		r, found := responseMap[fullHierarchy]

		if found {
			if value := r.Get("id"); value.Exists() {
				data.Floors[i].Id = types.StringValue(value.String())
			} else if data.Floors[i].Id.IsNull() {
				data.Floors[i].Id = types.StringNull()
			}
			// Read parentId from response
			if value := r.Get("parentId"); value.Exists() {
				data.Floors[i].ParentId = types.StringValue(value.String())
			} else {
				data.Floors[i].ParentId = types.StringNull()
			}
			// ParentNameHierarchy is not updated from response - it's user-specified and stays as-is
			if value := r.Get("name"); value.Exists() && !data.Floors[i].Name.IsNull() {
				data.Floors[i].Name = types.StringValue(value.String())
			} else {
				data.Floors[i].Name = types.StringNull()
			}
			if value := r.Get("floorNumber"); value.Exists() && !data.Floors[i].FloorNumber.IsNull() {
				data.Floors[i].FloorNumber = types.Int64Value(value.Int())
			} else {
				data.Floors[i].FloorNumber = types.Int64Null()
			}
			if value := r.Get("rfModel"); value.Exists() && !data.Floors[i].RfModel.IsNull() {
				data.Floors[i].RfModel = types.StringValue(value.String())
			} else {
				data.Floors[i].RfModel = types.StringNull()
			}
			if value := r.Get("width"); value.Exists() && !data.Floors[i].Width.IsNull() {
				data.Floors[i].Width = types.Float64Value(value.Float())
			} else {
				data.Floors[i].Width = types.Float64Null()
			}
			if value := r.Get("length"); value.Exists() && !data.Floors[i].Length.IsNull() {
				data.Floors[i].Length = types.Float64Value(value.Float())
			} else {
				data.Floors[i].Length = types.Float64Null()
			}
			if value := r.Get("height"); value.Exists() && !data.Floors[i].Height.IsNull() {
				data.Floors[i].Height = types.Float64Value(value.Float())
			} else {
				data.Floors[i].Height = types.Float64Null()
			}
			if value := r.Get("unitsOfMeasure"); value.Exists() && !data.Floors[i].UnitsOfMeasure.IsNull() {
				data.Floors[i].UnitsOfMeasure = types.StringValue(value.String())
			} else {
				data.Floors[i].UnitsOfMeasure = types.StringNull()
			}

			// Only add to final if found in API response and has valid ID to enable drift detection
			if data.Floors[i].Id != types.StringNull() {
				final = append(final, data.Floors[i])
			}
		}
		// If not found in API response, item is not added to final
	}
	data.Floors = final
}

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Floors) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")

	// Build a lookup map
	responseMap := make(map[string]gjson.Result)
	res.ForEach(func(_, v gjson.Result) bool {
		if v.Get("type").String() == "floor" {
			nameHierarchy := v.Get("nameHierarchy").String()
			responseMap[nameHierarchy] = v
		}
		return true
	})

	for i := range data.Floors {
		// Construct full nameHierarchy as parentNameHierarchy + "/" + name
		fullHierarchy := data.Floors[i].ParentNameHierarchy.ValueString() + "/" + data.Floors[i].Name.ValueString()

		r, found := responseMap[fullHierarchy]

		if found {
			if data.Floors[i].Id.IsUnknown() {
				if value := r.Get("id"); value.Exists() && !data.Floors[i].Id.IsNull() {
					data.Floors[i].Id = types.StringValue(value.String())
				} else {
					data.Floors[i].Id = types.StringNull()
				}
			}
			if data.Floors[i].ParentId.IsUnknown() {
				if value := r.Get("parentId"); value.Exists() {
					data.Floors[i].ParentId = types.StringValue(value.String())
				} else {
					data.Floors[i].ParentId = types.StringNull()
				}
			}
		}
	}
}

func (data *Floors) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Floors) > 0 {
		return false
	}
	return true
}
