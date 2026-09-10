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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AnchorGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &AnchorGroupDataSource{}
)

func NewAnchorGroupDataSource() datasource.DataSource {
	return &AnchorGroupDataSource{}
}

type AnchorGroupDataSource struct {
	client *cc.Client
}

func (d *AnchorGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_anchor_group"
}

func (d *AnchorGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Anchor Group.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"anchor_group_name": schema.StringAttribute{
				MarkdownDescription: "Anchor Group Name. Max length is 32 characters.",
				Computed:            true,
			},
			"mobility_anchors": schema.ListNestedAttribute{
				MarkdownDescription: "Mobility Anchor details. Maximum 3 anchors are allowed.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"device_name": schema.StringAttribute{
							MarkdownDescription: "Peer device host name.",
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Mobility public IP address of the anchor WLC.",
							Computed:            true,
						},
						"anchor_priority": schema.StringAttribute{
							MarkdownDescription: "Anchor priority. Only one priority value is allowed per anchor WLC.",
							Computed:            true,
						},
						"managed_anchor_wlc": schema.BoolAttribute{
							MarkdownDescription: "Whether the anchor WLC is managed by the Network Controller.",
							Computed:            true,
						},
						"peer_device_type": schema.StringAttribute{
							MarkdownDescription: "Indicates whether the peer device belongs to AireOS or IOS-XE family. Required for external (unmanaged) anchors.",
							Computed:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "Peer device mobility MAC address. Allowed formats: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11. Required for external (unmanaged) anchors.",
							Computed:            true,
						},
						"mobility_group_name": schema.StringAttribute{
							MarkdownDescription: "Peer device mobility group name. Must be alphanumeric, maximum 31 characters. Required for external (unmanaged) anchors.",
							Computed:            true,
						},
						"private_ip": schema.StringAttribute{
							MarkdownDescription: "Private management IP address of the anchor WLC.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AnchorGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *AnchorGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	applyProviderMeta(d.client, ctx, req.ProviderMeta)
	var config AnchorGroup

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
