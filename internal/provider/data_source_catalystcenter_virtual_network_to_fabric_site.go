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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VirtualNetworkToFabricSiteDataSource{}
	_ datasource.DataSourceWithConfigure = &VirtualNetworkToFabricSiteDataSource{}
)

func NewVirtualNetworkToFabricSiteDataSource() datasource.DataSource {
	return &VirtualNetworkToFabricSiteDataSource{}
}

type VirtualNetworkToFabricSiteDataSource struct {
	client *cc.Client
}

func (d *VirtualNetworkToFabricSiteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_network_to_fabric_site"
}

func (d *VirtualNetworkToFabricSiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Virtual Network To Fabric Site.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Virtual Network Name",
				Required:            true,
			},
			"virtual_network_id": schema.StringAttribute{
				MarkdownDescription: "ID of the Virtual Network",
				Computed:            true,
			},
			"fabric_site_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric this Layer 3 Virtual Network is to be assigned to.",
				Required:            true,
			},
		},
	}
}

func (d *VirtualNetworkToFabricSiteDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (d *VirtualNetworkToFabricSiteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VirtualNetworkToFabricSite

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?virtualNetworkName=" + url.QueryEscape(config.VirtualNetworkName.ValueString())
	params += "&fabricId=" + url.QueryEscape(config.FabricSiteId.ValueString())
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	// Verify that the requested fabric site ID is in the response
	fabricIds := res.Get("response.0.fabricIds").Array()
	found := false
	for _, id := range fabricIds {
		if id.String() == config.FabricSiteId.ValueString() {
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddError("Not Found", fmt.Sprintf("Virtual network '%s' is not assigned to fabric site '%s'", config.VirtualNetworkName.ValueString(), config.FabricSiteId.ValueString()))
		return
	}

	// Set the composite ID
	compositeId := config.FabricSiteId.ValueString() + "--" + config.VirtualNetworkName.ValueString()
	config.Id = types.StringValue(compositeId)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
