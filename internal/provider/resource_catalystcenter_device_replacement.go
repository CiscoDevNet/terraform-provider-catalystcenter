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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DeviceReplacementResource{}
var _ resource.ResourceWithImportState = &DeviceReplacementResource{}

func NewDeviceReplacementResource() resource.Resource {
	return &DeviceReplacementResource{}
}

type DeviceReplacementResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *DeviceReplacementResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_replacement"
}

func (r *DeviceReplacementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Device Replacement.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"faulty_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the faulty device to be replaced").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"replacement_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The replacement status of the device. Use MARKED-FOR-REPLACEMENT to mark the device for replacement.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"family": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The device family (computed from the faulty device)").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"faulty_device_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the faulty device (computed from the faulty device)").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"faulty_device_platform": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The platform of the faulty device (computed from the faulty device)").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"faulty_device_serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The serial number of the faulty device (computed from the faulty device)").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"replacement_device_platform": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The platform of the replacement device").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"replacement_device_serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The serial number of the replacement device").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"neighbour_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the neighbour device to create the DHCP server").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"network_readiness_task_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the network readiness task").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"workflow_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the replacement workflow").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"creation_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The creation time of the device replacement entry").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"replacement_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The replacement time of the device replacement entry").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"workflow_failed_step": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The failed step of the replacement workflow").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"readiness_check_task_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the readiness check task").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *DeviceReplacementResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below has custom code (not generated). Do not add template markers.
func (r *DeviceReplacementResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceReplacement

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, DeviceReplacement{})

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
	plan.Id = types.StringValue(res.Get("response.#(faultyDeviceId==\"" + plan.FaultyDeviceId.ValueString() + "\").id").String())

	// Read back the created object to populate computed attributes only.
	// We must NOT call fromBody here because it would overwrite plan values
	// (e.g. optional fields the user didn't set, or replacement_status which
	// the API may change from MARKED-FOR-REPLACEMENT to READY-FOR-REPLACEMENT).
	itemRes := res.Get("response.#(faultyDeviceId==\"" + plan.FaultyDeviceId.ValueString() + "\")")
	// Populate all computed-only fields from the API response
	if value := itemRes.Get("family"); value.Exists() && value.Type != gjson.Null {
		plan.Family = types.StringValue(value.String())
	} else {
		plan.Family = types.StringNull()
	}
	if value := itemRes.Get("faultyDeviceName"); value.Exists() && value.Type != gjson.Null {
		plan.FaultyDeviceName = types.StringValue(value.String())
	} else {
		plan.FaultyDeviceName = types.StringNull()
	}
	if value := itemRes.Get("faultyDevicePlatform"); value.Exists() && value.Type != gjson.Null {
		plan.FaultyDevicePlatform = types.StringValue(value.String())
	} else {
		plan.FaultyDevicePlatform = types.StringNull()
	}
	if value := itemRes.Get("faultyDeviceSerialNumber"); value.Exists() && value.Type != gjson.Null {
		plan.FaultyDeviceSerialNumber = types.StringValue(value.String())
	} else {
		plan.FaultyDeviceSerialNumber = types.StringNull()
	}
	if value := itemRes.Get("networkReadinessTaskId"); value.Exists() && value.Type != gjson.Null {
		plan.NetworkReadinessTaskId = types.StringValue(value.String())
	} else {
		plan.NetworkReadinessTaskId = types.StringNull()
	}
	if value := itemRes.Get("workflowId"); value.Exists() && value.Type != gjson.Null {
		plan.WorkflowId = types.StringValue(value.String())
	} else {
		plan.WorkflowId = types.StringNull()
	}
	if value := itemRes.Get("creationTime"); value.Exists() {
		data := value.String()
		if data != "" {
			plan.CreationTime = types.Int64Value(value.Int())
		} else {
			plan.CreationTime = types.Int64Null()
		}
	} else {
		plan.CreationTime = types.Int64Null()
	}
	if value := itemRes.Get("replacementTime"); value.Exists() {
		data := value.String()
		if data != "" {
			plan.ReplacementTime = types.Int64Value(value.Int())
		} else {
			plan.ReplacementTime = types.Int64Null()
		}
	} else {
		plan.ReplacementTime = types.Int64Null()
	}
	if value := itemRes.Get("workflowFailedStep"); value.Exists() && value.Type != gjson.Null {
		plan.WorkflowFailedStep = types.StringValue(value.String())
	} else {
		plan.WorkflowFailedStep = types.StringNull()
	}
	if value := itemRes.Get("readinesscheckTaskId"); value.Exists() && value.Type != gjson.Null {
		plan.ReadinessCheckTaskId = types.StringValue(value.String())
	} else {
		plan.ReadinessCheckTaskId = types.StringNull()
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *DeviceReplacementResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceReplacement

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
func (r *DeviceReplacementResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceReplacement

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Custom Delete: unmarks device by sending PUT with replacementStatus=NON-FAULTY (only if currently marked)
func (r *DeviceReplacementResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceReplacement

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// Check current replacement status before attempting to unmark
	getRes, getErr := r.client.Get(state.getPath())
	if getErr != nil {
		// If we can't read, try to unmark anyway
		tflog.Debug(ctx, fmt.Sprintf("%s: Failed to read current state, attempting unmark: %s", state.Id.ValueString(), getErr))
	} else {
		item := getRes.Get("response.#(id==\"" + state.Id.ValueString() + "\")")
		if !item.Exists() {
			// Entry doesn't exist in API — just remove from state
			tflog.Debug(ctx, fmt.Sprintf("%s: Device replacement entry not found, removing from state", state.Id.ValueString()))
			resp.State.RemoveResource(ctx)
			return
		}
		currentStatus := item.Get("replacementStatus").String()
		if currentStatus == "NON-FAULTY" || currentStatus == "" {
			// Already unmarked or empty — nothing to do, just remove from state
			tflog.Debug(ctx, fmt.Sprintf("%s: Device not in replacement workflow (status=%s), removing from state only", state.Id.ValueString(), currentStatus))
			resp.State.RemoveResource(ctx)
			return
		}
	}

	// Device is marked for replacement — unmark it
	body := `[{"id":"` + state.Id.ValueString() + `","replacementStatus":"NON-FAULTY"}]`
	res, err := r.client.Put(state.getPath(), body)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		errorCode := res.Get("response.errorCode").String()
		if strings.HasPrefix(errorCode, "NCND") {
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
func (r *DeviceReplacementResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *DeviceReplacementResource) ReadCache(ctx context.Context, req resource.ReadRequest, state DeviceReplacement, params string) (cc.Res, error) {
	var err error
	cacheKey := "DeviceReplacement::"

	_, cacheSuffix, found := strings.Cut(params, "?")
	queryPart, err := url.ParseQuery(cacheSuffix)
	if err == nil {
		delete(queryPart, "id")
		newQuery := queryPart.Encode()
		cacheSuffix = "?" + newQuery
		cacheKey += cacheSuffix
	}

	cachedValue, found := r.cache.Get(cacheKey)
	if found {
		tflog.Debug(ctx, fmt.Sprintf("hit cache for %s", cacheKey))
		ccRes, ok := cachedValue.(cc.Res)
		if ok {
			return ccRes, nil
		}
		tflog.Info(ctx, fmt.Sprintf("Invalid cache entry type for %s", cacheKey))
		r.cache.Delete(cacheKey)
	}
	res, err := r.client.Get(state.getPath() + params)
	singleRes := res
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache
