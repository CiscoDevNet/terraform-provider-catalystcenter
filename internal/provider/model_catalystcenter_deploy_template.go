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
type DeployTemplate struct {
	Id                           types.String                                 `tfsdk:"id"`
	TemplateId                   types.String                                 `tfsdk:"template_id"`
	Redeploy                     types.Bool                                   `tfsdk:"redeploy"`
	ForcePushTemplate            types.Bool                                   `tfsdk:"force_push_template"`
	IsComposite                  types.Bool                                   `tfsdk:"is_composite"`
	MainTemplateId               types.String                                 `tfsdk:"main_template_id"`
	MemberTemplateDeploymentInfo []DeployTemplateMemberTemplateDeploymentInfo `tfsdk:"member_template_deployment_info"`
	TargetInfo                   []DeployTemplateTargetInfo                   `tfsdk:"target_info"`
}

type DeployTemplateMemberTemplateDeploymentInfo struct {
	TemplateId        types.String                                           `tfsdk:"template_id"`
	ForcePushTemplate types.Bool                                             `tfsdk:"force_push_template"`
	IsComposite       types.Bool                                             `tfsdk:"is_composite"`
	CopyingConfig     types.Bool                                             `tfsdk:"copying_config"`
	MainTemplateId    types.String                                           `tfsdk:"main_template_id"`
	TargetInfo        []DeployTemplateMemberTemplateDeploymentInfoTargetInfo `tfsdk:"target_info"`
}

type DeployTemplateTargetInfo struct {
	HostName            types.String                             `tfsdk:"host_name"`
	Id                  types.String                             `tfsdk:"id"`
	Params              types.Map                                `tfsdk:"params"`
	ResourceParams      []DeployTemplateTargetInfoResourceParams `tfsdk:"resource_params"`
	Type                types.String                             `tfsdk:"type"`
	VersionedTemplateId types.String                             `tfsdk:"versioned_template_id"`
}

type DeployTemplateMemberTemplateDeploymentInfoTargetInfo struct {
	HostName            types.String                                                         `tfsdk:"host_name"`
	Id                  types.String                                                         `tfsdk:"id"`
	Params              types.Map                                                            `tfsdk:"params"`
	ResourceParams      []DeployTemplateMemberTemplateDeploymentInfoTargetInfoResourceParams `tfsdk:"resource_params"`
	Type                types.String                                                         `tfsdk:"type"`
	VersionedTemplateId types.String                                                         `tfsdk:"versioned_template_id"`
}

type DeployTemplateTargetInfoResourceParams struct {
	Type  types.String `tfsdk:"type"`
	Scope types.String `tfsdk:"scope"`
	Value types.String `tfsdk:"value"`
}

