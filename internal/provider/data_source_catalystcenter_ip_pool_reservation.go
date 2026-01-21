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
	"strings"

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
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.",
				Computed:            true,
			},
			"pool_type": schema.StringAttribute{
				MarkdownDescription: "The type of the reserve IP subpool. Once created, this cannot be changed.",
				Computed:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "The id of the non-Global site that this subpool belongs to.",
				Computed:            true,
			},
			"ipv4_subnet": schema.StringAttribute{
				MarkdownDescription: "The IPv4 IP address component of the CIDR notation for this subnet.",
				Computed:            true,
			},
			"ipv4_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "The IPv4 network mask length as a decimal for the CIDR notation of this subnet.",
				Computed:            true,
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: "The IPv4 gateway IP address for this subnet.",
				Computed:            true,
			},
			"ipv4_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: "The IPv4 DHCP server(s) for this subnet.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dns_servers": schema.SetAttribute{
				MarkdownDescription: "The IPv4 DNS server(s) for this subnet.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_total_addresses": schema.StringAttribute{
				MarkdownDescription: "The total number of addresses in the IPv4 pool (numeric string).",
				Computed:            true,
			},
			"ipv4_unassignable_addresses": schema.StringAttribute{
				MarkdownDescription: "The number of addresses in the IPv4 pool that cannot be assigned (numeric string).",
				Computed:            true,
			},
			"ipv4_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: "The number of addresses assigned from the IPv4 pool (numeric string).",
				Computed:            true,
			},
			"ipv4_default_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: "The number of addresses that are assigned from the IPv4 pool by default (numeric string).",
				Computed:            true,
			},
			"ipv4_global_pool_id": schema.StringAttribute{
				MarkdownDescription: "The non-tunnel global pool ID for this IPv4 reserve pool. Once added, this value cannot be changed.",
				Computed:            true,
			},
			"ipv6_subnet": schema.StringAttribute{
				MarkdownDescription: "The IPv6 IP address component of the CIDR notation for this subnet.",
				Computed:            true,
			},
			"ipv6_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "The IPv6 network mask length as a decimal for the CIDR notation of this subnet.",
				Computed:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: "The IPv6 gateway IP address for this subnet.",
				Computed:            true,
			},
			"ipv6_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: "The IPv6 DHCP server(s) for this subnet.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv6_dns_servers": schema.SetAttribute{
				MarkdownDescription: "The IPv6 DNS server(s) for this subnet.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv6_total_addresses": schema.StringAttribute{
				MarkdownDescription: "The total number of addresses in the IPv6 pool (numeric string, up to 128 bits).",
				Computed:            true,
			},
			"ipv6_unassignable_addresses": schema.StringAttribute{
				MarkdownDescription: "The number of addresses in the IPv6 pool that cannot be assigned (numeric string).",
				Computed:            true,
			},
			"ipv6_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: "The number of addresses assigned from the IPv6 pool (numeric string).",
				Computed:            true,
			},
			"ipv6_default_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: "The number of addresses that are assigned from the IPv6 pool by default (numeric string).",
				Computed:            true,
			},
			"ipv6_slaac_support": schema.BoolAttribute{
				MarkdownDescription: "If the IPv6 prefixLength is 64, this option may be enabled for SLAAC.",
				Computed:            true,
			},
			"ipv6_global_pool_id": schema.StringAttribute{
				MarkdownDescription: "The non-tunnel global pool ID for this IPv6 reserve pool. Once added, this value cannot be changed.",
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
	res, err := d.client.Get(config.getPath() + params)
	// Try fallback endpoint if primary fails with 404 or 500
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 500")) {
		tflog.Debug(ctx, fmt.Sprintf("%s: Primary endpoint returned 404, trying fallback endpoint", config.Id.ValueString()))
		res, err = d.client.Get(config.getFallbackPath() + params)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	res = res.Get("response.#(id==\"" + config.Id.ValueString() + "\")")

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
