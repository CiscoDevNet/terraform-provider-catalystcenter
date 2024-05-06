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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Image struct {
	Id                        types.String `tfsdk:"id"`
	ThirdPartyApplicationType types.String `tfsdk:"third_party_application_type"`
	Family                    types.String `tfsdk:"family"`
	SourcePath                types.String `tfsdk:"source_path"`
	Name                      types.String `tfsdk:"name"`
	ThirdPartyVendor          types.String `tfsdk:"third_party_vendor"`
	IsThirdParty              types.Bool   `tfsdk:"is_third_party"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Image) getPath() string {
	return "/dna/intent/api/v1/image/importation/source/file"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data Image) getPathDelete() string {
	return "/api/v1/image/importation"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Image) toBody(ctx context.Context, state Image) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.ThirdPartyApplicationType.IsNull() {
		body, _ = sjson.Set(body, "thirdPartyApplicationType", data.ThirdPartyApplicationType.ValueString())
	}
	if !data.Family.IsNull() {
		body, _ = sjson.Set(body, "0.thirdPartyImageFamily", data.Family.ValueString())
	}
	if !data.SourcePath.IsNull() {
		body, _ = sjson.Set(body, "sourcePath", data.SourcePath.ValueString())
	}
	if !data.ThirdPartyVendor.IsNull() {
		body, _ = sjson.Set(body, "0.thirdPartyVendor", data.ThirdPartyVendor.ValueString())
	}
	if !data.IsThirdParty.IsNull() {
		body, _ = sjson.Set(body, "0.isThirdParty", data.IsThirdParty.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Image) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.applicationType"); value.Exists() {
		data.ThirdPartyApplicationType = types.StringValue(value.String())
	} else {
		data.ThirdPartyApplicationType = types.StringNull()
	}
	if value := res.Get("response.0.family"); value.Exists() {
		data.Family = types.StringValue(value.String())
	} else {
		data.Family = types.StringNull()
	}
	if value := res.Get("response.0.vendor"); value.Exists() {
		data.ThirdPartyVendor = types.StringValue(value.String())
	} else {
		data.ThirdPartyVendor = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Image) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.applicationType"); value.Exists() && !data.ThirdPartyApplicationType.IsNull() {
		data.ThirdPartyApplicationType = types.StringValue(value.String())
	} else {
		data.ThirdPartyApplicationType = types.StringNull()
	}
	if value := res.Get("response.0.family"); value.Exists() && !data.Family.IsNull() {
		data.Family = types.StringValue(value.String())
	} else {
		data.Family = types.StringNull()
	}
	if value := res.Get("response.0.vendor"); value.Exists() && !data.ThirdPartyVendor.IsNull() {
		data.ThirdPartyVendor = types.StringValue(value.String())
	} else {
		data.ThirdPartyVendor = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Image) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.ThirdPartyApplicationType.IsNull() {
		return false
	}
	if !data.Family.IsNull() {
		return false
	}
	if !data.SourcePath.IsNull() {
		return false
	}
	if !data.ThirdPartyVendor.IsNull() {
		return false
	}
	if !data.IsThirdParty.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
