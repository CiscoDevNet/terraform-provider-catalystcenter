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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
var _ resource.Resource = &ProvisionAccessPointsResource{}

func NewProvisionAccessPointsResource() resource.Resource {
	return &ProvisionAccessPointsResource{}
}

type ProvisionAccessPointsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *ProvisionAccessPointsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_provision_access_points"
}

func (r *ProvisionAccessPointsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages the provisioning of Access Points within a single resource, specifying a list of access points to be provisioned in selected site.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_devices": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of access points to be provisioned").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the provisioned AP device").String,
							Optional:            true,
						},
						"device_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the Access Point to be provisioned").String,
							Required:            true,
						},
						"mesh_role": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Mesh Role (Applicable only when AP is in Bridge Mode)").String,
							Optional:            true,
						},
						"reprovision": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the device should be reprovisioned. If set to `true`, reprovisioning will be triggered on every Terraform apply").String,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplace(),
							},
						},
					},
				},
			},
			"rf_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RF Profile Name. RF Profile is not allowed for custom AP Zones.").AddStringEnumDescription("LOW", "TYPICAL", "HIGH").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("LOW", "TYPICAL", "HIGH"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ap_zone_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the site this network device needs to be provisioned").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *ProvisionAccessPointsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ProvisionAccessPointsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProvisionAccessPoints

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ProvisionAccessPoints{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

func (r *ProvisionAccessPointsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProvisionAccessPoints

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	// Iterate through all network devices and reset reprovision flag
	for i := range state.NetworkDevices {
		if !state.NetworkDevices[i].Reprovision.IsNull() && state.NetworkDevices[i].Reprovision.ValueBool() {
			tflog.Debug(ctx, fmt.Sprintf("Resetting reprovision flag for device_id=%s", state.NetworkDevices[i].DeviceId.ValueString()))
			state.NetworkDevices[i].Reprovision = types.BoolValue(false)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: All reprovision flags reset to false", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ProvisionAccessPointsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ProvisionAccessPoints

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *ProvisionAccessPointsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProvisionAccessPoints

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

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
