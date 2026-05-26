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
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AssignCredentials struct {
	Id            types.String `tfsdk:"id"`
	SiteId        types.String `tfsdk:"site_id"`
	CliId         types.String `tfsdk:"cli_id"`
	SnmpV2ReadId  types.String `tfsdk:"snmp_v2_read_id"`
	SnmpV2WriteId types.String `tfsdk:"snmp_v2_write_id"`
	SnmpV3Id      types.String `tfsdk:"snmp_v3_id"`
	HttpsReadId   types.String `tfsdk:"https_read_id"`
	HttpsWriteId  types.String `tfsdk:"https_write_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssignCredentials) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/deviceCredentials", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AssignCredentials) toBody(ctx context.Context, state AssignCredentials) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.CliId.IsNull() {
		body, _ = sjson.Set(body, "cliCredentialsId.credentialsId", data.CliId.ValueString())
	}
	if !data.SnmpV2ReadId.IsNull() {
		body, _ = sjson.Set(body, "snmpv2cReadCredentialsId.credentialsId", data.SnmpV2ReadId.ValueString())
	}
	if !data.SnmpV2WriteId.IsNull() {
		body, _ = sjson.Set(body, "snmpv2cWriteCredentialsId.credentialsId", data.SnmpV2WriteId.ValueString())
	}
	if !data.SnmpV3Id.IsNull() {
		body, _ = sjson.Set(body, "snmpv3CredentialsId.credentialsId", data.SnmpV3Id.ValueString())
	}
	if !data.HttpsReadId.IsNull() {
		body, _ = sjson.Set(body, "httpReadCredentialsId.credentialsId", data.HttpsReadId.ValueString())
	}
	if !data.HttpsWriteId.IsNull() {
		body, _ = sjson.Set(body, "httpWriteCredentialsId.credentialsId", data.HttpsWriteId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AssignCredentials) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.#(key=\"credential.cli\").value.0.objReferences.0"); value.Exists() {
		data.CliId = types.StringValue(value.String())
	} else {
		data.CliId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.snmp_v2_read\").value.0.objReferences.0"); value.Exists() {
		data.SnmpV2ReadId = types.StringValue(value.String())
	} else {
		data.SnmpV2ReadId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.snmp_v2_write\").value.0.objReferences.0"); value.Exists() {
		data.SnmpV2WriteId = types.StringValue(value.String())
	} else {
		data.SnmpV2WriteId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.snmp_v3\").value.0.objReferences.0"); value.Exists() {
		data.SnmpV3Id = types.StringValue(value.String())
	} else {
		data.SnmpV3Id = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.http.read\").value.0.objReferences.0"); value.Exists() {
		data.HttpsReadId = types.StringValue(value.String())
	} else {
		data.HttpsReadId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.http.write\").value.0.objReferences.0"); value.Exists() {
		data.HttpsWriteId = types.StringValue(value.String())
	} else {
		data.HttpsWriteId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AssignCredentials) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.#(key=\"credential.cli\").value.0.objReferences.0"); value.Exists() && !data.CliId.IsNull() {
		data.CliId = types.StringValue(value.String())
	} else {
		data.CliId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.snmp_v2_read\").value.0.objReferences.0"); value.Exists() && !data.SnmpV2ReadId.IsNull() {
		data.SnmpV2ReadId = types.StringValue(value.String())
	} else {
		data.SnmpV2ReadId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.snmp_v2_write\").value.0.objReferences.0"); value.Exists() && !data.SnmpV2WriteId.IsNull() {
		data.SnmpV2WriteId = types.StringValue(value.String())
	} else {
		data.SnmpV2WriteId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.snmp_v3\").value.0.objReferences.0"); value.Exists() && !data.SnmpV3Id.IsNull() {
		data.SnmpV3Id = types.StringValue(value.String())
	} else {
		data.SnmpV3Id = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.http.read\").value.0.objReferences.0"); value.Exists() && !data.HttpsReadId.IsNull() {
		data.HttpsReadId = types.StringValue(value.String())
	} else {
		data.HttpsReadId = types.StringNull()
	}
	if value := res.Get("response.#(key=\"credential.http.write\").value.0.objReferences.0"); value.Exists() && !data.HttpsWriteId.IsNull() {
		data.HttpsWriteId = types.StringValue(value.String())
	} else {
		data.HttpsWriteId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssignCredentials) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.CliId.IsNull() {
		return false
	}
	if !data.SnmpV2ReadId.IsNull() {
		return false
	}
	if !data.SnmpV2WriteId.IsNull() {
		return false
	}
	if !data.SnmpV3Id.IsNull() {
		return false
	}
	if !data.HttpsReadId.IsNull() {
		return false
	}
	if !data.HttpsWriteId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
