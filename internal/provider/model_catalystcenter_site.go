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
type Site struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	ParentId      types.String `tfsdk:"parent_id"`
	NameHierarchy types.String `tfsdk:"name_hierarchy"`
	Type          types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Site) getPath() string {
	return "/dna/intent/api/v1/sites"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Site) toBody(ctx context.Context, state Site) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.ParentId.IsNull() {
		body, _ = sjson.Set(body, "parentId", data.ParentId.ValueString())
	}
	if !data.NameHierarchy.IsNull() {
		body, _ = sjson.Set(body, "nameHierarchy", data.NameHierarchy.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Site) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.parentId"); value.Exists() {
		data.ParentId = types.StringValue(value.String())
	} else {
		data.ParentId = types.StringNull()
	}
	if value := res.Get("response.0.nameHierarchy"); value.Exists() {
		data.NameHierarchy = types.StringValue(value.String())
	} else {
		data.NameHierarchy = types.StringNull()
	}
	if value := res.Get("response.0.type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Site) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.parentId"); value.Exists() && !data.ParentId.IsNull() {
		data.ParentId = types.StringValue(value.String())
	} else {
		data.ParentId = types.StringNull()
	}
	if value := res.Get("response.0.nameHierarchy"); value.Exists() && !data.NameHierarchy.IsNull() {
		data.NameHierarchy = types.StringValue(value.String())
	} else {
		data.NameHierarchy = types.StringNull()
	}
	if value := res.Get("response.0.type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Site) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.ParentId.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
