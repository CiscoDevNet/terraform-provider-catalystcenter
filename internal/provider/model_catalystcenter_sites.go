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
)

// End of section. //template:end imports

type Sites struct {
	Sites []SitesSites `tfsdk:"sites"`
	Type  types.String `tfsdk:"type"`
}

type SitesSites struct {
	Id             types.String  `tfsdk:"id"`
	Name           types.String  `tfsdk:"name"`
	NameHierarchy  types.String  `tfsdk:"name_hierarchy"`
	Type           types.String  `tfsdk:"type"`
	Latitude       types.String  `tfsdk:"latitude"`
	Longitude      types.String  `tfsdk:"longitude"`
	Country        types.String  `tfsdk:"country"`
	Address        types.String  `tfsdk:"address"`
	ParentId       types.String  `tfsdk:"parent_id"`
	FloorNumber    types.Int64   `tfsdk:"floor_number"`
	RfModel        types.String  `tfsdk:"rf_model"`
	Width          types.Float64 `tfsdk:"width"`
	Length         types.Float64 `tfsdk:"length"`
	Height         types.Float64 `tfsdk:"height"`
	UnitsOfMeasure types.String  `tfsdk:"units_of_measure"`
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Sites) getPath() string {
	return "/dna/intent/api/v1/sites"
}

// End of section. //template:end getPath

func (data *Sites) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Sites = make([]SitesSites, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SitesSites{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("nameHierarchy"); cValue.Exists() {
				item.NameHierarchy = types.StringValue(cValue.String())
			} else {
				item.NameHierarchy = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("parentId"); cValue.Exists() {
				item.ParentId = types.StringValue(cValue.String())
			} else {
				item.ParentId = types.StringNull()
			}
			if cValue := v.Get("latitude"); cValue.Exists() {
				item.Latitude = types.StringValue(cValue.String())
			} else {
				item.Latitude = types.StringNull()
			}
			if cValue := v.Get("longitude"); cValue.Exists() {
				item.Longitude = types.StringValue(cValue.String())
			} else {
				item.Longitude = types.StringNull()
			}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			} else {
				item.Address = types.StringNull()
			}
			if cValue := v.Get("country"); cValue.Exists() {
				item.Country = types.StringValue(cValue.String())
			} else {
				item.Country = types.StringNull()
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
			data.Sites = append(data.Sites, item)
			return true
		})
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Sites) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Type.IsNull() {
		return false
	}
	if len(data.Sites) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
