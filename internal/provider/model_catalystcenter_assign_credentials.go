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
	Id                types.String `tfsdk:"id"`
	SiteId            types.String `tfsdk:"site_id"`
	CliId             types.String `tfsdk:"cli_id"`
	SnmpV2ReadId      types.String `tfsdk:"snmp_v2_read_id"`
	SnmpV2WriteId     types.String `tfsdk:"snmp_v2_write_id"`
	SnmpV3Id          types.String `tfsdk:"snmp_v3_id"`
	HttpsReadId       types.String `tfsdk:"https_read_id"`
	HttpsWriteId      types.String `tfsdk:"https_write_id"`
	PreserveUnmanaged types.Bool   `tfsdk:"preserve_unmanaged"`
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
// associated with a site). To fix this, on update (put) we send a top-level
// null for any slot that was previously set and is now removed. Per the API
// (PUT .../deviceCredentials, documented identically on 2.3.7.9-11, 3.1.5,
// 3.2.2 and 3.3.x): a null value makes the slot inherit from the parent site,
// an empty object {} unsets it (no credential of that type is used), and
// {"credentialsId": "<id>"} sets it. So a top-level null unassigns the
// site-local credential (letting the underlying credential object be deleted)
// while keeping the site consistent with the credentials declared higher in the
// hierarchy - which matches how the data model expresses "no override here,
// inherit from the parent". Slots that were never set are still omitted, so
// other inherited credentials (or credentials managed by another assignment)
// are left untouched.
//
// NOTE: an earlier attempt sent {"credentialsId": null} for the clear path. That
// is NOT the documented inherit form and Catalyst Center treats it as "unset"
// (the slot becomes null/none instead of inheriting from the parent), so it must
// not be used here.
func (data AssignCredentials) toBody(ctx context.Context, state AssignCredentials) string {
	body := ""
	put := state.Id.ValueString() != ""

	setSlot := func(slot string, plan, prior types.String) {
		if !plan.IsNull() {
			body, _ = sjson.Set(body, slot+".credentialsId", plan.ValueString())
		} else if put && !prior.IsNull() {
			// Slot was assigned and is now being cleared: send a top-level null
			// so the slot inherits from the parent site (matching the data
			// model), rather than {} which would leave it unset/none. Other
			// slots are omitted so unrelated inherited credentials stay intact.
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

// NOTE: toBodyGlobal is maintained manually (no generator markers) on purpose.
// The Global site is the common-settings root: it has no parent to inherit
// from, so Catalyst Center rejects a partial credential body (only some slots
// present) with NCND01090 on 2.3.7.x and NCND10603 ("site not found") on
// 3.2.2/3.2.3.
// This builds a complete body with every slot present so the write is accepted
// at the root, and is used only as the Create/Update fallback after a
// Global-root rejection. It mirrors the full-payload retry the Delete method
// performs for the Global site.
//
// Per slot the value is chosen from the plan (data), the prior state, the
// current controller assignment (current, from a GET; nil when not preserving)
// and the preserve flag:
//   - plan sets the slot            -> {"credentialsId": "<id>"}
//   - plan null, preserve, never    -> {"credentialsId": "<current id>"} (keep a
//     managed by TF (null in state)    credential configured outside Terraform)
//     and currently set on controller
//   - otherwise (not preserving, or -> {} ("unset"/none; nothing to inherit at
//     slot explicitly removed, i.e.    the root). Keeping removals as {} lets a
//     non-null in prior state, or      previously managed slot be cleared.
//     nothing to preserve)
func (data AssignCredentials) toBodyGlobal(ctx context.Context, state AssignCredentials, current *AssignCredentials, preserve bool) string {
	body := ""

	var cur AssignCredentials
	if current != nil {
		cur = *current
	}

	setSlot := func(slot string, plan, prior, curVal types.String) {
		switch {
		case !plan.IsNull():
			body, _ = sjson.Set(body, slot+".credentialsId", plan.ValueString())
		case preserve && prior.IsNull() && !curVal.IsNull():
			// Never managed by Terraform and currently set on the controller:
			// preserve the out-of-band credential.
			body, _ = sjson.Set(body, slot+".credentialsId", curVal.ValueString())
		default:
			body, _ = sjson.SetRaw(body, slot, "{}")
		}
	}

	setSlot("cliCredentialsId", data.CliId, state.CliId, cur.CliId)
	setSlot("snmpv2cReadCredentialsId", data.SnmpV2ReadId, state.SnmpV2ReadId, cur.SnmpV2ReadId)
	setSlot("snmpv2cWriteCredentialsId", data.SnmpV2WriteId, state.SnmpV2WriteId, cur.SnmpV2WriteId)
	setSlot("snmpv3CredentialsId", data.SnmpV3Id, state.SnmpV3Id, cur.SnmpV3Id)
	setSlot("httpReadCredentialsId", data.HttpsReadId, state.HttpsReadId, cur.HttpsReadId)
	setSlot("httpWriteCredentialsId", data.HttpsWriteId, state.HttpsWriteId, cur.HttpsWriteId)

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

// NOTE: updateFromBody is maintained manually (no generator markers) on purpose.
// It only reconciles the six credential slots from the deviceCredentials GET.
// The preserve_unmanaged attribute is a control-only flag that Catalyst Center
// never returns; the generated version would read it from the response, find it
// absent, and reset it to null on every refresh, producing a permanent diff.
// Keeping this hand-maintained leaves preserve_unmanaged untouched by Read.
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

// End of custom updateFromBody.

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
	if !data.PreserveUnmanaged.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
