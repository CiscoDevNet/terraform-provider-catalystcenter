// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://mozilla.org/MPL/2.0/
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
	"regexp"
	"time"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
var _ resource.Resource = &DeployTemplateResource{}

func NewDeployTemplateResource() resource.Resource {
	return &DeployTemplateResource{}
}

type DeployTemplateResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *DeployTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deploy_template"
}

func (r *DeployTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Deploy Template.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"template_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of template to be provisioned").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"redeploy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the template should be redeployed. If set to `true`, template will be redeployed on every Terraform apply").AddStringEnumDescription("ALWAYS", "ON_CHANGE", "NEVER").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ALWAYS", "ON_CHANGE", "NEVER"),
				},
			},
			"force_push_template": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Force Push Template").String,
				Optional:            true,
			},
			"copying_config": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Copy config from running into startup").String,
				Optional:            true,
			},
			"is_composite": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Composite template flag").String,
				Optional:            true,
			},
			"main_template_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Composite Template ID").String,
				Optional:            true,
			},
			"member_template_deployment_info": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Member Template Deployment Info").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Versioned Template ID").String,
							Required:            true,
						},
						"force_push_template": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Force Push Template").String,
							Optional:            true,
						},
						"is_composite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Composite template flag").String,
							Optional:            true,
						},
						"copying_config": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Copy config from running into startup").String,
							Optional:            true,
						},
						"main_template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Template ID").String,
							Optional:            true,
						},
						"target_info": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Target info to deploy template").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME").String,
										Optional:            true,
									},
									"redeploy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the template should be redeployed. If set to `true`, template will be redeployed on every Terraform apply").AddStringEnumDescription("ALWAYS", "ON_CHANGE", "NEVER").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ALWAYS", "ON_CHANGE", "NEVER"),
										},
									},
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ID of device is required if targetType is MANAGED_DEVICE_UUID").String,
										Optional:            true,
									},
									"params": schema.MapAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Template params/values to be provisioned").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"resource_params": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Resource params to be provisioned").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
													},
												},
												"scope": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Scope").String,
													Optional:            true,
												},
												"value": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Value").String,
													Optional:            true,
												},
											},
										},
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
										},
									},
									"versioned_template_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Versioned template ID to be provisioned").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"target_info": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Target info to deploy template").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME").String,
							Optional:            true,
						},
						"redeploy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the template should be redeployed. If set to `true`, template will be redeployed on every Terraform apply").AddStringEnumDescription("ALWAYS", "ON_CHANGE", "NEVER").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ALWAYS", "ON_CHANGE", "NEVER"),
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of device is required if `type` is MANAGED_DEVICE_UUID").String,
							Optional:            true,
						},
						"params": schema.MapAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Template params/values to be provisioned").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"resource_params": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Resource params to be provisioned").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
										},
									},
									"scope": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Scope").String,
										Optional:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Value").String,
										Optional:            true,
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Target type of device").AddStringEnumDescription("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MANAGED_DEVICE_IP", "MANAGED_DEVICE_UUID", "PRE_PROVISIONED_SERIAL", "PRE_PROVISIONED_MAC", "DEFAULT", "MANAGED_DEVICE_HOSTNAME"),
							},
						},
						"versioned_template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Versioned template ID to be provisioned").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DeployTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

