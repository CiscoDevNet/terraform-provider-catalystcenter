package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Based on the controller response: type and ignoreConflict are absent, and
// engineId is a JSON string even though the configuration uses an integer.
const applicationSingletonResponse = `{"response":[{
	"name":"my-custom-app",
	"parentScalableGroup":{"idRef":"application-set-id"},
	"networkApplications":[{
		"id":"network-application-id",
		"appProtocol":"TCP","applicationType":"CUSTOM",
		"categoryId":"category-id","dscp":"18","engineId":"6",
		"helpString":"This is a custom application","name":"my-custom-app",
		"rank":1,"serverName":"app.example.com",
		"trafficClass":"TRANSACTIONAL_DATA","url":"example.com"
	}]
}]}`

func applicationSingletonState() Application {
	return Application{
		Name:             types.StringValue("my-custom-app"),
		ApplicationSetId: types.StringValue("application-set-id"),
		NetworkApplications: []ApplicationNetworkApplications{{
			HelpString:      types.StringValue("This is a custom application"),
			ApplicationType: types.StringValue("CUSTOM"),
			ServerType:      types.StringValue("_server-ip"),
			Dscp:            types.StringValue("18"),
			AppProtocol:     types.StringValue("TCP"),
			ServerName:      types.StringValue("app.example.com"),
			Url:             types.StringValue("example.com"),
			TrafficClass:    types.StringValue("TRANSACTIONAL_DATA"),
			CategoryId:      types.StringValue("category-id"),
			IgnoreConflict:  types.BoolValue(false),
			Rank:            types.Int64Value(1),
			EngineId:        types.Int64Value(6),
		}},
	}
}

func TestApplicationSingletonRefresh(t *testing.T) {
	state := applicationSingletonState()
	want := applicationSingletonState()
	for i := 0; i < 2; i++ {
		state.updateFromBody(context.Background(), gjson.Parse(applicationSingletonResponse))
		if !reflect.DeepEqual(state, want) {
			t.Fatalf("refresh %d changed unchanged configuration: got %+v, want %+v", i, state, want)
		}
	}
	// An out-of-band change must be read, even though it changes a former match key.
	body, err := sjson.Set(applicationSingletonResponse, "response.0.networkApplications.0.dscp", "24")
	if err != nil {
		t.Fatal(err)
	}
	state.updateFromBody(context.Background(), gjson.Parse(body))
	want.NetworkApplications[0].Dscp = types.StringValue("24")
	if !reflect.DeepEqual(state, want) {
		t.Fatalf("refresh did not preserve request-only fields while detecting drift: %+v", state)
	}
}

func TestApplicationSingletonMissingOrAmbiguousResponse(t *testing.T) {
	for _, value := range []string{`[]`, `null`, `[{},{}]`} {
		t.Run(value, func(t *testing.T) {
			state := applicationSingletonState()
			body, err := sjson.SetRaw(applicationSingletonResponse, "response.0.networkApplications", value)
			if err != nil {
				t.Fatal(err)
			}
			state.updateFromBody(context.Background(), gjson.Parse(body))
			if len(state.NetworkApplications) != 0 {
				t.Fatal("missing or ambiguous singleton must not select an arbitrary entry")
			}
		})
	}
}

func TestApplicationSingletonSetValidation(t *testing.T) {
	ctx := context.Background()
	var resp resource.SchemaResponse
	(&ApplicationResource{}).Schema(ctx, resource.SchemaRequest{}, &resp)
	if diags := resp.Schema.ValidateImplementation(ctx); diags.HasError() {
		t.Fatalf("invalid schema: %v", diags)
	}
	collection := resp.Schema.Attributes["network_applications"].(schema.SetNestedAttribute)
	for _, count := range []int{0, 1, 2} {
		values := make([]attr.Value, count)
		for i := range values {
			values[i] = types.Int64Value(int64(i))
		}
		// Cardinality validators only inspect the element count.
		request := validator.SetRequest{ConfigValue: types.SetValueMust(types.Int64Type, values)}
		var response validator.SetResponse
		for _, v := range collection.Validators {
			v.ValidateSet(ctx, request, &response)
		}
		if response.Diagnostics.HasError() != (count != 1) {
			t.Fatalf("count %d: unexpected validation result: %v", count, response.Diagnostics)
		}
	}
}
