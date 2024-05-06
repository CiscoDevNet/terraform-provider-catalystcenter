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
type TemplateVersion struct {
	Id         types.String `tfsdk:"id"`
	TemplateId types.String `tfsdk:"template_id"`
	Comments   types.String `tfsdk:"comments"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TemplateVersion) getPath() string {
	return "/dna/intent/api/v1/template-programmer/template/version"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TemplateVersion) toBody(ctx context.Context, state TemplateVersion) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.TemplateId.IsNull() {
		body, _ = sjson.Set(body, "templateId", data.TemplateId.ValueString())
	}
	if !data.Comments.IsNull() {
		body, _ = sjson.Set(body, "comments", data.Comments.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TemplateVersion) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.templateId"); value.Exists() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TemplateVersion) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.templateId"); value.Exists() && !data.TemplateId.IsNull() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TemplateVersion) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Comments.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
