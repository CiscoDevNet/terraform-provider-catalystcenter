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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Role struct {
	Id            types.String        `tfsdk:"id"`
	Name          types.String        `tfsdk:"name"`
	Description   types.String        `tfsdk:"description"`
	ResourceTypes []RoleResourceTypes `tfsdk:"resource_types"`
}

type RoleResourceTypes struct {
	Type       types.String `tfsdk:"type"`
	Operations types.Set    `tfsdk:"operations"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Role) getPath() string {
	return "/dna/system/api/v1/role"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Role) toBody(ctx context.Context, state Role) string {
	body := ""
	if state.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "roleId", state.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "role", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.ResourceTypes) > 0 {
		body, _ = sjson.Set(body, "resourceTypes", []interface{}{})
		for _, item := range data.ResourceTypes {
			itemBody := ""
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Operations.IsNull() {
				var values []string
				item.Operations.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "operations", values)
			}
			body, _ = sjson.SetRaw(body, "resourceTypes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Role) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("resourceTypes"); value.Exists() {
		data.ResourceTypes = make([]RoleResourceTypes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RoleResourceTypes{}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("operations"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Operations = helpers.GetStringSet(cValue.Array())
			} else {
				item.Operations = types.SetNull(types.StringType)
			}
			data.ResourceTypes = append(data.ResourceTypes, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Role) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.ResourceTypes {
		keys := [...]string{"type"}
		keyValues := [...]string{data.ResourceTypes[i].Type.ValueString()}

		var r gjson.Result
		res.Get("resourceTypes").ForEach(
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
		if value := r.Get("type"); value.Exists() && !data.ResourceTypes[i].Type.IsNull() {
			data.ResourceTypes[i].Type = types.StringValue(value.String())
		} else {
			data.ResourceTypes[i].Type = types.StringNull()
		}
		if value := r.Get("operations"); value.Exists() && !data.ResourceTypes[i].Operations.IsNull() {
			data.ResourceTypes[i].Operations = helpers.GetStringSet(value.Array())
		} else {
			data.ResourceTypes[i].Operations = types.SetNull(types.StringType)
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Role) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if len(data.ResourceTypes) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
