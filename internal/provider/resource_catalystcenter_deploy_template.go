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
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
	cache                 *ThreadSafeCache
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
			"deployment_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("In-flight DNAC deployment id, set when an apply times out before SUCCESS/FAILURE. Cleared by Read on SUCCESS.").String,
				Optional:            true,
				Computed:            true,
			},
			"redeploy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Attribute that controls when the template should be redeployed. `ALWAYS` redeploys it on every Terraform apply, `ON_CHANGE` redeploys only when the template’s content changes, and `NEVER` prevents redeployment.").AddStringEnumDescription("ALWAYS", "ON_CHANGE", "NEVER").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ALWAYS", "ON_CHANGE", "NEVER"),
				},
			},
			"deployment_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum time in seconds to wait for the template deployment to reach `SUCCESS` or `FAILURE`. If it is still running when the timeout expires, a warning is reported and the deployment is reconciled on the next apply. Changing this value alone does not trigger a redeployment. Defaults to `300`.").AddIntegerRangeDescription(10, 86400).AddDefaultValueDescription("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 86400),
				},
				Default: int64default.StaticInt64(300),
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
										MarkdownDescription: helpers.NewAttributeDescription("Attribute that controls when the template should be redeployed. `ALWAYS` redeploys it on every Terraform apply, `ON_CHANGE` redeploys only when the template’s content changes, and `NEVER` prevents redeployment.").AddStringEnumDescription("ALWAYS", "ON_CHANGE", "NEVER").String,
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
										ElementType:         types.ListType{ElemType: types.StringType},
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
							ElementType:         types.ListType{ElemType: types.StringType},
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
			"secret_params": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secret parameters for existing deployment targets. Terraform >= 1.11 is required. Secret values must be supplied whenever a target is redeployed; only target references and rotation versions are stored in state. Entries are matched by target and member identity, not list position.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					deploymentSecretParamsListPlanModifier{},
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"target_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Target ID matching an existing target_info.id. Exactly one of target_id and target_host_name must be set.").String,
							Optional:            true,
						},
						"target_host_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname matching an existing target_info.host_name. Exactly one of target_id and target_host_name must be set.").String,
							Optional:            true,
						},
						"member_template_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Composite member template_id containing the target. Omit for a regular or top-level target.").String,
							Optional:            true,
						},
						"params_wo": schema.MapAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Write-only secret parameter values as lists of strings. Values are omitted from Terraform plan and state. Names must not also exist in the target's public params. Every key requires a matching params_wo_versions key. All secrets are resent whenever the target is deployed.").String,
							ElementType:         types.ListType{ElemType: types.StringType},
							Optional:            true,
							WriteOnly:           true,
							Sensitive:           true,
						},
						"params_wo_versions": schema.MapAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Positive integer rotation versions keyed by secret parameter name. Every known secret entry must supply this map with matching keys. Change a version whenever its secret changes. Versions remain in state and are never sent to CatC. Version changes respect the target's redeploy policy.").String,
							ElementType:         types.Int64Type,
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
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

