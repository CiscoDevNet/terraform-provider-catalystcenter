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
	_ datasource.DataSource              = &WirelessCleanAirConfigurationDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessCleanAirConfigurationDataSource{}
)

func NewWirelessCleanAirConfigurationDataSource() datasource.DataSource {
	return &WirelessCleanAirConfigurationDataSource{}
}

type WirelessCleanAirConfigurationDataSource struct {
	client *cc.Client
}

func (d *WirelessCleanAirConfigurationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_cleanair_configuration"
}

func (d *WirelessCleanAirConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless CleanAir Configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"design_name": schema.StringAttribute{
				MarkdownDescription: "The name of the CleanAir feature template. Forbidden characters are `% & < > [ ] ' /`.",
				Computed:            true,
			},
			"radio_band": schema.StringAttribute{
				MarkdownDescription: "The radio band the CleanAir configuration applies to.",
				Computed:            true,
			},
			"clean_air": schema.BoolAttribute{
				MarkdownDescription: "Enable CleanAir spectrum intelligence on the radio band.",
				Computed:            true,
			},
			"clean_air_device_reporting": schema.BoolAttribute{
				MarkdownDescription: "Enable reporting of CleanAir interferer devices.",
				Computed:            true,
			},
			"persistent_device_propagation": schema.BoolAttribute{
				MarkdownDescription: "Enable persistent device propagation.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the CleanAir feature template.",
				Computed:            true,
			},
			"unlocked_attributes": schema.ListAttribute{
				MarkdownDescription: "List of feature attributes that are unlocked (editable) for this template.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ble_beacon": schema.BoolAttribute{
				MarkdownDescription: "Detect BLE Beacon interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"bluetooth_paging_inquiry": schema.BoolAttribute{
				MarkdownDescription: "Detect Bluetooth Paging Inquiry interferers. Applicable to 2.4GHz.",
				Computed:            true,
			},
			"bluetooth_sco_acl": schema.BoolAttribute{
				MarkdownDescription: "Detect Bluetooth SCO and ACL interferers. Applicable to 2.4GHz.",
				Computed:            true,
			},
			"continuous_transmitter": schema.BoolAttribute{
				MarkdownDescription: "Detect Continuous Transmitter interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"generic_dect": schema.BoolAttribute{
				MarkdownDescription: "Detect Generic DECT interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"generic_tdd": schema.BoolAttribute{
				MarkdownDescription: "Detect Generic TDD Transmitter interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"jammer": schema.BoolAttribute{
				MarkdownDescription: "Detect Jammer interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"microwave_oven": schema.BoolAttribute{
				MarkdownDescription: "Detect Microwave Oven interferers. Applicable to 2.4GHz.",
				Computed:            true,
			},
			"motorola_canopy": schema.BoolAttribute{
				MarkdownDescription: "Detect Motorola Canopy interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"si_fhss": schema.BoolAttribute{
				MarkdownDescription: "Detect SI FHSS interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"spectrum_80211_fh": schema.BoolAttribute{
				MarkdownDescription: "Detect 802.11 FH interferers. Applicable to 2.4GHz.",
				Computed:            true,
			},
			"spectrum_80211_non_standard_channel": schema.BoolAttribute{
				MarkdownDescription: "Detect 802.11 Non Standard Channel interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"spectrum_802154": schema.BoolAttribute{
				MarkdownDescription: "Detect 802.15.4 interferers. Applicable to 2.4GHz.",
				Computed:            true,
			},
			"spectrum_inverted": schema.BoolAttribute{
				MarkdownDescription: "Detect Spectrum Inverted interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"super_ag": schema.BoolAttribute{
				MarkdownDescription: "Detect Super AG interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"video_camera": schema.BoolAttribute{
				MarkdownDescription: "Detect Video Camera interferers. Applicable to 2.4GHz and 5GHz.",
				Computed:            true,
			},
			"wimax_fixed": schema.BoolAttribute{
				MarkdownDescription: "Detect WiMax Fixed interferers. Applicable to 5GHz.",
				Computed:            true,
			},
			"wimax_mobile": schema.BoolAttribute{
				MarkdownDescription: "Detect WiMax Mobile interferers. Applicable to 5GHz.",
				Computed:            true,
			},
			"xbox": schema.BoolAttribute{
				MarkdownDescription: "Detect Xbox interferers. Applicable to 2.4GHz.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessCleanAirConfigurationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessCleanAirConfigurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessCleanAirConfiguration

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
