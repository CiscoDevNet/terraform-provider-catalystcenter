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
	_ datasource.DataSource              = &IPPoolsDataSource{}
	_ datasource.DataSourceWithConfigure = &IPPoolsDataSource{}
)

func NewIPPoolsDataSource() datasource.DataSource {
	return &IPPoolsDataSource{}
}

type IPPoolsDataSource struct {
	client *cc.Client
}

func (d *IPPoolsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_pools"
}

func (d *IPPoolsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches all global ip pools defined on the Catalyst Center. To retrieve detailed information about a specific pool, use the data source `data.catalystcenter_ip_pool`.",

		Attributes: map[string]schema.Attribute{
			"pools": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the IP pool",
							Computed:            true,
						},
						"pool_type": schema.StringAttribute{
							MarkdownDescription: "Pool type",
							Computed:            true,
						},
						"gateway_ip_address": schema.StringAttribute{
							MarkdownDescription: "The gateway for the IP pool",
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: "Subnet",
							Computed:            true,
						},
						"prefix_length": schema.StringAttribute{
							MarkdownDescription: "Prefix length",
							Computed:            true,
						},
						"dhcp_servers": schema.SetAttribute{
							MarkdownDescription: "List of DHCP Server IPs",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dns_servers": schema.SetAttribute{
							MarkdownDescription: "List of DNS Server IPs",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *IPPoolsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (d *IPPoolsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPPools

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "singleton: Beginning Read")

	params := ""
	res, err := d.client.Get(config.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 500")) {
		tflog.Debug(ctx, "Primary endpoint returned 404 or 500, trying fallback endpoint")
		res, err = d.client.Get(config.getFallbackPath() + params)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, "singleton: Read finished successfully")

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
