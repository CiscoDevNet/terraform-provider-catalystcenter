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
var _ resource.Resource = &DeviceResource{}
var _ resource.ResourceWithImportState = &DeviceResource{}

func NewDeviceResource() resource.Resource {
	return &DeviceResource{}
}

type DeviceResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *DeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (r *DeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Device.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cli_transport": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI transport").AddStringEnumDescription("telnet", "ssh").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("telnet", "ssh"),
				},
			},
			"compute_device": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compute device").String,
				Optional:            true,
			},
			"enable_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI enable password of the device").AddDeprecationDescription("The `enable_password` attribute stores the secret in Terraform state. Use `enable_password_wo` together with `enable_password_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"enable_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI enable password of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("enable_password_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("enable_password")),
				},
			},
			"enable_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `enable_password_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"extended_discovery_info": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("This field holds that info as whether to add device with canned data or not.").AddStringEnumDescription("DISCOVER_WITH_CANNED_DATA").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DISCOVER_WITH_CANNED_DATA"),
				},
			},
			"http_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("HTTP password of the device").AddDeprecationDescription("The `http_password` attribute stores the secret in Terraform state. Use `http_password_wo` together with `http_password_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"http_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("HTTP password of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("http_password_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("http_password")),
				},
			},
			"http_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `http_password_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"http_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("HTTP port of the device").String,
				Optional:            true,
			},
			"http_secure": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable HTTPS").String,
				Optional:            true,
			},
			"http_user_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("HTTP username of the device").String,
				Optional:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address of the device").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"meraki_org_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Selected Meraki organizations for which the devices needs to be imported").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"netconf_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NETCONF port of the device").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI password of the device").AddDeprecationDescription("The `password` attribute stores the secret in Terraform state. Use `password_wo` together with `password_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI password of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("password_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("password")),
				},
			},
			"password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `password_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Serial number of the device").String,
				Optional:            true,
			},
			"snmp_auth_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 authentication passphrase of the device").AddDeprecationDescription("The `snmp_auth_passphrase` attribute stores the secret in Terraform state. Use `snmp_auth_passphrase_wo` together with `snmp_auth_passphrase_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"snmp_auth_passphrase_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 authentication passphrase of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("snmp_auth_passphrase_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("snmp_auth_passphrase")),
				},
			},
			"snmp_auth_passphrase_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `snmp_auth_passphrase_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"snmp_auth_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 authentication protocol of the device").AddStringEnumDescription("sha", "md5").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha", "md5"),
				},
			},
			"snmp_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 mode of the device").AddStringEnumDescription("noAuthnoPriv", "authNoPriv", "authPriv").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("noAuthnoPriv", "authNoPriv", "authPriv"),
				},
			},
			"snmp_priv_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 privacy passphrase of the device").AddDeprecationDescription("The `snmp_priv_passphrase` attribute stores the secret in Terraform state. Use `snmp_priv_passphrase_wo` together with `snmp_priv_passphrase_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"snmp_priv_passphrase_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 privacy passphrase of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("snmp_priv_passphrase_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("snmp_priv_passphrase")),
				},
			},
			"snmp_priv_passphrase_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `snmp_priv_passphrase_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"snmp_priv_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 privacy protocol of the device").AddStringEnumDescription("AES128").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AES128"),
				},
			},
			"snmp_ro_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv2 read-only community of the device").AddDeprecationDescription("The `snmp_ro_community` attribute stores the secret in Terraform state. Use `snmp_ro_community_wo` together with `snmp_ro_community_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"snmp_ro_community_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv2 read-only community of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("snmp_ro_community_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("snmp_ro_community")),
				},
			},
			"snmp_ro_community_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `snmp_ro_community_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"snmp_rw_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv2 read-write community of the device").AddDeprecationDescription("The `snmp_rw_community` attribute stores the secret in Terraform state. Use `snmp_rw_community_wo` together with `snmp_rw_community_wo_version` instead, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"snmp_rw_community_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv2 read-write community of the device").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("snmp_rw_community_wo_version")),
					stringvalidator.ConflictsWith(path.MatchRoot("snmp_rw_community")),
				},
			},
			"snmp_rw_community_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `snmp_rw_community_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"snmp_retry": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP retry count").AddIntegerRangeDescription(0, 3).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3),
				},
			},
			"snmp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP timeout in seconds").AddIntegerRangeDescription(0, 300).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 300),
				},
			},
			"snmp_user_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMPv3 username of the device").String,
				Optional:            true,
			},
			"snmp_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP version of the device").AddStringEnumDescription("v2", "v3").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device type").AddStringEnumDescription("COMPUTE_DEVICE", "MERAKI_DASHBOARD", "NETWORK_DEVICE", "THIRD PARTY DEVICE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("COMPUTE_DEVICE", "MERAKI_DASHBOARD", "NETWORK_DEVICE", "THIRD PARTY DEVICE"),
				},
			},
			"update_mgmt_ip_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address of the device to be mapped to New IP address").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"exist_mgmt_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Existing IP address of the device").String,
							Optional:            true,
						},
						"new_mgmt_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("New IP address of the device").String,
							Optional:            true,
						},
					},
				},
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI username of the device").String,
				Optional:            true,
			},
		},
	}
}

