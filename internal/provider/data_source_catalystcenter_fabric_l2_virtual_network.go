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
	_ datasource.DataSource              = &FabricL2VirtualNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricL2VirtualNetworkDataSource{}
)

func NewFabricL2VirtualNetworkDataSource() datasource.DataSource {
	return &FabricL2VirtualNetworkDataSource{}
}

type FabricL2VirtualNetworkDataSource struct {
	client *cc.Client
}

func (d *FabricL2VirtualNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l2_virtual_network"
}

func (d *FabricL2VirtualNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric L2 Virtual Network.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric this layer 2 virtual network is to be assigned to",
				Required:            true,
			},
			"vlan_name": schema.StringAttribute{
				MarkdownDescription: "Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens",
				Required:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the VLAN of the layer 2 virtual network. Allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, and 2046. If deploying on a fabric zone, this vlanId must match the vlanId of the corresponding layer 2 virtual network on the fabric site",
				Computed:            true,
			},
			"traffic_type": schema.StringAttribute{
				MarkdownDescription: "The type of traffic that is served",
				Computed:            true,
			},
			"fabric_enabled_wireless": schema.BoolAttribute{
				MarkdownDescription: "Set to true to enable wireless. Default is false",
				Computed:            true,
			},
			"associated_l3_virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. The layer 3 virtual network must have already been added to the fabric before association. This field must either be present in all payload elements or none",
				Computed:            true,
			},
		},
	}
}

func (d *FabricL2VirtualNetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricL2VirtualNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricL2VirtualNetwork

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(config.FabricId.ValueString()) + "&vlanName=" + url.QueryEscape(config.VlanName.ValueString())
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
