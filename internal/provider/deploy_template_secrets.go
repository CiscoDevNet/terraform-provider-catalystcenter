package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithValidateConfig = &DeployTemplateResource{}

// secretTargetKey uses separate fields so parameter/hostname delimiters cannot collide.
type secretTargetKey struct{ member, id, host string }

func secretKey(entry DeployTemplateSecretParams) secretTargetKey {
	return secretTargetKey{entry.MemberTemplateId.ValueString(), entry.TargetId.ValueString(), entry.TargetHostName.ValueString()}
}

func (r *DeployTemplateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var entries types.List
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("secret_params"), &entries)...)
	if resp.Diagnostics.HasError() || entries.IsNull() || entries.IsUnknown() {
		return
	}
	var config DeployTemplate
	config.SecretParams = make([]DeployTemplateSecretParams, len(entries.Elements()))
	for i, entry := range entries.Elements() {
		if entry.IsUnknown() {
			// Preserve positions and validate known siblings without trying to
			// decode an unknown object into a Go struct.
			config.SecretParams[i] = DeployTemplateSecretParams{
				TargetId: types.StringUnknown(), TargetHostName: types.StringUnknown(), MemberTemplateId: types.StringUnknown(),
				ParamsWo: types.MapUnknown(types.ListType{ElemType: types.StringType}), ParamsWoVersions: types.MapUnknown(types.Int64Type),
			}
			continue
		}
		resp.Diagnostics.Append(tfsdk.ValueAs(ctx, entry, &config.SecretParams[i])...)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	var targets types.Set
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("target_info"), &targets)...)
	// The target model contains Go slices/structs that cannot represent unknown
	// nested collections or objects. Defer matching until the entire target set
	// is known, while still validating any known secret entries independently.
	terraformTargets, err := targets.ToTerraformValue(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Deployment Targets", "Unable to decode deployment target configuration.")
		return
	}
	targetsKnown := terraformTargets.IsFullyKnown()
	if !targets.IsNull() && targetsKnown {
		resp.Diagnostics.Append(targets.ElementsAs(ctx, &config.TargetInfo, false)...)
	}
	var members types.List
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("member_template_deployment_info"), &members)...)
	// A member object may contain a target set that is unknown until dependencies resolve.
	membersKnown := !members.IsUnknown()
	if !members.IsNull() && membersKnown {
		diags := members.ElementsAs(ctx, &config.MemberTemplateDeploymentInfo, false)
		if diags.HasError() {
			membersKnown = false
		} else {
			resp.Diagnostics.Append(diags...)
		}
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(validateDeploymentSecrets(ctx, config, targetsKnown && membersKnown, false)...)
}

