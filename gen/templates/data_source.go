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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-catalystcenter"
	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &{{camelCase .Name}}DataSource{}
	_ datasource.DataSourceWithConfigure = &{{camelCase .Name}}DataSource{}
)

func New{{camelCase .Name}}DataSource() datasource.DataSource {
	return &{{camelCase .Name}}DataSource{}
}

type {{camelCase .Name}}DataSource struct {
	client *cc.Client
}

func (d *{{camelCase .Name}}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

{{- $dataSourceQuery := hasDataSourceQuery .Attributes}}

func (d *{{camelCase .Name}}DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "{{.DsDescription}}",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				{{- if .DataSourceNoId}}
				Computed:            true,
				{{- else if not $dataSourceQuery}}
				Required:            true,
				{{- else}}
				Optional:            true,
				Computed:            true,
				{{- end}}
			},
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: "{{.Description}}",
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- else if eq .Type "Map"}}
				ElementType:         types.StringType,
				{{- end}}
				{{- if and (or .Reference .QueryParam) (not .DataSourceQuery)}}
				Required:            true,
				{{- else}}
				{{- if .DataSourceQuery}}
				Optional:            true,
				{{- end}}
				Computed:            true,
				{{- end}}
				{{- if isNestedListSet .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: "{{.Description}}",
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- else if eq .Type "Map"}}
							ElementType:         types.StringType,
							{{- end}}
							Computed:            true,
							{{- if isNestedListSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: "{{.Description}}",
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- else if eq .Type "Map"}}
										ElementType:         types.StringType,
										{{- end}}
										Computed:            true,
										{{- if isNestedListSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: "{{.Description}}",
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- else if eq .Type "Map"}}
													ElementType:         types.StringType,
													{{- end}}
													Computed:            true,
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}

{{- if and $dataSourceQuery (not .DataSourceNoId)}}
func (d *{{camelCase .Name}}DataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
    return []datasource.ConfigValidator{
        datasourcevalidator.ExactlyOneOf(
            path.MatchRoot("id"),
			{{- range  .Attributes}}
			{{- if .DataSourceQuery}}
            path.MatchRoot("{{.TfName}}"),
			{{- end}}
			{{- end}}
        ),
    }
}
{{- end}}

func (d *{{camelCase .Name}}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}
// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *{{camelCase .Name}}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config {{camelCase .Name}}

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	{{- if $dataSourceQuery}}
	{{- $getFromAllPath := getFromAllPath .}}
	{{- $idFromQueryPathAttribute := .IdFromQueryPathAttribute}}
	{{- $getRestEndpoint := .GetRestEndpoint}}
	{{- $getExtraQueryParams := .GetExtraQueryParams}}
	{{- range .Attributes}}
	{{- if .DataSourceQuery}}
	if config.Id.IsNull() && !config.{{toGoName .TfName}}.IsNull() {
		res, err := d.client.Get({{if $getRestEndpoint}}"{{$getRestEndpoint}}"{{else}}config.getPath(){{end}}{{if $getExtraQueryParams}} + "{{$getExtraQueryParams}}"{{end}})
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res{{if .ResponseDataPath}}.Get("{{firstPathElement .ResponseDataPath $getFromAllPath}}"){{end}}; len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.{{toGoName .TfName}}.ValueString() == v.Get("{{if .ResponseDataPath}}{{remainingPathElements .ResponseDataPath $getFromAllPath}}{{else}}{{if .DataPath}}{{.DataPath}}.{{end}}{{.ModelName}}{{end}}").String() {
					config.Id = types.StringValue(v.Get("{{if $idFromQueryPathAttribute}}{{$idFromQueryPathAttribute}}{{else}}id{{end}}").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with {{.ModelName}} '%v', id: %v", config.Id.String(), config.{{toGoName .TfName}}.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with {{.ModelName}}: %s", config.{{toGoName .TfName}}.ValueString()))
			return
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}

	params := ""
	{{- if .IdQueryParam}}
	params += "?{{.IdQueryParam}}=" + url.QueryEscape(config.Id.ValueString())
	{{- else if and (hasQueryParam .Attributes) (not .GetRequiresId)}}
		{{- $queryParams := generateQueryParamString "GET" "config" .Attributes }}
		{{- if $queryParams }}
	params += {{$queryParams}}
		{{- end}}
	{{- else if and (not .GetNoId) (not .GetFromAll)}}
	params += "/" + url.QueryEscape(config.Id.ValueString())
	{{- end}}
	{{- if .GetExtraQueryParams}}
	params += "{{.GetExtraQueryParams}}"
	{{- end}}
	res, err := d.client.Get({{if .GetRestEndpoint}}"{{.GetRestEndpoint}}"{{else}}config.getPath(){{end}} + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	{{- if .GetFromAll}}
		{{- if .IdFromAttribute}}
			{{- $id := getId .Attributes}}
	res = res.Get("{{.IdFromQueryPath}}.#({{if $id.ResponseModelName}}{{$id.ResponseModelName}}{{else}}{{$id.ModelName}}{{end}}==\"" + config.Id.ValueString() + "\")")
		{{- else}}
	res = res.Get("{{.IdFromQueryPath}}.#({{if .IdFromQueryPathAttribute}}{{.IdFromQueryPathAttribute}}{{else}}id{{end}}==\"" + config.Id.ValueString() + "\")")
		{{- end}}
	{{- end}}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end read
