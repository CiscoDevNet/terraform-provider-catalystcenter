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
	_ datasource.DataSource              = &ExtranetPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &ExtranetPolicyDataSource{}
)

func NewExtranetPolicyDataSource() datasource.DataSource {
	return &ExtranetPolicyDataSource{}
}

type ExtranetPolicyDataSource struct {
	client *cc.Client
}

func (d *ExtranetPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_extranet_policy"
}

func (d *ExtranetPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Extranet Policy.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"extranet_policy_name": schema.StringAttribute{
				MarkdownDescription: "Name of the extranet policy to be created",
				Required:            true,
			},
			"fabric_ids": schema.SetAttribute{
				MarkdownDescription: "IDs of the fabric sites to be associated with this extranet policy",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"provider_virtual_network_name": schema.StringAttribute{
				MarkdownDescription: "Name of the existing provider virtual network",
				Computed:            true,
			},
			"subscriber_virtual_network_names": schema.SetAttribute{
				MarkdownDescription: "Name of the subscriber virtual networks",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *ExtranetPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ExtranetPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ExtranetPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?extranetPolicyName=" + url.QueryEscape(config.ExtranetPolicyName.ValueString())
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
