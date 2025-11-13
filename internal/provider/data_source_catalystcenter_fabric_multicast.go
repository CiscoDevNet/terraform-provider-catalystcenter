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
	_ datasource.DataSource              = &FabricMulticastDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricMulticastDataSource{}
)

func NewFabricMulticastDataSource() datasource.DataSource {
	return &FabricMulticastDataSource{}
}

type FabricMulticastDataSource struct {
	client *cc.Client
}

func (d *FabricMulticastDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_multicast"
}

func (d *FabricMulticastDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric Multicast.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric to contain this anycast gateway",
				Required:            true,
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it",
				Required:            true,
			},
			"ip_pool_name": schema.StringAttribute{
				MarkdownDescription: "Name of the IP pool associated with the anycast gateway",
				Computed:            true,
			},
			"ipv4_ssm_ranges": schema.ListAttribute{
				MarkdownDescription: "IPv4 Source Specific Multicast (SSM) ranges",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"multicast_r_ps": schema.ListNestedAttribute{
				MarkdownDescription: "Multicast Rendezvous Points (RP).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rp_device_location": schema.StringAttribute{
							MarkdownDescription: "Device location of the RP",
							Computed:            true,
						},
						"network_device_ids": schema.ListAttribute{
							MarkdownDescription: "IDs of the network devices%",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *FabricMulticastDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricMulticastDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricMulticast

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(config.FabricId.ValueString()) + "&virtualNetworkName=" + url.QueryEscape(config.VirtualNetworkName.ValueString())
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
