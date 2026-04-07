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
type DeviceReplacement struct {
	Id                            types.String `tfsdk:"id"`
	FaultyDeviceId                types.String `tfsdk:"faulty_device_id"`
	ReplacementStatus             types.String `tfsdk:"replacement_status"`
	Family                        types.String `tfsdk:"family"`
	FaultyDeviceName              types.String `tfsdk:"faulty_device_name"`
	FaultyDevicePlatform          types.String `tfsdk:"faulty_device_platform"`
	FaultyDeviceSerialNumber      types.String `tfsdk:"faulty_device_serial_number"`
	ReplacementDevicePlatform     types.String `tfsdk:"replacement_device_platform"`
	ReplacementDeviceSerialNumber types.String `tfsdk:"replacement_device_serial_number"`
	NeighbourDeviceId             types.String `tfsdk:"neighbour_device_id"`
	NetworkReadinessTaskId        types.String `tfsdk:"network_readiness_task_id"`
	WorkflowId                    types.String `tfsdk:"workflow_id"`
	CreationTime                  types.Int64  `tfsdk:"creation_time"`
	ReplacementTime               types.Int64  `tfsdk:"replacement_time"`
	WorkflowFailedStep            types.String `tfsdk:"workflow_failed_step"`
	ReadinessCheckTaskId          types.String `tfsdk:"readiness_check_task_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DeviceReplacement) getPath() string {
	return "/dna/intent/api/v1/device-replacement"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DeviceReplacement) toBody(ctx context.Context, state DeviceReplacement) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.FaultyDeviceId.IsNull() {
		body, _ = sjson.Set(body, "0.faultyDeviceId", data.FaultyDeviceId.ValueString())
	}
	if !data.ReplacementStatus.IsNull() {
		body, _ = sjson.Set(body, "0.replacementStatus", data.ReplacementStatus.ValueString())
	}
	if data.Family.ValueString() != "" && !data.Family.IsNull() {
		body, _ = sjson.Set(body, "family", data.Family.ValueString())
	}
	if data.FaultyDeviceName.ValueString() != "" && !data.FaultyDeviceName.IsNull() {
		body, _ = sjson.Set(body, "faultyDeviceName", data.FaultyDeviceName.ValueString())
	}
	if data.FaultyDevicePlatform.ValueString() != "" && !data.FaultyDevicePlatform.IsNull() {
		body, _ = sjson.Set(body, "faultyDevicePlatform", data.FaultyDevicePlatform.ValueString())
	}
	if data.FaultyDeviceSerialNumber.ValueString() != "" && !data.FaultyDeviceSerialNumber.IsNull() {
		body, _ = sjson.Set(body, "faultyDeviceSerialNumber", data.FaultyDeviceSerialNumber.ValueString())
	}
	if !data.ReplacementDevicePlatform.IsNull() {
		body, _ = sjson.Set(body, "0.replacementDevicePlatform", data.ReplacementDevicePlatform.ValueString())
	}
	if !data.ReplacementDeviceSerialNumber.IsNull() {
		body, _ = sjson.Set(body, "0.replacementDeviceSerialNumber", data.ReplacementDeviceSerialNumber.ValueString())
	}
	if !data.NeighbourDeviceId.IsNull() {
		body, _ = sjson.Set(body, "0.neighbourDeviceId", data.NeighbourDeviceId.ValueString())
	}
	if data.NetworkReadinessTaskId.ValueString() != "" && !data.NetworkReadinessTaskId.IsNull() {
		body, _ = sjson.Set(body, "networkReadinessTaskId", data.NetworkReadinessTaskId.ValueString())
	}
	if data.WorkflowId.ValueString() != "" && !data.WorkflowId.IsNull() {
		body, _ = sjson.Set(body, "workflowId", data.WorkflowId.ValueString())
	}
	if data.CreationTime.ValueInt64() != 0 && !data.CreationTime.IsNull() {
		body, _ = sjson.Set(body, "creationTime", data.CreationTime.ValueInt64())
	}
	if data.ReplacementTime.ValueInt64() != 0 && !data.ReplacementTime.IsNull() {
		body, _ = sjson.Set(body, "replacementTime", data.ReplacementTime.ValueInt64())
	}
	if data.WorkflowFailedStep.ValueString() != "" && !data.WorkflowFailedStep.IsNull() {
		body, _ = sjson.Set(body, "workflowFailedStep", data.WorkflowFailedStep.ValueString())
	}
	if data.ReadinessCheckTaskId.ValueString() != "" && !data.ReadinessCheckTaskId.IsNull() {
		body, _ = sjson.Set(body, "readinesscheckTaskId", data.ReadinessCheckTaskId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DeviceReplacement) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("faultyDeviceId"); value.Exists() {
		data.FaultyDeviceId = types.StringValue(value.String())
	} else {
		data.FaultyDeviceId = types.StringNull()
	}
	if value := res.Get("replacementStatus"); value.Exists() {
		data.ReplacementStatus = types.StringValue(value.String())
	} else {
		data.ReplacementStatus = types.StringNull()
	}
	if value := res.Get("family"); value.Exists() {
		data.Family = types.StringValue(value.String())
	} else {
		data.Family = types.StringNull()
	}
	if value := res.Get("faultyDeviceName"); value.Exists() {
		data.FaultyDeviceName = types.StringValue(value.String())
	} else {
		data.FaultyDeviceName = types.StringNull()
	}
	if value := res.Get("faultyDevicePlatform"); value.Exists() {
		data.FaultyDevicePlatform = types.StringValue(value.String())
	} else {
		data.FaultyDevicePlatform = types.StringNull()
	}
	if value := res.Get("faultyDeviceSerialNumber"); value.Exists() {
		data.FaultyDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.FaultyDeviceSerialNumber = types.StringNull()
	}
	if value := res.Get("replacementDevicePlatform"); value.Exists() {
		data.ReplacementDevicePlatform = types.StringValue(value.String())
	} else {
		data.ReplacementDevicePlatform = types.StringNull()
	}
	if value := res.Get("replacementDeviceSerialNumber"); value.Exists() {
		data.ReplacementDeviceSerialNumber = types.StringValue(value.String())
	} else {
		data.ReplacementDeviceSerialNumber = types.StringNull()
	}
	if value := res.Get("neighbourDeviceId"); value.Exists() {
		data.NeighbourDeviceId = types.StringValue(value.String())
	} else {
		data.NeighbourDeviceId = types.StringNull()
	}
	if value := res.Get("networkReadinessTaskId"); value.Exists() {
		data.NetworkReadinessTaskId = types.StringValue(value.String())
	} else {
		data.NetworkReadinessTaskId = types.StringNull()
	}
	if value := res.Get("workflowId"); value.Exists() {
		data.WorkflowId = types.StringValue(value.String())
	} else {
		data.WorkflowId = types.StringNull()
	}
	if value := res.Get("creationTime"); value.Exists() {
		data.CreationTime = types.Int64Value(value.Int())
	} else {
		data.CreationTime = types.Int64Null()
	}
	if value := res.Get("replacementTime"); value.Exists() {
		data.ReplacementTime = types.Int64Value(value.Int())
	} else {
		data.ReplacementTime = types.Int64Null()
	}
	if value := res.Get("workflowFailedStep"); value.Exists() {
		data.WorkflowFailedStep = types.StringValue(value.String())
	} else {
		data.WorkflowFailedStep = types.StringNull()
	}
	if value := res.Get("readinesscheckTaskId"); value.Exists() {
		data.ReadinessCheckTaskId = types.StringValue(value.String())
	} else {
		data.ReadinessCheckTaskId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below has custom code (not generated). Do not add template markers.
// Custom updateFromBody: Only faultyDeviceId is updated from the API response.
// ALL other fields are preserved from existing state because the API either:
//   - computes them from faultyDeviceId (family, faultyDeviceName, etc.)
//   - auto-transitions them (replacementStatus)
//   - returns null for user-supplied optional values (replacementDevicePlatform,
//     replacementDeviceSerialNumber, neighbourDeviceId) — the API accepts these
//     in POST/PUT but does NOT persist or return them in GET.
//   - are API-generated IDs (workflowId, networkReadinessTaskId, etc.)
//
// Import uses fromBody which reads all fields from the API response.
func (data *DeviceReplacement) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("faultyDeviceId"); value.Exists() && !data.FaultyDeviceId.IsNull() {
		data.FaultyDeviceId = types.StringValue(value.String())
	} else {
		data.FaultyDeviceId = types.StringNull()
	}
	// All other fields: keep existing state.
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DeviceReplacement) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FaultyDeviceId.IsNull() {
		return false
	}
	if !data.ReplacementStatus.IsNull() {
		return false
	}
	if !data.Family.IsNull() {
		return false
	}
	if !data.FaultyDeviceName.IsNull() {
		return false
	}
	if !data.FaultyDevicePlatform.IsNull() {
		return false
	}
	if !data.FaultyDeviceSerialNumber.IsNull() {
		return false
	}
	if !data.ReplacementDevicePlatform.IsNull() {
		return false
	}
	if !data.ReplacementDeviceSerialNumber.IsNull() {
		return false
	}
	if !data.NeighbourDeviceId.IsNull() {
		return false
	}
	if !data.NetworkReadinessTaskId.IsNull() {
		return false
	}
	if !data.WorkflowId.IsNull() {
		return false
	}
	if !data.CreationTime.IsNull() {
		return false
	}
	if !data.ReplacementTime.IsNull() {
		return false
	}
	if !data.WorkflowFailedStep.IsNull() {
		return false
	}
	if !data.ReadinessCheckTaskId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
