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
	_ datasource.DataSource              = &ApplicationDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplicationDataSource{}
)

func NewApplicationDataSource() datasource.DataSource {
	return &ApplicationDataSource{}
}

type ApplicationDataSource struct {
	client *cc.Client
}

func (d *ApplicationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

func (d *ApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Reads a Catalyst Center application, including built-in NBAR applications. Both `name` and the controller UUID `id` are required.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the application",
				Required:            true,
			},
			"network_applications": schema.ListNestedAttribute{
				MarkdownDescription: "Exactly one network application.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"help_string": schema.StringAttribute{
							MarkdownDescription: "Description of the application shown in the GUI",
							Computed:            true,
						},
						"application_type": schema.StringAttribute{
							MarkdownDescription: "Application type",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "How the application is identified. `_servername` pairs with `server_name`, `_url` with `url`, and `_server-ip` with `network_identity`.",
							Computed:            true,
						},
						"dscp": schema.StringAttribute{
							MarkdownDescription: "DSCP value assigned to traffic matching this application",
							Computed:            true,
						},
						"app_protocol": schema.StringAttribute{
							MarkdownDescription: "Transport protocol. Omit for `_servername`; use `TCP` for `_url`.",
							Computed:            true,
						},
						"server_name": schema.StringAttribute{
							MarkdownDescription: "Server name identifying the application, when `server_type` is `_servername`",
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL identifying the application, when `server_type` is `_url`",
							Computed:            true,
						},
						"traffic_class": schema.StringAttribute{
							MarkdownDescription: "The traffic class the application is assigned to",
							Computed:            true,
						},
						"category_id": schema.StringAttribute{
							MarkdownDescription: "ID of the application category. Copy this UUID from an existing application in the desired category.",
							Computed:            true,
						},
						"ignore_conflict": schema.BoolAttribute{
							MarkdownDescription: "Whether to ignore conflicts with existing application signatures.",
							Computed:            true,
						},
						"rank": schema.Int64Attribute{
							MarkdownDescription: "Rank used to break ties between overlapping applications",
							Computed:            true,
						},
						"engine_id": schema.Int64Attribute{
							MarkdownDescription: "Engine id, should be set to 6",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "Controller-assigned network application ID.",
							Computed:            true,
						},
						"application_sub_type": schema.StringAttribute{
							MarkdownDescription: "Application subtype returned by the controller.",
							Computed:            true,
						},
						"display_name": schema.StringAttribute{
							MarkdownDescription: "Display name returned by the controller.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Network application name returned by the controller.",
							Computed:            true,
						},
					},
				},
			},
			"application_set_id": schema.StringAttribute{
				MarkdownDescription: "Id reference to parent application set",
				Computed:            true,
			},
			"network_identity": schema.SetNestedAttribute{
				MarkdownDescription: "L3/L4 signatures for `_server-ip` applications",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol of the signature",
							Computed:            true,
						},
						"ports": schema.StringAttribute{
							MarkdownDescription: "Comma separated list of discrete ports",
							Computed:            true,
						},
						"lower_port": schema.Int64Attribute{
							MarkdownDescription: "Start of a port range, used with `upper_port`. Leave `ports` empty when a range is used. The controller returns `0` when no range is set.",
							Computed:            true,
						},
						"upper_port": schema.Int64Attribute{
							MarkdownDescription: "End of a port range, used with `lower_port`. Leave `ports` empty when a range is used. The controller returns `0` when no range is set.",
							Computed:            true,
						},
						"ipv4_subnet": schema.SetAttribute{
							MarkdownDescription: "IPv4 subnets matching the application",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "Id",
							Computed:            true,
						},
						"display_name": schema.StringAttribute{
							MarkdownDescription: "Display name returned by the controller.",
							Computed:            true,
						},
					},
				},
			},
			"indicative_network_identity": schema.SetNestedAttribute{
				MarkdownDescription: "Indicative signatures used by NBAR applications. Available when reading built-in applications through the data source. Each protocol, ports, lower-port and upper-port combination must be unique.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol of the indicative signature.",
							Computed:            true,
						},
						"ports": schema.StringAttribute{
							MarkdownDescription: "Discrete ports, or an empty string when a range is used.",
							Computed:            true,
						},
						"ipv4_subnet": schema.SetAttribute{
							MarkdownDescription: "IPv4 subnets of the indicative signature.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipv6_subnet": schema.SetAttribute{
							MarkdownDescription: "IPv6 subnets of the indicative signature.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"lower_port": schema.Int64Attribute{
							MarkdownDescription: "First port in the range, or `0` for discrete ports.",
							Computed:            true,
						},
						"upper_port": schema.Int64Attribute{
							MarkdownDescription: "Last port in the range, or `0` for discrete ports.",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "Id",
							Computed:            true,
						},
						"display_name": schema.StringAttribute{
							MarkdownDescription: "Display name returned by the controller.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplicationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ApplicationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	applyProviderMeta(d.client, ctx, req.ProviderMeta)
	var config Application

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?name=" + url.QueryEscape(config.Name.ValueString())
	params += "&attributes=application&offset=1&limit=500"
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
