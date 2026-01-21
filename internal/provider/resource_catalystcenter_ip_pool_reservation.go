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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &IPPoolReservationResource{}
var _ resource.ResourceWithImportState = &IPPoolReservationResource{}

func NewIPPoolReservationResource() resource.Resource {
	return &IPPoolReservationResource{}
}

type IPPoolReservationResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *IPPoolReservationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_pool_reservation"
}

func (r *IPPoolReservationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages an IP reserve subpool using the new API schema.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"pool_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the reserve IP subpool. Once created, this cannot be changed.").AddStringEnumDescription("Generic", "LAN", "Management", "Service", "WAN").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Generic", "LAN", "Management", "Service", "WAN"),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The id of the non-Global site that this subpool belongs to.").String,
				Required:            true,
			},
			"ipv4_subnet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 IP address component of the CIDR notation for this subnet.").String,
				Required:            true,
			},
			"ipv4_prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 network mask length as a decimal for the CIDR notation of this subnet.").String,
				Required:            true,
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 gateway IP address for this subnet.").String,
				Optional:            true,
			},
			"ipv4_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 DHCP server(s) for this subnet.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_dns_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 DNS server(s) for this subnet.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_total_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The total number of addresses in the IPv4 pool (numeric string).").String,
				Optional:            true,
			},
			"ipv4_unassignable_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of addresses in the IPv4 pool that cannot be assigned (numeric string).").String,
				Optional:            true,
			},
			"ipv4_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of addresses assigned from the IPv4 pool (numeric string).").String,
				Optional:            true,
			},
			"ipv4_default_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of addresses that are assigned from the IPv4 pool by default (numeric string).").String,
				Optional:            true,
			},
			"ipv4_global_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The non-tunnel global pool ID for this IPv4 reserve pool. Once added, this value cannot be changed.").String,
				Required:            true,
			},
			"ipv6_subnet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 IP address component of the CIDR notation for this subnet.").String,
				Optional:            true,
			},
			"ipv6_prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 network mask length as a decimal for the CIDR notation of this subnet.").String,
				Optional:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 gateway IP address for this subnet.").String,
				Optional:            true,
			},
			"ipv6_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 DHCP server(s) for this subnet.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv6_dns_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 DNS server(s) for this subnet.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv6_total_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The total number of addresses in the IPv6 pool (numeric string, up to 128 bits).").String,
				Optional:            true,
			},
			"ipv6_unassignable_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of addresses in the IPv6 pool that cannot be assigned (numeric string).").String,
				Optional:            true,
			},
			"ipv6_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of addresses assigned from the IPv6 pool (numeric string).").String,
				Optional:            true,
			},
			"ipv6_default_assigned_addresses": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of addresses that are assigned from the IPv6 pool by default (numeric string).").String,
				Optional:            true,
			},
			"ipv6_slaac_support": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If the IPv6 prefixLength is 64, this option may be enabled for SLAAC.").String,
				Optional:            true,
			},
			"ipv6_global_pool_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The non-tunnel global pool ID for this IPv6 reserve pool. Once added, this value cannot be changed.").String,
				Optional:            true,
			},
		},
	}
}

func (r *IPPoolReservationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *IPPoolReservationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPPoolReservation
	cacheKey := "IPPoolReservation::"
	r.cache.DeletePattern(cacheKey)

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, IPPoolReservation{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		if r.AllowExistingOnCreate {
			tflog.Info(ctx, fmt.Sprintf("Failed to configure object (%s), got error: %s, %s. allow_existing_on_create is true, beginning update", "POST", err, res.String()))
		} else {
			if strings.Contains(err.Error(), "StatusCode 404") {
				res, err = r.client.Post(plan.getFallbackPath()+params, body)
			}
			// Check if error persists after fallback attempt
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
				return
			}
		}
	}
	params = ""
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.#(name==\"" + plan.Name.ValueString() + "\").id").String())
	if !r.AllowExistingOnCreate {
		tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))
	} else {
		params = ""
		body = plan.toBody(ctx, IPPoolReservation{Id: plan.Id})
		res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Fallback to update existing resource finished successfully", plan.Id.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *IPPoolReservationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPPoolReservation

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	res, err := r.ReadCache(ctx, req, state, params)

	// Try fallback endpoint if primary fails with 404 or 500
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 500")) {
		tflog.Debug(ctx, fmt.Sprintf("%s: Primary endpoint returned 404, trying fallback endpoint", state.Id.ValueString()))
		res, err = r.client.Get(state.getFallbackPath() + params)
	}
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	res = res.Get("response.#(id==\"" + state.Id.ValueString() + "\")")

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
func (r *IPPoolReservationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPPoolReservation
	cacheKey := "IPPoolReservation::"
	r.cache.DeletePattern(cacheKey)

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body)
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
func (r *IPPoolReservationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPPoolReservation
	cacheKey := "IPPoolReservation::"
	r.cache.DeletePattern(cacheKey)

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *IPPoolReservationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *IPPoolReservationResource) ReadCache(ctx context.Context, req resource.ReadRequest, state IPPoolReservation, params string) (cc.Res, error) {
	var err error
	cacheKey := "IPPoolReservation::"

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
