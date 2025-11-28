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
var _ resource.Resource = &AuthenticationPolicyServerResource{}
var _ resource.ResourceWithImportState = &AuthenticationPolicyServerResource{}

func NewAuthenticationPolicyServerResource() resource.Resource {
	return &AuthenticationPolicyServerResource{}
}

type AuthenticationPolicyServerResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AuthenticationPolicyServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_policy_server"
}

func (r *AuthenticationPolicyServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage an Authentication Policy Server.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"authentication_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication port of RADIUS server").AddIntegerRangeDescription(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"accounting_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accounting port of RADIUS server").AddIntegerRangeDescription(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"cisco_ise_dtos": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco ISE Server DTOs").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description about the Cisco ISE server").String,
							Optional:            true,
						},
						"fqdn": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Fully-qualified domain name of the Cisco ISE server").String,
							Required:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Password of the Cisco ISE server").String,
							Required:            true,
						},
						"sshkey": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSH key of the Cisco ISE server").String,
							Optional:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address of the Cisco ISE Server").String,
							Required:            true,
						},
						"subscriber_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Subscriber name of the Cisco ISE server").String,
							Required:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User name of the Cisco ISE server").String,
							Required:            true,
						},
					},
				},
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address of authentication and policy server").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"pxgrid_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value true for enable, false for disable. Default value is true").String,
				Optional:            true,
			},
			"use_dnac_cert_for_pxgrid": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value true to use DNAC certificate for Pxgrid. Default value is false").String,
				Optional:            true,
			},
			"is_ise_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value true for Cisco ISE Server. Default value is false").String,
				Optional:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port of TACACS server").AddIntegerRangeDescription(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS").AddStringEnumDescription("TACACS", "RADIUS", "RADIUS_TACACS").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TACACS", "RADIUS", "RADIUS_TACACS"),
				},
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of communication retries between devices and authentication and policy server. The range is from 1 to 3").AddIntegerRangeDescription(1, 3).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"role": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Role of authentication and policy server").AddStringEnumDescription("primary", "secondary").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("primary", "secondary"),
				},
			},
			"shared_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shared secret between devices and authentication and policy server").String,
				Required:            true,
			},
			"timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20").AddIntegerRangeDescription(2, 20).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 20),
				},
			},
			"encryption_scheme": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of encryption scheme for additional security").AddStringEnumDescription("KEYWRAP", "RADSEC").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("KEYWRAP", "RADSEC"),
				},
			},
			"message_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Message key used to encrypt shared secret").String,
				Optional:            true,
			},
			"encryption_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encryption key used to encrypt shared secret").String,
				Optional:            true,
			},
			"external_cisco_ise_ip_addr_dtos": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For future use").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"external_cisco_ise_ip_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"external_ip_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AuthenticationPolicyServerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AuthenticationPolicyServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AuthenticationPolicyServer

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AuthenticationPolicyServer{})

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
	plan.Id = types.StringValue(res.Get("response.#(ipAddress==\"" + plan.IpAddress.ValueString() + "\").instanceUuid").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AuthenticationPolicyServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AuthenticationPolicyServer

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
	res = res.Get("response.#(instanceUuid==\"" + state.Id.ValueString() + "\")")

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
func (r *AuthenticationPolicyServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AuthenticationPolicyServer

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body)
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
func (r *AuthenticationPolicyServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AuthenticationPolicyServer

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
func (r *AuthenticationPolicyServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
