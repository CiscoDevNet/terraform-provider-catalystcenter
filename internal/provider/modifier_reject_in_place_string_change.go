package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rejectInPlaceStringChangePlanModifier prevents replacing one configured string
// value with another while still permitting the value to be added or removed.
// This is useful for APIs that support enable/disable transitions but cannot move
// an existing association directly between two targets.
type rejectInPlaceStringChangePlanModifier struct{}

func (m rejectInPlaceStringChangePlanModifier) Description(context.Context) string {
	return "Rejects changing a configured string directly to a different configured value."
}

func (m rejectInPlaceStringChangePlanModifier) MarkdownDescription(context.Context) string {
	return "Rejects changing a configured string directly to a different configured value. Adding and removing the value remain allowed."
}

func (m rejectInPlaceStringChangePlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !isInPlaceStringChange(req.StateValue, req.PlanValue) {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"In-place value change is not supported",
		"This attribute cannot be changed directly from one configured value to another. Remove the current value and apply first, then configure the new value in a subsequent apply.",
	)
}

func isInPlaceStringChange(state, plan types.String) bool {
	if state.IsNull() || state.IsUnknown() || plan.IsNull() || plan.IsUnknown() {
		return false
	}

	stateValue := state.ValueString()
	planValue := plan.ValueString()
	return stateValue != "" && planValue != "" && stateValue != planValue
}
