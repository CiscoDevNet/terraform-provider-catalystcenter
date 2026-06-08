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
type SyncDeviceCredential struct {
	Id                 types.String `tfsdk:"id"`
	SiteId             types.String `tfsdk:"site_id"`
	DeviceCredentialId types.String `tfsdk:"device_credential_id"`
	ConfigureDevice    types.Bool   `tfsdk:"configure_device"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SyncDeviceCredential) getPath() string {
	return "/dna/intent/api/v1/sites/deviceCredentials/apply"
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
func (data SyncDeviceCredential) toBody(ctx context.Context, state SyncDeviceCredential) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if !data.DeviceCredentialId.IsNull() {
		body, _ = sjson.Set(body, "deviceCredentialId", data.DeviceCredentialId.ValueString())
	}
	if !data.ConfigureDevice.IsNull() {
		body, _ = sjson.Set(body, "configureDevice", data.ConfigureDevice.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SyncDeviceCredential) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("deviceCredentialId"); value.Exists() {
		data.DeviceCredentialId = types.StringValue(value.String())
	} else {
		data.DeviceCredentialId = types.StringNull()
	}
	if value := res.Get("configureDevice"); value.Exists() {
		data.ConfigureDevice = types.BoolValue(value.Bool())
	} else {
		data.ConfigureDevice = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SyncDeviceCredential) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("deviceCredentialId"); value.Exists() && !data.DeviceCredentialId.IsNull() {
		data.DeviceCredentialId = types.StringValue(value.String())
	} else {
		data.DeviceCredentialId = types.StringNull()
	}
	if value := res.Get("configureDevice"); value.Exists() && !data.ConfigureDevice.IsNull() {
		data.ConfigureDevice = types.BoolValue(value.Bool())
	} else {
		data.ConfigureDevice = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SyncDeviceCredential) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceCredentialId.IsNull() {
		return false
	}
	if !data.ConfigureDevice.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
