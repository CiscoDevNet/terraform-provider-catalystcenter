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
type NetworkProfile struct {
	Id        types.String              `tfsdk:"id"`
	Name      types.String              `tfsdk:"name"`
	Type      types.String              `tfsdk:"type"`
	Templates []NetworkProfileTemplates `tfsdk:"templates"`
}

type NetworkProfileTemplates struct {
	Type       types.String                        `tfsdk:"type"`
	Attributes []NetworkProfileTemplatesAttributes `tfsdk:"attributes"`
}

type NetworkProfileTemplatesAttributes struct {
	TemplateId types.String `tfsdk:"template_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NetworkProfile) getPath() string {
	return "/api/v1/siteprofile"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data NetworkProfile) toBody(ctx context.Context, state NetworkProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "namespace", data.Type.ValueString())
	}
	if len(data.Templates) > 0 {
		body, _ = sjson.Set(body, "profileAttributes", []interface{}{})
		for _, item := range data.Templates {
			itemBody := ""
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "key", item.Type.ValueString())
			}
			if len(item.Attributes) > 0 {
				itemBody, _ = sjson.Set(itemBody, "attribs", []interface{}{})
				for _, childItem := range item.Attributes {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "key", "template.id")
					if !childItem.TemplateId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.TemplateId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "attribs.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "profileAttributes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *NetworkProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.namespace"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("response.profileAttributes"); value.Exists() {
		data.Templates = make([]NetworkProfileTemplates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkProfileTemplates{}
			if cValue := v.Get("key"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("attribs"); cValue.Exists() {
				item.Attributes = make([]NetworkProfileTemplatesAttributes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := NetworkProfileTemplatesAttributes{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.TemplateId = types.StringValue(ccValue.String())
					} else {
						cItem.TemplateId = types.StringNull()
					}
					item.Attributes = append(item.Attributes, cItem)
					return true
				})
			}
			data.Templates = append(data.Templates, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *NetworkProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.namespace"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	for i := range data.Templates {
		keys := [...]string{"key"}
		keyValues := [...]string{data.Templates[i].Type.ValueString()}

		var r gjson.Result
		res.Get("response.profileAttributes").ForEach(
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
		if value := r.Get("key"); value.Exists() && !data.Templates[i].Type.IsNull() {
			data.Templates[i].Type = types.StringValue(value.String())
		} else {
			data.Templates[i].Type = types.StringNull()
		}
		for ci := range data.Templates[i].Attributes {
			keys := [...]string{"value"}
			keyValues := [...]string{data.Templates[i].Attributes[ci].TemplateId.ValueString()}

			var cr gjson.Result
			r.Get("attribs").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("value"); value.Exists() && !data.Templates[i].Attributes[ci].TemplateId.IsNull() {
				data.Templates[i].Attributes[ci].TemplateId = types.StringValue(value.String())
			} else {
				data.Templates[i].Attributes[ci].TemplateId = types.StringNull()
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *NetworkProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Type.IsNull() {
		return false
	}
	if len(data.Templates) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
