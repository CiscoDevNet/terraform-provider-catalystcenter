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
	_ datasource.DataSource              = &AAASettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &AAASettingsDataSource{}
)

func NewAAASettingsDataSource() datasource.DataSource {
	return &AAASettingsDataSource{}
}

type AAASettingsDataSource struct {
	client *cc.Client
}

func (d *AAASettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_settings"
}

func (d *AAASettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the AAA Settings.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "The site ID",
				Required:            true,
			},
			"network_aaa_server_type": schema.StringAttribute{
				MarkdownDescription: "Type of network AAA server",
				Computed:            true,
			},
			"network_aaa_protocol": schema.StringAttribute{
				MarkdownDescription: "Server protocol",
				Computed:            true,
			},
			"network_aaa_pan": schema.StringAttribute{
				MarkdownDescription: "Administration Node. Required for ISE",
				Computed:            true,
			},
			"network_aaa_primary_server_ip": schema.StringAttribute{
				MarkdownDescription: "The server to use as a primary",
				Computed:            true,
			},
			"network_aaa_secondary_server_ip": schema.StringAttribute{
				MarkdownDescription: "The server to use as a secondary",
				Computed:            true,
			},
			"network_aaa_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Only relevant for server type `ISE`, shared secret",
				Computed:            true,
			},
			"client_aaa_server_type": schema.StringAttribute{
				MarkdownDescription: "Type of client AAA server",
				Computed:            true,
			},
			"client_aaa_protocol": schema.StringAttribute{
				MarkdownDescription: "Server protocol",
				Computed:            true,
			},
			"client_aaa_pan": schema.StringAttribute{
				MarkdownDescription: "Administration Node. Required for ISE",
				Computed:            true,
			},
			"client_aaa_primary_server_ip": schema.StringAttribute{
				MarkdownDescription: "The server to use as a primary",
				Computed:            true,
			},
			"client_aaa_secondary_server_ip": schema.StringAttribute{
				MarkdownDescription: "The server to use as a secondary",
				Computed:            true,
			},
			"client_aaa_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Only relevant for server type `ISE`, shared secret",
				Computed:            true,
			},
		},
	}
}

func (d *AAASettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *AAASettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AAASettings

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?_inherited=true"
	res, err := d.client.Get(config.getPath() + params)
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