func validateDeploymentSecrets(ctx context.Context, data DeployTemplate, checkTargets, applying bool) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	seen := make(map[secretTargetKey]bool)
	for i, entry := range data.SecretParams {
		fail := func(message string) {
			diagnostics.AddError("Invalid Secret Parameters", fmt.Sprintf("secret_params entry %d: %s", i, message))
		}
		unknownSelector := entry.TargetId.IsUnknown() || entry.TargetHostName.IsUnknown() || entry.MemberTemplateId.IsUnknown()
		if unknownSelector {
			if applying {
				fail("target references must be known at apply.")
			}
		} else {
			id, host := entry.TargetId.ValueString(), entry.TargetHostName.ValueString()
			if (id == "") == (host == "") {
				fail("exactly one nonempty target_id or target_host_name must be set.")
			}
			key := secretKey(entry)
			if seen[key] {
				fail("duplicate target/member reference; combine secrets into one entry.")
			}
			for _, previous := range data.SecretParams[:i] {
				if secretKey(previous) != key && sameSecretTarget(data, entry, previous) {
					fail("duplicate target/member reference through ID and hostname; combine secrets into one entry.")
				}
			}
			seen[key] = true
			if checkTargets && (id != "" || host != "") {
				public, matches, unresolved := publicParamsForSecret(data, entry)
				if !unresolved && matches != 1 {
					fail("reference must match exactly one existing deployment target and, when specified, its member_template_id.")
				}
				for name := range entry.ParamsWoVersions.Elements() {
					if _, exists := public.Elements()[name]; exists {
						fail("a secret parameter name is also present in the target's public params.")
					}
				}
			}
		}
		valuesKnown, versionsKnown := !entry.ParamsWo.IsUnknown(), !entry.ParamsWoVersions.IsUnknown()
		if !valuesKnown || !versionsKnown {
			if applying {
				fail("secret values and versions must be known at apply.")
			}
		}
		values, versions := entry.ParamsWo.Elements(), entry.ParamsWoVersions.Elements()
		if (valuesKnown && len(values) == 0) || (versionsKnown && len(versions) == 0) {
			fail("params_wo and params_wo_versions must contain matching nonempty maps.")
		}
		for name, value := range values {
			if name == "" {
				fail("parameter names must be nonempty.")
			}
			if _, exists := versions[name]; versionsKnown && !exists {
				fail("each secret requires a matching params_wo_versions key.")
			}
			list := value.(types.List)
			if list.IsUnknown() {
				if applying {
					fail("secret lists must be known at apply.")
				}
				continue
			}
			if list.IsNull() || len(list.Elements()) == 0 {
				fail("secret values must be nonempty lists of strings.")
			}
			for _, raw := range list.Elements() {
				element := raw.(types.String)
				if element.IsNull() || (!element.IsUnknown() && element.ValueString() == "") {
					fail("secret values must be nonempty strings.")
				}
				if applying && element.IsUnknown() {
					fail("secret strings must be known at apply.")
				}
			}
		}
		for name, raw := range versions {
			if name == "" {
				fail("parameter names must be nonempty.")
			}
			if _, exists := values[name]; valuesKnown && !exists {
				fail("each version requires a matching params_wo key.")
			}
			version := raw.(types.Int64)
			if version.IsNull() || (!version.IsUnknown() && version.ValueInt64() < 1) {
				fail("rotation versions must be positive integers.")
			}
			if applying && version.IsUnknown() {
				fail("rotation versions must be known at apply.")
			}
		}
	}
	return diagnostics
}

func sameSecretTarget(data DeployTemplate, a, b DeployTemplateSecretParams) bool {
	member := a.MemberTemplateId.ValueString()
	if member != b.MemberTemplateId.ValueString() {
		return false
	}
	if member == "" {
		for _, target := range data.TargetInfo {
			if secretMatches(a, member, target.Id, target.HostName) && secretMatches(b, member, target.Id, target.HostName) {
				return true
			}
		}
	} else {
		for _, m := range data.MemberTemplateDeploymentInfo {
			if m.TemplateId.ValueString() != member {
				continue
			}
			for _, target := range m.TargetInfo {
				if secretMatches(a, member, target.Id, target.HostName) && secretMatches(b, member, target.Id, target.HostName) {
					return true
				}
			}
		}
	}
	return false
}

func secretMatches(entry DeployTemplateSecretParams, member string, id, host types.String) bool {
	if entry.MemberTemplateId.ValueString() != member {
		return false
	}
	if entry.TargetId.ValueString() != "" {
		return entry.TargetId.Equal(id)
	}
	return entry.TargetHostName.ValueString() != "" && entry.TargetHostName.Equal(host)
}

func publicParamsForSecret(data DeployTemplate, entry DeployTemplateSecretParams) (types.Map, int, bool) {
	result := types.MapNull(types.ListType{ElemType: types.StringType})
	matches, unresolved := 0, false
	member := entry.MemberTemplateId.ValueString()
	if member == "" {
		for _, target := range data.TargetInfo {
			unresolved = unresolved || target.Id.IsUnknown() || target.HostName.IsUnknown() || target.Params.IsUnknown()
			if secretMatches(entry, "", target.Id, target.HostName) {
				result = target.Params
				matches++
			}
		}
	} else {
		for _, m := range data.MemberTemplateDeploymentInfo {
			unresolved = unresolved || m.TemplateId.IsUnknown()
			if m.TemplateId.ValueString() != member {
				continue
			}
			for _, target := range m.TargetInfo {
				unresolved = unresolved || target.Id.IsUnknown() || target.HostName.IsUnknown() || target.Params.IsUnknown()
				if secretMatches(entry, member, target.Id, target.HostName) {
					result = target.Params
					matches++
				}
			}
		}
	}
	return result, matches, unresolved
}

func secretVersionsChanged(plan, state DeployTemplate, member string, id, host types.String) bool {
	var planned, previous *DeployTemplateSecretParams
	for i := range plan.SecretParams {
		if secretMatches(plan.SecretParams[i], member, id, host) {
			planned = &plan.SecretParams[i]
		}
	}
	for i := range state.SecretParams {
		if secretMatches(state.SecretParams[i], member, id, host) {
			previous = &state.SecretParams[i]
		}
	}
	if planned == nil || previous == nil {
		return planned != nil || previous != nil
	}
	return !planned.ParamsWoVersions.Equal(previous.ParamsWoVersions)
}

