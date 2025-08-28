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

type TemplateVersions struct {
	TemplateId       types.String                       `tfsdk:"template_id"`
	TemplateVersions []TemplateVersionsTemplateVersions `tfsdk:"template_versions"`
}

type TemplateVersionsTemplateVersions struct {
	Id             types.String `tfsdk:"id"`
	Version        types.String `tfsdk:"version"`
	VersionComment types.String `tfsdk:"version_comment"`
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TemplateVersions) getPath() string {
	return "/dna/intent/api/v1/template-programmer/template/version"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

func (data *TemplateVersions) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.templateId"); value.Exists() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
	if value := res.Get("0.versionsInfo"); value.Exists() && len(value.Array()) > 0 {
		data.TemplateVersions = make([]TemplateVersionsTemplateVersions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TemplateVersionsTemplateVersions{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("version"); cValue.Exists() {
				item.Version = types.StringValue(cValue.String())
			} else {
				item.Version = types.StringNull()
			}
			if cValue := v.Get("versionComment"); cValue.Exists() {
				item.VersionComment = types.StringValue(cValue.String())
			} else {
				item.VersionComment = types.StringNull()
			}
			data.TemplateVersions = append(data.TemplateVersions, item)
			return true
		})
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TemplateVersions) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.templateId"); value.Exists() && !data.TemplateId.IsNull() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
	for i := range data.TemplateVersions {
		keys := [...]string{"id", "version", "versionComment"}
		keyValues := [...]string{data.TemplateVersions[i].Id.ValueString(), data.TemplateVersions[i].Version.ValueString(), data.TemplateVersions[i].VersionComment.ValueString()}

		var r gjson.Result
		res.Get("0.versionsInfo").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.TemplateVersions[i].Id.IsNull() {
			data.TemplateVersions[i].Id = types.StringValue(value.String())
		} else {
			data.TemplateVersions[i].Id = types.StringNull()
		}
		if value := r.Get("version"); value.Exists() && !data.TemplateVersions[i].Version.IsNull() {
			data.TemplateVersions[i].Version = types.StringValue(value.String())
		} else {
			data.TemplateVersions[i].Version = types.StringNull()
		}
		if value := r.Get("versionComment"); value.Exists() && !data.TemplateVersions[i].VersionComment.IsNull() {
			data.TemplateVersions[i].VersionComment = types.StringValue(value.String())
		} else {
			data.TemplateVersions[i].VersionComment = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TemplateVersions) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.TemplateVersions) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
