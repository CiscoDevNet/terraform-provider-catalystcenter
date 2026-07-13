// modifier_fabric_devices_set_plan.go
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

	// Index prior-state elements by their natural key (network_device_id). A Terraform
	// set can technically hold two objects that share a network_device_id but differ in
	// other fields, which makes correlation ambiguous; skip such keys rather than risk
	// copying the wrong computed values.
	stateByKey := make(map[string]*FabricDevicesFabricDevices, len(state))
	ambiguousKeys := make(map[string]bool)
	for j := range state {
		s := &state[j]
		if s.NetworkDeviceId.IsNull() || s.NetworkDeviceId.IsUnknown() {
			continue
		}
		key := s.NetworkDeviceId.ValueString()
		if _, seen := stateByKey[key]; seen {
			ambiguousKeys[key] = true
			continue
		}
		stateByKey[key] = s
	}

	changed := false
	for i := range planned {
		p := &planned[i]
		if p.NetworkDeviceId.IsNull() || p.NetworkDeviceId.IsUnknown() {
			continue
		}
		key := p.NetworkDeviceId.ValueString()
		if ambiguousKeys[key] {
			continue
		}
		s, ok := stateByKey[key]
		if !ok {
			// Genuinely new element (no prior-state counterpart) - leave its computed id
			// unknown so a real create diff is produced.
			continue
		}

		// Reconcile the controller-added WIRELESS_CONTROLLER_NODE role first, so the role
		// comparison in fabricDeviceConfigEqual treats an auto-added-only difference as a
		// no-op. Copy state roles into the plan only when the plan is a strict subset of
		// the state and every extra role is a controller-added wireless role.
		if !p.DeviceRoles.IsNull() && !p.DeviceRoles.IsUnknown() &&
			!s.DeviceRoles.IsNull() && !s.DeviceRoles.IsUnknown() {
			var planRoles, stateRoles []string
			resp.Diagnostics.Append(p.DeviceRoles.ElementsAs(ctx, &planRoles, false)...)
			resp.Diagnostics.Append(s.DeviceRoles.ElementsAs(ctx, &stateRoles, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
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
		}

		// UseStateForUnknown for the computed-only `id`. fabric_devices is a Set, so element
		// identity is the hash of the whole nested object; a null/unknown `id` in the plan
		// hashes differently from the concrete `id` in state, making an otherwise-identical
		// element show up as remove+add. Propagate the state id only once the element's
		// user-managed fields already match state, so genuine changes still surface a diff.
		if (p.Id.IsNull() || p.Id.IsUnknown()) && !s.Id.IsNull() && !s.Id.IsUnknown() &&
			fabricDeviceConfigEqual(ctx, p, s, &resp.Diagnostics) {
			if resp.Diagnostics.HasError() {
				return
			}
			p.Id = s.Id
			changed = true
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

// fabricDeviceConfigEqual reports whether every user-managed (non-computed) field of the
// planned element already matches state. The computed `id` is excluded because it is the
// value being propagated. device_roles is compared after the auto-role normalization, so an
// auto-added-only difference is already treated as equal.
func fabricDeviceConfigEqual(ctx context.Context, p, s *FabricDevicesFabricDevices, diags *diag.Diagnostics) bool {
	if !p.NetworkDeviceId.Equal(s.NetworkDeviceId) ||
		!p.FabricId.Equal(s.FabricId) ||
		!p.LocalAutonomousSystemNumber.Equal(s.LocalAutonomousSystemNumber) ||
		!p.DefaultExit.Equal(s.DefaultExit) ||
		!p.ImportExternalRoutes.Equal(s.ImportExternalRoutes) ||
		!p.BorderPriority.Equal(s.BorderPriority) ||
		!p.PrependAutonomousSystemCount.Equal(s.PrependAutonomousSystemCount) {
		return false
	}
	return setsEqual(ctx, p.DeviceRoles, s.DeviceRoles, diags) &&
		setsEqual(ctx, p.BorderTypes, s.BorderTypes, diags)
}

// setsEqual compares two string sets by membership, treating null and empty as equal.
func setsEqual(ctx context.Context, a, b types.Set, diags *diag.Diagnostics) bool {
	if a.IsUnknown() || b.IsUnknown() {
		return false
	}
	var av, bv []string
	if !a.IsNull() {
		diags.Append(a.ElementsAs(ctx, &av, false)...)
	}
	if !b.IsNull() {
		diags.Append(b.ElementsAs(ctx, &bv, false)...)
	}
	if len(av) != len(bv) {
		return false
	}
	m := make(map[string]int, len(av))
	for _, r := range av {
		m[r]++
	}
	for _, r := range bv {
		m[r]--
		if m[r] < 0 {
			return false
		}
	}
	return true
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
