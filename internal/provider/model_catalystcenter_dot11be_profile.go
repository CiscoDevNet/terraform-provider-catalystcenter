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
type Dot11beProfile struct {
	Id             types.String `tfsdk:"id"`
	ProfileName    types.String `tfsdk:"profile_name"`
	OfdmaDownLink  types.Bool   `tfsdk:"ofdma_down_link"`
	OfdmaUpLink    types.Bool   `tfsdk:"ofdma_up_link"`
	MuMimoDownLink types.Bool   `tfsdk:"mu_mimo_down_link"`
	MuMimoUpLink   types.Bool   `tfsdk:"mu_mimo_up_link"`
	OfdmaMultiRu   types.Bool   `tfsdk:"ofdma_multi_ru"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Dot11beProfile) getPath() string {
	return "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data Dot11beProfile) getPathDelete() string {
	return "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Dot11beProfile) toBody(ctx context.Context, state Dot11beProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.ProfileName.IsNull() {
		body, _ = sjson.Set(body, "profileName", data.ProfileName.ValueString())
	}
	if !data.OfdmaDownLink.IsNull() {
		body, _ = sjson.Set(body, "ofdmaDownLink", data.OfdmaDownLink.ValueBool())
	}
	if !data.OfdmaUpLink.IsNull() {
		body, _ = sjson.Set(body, "ofdmaUpLink", data.OfdmaUpLink.ValueBool())
	}
	if !data.MuMimoDownLink.IsNull() {
		body, _ = sjson.Set(body, "muMimoDownLink", data.MuMimoDownLink.ValueBool())
	}
	if !data.MuMimoUpLink.IsNull() {
		body, _ = sjson.Set(body, "muMimoUpLink", data.MuMimoUpLink.ValueBool())
	}
	if !data.OfdmaMultiRu.IsNull() {
		body, _ = sjson.Set(body, "ofdmaMultiRu", data.OfdmaMultiRu.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Dot11beProfile) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.profileName"); value.Exists() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("response.0.ofdmaDownLink"); value.Exists() {
		data.OfdmaDownLink = types.BoolValue(value.Bool())
	} else {
		data.OfdmaDownLink = types.BoolValue(true)
	}
	if value := res.Get("response.0.ofdmaUpLink"); value.Exists() {
		data.OfdmaUpLink = types.BoolValue(value.Bool())
	} else {
		data.OfdmaUpLink = types.BoolValue(true)
	}
	if value := res.Get("response.0.muMimoDownLink"); value.Exists() {
		data.MuMimoDownLink = types.BoolValue(value.Bool())
	} else {
		data.MuMimoDownLink = types.BoolValue(false)
	}
	if value := res.Get("response.0.muMimoUpLink"); value.Exists() {
		data.MuMimoUpLink = types.BoolValue(value.Bool())
	} else {
		data.MuMimoUpLink = types.BoolValue(false)
	}
	if value := res.Get("response.0.ofdmaMultiRu"); value.Exists() {
		data.OfdmaMultiRu = types.BoolValue(value.Bool())
	} else {
		data.OfdmaMultiRu = types.BoolValue(false)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Dot11beProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.profileName"); value.Exists() && !data.ProfileName.IsNull() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("response.0.ofdmaDownLink"); value.Exists() && !data.OfdmaDownLink.IsNull() {
		data.OfdmaDownLink = types.BoolValue(value.Bool())
	} else if data.OfdmaDownLink.ValueBool() != true {
		data.OfdmaDownLink = types.BoolNull()
	}
	if value := res.Get("response.0.ofdmaUpLink"); value.Exists() && !data.OfdmaUpLink.IsNull() {
		data.OfdmaUpLink = types.BoolValue(value.Bool())
	} else if data.OfdmaUpLink.ValueBool() != true {
		data.OfdmaUpLink = types.BoolNull()
	}
	if value := res.Get("response.0.muMimoDownLink"); value.Exists() && !data.MuMimoDownLink.IsNull() {
		data.MuMimoDownLink = types.BoolValue(value.Bool())
	} else if data.MuMimoDownLink.ValueBool() != false {
		data.MuMimoDownLink = types.BoolNull()
	}
	if value := res.Get("response.0.muMimoUpLink"); value.Exists() && !data.MuMimoUpLink.IsNull() {
		data.MuMimoUpLink = types.BoolValue(value.Bool())
	} else if data.MuMimoUpLink.ValueBool() != false {
		data.MuMimoUpLink = types.BoolNull()
	}
	if value := res.Get("response.0.ofdmaMultiRu"); value.Exists() && !data.OfdmaMultiRu.IsNull() {
		data.OfdmaMultiRu = types.BoolValue(value.Bool())
	} else if data.OfdmaMultiRu.ValueBool() != false {
		data.OfdmaMultiRu = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Dot11beProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.OfdmaDownLink.IsNull() {
		return false
	}
	if !data.OfdmaUpLink.IsNull() {
		return false
	}
	if !data.MuMimoDownLink.IsNull() {
		return false
	}
	if !data.MuMimoUpLink.IsNull() {
		return false
	}
	if !data.OfdmaMultiRu.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
