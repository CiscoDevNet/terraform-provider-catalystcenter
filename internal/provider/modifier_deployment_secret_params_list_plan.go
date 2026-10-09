package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ planmodifier.List = deploymentSecretParamsListPlanModifier{}

type deploymentSecretParamsListPlanModifier struct{}

func (m deploymentSecretParamsListPlanModifier) Description(ctx context.Context) string {
	return "Keeps write-only parameter values null in planned secret entries while preserving unknown target references and versions."
}

func (m deploymentSecretParamsListPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// Unknown objects under a known list still have addressable write-only paths.
// Expand their metadata as unknown values and explicitly null the write-only map.
func (m deploymentSecretParamsListPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	entries := req.PlanValue
	if entries.IsNull() || entries.IsUnknown() {
		return
	}
	elementType := entries.ElementType(ctx).(types.ObjectType)
	elements, changed := entries.Elements(), false
	for i, raw := range elements {
		if raw.IsNull() {
			continue
		}
		entry := raw.(types.Object)
		attributes := entry.Attributes()
		if entry.IsUnknown() {
			attributes = make(map[string]attr.Value)
			for name, typ := range elementType.AttrTypes {
				value, err := typ.ValueFromTerraform(ctx, tftypes.NewValue(typ.TerraformType(ctx), tftypes.UnknownValue))
				if err != nil {
					resp.Diagnostics.AddError("Secret Parameter Plan Error", "Unable to preserve unknown secret metadata in the plan.")
					return
				}
				attributes[name] = value
			}
		} else if attributes["params_wo"].IsNull() {
			continue
		}
		woType := elementType.AttrTypes["params_wo"]
		value, err := woType.ValueFromTerraform(ctx, tftypes.NewValue(woType.TerraformType(ctx), nil))
		if err != nil {
			resp.Diagnostics.AddError("Secret Parameter Plan Error", "Unable to remove write-only values from the plan.")
			return
		}
		attributes["params_wo"] = value
		object, diagnostics := types.ObjectValue(elementType.AttrTypes, attributes)
		resp.Diagnostics.Append(diagnostics...)
		if resp.Diagnostics.HasError() {
			return
		}
		elements[i], changed = object, true
	}
	if changed {
		planned, diagnostics := types.ListValue(elementType, elements)
		resp.Diagnostics.Append(diagnostics...)
		if !resp.Diagnostics.HasError() {
			resp.PlanValue = planned
		}
	}
}
