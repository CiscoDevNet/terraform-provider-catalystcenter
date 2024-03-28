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
type ImageDistribution struct {
	Id         types.String `tfsdk:"id"`
	DeviceUuid types.String `tfsdk:"device_uuid"`
	ImageUuid  types.String `tfsdk:"image_uuid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ImageDistribution) getPath() string {
	return "/dna/intent/api/v1/image/distribution"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ImageDistribution) toBody(ctx context.Context, state ImageDistribution) string {
	body := ""
	if !data.DeviceUuid.IsNull() {
		body, _ = sjson.Set(body, "0.deviceUuid", data.DeviceUuid.ValueString())
	}
	if !data.ImageUuid.IsNull() {
		body, _ = sjson.Set(body, "0.imageUuid", data.ImageUuid.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ImageDistribution) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceUuid"); value.Exists() {
		data.DeviceUuid = types.StringValue(value.String())
	} else {
		data.DeviceUuid = types.StringNull()
	}
	if value := res.Get("0.imageUuid"); value.Exists() {
		data.ImageUuid = types.StringValue(value.String())
	} else {
		data.ImageUuid = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ImageDistribution) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("0.deviceUuid"); value.Exists() && !data.DeviceUuid.IsNull() {
		data.DeviceUuid = types.StringValue(value.String())
	} else {
		data.DeviceUuid = types.StringNull()
	}
	if value := res.Get("0.imageUuid"); value.Exists() && !data.ImageUuid.IsNull() {
		data.ImageUuid = types.StringValue(value.String())
	} else {
		data.ImageUuid = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ImageDistribution) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceUuid.IsNull() {
		return false
	}
	if !data.ImageUuid.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
