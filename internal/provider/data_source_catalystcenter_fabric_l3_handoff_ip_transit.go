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
	_ datasource.DataSource              = &FabricL3HandoffIPTransitDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricL3HandoffIPTransitDataSource{}
)

func NewFabricL3HandoffIPTransitDataSource() datasource.DataSource {
	return &FabricL3HandoffIPTransitDataSource{}
}

type FabricL3HandoffIPTransitDataSource struct {
	client *cc.Client
}

func (d *FabricL3HandoffIPTransitDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l3_handoff_ip_transit"
}

func (d *FabricL3HandoffIPTransitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric L3 Handoff IP Transit.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: "Network device ID of the fabric device",
				Required:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric this device belongs to",
				Required:            true,
			},
			"transit_network_id": schema.StringAttribute{
				MarkdownDescription: "ID of the transit network of the layer 3 handoff ip transit",
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface name of the layer 3 handoff ip transit",
				Computed:            true,
			},
			"external_connectivity_ip_pool_name": schema.StringAttribute{
				MarkdownDescription: "External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer",
				Computed:            true,
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Name of the virtual network associated with this fabric site",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094)",
				Computed:            true,
			},
			"tcp_mss_adjustment": schema.Int64Attribute{
				MarkdownDescription: "TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6",
				Computed:            true,
			},
			"local_ip_address": schema.StringAttribute{
				MarkdownDescription: "Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name",
				Computed:            true,
			},
			"remote_ip_address": schema.StringAttribute{
				MarkdownDescription: "Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name",
				Computed:            true,
			},
			"local_ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name",
				Computed:            true,
			},
			"remote_ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name",
				Computed:            true,
			},
		},
	}
}

func (d *FabricL3HandoffIPTransitDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricL3HandoffIPTransitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricL3HandoffIPTransit

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?networkDeviceId=" + url.QueryEscape(config.NetworkDeviceId.ValueString()) + "&fabricId=" + url.QueryEscape(config.FabricId.ValueString())
	res, err := d.client.Get(config.getPath() + params)
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
