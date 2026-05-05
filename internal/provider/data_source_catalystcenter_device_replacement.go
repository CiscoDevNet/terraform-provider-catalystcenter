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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceReplacementDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceReplacementDataSource{}
)

func NewDeviceReplacementDataSource() datasource.DataSource {
	return &DeviceReplacementDataSource{}
}

type DeviceReplacementDataSource struct {
	client *cc.Client
}

func (d *DeviceReplacementDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_replacement"
}

func (d *DeviceReplacementDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device Replacement.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"faulty_device_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the faulty device to be replaced",
				Optional:            true,
				Computed:            true,
			},
			"replacement_status": schema.StringAttribute{
				MarkdownDescription: "The replacement status of the device. Use MARKED-FOR-REPLACEMENT to mark the device for replacement.",
				Computed:            true,
			},
			"family": schema.StringAttribute{
				MarkdownDescription: "The device family (computed from the faulty device)",
				Computed:            true,
			},
			"faulty_device_name": schema.StringAttribute{
				MarkdownDescription: "The name of the faulty device (computed from the faulty device)",
				Computed:            true,
			},
			"faulty_device_platform": schema.StringAttribute{
				MarkdownDescription: "The platform of the faulty device (computed from the faulty device)",
				Computed:            true,
			},
			"faulty_device_serial_number": schema.StringAttribute{
				MarkdownDescription: "The serial number of the faulty device (computed from the faulty device)",
				Computed:            true,
			},
			"replacement_device_platform": schema.StringAttribute{
				MarkdownDescription: "The platform of the replacement device",
				Computed:            true,
			},
			"replacement_device_serial_number": schema.StringAttribute{
				MarkdownDescription: "The serial number of the replacement device",
				Computed:            true,
			},
			"neighbour_device_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the neighbour device to create the DHCP server",
				Computed:            true,
			},
			"network_readiness_task_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the network readiness task",
				Computed:            true,
			},
			"workflow_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the replacement workflow",
				Computed:            true,
			},
			"creation_time": schema.Int64Attribute{
				MarkdownDescription: "The creation time of the device replacement entry",
				Computed:            true,
			},
			"replacement_time": schema.Int64Attribute{
				MarkdownDescription: "The replacement time of the device replacement entry",
				Computed:            true,
			},
			"workflow_failed_step": schema.StringAttribute{
				MarkdownDescription: "The failed step of the replacement workflow",
				Computed:            true,
			},
			"readiness_check_task_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the readiness check task",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceReplacementDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("faulty_device_id"),
		),
	}
}

func (d *DeviceReplacementDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (d *DeviceReplacementDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceReplacement

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.FaultyDeviceId.IsNull() {
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res.Get("response"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.FaultyDeviceId.ValueString() == v.Get("faultyDeviceId").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with faultyDeviceId '%v', id: %v", config.Id.String(), config.FaultyDeviceId.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			// Device not found in replacement list — return empty state.
			// This is expected for devices that haven't been marked for replacement.
			tflog.Debug(ctx, fmt.Sprintf("No device replacement entry found for faultyDeviceId: %s, returning empty state", config.FaultyDeviceId.ValueString()))
			diags = resp.State.Set(ctx, &config)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	params := ""
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
