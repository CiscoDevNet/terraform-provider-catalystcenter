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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &AuthenticationPolicyServerDataSource{}
	_ datasource.DataSourceWithConfigure = &AuthenticationPolicyServerDataSource{}
)

func NewAuthenticationPolicyServerDataSource() datasource.DataSource {
	return &AuthenticationPolicyServerDataSource{}
}

type AuthenticationPolicyServerDataSource struct {
	client *cc.Client
}

func (d *AuthenticationPolicyServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_policy_server"
}

func (d *AuthenticationPolicyServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Authentication Policy Server.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"authentication_port": schema.Int64Attribute{
				MarkdownDescription: "Authentication port of RADIUS server",
				Computed:            true,
			},
			"accounting_port": schema.Int64Attribute{
				MarkdownDescription: "Accounting port of RADIUS server",
				Computed:            true,
			},
			"cisco_ise_dtos": schema.ListNestedAttribute{
				MarkdownDescription: "Cisco ISE Server DTOs",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							MarkdownDescription: "Description about the Cisco ISE server",
							Computed:            true,
						},
						"fqdn": schema.StringAttribute{
							MarkdownDescription: "Fully-qualified domain name of the Cisco ISE server",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Password of the Cisco ISE server",
							Computed:            true,
						},
						"sshkey": schema.StringAttribute{
							MarkdownDescription: "SSH key of the Cisco ISE server",
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP Address of the Cisco ISE Server",
							Computed:            true,
						},
						"subscriber_name": schema.StringAttribute{
							MarkdownDescription: "Subscriber name of the Cisco ISE server",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "User name of the Cisco ISE server",
							Computed:            true,
						},
					},
				},
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: "IP address of authentication and policy server",
				Computed:            true,
			},
			"pxgrid_enabled": schema.BoolAttribute{
				MarkdownDescription: "Value true for enable, false for disable. Default value is true",
				Computed:            true,
			},
			"use_dnac_cert_for_pxgrid": schema.BoolAttribute{
				MarkdownDescription: "Value true to use DNAC certificate for Pxgrid. Default value is false",
				Computed:            true,
			},
			"is_ise_enabled": schema.BoolAttribute{
				MarkdownDescription: "Value true for Cisco ISE Server. Default value is false",
				Computed:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Port of TACACS server",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS",
				Computed:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of communication retries between devices and authentication and policy server. The range is from 1 to 3",
				Computed:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "Role of authentication and policy server",
				Computed:            true,
			},
			"shared_secret": schema.StringAttribute{
				MarkdownDescription: "Shared secret between devices and authentication and policy server",
				Computed:            true,
			},
			"timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: "Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20",
				Computed:            true,
			},
			"encryption_scheme": schema.StringAttribute{
				MarkdownDescription: "Type of encryption scheme for additional security",
				Computed:            true,
			},
			"message_key": schema.StringAttribute{
				MarkdownDescription: "Message key used to encrypt shared secret",
				Computed:            true,
			},
			"encryption_key": schema.StringAttribute{
				MarkdownDescription: "Encryption key used to encrypt shared secret",
				Computed:            true,
			},
			"external_cisco_ise_ip_addr_dtos": schema.ListNestedAttribute{
				MarkdownDescription: "For future use",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"external_cisco_ise_ip_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"external_ip_address": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AuthenticationPolicyServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *AuthenticationPolicyServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AuthenticationPolicyServer

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	res = res.Get("response.#(instanceUuid==\"" + config.Id.ValueString() + "\")")

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
