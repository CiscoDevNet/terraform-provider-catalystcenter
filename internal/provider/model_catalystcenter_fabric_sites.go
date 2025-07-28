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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type FabricSites struct {
	Id    types.String       `tfsdk:"id"`
	Sites []FabricSitesSites `tfsdk:"sites"`
}

type FabricSitesSites struct {
	Id                        types.String `tfsdk:"id"`
	SiteId                    types.String `tfsdk:"site_id"`
	AuthenticationProfileName types.String `tfsdk:"authentication_profile_name"`
	IsPubSubEnabled           types.Bool   `tfsdk:"is_pub_sub_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricSites) getPath() string {
	return "/dna/intent/api/v1/sda/fabricSites"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricSites) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response"); value.Exists() && len(value.Array()) > 0 {
		data.Sites = make([]FabricSitesSites, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FabricSitesSites{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("siteId"); cValue.Exists() {
				item.SiteId = types.StringValue(cValue.String())
			} else {
				item.SiteId = types.StringNull()
			}
			if cValue := v.Get("authenticationProfileName"); cValue.Exists() {
				item.AuthenticationProfileName = types.StringValue(cValue.String())
			} else {
				item.AuthenticationProfileName = types.StringNull()
			}
			if cValue := v.Get("isPubSubEnabled"); cValue.Exists() {
				item.IsPubSubEnabled = types.BoolValue(cValue.Bool())
			} else {
				item.IsPubSubEnabled = types.BoolNull()
			}
			data.Sites = append(data.Sites, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricSites) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Sites {
		keys := [...]string{"id", "siteId", "authenticationProfileName", "isPubSubEnabled"}
		keyValues := [...]string{data.Sites[i].Id.ValueString(), data.Sites[i].SiteId.ValueString(), data.Sites[i].AuthenticationProfileName.ValueString(), strconv.FormatBool(data.Sites[i].IsPubSubEnabled.ValueBool())}

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
		if value := r.Get("id"); value.Exists() && !data.Sites[i].Id.IsNull() {
			data.Sites[i].Id = types.StringValue(value.String())
		} else {
			data.Sites[i].Id = types.StringNull()
		}
		if value := r.Get("siteId"); value.Exists() && !data.Sites[i].SiteId.IsNull() {
			data.Sites[i].SiteId = types.StringValue(value.String())
		} else {
			data.Sites[i].SiteId = types.StringNull()
		}
		if value := r.Get("authenticationProfileName"); value.Exists() && !data.Sites[i].AuthenticationProfileName.IsNull() {
			data.Sites[i].AuthenticationProfileName = types.StringValue(value.String())
		} else {
			data.Sites[i].AuthenticationProfileName = types.StringNull()
		}
		if value := r.Get("isPubSubEnabled"); value.Exists() && !data.Sites[i].IsPubSubEnabled.IsNull() {
			data.Sites[i].IsPubSubEnabled = types.BoolValue(value.Bool())
		} else {
			data.Sites[i].IsPubSubEnabled = types.BoolNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricSites) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.Sites) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