func deploymentPolicy(target, overall types.String) string {
	if !target.IsNull() {
		return target.ValueString()
	}
	if !overall.IsNull() {
		return overall.ValueString()
	}
	return "NEVER"
}

// Select using target identities, including composite members, instead of set positions.
func deploymentTargets(plan, state DeployTemplate) []DeployTemplateTargetInfo {
	var selected []DeployTemplateTargetInfo
	for _, target := range plan.TargetInfo {
		found, changed := false, false
		for _, old := range state.TargetInfo {
			if targetsMatch(target, old) {
				found = true
				changed = targetChanged(target, old)
				break
			}
		}
		changed = changed || secretVersionsChanged(plan, state, "", target.Id, target.HostName)
		policy := deploymentPolicy(target.Redeploy, plan.Redeploy)
		deploy := !found || policy == "ALWAYS" || (changed && policy == "ON_CHANGE")
		for _, member := range plan.MemberTemplateDeploymentInfo {
			for _, mt := range member.TargetInfo {
				if !targetsMatch(target, DeployTemplateTargetInfo{Id: mt.Id, HostName: mt.HostName}) {
					continue
				}
				memberFound, memberChanged := false, false
				for _, oldMember := range state.MemberTemplateDeploymentInfo {
					if !member.TemplateId.Equal(oldMember.TemplateId) {
						continue
					}
					for _, old := range oldMember.TargetInfo {
						if targetsMatch(DeployTemplateTargetInfo{Id: mt.Id, HostName: mt.HostName}, DeployTemplateTargetInfo{Id: old.Id, HostName: old.HostName}) {
							memberFound = true
							memberChanged = memberTargetChanged(mt, old) || !mt.Redeploy.Equal(old.Redeploy)
							break
						}
					}
				}
				memberChanged = memberChanged || secretVersionsChanged(plan, state, member.TemplateId.ValueString(), mt.Id, mt.HostName)
				memberPolicy := deploymentPolicy(mt.Redeploy, plan.Redeploy)
				deploy = deploy || !memberFound || memberPolicy == "ALWAYS" || (memberChanged && memberPolicy == "ON_CHANGE")
			}
		}
		if deploy {
			selected = append(selected, target)
		}
	}
	return selected
}

// deploymentBody merges config-only secrets into decoded JSON. It never mutates a
// plan/state target or its public map, even on timeout or failure.
func deploymentBody(ctx context.Context, data DeployTemplate, entries []DeployTemplateSecretParams, diagnostics *diag.Diagnostics) string {
	body := data.toBody(ctx, DeployTemplate{})
	if len(entries) == 0 {
		return body
	}
	var payload map[string]any
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		diagnostics.AddError("Internal Error", "Unable to construct template deployment payload.")
		return ""
	}
	merge := func(raw any, member string, id, host types.String) {
		target := raw.(map[string]any)
		for _, entry := range entries {
			if !secretMatches(entry, member, id, host) {
				continue
			}
			params, ok := target["params"].(map[string]any)
			if !ok {
				params = make(map[string]any)
			}
			for name, value := range entry.ParamsWo.Elements() {
				var strings []string
				diagnostics.Append(value.(types.List).ElementsAs(ctx, &strings, false)...)
				if len(strings) == 1 {
					params[name] = strings[0]
				} else {
					params[name] = strings
				}
			}
			target["params"] = params
		}
	}
	if raw, ok := payload["targetInfo"].([]any); ok {
		for i, target := range data.TargetInfo {
			merge(raw[i], "", target.Id, target.HostName)
		}
	}
	if raw, ok := payload["memberTemplateDeploymentInfo"].([]any); ok {
		for i, member := range data.MemberTemplateDeploymentInfo {
			if len(member.TargetInfo) == 0 {
				continue
			}
			targets := raw[i].(map[string]any)["targetInfo"].([]any)
			for j, target := range member.TargetInfo {
				merge(targets[j], member.TemplateId.ValueString(), target.Id, target.HostName)
			}
		}
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		diagnostics.AddError("Internal Error", "Unable to encode template deployment payload.")
		return ""
	}
	return string(bytes)
}
