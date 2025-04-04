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
	_ datasource.DataSource              = &FabricPortAssignmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricPortAssignmentsDataSource{}
)

func NewFabricPortAssignmentsDataSource() datasource.DataSource {
	return &FabricPortAssignmentsDataSource{}
}

type FabricPortAssignmentsDataSource struct {
	client *cc.Client
}

func (d *FabricPortAssignmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_port_assignments"
}

func (d *FabricPortAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric Port Assignments.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric the device is assigned to",
				Required:            true,
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: "Network device ID of the port assignment",
				Required:            true,
			},
			"port_assignments": schema.SetNestedAttribute{
				MarkdownDescription: "List of port assignments in SD-Access fabric",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "ID of the port assignment",
							Computed:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: "ID of the fabric the device is assigned to",
							Computed:            true,
						},
						"network_device_id": schema.StringAttribute{
							MarkdownDescription: "Network device ID of the port assignment",
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Interface name of the port assignment",
							Computed:            true,
						},
						"connected_device_type": schema.StringAttribute{
							MarkdownDescription: "Connected device type of the port assignment",
							Computed:            true,
						},
						"data_vlan_name": schema.StringAttribute{
							MarkdownDescription: "Data VLAN name of the port assignment",
							Computed:            true,
						},
						"voice_vlan_name": schema.StringAttribute{
							MarkdownDescription: "Voice VLAN name of the port assignment",
							Computed:            true,
						},
						"authenticate_template_name": schema.StringAttribute{
							MarkdownDescription: "Authenticate template name of the port assignment",
							Computed:            true,
						},
						"security_group_name": schema.StringAttribute{
							MarkdownDescription: "Security group name of the port assignment",
							Computed:            true,
						},
						"interface_description": schema.StringAttribute{
							MarkdownDescription: "Interface description of the port assignment",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *FabricPortAssignmentsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricPortAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricPortAssignments

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(config.FabricId.ValueString()) + "&networkDeviceId=" + url.QueryEscape(config.NetworkDeviceId.ValueString())
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
