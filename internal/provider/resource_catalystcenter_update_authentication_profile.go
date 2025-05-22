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
var _ resource.Resource = &UpdateAuthenticationProfileResource{}
var _ resource.ResourceWithImportState = &UpdateAuthenticationProfileResource{}

func NewUpdateAuthenticationProfileResource() resource.Resource {
	return &UpdateAuthenticationProfileResource{}
}

type UpdateAuthenticationProfileResource struct {
	client *cc.Client
}

func (r *UpdateAuthenticationProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_update_authentication_profile"
}

func (r *UpdateAuthenticationProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Updates an Authentication Profile. The No Authentication profile cannot be updated.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric this authentication profile is assigned to. To update a global authentication profile, either remove this property or set its value to null.").String,
				Optional:            true,
			},
			"authentication_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The default host authentication template").AddStringEnumDescription("Open Authentication", "Closed Authentication", "Low Impact").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Open Authentication", "Closed Authentication", "Low Impact"),
				},
			},
			"authentication_order": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("First authentication method").AddStringEnumDescription("dot1x", "mac").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dot1x", "mac"),
				},
			},
			"dot1x_to_mab_fallback_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("802.1x Timeout").AddIntegerRangeDescription(3, 120).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 120),
				},
			},
			"wake_on_lan": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wake on LAN").String,
				Required:            true,
			},
			"number_of_hosts": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of hosts").AddStringEnumDescription("Single", "Unlimited").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Single", "Unlimited"),
				},
			},
			"is_bpdu_guard_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to `Closed Authentication`").String,
				Optional:            true,
			},
			"pre_auth_acl_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Pre-Authentication ACL").String,
				Optional:            true,
			},
			"pre_auth_acl_implicit_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Implicit behaviour unless overridden (defaults to `DENY`)").AddStringEnumDescription("PERMIT", "DENY").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("PERMIT", "DENY"),
				},
			},
			"pre_auth_acl_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the Pre-Authentication ACL").String,
				Optional:            true,
			},
			"pre_auth_acl_access_contracts": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access contract list schema. Omitting this property or setting it to null, will reset the property to its default value.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Contract behaviour").AddStringEnumDescription("PERMIT", "DENY").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("PERMIT", "DENY"),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol for the access contract - UDP - TCP - TCP_UDP").String,
							Required:            true,
						},
						"port": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Port for the access contract. The port can only be used once in the Access Contract list. - domain - bootpc - bootps").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *UpdateAuthenticationProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (r *UpdateAuthenticationProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UpdateAuthenticationProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve authenticationProfileId
	params := ""
	params += "?authenticationProfileName=" + url.QueryEscape(plan.AuthenticationProfileName.ValueString())

	if !plan.FabricId.IsUnknown() && !plan.FabricId.IsNull() {
		params += "&fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
	}

	res, err := r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("response.0.id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, UpdateAuthenticationProfile{})

	params = ""
	res, err = r.client.Put(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *UpdateAuthenticationProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UpdateAuthenticationProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?authenticationProfileName=" + url.QueryEscape(state.AuthenticationProfileName.ValueString())

	if !state.FabricId.IsUnknown() && !state.FabricId.IsNull() {
		params += "&fabricId=" + url.QueryEscape(state.FabricId.ValueString())
	} else {
		params += "&isGlobalAuthenticationProfile=true"
	}

	res, err := r.client.Get(state.getPath() + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
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

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *UpdateAuthenticationProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UpdateAuthenticationProfile

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
	res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
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
func (r *UpdateAuthenticationProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UpdateAuthenticationProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

func (r *UpdateAuthenticationProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) == 1 {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("authentication_profile_name"), idParts[0])...)
	} else if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <fabric_id>,<authentication_profile_name>, or <authentication_profile_name>. Got: %q", req.ID),
		)
		return
	} else {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("authentication_profile_name"), idParts[1])...)
	}
}
