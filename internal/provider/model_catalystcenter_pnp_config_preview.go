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
type PNPConfigPreview struct {
	Id       types.String `tfsdk:"id"`
	DeviceId types.String `tfsdk:"device_id"`
	SiteId   types.String `tfsdk:"site_id"`
	Type     types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PNPConfigPreview) getPath() string {
	return "/dna/intent/api/v1/onboarding/pnp-device/site-config-preview"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PNPConfigPreview) toBody(ctx context.Context, state PNPConfigPreview) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceId.IsNull() {
		body, _ = sjson.Set(body, "deviceId", data.DeviceId.ValueString())
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PNPConfigPreview) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceId"); value.Exists() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PNPConfigPreview) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceId"); value.Exists() && !data.DeviceId.IsNull() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PNPConfigPreview) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceId.IsNull() {
		return false
	}
	if !data.SiteId.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
