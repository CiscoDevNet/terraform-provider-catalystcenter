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
	_ datasource.DataSource              = &AnycastGatewayDataSource{}
	_ datasource.DataSourceWithConfigure = &AnycastGatewayDataSource{}
)

func NewAnycastGatewayDataSource() datasource.DataSource {
	return &AnycastGatewayDataSource{}
}

type AnycastGatewayDataSource struct {
	client *cc.Client
}

func (d *AnycastGatewayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_anycast_gateway"
}

func (d *AnycastGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Anycast Gateway.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric to contain this anycast gateway",
				Computed:            true,
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it",
				Computed:            true,
			},
			"ip_pool_name": schema.StringAttribute{
				MarkdownDescription: "Name of the IP pool associated with the anycast gateway",
				Computed:            true,
			},
			"tcp_mss_adjustment": schema.Int64Attribute{
				MarkdownDescription: "TCP maximum segment size adjustment",
				Computed:            true,
			},
			"vlan_name": schema.StringAttribute{
				MarkdownDescription: "Name of the VLAN of the anycast gateway",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the VLAN of the anycast gateway. allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, 2046, and 4094. if deploying an anycast gateway on a fabric zone, this vlanId must match the vlanId of the corresponding anycast gateway on the fabric site",
				Computed:            true,
			},
			"traffic_type": schema.StringAttribute{
				MarkdownDescription: "The type of traffic the anycast gateway serves",
				Computed:            true,
			},
			"pool_type": schema.StringAttribute{
				MarkdownDescription: "The pool type of the anycast gateway (required for & applicable only to INFRA_VN)",
				Computed:            true,
			},
			"security_group_name": schema.StringAttribute{
				MarkdownDescription: "Name of the associated Security Group (not applicable to INFRA_VN)",
				Computed:            true,
			},
			"critical_pool": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable critical VLAN. if true, autoGenerateVlanName must also be true. (isCriticalPool is not applicable to INFRA_VN)",
				Computed:            true,
			},
			"l2_flooding_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable layer 2 flooding (not applicable to INFRA_VN)",
				Computed:            true,
			},
			"wireless_pool": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable fabric-enabled wireless (not applicable to INFRA_VN)",
				Computed:            true,
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable IP-directed broadcast (not applicable to INFRA_VN)",
				Computed:            true,
			},
			"intra_subnet_routing_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN)",
				Computed:            true,
			},
			"multiple_ip_to_mac_addresses": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN)",
				Computed:            true,
			},
			"supplicant_based_extended_node_onboarding": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN)",
				Computed:            true,
			},
			"auto_generate_vlan_name": schema.BoolAttribute{
				MarkdownDescription: "This field cannot be true when vlanName is provided. the vlanName will be generated as ipPoolGroupV4Cidr-virtualNetworkName for non-critical VLANs. for critical VLANs with DATA trafficType, vlanName will be CRITICAL_VLAN. for critical VLANs with VOICE trafficType, vlanName will be VOICE_VLAN",
				Computed:            true,
			},
		},
	}
}

func (d *AnycastGatewayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (d *AnycastGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AnycastGateway

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?id=" + url.QueryEscape(config.Id.ValueString())
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
