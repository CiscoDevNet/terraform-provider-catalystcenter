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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessProfileSiteTagSiteAssignment struct {
	Id                types.String `tfsdk:"id"`
	WirelessProfileId types.String `tfsdk:"wireless_profile_id"`
	SiteTagName       types.String `tfsdk:"site_tag_name"`
	SiteId            types.String `tfsdk:"site_id"`
	ApProfileName     types.String `tfsdk:"ap_profile_name"`
	FlexProfileName   types.String `tfsdk:"flex_profile_name"`
	SiteTagId         types.String `tfsdk:"site_tag_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessProfileSiteTagSiteAssignment) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags", url.QueryEscape(data.WirelessProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data WirelessProfileSiteTagSiteAssignment) toBody(ctx context.Context, state WirelessProfileSiteTagSiteAssignment) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SiteTagName.IsNull() {
		body, _ = sjson.Set(body, "siteTagName", data.SiteTagName.ValueString())
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if !data.ApProfileName.IsNull() {
		body, _ = sjson.Set(body, "apProfileName", data.ApProfileName.ValueString())
	}
	if !data.FlexProfileName.IsNull() {
		body, _ = sjson.Set(body, "flexProfileName", data.FlexProfileName.ValueString())
	}
	if data.SiteTagId.ValueString() != "" && !data.SiteTagId.IsNull() {
		body, _ = sjson.Set(body, "siteTagId", data.SiteTagId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessProfileSiteTagSiteAssignment) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteTagName"); value.Exists() {
		data.SiteTagName = types.StringValue(value.String())
	} else {
		data.SiteTagName = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("apProfileName"); value.Exists() {
		data.ApProfileName = types.StringValue(value.String())
	} else {
		data.ApProfileName = types.StringNull()
	}
	if value := res.Get("flexProfileName"); value.Exists() {
		data.FlexProfileName = types.StringValue(value.String())
	} else {
		data.FlexProfileName = types.StringValue("default-flex-profile")
	}
	if value := res.Get("siteTagId"); value.Exists() {
		data.SiteTagId = types.StringValue(value.String())
	} else {
		data.SiteTagId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessProfileSiteTagSiteAssignment) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteTagName"); value.Exists() && !data.SiteTagName.IsNull() {
		data.SiteTagName = types.StringValue(value.String())
	} else {
		data.SiteTagName = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("apProfileName"); value.Exists() && !data.ApProfileName.IsNull() {
		data.ApProfileName = types.StringValue(value.String())
	} else {
		data.ApProfileName = types.StringNull()
	}
	if value := res.Get("flexProfileName"); value.Exists() && !data.FlexProfileName.IsNull() {
		data.FlexProfileName = types.StringValue(value.String())
	} else if data.FlexProfileName.ValueString() != "default-flex-profile" {
		data.FlexProfileName = types.StringNull()
	}
	if value := res.Get("siteTagId"); value.Exists() && !data.SiteTagId.IsNull() {
		data.SiteTagId = types.StringValue(value.String())
	} else {
		data.SiteTagId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessProfileSiteTagSiteAssignment) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.SiteTagName.IsNull() {
		return false
	}
	if !data.SiteId.IsNull() {
		return false
	}
	if !data.ApProfileName.IsNull() {
		return false
	}
	if !data.FlexProfileName.IsNull() {
		return false
	}
	if !data.SiteTagId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull

// Custom helpers (not generated) supporting the read-modify-write CRUD in the resource file.

// getAssignmentId builds the stable composite id for this membership edge. Site Tag names cannot
// contain '/', so it is a safe separator.
func (data WirelessProfileSiteTagSiteAssignment) getAssignmentId() string {
	return fmt.Sprintf("%s/%s/%s", data.WirelessProfileId.ValueString(), data.SiteTagName.ValueString(), data.SiteId.ValueString())
}

// getPathListByName lists the Site Tags of the Wireless Profile filtered by name (server-side).
func (data WirelessProfileSiteTagSiteAssignment) getPathListByName() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags?siteTagName=%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(data.SiteTagName.ValueString()))
}

// getPathBulk is the POST endpoint used to first-create the Site Tag.
func (data WirelessProfileSiteTagSiteAssignment) getPathBulk() string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags/bulk", url.QueryEscape(data.WirelessProfileId.ValueString()))
}

// getPathTag is the PUT/DELETE endpoint for a specific Site Tag id.
func (data WirelessProfileSiteTagSiteAssignment) getPathTag(siteTagId string) string {
	return fmt.Sprintf("/dna/intent/api/v1/wirelessProfiles/%v/siteTags/%v", url.QueryEscape(data.WirelessProfileId.ValueString()), url.QueryEscape(siteTagId))
}

// createTagBody builds the POST .../siteTags/bulk body that first-creates the Site Tag with the
// given site ids.
func (data WirelessProfileSiteTagSiteAssignment) createTagBody(siteIds []string) string {
	body := ""
	body, _ = sjson.Set(body, "items.0.siteTagName", data.SiteTagName.ValueString())
	body, _ = sjson.Set(body, "items.0.apProfileName", data.ApProfileName.ValueString())
	body, _ = sjson.Set(body, "items.0.siteIds", siteIds)
	if !data.FlexProfileName.IsNull() && data.FlexProfileName.ValueString() != "" {
		body, _ = sjson.Set(body, "items.0.flexProfileName", data.FlexProfileName.ValueString())
	}
	return body
}

// putTagBody builds the PUT .../siteTags/{id} body that rewrites the whole membership list while
// preserving the tag's existing profile attributes.
func putTagBody(siteTagName, apProfileName, flexProfileName string, siteIds []string) string {
	body := ""
	body, _ = sjson.Set(body, "siteTagName", siteTagName)
	body, _ = sjson.Set(body, "apProfileName", apProfileName)
	body, _ = sjson.Set(body, "siteIds", siteIds)
	if flexProfileName != "" {
		body, _ = sjson.Set(body, "flexProfileName", flexProfileName)
	}
	return body
}