func (r *DeployTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeployTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Beginning Create for TemplateId: %s", plan.TemplateId.ValueString()))

	// As per requirement, plan.Id should always be TemplateId
	plan.Id = types.StringValue(fmt.Sprint(plan.TemplateId.ValueString()))

	// Create object body
	body := plan.toBody(ctx, DeployTemplate{})

	for _, v := range plan.TargetInfo {
		tflog.Debug(ctx, fmt.Sprintf("create deploying target id: %s", v.Id))
	}
	postPath := plan.getPath() // params is an empty string in the original code, so just getPath() is sufficient.

	// Perform deployment and monitor status using the new helper
	deploymentSuccess := r.performDeploymentAndMonitorStatus(ctx, postPath, body, &resp.Diagnostics)
	if !deploymentSuccess {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DeployTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeployTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// read always as on_change to create drift and push template every apply
	for i, s := range state.TargetInfo {
		if s.Redeploy.ValueString() == "ALWAYS" {
			state.TargetInfo[i].Redeploy = types.StringValue("NEVER")
		}
	}
	for i, s := range state.MemberTemplateDeploymentInfo {
		for k, d := range s.TargetInfo {
			if d.Redeploy.ValueString() == "ALWAYS" {
				state.MemberTemplateDeploymentInfo[i].TargetInfo[k].Redeploy = types.StringValue("NEVER")
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *DeployTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeployTemplate

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

	// Build list of target_info items that need redeployment
	var targetInfoToRedeploy []DeployTemplateTargetInfo

	// Check for changed or new target_info items
	for _, planTarget := range plan.TargetInfo {
		found := false
		changed := false

		for _, stateTarget := range state.TargetInfo {
			if targetsMatch(planTarget, stateTarget) {
				found = true
				// Check if any attributes changed
				if targetChanged(planTarget, stateTarget) {
					changed = true
				}
				break
			}
		}
		redeploy := "NEVER"
		if !planTarget.Redeploy.IsNull() {
			redeploy = planTarget.Redeploy.ValueString()
		} else if !plan.Redeploy.IsNull() {
			redeploy = plan.Redeploy.ValueString()
		}

		tflog.Debug(ctx, fmt.Sprintf("redeployaa: %s: %v", planTarget.Id.ValueString(), redeploy))
		// Add to redeploy list if new or changed
		if !found || changed && redeploy == "ON_CHANGE" || redeploy == "ALWAYS" {
			targetInfoToRedeploy = append(targetInfoToRedeploy, planTarget)
			if !found {
				tflog.Debug(ctx, fmt.Sprintf("New target_info item detected: %s", planTarget.Id.ValueString()))
			} else {
				tflog.Debug(ctx, fmt.Sprintf("Changed target_info item detected: %s", planTarget.Id.ValueString()))
			}
		}
	}

	// Items removed from state are handled implicitly - they just won't be in the new state
	for _, stateTarget := range state.TargetInfo {
		found := false
		for _, planTarget := range plan.TargetInfo {
			if targetsMatch(planTarget, stateTarget) {
				found = true
				break
			}
		}
		if !found {
			tflog.Debug(ctx, fmt.Sprintf("Removed target_info item (no-op): %s", stateTarget.Id.ValueString()))
		}
	}

	// Deploy to changed/new targets
	if len(targetInfoToRedeploy) > 0 {
		success := r.deployTargets(ctx, &plan, targetInfoToRedeploy, &resp.Diagnostics) // Pass resp.Diagnostics
		if !success {
			return
		}
	} else {
		tflog.Debug(ctx, "No target_info items need redeployment")
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *DeployTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeployTemplate

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

// Helper function to check if two target_info items match (same device)
func targetsMatch(target1, target2 DeployTemplateTargetInfo) bool {
	// Match by id (device UUID) or hostname
	if !target1.Id.IsNull() && !target2.Id.IsNull() {
		return target1.Id.Equal(target2.Id)
	}
	if !target1.HostName.IsNull() && !target2.HostName.IsNull() {
		return target1.HostName.Equal(target2.HostName)
	}
	return false
}

// Helper function to check if a target_info item has changed
func targetChanged(planTarget, stateTarget DeployTemplateTargetInfo) bool {
	// Check if params changed
	if !planTarget.Params.Equal(stateTarget.Params) {
		return true
	}

	// Check if type changed
	if !planTarget.Type.Equal(stateTarget.Type) {
		return true
	}

	// Check if versioned_template_id changed
	if !planTarget.VersionedTemplateId.Equal(stateTarget.VersionedTemplateId) {
		return true
	}

	// Check if resource_params changed
	if len(planTarget.ResourceParams) != len(stateTarget.ResourceParams) {
		return true
	}

	for i := range planTarget.ResourceParams {
		if i >= len(stateTarget.ResourceParams) {
			return true
		}
		if !planTarget.ResourceParams[i].Type.Equal(stateTarget.ResourceParams[i].Type) ||
			!planTarget.ResourceParams[i].Scope.Equal(stateTarget.ResourceParams[i].Scope) ||
			!planTarget.ResourceParams[i].Value.Equal(stateTarget.ResourceParams[i].Value) {
			return true
		}
	}

	return false
}

// New helper function to perform template deployment and monitor its status
func (r *DeployTemplateResource) performDeploymentAndMonitorStatus(ctx context.Context, postPath string, body interface{}, diag *diag.Diagnostics) bool {
	bodyString, ok := body.(string)
	if !ok {
		diag.AddError("Internal Error", "Failed to convert request body to string. The 'toBody' method is expected to return a string for the client.Post method.")
		return false
	}

	res, err := r.client.Post(postPath, bodyString) // Pass the type-asserted string
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to initiate template deployment (%s), got error: %s, %s", "POST", err, res.String()))
		return false
	}

	progress := res.Get("response.progress").String()
	re := regexp.MustCompile(`Template Deployemnt Id:\s*([a-f0-9-]+)`)
	matches := re.FindStringSubmatch(progress)

	if len(matches) == 0 {
		tflog.Warn(ctx, "Deployment Id was not found in response. Assuming immediate success or no deployment to track.")
		return true
	}

	deploymentId := matches[1]
	tflog.Debug(ctx, fmt.Sprintf("Deployment started with ID: %s", deploymentId))

	maxRetries := 30
	statusURL := fmt.Sprintf("/dna/intent/api/v1/template-programmer/template/deploy/status/%s", url.QueryEscape(deploymentId))

	waitingStatuses := map[string]bool{
		"INIT": true,
	}

	for retry := 0; retry < maxRetries; retry++ {
		statusRes, err := r.client.Get(statusURL)
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to retrieve deployment status for Id %s: %s", deploymentId, err))
			time.Sleep(10 * time.Second)
			continue
		}

		status := statusRes.Get("status").String()
		devices := statusRes.Get("devices").String()

		switch {
		case status == "SUCCESS":
			tflog.Debug(ctx, fmt.Sprintf("Template deployment %s finished successfully", deploymentId))
			return true
		case status == "FAILURE":
			diag.AddWarning(
				"Template Deployment Failed",
				fmt.Sprintf("Deployment %s failed with status: %s, on devices: %s", deploymentId, status, devices),
			)
			return false
		case waitingStatuses[status]:
			time.Sleep(10 * time.Second)
			continue
		default:
			diag.AddWarning(
				"Template Deployment Unknown Status",
				fmt.Sprintf("Deployment %s ended with unexpected status: %s", deploymentId, status),
			)
			return false
		}
	}

	// If maxRetries reached without success
	diag.AddWarning(
		"Template Deployment Timeout",
		fmt.Sprintf("Deployment %s did not complete within expected time", deploymentId),
	)
	return false
}

// Helper function to deploy to specific target_info items
func (r *DeployTemplateResource) deployTargets(ctx context.Context, plan *DeployTemplate, targets []DeployTemplateTargetInfo, diag *diag.Diagnostics) bool {
	tempPlan := DeployTemplate{
		TemplateId:        plan.TemplateId,
		ForcePushTemplate: plan.ForcePushTemplate,
		CopyingConfig:     plan.CopyingConfig,
		IsComposite:       plan.IsComposite,
		MainTemplateId:    plan.MainTemplateId,
		TargetInfo:        targets,
	}

	for _, v := range targets {
		tflog.Debug(ctx, fmt.Sprintf("deploying target id: %s", v.Id))
	}

	body := tempPlan.toBody(ctx, DeployTemplate{})
	postPath := plan.getPath() // params is an empty string

	// Use the unified helper for deployment and status monitoring
	deploymentSuccess := r.performDeploymentAndMonitorStatus(ctx, postPath, body, diag)
	return deploymentSuccess
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
