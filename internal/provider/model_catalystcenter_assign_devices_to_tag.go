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
type AssignDevicesToTag struct {
	Id        types.String `tfsdk:"id"`
	TagId     types.String `tfsdk:"tag_id"`
	DeviceIds types.Set    `tfsdk:"device_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssignDevicesToTag) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/tag/%v/member", url.QueryEscape(data.TagId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AssignDevicesToTag) toBody(ctx context.Context, state AssignDevicesToTag) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceIds.IsNull() {
		var values []string
		data.DeviceIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "networkdevice", values)
	}
	return body
}

// End of section. //template:end toBody

func (data *AssignDevicesToTag) fromBody(ctx context.Context, res gjson.Result) {
	// Extract and convert the list of instanceUuids
	instanceUuidAttrValues := extractInstanceUuids(res)

	// Check if instanceUuidAttrValues is not empty, assign it to DeviceIds, else assign a null list
	if len(instanceUuidAttrValues) > 0 {
		data.DeviceIds = types.SetValueMust(types.StringType, instanceUuidAttrValues)
	} else {
		data.DeviceIds = types.SetNull(types.StringType)
	}
}

func (data *AssignDevicesToTag) updateFromBody(ctx context.Context, res gjson.Result) {
	instanceUuidAttrValues := extractInstanceUuids(res)

	// Check if instanceUuidAttrValues is not empty, assign it to DeviceIds, else assign a null list
	if len(instanceUuidAttrValues) > 0 {
		data.DeviceIds = types.SetValueMust(types.StringType, instanceUuidAttrValues)
	} else {
		data.DeviceIds = types.SetNull(types.StringType)
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssignDevicesToTag) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceIds.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
