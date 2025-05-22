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
	_ datasource.DataSource              = &UpdateAuthenticationProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &UpdateAuthenticationProfileDataSource{}
)

func NewUpdateAuthenticationProfileDataSource() datasource.DataSource {
	return &UpdateAuthenticationProfileDataSource{}
}

type UpdateAuthenticationProfileDataSource struct {
	client *cc.Client
}

func (d *UpdateAuthenticationProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_update_authentication_profile"
}

func (d *UpdateAuthenticationProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Update Authentication Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"authentication_profile_id": schema.StringAttribute{
				MarkdownDescription: "ID of the authentication profile",
				Computed:            true,
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: "ID of the fabric this authentication profile is assigned to. To update a global authentication profile, either remove this property or set its value to null.",
				Optional:            true,
				Computed:            true,
			},
			"authentication_profile_name": schema.StringAttribute{
				MarkdownDescription: "The default host authentication template",
				Required:            true,
			},
			"authentication_order": schema.StringAttribute{
				MarkdownDescription: "First authentication method",
				Computed:            true,
			},
			"dot1x_to_mab_fallback_timeout": schema.Int64Attribute{
				MarkdownDescription: "802.1x Timeout",
				Computed:            true,
			},
			"wake_on_lan": schema.BoolAttribute{
				MarkdownDescription: "Wake on LAN",
				Computed:            true,
			},
			"number_of_hosts": schema.StringAttribute{
				MarkdownDescription: "Number of hosts",
				Computed:            true,
			},
			"is_bpdu_guard_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to `Closed Authentication`",
				Computed:            true,
			},
			"pre_auth_acl_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable/disable Pre-Authentication ACL",
				Computed:            true,
			},
			"pre_auth_acl_implicit_action": schema.StringAttribute{
				MarkdownDescription: "Implicit behaviour unless overridden (defaults to `DENY`)",
				Computed:            true,
			},
			"pre_auth_acl_description": schema.StringAttribute{
				MarkdownDescription: "Description of the Pre-Authentication ACL",
				Computed:            true,
			},
			"pre_auth_acl_access_contracts": schema.SetNestedAttribute{
				MarkdownDescription: "Access contract list schema. Omitting this property or setting it to null, will reset the property to its default value.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							MarkdownDescription: "Contract behaviour",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol for the access contract - UDP - TCP - TCP_UDP",
							Computed:            true,
						},
						"port": schema.StringAttribute{
							MarkdownDescription: "Port for the access contract. The port can only be used once in the Access Contract list. - domain - bootpc - bootps",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *UpdateAuthenticationProfileDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("fabric_id"),
		),
	}
}

func (d *UpdateAuthenticationProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *UpdateAuthenticationProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config UpdateAuthenticationProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.FabricId.IsNull() {
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res.Get("response"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.FabricId.ValueString() == v.Get("fabricId").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with fabricId '%v', id: %v", config.Id.String(), config.FabricId.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with fabricId: %s", config.FabricId.ValueString()))
			return
		}
	}

	params := ""
	params += "?authenticationProfileName=" + url.QueryEscape(config.AuthenticationProfileName.ValueString())
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
