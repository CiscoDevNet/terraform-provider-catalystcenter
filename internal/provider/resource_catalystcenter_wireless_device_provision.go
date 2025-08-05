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
var _ resource.Resource = &WirelessDeviceProvisionResource{}

func NewWirelessDeviceProvisionResource() resource.Resource {
	return &WirelessDeviceProvisionResource{}
}

type WirelessDeviceProvisionResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *WirelessDeviceProvisionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_device_provision"
}

func (r *WirelessDeviceProvisionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource is used to provision a wireless device. To reprovision a device, set the `reprovision` attribute to `true`. The resource will then trigger reprovisioning on every Terraform apply.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network Device ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ap_reboot_percentage": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("AP Reboot Percentage").String,
				Optional:            true,
			},
			"enable_rolling_ap_upgrade": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if Rolling AP Upgrade is enabled, else False").String,
				Optional:            true,
			},
			"skip_ap_provision": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if Skip AP Provision is enabled, else False").String,
				Optional:            true,
			},
			"ap_authorization_list_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AP Authorization List Name").String,
				Optional:            true,
			},
			"authorize_mesh_and_non_mesh_access_points": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("True if Mesh and Non-Mesh Access Points are authorized, else False").String,
				Optional:            true,
			},
			"reprovision": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the device should be reprovisioned. If set to `true`, reprovisioning will be triggered on every Terraform apply").String,
				Optional:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dynamic Interface Details. The required attributes depend on the device type").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface Name").String,
							Required:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN ID").AddIntegerRangeDescription(1, 4094).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"interface_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface IP Address").String,
							Optional:            true,
						},
						"interface_netmask": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface Netmask In CIDR").AddIntegerRangeDescription(1, 30).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 30),
							},
						},
						"interface_gateway": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface Gateway").String,
							Optional:            true,
						},
						"lag_or_port_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("LAG or Port Number").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *WirelessDeviceProvisionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

func (r *WirelessDeviceProvisionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessDeviceProvision

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessDeviceProvision{})

	if body == "" {
		body = "{}"
	}

	tflog.Debug(ctx, fmt.Sprintf("%v: BODYBODY", body))

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.NetworkDeviceId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WirelessDeviceProvisionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessDeviceProvision

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	state.Reprovision = types.BoolValue(false) // Reset reprovision flag to false on read

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *WirelessDeviceProvisionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessDeviceProvision

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
	if plan.Reprovision.ValueBool() {
		body := "{}" // Empty body for reprovisioning
		tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Re-Provisioning", plan.Id.ValueString()))
		params := ""
		res, err := r.client.Post(plan.getPathUpdate()+params, body)
		if err != nil {
			errorCode := res.Get("response.errorCode").String()
			if errorCode == "NCDP10000" {
				// Log a warning and continue execution when device is unreachable
				failureReason := res.Get("response.failureReason").String()
				resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (%s), got error: %s, %s", "PUT", err, res.String()))
				return
			}
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Device Re-Provisioning finished successfully", plan.Id.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *WirelessDeviceProvisionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessDeviceProvision

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	params := "?networkDeviceId=" + url.QueryEscape(state.NetworkDeviceId.ValueString())
	res, err := r.client.Delete(state.getPathDelete() + params)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
