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

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessPreAuthACLDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessPreAuthACLDataSource{}
)

func NewWirelessPreAuthACLDataSource() datasource.DataSource {
	return &WirelessPreAuthACLDataSource{}
}

type WirelessPreAuthACLDataSource struct {
	client *cc.Client
}

func (d *WirelessPreAuthACLDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_pre_auth_acl"
}

func (d *WirelessPreAuthACLDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless Pre Auth ACL.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"acl_list_name": schema.StringAttribute{
				MarkdownDescription: "Pre-Auth ACL List Name.",
				Computed:            true,
			},
			"acl_name": schema.StringAttribute{
				MarkdownDescription: "Pre-Auth ACL Name.",
				Computed:            true,
			},
			"ipv6_acl_enabled": schema.BoolAttribute{
				MarkdownDescription: "When set to true, only IPv6 ACL rules can be configured. When set to false, only IPv4 ACL rules can be configured. Default value of the attribute is false.",
				Computed:            true,
			},
			"include_auto_generated_rules": schema.BoolAttribute{
				MarkdownDescription: "If set to true, Predefined rules will be pushed, which include allowing DHCP, DNS, AAA, and other applicable ports and IPs that are created as part of the default pre-auth rule/ACE for the applicable SSID.",
				Computed:            true,
			},
			"ip_acl_rules": schema.ListNestedAttribute{
				MarkdownDescription: "List of IPv4 or IPv6 ACL rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_address": schema.StringAttribute{
							MarkdownDescription: "An IPv4 or IPv6 source address.",
							Computed:            true,
						},
						"source_subnet_mask_or_prefix": schema.Int64Attribute{
							MarkdownDescription: "Source subnet (IPv4) / Source prefix (IPv6). Value should be in CIDR notation.",
							Computed:            true,
						},
						"destination_address": schema.StringAttribute{
							MarkdownDescription: "An IPv4 or IPv6 destination address.",
							Computed:            true,
						},
						"destination_subnet_mask_or_prefix": schema.Int64Attribute{
							MarkdownDescription: "Destination subnet (IPv4) / Destination prefix (IPv6). Value should be in CIDR notation.",
							Computed:            true,
						},
						"source_ports": schema.StringAttribute{
							MarkdownDescription: "Source port. Required when protocol is TCP or UDP. Accepts a single number (e.g., 100) or a range with a hyphen and no whitespaces (e.g., 100-200). Valid values are between 0 and 65535. When not specified, Catalyst Center defaults this to the full range `0-65535`.",
							Computed:            true,
						},
						"destination_ports": schema.StringAttribute{
							MarkdownDescription: "Destination port. Required when protocol is TCP or UDP. Accepts a single number (e.g., 100) or a range with a hyphen and no whitespaces (e.g., 100-200). Valid values are between 0 and 65535. When not specified, Catalyst Center defaults this to the full range `0-65535`.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "An IPv4 or IPv6 protocol. IPv4 values: ANY, AHP, ESP, GRE, ICMP, IGMP, IP, IPINIP, NOS, OSPF, PCP, PIM, TCP, UDP. IPv6 values: ANY, AHP, ESP, ICMPV6, IPV6, PCP, SCTP, TCP, UDP.",
							Computed:            true,
						},
					},
				},
			},
			"walled_garden_urls": schema.SetAttribute{
				MarkdownDescription: "Walled Garden URLs. List of URLs that are allowed for access before a client device is fully authenticated to the network. Allowed special characters are hyphen(-), underscore(_), dot(.) and asterisk(*).",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *WirelessPreAuthACLDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *WirelessPreAuthACLDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	applyProviderMeta(d.client, ctx, req.ProviderMeta)
	var config WirelessPreAuthACL

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
	res = res.Get("response.#(id==\"" + config.Id.ValueString() + "\")")

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
