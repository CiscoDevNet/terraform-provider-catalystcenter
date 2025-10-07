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
	_ datasource.DataSource              = &FabricDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricDeviceDataSource{}
)

func NewFabricDeviceDataSource() datasource.DataSource {
	return &FabricDeviceDataSource{}
}

type FabricDeviceDataSource struct {
	client *cc.Client
}

func (d *FabricDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_device"
}

func (d *FabricDeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric Device.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: "Network device ID of the fabric device",
				Required:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric site/zone of this fabric device",
				Required:            true,
			},
			"device_roles": schema.SetAttribute{
				MarkdownDescription: "List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE]",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"border_types": schema.SetAttribute{
				MarkdownDescription: "List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3]",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"local_autonomous_system_number": schema.StringAttribute{
				MarkdownDescription: "BGP Local autonomous system number of the fabric border device",
				Computed:            true,
			},
			"default_exit": schema.BoolAttribute{
				MarkdownDescription: "Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes",
				Computed:            true,
			},
			"import_external_routes": schema.BoolAttribute{
				MarkdownDescription: "Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane",
				Computed:            true,
			},
			"border_priority": schema.Int64Attribute{
				MarkdownDescription: "Border priority of the fabric border device. A lower value indicates higher priority",
				Computed:            true,
			},
			"prepend_autonomous_system_count": schema.Int64Attribute{
				MarkdownDescription: "Prepend autonomous system count of the fabric border device",
				Computed:            true,
			},
		},
	}
}

func (d *FabricDeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricDevice

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

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
