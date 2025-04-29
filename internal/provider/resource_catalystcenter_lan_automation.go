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
var _ resource.Resource = &LANAutomationResource{}
var _ resource.ResourceWithImportState = &LANAutomationResource{}

func NewLANAutomationResource() resource.Resource {
	return &LANAutomationResource{}
}

type LANAutomationResource struct {
	client *cc.Client
}

func (r *LANAutomationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lan_automation"
}

func (r *LANAutomationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can Start LAN Automation on resource creation and Stop LAN Automation on resource deletion").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"discovered_device_site_name_hierarchy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Discovered device site name.").String,
				Required:            true,
			},
			"primary_device_management_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary seed management IP address.").String,
				Required:            true,
			},
			"peer_device_management_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary seed management IP address.").String,
				Optional:            true,
			},
			"primary_device_interface_names": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of interfaces on primary seed via which the discovered devices are connected.").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"ip_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of IP pools with its name and role.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_pool_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the IP pool.").String,
							Required:            true,
						},
						"ip_pool_role": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Role of the IP pool.").AddStringEnumDescription("MAIN_POOL", "PHYSICAL_LINK_POOL").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MAIN_POOL", "PHYSICAL_LINK_POOL"),
							},
						},
					},
				},
			},
			"multicast_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To enable underlay native multicast.").String,
				Optional:            true,
			},
			"host_name_prefix": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Host name prefix which shall be assigned to the discovered device.").String,
				Optional:            true,
			},
			"host_name_file_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("File ID of the CSV file containing the host name list.").String,
				Optional:            true,
			},
			"isis_domain_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ISIS domain password.").String,
				Optional:            true,
			},
			"redistribute_isis_to_bgp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise LAN Automation summary route into BGP.").String,
				Optional:            true,
			},
			"discovery_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier").AddIntegerRangeDescription(1, 5).AddDefaultValueDescription("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 5),
				},
				Default: int64default.StaticInt64(2),
			},
			"discovery_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Discovery timeout in minutes. Until this time, the stop processing will not be triggered.").AddIntegerRangeDescription(20, 10080).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(20, 10080),
				},
			},
			"discovery_devices": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of specific devices that will be LAN Automated in this session").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"device_serial_number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of the device").String,
							Required:            true,
						},
						"device_host_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname of the device").String,
							Optional:            true,
						},
						"device_site_name_hierarchy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type").String,
							Optional:            true,
						},
						"device_management_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Management IP Address of the device").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *LANAutomationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *LANAutomationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LANAutomation

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, LANAutomation{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, cc.NoWait)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.taskId").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *LANAutomationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LANAutomation

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get("/dna/intent/api/v1/lan-automation/status" + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
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
func (r *LANAutomationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state LANAutomation

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

func (r *LANAutomationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LANAutomation

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPathDelete() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 406") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *LANAutomationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
