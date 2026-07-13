// modifier_fabric_devices_set_plan.go
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// wirelessControllerAutoRoles are fabric roles that Catalyst Center adds on its own
// after provisioning (e.g. an embedded WLC on a Fabric-in-a-Box edge node). The API
// only ever returns WIRELESS_CONTROLLER_NODE, and it never removes it, so a config
// that omits the role produces a permanent diff. This modifier tolerates it.
var wirelessControllerAutoRoles = map[string]bool{
	"WIRELESS_CONTROLLER_NODE": true,
}

type fabricDevicesSetPlanModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m fabricDevicesSetPlanModifier) Description(ctx context.Context) string {
	return "Suppresses the perpetual diff on device_roles when Catalyst Center auto-adds the WIRELESS_CONTROLLER_NODE role."
}

// MarkdownDescription returns a markdown-formatted description of the plan modifier.
func (m fabricDevicesSetPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "For each planned fabric device (matched to prior state by `network_device_id`), if the state's " +
		"`device_roles` differ from the plan only by the controller-added `WIRELESS_CONTROLLER_NODE` role, " +
		"the state value is copied into the plan so no spurious update is generated. Genuine role changes are " +
		"left untouched."
}

// PlanModifySet is the method required by the planmodifier.Set interface.
func (m fabricDevicesSetPlanModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// Nothing to reconcile on create (no prior state) or when the plan is not yet known.
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() || req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	var planned []FabricDevicesFabricDevices
	resp.Diagnostics.Append(req.PlanValue.ElementsAs(ctx, &planned, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state []FabricDevicesFabricDevices
	resp.Diagnostics.Append(req.StateValue.ElementsAs(ctx, &state, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	changed := false
	for i := range planned {
		p := &planned[i]
		if p.NetworkDeviceId.IsNull() || p.NetworkDeviceId.IsUnknown() {
			continue
		}
		if p.DeviceRoles.IsNull() || p.DeviceRoles.IsUnknown() {
			continue
		}

		// Match the planned device to its prior-state counterpart by network_device_id.
		for j := range state {
			s := state[j]
			if !p.NetworkDeviceId.Equal(s.NetworkDeviceId) {
				continue
			}
			if s.DeviceRoles.IsNull() || s.DeviceRoles.IsUnknown() {
				break
			}

			var planRoles, stateRoles []string
			resp.Diagnostics.Append(p.DeviceRoles.ElementsAs(ctx, &planRoles, false)...)
			resp.Diagnostics.Append(s.DeviceRoles.ElementsAs(ctx, &stateRoles, false)...)
			if resp.Diagnostics.HasError() {
				return
			}

			// Copy the state roles into the plan only when the plan is a strict subset of
			// the state and every extra role in the state is a controller-added wireless
			// role. This suppresses the auto-added-role drift while leaving any genuine
			// role change (add/remove of a non-wireless role) intact.
			if rolesDifferOnlyByAutoRoles(planRoles, stateRoles) {
				p.DeviceRoles = s.DeviceRoles
				// The nested `id` is Computed with no UseStateForUnknown, so once this
				// element is seen as changed it becomes unknown in the plan. Left unknown,
				// the reconstructed set element would never match prior state and the diff
				// would persist despite the roles now matching. Re-adopt the state id.
				if p.Id.IsUnknown() || p.Id.IsNull() {
					p.Id = s.Id
				}
				changed = true
			}
			break
		}
	}

	if !changed {
		return
	}

	newSet, diags := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: fabricDevicesAttributeTypes()}, planned)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.PlanValue = newSet
}

// rolesDifferOnlyByAutoRoles reports whether stateRoles equals planRoles plus one or
// more controller-added wireless roles and nothing else. Returns false when the sets are
// identical (nothing to suppress) or when the difference includes any user-managed role.
func rolesDifferOnlyByAutoRoles(planRoles, stateRoles []string) bool {
	planSet := make(map[string]bool, len(planRoles))
	for _, r := range planRoles {
		planSet[r] = true
	}
	stateSet := make(map[string]bool, len(stateRoles))
	for _, r := range stateRoles {
		stateSet[r] = true
	}

	// Every planned role must exist in the state (plan is a subset of state).
	for r := range planSet {
		if !stateSet[r] {
			return false
		}
	}

	// Every extra role present only in the state must be a controller-added wireless role,
	// and there must be at least one such extra (otherwise the sets are identical).
	extras := 0
	for r := range stateSet {
		if planSet[r] {
			continue
		}
		if !wirelessControllerAutoRoles[r] {
			return false
		}
		extras++
	}
	return extras > 0
}

// fabricDevicesAttributeTypes must mirror the FabricDevicesFabricDevices schema.
func fabricDevicesAttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                              types.StringType,
		"network_device_id":               types.StringType,
		"fabric_id":                       types.StringType,
		"device_roles":                    types.SetType{ElemType: types.StringType},
		"border_types":                    types.SetType{ElemType: types.StringType},
		"local_autonomous_system_number":  types.StringType,
		"default_exit":                    types.BoolType,
		"import_external_routes":          types.BoolType,
		"border_priority":                 types.Int64Type,
		"prepend_autonomous_system_count": types.Int64Type,
	}
}
