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
	_ datasource.DataSource              = &LANAutomationDataSource{}
	_ datasource.DataSourceWithConfigure = &LANAutomationDataSource{}
)

func NewLANAutomationDataSource() datasource.DataSource {
	return &LANAutomationDataSource{}
}

type LANAutomationDataSource struct {
	client *cc.Client
}

func (d *LANAutomationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lan_automation"
}

func (d *LANAutomationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the LAN Automation.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"discovered_device_site_name_hierarchy": schema.StringAttribute{
				MarkdownDescription: "Discovered device site name.",
				Computed:            true,
			},
			"primary_device_management_ip_address": schema.StringAttribute{
				MarkdownDescription: "Primary seed management IP address.",
				Computed:            true,
			},
			"peer_device_management_ip_address": schema.StringAttribute{
				MarkdownDescription: "Secondary seed management IP address.",
				Computed:            true,
			},
			"primary_device_interface_names": schema.SetAttribute{
				MarkdownDescription: "The list of interfaces on primary seed via which the discovered devices are connected.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ip_pools": schema.ListNestedAttribute{
				MarkdownDescription: "The list of IP pools with its name and role.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_pool_name": schema.StringAttribute{
							MarkdownDescription: "Name of the IP pool.",
							Computed:            true,
						},
						"ip_pool_role": schema.StringAttribute{
							MarkdownDescription: "Role of the IP pool.",
							Computed:            true,
						},
					},
				},
			},
			"multicast_enabled": schema.BoolAttribute{
				MarkdownDescription: "To enable underlay native multicast.",
				Computed:            true,
			},
			"host_name_prefix": schema.StringAttribute{
				MarkdownDescription: "Host name prefix which shall be assigned to the discovered device.",
				Computed:            true,
			},
			"host_name_file_id": schema.StringAttribute{
				MarkdownDescription: "File ID of the CSV file containing the host name list.",
				Computed:            true,
			},
			"isis_domain_password": schema.StringAttribute{
				MarkdownDescription: "ISIS domain password.",
				Computed:            true,
			},
			"redistribute_isis_to_bgp": schema.BoolAttribute{
				MarkdownDescription: "Advertise LAN Automation summary route into BGP.",
				Computed:            true,
			},
			"discovery_level": schema.Int64Attribute{
				MarkdownDescription: "Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier",
				Computed:            true,
			},
			"discovery_timeout": schema.Int64Attribute{
				MarkdownDescription: "Discovery timeout in minutes. Until this time, the stop processing will not be triggered.",
				Computed:            true,
			},
			"discovery_devices": schema.SetNestedAttribute{
				MarkdownDescription: "List of specific devices that will be LAN Automated in this session",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"device_serial_number": schema.StringAttribute{
							MarkdownDescription: "Serial number of the device",
							Computed:            true,
						},
						"device_host_name": schema.StringAttribute{
							MarkdownDescription: "Hostname of the device",
							Computed:            true,
						},
						"device_site_name_hierarchy": schema.StringAttribute{
							MarkdownDescription: "Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type",
							Computed:            true,
						},
						"device_management_ip_address": schema.StringAttribute{
							MarkdownDescription: "Management IP Address of the device",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *LANAutomationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *LANAutomationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LANAutomation

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get("/dna/intent/api/v1/lan-automation/status" + params)
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
