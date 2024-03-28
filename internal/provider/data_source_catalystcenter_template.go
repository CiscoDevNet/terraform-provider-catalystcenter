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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &TemplateDataSource{}
)

func NewTemplateDataSource() datasource.DataSource {
	return &TemplateDataSource{}
}

type TemplateDataSource struct {
	client *cc.Client
}

func (d *TemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

func (d *TemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the project",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the template",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description",
				Computed:            true,
			},
			"device_types": schema.ListNestedAttribute{
				MarkdownDescription: "List of device types",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"product_family": schema.StringAttribute{
							MarkdownDescription: "Product family",
							Computed:            true,
						},
						"product_series": schema.StringAttribute{
							MarkdownDescription: "Product series",
							Computed:            true,
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: "Product type",
							Computed:            true,
						},
					},
				},
			},
			"language": schema.StringAttribute{
				MarkdownDescription: "Language of the template",
				Computed:            true,
			},
			"software_type": schema.StringAttribute{
				MarkdownDescription: "Software type",
				Computed:            true,
			},
			"software_variant": schema.StringAttribute{
				MarkdownDescription: "Software variant",
				Computed:            true,
			},
			"software_version": schema.StringAttribute{
				MarkdownDescription: "Software version",
				Computed:            true,
			},
			"template_content": schema.StringAttribute{
				MarkdownDescription: "Template content",
				Computed:            true,
			},
			"template_params": schema.ListNestedAttribute{
				MarkdownDescription: "List of template parameters",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"binding": schema.StringAttribute{
							MarkdownDescription: "Bind to source",
							Computed:            true,
						},
						"data_type": schema.StringAttribute{
							MarkdownDescription: "Datatype of template parameter",
							Computed:            true,
						},
						"default_value": schema.StringAttribute{
							MarkdownDescription: "Default value of template parameter",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of template parameter",
							Computed:            true,
						},
						"display_name": schema.StringAttribute{
							MarkdownDescription: "Display name of template parameter",
							Computed:            true,
						},
						"instruction_text": schema.StringAttribute{
							MarkdownDescription: "Instruction text",
							Computed:            true,
						},
						"not_param": schema.BoolAttribute{
							MarkdownDescription: "Is it not a variable",
							Computed:            true,
						},
						"param_array": schema.BoolAttribute{
							MarkdownDescription: "Is it an array",
							Computed:            true,
						},
						"parameter_name": schema.StringAttribute{
							MarkdownDescription: "Name of the template parameter",
							Computed:            true,
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: "List of ranges",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"max_value": schema.Int64Attribute{
										MarkdownDescription: "Max value of range",
										Computed:            true,
									},
									"min_value": schema.Int64Attribute{
										MarkdownDescription: "Min value of range",
										Computed:            true,
									},
								},
							},
						},
						"required": schema.BoolAttribute{
							MarkdownDescription: "Is parameter required",
							Computed:            true,
						},
						"default_selected_values": schema.SetAttribute{
							MarkdownDescription: "Default selection values",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"selection_type": schema.StringAttribute{
							MarkdownDescription: "Type of selection",
							Computed:            true,
						},
						"selection_values": schema.MapAttribute{
							MarkdownDescription: "Selection values",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *TemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *TemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Template

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get("/dna/intent/api/v1/template-programmer/template" + "?unCommitted=true")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res; len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					config.Id = types.StringValue(v.Get("templateId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	params += "?unCommitted=true"
	res, err := d.client.Get("/dna/intent/api/v1/template-programmer/template" + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
