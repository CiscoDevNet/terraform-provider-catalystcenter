// multicast_rps_set_plan_modifier.go
package provider

import (
	"context"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type multicastRpsSetPlanModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m multicastRpsSetPlanModifier) Description(ctx context.Context) string {
	return "Retains state ipv4_address/ipv6_address for FABRIC RPs if omitted in plan, matching by rp_device_location and network_device_ids."
}

// MarkdownDescription returns a markdown-formatted description of the plan modifier.
func (m multicastRpsSetPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "For FABRIC RPs where both `ipv4_address` and `ipv6_address` are omitted in the plan, " +
		"this modifier searches the prior state for a matching RP (by `rp_device_location` and `network_device_ids`). " +
		"If a match is found, the `ipv4_address` and/or `ipv6_address` from the state will be copied to the plan."
}

// PlanModifySet is the method required by the planmodifier.Set interface.
func (m multicastRpsSetPlanModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	var plannedRps []FabricMulticastVirtualNetworksVirtualNetworksMulticastRps
	resp.Diagnostics.Append(req.PlanValue.ElementsAs(ctx, &plannedRps, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateRps []FabricMulticastVirtualNetworksVirtualNetworksMulticastRps
	resp.Diagnostics.Append(req.StateValue.ElementsAs(ctx, &stateRps, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i := range plannedRps {
		currentPlannedRp := &plannedRps[i]

		if (currentPlannedRp.Ipv4Address.IsNull() || currentPlannedRp.Ipv4Address.IsUnknown()) &&
			(currentPlannedRp.Ipv6Address.IsNull() || currentPlannedRp.Ipv6Address.IsUnknown()) {

			if currentPlannedRp.RpDeviceLocation.ValueString() == "FABRIC" {
				for _, stateRp := range stateRps {
					if currentPlannedRp.RpDeviceLocation.Equal(stateRp.RpDeviceLocation) {
						var plannedNetDevIds []string
						var stateNetDevIds []string

						resp.Diagnostics.Append(currentPlannedRp.NetworkDeviceIds.ElementsAs(ctx, &plannedNetDevIds, false)...)
						resp.Diagnostics.Append(stateRp.NetworkDeviceIds.ElementsAs(ctx, &stateNetDevIds, false)...)
						if resp.Diagnostics.HasError() {
							return
						}

						sort.Strings(plannedNetDevIds)
						sort.Strings(stateNetDevIds)

						if len(plannedNetDevIds) == len(stateNetDevIds) &&
							strings.Join(plannedNetDevIds, ",") == strings.Join(stateNetDevIds, ",") {

							if !stateRp.Ipv4Address.IsNull() && !stateRp.Ipv4Address.IsUnknown() {
								currentPlannedRp.Ipv4Address = stateRp.Ipv4Address
							}
							if !stateRp.Ipv6Address.IsNull() && !stateRp.Ipv6Address.IsUnknown() {
								currentPlannedRp.Ipv6Address = stateRp.Ipv6Address
							}
							break
						}
					}
				}
				if currentPlannedRp.Ipv4Address.IsNull() || currentPlannedRp.Ipv4Address.IsUnknown() || currentPlannedRp.Ipv4Address.ValueString() == "" {
					currentPlannedRp.Ipv4Address = types.StringUnknown()
				}
				if currentPlannedRp.Ipv6Address.IsNull() {
					currentPlannedRp.Ipv6Address = types.StringUnknown()
				}
			}
		}
	}

	var newPlannedSet types.Set
	var diags diag.Diagnostics
	newPlannedSet, diags = types.SetValueFrom(ctx, types.ObjectType{AttrTypes: multicastRpsAttributeTypes()}, plannedRps)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.PlanValue = newPlannedSet
}

// It must accurately reflect the schema of the FabricMulticastVirtualNetworksVirtualNetworksMulticastRps struct.
func multicastRpsAttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"rp_device_location": types.StringType,
		"ipv4_address":       types.StringType,
		"ipv6_address":       types.StringType,
		"is_default_v4_rp":   types.BoolType,
		"is_default_v6_rp":   types.BoolType,
		"network_device_ids": types.SetType{ElemType: types.StringType},
		"ipv4_asm_ranges":    types.SetType{ElemType: types.StringType},
		"ipv6_asm_ranges":    types.SetType{ElemType: types.StringType},
	}
}
