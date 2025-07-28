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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FabricSitesDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricSitesDataSource{}
)

func NewFabricSitesDataSource() datasource.DataSource {
	return &FabricSitesDataSource{}
}

type FabricSitesDataSource struct {
	client *cc.Client
}

func (d *FabricSitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_sites"
}

func (d *FabricSitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches all fabric sites defined on the Catalyst Center. To retrieve detailed information about a specific site, use the data source `data.catalystcenter_fabric_site`.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"sites": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the fabric site",
							Computed:            true,
						},
						"site_id": schema.StringAttribute{
							MarkdownDescription: "ID of the network hierarchy associated with the fabric site",
							Computed:            true,
						},
						"authentication_profile_name": schema.StringAttribute{
							MarkdownDescription: "Authentication profile used for this fabric",
							Computed:            true,
						},
						"is_pub_sub_enabled": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether this fabric site will use pub/sub for control nodes",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *FabricSitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

func (d *FabricSitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricSites

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "singleton: Beginning Read")

	params := ""
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
