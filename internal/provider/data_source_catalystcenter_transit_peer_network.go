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
	_ datasource.DataSource              = &TransitPeerNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &TransitPeerNetworkDataSource{}
)

func NewTransitPeerNetworkDataSource() datasource.DataSource {
	return &TransitPeerNetworkDataSource{}
}

type TransitPeerNetworkDataSource struct {
	client *cc.Client
}

func (d *TransitPeerNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transit_peer_network"
}

func (d *TransitPeerNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transit Peer Network.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"transit_peer_network_name": schema.StringAttribute{
				MarkdownDescription: "Transit Peer Network Name",
				Required:            true,
			},
			"transit_peer_network_type": schema.StringAttribute{
				MarkdownDescription: "Transit Peer Network Type",
				Computed:            true,
			},
			"routing_protocol_name": schema.StringAttribute{
				MarkdownDescription: "Routing Protocol Name",
				Computed:            true,
			},
			"autonomous_system_number": schema.StringAttribute{
				MarkdownDescription: "Autonomous System Number",
				Computed:            true,
			},
			"transit_control_plane_settings": schema.ListNestedAttribute{
				MarkdownDescription: "Transit Control Plane Settings info",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"site_name_hierarchy": schema.StringAttribute{
							MarkdownDescription: "Site Name Hierarchy where device is provisioned",
							Computed:            true,
						},
						"device_management_ip_address": schema.StringAttribute{
							MarkdownDescription: "Device Management Ip Address of provisioned device",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *TransitPeerNetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransitPeerNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransitPeerNetwork

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?transitPeerNetworkName=" + url.QueryEscape(config.TransitPeerNetworkName.ValueString())
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
