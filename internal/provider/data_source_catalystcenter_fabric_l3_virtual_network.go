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
	_ datasource.DataSource              = &FabricL3VirtualNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricL3VirtualNetworkDataSource{}
)

func NewFabricL3VirtualNetworkDataSource() datasource.DataSource {
	return &FabricL3VirtualNetworkDataSource{}
}

type FabricL3VirtualNetworkDataSource struct {
	client *cc.Client
}

func (d *FabricL3VirtualNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l3_virtual_network"
}

func (d *FabricL3VirtualNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric L3 Virtual Network.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Name of the layer 3 virtual network. If `INFRA_VN` or `DEFAULT_VN` is used, those layer 3 virtual networks will be updated instead of created.",
				Required:            true,
			},
			"fabric_ids": schema.SetAttribute{
				MarkdownDescription: "IDs of the fabrics this layer 3 virtual network is to be assigned to.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"anchored_site_id": schema.StringAttribute{
				MarkdownDescription: "Fabric ID of the fabric site this layer 3 virtual network is to be anchored at.",
				Computed:            true,
			},
		},
	}
}

func (d *FabricL3VirtualNetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (d *FabricL3VirtualNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricL3VirtualNetwork

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var fabricIds []string
	config.FabricIds.ElementsAs(ctx, &fabricIds, false)
	params := ""
	params += "?virtualNetworkName=" + url.QueryEscape(config.VirtualNetworkName.ValueString())
	if len(fabricIds) > 0 {
		params += "&fabricIds=" + url.QueryEscape(fabricIds[0])
	}
	if !config.AnchoredSiteId.IsNull() && config.AnchoredSiteId.ValueString() != "" {
		params += "&anchoredSiteId=" + url.QueryEscape(config.AnchoredSiteId.ValueString())
	}
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
