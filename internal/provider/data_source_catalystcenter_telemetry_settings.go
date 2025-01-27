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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TelemetrySettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &TelemetrySettingsDataSource{}
)

func NewTelemetrySettingsDataSource() datasource.DataSource {
	return &TelemetrySettingsDataSource{}
}

type TelemetrySettingsDataSource struct {
	client *cc.Client
}

func (d *TelemetrySettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_telemetry_settings"
}

func (d *TelemetrySettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Telemetry Settings.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "The site ID",
				Required:            true,
			},
			"enable_wired_data_collection": schema.BoolAttribute{
				MarkdownDescription: "Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory",
				Computed:            true,
			},
			"enable_wireless_telemetry": schema.BoolAttribute{
				MarkdownDescription: "Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients",
				Computed:            true,
			},
			"use_builtin_trap_server": schema.BoolAttribute{
				MarkdownDescription: "Enable this server as a destination server for SNMP traps and messages from your network",
				Computed:            true,
			},
			"external_trap_servers": schema.SetAttribute{
				MarkdownDescription: "External SNMP trap servers",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"use_builtin_syslog_server": schema.BoolAttribute{
				MarkdownDescription: "Enable this server as a destination server for syslog messages",
				Computed:            true,
			},
			"external_syslog_servers": schema.SetAttribute{
				MarkdownDescription: "External syslog servers",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"netflow_collector": schema.StringAttribute{
				MarkdownDescription: "Application Visibility Collector Type",
				Computed:            true,
			},
			"netflow_collector_ip_address": schema.StringAttribute{
				MarkdownDescription: "IP Address. If collection type is `TelemetryBrokerOrUDPDirector`, this field value is mandatory otherwise it is optional",
				Computed:            true,
			},
			"netflow_collector_port": schema.Int64Attribute{
				MarkdownDescription: "If collection type is `TelemetryBrokerOrUDPDirector`, this field value is mandatory otherwise it is optional",
				Computed:            true,
			},
			"enable_netflow_collector_on_devices": schema.BoolAttribute{
				MarkdownDescription: "Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices",
				Computed:            true,
			},
		},
	}
}

func (d *TelemetrySettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TelemetrySettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TelemetrySettings

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?_inherited=true"
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
