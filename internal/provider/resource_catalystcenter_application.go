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
var _ resource.Resource = &ApplicationResource{}
var _ resource.ResourceWithImportState = &ApplicationResource{}

func NewApplicationResource() resource.Resource {
	return &ApplicationResource{}
}

type ApplicationResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *ApplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application"
}

func (r *ApplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages a custom Catalyst Center Application. The controller ships roughly 1500 built-in (NBAR) applications which are seeded and must not be managed by this resource. An application always belongs to an application set via `application_set_id`. `terraform import` expects the two-part identifier `<name>,<id>`.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the application").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"application_set_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the application set this application belongs to. Only `idRef` is sent on update; the `id` that `GET` returns alongside it must not be echoed back.").String,
				Required:            true,
			},
			"category_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the application category. The controller exposes no endpoint to list categories, so this opaque UUID must be copied from an existing application in the desired category.").String,
				Required:            true,
			},
			"traffic_class": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The traffic class the application is assigned to").AddStringEnumDescription("BROADCAST_VIDEO", "BULK_DATA", "MULTIMEDIA_CONFERENCING", "MULTIMEDIA_STREAMING", "NETWORK_CONTROL", "OPS_ADMIN_MGMT", "REAL_TIME_INTERACTIVE", "SIGNALING", "TRANSACTIONAL_DATA", "VOIP_TELEPHONY", "BEST_EFFORT", "SCAVENGER").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("BROADCAST_VIDEO", "BULK_DATA", "MULTIMEDIA_CONFERENCING", "MULTIMEDIA_STREAMING", "NETWORK_CONTROL", "OPS_ADMIN_MGMT", "REAL_TIME_INTERACTIVE", "SIGNALING", "TRANSACTIONAL_DATA", "VOIP_TELEPHONY", "BEST_EFFORT", "SCAVENGER"),
				},
			},
			"help_string": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the application shown in the GUI").String,
				Optional:            true,
			},
			"dscp": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DSCP value assigned to traffic matching this application").String,
				Optional:            true,
			},
			"rank": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rank used to break ties between overlapping applications").String,
				Optional:            true,
			},
			"app_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Transport protocol the application runs over").AddStringEnumDescription("TCP", "UDP", "TCP/UDP", "IP").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TCP", "UDP", "TCP/UDP", "IP"),
				},
			},
			"server_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("How the application is identified. `_servername` pairs with `server_name`, `_url` with `url`.").AddStringEnumDescription("_servername", "_url", "_server-ip").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("_servername", "_url", "_server-ip"),
				},
			},
			"server_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server name identifying the application, when `server_type` is `_servername`").String,
				Optional:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("URL identifying the application, when `server_type` is `_url`").String,
				Optional:            true,
			},
			"network_identity": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Optional L3/L4 signatures identifying the application").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol of the signature").AddStringEnumDescription("TCP_OR_UDP", "TCP", "UDP", "IP").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TCP_OR_UDP", "TCP", "UDP", "IP"),
							},
						},
						"ports": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma separated list of discrete ports. A hyphenated range is rejected by the controller; use `lower_port` and `upper_port` for a range. The controller requires this key to be present even when empty, so send an empty string when only a range is used.").String,
							Optional:            true,
						},
						"lower_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Start of a port range, used with `upper_port`. Leave `ports` empty when a range is used. The controller returns `0` when no range is set.").String,
							Optional:            true,
						},
						"upper_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("End of a port range, used with `lower_port`. Leave `ports` empty when a range is used. The controller returns `0` when no range is set.").String,
							Optional:            true,
						},
						"ipv4_subnet": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 subnets matching the application").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
			"instance_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned instance ID, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"display_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned display name, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"instance_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned instance version, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned namespace, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"qualifier": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned qualifier, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scalable_group_external_handle": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned external handle, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_application_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned ID of the nested network application, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_application_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned name of the nested network application, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"application_sub_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned application subtype, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_application_display_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned display name of the nested network application").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"engine_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Classification engine. Custom applications use `6`. The controller accepts an application without it, but the GUI then fails to render the application's IP/Port classifier.").String,
				Optional:            true,
			},
			"popularity": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned popularity, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"selector_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server-assigned selector ID, echoed back on update").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ApplicationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ApplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Application

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Application{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	params += "?name=" + url.QueryEscape(plan.Name.ValueString())
	params += "&attributes=application&limit=500"
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.#(name==\"" + plan.Name.ValueString() + "\").id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ApplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Application

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?name=" + url.QueryEscape(state.Name.ValueString())
	params += "&attributes=application&limit=500"
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	if data := res.Get("response"); !data.Exists() || len(data.Array()) == 0 {
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
func (r *ApplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Application

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
func (r *ApplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Application

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
func (r *ApplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
func (r *ApplicationResource) ReadCache(ctx context.Context, req resource.ReadRequest, state Application, params string) (cc.Res, error) {
	var err error
	cacheKey := "Application::"

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
