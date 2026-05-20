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

// Custom imports: template markers removed to exclude unused gjson import
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

// End of custom imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &PlannedAccessPointPositionDataSource{}
	_ datasource.DataSourceWithConfigure = &PlannedAccessPointPositionDataSource{}
)

func NewPlannedAccessPointPositionDataSource() datasource.DataSource {
	return &PlannedAccessPointPositionDataSource{}
}

type PlannedAccessPointPositionDataSource struct {
	client *cc.Client
}

func (d *PlannedAccessPointPositionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_planned_access_point_position"
}

func (d *PlannedAccessPointPositionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source reads a planned Access Point position from a floor map in Cisco Catalyst Center.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"floor_id": schema.StringAttribute{
				MarkdownDescription: "Floor ID where the planned AP is placed",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the planned Access Point",
				Required:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC address of the planned Access Point",
				Computed:            true,
			},
			"ap_type": schema.StringAttribute{
				MarkdownDescription: "AP model type. Use /dna/intent/api/v1/maps/supported-access-points to find supported models",
				Computed:            true,
			},
			"position_x": schema.Float64Attribute{
				MarkdownDescription: "X coordinate of the planned AP position in feet",
				Computed:            true,
			},
			"position_y": schema.Float64Attribute{
				MarkdownDescription: "Y coordinate of the planned AP position in feet",
				Computed:            true,
			},
			"position_z": schema.Float64Attribute{
				MarkdownDescription: "Z coordinate (height) of the planned AP position in feet",
				Computed:            true,
			},
			"radios": schema.SetNestedAttribute{
				MarkdownDescription: "Radio configurations for the planned AP (1-4 radios). Any change to any radio attribute forces destroy+create because the Catalyst Center bulkChange API has limited update support.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"radio_id": schema.StringAttribute{
							MarkdownDescription: "Radio ID (read-only, assigned by Catalyst Center)",
							Computed:            true,
						},
						"bands": schema.SetAttribute{
							MarkdownDescription: "Radio frequency bands in GHz. Valid values are 2.4, 5, and 6",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"channel": schema.Int64Attribute{
							MarkdownDescription: "Channel number for map visualization",
							Computed:            true,
						},
						"tx_power": schema.Int64Attribute{
							MarkdownDescription: "Transmit power in dBm for map visualization",
							Computed:            true,
						},
						"antenna_name": schema.StringAttribute{
							MarkdownDescription: "Antenna type name. Use /dna/intent/api/v1/maps/supported-access-points to find supported antennas",
							Computed:            true,
						},
						"antenna_azimuth": schema.Int64Attribute{
							MarkdownDescription: "Antenna azimuth angle in degrees (0-360)",
							Computed:            true,
						},
						"antenna_elevation": schema.Int64Attribute{
							MarkdownDescription: "Antenna elevation angle in degrees (-90 to 90)",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *PlannedAccessPointPositionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Custom Read: template markers removed because floor_id is a URL path parameter
// with no model_name, causing the generated query-by-name code to fail
func (d *PlannedAccessPointPositionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PlannedAccessPointPosition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	if !config.Name.IsNull() {
		params += "?name=" + url.QueryEscape(config.Name.ValueString())
	} else if !config.Id.IsNull() {
		params += "/" + url.QueryEscape(config.Id.ValueString())
	}
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if config.Id.IsNull() || config.Id.ValueString() == "" {
		config.Id = types.StringValue(res.Get("response.0.id").String())
	}
	if config.Id.IsNull() || config.Id.ValueString() == "" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of custom Read