func (r *DeviceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// ValidateConfig raises the deprecation warnings for secret attributes that have a
// write-only replacement. They are raised here, at resource level, rather than through
// the schema's DeprecationMessage: the framework raises that one against the attribute
// path, and Terraform renders an attribute-scoped diagnostic together with the offending
// configuration line, which would print the secret itself into plan output and CI logs.
// The message therefore names the attribute explicitly.
func (r *DeviceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var deprecatedEnablePassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("enable_password"), &deprecatedEnablePassword)...)
	if !deprecatedEnablePassword.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `enable_password` attribute stores the secret in Terraform state. Use `enable_password_wo` together with `enable_password_wo_version` instead, which keeps it out of state.")
	}
	var deprecatedHttpPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("http_password"), &deprecatedHttpPassword)...)
	if !deprecatedHttpPassword.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `http_password` attribute stores the secret in Terraform state. Use `http_password_wo` together with `http_password_wo_version` instead, which keeps it out of state.")
	}
	var deprecatedPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &deprecatedPassword)...)
	if !deprecatedPassword.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `password` attribute stores the secret in Terraform state. Use `password_wo` together with `password_wo_version` instead, which keeps it out of state.")
	}
	var deprecatedSnmpAuthPassphrase types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_auth_passphrase"), &deprecatedSnmpAuthPassphrase)...)
	if !deprecatedSnmpAuthPassphrase.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `snmp_auth_passphrase` attribute stores the secret in Terraform state. Use `snmp_auth_passphrase_wo` together with `snmp_auth_passphrase_wo_version` instead, which keeps it out of state.")
	}
	var deprecatedSnmpPrivPassphrase types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_priv_passphrase"), &deprecatedSnmpPrivPassphrase)...)
	if !deprecatedSnmpPrivPassphrase.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `snmp_priv_passphrase` attribute stores the secret in Terraform state. Use `snmp_priv_passphrase_wo` together with `snmp_priv_passphrase_wo_version` instead, which keeps it out of state.")
	}
	var deprecatedSnmpRoCommunity types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_ro_community"), &deprecatedSnmpRoCommunity)...)
	if !deprecatedSnmpRoCommunity.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `snmp_ro_community` attribute stores the secret in Terraform state. Use `snmp_ro_community_wo` together with `snmp_ro_community_wo_version` instead, which keeps it out of state.")
	}
	var deprecatedSnmpRwCommunity types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_rw_community"), &deprecatedSnmpRwCommunity)...)
	if !deprecatedSnmpRwCommunity.IsNull() {
		resp.Diagnostics.AddWarning("Attribute Deprecated", "The `snmp_rw_community` attribute stores the secret in Terraform state. Use `snmp_rw_community_wo` together with `snmp_rw_community_wo_version` instead, which keeps it out of state.")
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *DeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Device

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "enable_password_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("enable_password_wo"), &plan.EnablePasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "http_password_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("http_password_wo"), &plan.HttpPasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "password_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &plan.PasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_auth_passphrase_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_auth_passphrase_wo"), &plan.SnmpAuthPassphraseWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_priv_passphrase_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_priv_passphrase_wo"), &plan.SnmpPrivPassphraseWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_ro_community_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_ro_community_wo"), &plan.SnmpRoCommunityWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_rw_community_wo" is not stored in plan/state; read it from config so it can be sent to the API.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_rw_community_wo"), &plan.SnmpRwCommunityWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Device{})

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
	plan.Id = types.StringValue(res.Get("response.#(managementIpAddress==\"" + plan.IpAddress.ValueString() + "\").id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *DeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Device

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(state.Id.ValueString())
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
func (r *DeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Device

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
	// Write-only value "enable_password_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "enable_password_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("enable_password_wo"), &plan.EnablePasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "http_password_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "http_password_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("http_password_wo"), &plan.HttpPasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "password_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "password_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &plan.PasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_auth_passphrase_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "snmp_auth_passphrase_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_auth_passphrase_wo"), &plan.SnmpAuthPassphraseWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_priv_passphrase_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "snmp_priv_passphrase_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_priv_passphrase_wo"), &plan.SnmpPrivPassphraseWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_ro_community_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "snmp_ro_community_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_ro_community_wo"), &plan.SnmpRoCommunityWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Write-only value "snmp_rw_community_wo" is not stored in plan/state; read it from config so it can be sent to the API. It is read unconditionally on every Update because CatC updates are full-object replace PUTs (the whole toBody is sent), and the API requires the secret to be present on every write (omitting an unchanged secret is rejected, e.g. wireless_ssid NCND03006). The "snmp_rw_community_wo_version" companion still drives whether Terraform detects a change worth applying; it cannot make the on-wire PUT omit the field.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("snmp_rw_community_wo"), &plan.SnmpRwCommunityWo)...)
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
func (r *DeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Device

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
func (r *DeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
