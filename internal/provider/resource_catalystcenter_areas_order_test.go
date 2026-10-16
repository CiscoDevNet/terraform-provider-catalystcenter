package provider

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestAreasDeleteOrder(t *testing.T) {
	area := func(parent, name, id string) AreasAreas {
		a := AreasAreas{
			Name:                types.StringValue(name),
			ParentNameHierarchy: types.StringValue(parent),
			Id:                  types.StringNull(),
		}
		if id != "" {
			a.Id = types.StringValue(id)
		}
		return a
	}
	areas := map[string]AreasAreas{
		"Global/A":           area("Global", "A", "a"),
		"Global/A/B":         area("Global/A", "B", "b"),
		"Global/A/B/C":       area("Global/A/B", "C", "c"),
		"Global/A/B2":        area("Global/A", "B2", "b2"),
		"Global/A/B/C/D":     area("Global/A/B/C", "D", "d"),
		"Global/A/B/C/D/new": area("Global/A/B/C/D", "new", ""),
	}
	want := []string{"d", "c", "b", "b2", "a"}

	for i := 0; i < 100; i++ {
		var got []string
		for _, a := range areasDeleteOrder(areas) {
			got = append(got, a.Id.ValueString())
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("run %d: got %v, want %v", i, got, want)
		}
	}
}
