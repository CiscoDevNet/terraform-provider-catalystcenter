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
type Buildings struct {
	Id        types.String         `tfsdk:"id"`
	Buildings []BuildingsBuildings `tfsdk:"buildings"`
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

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Buildings) getPath() string {
	return "/dna/intent/api/v1/sites/bulk"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
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

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Buildings) fromBody(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	if value := res; value.Exists() && len(value.Array()) > 0 {
		data.Buildings = make([]BuildingsBuildings, 0)
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
			data.Buildings = append(data.Buildings, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Buildings) updateFromBody(ctx context.Context, res gjson.Result) {
	var final []BuildingsBuildings

	res = res.Get("response")
	for i := range data.Buildings {
		keys := [...]string{"nameHierarchy", "type"}
		// Construct full nameHierarchy as parentNameHierarchy + "/" + name
		fullHierarchy := data.Buildings[i].ParentNameHierarchy.ValueString() + "/" + data.Buildings[i].Name.ValueString()
		keyValues := [...]string{fullHierarchy, "building"}

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
			data.Buildings[i].Id = types.StringValue(value.String())
		} else if data.Buildings[i].Id.IsNull() {
			data.Buildings[i].Id = types.StringNull()
		}
		// Read parentId from response
		if value := r.Get("parentId"); value.Exists() {
			data.Buildings[i].ParentId = types.StringValue(value.String())
		} else {
			data.Buildings[i].ParentId = types.StringNull()
		}
		// ParentNameHierarchy is not updated from response - it's user-specified and stays as-is
		if value := r.Get("name"); value.Exists() && !data.Buildings[i].Name.IsNull() {
			data.Buildings[i].Name = types.StringValue(value.String())
		} else {
			data.Buildings[i].Name = types.StringNull()
		}
		if value := r.Get("country"); value.Exists() && !data.Buildings[i].Country.IsNull() {
			data.Buildings[i].Country = types.StringValue(value.String())
		} else {
			data.Buildings[i].Country = types.StringNull()
		}
		if value := r.Get("address"); value.Exists() && !data.Buildings[i].Address.IsNull() {
			data.Buildings[i].Address = types.StringValue(value.String())
		} else {
			data.Buildings[i].Address = types.StringNull()
		}
		if value := r.Get("latitude"); value.Exists() && !data.Buildings[i].Latitude.IsNull() {
			data.Buildings[i].Latitude = types.Float64Value(value.Float())
		} else {
			data.Buildings[i].Latitude = types.Float64Null()
		}
		if value := r.Get("longitude"); value.Exists() && !data.Buildings[i].Longitude.IsNull() {
			data.Buildings[i].Longitude = types.Float64Value(value.Float())
		} else {
			data.Buildings[i].Longitude = types.Float64Null()
		}
		if data.Buildings[i].Id != types.StringNull() {
			final = append(final, data.Buildings[i])
		}
	}
	data.Buildings = final
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Buildings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {

	res = res.Get("response")
	for i := range data.Buildings {
		var r gjson.Result
		res.ForEach(
			func(_, v gjson.Result) bool {
				// Match by name and type
				if v.Get("name").String() != data.Buildings[i].Name.ValueString() {
					return true
				}
				if v.Get("type").String() != "building" {
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
				if parentHierarchy != data.Buildings[i].ParentNameHierarchy.ValueString() {
					return true
				}
				// Found match
				r = v
				return false
			},
		)
		if data.Buildings[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.Buildings[i].Id.IsNull() {
				data.Buildings[i].Id = types.StringValue(value.String())
			} else {
				data.Buildings[i].Id = types.StringNull()
			}
		}
		if data.Buildings[i].ParentId.IsUnknown() {
			if value := r.Get("parentId"); value.Exists() {
				data.Buildings[i].ParentId = types.StringValue(value.String())
			} else {
				data.Buildings[i].ParentId = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Buildings) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Buildings) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
