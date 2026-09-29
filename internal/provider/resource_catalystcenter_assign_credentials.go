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
	"time"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AssignCredentialsResource{}
var _ resource.ResourceWithImportState = &AssignCredentialsResource{}

func NewAssignCredentialsResource() resource.Resource {
	return &AssignCredentialsResource{}
}

type AssignCredentialsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AssignCredentialsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assign_credentials"
}

func (r *AssignCredentialsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the assigned credentials of a site.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The site ID.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"cli_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the CLI credentials used to access devices at the site.").String,
				Optional:            true,
			},
			"snmp_v2_read_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the SNMPv2c Read credentials.").String,
				Optional:            true,
			},
			"snmp_v2_write_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the SNMPv2c Write credentials.").String,
				Optional:            true,
			},
			"snmp_v3_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the SNMPv3 credentials.").String,
				Optional:            true,
			},
			"https_read_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the HTTP(S) Read credentials.").String,
				Optional:            true,
			},
			"https_write_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the HTTP(S) Write credentials.").String,
				Optional:            true,
			},
			"preserve_unmanaged": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("When managing the Global site's credentials, preserve credential slots that are set outside Terraform instead of unsetting slots that are not present in the configuration. The Global site is the common-settings root and has no parent to inherit from, so a write must include every slot; with this enabled, unspecified slots that are currently assigned on the controller are re-sent (preserved) rather than cleared, while slots the configuration previously managed and then removed are still unset. Has no effect on non-Global sites, which inherit unspecified slots from their parent.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

func (r *AssignCredentialsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// getCurrentGlobalCredentials reads the current credential assignment for the
// Global site. It is used by the preserve_unmanaged fallback in Create/Update
// to keep credential slots that are set on the controller but not managed by
// this configuration. It mirrors the Read function's GET against the global
// common-settings endpoint and parses the six slots via fromBody.
func (r *AssignCredentialsResource) getCurrentGlobalCredentials(ctx context.Context, siteId string) (AssignCredentials, error) {
	var cur AssignCredentials
	res, err := r.client.Get("/api/v1/commonsetting/global/" + url.QueryEscape(siteId))
	if err != nil {
		return cur, err
	}
	cur.fromBody(ctx, res)
	return cur, nil
}

// Custom create: the generator markers are removed so the Global-root fallback
// below can be maintained by hand. Create PUTs the credential body and, if the
// Global site rejects the partial body (NCND01090 on 2.3.7.x, NCND10603 on
// 3.2.2/3.2.3), retries once with a complete six-slot body (unspecified slots sent as
// {} / unset), mirroring the Delete method's Global handling. Transient
// NCND00010 ("Global Settings Save is in progress") errors are then retried.
func (r *AssignCredentialsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AssignCredentials

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AssignCredentials{})

	params := ""
	res, err := r.client.Put(plan.getPath()+params, body)

	// Global-root fallback: the Global site has no parent to inherit from and
	// rejects a partial credential body (NCND01090 on 2.3.7.x, NCND10603 on
	// 3.2.2/3.2.3). Retry once with a complete six-slot body, mirroring the Delete
	// method. With preserve_unmanaged, slots set outside Terraform are kept
	// (see toBodyGlobal); otherwise unspecified slots are unset with {}.
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCND01090" || errorCode == "NCND10603" {
			tflog.Warn(ctx, fmt.Sprintf("%s: %s detected (likely Global site), retrying create with full credential payload", plan.Id.ValueString(), errorCode))
			preserve := plan.PreserveUnmanaged.ValueBool()
			var current *AssignCredentials
			if preserve {
				cur, gerr := r.getCurrentGlobalCredentials(ctx, plan.SiteId.ValueString())
				if gerr != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read current Global credentials for preserve_unmanaged, got error: %s", gerr))
					return
				}
				current = &cur
			}
			body = plan.toBodyGlobal(ctx, AssignCredentials{}, current, preserve)
			res, err = r.client.Put(plan.getPath()+params, body)
		}
	}

	if err != nil {
		retryErrorCodes := []string{"NCND00010"}
		errorCode := res.Get("response.errorCode").String()

		shouldRetry := false
		for _, code := range retryErrorCodes {
			if errorCode == code {
				shouldRetry = true
				break
			}
		}

		if shouldRetry {
			maxWaitTime := time.Duration(r.client.DefaultMaxAsyncWaitTime) * time.Second
			startTime := time.Now()
			retryInterval := 15 * time.Second

			for shouldRetry && time.Since(startTime) < maxWaitTime {
				tflog.Warn(ctx, fmt.Sprintf("%s: Error code %s encountered, waiting %v before retry (elapsed: %v, max: %v)",
					plan.Id.ValueString(), errorCode, retryInterval, time.Since(startTime), maxWaitTime))
				time.Sleep(retryInterval)
				res, err = r.client.Put(plan.getPath()+params, body)

				if err == nil {
					shouldRetry = false
				} else {
					errorCode = res.Get("response.errorCode").String()
					shouldRetry = false
					for _, code := range retryErrorCodes {
						if errorCode == code {
							shouldRetry = true
							break
						}
					}
				}
			}
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of custom create.

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AssignCredentialsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AssignCredentials

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get("/api/v1/commonsetting/global" + params)
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

// Custom update: the generator markers are removed so the Global-root fallback
// below can be maintained by hand. Update PUTs the credential body and, if the
// Global site rejects the partial body (NCND01090 on 2.3.7.x, NCND10603 on
// 3.2.2/3.2.3), retries once with a complete six-slot body (unspecified slots sent as
// {} / unset), mirroring the Delete method's Global handling. Transient
// NCND00010 ("Global Settings Save is in progress") errors are then retried.
func (r *AssignCredentialsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AssignCredentials

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

	// Global-root fallback: the Global site has no parent to inherit from and
	// rejects a partial credential body (NCND01090 on 2.3.7.x, NCND10603 on
	// 3.2.2/3.2.3). Retry once with a complete six-slot body, mirroring the Delete
	// method. With preserve_unmanaged, slots set outside Terraform (never
	// managed by TF, i.e. null in prior state) are kept while slots the config
	// removed are still unset; otherwise all unspecified slots are unset with {}.
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCND01090" || errorCode == "NCND10603" {
			tflog.Warn(ctx, fmt.Sprintf("%s: %s detected (likely Global site), retrying update with full credential payload", plan.Id.ValueString(), errorCode))
			preserve := plan.PreserveUnmanaged.ValueBool()
			var current *AssignCredentials
			if preserve {
				cur, gerr := r.getCurrentGlobalCredentials(ctx, plan.SiteId.ValueString())
				if gerr != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read current Global credentials for preserve_unmanaged, got error: %s", gerr))
					return
				}
				current = &cur
			}
			body = plan.toBodyGlobal(ctx, state, current, preserve)
			res, err = r.client.Put(plan.getPath()+params, body)
		}
	}

	if err != nil {
		retryErrorCodes := []string{"NCND00010"}
		errorCode := res.Get("response.errorCode").String()

		shouldRetry := false
		for _, code := range retryErrorCodes {
			if errorCode == code {
				shouldRetry = true
				break
			}
		}

		if shouldRetry {
			maxWaitTime := time.Duration(r.client.DefaultMaxAsyncWaitTime) * time.Second
			startTime := time.Now()
			retryInterval := 15 * time.Second

			for shouldRetry && time.Since(startTime) < maxWaitTime {
				tflog.Warn(ctx, fmt.Sprintf("%s: Error code %s encountered, waiting %v before retry (elapsed: %v, max: %v)",
					plan.Id.ValueString(), errorCode, retryInterval, time.Since(startTime), maxWaitTime))
				time.Sleep(retryInterval)
				res, err = r.client.Put(plan.getPath()+params, body)

				if err == nil {
					shouldRetry = false
				} else {
					errorCode = res.Get("response.errorCode").String()
					shouldRetry = false
					for _, code := range retryErrorCodes {
						if errorCode == code {
							shouldRetry = true
							break
						}
					}
				}
			}
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of custom update.

// NOTE: Delete is maintained manually (no generator markers) on purpose.
// Destroying an assign_credentials resource must unassign only the credential
// slots that Terraform actually manages at this site (the ones that are non-null
// in state), so that credentials inherited from a parent site are preserved. The
// generated version sends a static body that clears all six slots, which also
// wipes inherited credentials at child sites. Instead we build the body from
// state (reusing toBody with an empty plan) so each managed slot is sent as a
// top-level null (inherit from the parent site) and unmanaged/inherited slots
// are omitted. The global site has no parent to inherit from and rejects a
// partial body with NCND01090 (2.3.7.x) or NCND10603 ("site not found",
// 3.2.2/3.2.3); in that case we retry once with a complete
// six-slot body (via toBodyGlobal), unsetting each slot with an empty object
// ({}, the API's "unset" form). When preserve_unmanaged is set, slots that are
// currently assigned on the controller but were never managed by Terraform are
// re-sent instead of unset, so credentials configured outside Terraform are not
// wiped when the Global assignment is destroyed. Transient NCND00010
// ("Global Settings Save is in progress") errors are retried.
func (r *AssignCredentialsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AssignCredentials

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// Build the clear body from state: only the slots Terraform manages
	// (non-null in state) are unassigned (sent as a top-level null so they
	// inherit from the parent); inherited slots are omitted.
	var empty AssignCredentials
	body := empty.toBody(ctx, state)
	if body == "" {
		// Nothing managed by Terraform to unassign at this site.
		tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully (nothing to unassign)", state.Id.ValueString()))
		resp.State.RemoveResource(ctx)
		return
	}

	res, err := r.client.Put(state.getPath(), body)

	// The global site (no parent to inherit from) rejects a partial/null
	// credential payload with NCND01090 (2.3.7.x) or NCND10603 ("site not
	// found", 3.2.2/3.2.3). Retry once with a complete six-slot body built by
	// toBodyGlobal (with an empty plan): the slots Terraform manages (non-null
	// in state) are unset with {}. Remaining slots are also unset with {} -
	// unless preserve_unmanaged is set, in which case slots that are currently
	// assigned on the controller but were never managed by Terraform (null in
	// state) are re-sent so credentials configured outside Terraform survive the
	// destroy instead of being wiped. Create/Update handle the same two codes.
	if errorCode := res.Get("response.errorCode").String(); err != nil && (errorCode == "NCND01090" || errorCode == "NCND10603") {
		tflog.Warn(ctx, fmt.Sprintf("%s: %s detected (likely Global site), retrying delete with full credential payload", state.Id.ValueString(), errorCode))
		preserve := state.PreserveUnmanaged.ValueBool()
		var current *AssignCredentials
		if preserve {
			cur, gerr := r.getCurrentGlobalCredentials(ctx, state.SiteId.ValueString())
			if gerr != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read current Global credentials for preserve_unmanaged, got error: %s", gerr))
				return
			}
			current = &cur
		}
		body = empty.toBodyGlobal(ctx, state, current, preserve)
		res, err = r.client.Put(state.getPath(), body)
	}

	if err != nil {
		retryErrorCodes := []string{"NCND00010"}
		errorCode := res.Get("response.errorCode").String()

		shouldRetry := false
		for _, code := range retryErrorCodes {
			if errorCode == code {
				shouldRetry = true
				break
			}
		}

		if shouldRetry {
			maxWaitTime := time.Duration(r.client.DefaultMaxAsyncWaitTime) * time.Second
			startTime := time.Now()
			retryInterval := 15 * time.Second

			for shouldRetry && time.Since(startTime) < maxWaitTime {
				tflog.Warn(ctx, fmt.Sprintf("%s: Error code %s encountered, waiting %v before retry (elapsed: %v, max: %v)",
					state.Id.ValueString(), errorCode, retryInterval, time.Since(startTime), maxWaitTime))
				time.Sleep(retryInterval)
				res, err = r.client.Put(state.getPath(), body)

				if err == nil {
					shouldRetry = false
				} else {
					errorCode = res.Get("response.errorCode").String()
					shouldRetry = false
					for _, code := range retryErrorCodes {
						if errorCode == code {
							shouldRetry = true
							break
						}
					}
				}
			}
		}
	}
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

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *AssignCredentialsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <site_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

// End of section. //template:end import
