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
var _ resource.Resource = &ApplicationPolicyResource{}
var _ resource.ResourceWithImportState = &ApplicationPolicyResource{}

func NewApplicationPolicyResource() resource.Resource {
	return &ApplicationPolicyResource{}
}

type ApplicationPolicyResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *ApplicationPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_policy"
}

func (r *ApplicationPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages a Catalyst Center Application QoS Policy. A single logical policy is not one API object: the controller stores it as a set of sibling group-based policies that all share the same `policy_scope`. Each entry in `items` is one such sibling — typically one per application set (carrying a `BUSINESS_RELEVANCE` clause), plus a `<policy>_queuing_customization` entry carrying the queuing profile reference, and optionally a `<policy>_global_policy_configuration` entry carrying `APPLICATION_POLICY_KNOBS`. A site may only be used by one wired policy (`NCAS10157`). `terraform import` expects the policy scope.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_scope": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the policy. On the controller every sibling object carries this as its `policyScope`, and their names are prefixed with it.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The sibling group-based policies making up this policy. Modelled as a set because the controller does not preserve ordering.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the sibling policy, conventionally `<policy_scope>_<application set>`, `<policy_scope>_queuing_customization` or `<policy_scope>_global_policy_configuration`.").String,
							Required:            true,
						},
						"policy_scope": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Must equal the resource's `policy_scope`").String,
							Required:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority of the sibling policy. `100` normally, `4095` when the producer refers to an application scalable group.").String,
							Optional:            true,
						},
						"delete_policy_status": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Deployment state of the sibling policy").AddStringEnumDescription("NONE", "DELETED", "RESTORED").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NONE", "DELETED", "RESTORED"),
							},
						},
						"advanced_policy_scope_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name carried by the advanced policy scope, normally the policy scope").String,
							Optional:            true,
						},
						"site_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Site IDs this sibling policy is deployed to").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ssids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSIDs this sibling policy applies to").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"clause_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Kind of exclusive contract clause carried by this sibling policy").AddStringEnumDescription("BUSINESS_RELEVANCE", "APPLICATION_POLICY_KNOBS").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BUSINESS_RELEVANCE", "APPLICATION_POLICY_KNOBS"),
							},
						},
						"relevance_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Business relevance, when `clause_type` is `BUSINESS_RELEVANCE`").AddStringEnumDescription("BUSINESS_RELEVANT", "BUSINESS_IRRELEVANT", "DEFAULT").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BUSINESS_RELEVANT", "BUSINESS_IRRELEVANT", "DEFAULT"),
							},
						},
						"device_removal_behavior": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Device removal behaviour, when `clause_type` is `APPLICATION_POLICY_KNOBS`").AddStringEnumDescription("DELETE", "RESTORE", "IGNORE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("DELETE", "RESTORE", "IGNORE"),
							},
						},
						"host_tracking_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether host tracking is enabled, when `clause_type` is `APPLICATION_POLICY_KNOBS`").String,
							Optional:            true,
						},
						"queuing_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the queuing profile, set on the `<policy_scope>_queuing_customization` sibling policy only.").String,
							Optional:            true,
						},
						"application_set_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the application set this sibling policy applies to, set on `BUSINESS_RELEVANCE` sibling policies.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplicationPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ApplicationPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplicationPolicy{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.PolicyScope.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ApplicationPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?policyScope=" + url.QueryEscape(state.PolicyScope.ValueString())
	res, err := r.client.Get("/dna/intent/api/v1/app-policy" + params)
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

// Update and Delete are hand-written in resource_catalystcenter_application_policy_crud.go
// because this API uses a single POST endpoint with createList/updateList/deleteList
// buckets, which the generator cannot express.

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplicationPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <policy_scope>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_scope"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *ApplicationPolicyResource) ReadCache(ctx context.Context, req resource.ReadRequest, state ApplicationPolicy, params string) (cc.Res, error) {
	var err error
	cacheKey := "ApplicationPolicy::"

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
	res, err := r.client.Get("/dna/intent/api/v1/app-policy" + params)
	singleRes := res
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache
