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
	_ datasource.DataSource              = &FabricMulticastVirtualNetworksDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricMulticastVirtualNetworksDataSource{}
)

func NewFabricMulticastVirtualNetworksDataSource() datasource.DataSource {
	return &FabricMulticastVirtualNetworksDataSource{}
}

type FabricMulticastVirtualNetworksDataSource struct {
	client *cc.Client
}

func (d *FabricMulticastVirtualNetworksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_multicast_virtual_networks"
}

func (d *FabricMulticastVirtualNetworksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric Multicast Virtual Networks.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric site this multicast configuration is associated with",
				Required:            true,
			},
			"virtual_networks": schema.SetNestedAttribute{
				MarkdownDescription: "List of multicast virtual networks",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the multicast configuration",
							Computed:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: "ID of the fabric site this multicast configuration is associated with",
							Computed:            true,
						},
						"virtual_network_name": schema.StringAttribute{
							MarkdownDescription: "Name of the virtual network associated with the fabric site",
							Computed:            true,
						},
						"ip_pool_name": schema.StringAttribute{
							MarkdownDescription: "Name of the IP Pool associated with the fabric site",
							Computed:            true,
						},
						"ipv4_ssm_ranges": schema.SetAttribute{
							MarkdownDescription: "IPv4 Source Specific Multicast (SSM) ranges",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"multicast_rps": schema.SetNestedAttribute{
							MarkdownDescription: "Multicast Rendezvous Points (RP). Required for Any Source Multicast (ASM) scenario",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"rp_device_location": schema.StringAttribute{
										MarkdownDescription: "Device location of the RP",
										Computed:            true,
									},
									"ipv4_address": schema.StringAttribute{
										MarkdownDescription: "IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.",
										Computed:            true,
									},
									"ipv6_address": schema.StringAttribute{
										MarkdownDescription: "IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.",
										Computed:            true,
									},
									"is_default_v4_rp": schema.BoolAttribute{
										MarkdownDescription: "Specifies whether it is a default IPv4 RP",
										Computed:            true,
									},
									"is_default_v6_rp": schema.BoolAttribute{
										MarkdownDescription: "Specifies whether it is a default IPv6 RP",
										Computed:            true,
									},
									"network_device_ids": schema.SetAttribute{
										MarkdownDescription: "IDs of the network devices. This is a required field for fabric RPs",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"ipv4_asm_ranges": schema.SetAttribute{
										MarkdownDescription: "IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP.",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"ipv6_asm_ranges": schema.SetAttribute{
										MarkdownDescription: "IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP.",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *FabricMulticastVirtualNetworksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricMulticastVirtualNetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricMulticastVirtualNetworks

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(config.FabricId.ValueString())
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
