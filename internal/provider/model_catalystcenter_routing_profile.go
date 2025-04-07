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

type RoutingProfile struct {
	Id                types.String                      `tfsdk:"id"`
	Name              types.String                      `tfsdk:"name"`
	ProfileAttributes []RoutingProfileProfileAttributes `tfsdk:"profile_attributes"`
}

type RoutingProfileProfileAttributes struct {
	Key        types.String                      `tfsdk:"key"`
	Value      types.String                      `tfsdk:"value"`
	Attributes []RoutingProfileProfileAttributes `tfsdk:"attributes"`
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data RoutingProfile) getPath() string {
	return "/api/v1/siteprofile"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

func attributesToJSON(attrs []RoutingProfileProfileAttributes) []interface{} {
	result := make([]interface{}, 0)
	for _, attr := range attrs {
		item := map[string]interface{}{
			"key":   attr.Key.ValueString(),
			"value": attr.Value.ValueString(),
		}
		if len(attr.Attributes) > 0 {
			item["attribs"] = attributesToJSON(attr.Attributes)
		}
		result = append(result, item)
	}
	return result
}

func attributesFromJSON(json gjson.Result) []RoutingProfileProfileAttributes {
	attrs := make([]RoutingProfileProfileAttributes, 0)
	for _, item := range json.Array() {
		attr := RoutingProfileProfileAttributes{
			Key:   types.StringValue(item.Get("key").String()),
			Value: types.StringValue(item.Get("value").String()),
		}
		if item.Get("attribs").Exists() {
			attr.Attributes = attributesFromJSON(item.Get("attribs"))
		}
		attrs = append(attrs, attr)
	}
	return attrs
}

func (data RoutingProfile) toBody(ctx context.Context, state RoutingProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "namespace", "routing")
	if len(data.ProfileAttributes) > 0 {
		body, _ = sjson.Set(body, "profileAttributes", attributesToJSON(data.ProfileAttributes))
	}
	return body
}

func (data *RoutingProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.profileAttributes"); value.Exists() && len(value.Array()) > 0 {
		if value := res.Get("response.profileAttributes"); value.Exists() {
			data.ProfileAttributes = attributesFromJSON(value)
		}
	}
}

func findMatching(result gjson.Result, keys []string, values []string) gjson.Result {
	var match gjson.Result
	result.ForEach(func(_, item gjson.Result) bool {
		matched := true
		for i := range keys {
			if item.Get(keys[i]).String() != values[i] {
				matched = false
				break
			}
		}
		if matched {
			match = item
			return false // stop iteration
		}
		return true
	})
	return match
}

func updateAttributeFromResult(attr *RoutingProfileProfileAttributes, result gjson.Result) {
	if val := result.Get("key"); val.Exists() && !attr.Key.IsNull() {
		attr.Key = types.StringValue(val.String())
	}
	if val := result.Get("value"); val.Exists() && !attr.Value.IsNull() {
		attr.Value = types.StringValue(val.String())
	}
}

func (data *RoutingProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}

	attrs := res.Get("response.profileAttributes")
	for i := range data.ProfileAttributes {
		// Match top-level attribute
		r := findMatching(attrs, []string{"key"}, []string{data.ProfileAttributes[i].Key.ValueString()})
		updateAttributeFromResult(&data.ProfileAttributes[i], r)

		subAttrs := r.Get("attribs")
		for j := range data.ProfileAttributes[i].Attributes {
			cr := findMatching(subAttrs, []string{"value"}, []string{data.ProfileAttributes[i].Attributes[j].Value.ValueString()})
			updateAttributeFromResult(&data.ProfileAttributes[i].Attributes[j], cr)

			subSubAttrs := cr.Get("attribs")
			for k := range data.ProfileAttributes[i].Attributes[j].Attributes {
				ccr := findMatching(subSubAttrs, []string{"value"}, []string{data.ProfileAttributes[i].Attributes[j].Attributes[k].Value.ValueString()})
				updateAttributeFromResult(&data.ProfileAttributes[i].Attributes[j].Attributes[k], ccr)
			}
		}
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *RoutingProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if len(data.ProfileAttributes) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
