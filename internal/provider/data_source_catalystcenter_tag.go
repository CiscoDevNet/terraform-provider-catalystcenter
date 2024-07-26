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
	_ datasource.DataSource              = &TagDataSource{}
	_ datasource.DataSourceWithConfigure = &TagDataSource{}
)

func NewTagDataSource() datasource.DataSource {
	return &TagDataSource{}
}

type TagDataSource struct {
	client *cc.Client
}

func (d *TagDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (d *TagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Tag.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the tag",
				Computed:            true,
			},
			"system_tag": schema.BoolAttribute{
				MarkdownDescription: "true for system created tags, false for user defined tag",
				Computed:            true,
			},
			"dynamic_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Dynamic rules details",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"member_type": schema.StringAttribute{
							MarkdownDescription: "memberType of the tag (e.g. networkdevice, interface)",
							Computed:            true,
						},
						"rule_values": schema.ListAttribute{
							MarkdownDescription: "values of the parameter,Only one of the value or values can be used for the given parameter. (for managementIpAddress e.g. [\"10.197.124.90\",\"10.197.124.91\"])",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"rule_items": schema.ListNestedAttribute{
							MarkdownDescription: "items details, multiple rules can be defined by items",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"operation": schema.StringAttribute{
										MarkdownDescription: "Operation of the rule (e.g. OR,IN,EQ,LIKE,ILIKE,AND)",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Name of the parameter (e.g. managementIpAddress,hostname)",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Value of the parameter (e.g. %10%,%NA%)",
										Computed:            true,
									},
								},
							},
						},
						"rule_operation": schema.StringAttribute{
							MarkdownDescription: "Operation of the rule (e.g. OR,IN,EQ,LIKE,ILIKE,AND)",
							Computed:            true,
						},
						"rule_name": schema.StringAttribute{
							MarkdownDescription: "Name of the parameter (e.g. for interface:portName,adminStatus,speed,status,description. for networkdevice:family,series,hostname,managementIpAddress,groupNameHierarchy,softwareVersion)",
							Computed:            true,
						},
						"rule_value": schema.StringAttribute{
							MarkdownDescription: "Value of the parameter (e.g. for portName:1/0/1,for adminStatus,status:up/down, for speed: any integer value, for description: any valid string, for family:switches, for series:C3650, for managementIpAddress:10.197.124.90, groupNameHierarchy:Global, softwareVersion: 16.9.1)",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *TagDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Tag

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?name=" + url.QueryEscape(config.Name.ValueString())
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
