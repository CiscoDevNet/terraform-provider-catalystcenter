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
type GlobalCredentialNETCONF struct {
	Id             types.String `tfsdk:"id"`
	NetconfPort    types.String `tfsdk:"netconf_port"`
	Comments       types.String `tfsdk:"comments"`
	CredentialType types.String `tfsdk:"credential_type"`
	Description    types.String `tfsdk:"description"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data GlobalCredentialNETCONF) getPath() string {
	return "/dna/intent/api/v1/global-credential/netconf"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data GlobalCredentialNETCONF) getPathDelete() string {
	return "/dna/intent/api/v1/global-credential"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data GlobalCredentialNETCONF) toBody(ctx context.Context, state GlobalCredentialNETCONF) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "id", state.Id.ValueString())
	}
	_ = put
	if !data.NetconfPort.IsNull() {
		if put {
			body, _ = sjson.Set(body, "netconfPort", data.NetconfPort.ValueString())
		} else {
			body, _ = sjson.Set(body, "0.netconfPort", data.NetconfPort.ValueString())
		}
	}
	if !data.Comments.IsNull() {
		if put {
			body, _ = sjson.Set(body, "comments", data.Comments.ValueString())
		} else {
			body, _ = sjson.Set(body, "0.comments", data.Comments.ValueString())
		}
	}
	if !data.CredentialType.IsNull() {
		if put {
			body, _ = sjson.Set(body, "credentialType", data.CredentialType.ValueString())
		} else {
			body, _ = sjson.Set(body, "0.credentialType", data.CredentialType.ValueString())
		}
	}
	if !data.Description.IsNull() {
		if put {
			body, _ = sjson.Set(body, "description", data.Description.ValueString())
		} else {
			body, _ = sjson.Set(body, "0.description", data.Description.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *GlobalCredentialNETCONF) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("netconfPort"); value.Exists() {
		data.NetconfPort = types.StringValue(value.String())
	} else {
		data.NetconfPort = types.StringNull()
	}
	if value := res.Get("comments"); value.Exists() {
		data.Comments = types.StringValue(value.String())
	} else {
		data.Comments = types.StringNull()
	}
	if value := res.Get("credentialType"); value.Exists() {
		data.CredentialType = types.StringValue(value.String())
	} else {
		data.CredentialType = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *GlobalCredentialNETCONF) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("netconfPort"); value.Exists() && !data.NetconfPort.IsNull() {
		data.NetconfPort = types.StringValue(value.String())
	} else {
		data.NetconfPort = types.StringNull()
	}
	if value := res.Get("comments"); value.Exists() && !data.Comments.IsNull() {
		data.Comments = types.StringValue(value.String())
	} else {
		data.Comments = types.StringNull()
	}
	if value := res.Get("credentialType"); value.Exists() && !data.CredentialType.IsNull() {
		data.CredentialType = types.StringValue(value.String())
	} else {
		data.CredentialType = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *GlobalCredentialNETCONF) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.NetconfPort.IsNull() {
		return false
	}
	if !data.Comments.IsNull() {
		return false
	}
	if !data.CredentialType.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
