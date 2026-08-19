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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DeviceReplacementWorkflowResource{}

func NewDeviceReplacementWorkflowResource() resource.Resource {
	return &DeviceReplacementWorkflowResource{}
}

type DeviceReplacementWorkflowResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *DeviceReplacementWorkflowResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_replacement_workflow"
}

func (r *DeviceReplacementWorkflowResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource triggers the device replacement workflow on Catalyst Center. The faulty device must already be marked for replacement (see `catalystcenter_device_replacement`). The workflow is long-running; this resource only triggers it. <p/> When this resource is destroyed or refreshed, no actions are performed on Catalyst Center.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"faulty_device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the faulty device. Used at apply time to read the live replacement record; the physical workflow fires only when a record exists whose faulty serial differs from replacement_device_serial_number.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"replacement_device_serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Serial number of the replacement device").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"faulty_device_serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Serial number of the faulty device to be replaced, read from the live Catalyst Center replacement record.").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"executed": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the physical replacement workflow has been triggered on Catalyst Center for this instance.").String,
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *DeviceReplacementWorkflowResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below has custom code (not generated). Do not add template markers.
func (r *DeviceReplacementWorkflowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceReplacementWorkflow

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	r.reconcile(&plan, resp.Diagnostics)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// reconcile reads the live Catalyst Center replacement record for the plan's
// faulty_device_id and triggers the physical replacement workflow ONLY when a
// record exists whose faulty serial differs from the replacement serial. When
// no such record exists it records a pending (executed=false) state and does
// nothing physical, so ordinary provisioned devices are never replaced.
func (r *DeviceReplacementWorkflowResource) reconcile(plan *DeviceReplacementWorkflow, diags diag.Diagnostics) {
	plan.Id = types.StringValue(plan.FaultyDeviceId.ValueString() + ":" + plan.ReplacementDeviceSerialNumber.ValueString())

	record, err := r.client.Get("/dna/intent/api/v1/device-replacement")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to retrieve replacement record (GET), got error: %s, %s", err, record.String()))
		return
	}
	item := record.Get("response.#(faultyDeviceId==\"" + plan.FaultyDeviceId.ValueString() + "\")")

	if !item.Exists() {
		plan.Executed = types.BoolValue(false)
		plan.FaultyDeviceSerialNumber = types.StringNull()
		return
	}

	faultySerial := item.Get("faultyDeviceSerialNumber").String()
	plan.FaultyDeviceSerialNumber = types.StringValue(faultySerial)

	if faultySerial == "" || faultySerial == plan.ReplacementDeviceSerialNumber.ValueString() {
		plan.Executed = types.BoolValue(false)
		return
	}

	body := ""
	body, _ = sjson.Set(body, "faultyDeviceSerialNumber", faultySerial)
	body, _ = sjson.Set(body, "replacementDeviceSerialNumber", plan.ReplacementDeviceSerialNumber.ValueString())
	res, err := r.client.Post(plan.getPath(), body, cc.NoWait)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Executed = types.BoolValue(true)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *DeviceReplacementWorkflowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceReplacementWorkflow

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below has custom code (not generated). Do not add template markers.
func (r *DeviceReplacementWorkflowResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceReplacementWorkflow

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

	if state.Executed.ValueBool() {
		plan.Executed = types.BoolValue(true)
		plan.FaultyDeviceSerialNumber = state.FaultyDeviceSerialNumber
		plan.Id = state.Id
	} else {
		r.reconcile(&plan, resp.Diagnostics)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *DeviceReplacementWorkflowResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceReplacementWorkflow

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// ModifyPlan forces a pending (never-executed) instance to be re-evaluated on
// every apply so that a replacement record appearing later in Catalyst Center
// triggers the physical workflow. Once executed, the instance is left stable so
// the destructive workflow is never re-fired.
func (r *DeviceReplacementWorkflowResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var state DeviceReplacementWorkflow
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !state.Executed.ValueBool() {
		resp.Plan.SetAttribute(ctx, path.Root("executed"), types.BoolUnknown())
		resp.Plan.SetAttribute(ctx, path.Root("faulty_device_serial_number"), types.StringUnknown())
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *DeviceReplacementWorkflowResource) ReadCache(ctx context.Context, req resource.ReadRequest, state DeviceReplacementWorkflow, params string) (cc.Res, error) {
	var err error
	cacheKey := "DeviceReplacementWorkflow::"

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
