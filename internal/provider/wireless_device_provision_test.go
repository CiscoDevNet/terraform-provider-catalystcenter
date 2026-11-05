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
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cc "github.com/netascode/go-catalystcenter"
)

type wirelessDeviceProvisionTransport func(*http.Request) (*http.Response, error)

func (f wirelessDeviceProvisionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestWirelessDeviceProvisionDelete(t *testing.T) {
	const deviceID = "c21c291b-c2fe-483b-b747-a60528db0576"
	for _, test := range []struct {
		name       string
		status     int
		body       string
		getStatus  int
		getBody    string
		skipDelete bool
		wantError  bool
	}{
		{name: "success", status: http.StatusOK, body: `{}`},
		{name: "not found", status: http.StatusNotFound, body: `{}`},
		{name: "deleted controller", status: http.StatusBadRequest, body: `{"response":{"errorCode":"NCHS20416","message":"Bad Request","detail":"Device Delete failed. Field networkDeviceId is invalid."},"version":"1.0"}`},
		{name: "other bad request", status: http.StatusBadRequest, body: `{"response":{"errorCode":"NCHS20599"}}`, wantError: true},
		{name: "bad request without code", status: http.StatusBadRequest, body: `{}`, wantError: true},
		{name: "code only in detail", status: http.StatusBadRequest, body: `{"response":{"errorCode":"OTHER","detail":"NCHS20416"}}`, wantError: true},
		{name: "same code with forbidden status", status: http.StatusForbidden, body: `{"response":{"errorCode":"NCHS20416"}}`, wantError: true},
		{name: "server error", status: http.StatusInternalServerError, body: `{"response":{"errorCode":"NCHS20416"}}`, wantError: true},
		{name: "already absent on lookup", getStatus: http.StatusOK, getBody: `{"response":[]}`, skipDelete: true},
		{name: "not found on lookup", getStatus: http.StatusNotFound, getBody: `{}`, skipDelete: true},
		{name: "invalid device on lookup", getStatus: http.StatusBadRequest, getBody: `{"response":{"errorCode":"NCHS20416"}}`, skipDelete: true},
		{name: "lookup forbidden", getStatus: http.StatusForbidden, getBody: `{}`, skipDelete: true, wantError: true},
		{name: "lookup other bad request", getStatus: http.StatusBadRequest, getBody: `{"response":{"errorCode":"OTHER"}}`, skipDelete: true, wantError: true},
		{name: "lookup wrong device", getStatus: http.StatusOK, getBody: `{"response":[{"networkDeviceId":"another-device"}]}`, skipDelete: true, wantError: true},
		{name: "lookup malformed response", getStatus: http.StatusOK, getBody: `{}`, skipDelete: true, wantError: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			client, err := cc.NewClient("https://controller.example", "", "", cc.MaxRetries(0))
			if err != nil {
				t.Fatal(err)
			}
			client.Token = "test-token"
			var methods []string
			client.HttpClient.Transport = wirelessDeviceProvisionTransport(func(req *http.Request) (*http.Response, error) {
				methods = append(methods, req.Method)
				if req.URL.Path != "/dna/intent/api/v1/sda/provisionDevices" || req.URL.RawQuery != "networkDeviceId="+deviceID {
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL)
				}
				status, body := test.status, test.body
				switch req.Method {
				case http.MethodGet:
					status, body = test.getStatus, test.getBody
					if status == 0 {
						status = http.StatusOK
						body = `{"response":[{"id":"provisioning-object-id","networkDeviceId":"` + deviceID + `"}]}`
					}
				case http.MethodDelete:
					if test.skipDelete {
						t.Fatal("DELETE must not be called when lookup fails or device is absent")
					}
				default:
					t.Fatalf("unexpected method: %s", req.Method)
				}
				return &http.Response{StatusCode: status, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
			})
			r := &WirelessDeviceProvisionResource{client: &client}
			var schemaResponse resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
			state := tfsdk.State{Schema: schemaResponse.Schema}
			model := WirelessDeviceProvision{Id: types.StringValue(deviceID), NetworkDeviceId: types.StringValue(deviceID)}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatal(diags)
			}
			deleted := resource.DeleteResponse{State: state}
			r.Delete(ctx, resource.DeleteRequest{State: state}, &deleted)
			if deleted.Diagnostics.HasError() != test.wantError {
				t.Fatalf("expected error=%t, got diagnostics: %v", test.wantError, deleted.Diagnostics)
			}
			if test.wantError {
				if !deleted.State.Raw.Equal(state.Raw) {
					t.Fatal("failed deletion must retain the resource in state")
				}
			} else if !deleted.State.Raw.IsNull() {
				t.Fatal("successful deletion must remove the resource from state")
			}
			wantMethods := "GET,DELETE"
			if test.skipDelete {
				wantMethods = "GET"
			}
			if strings.Join(methods, ",") != wantMethods {
				t.Fatalf("expected methods %s, got %v", wantMethods, methods)
			}
		})
	}
}
