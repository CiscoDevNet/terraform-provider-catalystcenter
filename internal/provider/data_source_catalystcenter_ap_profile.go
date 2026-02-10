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
	_ datasource.DataSource              = &APProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &APProfileDataSource{}
)

func NewAPProfileDataSource() datasource.DataSource {
	return &APProfileDataSource{}
}

type APProfileDataSource struct {
	client *cc.Client
}

func (d *APProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ap_profile"
}

func (d *APProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the AP Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"ap_profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the Access Point profile. Max length is 32 characters.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the AP profile. Max length is 241 characters",
				Computed:            true,
			},
			"remote_worker_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS, Forensic Capture Enablement, Rogue Detection and Rogue Containment.",
				Computed:            true,
			},
			"auth_type": schema.StringAttribute{
				MarkdownDescription: "Authentication type used in the AP profile. These settings are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.",
				Computed:            true,
			},
			"dot1x_username": schema.StringAttribute{
				MarkdownDescription: "Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.",
				Computed:            true,
			},
			"dot1x_password": schema.StringAttribute{
				MarkdownDescription: "Password for 802.1X authentication. AP dot1x password length should not exceed 120.",
				Computed:            true,
			},
			"ssh_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if SSH is enabled on the AP. Enable SSH to add credentials for device management.",
				Computed:            true,
			},
			"telnet_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.",
				Computed:            true,
			},
			"management_user_name": schema.StringAttribute{
				MarkdownDescription: "Management username must have a minimum of 1 character and a maximum of 32 characters.",
				Computed:            true,
			},
			"management_password": schema.StringAttribute{
				MarkdownDescription: "Management password for the AP. Length must be 8-120 characters.",
				Computed:            true,
			},
			"management_enable_password": schema.StringAttribute{
				MarkdownDescription: "Enable password for managing the AP. Length must be 8-120 characters.",
				Computed:            true,
			},
			"cdp_state": schema.BoolAttribute{
				MarkdownDescription: "Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.",
				Computed:            true,
			},
			"awips_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if AWIPS is enabled on the AP.",
				Computed:            true,
			},
			"awips_forensic_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.",
				Computed:            true,
			},
			"rogue_detection": schema.BoolAttribute{
				MarkdownDescription: "Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters",
				Computed:            true,
			},
			"rogue_detection_min_rssi": schema.Int64Attribute{
				MarkdownDescription: "Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts",
				Computed:            true,
			},
			"rogue_detection_transient_interval": schema.Int64Attribute{
				MarkdownDescription: "Transient interval for rogue detection. Value should be 0 or from 120 to 1800.",
				Computed:            true,
			},
			"rogue_detection_report_interval": schema.Int64Attribute{
				MarkdownDescription: "Report interval for rogue detection. Value should be in range 10 and 300.",
				Computed:            true,
			},
			"pmf_denial_enabled": schema.BoolAttribute{
				MarkdownDescription: "Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.",
				Computed:            true,
			},
			"mesh_enabled": schema.BoolAttribute{
				MarkdownDescription: "This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.",
				Computed:            true,
			},
			"bridge_group_name": schema.StringAttribute{
				MarkdownDescription: "Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.",
				Computed:            true,
			},
			"backhaul_client_access": schema.BoolAttribute{
				MarkdownDescription: "Indicates if backhaul client access is enabled on the AP.",
				Computed:            true,
			},
			"range": schema.Int64Attribute{
				MarkdownDescription: "Range of the mesh network. Value should be between 150-132000",
				Computed:            true,
			},
			"ghz5_backhaul_data_rates": schema.StringAttribute{
				MarkdownDescription: "5GHz backhaul data rates.",
				Computed:            true,
			},
			"ghz24_backhaul_data_rates": schema.StringAttribute{
				MarkdownDescription: "2.4GHz backhaul data rates.",
				Computed:            true,
			},
			"rap_downlink_backhaul": schema.StringAttribute{
				MarkdownDescription: "Type of downlink backhaul used.",
				Computed:            true,
			},
			"ap_power_profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the existing AP power profile.",
				Computed:            true,
			},
			"power_profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.",
				Computed:            true,
			},
			"scheduler_type": schema.StringAttribute{
				MarkdownDescription: "Type of the scheduler.",
				Computed:            true,
			},
			"scheduler_start_time": schema.StringAttribute{
				MarkdownDescription: "Start time of the duration setting.",
				Computed:            true,
			},
			"scheduler_end_time": schema.StringAttribute{
				MarkdownDescription: "End time of the duration setting.",
				Computed:            true,
			},
			"scheduler_day": schema.StringAttribute{
				MarkdownDescription: "Applies every week on the selected days",
				Computed:            true,
			},
			"scheduler_date": schema.StringAttribute{
				MarkdownDescription: "Start and End date of the duration setting, applicable for MONTHLY schedulers.",
				Computed:            true,
			},
			"country_code": schema.StringAttribute{
				MarkdownDescription: "Country code for the AP profile.",
				Computed:            true,
			},
			"time_zone": schema.StringAttribute{
				MarkdownDescription: "In the Time Zone area, choose one of the following options. Not Configured - APs operate in the UTC time zone. Controller - APs operate in the Cisco Wireless Controller time zone. Delta from Controller - APs operate in the offset time from the wireless controller time zone.",
				Computed:            true,
			},
			"time_zone_offset_hour": schema.Int64Attribute{
				MarkdownDescription: "Enter the hour value (HH). The valid range is from -12 through 14.",
				Computed:            true,
			},
			"time_zone_offset_minutes": schema.Int64Attribute{
				MarkdownDescription: "Enter the minute value (MM). The valid range is from 0 through 59.",
				Computed:            true,
			},
			"client_limit": schema.Int64Attribute{
				MarkdownDescription: "Number of clients. Value should be between 0-1200.",
				Computed:            true,
			},
		},
	}
}

func (d *APProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *APProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config APProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?apProfileName=" + url.QueryEscape(config.ApProfileName.ValueString())
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
