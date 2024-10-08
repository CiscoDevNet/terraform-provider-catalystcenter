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
	_ datasource.DataSource              = &IPPoolReservationDataSource{}
	_ datasource.DataSourceWithConfigure = &IPPoolReservationDataSource{}
)

func NewIPPoolReservationDataSource() datasource.DataSource {
	return &IPPoolReservationDataSource{}
}

type IPPoolReservationDataSource struct {
	client *cc.Client
}

func (d *IPPoolReservationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_pool_reservation"
}

func (d *IPPoolReservationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the IP Pool Reservation.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "The site ID",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the IP pool reservation",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the IP pool reservation",
				Computed:            true,
			},
			"ipv6_address_space": schema.BoolAttribute{
				MarkdownDescription: "If the value is `false` only IPv4 input are required, otherwise both IPv6 and IPv4 are required",
				Computed:            true,
			},
			"ipv4_global_pool": schema.StringAttribute{
				MarkdownDescription: "IPv4 Global pool address with cidr, example: 175.175.0.0/16",
				Computed:            true,
			},
			"ipv4_prefix": schema.BoolAttribute{
				MarkdownDescription: "If this value is `true`, the `ipv4_prefix_length` attribute must be provided, if it is `false`, the `ipv4_total_host` attribute must be provided",
				Computed:            true,
			},
			"ipv4_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "The IPv4 prefix length is required when `ipv4_prefix` value is `true`.",
				Computed:            true,
			},
			"ipv4_subnet": schema.StringAttribute{
				MarkdownDescription: "The IPv4 subnet",
				Computed:            true,
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: "The gateway for the IP pool reservation",
				Computed:            true,
			},
			"ipv4_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: "List of DHCP Server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dns_servers": schema.SetAttribute{
				MarkdownDescription: "List of DNS Server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv6_global_pool": schema.StringAttribute{
				MarkdownDescription: "IPv6 Global pool address with cidr, example: 2001:db8:85a3::/64",
				Computed:            true,
			},
			"ipv6_prefix": schema.BoolAttribute{
				MarkdownDescription: "If this value is `true`, the `ipv6_prefix_length` attribute must be provided, if it is `false`, the `ipv6_total_host` attribute must be provided",
				Computed:            true,
			},
			"ipv6_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "The IPv6 prefix length is required when `ipv6_prefix` value is `true`.",
				Computed:            true,
			},
			"ipv6_subnet": schema.StringAttribute{
				MarkdownDescription: "The IPv6 subnet, for example `2001:db8:85a3:0:100::`",
				Computed:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: "The gateway for the IP pool reservation",
				Computed:            true,
			},
			"ipv6_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: "List of DHCP Server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv6_dns_servers": schema.SetAttribute{
				MarkdownDescription: "List of DNS Server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_total_host": schema.Int64Attribute{
				MarkdownDescription: "The total number of IPv4 hosts",
				Computed:            true,
			},
			"ipv6_total_host": schema.Int64Attribute{
				MarkdownDescription: "The total number of IPv6 hosts",
				Computed:            true,
			},
			"slaac_support": schema.BoolAttribute{
				MarkdownDescription: "Enable SLAAC support",
				Computed:            true,
			},
		},
	}
}

func (d *IPPoolReservationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *IPPoolReservationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPPoolReservation

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?siteId=" + url.QueryEscape(config.SiteId.ValueString()) + "&groupName=" + url.QueryEscape(config.Name.ValueString())
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
