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
type ExtranetPolicy struct {
	Id                            types.String `tfsdk:"id"`
	ExtranetPolicyName            types.String `tfsdk:"extranet_policy_name"`
	FabricIds                     types.Set    `tfsdk:"fabric_ids"`
	ProviderVirtualNetworkName    types.String `tfsdk:"provider_virtual_network_name"`
	SubscriberVirtualNetworkNames types.Set    `tfsdk:"subscriber_virtual_network_names"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ExtranetPolicy) getPath() string {
	return "/dna/intent/api/v1/sda/extranetPolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath
func (data ExtranetPolicy) getFallbackPath() string {
	return ""
}

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ExtranetPolicy) toBody(ctx context.Context, state ExtranetPolicy) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.ExtranetPolicyName.IsNull() {
		body, _ = sjson.Set(body, "0.extranetPolicyName", data.ExtranetPolicyName.ValueString())
	}
	if !data.FabricIds.IsNull() {
		var values []string
		data.FabricIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.fabricIds", values)
	}
	if !data.ProviderVirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "0.providerVirtualNetworkName", data.ProviderVirtualNetworkName.ValueString())
	}
	if !data.SubscriberVirtualNetworkNames.IsNull() {
		var values []string
		data.SubscriberVirtualNetworkNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.subscriberVirtualNetworkNames", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ExtranetPolicy) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.extranetPolicyName"); value.Exists() {
		data.ExtranetPolicyName = types.StringValue(value.String())
	} else {
		data.ExtranetPolicyName = types.StringNull()
	}
	if value := res.Get("response.0.fabricIds"); value.Exists() && len(value.Array()) > 0 {
		data.FabricIds = helpers.GetStringSet(value.Array())
	} else {
		data.FabricIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.providerVirtualNetworkName"); value.Exists() {
		data.ProviderVirtualNetworkName = types.StringValue(value.String())
	} else {
		data.ProviderVirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.subscriberVirtualNetworkNames"); value.Exists() && len(value.Array()) > 0 {
		data.SubscriberVirtualNetworkNames = helpers.GetStringSet(value.Array())
	} else {
		data.SubscriberVirtualNetworkNames = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ExtranetPolicy) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.extranetPolicyName"); value.Exists() && !data.ExtranetPolicyName.IsNull() {
		data.ExtranetPolicyName = types.StringValue(value.String())
	} else {
		data.ExtranetPolicyName = types.StringNull()
	}
	if value := res.Get("response.0.fabricIds"); value.Exists() && !data.FabricIds.IsNull() {
		data.FabricIds = helpers.GetStringSet(value.Array())
	} else {
		data.FabricIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.providerVirtualNetworkName"); value.Exists() && !data.ProviderVirtualNetworkName.IsNull() {
		data.ProviderVirtualNetworkName = types.StringValue(value.String())
	} else {
		data.ProviderVirtualNetworkName = types.StringNull()
	}
	if value := res.Get("response.0.subscriberVirtualNetworkNames"); value.Exists() && !data.SubscriberVirtualNetworkNames.IsNull() {
		data.SubscriberVirtualNetworkNames = helpers.GetStringSet(value.Array())
	} else {
		data.SubscriberVirtualNetworkNames = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ExtranetPolicy) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FabricIds.IsNull() {
		return false
	}
	if !data.ProviderVirtualNetworkName.IsNull() {
		return false
	}
	if !data.SubscriberVirtualNetworkNames.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
