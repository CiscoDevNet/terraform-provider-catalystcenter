package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

type wirelessDeviceProvisionTransport func(*http.Request) (*http.Response, error)

func (f wirelessDeviceProvisionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestWirelessDeviceProvisionUpdate(t *testing.T) {
	iface := WirelessDeviceProvisionInterfaces{InterfaceName: types.StringValue("PSK_SSID"), VlanId: types.Int64Value(560)}
	oldIface := iface
	oldIface.VlanId = types.Int64Value(559)
	other := WirelessDeviceProvisionInterfaces{InterfaceName: types.StringValue("other"), VlanId: types.Int64Value(100)}
	optional := iface
	optional.InterfaceIpAddress = types.StringValue("10.0.0.2")
	optional.InterfaceNetmask = types.Int64Value(24)
	optional.InterfaceGateway = types.StringValue("10.0.0.1")
	optional.LagOrPortNumber = types.Int64Value(1)
	tests := []struct {
		name          string
		before, after []WirelessDeviceProvisionInterfaces
		reprovision   bool
		wantPost      bool
		wantEmpty     bool
		fail          bool
		errorCode     string
		wantWarning   bool
	}{
		{name: "false interface VLAN change", before: []WirelessDeviceProvisionInterfaces{oldIface}, after: []WirelessDeviceProvisionInterfaces{iface}, wantPost: true},
		{name: "false interface addition", after: []WirelessDeviceProvisionInterfaces{iface}, wantPost: true},
		{name: "explicit reprovision includes interfaces", before: []WirelessDeviceProvisionInterfaces{optional}, after: []WirelessDeviceProvisionInterfaces{optional}, reprovision: true, wantPost: true},
		{name: "unchanged false no POST", before: []WirelessDeviceProvisionInterfaces{iface}, after: []WirelessDeviceProvisionInterfaces{iface}},
		{name: "reorder no POST", before: []WirelessDeviceProvisionInterfaces{iface, other}, after: []WirelessDeviceProvisionInterfaces{other, iface}},
		{name: "optional interface change", before: []WirelessDeviceProvisionInterfaces{iface}, after: []WirelessDeviceProvisionInterfaces{optional}, wantPost: true},
		{name: "explicit empty fallback", reprovision: true, wantPost: true, wantEmpty: true},
		{name: "interface removal empty fallback", before: []WirelessDeviceProvisionInterfaces{iface}, wantPost: true, wantEmpty: true},
		{name: "client error retains state", before: []WirelessDeviceProvisionInterfaces{oldIface}, after: []WirelessDeviceProvisionInterfaces{iface}, wantPost: true, fail: true, errorCode: "INVALID"},
		{name: "unreachable interface change retains state", before: []WirelessDeviceProvisionInterfaces{oldIface}, after: []WirelessDeviceProvisionInterfaces{iface}, wantPost: true, fail: true, errorCode: "NCDP10000"},
		{name: "unreachable explicit reprovision warns", before: []WirelessDeviceProvisionInterfaces{iface}, after: []WirelessDeviceProvisionInterfaces{iface}, reprovision: true, wantPost: true, errorCode: "NCDP10000", wantWarning: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client, err := cc.NewClient("https://controller.example", "", "")
			if err != nil {
				t.Fatal(err)
			}
			client.Token = "test-token"
			posts := 0
			client.HttpClient.Transport = wirelessDeviceProvisionTransport(func(req *http.Request) (*http.Response, error) {
				posts++
				if req.Method != http.MethodPost || req.URL.Path != "/dna/intent/api/v1/wirelessControllers/wlc-id/provision" || req.URL.RawQuery != "" {
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL)
				}
				payload, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatal(err)
				}
				if tt.wantEmpty {
					if string(payload) != "{}" {
						t.Fatalf("expected empty object, got %s", payload)
					}
				} else {
					parsed := gjson.ParseBytes(payload)
					if parsed.Get("interfaces.0.interfaceName").String() != "PSK_SSID" || parsed.Get("interfaces.0.vlanId").Int() != 560 {
						t.Fatalf("missing planned interface: %s", payload)
					}
					if parsed.Get("reprovision").Exists() {
						t.Fatalf("reprovision leaked into body: %s", payload)
					}
					if !tt.after[0].InterfaceIpAddress.IsNull() {
						if parsed.Get("interfaces.0.interfaceIPAddress").String() != "10.0.0.2" || parsed.Get("interfaces.0.interfaceNetmaskInCIDR").Int() != 24 || parsed.Get("interfaces.0.interfaceGateway").String() != "10.0.0.1" || parsed.Get("interfaces.0.lagOrPortNumber").Int() != 1 {
							t.Fatalf("optional interface fields dropped: %s", payload)
						}
					}
					if !parsed.Get("skipApProvision").Bool() {
						t.Fatalf("optional provisioning field dropped: %s", payload)
					}
				}
				status, body := http.StatusOK, `{}`
				if tt.errorCode != "" {
					status, body = http.StatusBadRequest, fmt.Sprintf(`{"response":{"errorCode":%q,"failureReason":"bad configuration"}}`, tt.errorCode)
				}
				return &http.Response{StatusCode: status, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(body)), Request: req}, nil
			})
			r := &WirelessDeviceProvisionResource{client: &client}
			var schemaResponse resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
			before := WirelessDeviceProvision{Id: types.StringValue("wlc-id"), NetworkDeviceId: types.StringValue("wlc-id"), Reprovision: types.BoolValue(false), Interfaces: tt.before}
			after := before
			after.Interfaces = tt.after
			after.Reprovision = types.BoolValue(tt.reprovision)
			if !tt.wantEmpty {
				after.SkipApProvision = types.BoolValue(true)
			}
			state := tfsdk.State{Schema: schemaResponse.Schema}
			if diags := state.Set(ctx, &before); diags.HasError() {
				t.Fatal(diags)
			}
			plan := tfsdk.Plan{Schema: schemaResponse.Schema}
			if diags := plan.Set(ctx, &after); diags.HasError() {
				t.Fatal(diags)
			}
			response := resource.UpdateResponse{State: state}
			r.Update(ctx, resource.UpdateRequest{Plan: plan, State: state}, &response)
			wantPosts := 0
			if tt.wantPost {
				wantPosts = 1
			}
			if posts != wantPosts {
				t.Fatalf("POST count = %d, want %d", posts, wantPosts)
			}
			if response.Diagnostics.HasError() != tt.fail {
				t.Fatalf("unexpected diagnostics: %v", response.Diagnostics)
			}
			if (response.Diagnostics.WarningsCount() > 0) != tt.wantWarning {
				t.Fatalf("unexpected warning diagnostics: %v", response.Diagnostics)
			}
			if tt.fail {
				if !response.State.Raw.Equal(state.Raw) {
					t.Fatal("failed request saved planned state")
				}
			} else if !response.State.Raw.Equal(plan.Raw) {
				t.Fatal("successful update did not save plan")
			}
		})
	}
}
