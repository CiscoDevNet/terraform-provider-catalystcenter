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

type Images struct {
	Images []ImagesImages `tfsdk:"images"`
}

type ImagesImages struct {
	Id              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Family          types.String `tfsdk:"family"`
	Vendor          types.String `tfsdk:"vendor"`
	Version         types.String `tfsdk:"version"`
	ApplicationType types.String `tfsdk:"application_type"`
	ImageType       types.String `tfsdk:"image_type"`
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Images) getPath() string {
	return "/dna/intent/api/v1/images"
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

func (data Images) toBody(ctx context.Context, state Images) string {
	body := ""
	if len(data.Images) > 0 {
		body, _ = sjson.Set(body, "response", []interface{}{})
		for _, item := range data.Images {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Family.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "family", item.Family.ValueString())
			}
			if !item.Vendor.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vendor", item.Vendor.ValueString())
			}
			if !item.Version.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "version", item.Version.ValueString())
			}
			if !item.ApplicationType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "applicationType", item.ApplicationType.ValueString())
			}
			if !item.ImageType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "imageType", item.ImageType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "response.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Images) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Images = make([]ImagesImages, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ImagesImages{}
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
			if cValue := v.Get("family"); cValue.Exists() {
				item.Family = types.StringValue(cValue.String())
			} else {
				item.Family = types.StringNull()
			}
			if cValue := v.Get("vendor"); cValue.Exists() {
				item.Vendor = types.StringValue(cValue.String())
			} else {
				item.Vendor = types.StringNull()
			}
			if cValue := v.Get("version"); cValue.Exists() {
				item.Version = types.StringValue(cValue.String())
			} else {
				item.Version = types.StringNull()
			}
			if cValue := v.Get("applicationType"); cValue.Exists() {
				item.ApplicationType = types.StringValue(cValue.String())
			} else {
				item.ApplicationType = types.StringNull()
			}
			if cValue := v.Get("imageType"); cValue.Exists() {
				item.ImageType = types.StringValue(cValue.String())
			} else {
				item.ImageType = types.StringNull()
			}
			data.Images = append(data.Images, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Images) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Images {
		keys := [...]string{"id", "name", "family", "vendor", "version", "applicationType", "imageType"}
		keyValues := [...]string{data.Images[i].Id.ValueString(), data.Images[i].Name.ValueString(), data.Images[i].Family.ValueString(), data.Images[i].Vendor.ValueString(), data.Images[i].Version.ValueString(), data.Images[i].ApplicationType.ValueString(), data.Images[i].ImageType.ValueString()}

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
		if value := r.Get("id"); value.Exists() && !data.Images[i].Id.IsNull() {
			data.Images[i].Id = types.StringValue(value.String())
		} else {
			data.Images[i].Id = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.Images[i].Name.IsNull() {
			data.Images[i].Name = types.StringValue(value.String())
		} else {
			data.Images[i].Name = types.StringNull()
		}
		if value := r.Get("family"); value.Exists() && !data.Images[i].Family.IsNull() {
			data.Images[i].Family = types.StringValue(value.String())
		} else {
			data.Images[i].Family = types.StringNull()
		}
		if value := r.Get("vendor"); value.Exists() && !data.Images[i].Vendor.IsNull() {
			data.Images[i].Vendor = types.StringValue(value.String())
		} else {
			data.Images[i].Vendor = types.StringNull()
		}
		if value := r.Get("version"); value.Exists() && !data.Images[i].Version.IsNull() {
			data.Images[i].Version = types.StringValue(value.String())
		} else {
			data.Images[i].Version = types.StringNull()
		}
		if value := r.Get("applicationType"); value.Exists() && !data.Images[i].ApplicationType.IsNull() {
			data.Images[i].ApplicationType = types.StringValue(value.String())
		} else {
			data.Images[i].ApplicationType = types.StringNull()
		}
		if value := r.Get("imageType"); value.Exists() && !data.Images[i].ImageType.IsNull() {
			data.Images[i].ImageType = types.StringValue(value.String())
		} else {
			data.Images[i].ImageType = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Images) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Images) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
