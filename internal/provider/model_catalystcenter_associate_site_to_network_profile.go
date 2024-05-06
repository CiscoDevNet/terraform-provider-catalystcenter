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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AssociateSiteToNetworkProfile struct {
	Id               types.String `tfsdk:"id"`
	NetworkProfileId types.String `tfsdk:"network_profile_id"`
	SiteId           types.String `tfsdk:"site_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssociateSiteToNetworkProfile) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/networkprofile/%v/site/%v", url.QueryEscape(data.NetworkProfileId.ValueString()), url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AssociateSiteToNetworkProfile) toBody(ctx context.Context, state AssociateSiteToNetworkProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AssociateSiteToNetworkProfile) fromBody(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AssociateSiteToNetworkProfile) updateFromBody(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssociateSiteToNetworkProfile) isNull(ctx context.Context, res gjson.Result) bool {
	return true
}

// End of section. //template:end isNull
