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
	_ datasource.DataSource              = &WirelessRRMFRAConfigurationDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessRRMFRAConfigurationDataSource{}
)

func NewWirelessRRMFRAConfigurationDataSource() datasource.DataSource {
	return &WirelessRRMFRAConfigurationDataSource{}
}

type WirelessRRMFRAConfigurationDataSource struct {
	client *cc.Client
}

func (d *WirelessRRMFRAConfigurationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rrm_fra_configuration"
}

func (d *WirelessRRMFRAConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless RRM FRA Configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"design_name": schema.StringAttribute{
				MarkdownDescription: "The name of the RRM FRA feature template. Forbidden characters are `% & < > [ ] ' /`.",
				Computed:            true,
			},
			"radio_band": schema.StringAttribute{
				MarkdownDescription: "Radio band pair the FRA configuration applies to. `5GHZ_6GHZ` is supported only on IOS-XE Wireless Controllers running 17.9.1 and above.",
				Computed:            true,
			},
			"fra_freeze": schema.BoolAttribute{
				MarkdownDescription: "Enable Flexible Radio Assignment freeze. Supported on IOS-XE Wireless Controllers running version >= 17.6 for `2_4GHZ_5GHZ` and version >= 17.9 for `5GHZ_6GHZ`.",
				Computed:            true,
			},
			"fra_status": schema.BoolAttribute{
				MarkdownDescription: "Enable Flexible Radio Assignment.",
				Computed:            true,
			},
			"fra_interval": schema.Int64Attribute{
				MarkdownDescription: "Flexible Radio Assignment interval in hours (1-24).",
				Computed:            true,
			},
			"fra_sensitivity": schema.StringAttribute{
				MarkdownDescription: "Flexible Radio Assignment sensitivity. Values `Higher`, `Even Higher`, and `Super High` are supported only on IOS-XE Wireless Controllers running 17.5 and above and only for the `2_4GHZ_5GHZ` radio band.",
				Computed:            true,
			},
			"unlocked_attributes": schema.ListAttribute{
				MarkdownDescription: "Feature attributes unlocked in the design that can be changed at device provision time. Only attributes defined under `featureAttributes` are allowed.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *WirelessRRMFRAConfigurationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessRRMFRAConfigurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	applyProviderMeta(d.client, ctx, req.ProviderMeta)
	var config WirelessRRMFRAConfiguration

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
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
