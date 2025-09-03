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
	_ datasource.DataSource              = &WirelessProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessProfileDataSource{}
)

func NewWirelessProfileDataSource() datasource.DataSource {
	return &WirelessProfileDataSource{}
}

type WirelessProfileDataSource struct {
	client *cc.Client
}

func (d *WirelessProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_profile"
}

func (d *WirelessProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"wireless_profile_name": schema.StringAttribute{
				MarkdownDescription: "Wireless Network Profile Name",
				Required:            true,
			},
			"ssid_details": schema.SetNestedAttribute{
				MarkdownDescription: "SSID Details",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ssid_name": schema.StringAttribute{
							MarkdownDescription: "SSID Name",
							Computed:            true,
						},
						"enable_fabric": schema.BoolAttribute{
							MarkdownDescription: "True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time",
							Computed:            true,
						},
						"enable_flex_connect": schema.BoolAttribute{
							MarkdownDescription: "True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time",
							Computed:            true,
						},
						"local_to_vlan": schema.Int64Attribute{
							MarkdownDescription: "Local To Vlan Id",
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Interface Name",
							Computed:            true,
						},
						"wlan_profile_name": schema.StringAttribute{
							MarkdownDescription: "WLAN Profile Name",
							Computed:            true,
						},
						"dot11be_profile_id": schema.StringAttribute{
							MarkdownDescription: "802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured",
							Computed:            true,
						},
					},
				},
			},
			"additional_interfaces": schema.SetAttribute{
				MarkdownDescription: "These additional interfaces will be configured on the device as independent interfaces in addition to the interfaces mapped to SSIDs. Max Limit 4094",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ap_zones": schema.SetNestedAttribute{
				MarkdownDescription: "AP Zones",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ap_zone_name": schema.StringAttribute{
							MarkdownDescription: "AP Zone Name",
							Computed:            true,
						},
						"rf_profile_name": schema.StringAttribute{
							MarkdownDescription: "AP Zone Name",
							Computed:            true,
						},
						"ssids": schema.SetAttribute{
							MarkdownDescription: "ssids part of apZone",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WirelessProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?wirelessProfileName=" + url.QueryEscape(config.WirelessProfileName.ValueString())
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
