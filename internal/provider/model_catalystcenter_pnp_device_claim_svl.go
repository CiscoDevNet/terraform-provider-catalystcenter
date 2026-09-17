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
type PnPDeviceClaimSVL struct {
	Id                 types.String                          `tfsdk:"id"`
	DeviceId           types.String                          `tfsdk:"device_id"`
	SiteId             types.String                          `tfsdk:"site_id"`
	Hostname           types.String                          `tfsdk:"hostname"`
	ImageId            types.String                          `tfsdk:"image_id"`
	RemoveInactive     types.Bool                            `tfsdk:"remove_inactive"`
	TemplateId         types.String                          `tfsdk:"template_id"`
	TemplateParameters []PnPDeviceClaimSVLTemplateParameters `tfsdk:"template_parameters"`
	Domain             types.Int64                           `tfsdk:"domain"`
	SvlMembers         []PnPDeviceClaimSVLSvlMembers         `tfsdk:"svl_members"`
}

type PnPDeviceClaimSVLTemplateParameters struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type PnPDeviceClaimSVLSvlMembers struct {
	SerialNumber    types.String                          `tfsdk:"serial_number"`
	Role            types.String                          `tfsdk:"role"`
	SvlLinks        []PnPDeviceClaimSVLSvlMembersSvlLinks `tfsdk:"svl_links"`
	LocalInterface  types.String                          `tfsdk:"local_interface"`
	RemoteInterface types.String                          `tfsdk:"remote_interface"`
}

