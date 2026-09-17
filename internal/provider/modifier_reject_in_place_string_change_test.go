package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestRejectInPlaceStringChangePlanModifier(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		state     types.String
		plan      types.String
		wantError bool
	}{
		"add":                  {state: types.StringNull(), plan: types.StringValue("site-a")},
		"remove":               {state: types.StringValue("site-a"), plan: types.StringNull()},
		"unchanged":            {state: types.StringValue("site-a"), plan: types.StringValue("site-a")},
		"replace":              {state: types.StringValue("site-a"), plan: types.StringValue("site-b"), wantError: true},
		"empty state is unset": {state: types.StringValue(""), plan: types.StringValue("site-a")},
		"empty plan is unset":  {state: types.StringValue("site-a"), plan: types.StringValue("")},
		"unknown state":        {state: types.StringUnknown(), plan: types.StringValue("site-a")},
		"unknown plan":         {state: types.StringValue("site-a"), plan: types.StringUnknown()},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			request := planmodifier.StringRequest{
				Path:       path.Root("test_attribute"),
				StateValue: test.state,
				PlanValue:  test.plan,
			}
			response := &planmodifier.StringResponse{PlanValue: test.plan}

			rejectInPlaceStringChangePlanModifier{}.PlanModifyString(context.Background(), request, response)

			if response.Diagnostics.HasError() != test.wantError {
				t.Fatalf("HasError() = %t, want %t; diagnostics: %v", response.Diagnostics.HasError(), test.wantError, response.Diagnostics)
			}
		})
	}
}

func TestFabricL3VirtualNetworkAnchoredSiteUsesRejectInPlaceStringChangeModifier(t *testing.T) {
	t.Parallel()

	response := &resource.SchemaResponse{}
	(&FabricL3VirtualNetworkResource{}).Schema(context.Background(), resource.SchemaRequest{}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("Schema() returned errors: %v", response.Diagnostics)
	}

	attribute, ok := response.Schema.Attributes["anchored_site_id"].(resourceschema.StringAttribute)
	if !ok {
		t.Fatalf("anchored_site_id has type %T, want schema.StringAttribute", response.Schema.Attributes["anchored_site_id"])
	}
	if len(attribute.PlanModifiers) != 1 {
		t.Fatalf("anchored_site_id has %d plan modifiers, want 1", len(attribute.PlanModifiers))
	}
	if _, ok := attribute.PlanModifiers[0].(rejectInPlaceStringChangePlanModifier); !ok {
		t.Fatalf("anchored_site_id plan modifier has type %T, want rejectInPlaceStringChangePlanModifier", attribute.PlanModifiers[0])
	}
}
