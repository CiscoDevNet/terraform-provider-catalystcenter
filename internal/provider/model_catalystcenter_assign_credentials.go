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

// NOTE: toBody is maintained manually (no generator markers) on purpose.
// The generated version only sets slots whose ID is non-null, which means an
// in-place update that clears a single credential (set -> null) produces a PUT
// body that omits that slot, so Catalyst Center leaves the existing assignment
// untouched. The credential object can then not be deleted (NCIM01100 - still
// associated with a site). To fix this, on update (put) we send null for any
// slot that was previously set and is now removed. Per the API
// (PUT .../deviceCredentials), a null value makes the slot inherit from the
// parent site: this unassigns the site-local credential (so the underlying
// credential object can be deleted) while keeping the site consistent with the
// credentials declared higher in the hierarchy - which matches how the data
// model expresses "no override here, inherit from the parent". Slots that were
// never set are still omitted, so other inherited credentials (or credentials
// managed by another assignment) are left untouched. (An empty object {}, by
// contrast, marks the slot as "unset"/none instead of inheriting, so it is not
// used for clearing here.)
func (data AssignCredentials) toBody(ctx context.Context, state AssignCredentials) string {
	body := ""
	put := state.Id.ValueString() != ""

	setSlot := func(slot string, plan, prior types.String) {
		if !plan.IsNull() {
			body, _ = sjson.Set(body, slot+".credentialsId", plan.ValueString())
		} else if put && !prior.IsNull() {
			// Slot was assigned and is now being cleared: send null so the slot
			// inherits from the parent site (matching the data model), rather than
			// {} which would leave it unset/none. Other slots are omitted so
			// unrelated inherited credentials stay intact.
			body, _ = sjson.SetRaw(body, slot, "null")
		}
	}

	setSlot("cliCredentialsId", data.CliId, state.CliId)
	setSlot("snmpv2cReadCredentialsId", data.SnmpV2ReadId, state.SnmpV2ReadId)
	setSlot("snmpv2cWriteCredentialsId", data.SnmpV2WriteId, state.SnmpV2WriteId)
	setSlot("snmpv3CredentialsId", data.SnmpV3Id, state.SnmpV3Id)
	setSlot("httpReadCredentialsId", data.HttpsReadId, state.HttpsReadId)
	setSlot("httpWriteCredentialsId", data.HttpsWriteId, state.HttpsWriteId)

	return body
}

// NOTE: fromBody is maintained manually (no generator markers) on purpose.
// Catalyst Center returns an empty-string objReference ("") for a slot that has
// no locally assigned credential (e.g. HTTP read/write on a site that inherits
// or does not set them). The generated version treats "" as a real value
// (types.StringValue("")), which then differs from a null configuration and
// produces a permanent diff on every refresh. An empty objReference is never a
// valid credential id, so we treat it the same as an absent one, i.e. as null.
func (data *AssignCredentials) fromBody(ctx context.Context, res gjson.Result) {
	readSlot := func(key string) types.String {
		if value := res.Get("response.#(key=\"" + key + "\").value.0.objReferences.0"); value.Exists() && value.String() != "" {
			return types.StringValue(value.String())
		}
		return types.StringNull()
	}

	data.CliId = readSlot("credential.cli")
	data.SnmpV2ReadId = readSlot("credential.snmp_v2_read")
	data.SnmpV2WriteId = readSlot("credential.snmp_v2_write")
	data.SnmpV3Id = readSlot("credential.snmp_v3")
	data.HttpsReadId = readSlot("credential.http.read")
	data.HttpsWriteId = readSlot("credential.http.write")
}

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