type PnPDeviceClaimSVLSvlMembersSvlLinks struct {
	LocalInterface  types.String `tfsdk:"local_interface"`
	RemoteInterface types.String `tfsdk:"remote_interface"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PnPDeviceClaimSVL) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/pnpNetworkDevices/%v/claim", url.QueryEscape(data.DeviceId.ValueString()))
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
func (data PnPDeviceClaimSVL) toBody(ctx context.Context, state PnPDeviceClaimSVL) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	body, _ = sjson.Set(body, "deviceType", "SVL")
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
func (data *PnPDeviceClaimSVL) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("hostname"); value.Exists() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get("imageInfo.imageId"); value.Exists() {
		data.ImageId = types.StringValue(value.String())
	} else {
		data.ImageId = types.StringNull()
	}
	if value := res.Get("imageInfo.removeInactive"); value.Exists() {
		data.RemoveInactive = types.BoolValue(value.Bool())
	} else {
		data.RemoveInactive = types.BoolNull()
	}
	if value := res.Get("templateInfo.templateId"); value.Exists() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
	if value := res.Get("templateInfo.templateParameters"); value.Exists() && len(value.Array()) > 0 {
		data.TemplateParameters = make([]PnPDeviceClaimSVLTemplateParameters, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PnPDeviceClaimSVLTemplateParameters{}
			if cValue := v.Get("key"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("value"); cValue.Exists() {
				item.Value = types.StringValue(cValue.String())
			} else {
				item.Value = types.StringNull()
			}
			data.TemplateParameters = append(data.TemplateParameters, item)
			return true
		})
	}
	if value := res.Get("svlConfig.domain"); value.Exists() {
		data.Domain = types.Int64Value(value.Int())
	} else {
		data.Domain = types.Int64Null()
	}
	if value := res.Get("svlConfig.svlMembers"); value.Exists() && len(value.Array()) > 0 {
		data.SvlMembers = make([]PnPDeviceClaimSVLSvlMembers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PnPDeviceClaimSVLSvlMembers{}
			if cValue := v.Get("serialNumber"); cValue.Exists() {
				item.SerialNumber = types.StringValue(cValue.String())
			} else {
				item.SerialNumber = types.StringNull()
			}
			if cValue := v.Get("role"); cValue.Exists() {
				item.Role = types.StringValue(cValue.String())
			} else {
				item.Role = types.StringNull()
			}
			if cValue := v.Get("svlLinks"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.SvlLinks = make([]PnPDeviceClaimSVLSvlMembersSvlLinks, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PnPDeviceClaimSVLSvlMembersSvlLinks{}
					if ccValue := cv.Get("localInterface"); ccValue.Exists() {
						cItem.LocalInterface = types.StringValue(ccValue.String())
					} else {
						cItem.LocalInterface = types.StringNull()
					}
					if ccValue := cv.Get("remoteInterface"); ccValue.Exists() {
						cItem.RemoteInterface = types.StringValue(ccValue.String())
					} else {
						cItem.RemoteInterface = types.StringNull()
					}
					item.SvlLinks = append(item.SvlLinks, cItem)
					return true
				})
			}
			if cValue := v.Get("dadLink.localInterface"); cValue.Exists() {
				item.LocalInterface = types.StringValue(cValue.String())
			} else {
				item.LocalInterface = types.StringNull()
			}
			if cValue := v.Get("dadLink.remoteInterface"); cValue.Exists() {
				item.RemoteInterface = types.StringValue(cValue.String())
			} else {
				item.RemoteInterface = types.StringNull()
			}
			data.SvlMembers = append(data.SvlMembers, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PnPDeviceClaimSVL) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
	if value := res.Get("hostname"); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get("imageInfo.imageId"); value.Exists() && !data.ImageId.IsNull() {
		data.ImageId = types.StringValue(value.String())
	} else {
		data.ImageId = types.StringNull()
	}
	if value := res.Get("imageInfo.removeInactive"); value.Exists() && !data.RemoveInactive.IsNull() {
		data.RemoveInactive = types.BoolValue(value.Bool())
	} else {
		data.RemoveInactive = types.BoolNull()
	}
	if value := res.Get("templateInfo.templateId"); value.Exists() && !data.TemplateId.IsNull() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
	for i := range data.TemplateParameters {
		keys := [...]string{"key"}
		keyValues := [...]string{data.TemplateParameters[i].Name.ValueString()}

		var r gjson.Result
		res.Get("templateInfo.templateParameters").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("key"); value.Exists() && !data.TemplateParameters[i].Name.IsNull() {
			data.TemplateParameters[i].Name = types.StringValue(value.String())
		} else {
			data.TemplateParameters[i].Name = types.StringNull()
		}
		if value := r.Get("value"); value.Exists() && !data.TemplateParameters[i].Value.IsNull() {
			data.TemplateParameters[i].Value = types.StringValue(value.String())
		} else {
			data.TemplateParameters[i].Value = types.StringNull()
		}
	}
	if value := res.Get("svlConfig.domain"); value.Exists() && !data.Domain.IsNull() {
		data.Domain = types.Int64Value(value.Int())
	} else {
		data.Domain = types.Int64Null()
	}
	for i := range data.SvlMembers {
		keys := [...]string{"serialNumber"}
		keyValues := [...]string{data.SvlMembers[i].SerialNumber.ValueString()}

		var r gjson.Result
		res.Get("svlConfig.svlMembers").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("serialNumber"); value.Exists() && !data.SvlMembers[i].SerialNumber.IsNull() {
			data.SvlMembers[i].SerialNumber = types.StringValue(value.String())
		} else {
			data.SvlMembers[i].SerialNumber = types.StringNull()
		}
		if value := r.Get("role"); value.Exists() && !data.SvlMembers[i].Role.IsNull() {
			data.SvlMembers[i].Role = types.StringValue(value.String())
		} else {
			data.SvlMembers[i].Role = types.StringNull()
		}
		for ci := range data.SvlMembers[i].SvlLinks {
			keys := [...]string{"localInterface"}
			keyValues := [...]string{data.SvlMembers[i].SvlLinks[ci].LocalInterface.ValueString()}

			var cr gjson.Result
			r.Get("svlLinks").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("localInterface"); value.Exists() && !data.SvlMembers[i].SvlLinks[ci].LocalInterface.IsNull() {
				data.SvlMembers[i].SvlLinks[ci].LocalInterface = types.StringValue(value.String())
			} else {
				data.SvlMembers[i].SvlLinks[ci].LocalInterface = types.StringNull()
			}
			if value := cr.Get("remoteInterface"); value.Exists() && !data.SvlMembers[i].SvlLinks[ci].RemoteInterface.IsNull() {
				data.SvlMembers[i].SvlLinks[ci].RemoteInterface = types.StringValue(value.String())
			} else {
				data.SvlMembers[i].SvlLinks[ci].RemoteInterface = types.StringNull()
			}
		}
		if value := r.Get("dadLink.localInterface"); value.Exists() && !data.SvlMembers[i].LocalInterface.IsNull() {
			data.SvlMembers[i].LocalInterface = types.StringValue(value.String())
		} else {
			data.SvlMembers[i].LocalInterface = types.StringNull()
		}
		if value := r.Get("dadLink.remoteInterface"); value.Exists() && !data.SvlMembers[i].RemoteInterface.IsNull() {
			data.SvlMembers[i].RemoteInterface = types.StringValue(value.String())
		} else {
			data.SvlMembers[i].RemoteInterface = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PnPDeviceClaimSVL) isNull(ctx context.Context, res gjson.Result) bool {
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
	if !data.Domain.IsNull() {
		return false
	}
	if len(data.SvlMembers) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
