//go:build ignore
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}
// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data {{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}
// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete
{{if .DeleteRestEndpoint}}
func (data {{camelCase .Name}}) getPathDelete() string {
	{{- if and (hasReference .Attributes) (strContains .DeleteRestEndpoint "%v")}}
		return fmt.Sprintf("{{.DeleteRestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.DeleteRestEndpoint}}"
	{{- end}}
}
{{end}}
// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data {{camelCase .Name}}) toBody(ctx context.Context, state {{camelCase .Name}}) string {
	{{- if .RootList}}
	body := "[]"
	{{- else}}
	body := ""
	{{- end}}
	put := false
	if state.Id.ValueString() != "" {
		put = true
		{{- if .PutIdIncludePath}}
		body, _ = sjson.Set(body, "{{ .PutIdIncludePath}}", state.Id.ValueString())
		{{- end}}
	}
	_ = put
	{{- range .Attributes}}
	{{- if .Value}}
	{{- if .ValueCondition}}
	if {{.ValueCondition}} {
		body, _ = sjson.Set(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	}
	{{- else}}
	body, _ = sjson.Set(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- end}}
	{{- else if and (not .Reference) (not .CreateQueryPath) (not .QueryParamNoBody)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !data.{{toGoName .TfName}}.IsNull() {{if .ExcludeFromPut}}&& put == false{{end}} {
		body, _ = sjson.Set(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", data.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if isListSet .}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if eq .Type "Map"}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values map[string]string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if isNestedListSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			{{- if .ValueCondition}}
			if {{.ValueCondition}} {
				itemBody, _ = sjson.Set(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			}
			{{- else}}
			itemBody, _ = sjson.Set(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			{{- end}}
			{{- else if not .Reference}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if !item.{{toGoName .TfName}}.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", item.{{toGoName .TfName}}.Value{{.Type}}())
			}
			{{- else if isListSet .}}
			if !item.{{toGoName .TfName}}.IsNull() {
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if eq .Type "Map"}}
			if !item.{{toGoName .TfName}}.IsNull() {
				var values map[string]string
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if isNestedListSet .}}
			if len(item.{{toGoName .TfName}}) > 0 {
				itemBody, _ = sjson.Set(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					{{- if .ValueCondition}}
					if {{.ValueCondition}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					}
					{{- else}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					{{- end}}
					{{- else if not .Reference}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", childItem.{{toGoName .TfName}}.Value{{.Type}}())
					}
					{{- else if isListSet .}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if eq .Type "Map"}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						var values map[string]string
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if isNestedListSet .}}
					if len(childItem.{{toGoName .TfName}}) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							{{- if .ValueCondition}}
							if {{.ValueCondition}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							}
							{{- else}}
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							{{- end}}
							{{- else if not .Reference}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
							}
							{{- else if isListSet .}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
							}
							{{- else if eq .Type "Map"}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								var values map[string]string
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, "{{if .DataPath}}{{.DataPath}}.{{end}}{{if .ModelName}}{{.ModelName}}.{{end}}-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return body
}
// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	{{- if .RootList}}
	{{- range .Attributes}}
	{{if .ResponseDataPath}}
	res = res.Get("{{.ResponseDataPath}}")
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if .DataSourceNoId}}
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("{{if .IdFromQueryPath}}{{if eq .IdFromQueryPath "." }}{{else}}{{.IdFromQueryPath}}.{{end}}{{if .IdFromQueryPathAttribute}}{{.IdFromQueryPathAttribute}}{{else}}id{{end}}{{end}}"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	{{- end}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference) (not .CreateQueryPath) (not .QueryParamNoBody)}}
	{{- $cname := toGoName .TfName}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {
		{{- if .DefaultValue}}
		data.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
		{{- else}}
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
		{{- end}}
	}
	{{- else if isListSet .}}
	if value := res.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"); value.Exists() && len(value.Array()) > 0 {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if eq .Type "Map"}}
	if value := res.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"); value.Exists() && len(value.Map()) > 0 {
		data.{{toGoName .TfName}} = helpers.GetStringMap(value.Map())
	} else {
		data.{{toGoName .TfName}} = types.MapNull(types.StringType)
	}
	{{- else if isNestedListSet .}}
	if value := res{{if .ModelName}}.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"){{end}}; value.Exists() && len(value.Array()) > 0 {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if cValue := v.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.{{.Type}}Value(cValue.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {
				{{- if .DefaultValue}}
				item.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
				{{- else}}
				item.{{toGoName .TfName}} = types.{{.Type}}Null()
				{{- end}}
			}
			{{- else if isListSet .}}
			if cValue := v.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if eq .Type "Map"}}
			if cValue := v.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cValue.Exists() && len(cValue.Map()) > 0 {
				item.{{toGoName .TfName}} = helpers.GetStringMap(cValue.Map())
			} else {
				item.{{toGoName .TfName}} = types.MapNull(types.StringType)
			}
			{{- else if isNestedListSet .}}
			if cValue := v.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if ccValue := cv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.{{.Type}}Value(ccValue.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
					} else {
						{{- if .DefaultValue}}
						cItem.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
						{{- else}}
						cItem.{{toGoName .TfName}} = types.{{.Type}}Null()
						{{- end}}
					}
					{{- else if isListSet .}}
					if ccValue := cv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
					}
					{{- else if eq .Type "Map"}}
					if ccValue := cv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); ccValue.Exists() && len(ccValue.Map()) > 0 {
						cItem.{{toGoName .TfName}} = helpers.GetStringMap(ccValue.Map())
					} else {
						cItem.{{toGoName .TfName}} = types.MapNull(types.StringType)
					}
					{{- else if isNestedListSet .}}
					if ccValue := cv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if cccValue := ccv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cccValue.Exists() {
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Value(cccValue.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
							} else {
								{{- if .DefaultValue}}
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
								{{- else}}
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Null()
								{{- end}}
							}
							{{- else if isListSet .}}
							if cccValue := ccv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cccValue.Exists() && len(cccValue.Array()) > 0 {
								ccItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cccValue.Array())
							} else {
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
							}
							{{- else if eq .Type "Map"}}
							if cccValue := ccv.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); cccValue.Exists() && len(cccValue.Map()) > 0 {
								ccItem.{{toGoName .TfName}} = helpers.GetStringMap(cccValue.Map())
							} else {
								ccItem.{{toGoName .TfName}} = types.MapNull(types.StringType)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}
// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *{{camelCase .Name}}) updateFromBody(ctx context.Context, res gjson.Result) {
	{{- if .RootList}}
	{{- range .Attributes}}
	{{if .ResponseDataPath}}
	res = res.Get("{{.ResponseDataPath}}")
	{{- end}}
	{{- end}}
	{{- end}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference) (not .CreateQueryPath) (not .QueryParamNoBody)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {{if .DefaultValue}}if data.{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
	}
	{{- else if isListSet .}}
	if value := res.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if eq .Type "Map"}}
	if value := res.Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = helpers.GetStringMap(value.Map())
	} else {
		data.{{toGoName .TfName}} = types.MapNull(types.StringType)
	}
	{{- else if isNestedListSet .}}
	{{- $list := (toGoName .TfName)}}
	for i := range data.{{toGoName .TfName}} {
		keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
		keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

		var r gjson.Result
		res.{{if .ModelName}}Get("{{if .ResponseDataPath}}{{.ResponseDataPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}").{{end}}ForEach(
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

		{{- range .Attributes}}
		{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
		if value := r.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
		} else {{if .DefaultValue}}if data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
			data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Null()
		}
		{{- else if isListSet .}}
		if value := r.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
		}
		{{- else if eq .Type "Map"}}
		if value := r.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{toGoName .TfName}}.IsNull() {
			data.{{$list}}[i].{{toGoName .TfName}} = helpers.GetStringMap(value.Map())
		} else {
			data.{{$list}}[i].{{toGoName .TfName}} = types.MapNull(types.StringType)
		}
		{{- else if isNestedListSet .}}
		{{- $clist := (toGoName .TfName)}}
		for ci := range data.{{$list}}[i].{{toGoName .TfName}} {
			keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
			keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

			var cr gjson.Result
			r.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}").ForEach(
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

			{{- range .Attributes}}
			{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if value := cr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {{if .DefaultValue}}if data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Null()
			}
			{{- else if isListSet .}}
			if value := cr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if eq .Type "Map"}}
			if value := cr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.IsNull() {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = helpers.GetStringMap(value.Map())
			} else {
				data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.MapNull(types.StringType)
			}
			{{- else if isNestedListSet .}}
			{{- $cclist := (toGoName .TfName)}}
			for cci := range data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} {
				keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
				keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

				var ccr gjson.Result
				cr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}").ForEach(
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

				{{- range .Attributes}}
				{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
				{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
				if value := ccr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.IsNull() {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
				} else {{if .DefaultValue}}if data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Null()
				}
				{{- else if isListSet .}}
				if value := ccr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.IsNull() {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
				} else {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
				}
				{{- else if eq .Type "Map"}}
				if value := ccr.Get("{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.IsNull() {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = helpers.GetStringMap(value.Map())
				} else {
					data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.MapNull(types.StringType)
				}
				{{- end}}
				{{- end}}
				{{- end}}
			}

			{{- end}}
			{{- end}}
			{{- end}}
		}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
}
// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *{{camelCase .Name}}) isNull(ctx context.Context, res gjson.Result) bool {
	{{- range .Attributes}}
	{{- if and (not .Value) (not .Id) (not .Reference) (not .QueryParam)}}
	{{- if isNestedListSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		return false
	}
	{{- else}}
	if !data.{{toGoName .TfName}}.IsNull() {
		return false
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return true
}
// End of section. //template:end isNull
