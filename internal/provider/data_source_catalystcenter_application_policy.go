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
	_ datasource.DataSource              = &ApplicationPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplicationPolicyDataSource{}
)

func NewApplicationPolicyDataSource() datasource.DataSource {
	return &ApplicationPolicyDataSource{}
}

type ApplicationPolicyDataSource struct {
	client *cc.Client
}

func (d *ApplicationPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_policy"
}

func (d *ApplicationPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Application Policy.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"policy_scope": schema.StringAttribute{
				MarkdownDescription: "Name of the policy. On the controller every sibling object carries this as its `policyScope`, and their names are prefixed with it.",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The sibling group-based policies making up this policy. Modelled as a set because the controller does not preserve ordering.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the sibling policy, conventionally `<policy_scope>_<application set>`, `<policy_scope>_queuing_customization` or `<policy_scope>_global_policy_configuration`.",
							Computed:            true,
						},
						"policy_scope": schema.StringAttribute{
							MarkdownDescription: "Must equal the resource's `policy_scope`",
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Priority of the sibling policy. `100` normally, `4095` when the producer refers to an application scalable group.",
							Computed:            true,
						},
						"delete_policy_status": schema.StringAttribute{
							MarkdownDescription: "Deployment state of the sibling policy",
							Computed:            true,
						},
						"advanced_policy_scope_name": schema.StringAttribute{
							MarkdownDescription: "Name carried by the advanced policy scope, normally the policy scope",
							Computed:            true,
						},
						"site_ids": schema.SetAttribute{
							MarkdownDescription: "Site IDs this sibling policy is deployed to",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ssids": schema.SetAttribute{
							MarkdownDescription: "SSIDs this sibling policy applies to",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"clause_type": schema.StringAttribute{
							MarkdownDescription: "Kind of exclusive contract clause carried by this sibling policy",
							Computed:            true,
						},
						"relevance_level": schema.StringAttribute{
							MarkdownDescription: "Business relevance, when `clause_type` is `BUSINESS_RELEVANCE`",
							Computed:            true,
						},
						"device_removal_behavior": schema.StringAttribute{
							MarkdownDescription: "Device removal behaviour, when `clause_type` is `APPLICATION_POLICY_KNOBS`",
							Computed:            true,
						},
						"host_tracking_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether host tracking is enabled, when `clause_type` is `APPLICATION_POLICY_KNOBS`",
							Computed:            true,
						},
						"queuing_profile_id": schema.StringAttribute{
							MarkdownDescription: "ID of the queuing profile, set on the `<policy_scope>_queuing_customization` sibling policy only.",
							Computed:            true,
						},
						"application_set_id": schema.StringAttribute{
							MarkdownDescription: "ID of the application set this sibling policy applies to, set on `BUSINESS_RELEVANCE` sibling policies.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplicationPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ApplicationPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplicationPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?policyScope=" + url.QueryEscape(config.PolicyScope.ValueString())
	res, err := d.client.Get("/dna/intent/api/v1/app-policy" + params)
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
