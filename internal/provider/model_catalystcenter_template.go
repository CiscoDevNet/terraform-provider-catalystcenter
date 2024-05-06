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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Template struct {
	Id              types.String             `tfsdk:"id"`
	ProjectId       types.String             `tfsdk:"project_id"`
	Name            types.String             `tfsdk:"name"`
	Description     types.String             `tfsdk:"description"`
	DeviceTypes     []TemplateDeviceTypes    `tfsdk:"device_types"`
	Language        types.String             `tfsdk:"language"`
	SoftwareType    types.String             `tfsdk:"software_type"`
	SoftwareVariant types.String             `tfsdk:"software_variant"`
	SoftwareVersion types.String             `tfsdk:"software_version"`
	TemplateContent types.String             `tfsdk:"template_content"`
	TemplateParams  []TemplateTemplateParams `tfsdk:"template_params"`
}

type TemplateDeviceTypes struct {
	ProductFamily types.String `tfsdk:"product_family"`
	ProductSeries types.String `tfsdk:"product_series"`
	ProductType   types.String `tfsdk:"product_type"`
}

type TemplateTemplateParams struct {
	Binding               types.String                   `tfsdk:"binding"`
	DataType              types.String                   `tfsdk:"data_type"`
	DefaultValue          types.String                   `tfsdk:"default_value"`
	Description           types.String                   `tfsdk:"description"`
	DisplayName           types.String                   `tfsdk:"display_name"`
	InstructionText       types.String                   `tfsdk:"instruction_text"`
	NotParam              types.Bool                     `tfsdk:"not_param"`
	ParamArray            types.Bool                     `tfsdk:"param_array"`
	ParameterName         types.String                   `tfsdk:"parameter_name"`
	Ranges                []TemplateTemplateParamsRanges `tfsdk:"ranges"`
	Required              types.Bool                     `tfsdk:"required"`
	DefaultSelectedValues types.Set                      `tfsdk:"default_selected_values"`
	SelectionType         types.String                   `tfsdk:"selection_type"`
	SelectionValues       types.Map                      `tfsdk:"selection_values"`
}

