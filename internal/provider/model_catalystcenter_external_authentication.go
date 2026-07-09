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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ExternalAuthentication struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ExternalAuthentication) getPath() string {
	return "/dna/system/api/v1/users/external-authentication"
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ExternalAuthentication) toBody(ctx context.Context, state ExternalAuthentication) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enable", data.Enabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Custom (frozen) fromBody: the data source has no "name" input (data_source_no_id),
// so default the synthetic identifier to the fixed value when it is not set from config.
// Kept outside the generator markers so `go generate` preserves it.
func (data *ExternalAuthentication) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if !data.Name.IsNull() {
		data.Id = types.StringValue(fmt.Sprint(data.Name.ValueString()))
	} else {
		data.Id = types.StringValue("external_authentication")
		data.Name = types.StringValue("external_authentication")
	}
	if value := res.Get("response.external-authentication-flag.0.enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
}

// End of custom (frozen) section.

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ExternalAuthentication) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.external-authentication-flag.0.enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ExternalAuthentication) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Enabled.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
