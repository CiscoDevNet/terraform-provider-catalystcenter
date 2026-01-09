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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FloorsDataSource{}
	_ datasource.DataSourceWithConfigure = &FloorsDataSource{}
)

func NewFloorsDataSource() datasource.DataSource {
	return &FloorsDataSource{}
}

type FloorsDataSource struct {
	client *cc.Client
}

func (d *FloorsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_floors"
}

func (d *FloorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Floors.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"floors": schema.SetNestedAttribute{
				MarkdownDescription: "List of floors",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the floor",
							Computed:            true,
						},
						"parent_id": schema.StringAttribute{
							MarkdownDescription: "The parent ID of the floor",
							Computed:            true,
						},
						"parent_name_hierarchy": schema.StringAttribute{
							MarkdownDescription: "Parent hierarchical name (e.g., Global/USA/Building1)",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Floor name",
							Computed:            true,
						},
						"floor_number": schema.Int64Attribute{
							MarkdownDescription: "Floor number",
							Computed:            true,
						},
						"rf_model": schema.StringAttribute{
							MarkdownDescription: "The RF model",
							Computed:            true,
						},
						"width": schema.Float64Attribute{
							MarkdownDescription: "Width",
							Computed:            true,
						},
						"length": schema.Float64Attribute{
							MarkdownDescription: "Length",
							Computed:            true,
						},
						"height": schema.Float64Attribute{
							MarkdownDescription: "Height",
							Computed:            true,
						},
						"units_of_measure": schema.StringAttribute{
							MarkdownDescription: "The unit of measurement",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *FloorsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FloorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Floors

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get("/dna/intent/api/v1/sites" + params)
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
