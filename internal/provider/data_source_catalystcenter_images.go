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
	_ datasource.DataSource              = &ImagesDataSource{}
	_ datasource.DataSourceWithConfigure = &ImagesDataSource{}
)

func NewImagesDataSource() datasource.DataSource {
	return &ImagesDataSource{}
}

type ImagesDataSource struct {
	client *cc.Client
}

func (d *ImagesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_images"
}

func (d *ImagesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches all software images uploaded to the Catalyst Center.\n\n" +
			"Use it to look up an image's `id` by name — useful as input to `catalystcenter_image_distribution`\n" +
			"or `catalystcenter_image_activation` when the image was uploaded out-of-band (not via the\n" +
			"`catalystcenter_image` resource).",

		Attributes: map[string]schema.Attribute{
			"images": schema.ListNestedAttribute{
				MarkdownDescription: "List of software images on the Catalyst Center.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "UUID of the software image. Suitable as input to `catalystcenter_image_distribution` / `_activation`.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Image file name.",
							Computed:            true,
						},
						"family": schema.StringAttribute{
							MarkdownDescription: "Image family.",
							Computed:            true,
						},
						"vendor": schema.StringAttribute{
							MarkdownDescription: "Image vendor.",
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: "Image software version.",
							Computed:            true,
						},
						"application_type": schema.StringAttribute{
							MarkdownDescription: "Third party application type (for non-Cisco images).",
							Computed:            true,
						},
						"image_type": schema.StringAttribute{
							MarkdownDescription: "Image type (e.g. `SYSTEM_SW`).",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ImagesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

func (d *ImagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Images

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
