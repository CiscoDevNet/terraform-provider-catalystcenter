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
	_ datasource.DataSource              = &FabricL2HandoffDataSource{}
	_ datasource.DataSourceWithConfigure = &FabricL2HandoffDataSource{}
)

func NewFabricL2HandoffDataSource() datasource.DataSource {
	return &FabricL2HandoffDataSource{}
}

type FabricL2HandoffDataSource struct {
	client *cc.Client
}

func (d *FabricL2HandoffDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_l2_handoff"
}

func (d *FabricL2HandoffDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Fabric L2 Handoff.",

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
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface name of the layer 2 handoff. E.g., GigabitEthernet1/0/4",
				Computed:            true,
			},
			"internal_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN number associated with this fabric. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094)",
				Computed:            true,
			},
			"external_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "External VLAN number into which the fabric must be extended. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094)",
				Computed:            true,
			},
		},
	}
}

func (d *FabricL2HandoffDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *FabricL2HandoffDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FabricL2Handoff

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
