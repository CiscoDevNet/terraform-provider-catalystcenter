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
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"accounting_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accounting port of RADIUS server").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
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
							MarkdownDescription: helpers.NewAttributeDescription("Password of the Cisco ISE server").AddMutualExclusivityDescription("**Required**: exactly one of `password` and `password_wo` must be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `password_wo` together with `password_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
						},
						"password_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Password of the Cisco ISE server").AddMutualExclusivityDescription("**Required**: exactly one of `password` and `password_wo` must be set.").String,
							Optional:            true,
							WriteOnly:           true,
							Sensitive:           true,
						},
						"password_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `password_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
						"sshkey": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSH key of the Cisco ISE server").AddMutualExclusivityDescription("Only one of `sshkey` and `sshkey_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `sshkey_wo` together with `sshkey_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
						},
						"sshkey_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSH key of the Cisco ISE server").AddMutualExclusivityDescription("Only one of `sshkey` and `sshkey_wo` can be set.").String,
							Optional:            true,
							WriteOnly:           true,
							Sensitive:           true,
						},
						"sshkey_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `sshkey_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
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
				Optional:            true,
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
				MarkdownDescription: helpers.NewAttributeDescription("Shared secret between devices and authentication and policy server").AddMutualExclusivityDescription("**Required**: exactly one of `shared_secret` and `shared_secret_wo` must be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `shared_secret_wo` together with `shared_secret_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"shared_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shared secret between devices and authentication and policy server").AddMutualExclusivityDescription("**Required**: exactly one of `shared_secret` and `shared_secret_wo` must be set.").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"shared_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `shared_secret_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
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
				MarkdownDescription: helpers.NewAttributeDescription("Message key used to encrypt shared secret").AddMutualExclusivityDescription("Only one of `message_key` and `message_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `message_key_wo` together with `message_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"message_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Message key used to encrypt shared secret").AddMutualExclusivityDescription("Only one of `message_key` and `message_key_wo` can be set.").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"message_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `message_key_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"encryption_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encryption key used to encrypt shared secret").AddMutualExclusivityDescription("Only one of `encryption_key` and `encryption_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `encryption_key_wo` together with `encryption_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"encryption_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encryption key used to encrypt shared secret").AddMutualExclusivityDescription("Only one of `encryption_key` and `encryption_key_wo` can be set.").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"encryption_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `encryption_key_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
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

// ValidateConfig enforces the relationship between a secret attribute, its write-only
// "_wo" counterpart and the "_wo_version" rotation trigger.
//
// These checks live here, at resource level, rather than as schema validators. The
// equivalent validators (ConflictsWith, ExactlyOneOf, AlsoRequires) report against an
// attribute path, and Terraform renders an attribute-scoped diagnostic together with the
// offending configuration line - which for a secret prints the value itself into plan
// output and CI logs. A resource-scoped diagnostic is rendered against the resource block
// header instead, so the messages name the attributes explicitly, and identify the list
// element by index for secrets nested inside a list.
func (r *AuthenticationPolicyServerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var legacySharedSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("shared_secret"), &legacySharedSecret)...)
	var woSharedSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("shared_secret_wo"), &woSharedSecret)...)
	var woVersionSharedSecret types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("shared_secret_wo_version"), &woVersionSharedSecret)...)
	if !legacySharedSecret.IsUnknown() && !woSharedSecret.IsUnknown() && !legacySharedSecret.IsNull() && !woSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `shared_secret` and `shared_secret_wo` can be set.",
		)
	}
	if !legacySharedSecret.IsUnknown() && !woSharedSecret.IsUnknown() && legacySharedSecret.IsNull() && woSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Exactly one of `shared_secret` and `shared_secret_wo` must be set.",
		)
	}
	if !woSharedSecret.IsUnknown() && !woVersionSharedSecret.IsUnknown() && !woSharedSecret.IsNull() && woVersionSharedSecret.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`shared_secret_wo_version` must be set when `shared_secret_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	var legacyMessageKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("message_key"), &legacyMessageKey)...)
	var woMessageKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("message_key_wo"), &woMessageKey)...)
	var woVersionMessageKey types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("message_key_wo_version"), &woVersionMessageKey)...)
	if !legacyMessageKey.IsUnknown() && !woMessageKey.IsUnknown() && !legacyMessageKey.IsNull() && !woMessageKey.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `message_key` and `message_key_wo` can be set.",
		)
	}
	if !woMessageKey.IsUnknown() && !woVersionMessageKey.IsUnknown() && !woMessageKey.IsNull() && woVersionMessageKey.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`message_key_wo_version` must be set when `message_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	var legacyEncryptionKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("encryption_key"), &legacyEncryptionKey)...)
	var woEncryptionKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("encryption_key_wo"), &woEncryptionKey)...)
	var woVersionEncryptionKey types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("encryption_key_wo_version"), &woVersionEncryptionKey)...)
	if !legacyEncryptionKey.IsUnknown() && !woEncryptionKey.IsUnknown() && !legacyEncryptionKey.IsNull() && !woEncryptionKey.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `encryption_key` and `encryption_key_wo` can be set.",
		)
	}
	if !woEncryptionKey.IsUnknown() && !woVersionEncryptionKey.IsUnknown() && !woEncryptionKey.IsNull() && woVersionEncryptionKey.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`encryption_key_wo_version` must be set when `encryption_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	{
		// Secrets nested in "cisco_ise_dtos" are validated per element. A list that cannot be
		// read as a whole - because it is still unknown at validation time - is skipped
		// rather than reported, since there is nothing to check yet.
		var cfgCiscoIseDtos []AuthenticationPolicyServerCiscoIseDtos
		if diags := req.Config.GetAttribute(ctx, path.Root("cisco_ise_dtos"), &cfgCiscoIseDtos); !diags.HasError() {
			for i := range cfgCiscoIseDtos {
				if !cfgCiscoIseDtos[i].Password.IsUnknown() && !cfgCiscoIseDtos[i].PasswordWo.IsUnknown() && !cfgCiscoIseDtos[i].Password.IsNull() && !cfgCiscoIseDtos[i].PasswordWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `password` and `password_wo` can be set in `cisco_ise_dtos` element %d.", i),
					)
				}
				if !cfgCiscoIseDtos[i].Password.IsUnknown() && !cfgCiscoIseDtos[i].PasswordWo.IsUnknown() && cfgCiscoIseDtos[i].Password.IsNull() && cfgCiscoIseDtos[i].PasswordWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Exactly one of `password` and `password_wo` must be set in `cisco_ise_dtos` element %d.", i),
					)
				}
				if !cfgCiscoIseDtos[i].PasswordWo.IsUnknown() && !cfgCiscoIseDtos[i].PasswordWoVersion.IsUnknown() && !cfgCiscoIseDtos[i].PasswordWo.IsNull() && cfgCiscoIseDtos[i].PasswordWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`password_wo_version` must be set when `password_wo` is used in `cisco_ise_dtos` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !cfgCiscoIseDtos[i].Sshkey.IsUnknown() && !cfgCiscoIseDtos[i].SshkeyWo.IsUnknown() && !cfgCiscoIseDtos[i].Sshkey.IsNull() && !cfgCiscoIseDtos[i].SshkeyWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `sshkey` and `sshkey_wo` can be set in `cisco_ise_dtos` element %d.", i),
					)
				}
				if !cfgCiscoIseDtos[i].SshkeyWo.IsUnknown() && !cfgCiscoIseDtos[i].SshkeyWoVersion.IsUnknown() && !cfgCiscoIseDtos[i].SshkeyWo.IsNull() && cfgCiscoIseDtos[i].SshkeyWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`sshkey_wo_version` must be set when `sshkey_wo` is used in `cisco_ise_dtos` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
			}
		}
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AuthenticationPolicyServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan AuthenticationPolicyServer

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "shared_secret_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("shared_secret_wo"), &plan.SharedSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "message_key_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("message_key_wo"), &plan.MessageKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "encryption_key_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("encryption_key_wo"), &plan.EncryptionKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only values in list "cisco_ise_dtos" are not stored in plan/state; read the parent list from config and copy them into plan element-by-element.
	{
		var cfgCiscoIseDtos []AuthenticationPolicyServerCiscoIseDtos
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_ise_dtos"), &cfgCiscoIseDtos)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.CiscoIseDtos {
			if i < len(cfgCiscoIseDtos) {
				plan.CiscoIseDtos[i].PasswordWo = cfgCiscoIseDtos[i].PasswordWo
				plan.CiscoIseDtos[i].SshkeyWo = cfgCiscoIseDtos[i].SshkeyWo
			}
		}
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
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
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
func (r *AuthenticationPolicyServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
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
	// Write-only value "shared_secret_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "shared_secret_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("shared_secret_wo"), &plan.SharedSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "message_key_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "message_key_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("message_key_wo"), &plan.MessageKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "encryption_key_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "encryption_key_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("encryption_key_wo"), &plan.EncryptionKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only values in list "cisco_ise_dtos" are not stored in plan/state; read the parent list from config and copy them into plan element-by-element. Read unconditionally for the same reason as the top-level secrets above (full-object replace PUT requires every secret present).
	{
		var cfgCiscoIseDtos []AuthenticationPolicyServerCiscoIseDtos
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_ise_dtos"), &cfgCiscoIseDtos)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.CiscoIseDtos {
			if i < len(cfgCiscoIseDtos) {
				plan.CiscoIseDtos[i].PasswordWo = cfgCiscoIseDtos[i].PasswordWo
				plan.CiscoIseDtos[i].SshkeyWo = cfgCiscoIseDtos[i].SshkeyWo
			}
		}
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
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
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
