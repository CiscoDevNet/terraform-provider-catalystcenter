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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessProfileSiteTagDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessProfileSiteTagDataSource{}
)

func NewWirelessProfileSiteTagDataSource() datasource.DataSource {
	return &WirelessProfileSiteTagDataSource{}
}

type WirelessProfileSiteTagDataSource struct {
	client *cc.Client
}

func (d *WirelessProfileSiteTagDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_profile_site_tag"
}

func (d *WirelessProfileSiteTagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless Profile Site Tag.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"wireless_profile_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the Wireless Profile",
				Required:            true,
			},
			"site_tag_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the Site Tag (computed after creation)",
				Required:            true,
			},
			"site_tag_name": schema.StringAttribute{
				MarkdownDescription: "Name of the Site Tag. Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space. Cannot be modified after creation.",
				Computed:            true,
			},
			"ap_profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the AP Profile to associate with this Site Tag",
				Computed:            true,
			},
			"site_ids": schema.SetAttribute{
				MarkdownDescription: "Set of Site IDs where this Site Tag applies. Must be Area, Building, or Floor level sites (not Global).",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"flex_profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the Flex Profile. Required for Flex-enabled Wireless Profiles. If not provided for a Flex profile, the system will auto-generate one.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessProfileSiteTagDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessProfileSiteTagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessProfileSiteTag

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	res, err := d.client.Get(config.getPathGet() + params)
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
