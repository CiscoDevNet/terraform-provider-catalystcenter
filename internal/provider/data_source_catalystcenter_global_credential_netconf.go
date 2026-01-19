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
	_ datasource.DataSource              = &GlobalCredentialNETCONFDataSource{}
	_ datasource.DataSourceWithConfigure = &GlobalCredentialNETCONFDataSource{}
)

func NewGlobalCredentialNETCONFDataSource() datasource.DataSource {
	return &GlobalCredentialNETCONFDataSource{}
}

type GlobalCredentialNETCONFDataSource struct {
	client *cc.Client
}

func (d *GlobalCredentialNETCONFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_global_credential_netconf"
}

func (d *GlobalCredentialNETCONFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Global Credential NETCONF.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"netconf_port": schema.StringAttribute{
				MarkdownDescription: "Netconf port on the device. Valid port should be in the range of 1 to 65535.",
				Computed:            true,
			},
			"comments": schema.StringAttribute{
				MarkdownDescription: "Comments to identify the netconf credential",
				Computed:            true,
			},
			"credential_type": schema.StringAttribute{
				MarkdownDescription: "Credential type to identify the application that uses the netconf credential",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for Netconf Credentials",
				Computed:            true,
			},
		},
	}
}

func (d *GlobalCredentialNETCONFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *GlobalCredentialNETCONFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GlobalCredentialNETCONF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?credentialSubType=NETCONF"
	res, err := d.client.Get("/dna/intent/api/v1/global-credential" + params)
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

// End of section. //template:end read
