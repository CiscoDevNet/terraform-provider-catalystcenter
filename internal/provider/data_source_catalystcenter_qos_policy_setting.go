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

// Custom (frozen) section: this data source is a singleton whose synthetic "name"
// identifier is marked query_param_no_body and would therefore be dropped from the
// generated schema. It is added here as a computed attribute so the struct and schema
// stay in sync. Kept outside the generator markers so `go generate` preserves it.

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &QoSPolicySettingDataSource{}
	_ datasource.DataSourceWithConfigure = &QoSPolicySettingDataSource{}
)

func NewQoSPolicySettingDataSource() datasource.DataSource {
	return &QoSPolicySettingDataSource{}
}

type QoSPolicySettingDataSource struct {
	client *cc.Client
}

func (d *QoSPolicySettingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_policy_setting"
}

func (d *QoSPolicySettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the QoS Policy Setting.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Fixed resource identifier.",
				Computed:            true,
			},
			"deploy_by_default_on_wired_devices": schema.BoolAttribute{
				MarkdownDescription: "Whether a QoS policy is deployed automatically to a wired network device when it is provisioned. Applies only where the device is assigned to a site that has a QoS policy configured.",
				Computed:            true,
			},
		},
	}
}

func (d *QoSPolicySettingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of custom (frozen) section.

// Custom (frozen) read: the data source has no "name" input (data_source_no_id), so the
// synthetic identifier is defaulted to the fixed value before fromBody derives the id from
// it. Kept outside the generator markers so `go generate` preserves it.
func (d *QoSPolicySettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	applyProviderMeta(d.client, ctx, req.ProviderMeta)
	var config QoSPolicySetting

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Name.IsNull() {
		config.Name = types.StringValue("qos_policy_setting")
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
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

// End of custom (frozen) section.
