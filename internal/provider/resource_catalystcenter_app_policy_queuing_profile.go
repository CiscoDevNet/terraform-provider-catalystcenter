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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &AppPolicyQueuingProfileResource{}
var _ resource.ResourceWithImportState = &AppPolicyQueuingProfileResource{}

func NewAppPolicyQueuingProfileResource() resource.Resource {
	return &AppPolicyQueuingProfileResource{}
}

type AppPolicyQueuingProfileResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AppPolicyQueuingProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_policy_queuing_profile"
}

func (r *AppPolicyQueuingProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages a Catalyst Center Application QoS Queuing Profile. A profile carries a `BANDWIDTH` clause, a `DSCP_CUSTOMIZATION` clause, or both. Built-in profiles such as `CVD_QUEUING_PROFILE` are seeded by the controller and must not be managed by this resource; reference them from an application policy instead. `terraform import` expects the two-part identifier `<name>,<id>` — the profile name followed by the controller-assigned UUID — because the name is the attribute used to match the object on the collection endpoint.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the queuing profile").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the queuing profile").String,
				Optional:            true,
			},
			"clauses": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The clauses carried by this profile. Supply a `BANDWIDTH` clause, a `DSCP_CUSTOMIZATION` clause, or both. The controller does not add a missing clause automatically.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The kind of clause").AddStringEnumDescription("BANDWIDTH", "DSCP_CUSTOMIZATION").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BANDWIDTH", "DSCP_CUSTOMIZATION"),
							},
						},
						"is_common_between_all_interface_speeds": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the same bandwidth allocation applies to every interface speed. When `true` supply a single `ALL` entry in `interface_speed_bandwidth_clauses`. Only valid on a `BANDWIDTH` clause.").String,
							Optional:            true,
						},
						"interface_speed_bandwidth_clauses": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Per-interface-speed bandwidth allocation. Only valid on a `BANDWIDTH` clause.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_speed": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The interface speed this allocation applies to").AddStringEnumDescription("ALL", "HUNDRED_GBPS", "TEN_GBPS", "ONE_GBPS", "HUNDRED_MBPS", "TEN_MBPS", "ONE_MBPS").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ALL", "HUNDRED_GBPS", "TEN_GBPS", "ONE_GBPS", "HUNDRED_MBPS", "TEN_MBPS", "ONE_MBPS"),
										},
									},
									"tc_bandwidth_settings": schema.SetNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Bandwidth percentage per traffic class. All twelve traffic classes must be supplied and the percentages must total 100.").String,
										Required:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"traffic_class": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("The traffic class").AddStringEnumDescription("BROADCAST_VIDEO", "BULK_DATA", "MULTIMEDIA_CONFERENCING", "MULTIMEDIA_STREAMING", "NETWORK_CONTROL", "OPS_ADMIN_MGMT", "REAL_TIME_INTERACTIVE", "SIGNALING", "TRANSACTIONAL_DATA", "VOIP_TELEPHONY", "BEST_EFFORT", "SCAVENGER").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("BROADCAST_VIDEO", "BULK_DATA", "MULTIMEDIA_CONFERENCING", "MULTIMEDIA_STREAMING", "NETWORK_CONTROL", "OPS_ADMIN_MGMT", "REAL_TIME_INTERACTIVE", "SIGNALING", "TRANSACTIONAL_DATA", "VOIP_TELEPHONY", "BEST_EFFORT", "SCAVENGER"),
													},
												},
												"bandwidth_percentage": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Percentage of bandwidth assigned to this traffic class").AddIntegerRangeDescription(0, 100).String,
													Required:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 100),
													},
												},
											},
										},
									},
								},
							},
						},
						"tc_dscp_settings": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DSCP value per traffic class. All twelve traffic classes must be supplied and every DSCP value must be unique across them. Only valid on a `DSCP_CUSTOMIZATION` clause.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"traffic_class": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The traffic class").AddStringEnumDescription("BROADCAST_VIDEO", "BULK_DATA", "MULTIMEDIA_CONFERENCING", "MULTIMEDIA_STREAMING", "NETWORK_CONTROL", "OPS_ADMIN_MGMT", "REAL_TIME_INTERACTIVE", "SIGNALING", "TRANSACTIONAL_DATA", "VOIP_TELEPHONY", "BEST_EFFORT", "SCAVENGER").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("BROADCAST_VIDEO", "BULK_DATA", "MULTIMEDIA_CONFERENCING", "MULTIMEDIA_STREAMING", "NETWORK_CONTROL", "OPS_ADMIN_MGMT", "REAL_TIME_INTERACTIVE", "SIGNALING", "TRANSACTIONAL_DATA", "VOIP_TELEPHONY", "BEST_EFFORT", "SCAVENGER"),
										},
									},
									"dscp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP value assigned to this traffic class").String,
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *AppPolicyQueuingProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AppPolicyQueuingProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AppPolicyQueuingProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AppPolicyQueuingProfile{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
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
	plan.Id = types.StringValue(res.Get("response.#(name==\"" + plan.Name.ValueString() + "\").id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AppPolicyQueuingProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AppPolicyQueuingProfile

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
func (r *AppPolicyQueuingProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AppPolicyQueuingProfile

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
func (r *AppPolicyQueuingProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AppPolicyQueuingProfile

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
func (r *AppPolicyQueuingProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <name>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *AppPolicyQueuingProfileResource) ReadCache(ctx context.Context, req resource.ReadRequest, state AppPolicyQueuingProfile, params string) (cc.Res, error) {
	var err error
	cacheKey := "AppPolicyQueuingProfile::"

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
