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
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-catalystcenter"
	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &{{camelCase .Name}}Resource{}
{{- if not .NoImport}}
var _ resource.ResourceWithImportState = &{{camelCase .Name}}Resource{}
{{- end}}

func New{{camelCase .Name}}Resource() resource.Resource {
	return &{{camelCase .Name}}Resource{}
}

type {{camelCase .Name}}Resource struct {
	client *cc.Client
}

func (r *{{camelCase .Name}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

func (r *{{camelCase .Name}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ResDescription}}").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				{{- if not .PutUpdateId}}
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				{{- end}}
			},
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
					.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
					{{- end -}}
					{{- if .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- else if eq .Type "Map"}}
				ElementType:         types.StringType,
				{{- end}}
				{{- if or .Id .MatchId .Reference .Mandatory}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- end}}
				{{- if or (len .DefaultValue) .Computed}}
				Computed:            true,
				{{- end}}
				{{- if len .EnumValues}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
				Validators: []validator.String{
					{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
					{{- end}}
					{{- range .StringPatterns}}
					stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
					{{- end}}
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
				Validators: []validator.Float64{
					float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
				},
				{{- end}}
				{{- if and (len .DefaultValue) (eq .Type "Int64")}}
				Default:             int64default.StaticInt64({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
				Default:             booldefault.StaticBool({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "String")}}
				Default:             stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- if or .Id .MatchId .Reference .RequiresReplace}}
				PlanModifiers: []planmodifier.{{if eq .Type "StringList"}}List{{else}}{{.Type}}{{end}}{
					{{if eq .Type "StringList"}}list{{else}}{{snakeCase .Type}}{{end}}planmodifier.RequiresReplace(),
				},
				{{- end}}
				{{- if .Computed}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
				},
				{{- end}}
				{{- if isNestedListSet .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
								{{- if len .EnumValues -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
								.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
								{{- end -}}
								{{- if .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- else if eq .Type "Map"}}
							ElementType:         types.StringType,
							{{- end}}
							{{- if or .Id .Reference .Mandatory}}
							Required:            true,
							{{- else if not .Computed}}
							Optional:            true,
							{{- end}}
							{{- if or (len .DefaultValue) .Computed}}
							Computed:            true,
							{{- end}}
							{{- if len .EnumValues}}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
							Validators: []validator.String{
								{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
								stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
								{{- end}}
								{{- range .StringPatterns}}
								stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
								{{- end}}
							},
							{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
							Validators: []validator.Int64{
								int64validator.Between({{.MinInt}}, {{.MaxInt}}),
							},
							{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
							Validators: []validator.Float64{
								float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
							},
							{{- end}}
							{{- if and (len .DefaultValue) (eq .Type "Int64")}}
							Default:             int64default.StaticInt64({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
							Default:             booldefault.StaticBool({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "String")}}
							Default:             stringdefault.StaticString("{{.DefaultValue}}"),
							{{- end}}
							{{- if .RequiresReplace}}
							PlanModifiers: []planmodifier.{{if eq .Type "StringList"}}List{{else}}{{.Type}}{{end}}{
								{{if eq .Type "StringList"}}list{{else}}{{snakeCase .Type}}{{end}}planmodifier.RequiresReplace(),
							},
							{{- end}}
							{{- if .Computed}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
							},
							{{- end}}
							{{- if isNestedListSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
											{{- if len .EnumValues -}}
											.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
											{{- end -}}
											{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
											.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
											{{- end -}}
											{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
											.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
											{{- end -}}
											{{- if .DefaultValue -}}
											.AddDefaultValueDescription("{{.DefaultValue}}")
											{{- end -}}
											.String,
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- else if eq .Type "Map"}}
										ElementType:         types.StringType,
										{{- end}}
										{{- if or .Id .Reference .Mandatory}}
										Required:            true,
										{{- else}}
										Optional:            true,
										{{- end}}
										{{- if or (len .DefaultValue) .Computed}}
										Computed:            true,
										{{- end}}
										{{- if len .EnumValues}}
										Validators: []validator.String{
											stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
										},
										{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
										Validators: []validator.String{
											{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
											stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
											{{- end}}
											{{- range .StringPatterns}}
											stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
											{{- end}}
										},
										{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
										Validators: []validator.Int64{
											int64validator.Between({{.MinInt}}, {{.MaxInt}}),
										},
										{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
										Validators: []validator.Float64{
											float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
										},
										{{- end}}
										{{- if and (len .DefaultValue) (eq .Type "Int64")}}
										Default:             int64default.StaticInt64({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
										Default:             booldefault.StaticBool({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "String")}}
										Default:             stringdefault.StaticString("{{.DefaultValue}}"),
										{{- end}}
										{{- if .RequiresReplace}}
										PlanModifiers: []planmodifier.{{if eq .Type "StringList"}}List{{else}}{{.Type}}{{end}}{
											{{if eq .Type "StringList"}}list{{else}}{{snakeCase .Type}}{{end}}planmodifier.RequiresReplace(),
										},
										{{- end}}
										{{- if .Computed}}
										PlanModifiers: []planmodifier.{{.Type}}{
											{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
										},
										{{- end}}
										{{- if isNestedListSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
														{{- if len .EnumValues -}}
														.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
														{{- end -}}
														{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
														.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
														{{- end -}}
														{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
														.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
														{{- end -}}
														{{- if .DefaultValue -}}
														.AddDefaultValueDescription("{{.DefaultValue}}")
														{{- end -}}
														.String,
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- else if eq .Type "Map"}}
													ElementType:         types.StringType,
													{{- end}}
													{{- if or .Id .Reference .Mandatory}}
													Required:            true,
													{{- else}}
													Optional:            true,
													{{- end}}
													{{- if or (len .DefaultValue) .Computed}}
													Computed:            true,
													{{- end}}
													{{- if len .EnumValues}}
													Validators: []validator.String{
														stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
													},
													{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
													Validators: []validator.String{
														{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
														stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
														{{- end}}
														{{- range .StringPatterns}}
														stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
														{{- end}}
													},
													{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
													Validators: []validator.Int64{
														int64validator.Between({{.MinInt}}, {{.MaxInt}}),
													},
													{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
													Validators: []validator.Float64{
														float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
													},
													{{- end}}
													{{- if and (len .DefaultValue) (eq .Type "Int64")}}
													Default:             int64default.StaticInt64({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
													Default:             booldefault.StaticBool({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "String")}}
													Default:             stringdefault.StaticString("{{.DefaultValue}}"),
													{{- end}}
													{{- if .RequiresReplace}}
													PlanModifiers: []planmodifier.{{if eq .Type "StringList"}}List{{else}}{{.Type}}{{end}}{
														{{if eq .Type "StringList"}}list{{else}}{{snakeCase .Type}}{{end}}planmodifier.RequiresReplace(),
													},
													{{- end}}
													{{- if .Computed}}
													PlanModifiers: []planmodifier.{{.Type}}{
														{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
													},
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- if or (ne .MinList 0) (ne .MaxList 0)}}
										Validators: []validator.List{
											{{- if ne .MinList 0}}
											listvalidator.SizeAtLeast({{.MinList}}),
											{{- end}}
											{{- if ne .MaxList 0}}
											listvalidator.SizeAtMost({{.MaxList}}),
											{{- end}}
										},
										{{- end}}
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- if or (ne .MinList 0) (ne .MaxList 0)}}
							Validators: []validator.List{
								{{- if ne .MinList 0}}
								listvalidator.SizeAtLeast({{.MinList}}),
								{{- end}}
								{{- if ne .MaxList 0}}
								listvalidator.SizeAtMost({{.MaxList}}),
								{{- end}}
							},
							{{- end}}
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- if or (ne .MinList 0) (ne .MaxList 0)}}
				Validators: []validator.List{
					{{- if ne .MinList 0}}
					listvalidator.SizeAtLeast({{.MinList}}),
					{{- end}}
					{{- if ne .MaxList 0}}
					listvalidator.SizeAtMost({{.MaxList}}),
					{{- end}}
				},
				{{- end}}
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}

func (r *{{camelCase .Name}}Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}
// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *{{camelCase .Name}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, {{camelCase .Name}}{})

	params := ""
	{{- if hasCreateQueryPath .Attributes}}
		{{- $createQueryPath := getCreateQueryPath .Attributes}}
	params += "/" + url.QueryEscape(plan.{{toGoName $createQueryPath.TfName}}.Value{{$createQueryPath.Type}}())
	{{- end}}
	{{- if .PutCreate}}
	res, err := r.client.Put(plan.getPath() + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- else}}
	res, err := r.client.Post(plan.getPath() + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- end}}
	if err != nil {
	{{- if .DeviceUnreachabilityWarning}}
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCDP10000" {
			// Log a warning and continue execution when device is unreachable
			failureReason := res.Get("response.failureReason").String()
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", {{- if .PutCreate }} "PUT" {{- else }} "POST" {{- end }}, err, res.String()))
			return
		}
	{{- else}}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", {{- if .PutCreate }} "PUT" {{- else }} "POST" {{- end }}, err, res.String()))
		return
	{{- end}}
	}
	{{- /* Check if id can be resolved directly from response */}}
	{{- if .IdPath}}
	plan.Id = types.StringValue(res.Get("{{.IdPath}}").String())
	{{- /* Check if we need an extra query to resolve the id */}}
	{{- else if and .IdFromQueryPath (not .IdFromAttribute)}}
		{{- $id := getMatchId .Attributes}}
	params = ""
		{{- if hasQueryParam .Attributes}}
		{{- $queryParams := generateQueryParamString "GET" "plan" .Attributes }}
		{{- if $queryParams }}
	params += {{$queryParams}}
		{{- end}}
		{{- end}}
	{{- if .GetExtraQueryParams}}
	params += "{{.GetExtraQueryParams}}"
	{{- end}}
	res, err = r.client.Get({{if .GetRestEndpoint}}"{{.GetRestEndpoint}}"{{else}}plan.getPath(){{end}} + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	{{- if and .IdFromQueryPathAttribute .IdFromQueryPath (not .GetExtraQueryParams) (not .GetFromAll) }}
	plan.Id = types.StringValue(res.Get("{{if eq .IdFromQueryPath "." }}{{else}}{{.IdFromQueryPath}}.{{end}}{{.IdFromQueryPathAttribute}}").String())
	{{- else}}
	plan.Id = types.StringValue(res.Get("{{.IdFromQueryPath}}.#({{if $id.ResponseModelName}}{{$id.ResponseModelName}}{{else}}{{$id.ModelName}}{{end}}==\""+ {{if eq $id.Type "Int64"}}strconv.FormatInt(plan.{{toGoName $id.TfName}}.ValueInt64(), 10){{else}}plan.{{toGoName $id.TfName}}.Value{{$id.Type}}(){{end}} +"\").{{if .IdFromQueryPathAttribute}}{{.IdFromQueryPathAttribute}}{{else}}id{{end}}").String())
	{{- end}}
	{{- /* If we have an id attribute we will use that as id */}}
	{{- else if hasId .Attributes}}
		{{- $id := getId .Attributes}}
	plan.Id = types.StringValue(fmt.Sprint(plan.{{toGoName $id.TfName}}.Value{{$id.Type}}()))
	{{- end}}

	{{- if .UpdateComputed }}
	params = ""
	{{- $queryParams := generateQueryParamString "GET" "plan" .Attributes }}

	{{- if .IdQueryParam}}
	params += "?{{.IdQueryParam}}=" + url.QueryEscape(plan.Id.ValueString())
	{{- else if and (hasQueryParam .Attributes) (not .GetRequiresId)}}
	{{- if $queryParams }}
	params += {{$queryParams}}
	{{- end}}
	{{- else if and (not .GetNoId) (not .GetFromAll)}}
	params += "/" + url.QueryEscape(plan.Id.ValueString())
	{{- end}}
	{{- if .GetExtraQueryParams}}
	params += "{{.GetExtraQueryParams}}"
	{{- end}}
	res, err = r.client.Get({{if .GetRestEndpoint}}"{{.GetRestEndpoint}}"{{else}}plan.getPath(){{end}} + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	{{- if .GetFromAll}}
		{{- if .IdFromAttribute}}
			{{- $id := getId .Attributes}}
	res = res.Get("{{.IdFromQueryPath}}.#({{if $id.ResponseModelName}}{{$id.ResponseModelName}}{{else}}{{$id.ModelName}}{{end}}==\"" + state.{{toGoName $id.TfName}}.Value{{$id.Type}}() + "\")")
		{{- else}}
	res = res.Get("{{.IdFromQueryPath}}.#({{if .IdFromQueryPathAttribute}}{{.IdFromQueryPathAttribute}}{{else}}id{{end}}==\"" + state.Id.ValueString() + "\")")
		{{- end}}
	{{- end}}
	plan.fromBodyUnknowns(ctx, res)
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *{{camelCase .Name}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	{{- if not .NoRead}}

	params := ""
	{{- $queryParams := generateQueryParamString "GET" "state" .Attributes }}

	{{- if .IdQueryParam}}
	params += "?{{.IdQueryParam}}=" + url.QueryEscape(state.Id.ValueString())
	{{- else if and (hasQueryParam .Attributes) (not .GetRequiresId)}}
	{{- if $queryParams }}
	params += {{$queryParams}}
	{{- end}}
	{{- else if and (not .GetNoId) (not .GetFromAll)}}
	params += "/" + url.QueryEscape(state.Id.ValueString())
	{{- end}}
	{{- if .GetExtraQueryParams}}
	params += "{{.GetExtraQueryParams}}"
	{{- end}}
	res, err := r.client.Get({{if .GetRestEndpoint}}"{{.GetRestEndpoint}}"{{else}}state.getPath(){{end}} + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	{{- if .GetFromAll}}
		{{- if .IdFromAttribute}}
			{{- $id := getId .Attributes}}
	res = res.Get("{{.IdFromQueryPath}}.#({{if $id.ResponseModelName}}{{$id.ResponseModelName}}{{else}}{{$id.ModelName}}{{end}}==\"" + state.{{toGoName $id.TfName}}.Value{{$id.Type}}() + "\")")
		{{- else}}
	res = res.Get("{{.IdFromQueryPath}}.#({{if .IdFromQueryPathAttribute}}{{.IdFromQueryPathAttribute}}{{else}}id{{end}}==\"" + state.Id.ValueString() + "\")")
		{{- end}}
	{{- end}}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *{{camelCase .Name}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))
	{{- if not .NoUpdate}}

	body := plan.toBody(ctx, state)
	params := ""
	{{- if hasCreateQueryPath .Attributes}}
		{{- $createQueryPath := getCreateQueryPath .Attributes}}
	params += "/" + url.QueryEscape(plan.{{toGoName $createQueryPath.TfName}}.Value{{$createQueryPath.Type}}())
	{{- end}}
	{{- if .PutIdQueryParam}}
	params += "?{{.PutIdQueryParam}}=" + url.QueryEscape(plan.Id.ValueString())
	{{- end}}
	{{- if .PostUpdate}}
	res, err := r.client.Post(plan.getPath() + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- else if .PutCreate}}
	res, err := r.client.Put(plan.getPath() + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- else if hasQueryParam .Attributes}}
	res, err := r.client.Put({{if .PutRestEndpoint}}"{{.PutRestEndpoint}}"{{else}}plan.getPath(){{end}} + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- else if .PutNoId}}
	res, err := r.client.Put({{if .PutRestEndpoint}}"{{.PutRestEndpoint}}"{{else}}plan.getPath(){{end}} + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- else}}
	res, err := r.client.Put({{if .PutRestEndpoint}}"{{.PutRestEndpoint}}"{{else}}plan.getPath(){{end}} + "/" + url.QueryEscape(plan.Id.ValueString()) + params, body {{- if .MaxAsyncWaitTime }}, func(r *cc.Req) { r.MaxAsyncWaitTime={{.MaxAsyncWaitTime}} }{{end}})
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	{{- if and .IdPath .PutUpdateId}}
	plan.Id = types.StringValue(res.Get("{{.IdPath}}").String())
	{{- end}}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *{{camelCase .Name}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	{{- if not .NoDelete}}
	{{- if .PutDelete}}
	{{- if .DeleteNoId}}
	res, err := r.client.Put({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}}, "{}")
	{{- else if .DeleteIdQueryParam}}
	res, err := r.client.Put({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + "?{{.DeleteIdQueryParam}}=" + url.QueryEscape(state.Id.ValueString()), "{}")
	{{- else if hasDeleteQueryParam .Attributes }}
	{{- $queryParams := generateQueryParamString "DELETE" "state" .Attributes }}
	{{- if $queryParams }}
	params := {{$queryParams}}
	{{- end}}
	res, err := r.client.Put({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + params, "{}")
	{{- else}}
	res, err := r.client.Put({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + "/" + url.QueryEscape(state.Id.ValueString()), "{}")
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	{{- else}}
	{{- if .DeleteNoId}}
	res, err := r.client.Delete({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}})
	{{- else if .DeleteIdQueryParam}}
	res, err := r.client.Delete({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + "?{{.DeleteIdQueryParam}}=" + url.QueryEscape(state.Id.ValueString()))
	{{- else if hasDeleteQueryParam .Attributes }}
	{{- $queryParams := generateQueryParamString "DELETE" "state" .Attributes }}
	{{- if $queryParams }}
	params := {{$queryParams}}
	{{- end}}
	res, err := r.client.Delete({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + params)
	{{- else}}
	res, err := r.client.Delete({{if .DeleteRestEndpoint}}state.getPathDelete(){{else}}state.getPath(){{end}} + "/" + url.QueryEscape(state.Id.ValueString()))
	{{- end}}
	if err != nil {
		{{- if .DeviceUnreachabilityWarning}}
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCDP10000" {
			// Log a warning and continue execution when device is unreachable
			failureReason := res.Get("response.failureReason").String()
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "DELETE", err, res.String()))
			return
		}
	{{- else}}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	{{- end}}
	}
	{{- end}}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}
// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
{{- if not .NoImport}}
func (r *{{camelCase .Name}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	{{- if importAttributes .}}
	idParts := strings.Split(req.ID, ",")

    if len(idParts) != {{len (importAttributes .)}}{{range $index, $attr := importAttributes .}} || idParts[{{$index}}] == ""{{end}} {
        resp.Diagnostics.AddError(
            "Unexpected Import Identifier",
            fmt.Sprintf("Expected import identifier with format: {{range $index, $attr := importAttributes .}}{{if $index}},{{end}}<{{$attr.TfName}}>{{end}}. Got: %q", req.ID),
        )
        return
    }

	{{- $config := .}}
	{{- range $index, $attr := importAttributes .}}
    resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("{{$attr.TfName}}"), idParts[{{$index}}])...)
	{{- if and $config.IdFromAttribute $attr.Id}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[{{$index}}])...)
	{{- end}}
	{{- end}}
	{{- else}}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	{{- end}}
}
{{- end}}
// End of section. //template:end import
