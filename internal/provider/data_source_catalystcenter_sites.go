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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SitesDataSource{}
	_ datasource.DataSourceWithConfigure = &SitesDataSource{}
)

func NewSitesDataSource() datasource.DataSource {
	return &SitesDataSource{}
}

type SitesDataSource struct {
	client *cc.Client
}

func (d *SitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sites"
}

func (d *SitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source fetches all sites defined on the Catalyst Center.\n\n" +
			"Sites represent the logical and physical hierarchy of your network environment, including areas, buildings, and floors.\n" +
			"You can optionally specify the `type` attribute to filter the results and retrieve only sites of a specific type—such as `area`, `building` or `floor`.\n" +
			"To retrieve detailed information about a specific site, use the data source `data.catalystcenter_site`.",

		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "Site type (eg. area, building, floor)",
				Optional:            true,
			},
			"sites": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"name_hierarchy": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"parent_id": schema.StringAttribute{
							MarkdownDescription: "Parent Id",
							Computed:            true,
						},
						"latitude": schema.StringAttribute{
							MarkdownDescription: "Building latitude",
							Computed:            true,
						},
						"longitude": schema.StringAttribute{
							MarkdownDescription: "Building longitude",
							Computed:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: "Building address",
							Computed:            true,
						},
						"country": schema.StringAttribute{
							MarkdownDescription: "Country name for the building",
							Computed:            true,
						},
						"floor_number": schema.Int64Attribute{
							MarkdownDescription: "Floor number",
							Computed:            true,
						},
						"rf_model": schema.StringAttribute{
							MarkdownDescription: "Floor RF model",
							Computed:            true,
						},
						"width": schema.Float64Attribute{
							MarkdownDescription: "Floor width",
							Computed:            true,
						},
						"length": schema.Float64Attribute{
							MarkdownDescription: "Floor length",
							Computed:            true,
						},
						"height": schema.Float64Attribute{
							MarkdownDescription: "Floor height",
							Computed:            true,
						},
						"units_of_measure": schema.StringAttribute{
							MarkdownDescription: "Floor unit of measure",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

const sitesReadMaxAttempts = 4

func (d *SitesDataSource) sitesCount(typeFilter string) (int64, error) {
	path := "/dna/intent/api/v1/sites/count"
	if typeFilter != "" {
		path += "?type=" + typeFilter
	}
	res, err := d.client.Get(path)
	if err != nil {
		return 0, err
	}
	return res.Get("response.count").Int(), nil
}

func dedupeSitesByID(res gjson.Result) (gjson.Result, int64) {
	items := res.Get("response").Array()
	seen := make(map[string]struct{}, len(items))
	uniqueRaw := make([]string, 0, len(items))
	for _, item := range items {
		id := item.Get("id").String()
		if id != "" {
			if _, dup := seen[id]; dup {
				continue
			}
			seen[id] = struct{}{}
		}
		uniqueRaw = append(uniqueRaw, item.Raw)
	}
	return gjson.Parse(`{"response":[` + strings.Join(uniqueRaw, ",") + `]}`), int64(len(uniqueRaw))
}

func (d *SitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Sites

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "sites: Beginning Read")

	typeFilter := ""
	if !config.Type.IsNull() {
		typeFilter = config.Type.ValueString()
	}
	params := ""
	if typeFilter != "" {
		params = "?type=" + typeFilter
	}

	var (
		lastBefore, lastWalk, lastAfter int64 = -1, -1, -1
		lastDeduped                     gjson.Result
	)
	for attempt := 1; attempt <= sitesReadMaxAttempts; attempt++ {
		countBefore, err := d.sitesCount(typeFilter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve site count: %s", err))
			return
		}

		res, err := d.client.Get(config.getPath() + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		deduped, walked := dedupeSitesByID(res)

		countAfter, err := d.sitesCount(typeFilter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve site count: %s", err))
			return
		}

		lastBefore, lastWalk, lastAfter = countBefore, walked, countAfter
		lastDeduped = deduped

		if walked == countBefore && walked == countAfter {
			config.fromBody(ctx, deduped)
			tflog.Debug(ctx, fmt.Sprintf("sites: Read finished successfully on attempt %d, total: %d", attempt, walked))
			diags = resp.State.Set(ctx, &config)
			resp.Diagnostics.Append(diags...)
			return
		}

		tflog.Warn(ctx, fmt.Sprintf("sites: walk count mismatch on attempt %d (countBefore=%d, walked=%d, countAfter=%d) - site set was mutated mid-walk, retrying", attempt, countBefore, walked, countAfter))
	}

	config.fromBody(ctx, lastDeduped)
	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
	resp.Diagnostics.AddWarning(
		"Inconsistent Read - site list may be incomplete",
		fmt.Sprintf("Couldn't get all sites; some sites might be missing.\n\nThe site list changed during every read attempt over %d retries (last counts: before=%d, walked=%d, after=%d). Concurrent site mutation prevented obtaining a consistent snapshot. The data source contains the last walk result; rerun once mutation activity settles for a complete list.",
			sitesReadMaxAttempts, lastBefore, lastWalk, lastAfter),
	)
}
