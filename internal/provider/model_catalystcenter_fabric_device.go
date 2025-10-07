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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type FabricDevice struct {
	Id                           types.String `tfsdk:"id"`
	NetworkDeviceId              types.String `tfsdk:"network_device_id"`
	FabricId                     types.String `tfsdk:"fabric_id"`
	DeviceRoles                  types.Set    `tfsdk:"device_roles"`
	BorderTypes                  types.Set    `tfsdk:"border_types"`
	LocalAutonomousSystemNumber  types.String `tfsdk:"local_autonomous_system_number"`
	DefaultExit                  types.Bool   `tfsdk:"default_exit"`
	ImportExternalRoutes         types.Bool   `tfsdk:"import_external_routes"`
	BorderPriority               types.Int64  `tfsdk:"border_priority"`
	PrependAutonomousSystemCount types.Int64  `tfsdk:"prepend_autonomous_system_count"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricDevice) getPath() string {
	return "/dna/intent/api/v1/sda/fabricDevices"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricDevice) toBody(ctx context.Context, state FabricDevice) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.NetworkDeviceId.IsNull() {
		body, _ = sjson.Set(body, "0.networkDeviceId", data.NetworkDeviceId.ValueString())
	}
	if !data.FabricId.IsNull() {
		body, _ = sjson.Set(body, "0.fabricId", data.FabricId.ValueString())
	}
	if !data.DeviceRoles.IsNull() {
		var values []string
		data.DeviceRoles.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.deviceRoles", values)
	}
	if !data.BorderTypes.IsNull() {
		var values []string
		data.BorderTypes.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "0.borderDeviceSettings.borderTypes", values)
	}
	if !data.LocalAutonomousSystemNumber.IsNull() {
		body, _ = sjson.Set(body, "0.borderDeviceSettings.layer3Settings.localAutonomousSystemNumber", data.LocalAutonomousSystemNumber.ValueString())
	}
	if !data.DefaultExit.IsNull() {
		body, _ = sjson.Set(body, "0.borderDeviceSettings.layer3Settings.isDefaultExit", data.DefaultExit.ValueBool())
	}
	if !data.ImportExternalRoutes.IsNull() {
		body, _ = sjson.Set(body, "0.borderDeviceSettings.layer3Settings.importExternalRoutes", data.ImportExternalRoutes.ValueBool())
	}
	if !data.BorderPriority.IsNull() {
		body, _ = sjson.Set(body, "0.borderDeviceSettings.layer3Settings.borderPriority", data.BorderPriority.ValueInt64())
	}
	if !data.PrependAutonomousSystemCount.IsNull() {
		body, _ = sjson.Set(body, "0.borderDeviceSettings.layer3Settings.prependAutonomousSystemCount", data.PrependAutonomousSystemCount.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricDevice) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.networkDeviceId"); value.Exists() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.deviceRoles"); value.Exists() && len(value.Array()) > 0 {
		data.DeviceRoles = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceRoles = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.borderDeviceSettings.borderTypes"); value.Exists() && len(value.Array()) > 0 {
		data.BorderTypes = helpers.GetStringSet(value.Array())
	} else {
		data.BorderTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.localAutonomousSystemNumber"); value.Exists() {
		data.LocalAutonomousSystemNumber = types.StringValue(value.String())
	} else {
		data.LocalAutonomousSystemNumber = types.StringNull()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.isDefaultExit"); value.Exists() {
		data.DefaultExit = types.BoolValue(value.Bool())
	} else {
		data.DefaultExit = types.BoolNull()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.importExternalRoutes"); value.Exists() {
		data.ImportExternalRoutes = types.BoolValue(value.Bool())
	} else {
		data.ImportExternalRoutes = types.BoolNull()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.borderPriority"); value.Exists() {
		data.BorderPriority = types.Int64Value(value.Int())
	} else {
		data.BorderPriority = types.Int64Null()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.prependAutonomousSystemCount"); value.Exists() {
		data.PrependAutonomousSystemCount = types.Int64Value(value.Int())
	} else {
		data.PrependAutonomousSystemCount = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricDevice) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.networkDeviceId"); value.Exists() && !data.NetworkDeviceId.IsNull() {
		data.NetworkDeviceId = types.StringValue(value.String())
	} else {
		data.NetworkDeviceId = types.StringNull()
	}
	if value := res.Get("response.0.fabricId"); value.Exists() && !data.FabricId.IsNull() {
		data.FabricId = types.StringValue(value.String())
	} else {
		data.FabricId = types.StringNull()
	}
	if value := res.Get("response.0.deviceRoles"); value.Exists() && !data.DeviceRoles.IsNull() {
		data.DeviceRoles = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceRoles = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.borderDeviceSettings.borderTypes"); value.Exists() && !data.BorderTypes.IsNull() {
		data.BorderTypes = helpers.GetStringSet(value.Array())
	} else {
		data.BorderTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.localAutonomousSystemNumber"); value.Exists() && !data.LocalAutonomousSystemNumber.IsNull() {
		data.LocalAutonomousSystemNumber = types.StringValue(value.String())
	} else {
		data.LocalAutonomousSystemNumber = types.StringNull()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.isDefaultExit"); value.Exists() && !data.DefaultExit.IsNull() {
		data.DefaultExit = types.BoolValue(value.Bool())
	} else {
		data.DefaultExit = types.BoolNull()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.importExternalRoutes"); value.Exists() && !data.ImportExternalRoutes.IsNull() {
		data.ImportExternalRoutes = types.BoolValue(value.Bool())
	} else {
		data.ImportExternalRoutes = types.BoolNull()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.borderPriority"); value.Exists() && !data.BorderPriority.IsNull() {
		data.BorderPriority = types.Int64Value(value.Int())
	} else {
		data.BorderPriority = types.Int64Null()
	}
	if value := res.Get("response.0.borderDeviceSettings.layer3Settings.prependAutonomousSystemCount"); value.Exists() && !data.PrependAutonomousSystemCount.IsNull() {
		data.PrependAutonomousSystemCount = types.Int64Value(value.Int())
	} else {
		data.PrependAutonomousSystemCount = types.Int64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricDevice) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceRoles.IsNull() {
		return false
	}
	if !data.BorderTypes.IsNull() {
		return false
	}
	if !data.LocalAutonomousSystemNumber.IsNull() {
		return false
	}
	if !data.DefaultExit.IsNull() {
		return false
	}
	if !data.ImportExternalRoutes.IsNull() {
		return false
	}
	if !data.BorderPriority.IsNull() {
		return false
	}
	if !data.PrependAutonomousSystemCount.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
