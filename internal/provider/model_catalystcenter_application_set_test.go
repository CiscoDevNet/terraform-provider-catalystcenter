package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func TestApplicationSetImportReadsExternalHandle(t *testing.T) {
	ctx := context.Background()
	const id = "75820325-2c31-4cfa-8496-6eeaae1e628f"
	state := ApplicationSet{Id: types.StringValue(id)}
	response := gjson.Parse(`{
		"response": [{
			"id": "75820325-2c31-4cfa-8496-6eeaae1e628f",
			"name": "my-app-set2",
			"defaultBusinessRelevance": "BUSINESS_IRRELEVANT",
			"scalableGroupExternalHandle": "my-app-set2"
		}]
	}`).Get(`response.#(id=="` + id + `")`)

	if !state.isNull(ctx, response) {
		t.Fatal("UUID-only import should read all attributes")
	}
	state.fromBody(ctx, response)
	if state.Id.ValueString() != id {
		t.Fatalf("import changed the UUID to %q", state.Id.ValueString())
	}
	if !state.ScalableGroupExternalHandle.Equal(types.StringValue("my-app-set2")) {
		t.Fatalf("import did not restore the external handle: %s", state.ScalableGroupExternalHandle)
	}

	// Refresh must also detect an external change to this replacement attribute.
	state.updateFromBody(ctx, gjson.Parse(`{
		"name": "my-app-set2",
		"defaultBusinessRelevance": "BUSINESS_IRRELEVANT",
		"scalableGroupExternalHandle": "changed-handle"
	}`))
	if !state.ScalableGroupExternalHandle.Equal(types.StringValue("changed-handle")) {
		t.Fatalf("refresh did not read the external handle: %s", state.ScalableGroupExternalHandle)
	}
}
