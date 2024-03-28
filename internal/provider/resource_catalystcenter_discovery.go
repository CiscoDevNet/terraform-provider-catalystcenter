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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DiscoveryResource{}

func NewDiscoveryResource() resource.Resource {
	return &DiscoveryResource{}
}

type DiscoveryResource struct {
	client *cc.Client
}

func (r *DiscoveryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_discovery"
}

func (r *DiscoveryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("After discovery resource has been created, the Catalyst Center would contain device entries (if discovered). All the device entries can be subsequently obtained from data source `catalystcenter_network_devices`. Terraform currently is not able to handle `for_each` from a data source that depends on any managed resource, therefore to work around that limitation the `catalystcenter_discovery` can be placed in a different tfstate (root module) than `catalystcenter_network_devices` when the latter is used as a source of `for_each`. <p/> The discovery resource does not support updates, it needs to be destroyed and re-created instead.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cdp_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("CDP level is the number of hops across neighbor devices.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"discovery_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of Discovery.").AddStringEnumDescription("Single", "Range", "Multi Range", "CDP", "LLDP", "CIDR").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Single", "Range", "Multi Range", "CDP", "LLDP", "CIDR"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"enable_password_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable passwords of the devices to be discovered.").String,
				ElementType:         types.StringType,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
			},
			"global_credential_id_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of IDs, which must include SNMP credential and CLI credential.").String,
				ElementType:         types.StringType,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
			},
			"http_read_credential": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"http_write_credential": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ip_address_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A string of IP address ranges to discover.  E.g.: '172.30.0.1' for discovery_type Single, CDP and LLDP; '172.30.0.1-172.30.0.4' for Range; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for Multi Range; '172.30.0.1/20' for CIDR.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ip_filter_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of IP address ranges to exclude from the discovery.").String,
				ElementType:         types.StringType,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
			},
			"lldp_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("LLDP level to which neighbor devices to be discovered.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A name of the discovery.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"netconf_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port number for netconf as a string. It requires SSH protocol to work.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"password_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Passwords of the devices to be discovered.").String,
				ElementType:         types.StringType,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
			},
			"preferred_ip_method": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Preferred method for selecting management IP address.").AddStringEnumDescription("None", "UseLoopBack").AddDefaultValueDescription("None").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("None", "UseLoopBack"),
				},
				Default: stringdefault.StaticString("None"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"protocol_order": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A string of comma-separated protocols (SSH/Telnet), in the same order in which the connections to each device are attempted. E.g.: 'Telnet': only telnet; 'SSH,Telnet': ssh first, with telnet fallback.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"retry": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of times to try establishing SSH/Telnet connection to a device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"snmp_auth_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Auth passphrase for SNMP.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_auth_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP auth protocol.").AddStringEnumDescription("SHA", "MD5").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SHA", "MD5"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mode of SNMP. The `snmp_auth_protocol` and `snmp_auth_passphrase` are required for \"AuthNoPriv\" mode. Additionally, `snmp_priv_protocol` and `snmp_priv_passphrase` are required for \"AuthPriv\" mode.").AddStringEnumDescription("AuthPriv", "AuthNoPriv", "NoAuthNoPriv").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AuthPriv", "AuthNoPriv", "NoAuthNoPriv"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_priv_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Passphrase for SNMP privacy.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_priv_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP privacy protocol.").AddStringEnumDescription("DES", "AES128").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DES", "AES128"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_ro_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP RO community of the devices to be discovered.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_ro_community_desc": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description for snmp_ro_community.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_rw_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP RW community of the devices to be discovered.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_rw_community_desc": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description for snmp_rw_community").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_user_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP username of the devices to be discovered.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snmp_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP version").AddStringEnumDescription("v2", "v3").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of seconds to wait for each SSH/Telnet connection to a device.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"user_name_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Usernames for the devices to be discovered.").String,
				ElementType:         types.StringType,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *DiscoveryResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *DiscoveryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Discovery

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Discovery{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, func(r *cc.Req) { r.MaxAsyncWaitTime = 600 })
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	params = ""
	res, err = r.client.Get("/dna/intent/api/v1/discovery/1/500" + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.#(name==\"" + plan.Name.ValueString() + "\").id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *DiscoveryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Discovery

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	res, err := r.client.Get("/dna/intent/api/v1/discovery/1/500" + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	res = res.Get("response.#(id==\"" + state.Id.ValueString() + "\")")

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
func (r *DiscoveryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Discovery

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
func (r *DiscoveryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Discovery

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
