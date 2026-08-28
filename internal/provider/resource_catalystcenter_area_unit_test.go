// Copyright © 2026 Cisco Systems, Inc. and its affiliates.
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

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cc "github.com/netascode/go-catalystcenter"
)

func TestAreaReadRemovesMissingSiteFromState(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		fmt.Fprint(w, `{"response":[]}`)
	}))
	defer server.Close()

	client, err := cc.NewClient(server.URL, "", "", cc.MaxRetries(0))
	if err != nil {
		t.Fatal(err)
	}
	client.Token = "test"
	area := AreaResource{client: &client, cache: NewThreadSafeCache()}
	ctx := context.Background()
	var schemaResp resource.SchemaResponse
	area.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	newState := func() tfsdk.State {
		state := tfsdk.State{Schema: schemaResp.Schema}
		if diags := state.Set(ctx, &Area{Id: types.StringValue("missing")}); diags.HasError() {
			t.Fatal(diags.Errors())
		}
		return state
	}

	for _, name := range []string{"cold cache", "cache hit"} {
		t.Run(name, func(t *testing.T) {
			state := newState()
			resp := resource.ReadResponse{State: state}
			area.Read(ctx, resource.ReadRequest{State: state}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatal(resp.Diagnostics.Errors())
			}
			if !resp.State.Raw.IsNull() {
				t.Fatal("missing site was not removed from state")
			}
		})
	}

	if requests != 1 {
		t.Fatalf("expected one API request followed by a cache hit, got %d requests", requests)
	}
}
