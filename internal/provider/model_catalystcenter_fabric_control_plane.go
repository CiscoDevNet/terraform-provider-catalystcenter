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
type FabricControlPlane struct {
	Id                        types.String `tfsdk:"id"`
	DeviceManagementIpAddress types.String `tfsdk:"device_management_ip_address"`
	SiteNameHierarchy         types.String `tfsdk:"site_name_hierarchy"`
	RouteDistributionProtocol types.String `tfsdk:"route_distribution_protocol"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FabricControlPlane) getPath() string {
	return "/dna/intent/api/v1/business/sda/control-plane-device"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FabricControlPlane) toBody(ctx context.Context, state FabricControlPlane) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceManagementIpAddress.IsNull() {
		body, _ = sjson.Set(body, "deviceManagementIpAddress", data.DeviceManagementIpAddress.ValueString())
	}
	if !data.SiteNameHierarchy.IsNull() {
		body, _ = sjson.Set(body, "siteNameHierarchy", data.SiteNameHierarchy.ValueString())
	}
	if !data.RouteDistributionProtocol.IsNull() {
		body, _ = sjson.Set(body, "routeDistributionProtocol", data.RouteDistributionProtocol.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FabricControlPlane) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get(""); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("deviceManagementIpAddress"); value.Exists() {
		data.DeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.DeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("siteNameHierarchy"); value.Exists() {
		data.SiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.SiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("routeDistributionProtocol"); value.Exists() {
		data.RouteDistributionProtocol = types.StringValue(value.String())
	} else {
		data.RouteDistributionProtocol = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *FabricControlPlane) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceManagementIpAddress"); value.Exists() && !data.DeviceManagementIpAddress.IsNull() {
		data.DeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.DeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("siteNameHierarchy"); value.Exists() && !data.SiteNameHierarchy.IsNull() {
		data.SiteNameHierarchy = types.StringValue(value.String())
	} else {
		data.SiteNameHierarchy = types.StringNull()
	}
	if value := res.Get("routeDistributionProtocol"); value.Exists() && !data.RouteDistributionProtocol.IsNull() {
		data.RouteDistributionProtocol = types.StringValue(value.String())
	} else {
		data.RouteDistributionProtocol = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FabricControlPlane) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.DeviceManagementIpAddress.IsUnknown() {
		if value := res.Get("deviceManagementIpAddress"); value.Exists() && !data.DeviceManagementIpAddress.IsNull() {
			data.DeviceManagementIpAddress = types.StringValue(value.String())
		} else {
			data.DeviceManagementIpAddress = types.StringNull()
		}
	}
	if data.SiteNameHierarchy.IsUnknown() {
		if value := res.Get("siteNameHierarchy"); value.Exists() && !data.SiteNameHierarchy.IsNull() {
			data.SiteNameHierarchy = types.StringValue(value.String())
		} else {
			data.SiteNameHierarchy = types.StringNull()
		}
	}
	if data.RouteDistributionProtocol.IsUnknown() {
		if value := res.Get("routeDistributionProtocol"); value.Exists() && !data.RouteDistributionProtocol.IsNull() {
			data.RouteDistributionProtocol = types.StringValue(value.String())
		} else {
			data.RouteDistributionProtocol = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *FabricControlPlane) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.SiteNameHierarchy.IsNull() {
		return false
	}
	if !data.RouteDistributionProtocol.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