type TemplateTemplateParamsRanges struct {
	MaxValue types.Int64 `tfsdk:"max_value"`
	MinValue types.Int64 `tfsdk:"min_value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Template) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/template-programmer/project/%v/template", url.QueryEscape(data.ProjectId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data Template) getPathDelete() string {
	return "/dna/intent/api/v1/template-programmer/template"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Template) toBody(ctx context.Context, state Template) string {
	body := ""
	if state.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", state.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.DeviceTypes) > 0 {
		body, _ = sjson.Set(body, "deviceTypes", []interface{}{})
		for _, item := range data.DeviceTypes {
			itemBody := ""
			if !item.ProductFamily.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "productFamily", item.ProductFamily.ValueString())
			}
			if !item.ProductSeries.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "productSeries", item.ProductSeries.ValueString())
			}
			if !item.ProductType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "productType", item.ProductType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "deviceTypes.-1", itemBody)
		}
	}
	if !data.Language.IsNull() {
		body, _ = sjson.Set(body, "language", data.Language.ValueString())
	}
	if !data.SoftwareType.IsNull() {
		body, _ = sjson.Set(body, "softwareType", data.SoftwareType.ValueString())
	}
	if !data.SoftwareVariant.IsNull() {
		body, _ = sjson.Set(body, "softwareVariant", data.SoftwareVariant.ValueString())
	}
	if !data.SoftwareVersion.IsNull() {
		body, _ = sjson.Set(body, "softwareVersion", data.SoftwareVersion.ValueString())
	}
	if !data.TemplateContent.IsNull() {
		body, _ = sjson.Set(body, "templateContent", data.TemplateContent.ValueString())
	}
	if len(data.TemplateParams) > 0 {
		body, _ = sjson.Set(body, "templateParams", []interface{}{})
		for _, item := range data.TemplateParams {
			itemBody := ""
			if !item.Binding.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "binding", item.Binding.ValueString())
			}
			if !item.DataType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dataType", item.DataType.ValueString())
			}
			if !item.DefaultValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "defaultValue", item.DefaultValue.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.DisplayName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "displayName", item.DisplayName.ValueString())
			}
			if !item.InstructionText.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "instructionText", item.InstructionText.ValueString())
			}
			if !item.NotParam.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "notParam", item.NotParam.ValueBool())
			}
			if !item.ParamArray.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "paramArray", item.ParamArray.ValueBool())
			}
			if !item.ParameterName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parameterName", item.ParameterName.ValueString())
			}
			if len(item.Ranges) > 0 {
				itemBody, _ = sjson.Set(itemBody, "range", []interface{}{})
				for _, childItem := range item.Ranges {
					itemChildBody := ""
					if !childItem.MaxValue.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "maxValue", childItem.MaxValue.ValueInt64())
					}
					if !childItem.MinValue.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "minValue", childItem.MinValue.ValueInt64())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "range.-1", itemChildBody)
				}
			}
			if !item.Required.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "required", item.Required.ValueBool())
			}
			if !item.DefaultSelectedValues.IsNull() {
				var values []string
				item.DefaultSelectedValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "selection.defaultSelectedValues", values)
			}
			if !item.SelectionType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "selection.selectionType", item.SelectionType.ValueString())
			}
			if !item.SelectionValues.IsNull() {
				var values map[string]string
				item.SelectionValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "selection.selectionValues", values)
			}
			body, _ = sjson.SetRaw(body, "templateParams.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Template) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("deviceTypes"); value.Exists() {
		data.DeviceTypes = make([]TemplateDeviceTypes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TemplateDeviceTypes{}
			if cValue := v.Get("productFamily"); cValue.Exists() {
				item.ProductFamily = types.StringValue(cValue.String())
			} else {
				item.ProductFamily = types.StringNull()
			}
			if cValue := v.Get("productSeries"); cValue.Exists() {
				item.ProductSeries = types.StringValue(cValue.String())
			} else {
				item.ProductSeries = types.StringNull()
			}
			if cValue := v.Get("productType"); cValue.Exists() {
				item.ProductType = types.StringValue(cValue.String())
			} else {
				item.ProductType = types.StringNull()
			}
			data.DeviceTypes = append(data.DeviceTypes, item)
			return true
		})
	}
	if value := res.Get("language"); value.Exists() {
		data.Language = types.StringValue(value.String())
	} else {
		data.Language = types.StringNull()
	}
	if value := res.Get("softwareType"); value.Exists() {
		data.SoftwareType = types.StringValue(value.String())
	} else {
		data.SoftwareType = types.StringNull()
	}
	if value := res.Get("softwareVariant"); value.Exists() {
		data.SoftwareVariant = types.StringValue(value.String())
	} else {
		data.SoftwareVariant = types.StringNull()
	}
	if value := res.Get("softwareVersion"); value.Exists() {
		data.SoftwareVersion = types.StringValue(value.String())
	} else {
		data.SoftwareVersion = types.StringNull()
	}
	if value := res.Get("templateContent"); value.Exists() {
		data.TemplateContent = types.StringValue(value.String())
	} else {
		data.TemplateContent = types.StringNull()
	}
	if value := res.Get("templateParams"); value.Exists() {
		data.TemplateParams = make([]TemplateTemplateParams, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TemplateTemplateParams{}
			if cValue := v.Get("binding"); cValue.Exists() {
				item.Binding = types.StringValue(cValue.String())
			} else {
				item.Binding = types.StringNull()
			}
			if cValue := v.Get("dataType"); cValue.Exists() {
				item.DataType = types.StringValue(cValue.String())
			} else {
				item.DataType = types.StringNull()
			}
			if cValue := v.Get("defaultValue"); cValue.Exists() {
				item.DefaultValue = types.StringValue(cValue.String())
			} else {
				item.DefaultValue = types.StringNull()
			}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			} else {
				item.Description = types.StringNull()
			}
			if cValue := v.Get("displayName"); cValue.Exists() {
				item.DisplayName = types.StringValue(cValue.String())
			} else {
				item.DisplayName = types.StringNull()
			}
			if cValue := v.Get("instructionText"); cValue.Exists() {
				item.InstructionText = types.StringValue(cValue.String())
			} else {
				item.InstructionText = types.StringNull()
			}
			if cValue := v.Get("notParam"); cValue.Exists() {
				item.NotParam = types.BoolValue(cValue.Bool())
			} else {
				item.NotParam = types.BoolNull()
			}
			if cValue := v.Get("paramArray"); cValue.Exists() {
				item.ParamArray = types.BoolValue(cValue.Bool())
			} else {
				item.ParamArray = types.BoolNull()
			}
			if cValue := v.Get("parameterName"); cValue.Exists() {
				item.ParameterName = types.StringValue(cValue.String())
			} else {
				item.ParameterName = types.StringNull()
			}
			if cValue := v.Get("range"); cValue.Exists() {
				item.Ranges = make([]TemplateTemplateParamsRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TemplateTemplateParamsRanges{}
					if ccValue := cv.Get("maxValue"); ccValue.Exists() {
						cItem.MaxValue = types.Int64Value(ccValue.Int())
					} else {
						cItem.MaxValue = types.Int64Null()
					}
					if ccValue := cv.Get("minValue"); ccValue.Exists() {
						cItem.MinValue = types.Int64Value(ccValue.Int())
					} else {
						cItem.MinValue = types.Int64Null()
					}
					item.Ranges = append(item.Ranges, cItem)
					return true
				})
			}
			if cValue := v.Get("required"); cValue.Exists() {
				item.Required = types.BoolValue(cValue.Bool())
			} else {
				item.Required = types.BoolNull()
			}
			if cValue := v.Get("selection.defaultSelectedValues"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.DefaultSelectedValues = helpers.GetStringSet(cValue.Array())
			} else {
				item.DefaultSelectedValues = types.SetNull(types.StringType)
			}
			if cValue := v.Get("selection.selectionType"); cValue.Exists() {
				item.SelectionType = types.StringValue(cValue.String())
			} else {
				item.SelectionType = types.StringNull()
			}
			if cValue := v.Get("selection.selectionValues"); cValue.Exists() {
				item.SelectionValues = helpers.GetStringMap(cValue.Map())
			} else {
				item.SelectionValues = types.MapNull(types.StringType)
			}
			data.TemplateParams = append(data.TemplateParams, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Template) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.DeviceTypes {
		keys := [...]string{"productFamily", "productSeries", "productType"}
		keyValues := [...]string{data.DeviceTypes[i].ProductFamily.ValueString(), data.DeviceTypes[i].ProductSeries.ValueString(), data.DeviceTypes[i].ProductType.ValueString()}

		var r gjson.Result
		res.Get("deviceTypes").ForEach(
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
		if value := r.Get("productFamily"); value.Exists() && !data.DeviceTypes[i].ProductFamily.IsNull() {
			data.DeviceTypes[i].ProductFamily = types.StringValue(value.String())
		} else {
			data.DeviceTypes[i].ProductFamily = types.StringNull()
		}
		if value := r.Get("productSeries"); value.Exists() && !data.DeviceTypes[i].ProductSeries.IsNull() {
			data.DeviceTypes[i].ProductSeries = types.StringValue(value.String())
		} else {
			data.DeviceTypes[i].ProductSeries = types.StringNull()
		}
		if value := r.Get("productType"); value.Exists() && !data.DeviceTypes[i].ProductType.IsNull() {
			data.DeviceTypes[i].ProductType = types.StringValue(value.String())
		} else {
			data.DeviceTypes[i].ProductType = types.StringNull()
		}
	}
	if value := res.Get("language"); value.Exists() && !data.Language.IsNull() {
		data.Language = types.StringValue(value.String())
	} else {
		data.Language = types.StringNull()
	}
	if value := res.Get("softwareType"); value.Exists() && !data.SoftwareType.IsNull() {
		data.SoftwareType = types.StringValue(value.String())
	} else {
		data.SoftwareType = types.StringNull()
	}
	if value := res.Get("softwareVariant"); value.Exists() && !data.SoftwareVariant.IsNull() {
		data.SoftwareVariant = types.StringValue(value.String())
	} else {
		data.SoftwareVariant = types.StringNull()
	}
	if value := res.Get("softwareVersion"); value.Exists() && !data.SoftwareVersion.IsNull() {
		data.SoftwareVersion = types.StringValue(value.String())
	} else {
		data.SoftwareVersion = types.StringNull()
	}
	if value := res.Get("templateContent"); value.Exists() && !data.TemplateContent.IsNull() {
		data.TemplateContent = types.StringValue(value.String())
	} else {
		data.TemplateContent = types.StringNull()
	}
	for i := range data.TemplateParams {
		keys := [...]string{"binding", "dataType", "defaultValue", "description", "displayName", "instructionText", "notParam", "paramArray", "parameterName", "required", "selection.selectionType"}
		keyValues := [...]string{data.TemplateParams[i].Binding.ValueString(), data.TemplateParams[i].DataType.ValueString(), data.TemplateParams[i].DefaultValue.ValueString(), data.TemplateParams[i].Description.ValueString(), data.TemplateParams[i].DisplayName.ValueString(), data.TemplateParams[i].InstructionText.ValueString(), strconv.FormatBool(data.TemplateParams[i].NotParam.ValueBool()), strconv.FormatBool(data.TemplateParams[i].ParamArray.ValueBool()), data.TemplateParams[i].ParameterName.ValueString(), strconv.FormatBool(data.TemplateParams[i].Required.ValueBool()), data.TemplateParams[i].SelectionType.ValueString()}

		var r gjson.Result
		res.Get("templateParams").ForEach(
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
		if value := r.Get("binding"); value.Exists() && !data.TemplateParams[i].Binding.IsNull() {
			data.TemplateParams[i].Binding = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].Binding = types.StringNull()
		}
		if value := r.Get("dataType"); value.Exists() && !data.TemplateParams[i].DataType.IsNull() {
			data.TemplateParams[i].DataType = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].DataType = types.StringNull()
		}
		if value := r.Get("defaultValue"); value.Exists() && !data.TemplateParams[i].DefaultValue.IsNull() {
			data.TemplateParams[i].DefaultValue = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].DefaultValue = types.StringNull()
		}
		if value := r.Get("description"); value.Exists() && !data.TemplateParams[i].Description.IsNull() {
			data.TemplateParams[i].Description = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].Description = types.StringNull()
		}
		if value := r.Get("displayName"); value.Exists() && !data.TemplateParams[i].DisplayName.IsNull() {
			data.TemplateParams[i].DisplayName = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].DisplayName = types.StringNull()
		}
		if value := r.Get("instructionText"); value.Exists() && !data.TemplateParams[i].InstructionText.IsNull() {
			data.TemplateParams[i].InstructionText = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].InstructionText = types.StringNull()
		}
		if value := r.Get("notParam"); value.Exists() && !data.TemplateParams[i].NotParam.IsNull() {
			data.TemplateParams[i].NotParam = types.BoolValue(value.Bool())
		} else {
			data.TemplateParams[i].NotParam = types.BoolNull()
		}
		if value := r.Get("paramArray"); value.Exists() && !data.TemplateParams[i].ParamArray.IsNull() {
			data.TemplateParams[i].ParamArray = types.BoolValue(value.Bool())
		} else {
			data.TemplateParams[i].ParamArray = types.BoolNull()
		}
		if value := r.Get("parameterName"); value.Exists() && !data.TemplateParams[i].ParameterName.IsNull() {
			data.TemplateParams[i].ParameterName = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].ParameterName = types.StringNull()
		}
		for ci := range data.TemplateParams[i].Ranges {
			keys := [...]string{"maxValue", "minValue"}
			keyValues := [...]string{strconv.FormatInt(data.TemplateParams[i].Ranges[ci].MaxValue.ValueInt64(), 10), strconv.FormatInt(data.TemplateParams[i].Ranges[ci].MinValue.ValueInt64(), 10)}

			var cr gjson.Result
			r.Get("range").ForEach(
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
			if value := cr.Get("maxValue"); value.Exists() && !data.TemplateParams[i].Ranges[ci].MaxValue.IsNull() {
				data.TemplateParams[i].Ranges[ci].MaxValue = types.Int64Value(value.Int())
			} else {
				data.TemplateParams[i].Ranges[ci].MaxValue = types.Int64Null()
			}
			if value := cr.Get("minValue"); value.Exists() && !data.TemplateParams[i].Ranges[ci].MinValue.IsNull() {
				data.TemplateParams[i].Ranges[ci].MinValue = types.Int64Value(value.Int())
			} else {
				data.TemplateParams[i].Ranges[ci].MinValue = types.Int64Null()
			}
		}
		if value := r.Get("required"); value.Exists() && !data.TemplateParams[i].Required.IsNull() {
			data.TemplateParams[i].Required = types.BoolValue(value.Bool())
		} else {
			data.TemplateParams[i].Required = types.BoolNull()
		}
		if value := r.Get("selection.defaultSelectedValues"); value.Exists() && !data.TemplateParams[i].DefaultSelectedValues.IsNull() {
			data.TemplateParams[i].DefaultSelectedValues = helpers.GetStringSet(value.Array())
		} else {
			data.TemplateParams[i].DefaultSelectedValues = types.SetNull(types.StringType)
		}
		if value := r.Get("selection.selectionType"); value.Exists() && !data.TemplateParams[i].SelectionType.IsNull() {
			data.TemplateParams[i].SelectionType = types.StringValue(value.String())
		} else {
			data.TemplateParams[i].SelectionType = types.StringNull()
		}
		if value := r.Get("selection.selectionValues"); value.Exists() && !data.TemplateParams[i].SelectionValues.IsNull() {
			data.TemplateParams[i].SelectionValues = helpers.GetStringMap(value.Map())
		} else {
			data.TemplateParams[i].SelectionValues = types.MapNull(types.StringType)
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Template) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if len(data.DeviceTypes) > 0 {
		return false
	}
	if !data.Language.IsNull() {
		return false
	}
	if !data.SoftwareType.IsNull() {
		return false
	}
	if !data.SoftwareVariant.IsNull() {
		return false
	}
	if !data.SoftwareVersion.IsNull() {
		return false
	}
	if !data.TemplateContent.IsNull() {
		return false
	}
	if len(data.TemplateParams) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
