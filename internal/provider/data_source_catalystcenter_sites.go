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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SitesDataSource{}
	_ datasource.DataSourceWithConfigure = &SitesDataSource{}
)

func NewSitesDataSource() datasource.DataSource {
	return &SitesDataSource{}
}

type SitesDataSource struct {
	client *cc.Client
}

func (d *SitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sites"
}

func (d *SitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches all sites defined on the Catalyst Center.\n\n" +
			"Sites represent the logical and physical hierarchy of your network environment, including areas, buildings, and floors.\n" +
			"You can optionally specify the `type` attribute to filter the results and retrieve only sites of a specific type—such as `area`, `building` or `floor`.\n" +
			"To retrieve detailed information about a specific site, use the data source `data.catalystcenter_site`.",

		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "Site type (eg. area, building, floor)",
				Optional:            true,
			},
			"sites": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"name_hierarchy": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"parent_id": schema.StringAttribute{
							MarkdownDescription: "Parent Id",
							Computed:            true,
						},
						"latitude": schema.StringAttribute{
							MarkdownDescription: "Building latitude",
							Computed:            true,
						},
						"longitude": schema.StringAttribute{
							MarkdownDescription: "Building longitude",
							Computed:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: "Building address",
							Computed:            true,
						},
						"country": schema.StringAttribute{
							MarkdownDescription: "Country name for the building",
							Computed:            true,
						},
						"floor_number": schema.Int64Attribute{
							MarkdownDescription: "Floor number",
							Computed:            true,
						},
						"rf_model": schema.StringAttribute{
							MarkdownDescription: "Floor RF model",
							Computed:            true,
						},
						"width": schema.Float64Attribute{
							MarkdownDescription: "Floor width",
							Computed:            true,
						},
						"length": schema.Float64Attribute{
							MarkdownDescription: "Floor length",
							Computed:            true,
						},
						"height": schema.Float64Attribute{
							MarkdownDescription: "Floor height",
							Computed:            true,
						},
						"units_of_measure": schema.StringAttribute{
							MarkdownDescription: "Floor unit of measure",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (d *SitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Sites

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "singleton: Beginning Read")
	params := ""
	if !config.Type.IsNull() {
		params = "?type=" + config.Type.ValueString()
	} else {
		params = ""
	}
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, "singleton: Read finished successfully")

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
