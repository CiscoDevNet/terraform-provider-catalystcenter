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
	_ datasource.DataSource              = &IPPoolDataSource{}
	_ datasource.DataSourceWithConfigure = &IPPoolDataSource{}
)

func NewIPPoolDataSource() datasource.DataSource {
	return &IPPoolDataSource{}
}

type IPPoolDataSource struct {
	client *cc.Client
}

func (d *IPPoolDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_pool"
}

func (d *IPPoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the IP Pool.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name for this global IP pool. Only letters, numbers, '-', '_', '.', and '/' are allowed.",
				Computed:            true,
			},
			"pool_type": schema.StringAttribute{
				MarkdownDescription: "The type of the global IP pool. Once created, this cannot be changed.",
				Computed:            true,
			},
			"address_space_subnet": schema.StringAttribute{
				MarkdownDescription: "The IP address component of the CIDR notation for this subnet.",
				Computed:            true,
			},
			"address_space_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "The network mask component, as a decimal, for the CIDR notation of this subnet.",
				Computed:            true,
			},
			"address_space_gateway": schema.StringAttribute{
				MarkdownDescription: "The gateway IP address for this subnet.",
				Computed:            true,
			},
			"address_space_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: "The DHCP server(s) for this subnet.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"address_space_dns_servers": schema.SetAttribute{
				MarkdownDescription: "The DNS server(s) for this subnet.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *IPPoolDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *IPPoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPPool

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
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
