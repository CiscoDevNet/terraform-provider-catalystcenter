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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
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
		MarkdownDescription: "This data source can read the Application.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the application",
				Optional:            true,
				Computed:            true,
			},
			"application_set_id": schema.StringAttribute{
				MarkdownDescription: "ID of the application set this application belongs to. Only `idRef` is sent on update; the `id` that `GET` returns alongside it must not be echoed back.",
				Computed:            true,
			},
			"category_id": schema.StringAttribute{
				MarkdownDescription: "ID of the application category. The controller exposes no endpoint to list categories, so this opaque UUID must be copied from an existing application in the desired category.",
				Computed:            true,
			},
			"traffic_class": schema.StringAttribute{
				MarkdownDescription: "The traffic class the application is assigned to",
				Computed:            true,
			},
			"help_string": schema.StringAttribute{
				MarkdownDescription: "Description of the application shown in the GUI",
				Computed:            true,
			},
			"dscp": schema.StringAttribute{
				MarkdownDescription: "DSCP value assigned to traffic matching this application",
				Computed:            true,
			},
			"rank": schema.Int64Attribute{
				MarkdownDescription: "Rank used to break ties between overlapping applications",
				Computed:            true,
			},
			"app_protocol": schema.StringAttribute{
				MarkdownDescription: "Transport protocol the application runs over",
				Computed:            true,
			},
			"server_type": schema.StringAttribute{
				MarkdownDescription: "How the application is identified. `_servername` pairs with `server_name`, `_url` with `url`.",
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
			"network_identity": schema.SetNestedAttribute{
				MarkdownDescription: "Optional L3/L4 signatures identifying the application",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol of the signature",
							Computed:            true,
						},
						"ports": schema.StringAttribute{
							MarkdownDescription: "Comma separated list of ports",
							Computed:            true,
						},
						"ipv4_subnet": schema.SetAttribute{
							MarkdownDescription: "IPv4 subnets matching the application",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"instance_id": schema.Int64Attribute{
				MarkdownDescription: "Server-assigned instance ID, echoed back on update",
				Computed:            true,
			},
			"display_name": schema.StringAttribute{
				MarkdownDescription: "Server-assigned display name, echoed back on update",
				Computed:            true,
			},
			"instance_version": schema.Int64Attribute{
				MarkdownDescription: "Server-assigned instance version, echoed back on update",
				Computed:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Server-assigned namespace, echoed back on update",
				Computed:            true,
			},
			"qualifier": schema.StringAttribute{
				MarkdownDescription: "Server-assigned qualifier, echoed back on update",
				Computed:            true,
			},
			"scalable_group_external_handle": schema.StringAttribute{
				MarkdownDescription: "Server-assigned external handle, echoed back on update",
				Computed:            true,
			},
			"network_application_id": schema.StringAttribute{
				MarkdownDescription: "Server-assigned ID of the nested network application, echoed back on update",
				Computed:            true,
			},
			"network_application_name": schema.StringAttribute{
				MarkdownDescription: "Server-assigned name of the nested network application, echoed back on update",
				Computed:            true,
			},
			"application_sub_type": schema.StringAttribute{
				MarkdownDescription: "Server-assigned application subtype, echoed back on update",
				Computed:            true,
			},
			"network_application_display_name": schema.StringAttribute{
				MarkdownDescription: "Server-assigned display name of the nested network application",
				Computed:            true,
			},
			"engine_id": schema.StringAttribute{
				MarkdownDescription: "Server-assigned engine ID, echoed back on update",
				Computed:            true,
			},
			"popularity": schema.Int64Attribute{
				MarkdownDescription: "Server-assigned popularity, echoed back on update",
				Computed:            true,
			},
			"selector_id": schema.StringAttribute{
				MarkdownDescription: "Server-assigned selector ID, echoed back on update",
				Computed:            true,
			},
		},
	}
}
func (d *ApplicationDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
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
	var config Application

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get(config.getPath() + "&attributes=application&limit=500")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res.Get("response"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	params := ""
	params += "?name=" + url.QueryEscape(config.Name.ValueString())
	params += "&attributes=application&limit=500"
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
