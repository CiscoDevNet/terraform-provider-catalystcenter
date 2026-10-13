//go:build ignore

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Run with: go test gen/generator.go gen/generator_singleton_test.go
// Compile and execute generated read methods, including computed values and
// three collection levels, without requiring a controller.
func TestGeneratedSingletonModels(t *testing.T) {
	root, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "gen/templates/model.go")); err != nil {
		root = filepath.Dir(root)
	}
	for _, collectionType := range []string{"List", "Set"} {
		t.Run(collectionType, func(t *testing.T) {
			config := YamlConfig{Name: "Singleton Fixture", UpdateComputed: true, Attributes: []YamlConfigAttribute{
				{ModelName: "entries", DataPath: "response", Type: collectionType, MinList: 1, MaxList: 1, Attributes: []YamlConfigAttribute{
					{ModelName: "name", Type: "String", Id: true},
					{ModelName: "requestOnly", Type: "String", WriteOnly: true},
					{ModelName: "constant", Type: "Int64", Value: "6"},
					{ModelName: "uuid", Type: "String", Computed: true},
					{ModelName: "children", Type: collectionType, Computed: true, MaxList: 1, Attributes: []YamlConfigAttribute{
						{ModelName: "name", Type: "String", Id: true},
						{ModelName: "uuid", Type: "String", Computed: true},
						{ModelName: "grandchildren", Type: collectionType, MaxList: 1, Attributes: []YamlConfigAttribute{
							{ModelName: "name", Type: "String", Id: true},
							{ModelName: "uuid", Type: "String", Computed: true},
						}},
					}},
				}},
			}}
			augmentConfig(&config)
			dir := t.TempDir()
			modelPath := filepath.Join(dir, "model.go")
			renderTemplate(filepath.Join(root, "gen/templates/model.go"), modelPath, config)
			model, err := os.ReadFile(modelPath)
			if err != nil {
				t.Fatal(err)
			}
			// Isolate read methods so unused imports from unrelated templates do not
			// require running goimports in each fixture.
			source := "package provider\nimport (\"context\"; \"github.com/hashicorp/terraform-plugin-framework/types\"; \"github.com/tidwall/gjson\")\n"
			for _, section := range []string{"types", "fromBody", "updateFromBody", "fromBodyUnknowns"} {
				source += getTemplateSection(string(model), section)
			}
			if err := os.WriteFile(modelPath, []byte(source), 0600); err != nil {
				t.Fatal(err)
			}
			for _, filename := range []string{"go.mod", "go.sum"} {
				data, err := os.ReadFile(filepath.Join(root, filename))
				if err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(dir, filename), data, 0600); err != nil {
					t.Fatal(err)
				}
			}
			if err := os.WriteFile(filepath.Join(dir, "model_test.go"), []byte(singletonRuntimeTests), 0600); err != nil {
				t.Fatal(err)
			}
			cmd := exec.Command("go", "test", "-mod=readonly", ".")
			cmd.Dir = dir
			if output, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("generated %s model failed: %v\n%s", collectionType, err, output)
			}
			// Verify validators follow the collection type at every nesting level.
			resourcePath := filepath.Join(dir, "resource.go")
			renderTemplate(filepath.Join(root, "gen/templates/resource.go"), resourcePath, config)
			resource, err := os.ReadFile(resourcePath)
			if err != nil {
				t.Fatal(err)
			}
			if n := strings.Count(string(resource), "Validators: []validator."+collectionType+"{"); n != 3 {
				t.Fatalf("expected three %s validator blocks, got %d", collectionType, n)
			}
		})
	}
}

const singletonRuntimeTests = `package provider
import (
    "context"
    "testing"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/tidwall/gjson"
)
func fixture() SingletonFixture {
    return SingletonFixture{Entries: []SingletonFixtureEntries{{
        Name: types.StringValue("before"), RequestOnly: types.StringValue("keep"), Uuid: types.StringUnknown(),
        Children: []SingletonFixtureEntriesChildren{{Name: types.StringValue("before"), Uuid: types.StringUnknown(),
            Grandchildren: []SingletonFixtureEntriesChildrenGrandchildren{{Name: types.StringValue("before"), Uuid: types.StringUnknown()}},
        }},
    }}}
}
const response = "{\"response\":{\"entries\":[{\"name\":\"after\",\"uuid\":\"parent\",\"children\":[{\"name\":\"after\",\"uuid\":\"child\",\"grandchildren\":[{\"name\":\"after\",\"uuid\":\"grandchild\"}]}]}]}}"
func TestRefresh(t *testing.T) {
    data := fixture()
    data.updateFromBody(context.Background(), gjson.Parse(response))
    row := data.Entries[0]
    if row.Name.ValueString() != "after" || row.Children[0].Name.ValueString() != "after" || row.Children[0].Grandchildren[0].Name.ValueString() != "after" { t.Fatal("singleton matching used mutable keys") }
    if row.RequestOnly.ValueString() != "keep" { t.Fatal("lost request-only value") }
    empty := gjson.Parse("{\"response\":{\"entries\":[{\"name\":\"after\",\"children\":[]}]}}")
    data.updateFromBody(context.Background(), empty)
    if len(data.Entries[0].Children) != 0 { t.Fatal("empty optional singleton was not cleared") }
}
func TestComputed(t *testing.T) {
    data := fixture()
    data.fromBodyUnknowns(context.Background(), gjson.Parse(response))
    row := data.Entries[0]
    if row.Uuid.ValueString() != "parent" || row.Children[0].Uuid.ValueString() != "child" || row.Children[0].Grandchildren[0].Uuid.ValueString() != "grandchild" { t.Fatal("computed IDs were not populated") }
    if row.Name.ValueString() != "before" || row.RequestOnly.ValueString() != "keep" { t.Fatal("known configured fields changed") }
}
func TestInvalidNestedResponse(t *testing.T) {
    data := fixture()
    response := gjson.Parse("{\"response\":{\"entries\":[{\"name\":\"changed\",\"children\":[{},{}]}]}}")
    data.updateFromBody(context.Background(), response)
    if len(data.Entries[0].Children) != 0 { t.Fatal("ambiguous singleton selected an arbitrary entry") }
}
`
