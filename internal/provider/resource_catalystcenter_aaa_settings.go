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
var _ resource.Resource = &AAASettingsResource{}
var _ resource.ResourceWithImportState = &AAASettingsResource{}

func NewAAASettingsResource() resource.Resource {
	return &AAASettingsResource{}
}

type AAASettingsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AAASettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_settings"
}

func (r *AAASettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage an AAA Settings.").String,

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
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"network_aaa_server_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of network AAA server").AddStringEnumDescription("AAA", "ISE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AAA", "ISE"),
				},
			},
			"network_aaa_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server protocol").AddStringEnumDescription("RADIUS", "TACACS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("RADIUS", "TACACS"),
				},
			},
			"network_aaa_pan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administration Node. Required for ISE").String,
				Optional:            true,
			},
			"network_aaa_primary_server_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The server to use as a primary").String,
				Optional:            true,
			},
			"network_aaa_secondary_server_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The server to use as a secondary").String,
				Optional:            true,
			},
			"network_aaa_shared_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only relevant for server type `ISE`, shared secret").AddMutualExclusivityDescription("Only one of `network_aaa_shared_secret` and `network_aaa_shared_secret_wo` can be set.").AddDeprecationDescription("The `network_aaa_shared_secret` attribute stores the secret in Terraform state. Use `network_aaa_shared_secret_wo` together with `network_aaa_shared_secret_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"network_aaa_shared_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only relevant for server type `ISE`, shared secret").AddMutualExclusivityDescription("Only one of `network_aaa_shared_secret` and `network_aaa_shared_secret_wo` can be set.").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"network_aaa_shared_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `network_aaa_shared_secret_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"client_aaa_server_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of client AAA server").AddStringEnumDescription("AAA", "ISE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AAA", "ISE"),
				},
			},
			"client_aaa_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Server protocol").AddStringEnumDescription("RADIUS", "TACACS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("RADIUS", "TACACS"),
				},
			},
			"client_aaa_pan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administration Node. Required for ISE").String,
				Optional:            true,
			},
			"client_aaa_primary_server_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The server to use as a primary").String,
				Optional:            true,
			},
			"client_aaa_secondary_server_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The server to use as a secondary").String,
				Optional:            true,
			},
			"client_aaa_shared_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only relevant for server type `ISE`, shared secret").AddMutualExclusivityDescription("Only one of `client_aaa_shared_secret` and `client_aaa_shared_secret_wo` can be set.").AddDeprecationDescription("The `client_aaa_shared_secret` attribute stores the secret in Terraform state. Use `client_aaa_shared_secret_wo` together with `client_aaa_shared_secret_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"client_aaa_shared_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only relevant for server type `ISE`, shared secret").AddMutualExclusivityDescription("Only one of `client_aaa_shared_secret` and `client_aaa_shared_secret_wo` can be set.").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"client_aaa_shared_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `client_aaa_shared_secret_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
		},
	}
}

func (r *AAASettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// ValidateConfig enforces the relationship between a deprecated secret attribute, its
// write-only "_wo" replacement and the "_wo_version" rotation trigger, and raises the
// deprecation warning for the old attribute.
//
// These checks live here, at resource level, rather than as schema validators. The
// equivalent validators (ConflictsWith, ExactlyOneOf, AlsoRequires) report against an
// attribute path, and Terraform renders an attribute-scoped diagnostic together with the
// offending configuration line - which for a secret prints the value itself into plan
// output and CI logs. A resource-scoped diagnostic is rendered against the resource block
// header instead, so the messages name the attributes explicitly.
func (r *AAASettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var legacyNetworkAaaSharedSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("network_aaa_shared_secret"), &legacyNetworkAaaSharedSecret)...)
	var woNetworkAaaSharedSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("network_aaa_shared_secret_wo"), &woNetworkAaaSharedSecret)...)
	var woVersionNetworkAaaSharedSecret types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("network_aaa_shared_secret_wo_version"), &woVersionNetworkAaaSharedSecret)...)
	if !legacyNetworkAaaSharedSecret.IsNull() && !woNetworkAaaSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `network_aaa_shared_secret` and `network_aaa_shared_secret_wo` can be set.",
		)
	}
	if !woNetworkAaaSharedSecret.IsNull() && woVersionNetworkAaaSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`network_aaa_shared_secret_wo_version` must be set when `network_aaa_shared_secret_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !legacyNetworkAaaSharedSecret.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `network_aaa_shared_secret` attribute stores the secret in Terraform state. Use `network_aaa_shared_secret_wo` together with `network_aaa_shared_secret_wo_version` instead, which keeps it out of state.")
	}
	var legacyClientAaaSharedSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("client_aaa_shared_secret"), &legacyClientAaaSharedSecret)...)
	var woClientAaaSharedSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("client_aaa_shared_secret_wo"), &woClientAaaSharedSecret)...)
	var woVersionClientAaaSharedSecret types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("client_aaa_shared_secret_wo_version"), &woVersionClientAaaSharedSecret)...)
	if !legacyClientAaaSharedSecret.IsNull() && !woClientAaaSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `client_aaa_shared_secret` and `client_aaa_shared_secret_wo` can be set.",
		)
	}
	if !woClientAaaSharedSecret.IsNull() && woVersionClientAaaSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`client_aaa_shared_secret_wo_version` must be set when `client_aaa_shared_secret_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !legacyClientAaaSharedSecret.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `client_aaa_shared_secret` attribute stores the secret in Terraform state. Use `client_aaa_shared_secret_wo` together with `client_aaa_shared_secret_wo_version` instead, which keeps it out of state.")
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AAASettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AAASettings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "network_aaa_shared_secret_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("network_aaa_shared_secret_wo"), &plan.NetworkAaaSharedSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "client_aaa_shared_secret_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("client_aaa_shared_secret_wo"), &plan.ClientAaaSharedSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AAASettings{})

	params := ""
	res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AAASettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AAASettings

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?_inherited=true"
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
func (r *AAASettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AAASettings

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
	// Write-only value "network_aaa_shared_secret_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "network_aaa_shared_secret_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("network_aaa_shared_secret_wo"), &plan.NetworkAaaSharedSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "client_aaa_shared_secret_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "client_aaa_shared_secret_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("client_aaa_shared_secret_wo"), &plan.ClientAaaSharedSecretWo)...)
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
func (r *AAASettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AAASettings

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Put(state.getPath(), `{"aaaNetwork": {}, "aaaClient": {}}`, cc.UseMutex)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		errorCode := res.Get("response.errorCode").String()
		if strings.HasPrefix(errorCode, "NCND") {
			// Log a warning and continue execution when NCND**** error is detected
			failureReason := res.Get("response.failureReason").String()
			resp.Diagnostics.AddWarning("Empty input Warning", fmt.Sprintf("Empty input detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "PUT", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *AAASettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <site_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
}

// End of section. //template:end import