func (r *DeployTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan DeployTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var secrets []DeployTemplateSecretParams
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("secret_params"), &secrets)...)
	config := plan
	config.SecretParams = secrets
	resp.Diagnostics.Append(validateDeploymentSecrets(ctx, config, true, true)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Beginning Create for TemplateId: %s", plan.TemplateId.ValueString()))

	// As per requirement, plan.Id should always be TemplateId
	plan.Id = types.StringValue(fmt.Sprint(plan.TemplateId.ValueString()))

	// Create object body
	body := deploymentBody(ctx, plan, secrets, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, v := range plan.TargetInfo {
		tflog.Debug(ctx, fmt.Sprintf("create deploying target id: %s", v.Id))
	}
	postPath := plan.getPath() // params is an empty string in the original code, so just getPath() is sufficient.

	// Perform deployment and monitor status using the new helper
	success, deploymentId, _ := r.performDeploymentAndMonitorStatus(ctx, postPath, body, plan.DeploymentTimeout, &resp.Diagnostics, len(secrets) > 0)
	if !success {
		if deploymentId != "" {
			// Persist deployment_id so the next Read can reconcile (TIMEOUT, FAILURE, or unknown status).
			plan.DeploymentId = types.StringValue(deploymentId)
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
		}
		return
	}

	plan.DeploymentId = types.StringNull()

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DeployTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state DeployTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Reconcile a prior timed-out deployment by re-polling its status.
	if !state.DeploymentId.IsNull() && state.DeploymentId.ValueString() != "" {
		deploymentId := state.DeploymentId.ValueString()
		statusURL := fmt.Sprintf("/dna/intent/api/v1/template-programmer/template/deploy/status/%s", url.QueryEscape(deploymentId))
		statusRes, err := r.client.Get(statusURL, cc.NoLogPayload)
		if err != nil {
			// Id no longer queryable; assume success optimistically.
			tflog.Warn(ctx, fmt.Sprintf("Deployment %s no longer queryable (%s); assuming success", deploymentId, err))
			state.DeploymentId = types.StringNull()
		} else {
			status := statusRes.Get("status").String()
			switch status {
			case "SUCCESS":
				state.DeploymentId = types.StringNull()
			case "FAILURE":
				revertVersionedTemplateIds(&state)
				resp.Diagnostics.AddWarning(
					"Previous Template Deployment Failed",
					fmt.Sprintf("Deployment %s reported FAILURE. State reverted to force redeploy.", deploymentId),
				)
			case "INIT", "IN_PROGRESS", "":
				// Still in flight; leave state and id for next Read.
			default:
				revertVersionedTemplateIds(&state)
				resp.Diagnostics.AddWarning(
					"Previous Template Deployment Ended With Unexpected Status",
					fmt.Sprintf("Deployment %s reported an unexpected status. API details are omitted because they may echo template parameters. State reverted to force redeploy.", deploymentId),
				)
			}
		}
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
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
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

	var secrets []DeployTemplateSecretParams
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("secret_params"), &secrets)...)
	config := plan
	config.SecretParams = secrets
	resp.Diagnostics.Append(validateDeploymentSecrets(ctx, config, true, true)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	targetsToUse := deploymentTargets(plan, state)
	if len(targetsToUse) > 0 {
		success, deploymentId, _ := r.deployTargets(ctx, &plan, targetsToUse, secrets, &resp.Diagnostics)
		if !success {
			if deploymentId != "" {
				// Persist deployment_id so the next Read can reconcile (TIMEOUT, FAILURE, or unknown status).
				plan.DeploymentId = types.StringValue(deploymentId)
				resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			}
			return
		}
		plan.DeploymentId = types.StringNull()
	} else {
		tflog.Debug(ctx, "No target_info items need redeployment")
		plan.DeploymentId = state.DeploymentId
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *DeployTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
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

// Helper function to check if a member template target_info item has changed
func memberTargetChanged(planTarget, stateTarget DeployTemplateMemberTemplateDeploymentInfoTargetInfo) bool {
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

// revertVersionedTemplateIds clears versioned_template_id on all targets and
// secret rotation versions and deployment_id, so the next plan diffs and Update retries.
func revertVersionedTemplateIds(state *DeployTemplate) {
	for i := range state.TargetInfo {
		state.TargetInfo[i].VersionedTemplateId = types.StringNull()
	}
	for i := range state.MemberTemplateDeploymentInfo {
		for k := range state.MemberTemplateDeploymentInfo[i].TargetInfo {
			state.MemberTemplateDeploymentInfo[i].TargetInfo[k].VersionedTemplateId = types.StringNull()
		}
	}
	for i := range state.SecretParams {
		state.SecretParams[i].ParamsWoVersions = types.MapNull(types.Int64Type)
	}
	state.DeploymentId = types.StringNull()
}

const (
	defaultDeploymentTimeoutSeconds = 300
	deploymentPollIntervalSeconds   = 10
)

// performDeploymentAndMonitorStatus runs the deploy POST and polls its status.
// Returns (success, deploymentId, terminalStatus) where terminalStatus is
// "SUCCESS", "TIMEOUT" (warning, caller should persist id), or "FAILURE"/other
// (Error already added to diag).
func (r *DeployTemplateResource) performDeploymentAndMonitorStatus(ctx context.Context, postPath string, body interface{}, timeout types.Int64, diag *diag.Diagnostics, sensitive bool) (bool, string, string) {
	bodyString, ok := body.(string)
	if !ok {
		diag.AddError("Internal Error", "Failed to convert request body to string. The 'toBody' method is expected to return a string for the client.Post method.")
		return false, "", ""
	}

	// This resource monitors deployment status itself. Skip the client's generic
	// task polling, which logs unstructured task details that may contain secrets.
	res, err := r.client.Post(postPath, bodyString, cc.NoLogPayload, cc.NoWait)
	if err != nil {
		if sensitive {
			diag.AddError("Client Error", "Failed to initiate template deployment (POST). API error details are omitted because they may echo secret parameters.")
		} else {
			diag.AddError("Client Error", fmt.Sprintf("Failed to initiate template deployment (%s), got error: %s, %s", "POST", err, res.String()))
		}
		return false, "", ""
	}

	timeoutSeconds := int64(defaultDeploymentTimeoutSeconds)
	if !timeout.IsNull() && !timeout.IsUnknown() {
		timeoutSeconds = timeout.ValueInt64()
	}
	deadline := time.Now().Add(time.Duration(timeoutSeconds) * time.Second)

	// The v2 endpoint first returns an asynchronous task. Its completed response
	// contains the deployment ID. Poll it here instead of using the client's task
	// monitor, which logs progress and failure details that can echo secrets.
	taskID := res.Get("response.taskId").String()
	if taskID != "" {
		if uuid.Validate(taskID) != nil {
			diag.AddError("Template Deployment Task Error", "The controller returned an invalid task ID.")
			return false, "", ""
		}
		for {
			taskRes, taskErr := r.client.Get("/api/v1/task/"+url.QueryEscape(taskID), cc.NoLogPayload)
			if taskErr != nil {
				diag.AddError("Template Deployment Task Error", fmt.Sprintf("Unable to retrieve deployment task %s. API details are omitted because they may echo template parameters.", taskID))
				return false, "", ""
			}
			if taskRes.Get("response.isError").Bool() {
				diag.AddError("Template Deployment Task Failed", fmt.Sprintf("Deployment task %s failed. API details are omitted because they may echo template parameters.", taskID))
				return false, "", ""
			}
			if taskRes.Get("response.endTime").Int() > 0 {
				res = taskRes
				break
			}
			remaining := time.Until(deadline)
			if remaining <= 0 {
				diag.AddError("Template Deployment Task Timeout", fmt.Sprintf("Deployment task %s did not complete within %d seconds. Check the controller task before retrying.", taskID, timeoutSeconds))
				return false, "", ""
			}
			select {
			case <-ctx.Done():
				diag.AddError("Template Deployment Cancelled", "Waiting for the deployment task was cancelled.")
				return false, "", ""
			case <-time.After(min(time.Second, remaining)):
			}
		}
	}

	progress := res.Get("response.progress").String()
	re := regexp.MustCompile(`Template Deployemnt Id:\s*([a-f0-9-]+)`)
	matches := re.FindStringSubmatch(progress)

	if len(matches) == 0 || uuid.Validate(matches[1]) != nil {
		if taskID != "" {
			diag.AddError("Template Deployment ID Missing", "The completed task did not contain a valid deployment ID. API details are omitted because they may echo template parameters.")
			return false, "", ""
		}
		tflog.Warn(ctx, "Deployment Id was not found in response. Assuming immediate success or no deployment to track.")
		return true, "", "SUCCESS"
	}

	deploymentId := matches[1]
	tflog.Debug(ctx, fmt.Sprintf("Deployment started with ID: %s", deploymentId))

	statusURL := fmt.Sprintf("/dna/intent/api/v1/template-programmer/template/deploy/status/%s", url.QueryEscape(deploymentId))

	waitingStatuses := map[string]bool{
		"INIT":        true,
		"IN_PROGRESS": true,
	}

	for {
		statusRes, err := r.client.Get(statusURL, cc.NoLogPayload)
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to retrieve deployment status for Id %s: %s", deploymentId, err))
		} else {
			status := statusRes.Get("status").String()

			switch {
			case status == "SUCCESS":
				tflog.Debug(ctx, fmt.Sprintf("Template deployment %s finished successfully", deploymentId))
				return true, deploymentId, "SUCCESS"
			case status == "FAILURE":
				diag.AddWarning(
					"Template Deployment Failed",
					fmt.Sprintf("Deployment %s failed with status: %s. State persisted with deployment_id; next apply will reconcile via Read and retry.", deploymentId, status),
				)
				return false, deploymentId, "FAILURE"
			case waitingStatuses[status]:
			default:
				diag.AddWarning(
					"Template Deployment Unknown Status",
					fmt.Sprintf("Deployment %s ended with an unexpected status. API details are omitted because they may echo template parameters. State persisted with deployment_id; next apply will reconcile via Read.", deploymentId),
				)
				return false, deploymentId, "UNKNOWN"
			}
		}

		// Never sleep past the deadline, so the last status check happens at it.
		remaining := time.Until(deadline)
		if remaining <= 0 {
			break
		}
		time.Sleep(min(deploymentPollIntervalSeconds*time.Second, remaining))
	}

	// Deadline reached: deploy may still be in flight on DNAC. Warning only;
	// caller persists deploymentId so Read can reconcile on the next apply.
	diag.AddWarning(
		"Template Deployment Timeout",
		fmt.Sprintf("Deployment %s did not complete within %d seconds. Will be reconciled on the next apply via deployment_id. Increase deployment_timeout for long-running deployments.", deploymentId, timeoutSeconds),
	)
	return false, deploymentId, "TIMEOUT"
}

// deployTargets deploys to a subset of target_info items.
// Returns (success, deploymentId, terminalStatus) — see performDeploymentAndMonitorStatus.
func (r *DeployTemplateResource) deployTargets(ctx context.Context, plan *DeployTemplate, targets []DeployTemplateTargetInfo, secrets []DeployTemplateSecretParams, diag *diag.Diagnostics) (bool, string, string) {
	var filteredMemberInfo []DeployTemplateMemberTemplateDeploymentInfo
	for _, member := range plan.MemberTemplateDeploymentInfo {
		var filteredTargets []DeployTemplateMemberTemplateDeploymentInfoTargetInfo
		for _, t := range member.TargetInfo {
			for _, selected := range targets {
				if targetsMatch(selected, DeployTemplateTargetInfo{Id: t.Id, HostName: t.HostName}) {
					filteredTargets = append(filteredTargets, t)
					break
				}
			}
		}
		if len(filteredTargets) > 0 {
			memberCopy := DeployTemplateMemberTemplateDeploymentInfo{
				TemplateId:        member.TemplateId,
				ForcePushTemplate: member.ForcePushTemplate,
				IsComposite:       member.IsComposite,
				CopyingConfig:     member.CopyingConfig,
				MainTemplateId:    member.MainTemplateId,
				TargetInfo:        filteredTargets,
			}
			filteredMemberInfo = append(filteredMemberInfo, memberCopy)
		}
	}

	tempPlan := DeployTemplate{
		TemplateId:                   plan.TemplateId,
		ForcePushTemplate:            plan.ForcePushTemplate,
		CopyingConfig:                plan.CopyingConfig,
		IsComposite:                  plan.IsComposite,
		MainTemplateId:               plan.MainTemplateId,
		MemberTemplateDeploymentInfo: filteredMemberInfo,
		TargetInfo:                   targets,
	}

	for _, v := range targets {
		tflog.Debug(ctx, fmt.Sprintf("deploying target id: %s", v.Id))
	}

	body := deploymentBody(ctx, tempPlan, secrets, diag)
	if diag.HasError() {
		return false, "", ""
	}
	postPath := plan.getPath() // params is an empty string

	// Use the unified helper for deployment and status monitoring
	return r.performDeploymentAndMonitorStatus(ctx, postPath, body, plan.DeploymentTimeout, diag, len(secrets) > 0)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
