package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	cc "github.com/netascode/go-catalystcenter"
)

func newSiteSettingsTestClient(t *testing.T) *cc.Client {
	t.Helper()
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/dna/system/api/v1/auth/token":
			_, _ = w.Write([]byte(`{"Token": "test"}`))
		case r.URL.Path == "/dna/intent/api/v1/site" && r.URL.Query().Get("siteId") == "global":
			_, _ = w.Write([]byte(`{"response": {"id": "global", "name": "Global"}}`))
		case r.URL.Path == "/dna/intent/api/v1/site" && r.URL.Query().Get("siteId") == "child":
			_, _ = w.Write([]byte(`{"response": {"id": "child", "parentId": "parent"}}`))
		case r.URL.Path == "/dna/intent/api/v1/site" && r.URL.Query().Get("siteId") == "deleted":
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"response": {"errorCode": "NCGR10008", "message": "NCGR10008: Group with UUID deleted does not exist"}}`))
		case r.URL.Path == "/dna/intent/api/v1/sites/parent/aaaSettings" && r.URL.Query().Get("_inherited") == "true":
			_, _ = w.Write([]byte(`{"response": {"aaaNetwork": {"serverType": "AAA", "protocol": "RADIUS", "primaryServerIp": "192.0.2.1", "inheritedSiteId": "global", "inheritedSiteName": "Global"}, "aaaClient": null}}`))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)
	client, err := cc.NewClient(server.URL, "user", "test", cc.MaxRetries(0))
	if err != nil {
		t.Fatal(err)
	}
	return &client
}

func TestSiteSettingsDeleteBody(t *testing.T) {
	client := newSiteSettingsTestClient(t)
	const endpoint = "/dna/intent/api/v1/sites/%v/aaaSettings"
	const globalBody = `{"aaaNetwork": {}, "aaaClient": {}}`
	// Not served by the test server: an inherit body must not read the parent's settings.
	const unservedEndpoint = "/dna/intent/api/v1/sites/%v/dnsSettings"
	const inheritBody = `{"dns": null}`

	tests := []struct {
		name        string
		siteId      string
		inheritBody string
		want        string
	}{
		{"global unsets", "global", "", globalBody},
		{"deleted site keeps delete body", "deleted", "", globalBody},
		{"child gets parent values", "child", "", `{"aaaNetwork":{"serverType":"AAA","protocol":"RADIUS","primaryServerIp":"192.0.2.1"},"aaaClient":null}`},
		{"global ignores inherit body", "global", inheritBody, globalBody},
		{"deleted site ignores inherit body", "deleted", inheritBody, globalBody},
		{"child gets inherit body", "child", inheritBody, inheritBody},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := endpoint
			if tt.inheritBody != "" {
				ep = unservedEndpoint
			}
			got, err := siteSettingsDeleteBody(client, ep, tt.siteId, globalBody, tt.inheritBody)
			if err != nil {
				t.Fatal(err)
			}
			var gotJson, wantJson any
			if err := json.Unmarshal([]byte(got), &gotJson); err != nil {
				t.Fatalf("invalid JSON %s: %v", got, err)
			}
			_ = json.Unmarshal([]byte(tt.want), &wantJson)
			if !reflect.DeepEqual(gotJson, wantJson) {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSiteSettingsDeleteBodyLookupError(t *testing.T) {
	client := newSiteSettingsTestClient(t)
	if _, err := siteSettingsDeleteBody(client, "/dna/intent/api/v1/sites/%v/aaaSettings", "missing", "{}", ""); err == nil {
		t.Error("expected error for unknown site")
	}
}
