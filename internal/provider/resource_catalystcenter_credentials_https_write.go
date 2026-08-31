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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CredentialsHTTPSWriteResource{}
var _ resource.ResourceWithImportState = &CredentialsHTTPSWriteResource{}

func NewCredentialsHTTPSWriteResource() resource.Resource {
	return &CredentialsHTTPSWriteResource{}
}

type CredentialsHTTPSWriteResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *CredentialsHTTPSWriteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credentials_https_write"
}

func (r *CredentialsHTTPSWriteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Credentials HTTPS Write.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the HTTPS credentials").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Username").String,
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password").AddMutualExclusivityDescription("**Required**: exactly one of `password` and `password_wo` must be set.").AddDeprecationDescription("The `password` attribute stores the secret in Terraform state. Use `password_wo` together with `password_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password").AddMutualExclusivityDescription("**Required**: exactly one of `password` and `password_wo` must be set.").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `password_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("HTTPS port").AddDefaultValueDescription("443").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(443),
			},
		},
	}
}

func (r *CredentialsHTTPSWriteResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
// header instead, so the messages name the attributes explicitly, and identify the list
// element by index for secrets nested inside a list.
func (r *CredentialsHTTPSWriteResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var legacyPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &legacyPassword)...)
	var woPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &woPassword)...)
	var woVersionPassword types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo_version"), &woVersionPassword)...)
	if !legacyPassword.IsNull() && !woPassword.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `password` and `password_wo` can be set.",
		)
	}
	if legacyPassword.IsNull() && woPassword.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Exactly one of `password` and `password_wo` must be set.",
		)
	}
	if !woPassword.IsNull() && woVersionPassword.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`password_wo_version` must be set when `password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !legacyPassword.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `password` attribute stores the secret in Terraform state. Use `password_wo` together with `password_wo_version` instead, which keeps it out of state.")
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CredentialsHTTPSWriteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CredentialsHTTPSWrite

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "password_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &plan.PasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, CredentialsHTTPSWrite{})

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
	plan.Id = types.StringValue(res.Get("response.httpsWrite.#(description==\"" + plan.Description.ValueString() + "\").id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CredentialsHTTPSWriteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CredentialsHTTPSWrite

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
	res = res.Get("response.httpsWrite.#(id==\"" + state.Id.ValueString() + "\")")
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
func (r *CredentialsHTTPSWriteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CredentialsHTTPSWrite

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
	// Write-only value "password_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "password_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &plan.PasswordWo)...)
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
func (r *CredentialsHTTPSWriteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CredentialsHTTPSWrite

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
func (r *CredentialsHTTPSWriteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import
