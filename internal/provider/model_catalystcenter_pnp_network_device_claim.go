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
type PnPNetworkDeviceClaim struct {
	Id                     types.String                              `tfsdk:"id"`
	DeviceId               types.String                              `tfsdk:"device_id"`
	DeviceType             types.String                              `tfsdk:"device_type"`
	SiteId                 types.String                              `tfsdk:"site_id"`
	Hostname               types.String                              `tfsdk:"hostname"`
	ImageId                types.String                              `tfsdk:"image_id"`
	RemoveInactive         types.Bool                                `tfsdk:"remove_inactive"`
	TemplateId             types.String                              `tfsdk:"template_id"`
	TemplateParameters     []PnPNetworkDeviceClaimTemplateParameters `tfsdk:"template_parameters"`
	RfProfile              types.String                              `tfsdk:"rf_profile"`
	TopOfStackSerialNumber types.String                              `tfsdk:"top_of_stack_serial_number"`
	CablingScheme          types.String                              `tfsdk:"cabling_scheme"`
	SensorProfile          types.String                              `tfsdk:"sensor_profile"`
	StaticIpAddress        types.String                              `tfsdk:"static_ip_address"`
	SubnetMask             types.String                              `tfsdk:"subnet_mask"`
	Gateway                types.String                              `tfsdk:"gateway"`
	VlanId                 types.Int64                               `tfsdk:"vlan_id"`
	IpInterfaceName        types.String                              `tfsdk:"ip_interface_name"`
	Domain                 types.Int64                               `tfsdk:"domain"`
	SvlMembers             []PnPNetworkDeviceClaimSvlMembers         `tfsdk:"svl_members"`
}

type PnPNetworkDeviceClaimTemplateParameters struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type PnPNetworkDeviceClaimSvlMembers struct {
	SerialNumber    types.String                              `tfsdk:"serial_number"`
	Role            types.String                              `tfsdk:"role"`
	SvlLinks        []PnPNetworkDeviceClaimSvlMembersSvlLinks `tfsdk:"svl_links"`
	LocalInterface  types.String                              `tfsdk:"local_interface"`
	RemoteInterface types.String                              `tfsdk:"remote_interface"`
}

type PnPNetworkDeviceClaimSvlMembersSvlLinks struct {
	LocalInterface  types.String `tfsdk:"local_interface"`
	RemoteInterface types.String `tfsdk:"remote_interface"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PnPNetworkDeviceClaim) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/pnpNetworkDevices/%v/claim", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data PnPNetworkDeviceClaim) getPathDelete() string {
	return fmt.Sprintf("/dna/intent/api/v1/pnpNetworkDevices/%v", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

func (data PnPNetworkDeviceClaim) getPathGet() string {
	return fmt.Sprintf("/dna/intent/api/v1/pnpNetworkDevices/%v", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PnPNetworkDeviceClaim) toBody(ctx context.Context, state PnPNetworkDeviceClaim) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceType.IsNull() {
		body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, "hostname", data.Hostname.ValueString())
	}
	if !data.ImageId.IsNull() {
		body, _ = sjson.Set(body, "imageInfo.imageId", data.ImageId.ValueString())
	}
	if !data.RemoveInactive.IsNull() {
		body, _ = sjson.Set(body, "imageInfo.removeInactive", data.RemoveInactive.ValueBool())
	}
	if !data.TemplateId.IsNull() {
		body, _ = sjson.Set(body, "templateInfo.templateId", data.TemplateId.ValueString())
	}
	if len(data.TemplateParameters) > 0 {
		body, _ = sjson.Set(body, "templateInfo.templateParameters", []interface{}{})
		for _, item := range data.TemplateParameters {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "key", item.Name.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "templateInfo.templateParameters.-1", itemBody)
		}
	}
	if !data.RfProfile.IsNull() {
		body, _ = sjson.Set(body, "rfProfile", data.RfProfile.ValueString())
	}
	if !data.TopOfStackSerialNumber.IsNull() {
		body, _ = sjson.Set(body, "topOfStackSerialNumber", data.TopOfStackSerialNumber.ValueString())
	}
	if !data.CablingScheme.IsNull() {
		body, _ = sjson.Set(body, "cablingScheme", data.CablingScheme.ValueString())
	}
	if !data.SensorProfile.IsNull() {
		body, _ = sjson.Set(body, "sensorProfile", data.SensorProfile.ValueString())
	}
	if !data.StaticIpAddress.IsNull() {
		body, _ = sjson.Set(body, "staticIpAddress", data.StaticIpAddress.ValueString())
	}
	if !data.SubnetMask.IsNull() {
		body, _ = sjson.Set(body, "subnetMask", data.SubnetMask.ValueString())
	}
	if !data.Gateway.IsNull() {
		body, _ = sjson.Set(body, "gateway", data.Gateway.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.IpInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "ipInterfaceName", data.IpInterfaceName.ValueString())
	}
	if !data.Domain.IsNull() {
		body, _ = sjson.Set(body, "svlConfig.domain", data.Domain.ValueInt64())
	}
	if len(data.SvlMembers) > 0 {
		body, _ = sjson.Set(body, "svlConfig.svlMembers", []interface{}{})
		for _, item := range data.SvlMembers {
			itemBody := ""
			if !item.SerialNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serialNumber", item.SerialNumber.ValueString())
			}
			if !item.Role.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "role", item.Role.ValueString())
			}
			if len(item.SvlLinks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "svlLinks", []interface{}{})
				for _, childItem := range item.SvlLinks {
					itemChildBody := ""
					if !childItem.LocalInterface.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "localInterface", childItem.LocalInterface.ValueString())
					}
					if !childItem.RemoteInterface.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "remoteInterface", childItem.RemoteInterface.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "svlLinks.-1", itemChildBody)
				}
			}
			if !item.LocalInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dadLink.localInterface", item.LocalInterface.ValueString())
			}
			if !item.RemoteInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dadLink.remoteInterface", item.RemoteInterface.ValueString())
			}
			body, _ = sjson.SetRaw(body, "svlConfig.svlMembers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PnPNetworkDeviceClaim) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("response.hostname"); value.Exists() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PnPNetworkDeviceClaim) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("response.hostname"); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PnPNetworkDeviceClaim) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceType.IsNull() {
		return false
	}
	if !data.SiteId.IsNull() {
		return false
	}
	if !data.Hostname.IsNull() {
		return false
	}
	if !data.ImageId.IsNull() {
		return false
	}
	if !data.RemoveInactive.IsNull() {
		return false
	}
	if !data.TemplateId.IsNull() {
		return false
	}
	if len(data.TemplateParameters) > 0 {
		return false
	}
	if !data.RfProfile.IsNull() {
		return false
	}
	if !data.TopOfStackSerialNumber.IsNull() {
		return false
	}
	if !data.CablingScheme.IsNull() {
		return false
	}
	if !data.SensorProfile.IsNull() {
		return false
	}
	if !data.StaticIpAddress.IsNull() {
		return false
	}
	if !data.SubnetMask.IsNull() {
		return false
	}
	if !data.Gateway.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.IpInterfaceName.IsNull() {
		return false
	}
	if !data.Domain.IsNull() {
		return false
	}
	if len(data.SvlMembers) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
