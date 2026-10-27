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
	_ datasource.DataSource              = &PowerProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &PowerProfileDataSource{}
)

func NewPowerProfileDataSource() datasource.DataSource {
	return &PowerProfileDataSource{}
}

type PowerProfileDataSource struct {
	client *cc.Client
}

func (d *PowerProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_power_profile"
}

func (d *PowerProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Power Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: "Name of the Power Profile",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the Power Profile",
				Computed:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "Ordered list of rules for the Power Profile",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_type": schema.StringAttribute{
							MarkdownDescription: "Interface Type for the rule",
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Interface ID for the rule",
							Computed:            true,
						},
						"parameter_type": schema.StringAttribute{
							MarkdownDescription: "Parameter Type for the rule",
							Computed:            true,
						},
						"parameter_value": schema.StringAttribute{
							MarkdownDescription: "Parameter Value for the rule",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *PowerProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *PowerProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	applyProviderMeta(d.client, ctx, req.ProviderMeta)
	var config PowerProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?profileName=" + url.QueryEscape(config.ProfileName.ValueString())
	res, err := d.client.Get("/dna/intent/api/v1/wirelessSettings/powerProfiles" + params)
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
