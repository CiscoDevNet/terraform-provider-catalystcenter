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
type WirelessProfileSiteTag struct {
	Id                types.String `tfsdk:"id"`
	WirelessProfileId types.String `tfsdk:"wireless_profile_id"`
	SiteTagId         types.String `tfsdk:"site_tag_id"`
	SiteTagName       types.String `tfsdk:"site_tag_name"`
	ApProfileName     types.String `tfsdk:"ap_profile_name"`
	SiteIds           types.Set    `tfsdk:"site_ids"`
	FlexProfileName   types.String `tfsdk:"flex_profile_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessProfileSiteTag) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags/bulk", url.QueryEscape(data.WirelessProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data WirelessProfileSiteTag) getPathDelete() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.SiteTagId.ValueString()))
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

func (data WirelessProfileSiteTag) getPathGet() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.SiteTagId.ValueString()))
}

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

func (data WirelessProfileSiteTag) getPathPut() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.SiteTagId.ValueString()))
}

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

func (data WirelessProfileSiteTag) getPathIdQuery() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags", url.QueryEscape(data.WirelessProfileId.ValueString()))
}

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessProfileSiteTag) toBody(ctx context.Context, state WirelessProfileSiteTag) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SiteTagName.IsNull() {
		if put {
			body, _ = sjson.Set(body, "siteTagName", data.SiteTagName.ValueString())
		} else {
			body, _ = sjson.Set(body, "items.0.siteTagName", data.SiteTagName.ValueString())
		}
	}
	if !data.ApProfileName.IsNull() {
		if put {
			body, _ = sjson.Set(body, "apProfileName", data.ApProfileName.ValueString())
		} else {
			body, _ = sjson.Set(body, "items.0.apProfileName", data.ApProfileName.ValueString())
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
	if !data.FlexProfileName.IsNull() {
		if put {
			body, _ = sjson.Set(body, "flexProfileName", data.FlexProfileName.ValueString())
		} else {
			body, _ = sjson.Set(body, "items.0.flexProfileName", data.FlexProfileName.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessProfileSiteTag) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.siteTagName"); value.Exists() {
		data.SiteTagName = types.StringValue(value.String())
	} else {
		data.SiteTagName = types.StringNull()
	}
	if value := res.Get("response.apProfileName"); value.Exists() {
		data.ApProfileName = types.StringValue(value.String())
	} else {
		data.ApProfileName = types.StringNull()
	}
	if value := res.Get("response.siteIds"); value.Exists() && len(value.Array()) > 0 {
		data.SiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.SiteIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.flexProfileName"); value.Exists() {
		data.FlexProfileName = types.StringValue(value.String())
	} else {
		data.FlexProfileName = types.StringValue("default-flex-profile")
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessProfileSiteTag) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.siteTagName"); value.Exists() && !data.SiteTagName.IsNull() {
		data.SiteTagName = types.StringValue(value.String())
	} else {
		data.SiteTagName = types.StringNull()
	}
	if value := res.Get("response.apProfileName"); value.Exists() && !data.ApProfileName.IsNull() {
		data.ApProfileName = types.StringValue(value.String())
	} else {
		data.ApProfileName = types.StringNull()
	}
	if value := res.Get("response.siteIds"); value.Exists() && !data.SiteIds.IsNull() {
		data.SiteIds = helpers.GetStringSet(value.Array())
	} else {
		data.SiteIds = types.SetNull(types.StringType)
	}
	if value := res.Get("response.flexProfileName"); value.Exists() && !data.FlexProfileName.IsNull() {
		data.FlexProfileName = types.StringValue(value.String())
	} else if data.FlexProfileName.ValueString() != "default-flex-profile" {
		data.FlexProfileName = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessProfileSiteTag) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.SiteTagName.IsNull() {
		return false
	}
	if !data.ApProfileName.IsNull() {
		return false
	}
	if !data.SiteIds.IsNull() {
		return false
	}
	if !data.FlexProfileName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
