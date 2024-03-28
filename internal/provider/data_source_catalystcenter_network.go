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
	_ datasource.DataSource              = &NetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkDataSource{}
)

func NewNetworkDataSource() datasource.DataSource {
	return &NetworkDataSource{}
}

type NetworkDataSource struct {
	client *cc.Client
}

func (d *NetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

func (d *NetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Network.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "The site ID",
				Required:            true,
			},
			"dhcp_servers": schema.SetAttribute{
				MarkdownDescription: "List of DHCP server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"domain_name": schema.StringAttribute{
				MarkdownDescription: "Domain name",
				Computed:            true,
			},
			"primary_dns_server": schema.StringAttribute{
				MarkdownDescription: "Primary DNS server IP",
				Computed:            true,
			},
			"secondary_dns_server": schema.StringAttribute{
				MarkdownDescription: "Secondary DNS server IP",
				Computed:            true,
			},
			"syslog_servers": schema.SetAttribute{
				MarkdownDescription: "List of Syslog server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"catalyst_center_as_syslog_server": schema.BoolAttribute{
				MarkdownDescription: "Use Catalyst Center as Syslog server",
				Computed:            true,
			},
			"snmp_servers": schema.SetAttribute{
				MarkdownDescription: "List of SNMP server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"catalyst_center_as_snmp_server": schema.BoolAttribute{
				MarkdownDescription: "Use Catalyst Center as SNMP server",
				Computed:            true,
			},
			"netflow_collector": schema.StringAttribute{
				MarkdownDescription: "Netflow collector IP",
				Computed:            true,
			},
			"netflow_collector_port": schema.Int64Attribute{
				MarkdownDescription: "Netflow collector port",
				Computed:            true,
			},
			"ntp_servers": schema.SetAttribute{
				MarkdownDescription: "List of NTP server IPs",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "Timezone, e.g. `Africa/Abidjan`",
				Computed:            true,
			},
			"network_aaa_server_type": schema.StringAttribute{
				MarkdownDescription: "Type of network AAA server",
				Computed:            true,
			},
			"network_aaa_server_primary_ip": schema.StringAttribute{
				MarkdownDescription: "In case of `ISE` server type, this is the PAN IP address, in case of `AAA` this is the primary IP address",
				Computed:            true,
			},
			"network_aaa_server_secondary_ip": schema.StringAttribute{
				MarkdownDescription: "In case of `ISE` server type, this is the PSN IP address, in case of `AAA` this is a secondary IP address",
				Computed:            true,
			},
			"network_aaa_server_protocol": schema.StringAttribute{
				MarkdownDescription: "Server protocol",
				Computed:            true,
			},
			"network_aaa_server_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Only relevant for type `ISE`, shared secret",
				Computed:            true,
			},
			"endpoint_aaa_server_type": schema.StringAttribute{
				MarkdownDescription: "Type of network AAA server",
				Computed:            true,
			},
			"endpoint_aaa_server_primary_ip": schema.StringAttribute{
				MarkdownDescription: "In case of `ISE` server type, this is the PAN IP address, in case of `AAA` this is the primary IP address",
				Computed:            true,
			},
			"endpoint_aaa_server_secondary_ip": schema.StringAttribute{
				MarkdownDescription: "In case of `ISE` server type, this is the PSN IP address, in case of `AAA` this is a secondary IP address",
				Computed:            true,
			},
			"endpoint_aaa_server_protocol": schema.StringAttribute{
				MarkdownDescription: "Server protocol",
				Computed:            true,
			},
			"endpoint_aaa_server_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Only relevant for type `ISE`, shared secret",
				Computed:            true,
			},
		},
	}
}

func (d *NetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *NetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Network

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get("/api/v1/commonsetting/global" + params)
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
