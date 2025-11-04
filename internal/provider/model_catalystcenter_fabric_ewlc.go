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
type FabricEWLC struct {
	Id                     types.String `tfsdk:"id"`
	FabricId               types.String `tfsdk:"fabric_id"`
	NetworkDeviceId        types.String `tfsdk:"network_device_id"`
	EnableWireless         types.Bool   `tfsdk:"enable_wireless"`
	EnableRollingApUpgrade types.Bool   `tfsdk:"enable_rolling_ap_upgrade"`
	ApRebootPercentage     types.Int64  `tfsdk:"ap_reboot_percentage"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricEWLC) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sda/fabrics/%v/switchWirelessSetting", url.QueryEscape(data.FabricId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricEWLC) toBody(ctx context.Context, state FabricEWLC) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "id", data.NetworkDeviceId.ValueString())
	}
	if !data.EnableWireless.IsNull() {
		body, _ = sjson.Set(body, "enableWireless", data.EnableWireless.ValueBool())
	}
	if !data.EnableRollingApUpgrade.IsNull() {
		body, _ = sjson.Set(body, "rollingApUpgrade.enableRollingApUpgrade", data.EnableRollingApUpgrade.ValueBool())
	}
	if !data.ApRebootPercentage.IsNull() {
		body, _ = sjson.Set(body, "rollingApUpgrade.apRebootPercentage", data.ApRebootPercentage.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricEWLC) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("enableWireless"); value.Exists() {
		data.EnableWireless = types.BoolValue(value.Bool())
	} else {
		data.EnableWireless = types.BoolNull()
	}
	if value := res.Get("rollingApUpgrade.enableRollingApUpgrade"); value.Exists() {
		data.EnableRollingApUpgrade = types.BoolValue(value.Bool())
	} else {
		data.EnableRollingApUpgrade = types.BoolNull()
	}
	if value := res.Get("rollingApUpgrade.apRebootPercentage"); value.Exists() {
		data.ApRebootPercentage = types.Int64Value(value.Int())
	} else {
		data.ApRebootPercentage = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricEWLC) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() && !data.NetworkDeviceId.IsNull() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("enableWireless"); value.Exists() && !data.EnableWireless.IsNull() {
		data.EnableWireless = types.BoolValue(value.Bool())
	} else {
		data.EnableWireless = types.BoolNull()
	}
	if value := res.Get("rollingApUpgrade.enableRollingApUpgrade"); value.Exists() && !data.EnableRollingApUpgrade.IsNull() {
		data.EnableRollingApUpgrade = types.BoolValue(value.Bool())
	} else {
		data.EnableRollingApUpgrade = types.BoolNull()
	}
	if value := res.Get("rollingApUpgrade.apRebootPercentage"); value.Exists() && !data.ApRebootPercentage.IsNull() {
		data.ApRebootPercentage = types.Int64Value(value.Int())
	} else {
		data.ApRebootPercentage = types.Int64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricEWLC) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.EnableWireless.IsNull() {
		return false
	}
	if !data.EnableRollingApUpgrade.IsNull() {
		return false
	}
	if !data.ApRebootPercentage.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
