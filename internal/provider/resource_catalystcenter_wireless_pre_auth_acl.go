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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WirelessPreAuthACLResource{}
var _ resource.ResourceWithImportState = &WirelessPreAuthACLResource{}

func NewWirelessPreAuthACLResource() resource.Resource {
	return &WirelessPreAuthACLResource{}
}

type WirelessPreAuthACLResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *WirelessPreAuthACLResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_pre_auth_acl"
}

func (r *WirelessPreAuthACLResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages wireless Pre-Auth ACLs. Pre-Auth ACLs define the traffic permitted before a client is fully authenticated, used by Guest Web-Auth (CWA/EWA) and Enterprise-Posturing SSIDs. Available in Catalyst Center 3.x.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"acl_list_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Pre-Auth ACL List Name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"acl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Pre-Auth ACL Name.").String,
				Required:            true,
			},
			"ipv6_acl_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When set to true, only IPv6 ACL rules can be configured. When set to false, only IPv4 ACL rules can be configured. Default value of the attribute is false.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"include_auto_generated_rules": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If set to true, Predefined rules will be pushed, which include allowing DHCP, DNS, AAA, and other applicable ports and IPs that are created as part of the default pre-auth rule/ACE for the applicable SSID.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ip_acl_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of IPv4 or IPv6 ACL rules.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An IPv4 or IPv6 source address.").String,
							Required:            true,
						},
						"source_subnet_mask_or_prefix": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source subnet (IPv4) / Source prefix (IPv6). Value should be in CIDR notation.").String,
							Required:            true,
						},
						"destination_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An IPv4 or IPv6 destination address.").String,
							Required:            true,
						},
						"destination_subnet_mask_or_prefix": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination subnet (IPv4) / Destination prefix (IPv6). Value should be in CIDR notation.").String,
							Required:            true,
						},
						"source_ports": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source port. Required when protocol is TCP or UDP. Accepts a single number (e.g., 100) or a range with a hyphen and no whitespaces (e.g., 100-200). Valid values are between 0 and 65535. When not specified, Catalyst Center defaults this to the full range `0-65535`.").AddDefaultValueDescription("0-65535").String,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("0-65535"),
						},
						"destination_ports": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination port. Required when protocol is TCP or UDP. Accepts a single number (e.g., 100) or a range with a hyphen and no whitespaces (e.g., 100-200). Valid values are between 0 and 65535. When not specified, Catalyst Center defaults this to the full range `0-65535`.").AddDefaultValueDescription("0-65535").String,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("0-65535"),
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An IPv4 or IPv6 protocol. IPv4 values: ANY, AHP, ESP, GRE, ICMP, IGMP, IP, IPINIP, NOS, OSPF, PCP, PIM, TCP, UDP. IPv6 values: ANY, AHP, ESP, ICMPV6, IPV6, PCP, SCTP, TCP, UDP.").AddStringEnumDescription("ANY", "AHP", "ESP", "GRE", "ICMP", "IGMP", "IP", "IPINIP", "NOS", "OSPF", "PCP", "PIM", "TCP", "UDP", "ICMPV6", "IPV6", "SCTP").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ANY", "AHP", "ESP", "GRE", "ICMP", "IGMP", "IP", "IPINIP", "NOS", "OSPF", "PCP", "PIM", "TCP", "UDP", "ICMPV6", "IPV6", "SCTP"),
							},
						},
					},
				},
			},
			"walled_garden_urls": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Walled Garden URLs. List of URLs that are allowed for access before a client device is fully authenticated to the network. Allowed special characters are hyphen(-), underscore(_), dot(.) and asterisk(*).").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *WirelessPreAuthACLResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *WirelessPreAuthACLResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessPreAuthACL

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessPreAuthACL{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.#(aclListName==\"" + plan.AclListName.ValueString() + "\").id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *WirelessPreAuthACLResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessPreAuthACL

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	res = res.Get("response.#(id==\"" + state.Id.ValueString() + "\")")
	if !res.Exists() {
		resp.State.RemoveResource(ctx)
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *WirelessPreAuthACLResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessPreAuthACL

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	params := ""
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *WirelessPreAuthACLResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessPreAuthACL

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), cc.UseMutex)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *WirelessPreAuthACLResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *WirelessPreAuthACLResource) ReadCache(ctx context.Context, req resource.ReadRequest, state WirelessPreAuthACL, params string) (cc.Res, error) {
	var err error
	cacheKey := "WirelessPreAuthACL::"

	_, cacheSuffix, found := strings.Cut(params, "?")
	queryPart, err := url.ParseQuery(cacheSuffix)
	if err == nil {
		delete(queryPart, "id")
		newQuery := queryPart.Encode()
		cacheSuffix = "?" + newQuery
		cacheKey += cacheSuffix
	}

	cachedValue, found := r.cache.Get(cacheKey)
	if found {
		tflog.Debug(ctx, fmt.Sprintf("hit cache for %s", cacheKey))
		ccRes, ok := cachedValue.(cc.Res)
		if ok {
			return ccRes, nil
		}
		tflog.Info(ctx, fmt.Sprintf("Invalid cache entry type for %s", cacheKey))
		r.cache.Delete(cacheKey)
	}
	res, err := r.client.Get(state.getPath() + params)
	singleRes := res
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache
