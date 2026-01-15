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

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &NetworkDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkDevicesDataSource{}
)

func NewNetworkDevicesDataSource() datasource.DataSource {
	return &NetworkDevicesDataSource{}
}

type NetworkDevicesDataSource struct {
	client *cc.Client
}

func (d *NetworkDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_devices"
}

func (d *NetworkDevicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches all network devices defined on the Catalyst Center.\n\n" +
			"This data source does not report physical location (site) information. Obtain it instead from the data source\n" +
			"`data.catalystcenter_device_detail`. To determine physical locations of multiple devices use the same data\n" +
			"source with `for_each` Terraform meta-argument.",

		Attributes: map[string]schema.Attribute{
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "UUID of the network device",
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname of the network device",
							Computed:            true,
						},
						"management_ip_address": schema.StringAttribute{
							MarkdownDescription: "Management IP address",
							Computed:            true,
						},
						"management_state": schema.StringAttribute{
							MarkdownDescription: "Management state of a network device. If it is not \"Managed\" for a device, then the `catalystcenter_device_detail` may fail for that device.",
							Computed:            true,
						},
						"platform_id": schema.StringAttribute{
							MarkdownDescription: "Platform identifier",
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "Role of the network device, such as `ACCESS` or `DISTRIBUTION`.",
							Computed:            true,
						},
						"software_type": schema.StringAttribute{
							MarkdownDescription: "Type of software",
							Computed:            true,
						},
						"reachability_status": schema.StringAttribute{
							MarkdownDescription: "Device reachability status as `Reachable` or `Unreachable`.",
							Computed:            true,
						},
						"reachability_failure_reason": schema.StringAttribute{
							MarkdownDescription: "Failure reason for unreachable devices.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

func (d *NetworkDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkDevices

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "singleton: Beginning Read")

	params := ""
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, "singleton: Read finished successfully")

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
