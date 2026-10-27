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

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

type qosPolicySettingTransport func(*http.Request) (*http.Response, error)

func (f qosPolicySettingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestQoSPolicySettingLifecycle(t *testing.T) {
	ctx := context.Background()
	client, err := cc.NewClient("https://controller.example", "", "")
	if err != nil {
		t.Fatal(err)
	}
	client.Token = "test-token"
	var enabled bool
	var methods []string
	client.HttpClient.Transport = qosPolicySettingTransport(func(req *http.Request) (*http.Response, error) {
		methods = append(methods, req.Method)
		if req.URL.Path != "/dna/intent/api/v1/qosPolicySetting" || req.URL.RawQuery != "" {
			t.Fatalf("unexpected request URL: %s", req.URL)
		}
		body := ""
		status := http.StatusOK
		switch req.Method {
		case http.MethodPut:
			payload, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			parsed := gjson.ParseBytes(payload)
			if len(parsed.Map()) != 1 || parsed.Get("deployByDefaultOnWiredDevices").Type != gjson.True && parsed.Get("deployByDefaultOnWiredDevices").Type != gjson.False {
				t.Fatalf("unexpected PUT body: %s", payload)
			}
			enabled = parsed.Get("deployByDefaultOnWiredDevices").Bool()
			status = http.StatusNoContent
		case http.MethodGet:
			body = fmt.Sprintf(`{"response":{"deployByDefaultOnWiredDevices":%t},"version":"1.0"}`, enabled)
		default:
			t.Fatalf("unexpected HTTP method: %s", req.Method)
		}
		return &http.Response{StatusCode: status, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
	})
	r := &QoSPolicySettingResource{client: &client}
	var schemaResponse resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	if len(schemaResponse.Schema.Attributes) != 3 {
		t.Fatal("resource should expose id, name, and the deployment flag")
	}
	plan := tfsdk.Plan{Schema: schemaResponse.Schema}
	const expectedID = "qos_policy_setting"
	model := QoSPolicySetting{Id: types.StringUnknown(), Name: types.StringValue(expectedID), DeployByDefaultOnWiredDevices: types.BoolValue(true)}
	if diags := plan.Set(ctx, &model); diags.HasError() {
		t.Fatal(diags)
	}
	create := resource.CreateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	r.Create(ctx, resource.CreateRequest{Plan: plan}, &create)
	if create.Diagnostics.HasError() {
		t.Fatal(create.Diagnostics)
	}
	if diags := create.State.Get(ctx, &model); diags.HasError() {
		t.Fatal(diags)
	}
	if model.Id.ValueString() != expectedID || !enabled {
		t.Fatalf("create did not set the singleton ID and deployment flag: %+v", model)
	}
	model.DeployByDefaultOnWiredDevices = types.BoolValue(false)
	if diags := plan.Set(ctx, &model); diags.HasError() {
		t.Fatal(diags)
	}
	update := resource.UpdateResponse{State: create.State}
	r.Update(ctx, resource.UpdateRequest{Plan: plan, State: create.State}, &update)
	if update.Diagnostics.HasError() || enabled {
		t.Fatalf("update failed to disable deployment: %v", update.Diagnostics)
	}
	// A refresh must detect a change made outside Terraform.
	enabled = true
	refresh := resource.ReadResponse{State: update.State}
	r.Read(ctx, resource.ReadRequest{State: update.State}, &refresh)
	if refresh.Diagnostics.HasError() {
		t.Fatal(refresh.Diagnostics)
	}
	if diags := refresh.State.Get(ctx, &model); diags.HasError() {
		t.Fatal(diags)
	}
	if !model.DeployByDefaultOnWiredDevices.ValueBool() || model.Id.ValueString() != expectedID {
		t.Fatal("refresh did not retain the ID and read controller drift")
	}
	// Import starts with only the fixed ID, then reads the controller flag.
	imported := resource.ImportStateResponse{State: tfsdk.State{Schema: schemaResponse.Schema, Raw: tftypes.NewValue(schemaResponse.Schema.Type().TerraformType(ctx), nil)}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: expectedID}, &imported)
	if imported.Diagnostics.HasError() {
		t.Fatal(imported.Diagnostics)
	}
	readImport := resource.ReadResponse{State: imported.State}
	r.Read(ctx, resource.ReadRequest{State: imported.State}, &readImport)
	if readImport.Diagnostics.HasError() {
		t.Fatal(readImport.Diagnostics)
	}
	if !readImport.State.Raw.Equal(refresh.State.Raw) {
		t.Fatal("imported state differs from refreshed state")
	}
	invalid := resource.ImportStateResponse{}
	r.ImportState(ctx, resource.ImportStateRequest{ID: ""}, &invalid)
	if !invalid.Diagnostics.HasError() {
		t.Fatal("import accepted an empty identifier")
	}
	d := &QoSPolicySettingDataSource{client: &client}
	var dataSchema datasource.SchemaResponse
	d.Schema(ctx, datasource.SchemaRequest{}, &dataSchema)
	if len(dataSchema.Schema.Attributes) != 3 {
		t.Fatal("data source should expose id, name, and the deployment flag")
	}
	dataRead := datasource.ReadResponse{State: tfsdk.State{Schema: dataSchema.Schema}}
	d.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: dataSchema.Schema, Raw: tftypes.NewValue(dataSchema.Schema.Type().TerraformType(ctx), map[string]tftypes.Value{
		"id":                                 tftypes.NewValue(tftypes.String, nil),
		"name":                               tftypes.NewValue(tftypes.String, expectedID),
		"deploy_by_default_on_wired_devices": tftypes.NewValue(tftypes.Bool, nil),
	})}}, &dataRead)
	if dataRead.Diagnostics.HasError() {
		t.Fatal(dataRead.Diagnostics)
	}
	if !dataRead.State.Raw.Equal(refresh.State.Raw) {
		t.Fatal("data source did not read the singleton state")
	}
	beforeDelete := len(methods)
	deleted := resource.DeleteResponse{State: refresh.State}
	r.Delete(ctx, resource.DeleteRequest{State: refresh.State}, &deleted)
	if deleted.Diagnostics.HasError() || !deleted.State.Raw.IsNull() || len(methods) != beforeDelete || !enabled {
		t.Fatal("destroy should remove only Terraform state")
	}
	if strings.Join(methods, ",") != "PUT,PUT,GET,GET,GET" {
		t.Fatalf("unexpected lifecycle requests: %v", methods)
	}
}
