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
	_ datasource.DataSource              = &FabricL3HandoffSDATransitDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricL3HandoffSDATransitDataSource{}
)

func NewFabricL3HandoffSDATransitDataSource() datasource.DataSource {
	return &FabricL3HandoffSDATransitDataSource{}
}

type FabricL3HandoffSDATransitDataSource struct {
	client *cc.Client
}

func (d *FabricL3HandoffSDATransitDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l3_handoff_sda_transit"
}

func (d *FabricL3HandoffSDATransitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric L3 Handoff SDA Transit.",

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
				MarkdownDescription: "ID of the fabric this device is assigned to",
				Required:            true,
			},
			"transit_network_id": schema.StringAttribute{
				MarkdownDescription: "ID of the transit network of the layer 3 handoff sda transit",
				Computed:            true,
			},
			"affinity_id_prime": schema.Int64Attribute{
				MarkdownDescription: "Affinity id prime value of the border node. It supersedes the border priority to determine border node preference.",
				Computed:            true,
			},
			"affinity_id_decider": schema.Int64Attribute{
				MarkdownDescription: "Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker.",
				Computed:            true,
			},
			"connected_to_internet": schema.BoolAttribute{
				MarkdownDescription: "Set this true to allow associated site to provide internet access to other sites through sd-access.",
				Computed:            true,
			},
			"is_multicast_over_transit_enabled": schema.BoolAttribute{
				MarkdownDescription: "Set this true to configure native multicast over multiple sites that are connected to an sd-access transit.",
				Computed:            true,
			},
		},
	}
}

func (d *FabricL3HandoffSDATransitDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricL3HandoffSDATransitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricL3HandoffSDATransit

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
