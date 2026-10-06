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
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessProfilePolicyTag struct {
	Id                types.String `tfsdk:"id"`
	WirelessProfileId types.String `tfsdk:"wireless_profile_id"`
	PolicyTagId       types.String `tfsdk:"policy_tag_id"`
	PolicyTagName     types.String `tfsdk:"policy_tag_name"`
	SiteIds           types.Set    `tfsdk:"site_ids"`
	ApZones           types.Set    `tfsdk:"ap_zones"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessProfilePolicyTag) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/policyTags/bulk", url.QueryEscape(data.WirelessProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data WirelessProfilePolicyTag) getPathDelete() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/policyTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.PolicyTagId.ValueString()))
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

func (data WirelessProfilePolicyTag) getPathGet() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/policyTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.PolicyTagId.ValueString()))
}

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

func (data WirelessProfilePolicyTag) getPathPut() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/policyTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.PolicyTagId.ValueString()))
}

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

func (data WirelessProfilePolicyTag) getPathIdQuery() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/policyTags", url.QueryEscape(data.WirelessProfileId.ValueString()))
}

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessProfilePolicyTag) toBody(ctx context.Context, state WirelessProfilePolicyTag) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.PolicyTagName.IsNull() {
		if put {
			body, _ = sjson.Set(body, "policyTagName", data.PolicyTagName.ValueString())
		} else {
			body, _ = sjson.Set(body, "items.0.policyTagName", data.PolicyTagName.ValueString())
		}
	}
	if !data.SiteIds.IsNull() {
		var values []string
		data.SiteIds.ElementsAs(ctx, &values, false)
		if put {
			body, _ = sjson.Set(body, "siteIds", values)
		} else {
			body, _ = sjson.Set(body, "items.0.siteIds", values)
		}
	}
	if !data.ApZones.IsNull() {
		var values []string
		data.ApZones.ElementsAs(ctx, &values, false)
		if put {
			body, _ = sjson.Set(body, "apZones", values)
		} else {
			body, _ = sjson.Set(body, "items.0.apZones", values)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessProfilePolicyTag) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.policyTagName"); value.Exists() {
		data.PolicyTagName = types.StringValue(value.String())
	} else {
		data.PolicyTagName = types.StringNull()
	}
	if value := res.Get("response.siteIds"); value.Exists() && len(value.Array()) > 0 {
		data.SiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.SiteIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.apZones"); value.Exists() && len(value.Array()) > 0 {
		data.ApZones = helpers.GetStringSet(value.Array())
	} else {
		data.ApZones = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessProfilePolicyTag) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.policyTagName"); value.Exists() && !data.PolicyTagName.IsNull() {
		data.PolicyTagName = types.StringValue(value.String())
	} else {
		data.PolicyTagName = types.StringNull()
	}
	if value := res.Get("response.siteIds"); value.Exists() && !data.SiteIds.IsNull() {
		data.SiteIds = helpers.KeepStringSetIn(ctx, data.SiteIds, value.Array())
	} else {
		data.SiteIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.apZones"); value.Exists() && !data.ApZones.IsNull() {
		data.ApZones = helpers.GetStringSet(value.Array())
	} else {
		data.ApZones = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessProfilePolicyTag) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.PolicyTagName.IsNull() {
		return false
	}
	if !data.SiteIds.IsNull() {
		return false
	}
	if !data.ApZones.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
