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
type CredentialsCLI struct {
	Id             types.String `tfsdk:"id"`
	Description    types.String `tfsdk:"description"`
	Username       types.String `tfsdk:"username"`
	Password       types.String `tfsdk:"password"`
	EnablePassword types.String `tfsdk:"enable_password"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CredentialsCLI) getPath() string {
	return "/dna/intent/api/v2/global-credential"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CredentialsCLI) toBody(ctx context.Context, state CredentialsCLI) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "cliCredential.0.id", state.Id.ValueString())
	}
	_ = put
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "cliCredential.0.description", data.Description.ValueString())
	}
	if !data.Username.IsNull() {
		body, _ = sjson.Set(body, "cliCredential.0.username", data.Username.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, "cliCredential.0.password", data.Password.ValueString())
	}
	if !data.EnablePassword.IsNull() {
		body, _ = sjson.Set(body, "cliCredential.0.enablePassword", data.EnablePassword.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CredentialsCLI) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("username"); value.Exists() {
		data.Username = types.StringValue(value.String())
	} else {
		data.Username = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *CredentialsCLI) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("username"); value.Exists() && !data.Username.IsNull() {
		data.Username = types.StringValue(value.String())
	} else {
		data.Username = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *CredentialsCLI) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Description.IsNull() {
		return false
	}
	if !data.Username.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.EnablePassword.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
