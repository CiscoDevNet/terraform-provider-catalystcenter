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
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
var _ resource.Resource = &FabricPortChannelResource{}
var _ resource.ResourceWithImportState = &FabricPortChannelResource{}

func NewFabricPortChannelResource() resource.Resource {
	return &FabricPortChannelResource{}
}

type FabricPortChannelResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *FabricPortChannelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_port_channel"
}

func (r *FabricPortChannelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages a single port channel in SD-Access fabric.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the fabric device").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric the device is assigned to").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"port_channel_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port channel name auto-assigned by Catalyst Center on creation; read-only (any value set in config is ignored).").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_names": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set of interface names for the port channel").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"connected_device_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Connected device type of the port channel").AddStringEnumDescription("TRUNK", "EXTENDED_NODE").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TRUNK", "EXTENDED_NODE"),
				},
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protocol of the port channel").AddStringEnumDescription("ON", "LACP", "PAGP").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ON", "LACP", "PAGP"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the port channel").String,
				Optional:            true,
			},
			"native_vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Native VLAN ID of the port channel").String,
				Optional:            true,
			},
			"allowed_vlan_ranges": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allowed VLAN ranges of the port channel").String,
				Optional:            true,
			},
		},
	}
}

func (r *FabricPortChannelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

func (r *FabricPortChannelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricPortChannel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	listParams := "?networkDeviceId=" + url.QueryEscape(plan.NetworkDeviceId.ValueString()) +
		"&fabricId=" + url.QueryEscape(plan.FabricId.ValueString())

	var planIfaces []string
	plan.InterfaceNames.ElementsAs(ctx, &planIfaces, false)

	if r.AllowExistingOnCreate {
		preRes, err := r.client.Get(plan.getPath() + listParams)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed pre-create GET, got error: %s, %s", err, preRes.String()))
			return
		}
		if id := findPortChannelByInterfaceSet(preRes, planIfaces); id != "" {
			tflog.Info(ctx, fmt.Sprintf("Found existing port channel id=%s matching interface set; adopting and reconciling via PUT", id))
			if name := preRes.Get(`response.#(id=="` + id + `").portChannelName`); name.Exists() {
				plan.PortChannelName = types.StringValue(name.String())
			}
			plan.Id = types.StringValue(id)
			body := plan.toBody(ctx, plan)
			res, err := r.client.Put(plan.getPath(), body, cc.UseMutex)
			if err != nil {
				if !handleDeviceUnreachable(ctx, res, err, "PUT", &resp.Diagnostics) {
					return
				}
			}
			diags = resp.State.Set(ctx, &plan)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	postPlan := plan
	postPlan.PortChannelName = types.StringNull()
	body := postPlan.toBody(ctx, FabricPortChannel{})
	res, err := r.client.Post(plan.getPath(), body, cc.UseMutex)
	if err != nil {
		if !handleDeviceUnreachable(ctx, res, err, "POST", &resp.Diagnostics) {
			return
		}
	}

	postRes, err := r.client.Get(plan.getPath() + listParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, postRes.String()))
		return
	}

	chosenId := findPortChannelByInterfaceSet(postRes, planIfaces)
	if chosenId == "" {
		resp.Diagnostics.AddError(
			"Create Verification Failed",
			"POST succeeded but no port channel matching the plan's interface set was found after creation. This may indicate silent rejection by Catalyst Center.",
		)
		return
	}

	plan.Id = types.StringValue(chosenId)
	if name := postRes.Get(`response.#(id=="` + chosenId + `").portChannelName`); name.Exists() {
		plan.PortChannelName = types.StringValue(name.String())
	} else {
		plan.PortChannelName = types.StringNull()
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *FabricPortChannelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricPortChannel

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?networkDeviceId=" + url.QueryEscape(state.NetworkDeviceId.ValueString()) + "&fabricId=" + url.QueryEscape(state.FabricId.ValueString())
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	res = res.Get("response.#(id==\"" + state.Id.ValueString() + "\")")
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
func (r *FabricPortChannelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricPortChannel

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
	res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		failureReason := res.Get("response.failureReason").String()
		deviceFailureMatch, _ := regexp.MatchString(`(?i)Operation failed on '\d+' devices`, failureReason)
		if errorCode == "NCDP10000" || deviceFailureMatch {
			// Log a warning and continue execution when device is unreachable
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (%s), got error: %s, %s", "PUT", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *FabricPortChannelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricPortChannel

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	// Skip delete if ID is empty or null to prevent sending DELETE to base endpoint
	if state.Id.IsNull() || state.Id.IsUnknown() || state.Id.ValueString() == "" {
		tflog.Debug(ctx, fmt.Sprintf("%s: Skipping delete - ID is empty or null", state.Id.ValueString()))
		resp.State.RemoveResource(ctx)
		return
	}
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), cc.UseMutex)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		errorCode := res.Get("response.errorCode").String()
		failureReason := res.Get("response.failureReason").String()
		deviceFailureMatch, _ := regexp.MatchString(`(?i)Operation failed on '\d+' devices`, failureReason)
		if errorCode == "NCDP10000" || deviceFailureMatch {
			// Log a warning and continue execution when device is unreachable
			resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "DELETE", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *FabricPortChannelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_device_id>,<fabric_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_device_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[2])...)
}

// End of section. //template:end import

// findPortChannelByInterfaceSet scans a /sda/portChannels list response and
// returns the id of the entry whose interfaceNames field — taken as a set —
// equals the provided plan set.
func findPortChannelByInterfaceSet(res cc.Res, planIfaces []string) string {
	for _, item := range res.Get("response").Array() {
		var itemIfaces []string
		for _, v := range item.Get("interfaceNames").Array() {
			itemIfaces = append(itemIfaces, v.String())
		}
		if helpers.StringSetEqual(planIfaces, itemIfaces) {
			return item.Get("id").String()
		}
	}
	return ""
}

func handleDeviceUnreachable(_ context.Context, res cc.Res, err error, verb string, diags *diag.Diagnostics) bool {
	errorCode := res.Get("response.errorCode").String()
	failureReason := res.Get("response.failureReason").String()
	deviceFailureMatch, _ := regexp.MatchString(`(?i)Operation failed on '\d+' devices`, failureReason)
	if errorCode == "NCDP10000" || deviceFailureMatch {
		diags.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
		return true
	}
	diags.AddError("Client Error", fmt.Sprintf("Failed to %s object, got error: %s, %s", verb, err, res.String()))
	return false
}
