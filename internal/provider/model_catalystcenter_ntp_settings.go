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
type NTPSettings struct {
	Id      types.String `tfsdk:"id"`
	SiteId  types.String `tfsdk:"site_id"`
	Servers types.Set    `tfsdk:"servers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NTPSettings) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/ntpSettings", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data NTPSettings) toBody(ctx context.Context, state NTPSettings) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Servers.IsNull() {
		var values []string
		data.Servers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ntp.servers", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *NTPSettings) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get(""); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.ntp.servers"); value.Exists() && len(value.Array()) > 0 {
		data.Servers = helpers.GetStringSet(value.Array())
	} else {
		data.Servers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *NTPSettings) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.ntp.servers"); value.Exists() && !data.Servers.IsNull() {
		data.Servers = helpers.GetStringSet(value.Array())
	} else {
		data.Servers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NTPSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.ntp.servers"); value.Exists() && !data.Servers.IsNull() {
		data.Servers = helpers.GetStringSet(value.Array())
	} else {
		data.Servers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *NTPSettings) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Servers.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
