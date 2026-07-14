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

import (
	"context"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Custom normalizeFraSensitivity: API returns title-case values (e.g. "Medium", "Even Higher")
// while the schema uses uppercase enum values (e.g. "MEDIUM", "EVEN_HIGHER").
func normalizeFraSensitivity(apiValue string) string {
	if apiValue == "" {
		return apiValue
	}
	return strings.ToUpper(strings.ReplaceAll(apiValue, " ", "_"))
}

// fraSensitivityToAPI converts Terraform enum values to API title-case strings.
func fraSensitivityToAPI(tfValue string) string {
	if tfValue == "" {
		return tfValue
	}
	words := strings.Split(strings.ToLower(tfValue), "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessRRMFRAConfiguration struct {
	Id                 types.String `tfsdk:"id"`
	DesignName         types.String `tfsdk:"design_name"`
	RadioBand          types.String `tfsdk:"radio_band"`
	FraFreeze          types.Bool   `tfsdk:"fra_freeze"`
	FraStatus          types.Bool   `tfsdk:"fra_status"`
	FraInterval        types.Int64  `tfsdk:"fra_interval"`
	FraSensitivity     types.String `tfsdk:"fra_sensitivity"`
	UnlockedAttributes types.List   `tfsdk:"unlocked_attributes"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessRRMFRAConfiguration) getPath() string {
	return "/dna/intent/api/v1/featureTemplates/wireless/rrmFraConfigurations"
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

// Custom toBody: fraSensitivity converted to API title-case via fraSensitivityToAPI.
func (data WirelessRRMFRAConfiguration) toBody(ctx context.Context, state WirelessRRMFRAConfiguration) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DesignName.IsNull() {
		body, _ = sjson.Set(body, "designName", data.DesignName.ValueString())
	}
	if !data.RadioBand.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.radioBand", data.RadioBand.ValueString())
	}
	if !data.FraFreeze.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.fraFreeze", data.FraFreeze.ValueBool())
	}
	if !data.FraStatus.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.fraStatus", data.FraStatus.ValueBool())
	}
	if !data.FraInterval.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.fraInterval", data.FraInterval.ValueInt64())
	}
	if !data.FraSensitivity.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.fraSensitivity", fraSensitivityToAPI(data.FraSensitivity.ValueString()))
	}
	if !data.UnlockedAttributes.IsNull() {
		var values []string
		data.UnlockedAttributes.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "unlockedAttributes", values)
	}
	return body
}

// Custom fromBody: fraSensitivity normalized via normalizeFraSensitivity.
func (data *WirelessRRMFRAConfiguration) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.designName"); value.Exists() {
		data.DesignName = types.StringValue(value.String())
	} else {
		data.DesignName = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.radioBand"); value.Exists() {
		data.RadioBand = types.StringValue(value.String())
	} else {
		data.RadioBand = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.fraFreeze"); value.Exists() {
		data.FraFreeze = types.BoolValue(value.Bool())
	} else {
		data.FraFreeze = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.fraStatus"); value.Exists() {
		data.FraStatus = types.BoolValue(value.Bool())
	} else {
		data.FraStatus = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.fraInterval"); value.Exists() {
		data.FraInterval = types.Int64Value(value.Int())
	} else {
		data.FraInterval = types.Int64Null()
	}
	if value := res.Get("response.featureAttributes.fraSensitivity"); value.Exists() {
		data.FraSensitivity = types.StringValue(normalizeFraSensitivity(value.String()))
	} else {
		data.FraSensitivity = types.StringNull()
	}
	if value := res.Get("response.unlockedAttributes"); value.Exists() && len(value.Array()) > 0 {
		data.UnlockedAttributes = helpers.GetStringList(value.Array())
	} else {
		data.UnlockedAttributes = types.ListNull(types.StringType)
	}
}

// Custom updateFromBody: fraSensitivity normalized via normalizeFraSensitivity.
func (data *WirelessRRMFRAConfiguration) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.designName"); value.Exists() && !data.DesignName.IsNull() {
		data.DesignName = types.StringValue(value.String())
	} else {
		data.DesignName = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.radioBand"); value.Exists() && !data.RadioBand.IsNull() {
		data.RadioBand = types.StringValue(value.String())
	} else {
		data.RadioBand = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.fraFreeze"); value.Exists() && !data.FraFreeze.IsNull() {
		data.FraFreeze = types.BoolValue(value.Bool())
	} else {
		data.FraFreeze = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.fraStatus"); value.Exists() && !data.FraStatus.IsNull() {
		data.FraStatus = types.BoolValue(value.Bool())
	} else {
		data.FraStatus = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.fraInterval"); value.Exists() && !data.FraInterval.IsNull() {
		data.FraInterval = types.Int64Value(value.Int())
	} else {
		data.FraInterval = types.Int64Null()
	}
	if value := res.Get("response.featureAttributes.fraSensitivity"); value.Exists() && !data.FraSensitivity.IsNull() {
		data.FraSensitivity = types.StringValue(normalizeFraSensitivity(value.String()))
	} else {
		data.FraSensitivity = types.StringNull()
	}
	if value := res.Get("response.unlockedAttributes"); value.Exists() && !data.UnlockedAttributes.IsNull() {
		data.UnlockedAttributes = helpers.GetStringList(value.Array())
	} else {
		data.UnlockedAttributes = types.ListNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessRRMFRAConfiguration) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DesignName.IsNull() {
		return false
	}
	if !data.RadioBand.IsNull() {
		return false
	}
	if !data.FraFreeze.IsNull() {
		return false
	}
	if !data.FraStatus.IsNull() {
		return false
	}
	if !data.FraInterval.IsNull() {
		return false
	}
	if !data.FraSensitivity.IsNull() {
		return false
	}
	if !data.UnlockedAttributes.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
