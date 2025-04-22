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
type UpdateDeviceManagementAddress struct {
	Id       types.String `tfsdk:"id"`
	DeviceId types.String `tfsdk:"device_id"`
	NewIp    types.String `tfsdk:"new_ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data UpdateDeviceManagementAddress) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/network-device/%v/management-address", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data UpdateDeviceManagementAddress) toBody(ctx context.Context, state UpdateDeviceManagementAddress) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.NewIp.IsNull() {
		body, _ = sjson.Set(body, "newIP", data.NewIp.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *UpdateDeviceManagementAddress) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.managementIpAddress"); value.Exists() {
		data.NewIp = types.StringValue(value.String())
	} else {
		data.NewIp = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *UpdateDeviceManagementAddress) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.managementIpAddress"); value.Exists() && !data.NewIp.IsNull() {
		data.NewIp = types.StringValue(value.String())
	} else {
		data.NewIp = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *UpdateDeviceManagementAddress) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.NewIp.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
