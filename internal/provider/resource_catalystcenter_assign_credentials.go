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

// Section below is generated&owned by "gen/generator.go". //template:begin create
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

// End of section. //template:end create

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

// Section below is generated&owned by "gen/generator.go". //template:begin update
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

// End of section. //template:end update

// NOTE: Delete is maintained manually (no generator markers) on purpose.
// Destroying an assign_credentials resource must unassign only the credential
// slots that Terraform actually manages at this site (the ones that are non-null
// in state), so that credentials inherited from a parent site are preserved. The
// generated version sends a static body that clears all six slots, which also
// wipes inherited credentials at child sites. Instead we build the body from
// state (reusing toBody with an empty plan) so each managed slot is sent as null
// (inherit from the parent site) and unmanaged/inherited slots are omitted. The
// global site has no parent to inherit from and rejects a partial body with
// NCND01090; in that case we retry once with every slot present as an empty
// object ({}, the API's "unset" form), which is correct at the global root since
// there is nothing to inherit. Transient NCND00010 ("Global Settings Save is in
// progress") errors are retried.
func (r *AssignCredentialsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AssignCredentials

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// Build the clear body from state: only the slots Terraform manages (non-null
	// in state) are unassigned (sent as null so they inherit from the parent);
	// inherited slots are omitted.
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
	// credential payload with NCND01090. Retry once with every slot present as an
	// empty object ({}, "unset").
	if err != nil && res.Get("response.errorCode").String() == "NCND01090" {
		tflog.Warn(ctx, fmt.Sprintf("%s: NCND01090 detected (likely Global site), retrying delete with full credential payload", state.Id.ValueString()))
		body = `{"cliCredentialsId":{},"snmpv2cReadCredentialsId":{},"snmpv2cWriteCredentialsId":{},"snmpv3CredentialsId":{},"httpReadCredentialsId":{},"httpWriteCredentialsId":{}}`
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
