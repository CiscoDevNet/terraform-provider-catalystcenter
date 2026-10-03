// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ApplicationPolicy struct {
	Id          types.String             `tfsdk:"id"`
	PolicyScope types.String             `tfsdk:"policy_scope"`
	Items       []ApplicationPolicyItems `tfsdk:"items"`
}

type ApplicationPolicyItems struct {
	Name                    types.String `tfsdk:"name"`
	PolicyScope             types.String `tfsdk:"policy_scope"`
	Priority                types.String `tfsdk:"priority"`
	DeletePolicyStatus      types.String `tfsdk:"delete_policy_status"`
	AdvancedPolicyScopeName types.String `tfsdk:"advanced_policy_scope_name"`
	SiteIds                 types.Set    `tfsdk:"site_ids"`
	Ssids                   types.Set    `tfsdk:"ssids"`
	ClauseType              types.String `tfsdk:"clause_type"`
	RelevanceLevel          types.String `tfsdk:"relevance_level"`
	DeviceRemovalBehavior   types.String `tfsdk:"device_removal_behavior"`
	HostTrackingEnabled     types.Bool   `tfsdk:"host_tracking_enabled"`
	QueuingProfileId        types.String `tfsdk:"queuing_profile_id"`
	ApplicationSetId        types.String `tfsdk:"application_set_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ApplicationPolicy) getPath() string {
	return "/dna/intent/api/v1/app-policy-intent"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

func (data ApplicationPolicy) getPathGet() string {
	return "/dna/intent/api/v1/app-policy"
}

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ApplicationPolicy) toBody(ctx context.Context, state ApplicationPolicy) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.PolicyScope.IsNull() {
		body, _ = sjson.Set(body, "policyScope", data.PolicyScope.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "createList", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.PolicyScope.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policyScope", item.PolicyScope.ValueString())
			}
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueString())
			}
			if !item.DeletePolicyStatus.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deletePolicyStatus", item.DeletePolicyStatus.ValueString())
			}
			if !item.AdvancedPolicyScopeName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "advancedPolicyScope.name", item.AdvancedPolicyScopeName.ValueString())
			}
			if !item.SiteIds.IsNull() {
				var values []string
				item.SiteIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "advancedPolicyScope.advancedPolicyScopeElement.0.groupId", values)
			}
			if !item.Ssids.IsNull() {
				var values []string
				item.Ssids.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "advancedPolicyScope.advancedPolicyScopeElement.0.ssid", values)
			}
			if !item.ClauseType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exclusiveContract.clause.0.type", item.ClauseType.ValueString())
			}
			if !item.RelevanceLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exclusiveContract.clause.0.relevanceLevel", item.RelevanceLevel.ValueString())
			}
			if !item.DeviceRemovalBehavior.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exclusiveContract.clause.0.deviceRemovalBehavior", item.DeviceRemovalBehavior.ValueString())
			}
			if !item.HostTrackingEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exclusiveContract.clause.0.hostTrackingEnabled", item.HostTrackingEnabled.ValueBool())
			}
			if !item.QueuingProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "contract.idRef", item.QueuingProfileId.ValueString())
			}
			if !item.ApplicationSetId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "producer.scalableGroup.0.idRef", item.ApplicationSetId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "createList.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ApplicationPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Items = make([]ApplicationPolicyItems, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationPolicyItems{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("policyScope"); cValue.Exists() {
				item.PolicyScope = types.StringValue(cValue.String())
			} else {
				item.PolicyScope = types.StringNull()
			}
			if cValue := v.Get("priority"); cValue.Exists() {
				item.Priority = types.StringValue(cValue.String())
			} else {
				item.Priority = types.StringNull()
			}
			if cValue := v.Get("deletePolicyStatus"); cValue.Exists() {
				item.DeletePolicyStatus = types.StringValue(cValue.String())
			} else {
				item.DeletePolicyStatus = types.StringNull()
			}
			if cValue := v.Get("advancedPolicyScope.name"); cValue.Exists() {
				item.AdvancedPolicyScopeName = types.StringValue(cValue.String())
			} else {
				item.AdvancedPolicyScopeName = types.StringNull()
			}
			if cValue := v.Get("advancedPolicyScope.advancedPolicyScopeElement.0.groupId"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.SiteIds = helpers.GetStringSet(cValue.Array())
			} else {
				item.SiteIds = types.SetNull(types.StringType)
			}
			if cValue := v.Get("advancedPolicyScope.advancedPolicyScopeElement.0.ssid"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ssids = helpers.GetStringSet(cValue.Array())
			} else {
				item.Ssids = types.SetNull(types.StringType)
			}
			if cValue := v.Get("exclusiveContract.clause.0.type"); cValue.Exists() {
				item.ClauseType = types.StringValue(cValue.String())
			} else {
				item.ClauseType = types.StringNull()
			}
			if cValue := v.Get("exclusiveContract.clause.0.relevanceLevel"); cValue.Exists() {
				item.RelevanceLevel = types.StringValue(cValue.String())
			} else {
				item.RelevanceLevel = types.StringNull()
			}
			if cValue := v.Get("exclusiveContract.clause.0.deviceRemovalBehavior"); cValue.Exists() {
				item.DeviceRemovalBehavior = types.StringValue(cValue.String())
			} else {
				item.DeviceRemovalBehavior = types.StringNull()
			}
			if cValue := v.Get("exclusiveContract.clause.0.hostTrackingEnabled"); cValue.Exists() {
				item.HostTrackingEnabled = types.BoolValue(cValue.Bool())
			} else {
				item.HostTrackingEnabled = types.BoolNull()
			}
			if cValue := v.Get("contract.idRef"); cValue.Exists() {
				item.QueuingProfileId = types.StringValue(cValue.String())
			} else {
				item.QueuingProfileId = types.StringNull()
			}
			if cValue := v.Get("producer.scalableGroup.0.idRef"); cValue.Exists() {
				item.ApplicationSetId = types.StringValue(cValue.String())
			} else {
				item.ApplicationSetId = types.StringNull()
			}
			data.Items = append(data.Items, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ApplicationPolicy) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Items[i].Name.ValueString()}

		var r gjson.Result
		res.Get("response").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.Items[i].Name.IsNull() {
			data.Items[i].Name = types.StringValue(value.String())
		} else {
			data.Items[i].Name = types.StringNull()
		}
		if value := r.Get("policyScope"); value.Exists() && !data.Items[i].PolicyScope.IsNull() {
			data.Items[i].PolicyScope = types.StringValue(value.String())
		} else {
			data.Items[i].PolicyScope = types.StringNull()
		}
		if value := r.Get("priority"); value.Exists() && !data.Items[i].Priority.IsNull() {
			data.Items[i].Priority = types.StringValue(value.String())
		} else {
			data.Items[i].Priority = types.StringNull()
		}
		if value := r.Get("deletePolicyStatus"); value.Exists() && !data.Items[i].DeletePolicyStatus.IsNull() {
			data.Items[i].DeletePolicyStatus = types.StringValue(value.String())
		} else {
			data.Items[i].DeletePolicyStatus = types.StringNull()
		}
		if value := r.Get("advancedPolicyScope.name"); value.Exists() && !data.Items[i].AdvancedPolicyScopeName.IsNull() {
			data.Items[i].AdvancedPolicyScopeName = types.StringValue(value.String())
		} else {
			data.Items[i].AdvancedPolicyScopeName = types.StringNull()
		}
		if value := r.Get("advancedPolicyScope.advancedPolicyScopeElement.0.groupId"); value.Exists() && !data.Items[i].SiteIds.IsNull() {
			data.Items[i].SiteIds = helpers.GetStringSet(value.Array())
		} else {
			data.Items[i].SiteIds = types.SetNull(types.StringType)
		}
		if value := r.Get("advancedPolicyScope.advancedPolicyScopeElement.0.ssid"); value.Exists() && !data.Items[i].Ssids.IsNull() {
			data.Items[i].Ssids = helpers.GetStringSet(value.Array())
		} else {
			data.Items[i].Ssids = types.SetNull(types.StringType)
		}
		if value := r.Get("exclusiveContract.clause.0.type"); value.Exists() && !data.Items[i].ClauseType.IsNull() {
			data.Items[i].ClauseType = types.StringValue(value.String())
		} else {
			data.Items[i].ClauseType = types.StringNull()
		}
		if value := r.Get("exclusiveContract.clause.0.relevanceLevel"); value.Exists() && !data.Items[i].RelevanceLevel.IsNull() {
			data.Items[i].RelevanceLevel = types.StringValue(value.String())
		} else {
			data.Items[i].RelevanceLevel = types.StringNull()
		}
		if value := r.Get("exclusiveContract.clause.0.deviceRemovalBehavior"); value.Exists() && !data.Items[i].DeviceRemovalBehavior.IsNull() {
			data.Items[i].DeviceRemovalBehavior = types.StringValue(value.String())
		} else {
			data.Items[i].DeviceRemovalBehavior = types.StringNull()
		}
		if value := r.Get("exclusiveContract.clause.0.hostTrackingEnabled"); value.Exists() && !data.Items[i].HostTrackingEnabled.IsNull() {
			data.Items[i].HostTrackingEnabled = types.BoolValue(value.Bool())
		} else {
			data.Items[i].HostTrackingEnabled = types.BoolNull()
		}
		if value := r.Get("contract.idRef"); value.Exists() && !data.Items[i].QueuingProfileId.IsNull() {
			data.Items[i].QueuingProfileId = types.StringValue(value.String())
		} else {
			data.Items[i].QueuingProfileId = types.StringNull()
		}
		if value := r.Get("producer.scalableGroup.0.idRef"); value.Exists() && !data.Items[i].ApplicationSetId.IsNull() {
			data.Items[i].ApplicationSetId = types.StringValue(value.String())
		} else {
			data.Items[i].ApplicationSetId = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ApplicationPolicy) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Items) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