type DeployTemplateMemberTemplateDeploymentInfoTargetInfoResourceParams struct {
	Type  types.String `tfsdk:"type"`
	Scope types.String `tfsdk:"scope"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DeployTemplate) getPath() string {
	return "/dna/intent/api/v2/template-programmer/template/deploy"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DeployTemplate) toBody(ctx context.Context, state DeployTemplate) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.TemplateId.IsNull() {
		body, _ = sjson.Set(body, "templateId", data.TemplateId.ValueString())
	}
	if !data.Redeploy.IsNull() {
		body, _ = sjson.Set(body, "", data.Redeploy.ValueBool())
	}
	if !data.ForcePushTemplate.IsNull() {
		body, _ = sjson.Set(body, "forcePushTemplate", data.ForcePushTemplate.ValueBool())
	}
	if !data.IsComposite.IsNull() {
		body, _ = sjson.Set(body, "isComposite", data.IsComposite.ValueBool())
	}
	if !data.MainTemplateId.IsNull() {
		body, _ = sjson.Set(body, "mainTemplateId", data.MainTemplateId.ValueString())
	}
	if len(data.MemberTemplateDeploymentInfo) > 0 {
		body, _ = sjson.Set(body, "memberTemplateDeploymentInfo", []interface{}{})
		for _, item := range data.MemberTemplateDeploymentInfo {
			itemBody := ""
			if !item.TemplateId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "templateId", item.TemplateId.ValueString())
			}
			if !item.ForcePushTemplate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "forcePushTemplate", item.ForcePushTemplate.ValueBool())
			}
			if !item.IsComposite.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isComposite", item.IsComposite.ValueBool())
			}
			if !item.CopyingConfig.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "copyingConfig", item.CopyingConfig.ValueBool())
			}
			if !item.MainTemplateId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mainTemplateId", item.MainTemplateId.ValueString())
			}
			if len(item.TargetInfo) > 0 {
				itemBody, _ = sjson.Set(itemBody, "targetInfo", []interface{}{})
				for _, childItem := range item.TargetInfo {
					itemChildBody := ""
					if !childItem.HostName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "hostName", childItem.HostName.ValueString())
					}
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Params.IsNull() {
						var values map[string]string
						childItem.Params.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "params", values)
					}
					if len(childItem.ResourceParams) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "resourceParams", []interface{}{})
						for _, childChildItem := range childItem.ResourceParams {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "type", childChildItem.Type.ValueString())
							}
							if !childChildItem.Scope.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "scope", childChildItem.Scope.ValueString())
							}
							if !childChildItem.Value.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.Value.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "resourceParams.-1", itemChildChildBody)
						}
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.VersionedTemplateId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "versionedTemplateId", childItem.VersionedTemplateId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "targetInfo.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "memberTemplateDeploymentInfo.-1", itemBody)
		}
	}
	if len(data.TargetInfo) > 0 {
		body, _ = sjson.Set(body, "targetInfo", []interface{}{})
		for _, item := range data.TargetInfo {
			itemBody := ""
			if !item.HostName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hostName", item.HostName.ValueString())
			}
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Params.IsNull() {
				var values map[string]string
				item.Params.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "params", values)
			}
			if len(item.ResourceParams) > 0 {
				itemBody, _ = sjson.Set(itemBody, "resourceParams", []interface{}{})
				for _, childItem := range item.ResourceParams {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Scope.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "scope", childItem.Scope.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "resourceParams.-1", itemChildBody)
				}
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.VersionedTemplateId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "versionedTemplateId", item.VersionedTemplateId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "targetInfo.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DeployTemplate) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("templateId"); value.Exists() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
	if value := res.Get(""); value.Exists() {
		data.Redeploy = types.BoolValue(value.Bool())
	} else {
		data.Redeploy = types.BoolNull()
	}
	if value := res.Get("forcePushTemplate"); value.Exists() {
		data.ForcePushTemplate = types.BoolValue(value.Bool())
	} else {
		data.ForcePushTemplate = types.BoolNull()
	}
	if value := res.Get("isComposite"); value.Exists() {
		data.IsComposite = types.BoolValue(value.Bool())
	} else {
		data.IsComposite = types.BoolNull()
	}
	if value := res.Get("mainTemplateId"); value.Exists() {
		data.MainTemplateId = types.StringValue(value.String())
	} else {
		data.MainTemplateId = types.StringNull()
	}
	if value := res.Get("memberTemplateDeploymentInfo"); value.Exists() && len(value.Array()) > 0 {
		data.MemberTemplateDeploymentInfo = make([]DeployTemplateMemberTemplateDeploymentInfo, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeployTemplateMemberTemplateDeploymentInfo{}
			if cValue := v.Get("templateId"); cValue.Exists() {
				item.TemplateId = types.StringValue(cValue.String())
			} else {
				item.TemplateId = types.StringNull()
			}
			if cValue := v.Get("forcePushTemplate"); cValue.Exists() {
				item.ForcePushTemplate = types.BoolValue(cValue.Bool())
			} else {
				item.ForcePushTemplate = types.BoolNull()
			}
			if cValue := v.Get("isComposite"); cValue.Exists() {
				item.IsComposite = types.BoolValue(cValue.Bool())
			} else {
				item.IsComposite = types.BoolNull()
			}
			if cValue := v.Get("copyingConfig"); cValue.Exists() {
				item.CopyingConfig = types.BoolValue(cValue.Bool())
			} else {
				item.CopyingConfig = types.BoolNull()
			}
			if cValue := v.Get("mainTemplateId"); cValue.Exists() {
				item.MainTemplateId = types.StringValue(cValue.String())
			} else {
				item.MainTemplateId = types.StringNull()
			}
			if cValue := v.Get("targetInfo"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.TargetInfo = make([]DeployTemplateMemberTemplateDeploymentInfoTargetInfo, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := DeployTemplateMemberTemplateDeploymentInfoTargetInfo{}
					if ccValue := cv.Get("hostName"); ccValue.Exists() {
						cItem.HostName = types.StringValue(ccValue.String())
					} else {
						cItem.HostName = types.StringNull()
					}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("params"); ccValue.Exists() && len(ccValue.Map()) > 0 {
						cItem.Params = helpers.GetStringMap(ccValue.Map())
					} else {
						cItem.Params = types.MapNull(types.StringType)
					}
					if ccValue := cv.Get("resourceParams"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.ResourceParams = make([]DeployTemplateMemberTemplateDeploymentInfoTargetInfoResourceParams, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := DeployTemplateMemberTemplateDeploymentInfoTargetInfoResourceParams{}
							if cccValue := ccv.Get("type"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							if cccValue := ccv.Get("scope"); cccValue.Exists() {
								ccItem.Scope = types.StringValue(cccValue.String())
							} else {
								ccItem.Scope = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() {
								ccItem.Value = types.StringValue(cccValue.String())
							} else {
								ccItem.Value = types.StringNull()
							}
							cItem.ResourceParams = append(cItem.ResourceParams, ccItem)
							return true
						})
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("versionedTemplateId"); ccValue.Exists() {
						cItem.VersionedTemplateId = types.StringValue(ccValue.String())
					} else {
						cItem.VersionedTemplateId = types.StringNull()
					}
					item.TargetInfo = append(item.TargetInfo, cItem)
					return true
				})
			}
			data.MemberTemplateDeploymentInfo = append(data.MemberTemplateDeploymentInfo, item)
			return true
		})
	}
	if value := res.Get("targetInfo"); value.Exists() && len(value.Array()) > 0 {
		data.TargetInfo = make([]DeployTemplateTargetInfo, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeployTemplateTargetInfo{}
			if cValue := v.Get("hostName"); cValue.Exists() {
				item.HostName = types.StringValue(cValue.String())
			} else {
				item.HostName = types.StringNull()
			}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("params"); cValue.Exists() && len(cValue.Map()) > 0 {
				item.Params = helpers.GetStringMap(cValue.Map())
			} else {
				item.Params = types.MapNull(types.StringType)
			}
			if cValue := v.Get("resourceParams"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ResourceParams = make([]DeployTemplateTargetInfoResourceParams, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := DeployTemplateTargetInfoResourceParams{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("scope"); ccValue.Exists() {
						cItem.Scope = types.StringValue(ccValue.String())
					} else {
						cItem.Scope = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.ResourceParams = append(item.ResourceParams, cItem)
					return true
				})
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("versionedTemplateId"); cValue.Exists() {
				item.VersionedTemplateId = types.StringValue(cValue.String())
			} else {
				item.VersionedTemplateId = types.StringNull()
			}
			data.TargetInfo = append(data.TargetInfo, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DeployTemplate) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("templateId"); value.Exists() && !data.TemplateId.IsNull() {
		data.TemplateId = types.StringValue(value.String())
	} else {
		data.TemplateId = types.StringNull()
	}
	if value := res.Get(""); value.Exists() && !data.Redeploy.IsNull() {
		data.Redeploy = types.BoolValue(value.Bool())
	} else {
		data.Redeploy = types.BoolNull()
	}
	if value := res.Get("forcePushTemplate"); value.Exists() && !data.ForcePushTemplate.IsNull() {
		data.ForcePushTemplate = types.BoolValue(value.Bool())
	} else {
		data.ForcePushTemplate = types.BoolNull()
	}
	if value := res.Get("isComposite"); value.Exists() && !data.IsComposite.IsNull() {
		data.IsComposite = types.BoolValue(value.Bool())
	} else {
		data.IsComposite = types.BoolNull()
	}
	if value := res.Get("mainTemplateId"); value.Exists() && !data.MainTemplateId.IsNull() {
		data.MainTemplateId = types.StringValue(value.String())
	} else {
		data.MainTemplateId = types.StringNull()
	}
	for i := range data.MemberTemplateDeploymentInfo {
		keys := [...]string{"templateId"}
		keyValues := [...]string{data.MemberTemplateDeploymentInfo[i].TemplateId.ValueString()}

		var r gjson.Result
		res.Get("memberTemplateDeploymentInfo").ForEach(
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
		if value := r.Get("templateId"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TemplateId.IsNull() {
			data.MemberTemplateDeploymentInfo[i].TemplateId = types.StringValue(value.String())
		} else {
			data.MemberTemplateDeploymentInfo[i].TemplateId = types.StringNull()
		}
		if value := r.Get("forcePushTemplate"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].ForcePushTemplate.IsNull() {
			data.MemberTemplateDeploymentInfo[i].ForcePushTemplate = types.BoolValue(value.Bool())
		} else {
			data.MemberTemplateDeploymentInfo[i].ForcePushTemplate = types.BoolNull()
		}
		if value := r.Get("isComposite"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].IsComposite.IsNull() {
			data.MemberTemplateDeploymentInfo[i].IsComposite = types.BoolValue(value.Bool())
		} else {
			data.MemberTemplateDeploymentInfo[i].IsComposite = types.BoolNull()
		}
		if value := r.Get("copyingConfig"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].CopyingConfig.IsNull() {
			data.MemberTemplateDeploymentInfo[i].CopyingConfig = types.BoolValue(value.Bool())
		} else {
			data.MemberTemplateDeploymentInfo[i].CopyingConfig = types.BoolNull()
		}
		if value := r.Get("mainTemplateId"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].MainTemplateId.IsNull() {
			data.MemberTemplateDeploymentInfo[i].MainTemplateId = types.StringValue(value.String())
		} else {
			data.MemberTemplateDeploymentInfo[i].MainTemplateId = types.StringNull()
		}
		for ci := range data.MemberTemplateDeploymentInfo[i].TargetInfo {
			keys := [...]string{"hostName", "id", "type", "versionedTemplateId"}
			keyValues := [...]string{data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].HostName.ValueString(), data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Id.ValueString(), data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Type.ValueString(), data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].VersionedTemplateId.ValueString()}

			var cr gjson.Result
			r.Get("targetInfo").ForEach(
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
			if value := cr.Get("hostName"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].HostName.IsNull() {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].HostName = types.StringValue(value.String())
			} else {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].HostName = types.StringNull()
			}
			if value := cr.Get("id"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Id.IsNull() {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Id = types.StringValue(value.String())
			} else {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Id = types.StringNull()
			}
			if value := cr.Get("params"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Params.IsNull() {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Params = helpers.GetStringMap(value.Map())
			} else {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Params = types.MapNull(types.StringType)
			}
			for cci := range data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams {
				keys := [...]string{"type", "scope", "value"}
				keyValues := [...]string{data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Type.ValueString(), data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Scope.ValueString(), data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Value.ValueString()}

				var ccr gjson.Result
				cr.Get("resourceParams").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				if value := ccr.Get("type"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Type.IsNull() {
					data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Type = types.StringValue(value.String())
				} else {
					data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Type = types.StringNull()
				}
				if value := ccr.Get("scope"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Scope.IsNull() {
					data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Scope = types.StringValue(value.String())
				} else {
					data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Scope = types.StringNull()
				}
				if value := ccr.Get("value"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Value.IsNull() {
					data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Value = types.StringValue(value.String())
				} else {
					data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].ResourceParams[cci].Value = types.StringNull()
				}
			}
			if value := cr.Get("type"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Type.IsNull() {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Type = types.StringValue(value.String())
			} else {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].Type = types.StringNull()
			}
			if value := cr.Get("versionedTemplateId"); value.Exists() && !data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].VersionedTemplateId.IsNull() {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].VersionedTemplateId = types.StringValue(value.String())
			} else {
				data.MemberTemplateDeploymentInfo[i].TargetInfo[ci].VersionedTemplateId = types.StringNull()
			}
		}
	}
	for i := range data.TargetInfo {
		keys := [...]string{"hostName", "id", "type", "versionedTemplateId"}
		keyValues := [...]string{data.TargetInfo[i].HostName.ValueString(), data.TargetInfo[i].Id.ValueString(), data.TargetInfo[i].Type.ValueString(), data.TargetInfo[i].VersionedTemplateId.ValueString()}

		var r gjson.Result
		res.Get("targetInfo").ForEach(
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
		if value := r.Get("hostName"); value.Exists() && !data.TargetInfo[i].HostName.IsNull() {
			data.TargetInfo[i].HostName = types.StringValue(value.String())
		} else {
			data.TargetInfo[i].HostName = types.StringNull()
		}
		if value := r.Get("id"); value.Exists() && !data.TargetInfo[i].Id.IsNull() {
			data.TargetInfo[i].Id = types.StringValue(value.String())
		} else {
			data.TargetInfo[i].Id = types.StringNull()
		}
		if value := r.Get("params"); value.Exists() && !data.TargetInfo[i].Params.IsNull() {
			data.TargetInfo[i].Params = helpers.GetStringMap(value.Map())
		} else {
			data.TargetInfo[i].Params = types.MapNull(types.StringType)
		}
		for ci := range data.TargetInfo[i].ResourceParams {
			keys := [...]string{"type", "scope", "value"}
			keyValues := [...]string{data.TargetInfo[i].ResourceParams[ci].Type.ValueString(), data.TargetInfo[i].ResourceParams[ci].Scope.ValueString(), data.TargetInfo[i].ResourceParams[ci].Value.ValueString()}

			var cr gjson.Result
			r.Get("resourceParams").ForEach(
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
			if value := cr.Get("type"); value.Exists() && !data.TargetInfo[i].ResourceParams[ci].Type.IsNull() {
				data.TargetInfo[i].ResourceParams[ci].Type = types.StringValue(value.String())
			} else {
				data.TargetInfo[i].ResourceParams[ci].Type = types.StringNull()
			}
			if value := cr.Get("scope"); value.Exists() && !data.TargetInfo[i].ResourceParams[ci].Scope.IsNull() {
				data.TargetInfo[i].ResourceParams[ci].Scope = types.StringValue(value.String())
			} else {
				data.TargetInfo[i].ResourceParams[ci].Scope = types.StringNull()
			}
			if value := cr.Get("value"); value.Exists() && !data.TargetInfo[i].ResourceParams[ci].Value.IsNull() {
				data.TargetInfo[i].ResourceParams[ci].Value = types.StringValue(value.String())
			} else {
				data.TargetInfo[i].ResourceParams[ci].Value = types.StringNull()
			}
		}
		if value := r.Get("type"); value.Exists() && !data.TargetInfo[i].Type.IsNull() {
			data.TargetInfo[i].Type = types.StringValue(value.String())
		} else {
			data.TargetInfo[i].Type = types.StringNull()
		}
		if value := r.Get("versionedTemplateId"); value.Exists() && !data.TargetInfo[i].VersionedTemplateId.IsNull() {
			data.TargetInfo[i].VersionedTemplateId = types.StringValue(value.String())
		} else {
			data.TargetInfo[i].VersionedTemplateId = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DeployTemplate) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Redeploy.IsNull() {
		return false
	}
	if !data.ForcePushTemplate.IsNull() {
		return false
	}
	if !data.IsComposite.IsNull() {
		return false
	}
	if !data.MainTemplateId.IsNull() {
		return false
	}
	if len(data.MemberTemplateDeploymentInfo) > 0 {
		return false
	}
	if len(data.TargetInfo) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
