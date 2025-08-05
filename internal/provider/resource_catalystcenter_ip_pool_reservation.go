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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
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
}

func (r *IPPoolReservationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_pool_reservation"
}

func (r *IPPoolReservationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage an IP Pool Reservation.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The site ID").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the IP pool reservation").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the IP pool reservation").AddStringEnumDescription("Generic", "LAN", "WAN", "management", "service").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Generic", "LAN", "WAN", "management", "service"),
				},
			},
			"ipv6_address_space": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If the value is `false` only IPv4 input are required, otherwise both IPv6 and IPv4 are required").String,
				Optional:            true,
			},
			"ipv4_global_pool": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Global pool address with cidr, example: 175.175.0.0/16").String,
				Required:            true,
			},
			"ipv4_prefix": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If this value is `true`, the `ipv4_prefix_length` attribute must be provided, if it is `false`, the `ipv4_total_host` attribute must be provided").String,
				Required:            true,
			},
			"ipv4_prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 prefix length is required when `ipv4_prefix` value is `true`.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"ipv4_subnet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv4 subnet").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The gateway for the IP pool reservation").String,
				Optional:            true,
			},
			"ipv4_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DHCP Server IPs").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_dns_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DNS Server IPs").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv6_global_pool": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Global pool address with cidr, example: 2001:db8:85a3::/64").String,
				Optional:            true,
			},
			"ipv6_prefix": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If this value is `true`, the `ipv6_prefix_length` attribute must be provided, if it is `false`, the `ipv6_total_host` attribute must be provided").String,
				Optional:            true,
			},
			"ipv6_prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 prefix length is required when `ipv6_prefix` value is `true`.").String,
				Optional:            true,
			},
			"ipv6_subnet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPv6 subnet, for example `2001:db8:85a3:0:100::`").String,
				Optional:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The gateway for the IP pool reservation").String,
				Optional:            true,
			},
			"ipv6_dhcp_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DHCP Server IPs").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv6_dns_servers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DNS Server IPs").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_total_host": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The total number of IPv4 hosts").String,
				Optional:            true,
			},
			"ipv6_total_host": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The total number of IPv6 hosts").String,
				Optional:            true,
			},
			"slaac_support": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SLAAC support").String,
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
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *IPPoolReservationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPPoolReservation

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
	params += "/" + url.QueryEscape(plan.SiteId.ValueString())
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	params += "?siteId=" + url.QueryEscape(plan.SiteId.ValueString()) + "&groupName=" + url.QueryEscape(plan.Name.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.0.id").String())

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
	params += "?siteId=" + url.QueryEscape(state.SiteId.ValueString()) + "&groupName=" + url.QueryEscape(state.Name.ValueString())
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
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
func (r *IPPoolReservationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPPoolReservation

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
	params += "/" + url.QueryEscape(plan.SiteId.ValueString())
	params += "?id=" + url.QueryEscape(plan.Id.ValueString())
	res, err := r.client.Put(plan.getPath()+params, body)
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
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <site_id>,<name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[1])...)
}

// End of section. //template:end import
