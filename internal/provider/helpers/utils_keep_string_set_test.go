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

package helpers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func TestKeepStringSetInIgnoresAPIExtras(t *testing.T) {
	ctx := context.Background()
	state := types.SetValueMust(types.StringType, []attr.Value{
		types.StringValue("parent-a"),
		types.StringValue("parent-b"),
	})
	api := gjson.Parse(`["parent-a","child-1","parent-b","child-2"]`).Array()

	got := KeepStringSetIn(ctx, state, api)
	var gotVals []string
	if diags := got.ElementsAs(ctx, &gotVals, false); diags.HasError() {
		t.Fatalf("ElementsAs: %v", diags)
	}
	if !StringSetEqual(gotVals, []string{"parent-a", "parent-b"}) {
		t.Fatalf("got %v, want parent-a and parent-b", gotVals)
	}
}

func TestKeepStringSetInDropsValuesMissingFromAPI(t *testing.T) {
	ctx := context.Background()
	state := types.SetValueMust(types.StringType, []attr.Value{
		types.StringValue("keep"),
		types.StringValue("gone"),
	})
	api := gjson.Parse(`["keep"]`).Array()

	got := KeepStringSetIn(ctx, state, api)
	var gotVals []string
	if diags := got.ElementsAs(ctx, &gotVals, false); diags.HasError() {
		t.Fatalf("ElementsAs: %v", diags)
	}
	if !StringSetEqual(gotVals, []string{"keep"}) {
		t.Fatalf("got %v, want keep only", gotVals)
	}
}

func TestKeepStringSetInExactMatchIsNoOp(t *testing.T) {
	ctx := context.Background()
	state := types.SetValueMust(types.StringType, []attr.Value{
		types.StringValue("parent-a"),
		types.StringValue("parent-b"),
	})
	api := gjson.Parse(`["parent-b","parent-a"]`).Array()

	got := KeepStringSetIn(ctx, state, api)
	var gotVals []string
	if diags := got.ElementsAs(ctx, &gotVals, false); diags.HasError() {
		t.Fatalf("ElementsAs: %v", diags)
	}
	if !StringSetEqual(gotVals, []string{"parent-a", "parent-b"}) {
		t.Fatalf("got %v, want exact configured set", gotVals)
	}
}
