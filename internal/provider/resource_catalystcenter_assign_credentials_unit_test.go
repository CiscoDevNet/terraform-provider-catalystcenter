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

// This file is hand-maintained (not generated) and covers the manual toBody
// clearing logic of catalystcenter_assign_credentials, which only unassigns the
// slots Terraform manages (so credentials inherited from a parent site are not
// wiped) and sets them to null (the API's "inherit from parent" form) to clear.

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func ccStr(s string) types.String { return types.StringValue(s) }

// assertNullSlot asserts the slot exists and its value is JSON null (the API's
// "inherit from the parent site" form).
func assertNullSlot(t *testing.T, body, slot string) {
	t.Helper()
	v := gjson.Get(body, slot)
	if !v.Exists() {
		t.Fatalf("expected slot %q to be present, body=%s", slot, body)
	}
	if v.Type != gjson.Null {
		t.Fatalf("expected slot %q to be JSON null, got %q, body=%s", slot, v.Raw, body)
	}
}

// assertSetSlot asserts the slot's credentialsId equals the given id.
func assertSetSlot(t *testing.T, body, slot, id string) {
	t.Helper()
	v := gjson.Get(body, slot+".credentialsId")
	if v.String() != id {
		t.Fatalf("expected slot %q credentialsId %q, got %q, body=%s", slot, id, v.String(), body)
	}
}

// assertAbsentSlot asserts the slot is not present at all (left to inherit).
func assertAbsentSlot(t *testing.T, body, slot string) {
	t.Helper()
	if gjson.Get(body, slot).Exists() {
		t.Fatalf("expected slot %q to be absent (inherited), body=%s", slot, body)
	}
}

// On Create (no prior state) only the configured slots are sent; nothing is
// null-cleared and unset slots are omitted so inheritance is untouched.
func TestAssignCredentialsToBodyCreate(t *testing.T) {
	plan := AssignCredentials{CliId: ccStr("CLI"), HttpsReadId: ccStr("HR")}
	body := plan.toBody(context.Background(), AssignCredentials{})

	assertSetSlot(t, body, "cliCredentialsId", "CLI")
	assertSetSlot(t, body, "httpReadCredentialsId", "HR")
	assertAbsentSlot(t, body, "snmpv2cReadCredentialsId")
	assertAbsentSlot(t, body, "snmpv2cWriteCredentialsId")
	assertAbsentSlot(t, body, "snmpv3CredentialsId")
	assertAbsentSlot(t, body, "httpWriteCredentialsId")
}

// On Update that clears a single managed slot, only that slot is sent as null
// (inherit); the still-configured slot keeps its value and never-set (inherited)
// slots stay omitted.
func TestAssignCredentialsToBodyUpdateClearOne(t *testing.T) {
	state := AssignCredentials{Id: ccStr("site1"), CliId: ccStr("CLI"), HttpsReadId: ccStr("HR")}
	plan := AssignCredentials{Id: ccStr("site1"), HttpsReadId: ccStr("HR")} // cli removed
	body := plan.toBody(context.Background(), state)

	assertNullSlot(t, body, "cliCredentialsId")
	assertSetSlot(t, body, "httpReadCredentialsId", "HR")
	assertAbsentSlot(t, body, "snmpv3CredentialsId")
	assertAbsentSlot(t, body, "snmpv2cReadCredentialsId")
}

// The Delete body (empty plan against current state) unassigns only the slots
// Terraform manages (non-null in state), each as null (inherit); inherited slots
// are omitted so they are preserved.
func TestAssignCredentialsDeleteBody(t *testing.T) {
	state := AssignCredentials{Id: ccStr("site1"), CliId: ccStr("CLI"), HttpsReadId: ccStr("HR")}
	var empty AssignCredentials
	body := empty.toBody(context.Background(), state)

	assertNullSlot(t, body, "cliCredentialsId")
	assertNullSlot(t, body, "httpReadCredentialsId")
	assertAbsentSlot(t, body, "snmpv2cReadCredentialsId")
	assertAbsentSlot(t, body, "snmpv2cWriteCredentialsId")
	assertAbsentSlot(t, body, "snmpv3CredentialsId")
	assertAbsentSlot(t, body, "httpWriteCredentialsId")
}

// When state has no managed slots, the clear body is empty so Delete can skip
// the PUT entirely instead of clearing all six slots.
func TestAssignCredentialsDeleteBodyEmpty(t *testing.T) {
	state := AssignCredentials{Id: ccStr("site1")}
	var empty AssignCredentials
	body := empty.toBody(context.Background(), state)
	if body != "" {
		t.Fatalf("expected empty clear body when no slots are managed, got %s", body)
	}
}
