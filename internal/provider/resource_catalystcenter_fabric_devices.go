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
	"reflect"
	"sort"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &FabricDevicesResource{}
var _ resource.ResourceWithImportState = &FabricDevicesResource{}

func NewFabricDevicesResource() resource.Resource {
	return &FabricDevicesResource{}
}

type FabricDevicesResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *FabricDevicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_devices"
}

func (r *FabricDevicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages Fabric Devices within a single resource, specifying a list of fabric_devices as input").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"fabric_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric this device belongs to").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"fabric_devices": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of fabric devices to be managed in bulk").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric device").String,
							Computed:            true,
						},
						"network_device_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the fabric device").String,
							Required:            true,
						},
						"fabric_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the fabric of this fabric device").String,
							Required:            true,
						},
						"device_roles": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE]").String,
							ElementType:         types.StringType,
							Required:            true,
						},
						"border_types": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3]").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"local_autonomous_system_number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("BGP Local autonomous system number of the fabric border device. Allowed range is [1 to 4294967295].").String,
							Optional:            true,
						},
						"default_exit": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes").String,
							Optional:            true,
						},
						"import_external_routes": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane.").String,
							Optional:            true,
						},
						"border_priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Border priority of the fabric border device. Allowed range is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.").AddIntegerRangeDescription(1, 9).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 9),
							},
						},
						"prepend_autonomous_system_count": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prepend autonomous system count of the fabric border device. Allowed range is [1 to 10].").AddIntegerRangeDescription(1, 10).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
					},
				},
			},
		},
	}
}

func (r *FabricDevicesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *FabricDevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricDevices

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, FabricDevices{})
	var planList []FabricDevices
	maxElementsPerShard := 10
	originalList := plan.FabricDevices
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.FabricDevices = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, FabricDevices{})
		res, err = r.client.Post(plan.getPath()+params, body, cc.UseMutex)
		if err != nil {
			errorCode := res.Get("response.errorCode").String()
			if errorCode == "NCDP10000" {
				// Log a warning and continue execution when device is unreachable
				failureReason := res.Get("response.failureReason").String()
				resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
			} else {
				resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
				break
			}
		}
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.FabricId.ValueString()))
	params = ""
	params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *FabricDevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?fabricId=" + url.QueryEscape(state.FabricId.ValueString())
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
func (r *FabricDevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricDevices

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

	// Initialize toDelete, toCreate, toReplace, and toUpdate with empty slices
	var toDelete = FabricDevices{
		FabricDevices: []FabricDevicesFabricDevices{},
	}
	var toCreate = FabricDevices{
		FabricDevices: []FabricDevicesFabricDevices{},
	}
	var toUpdate = FabricDevices{
		FabricDevices: []FabricDevicesFabricDevices{},
	}
	var toReplace = FabricDevices{
		FabricDevices: []FabricDevicesFabricDevices{},
	}

	planMap := make(map[string]FabricDevicesFabricDevices)
	stateMap := make(map[string]FabricDevicesFabricDevices)

	// Populate state map
	for _, v := range state.FabricDevices {
		stateMap[v.NetworkDeviceId.ValueString()] = v
	}

	// Populate plan map
	for _, v := range plan.FabricDevices {
		planMap[v.NetworkDeviceId.ValueString()] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.FabricDevices = append(toDelete.FabricDevices, stateItem)
		}
	}

	// Find items to create update and replace
	for planKey, planItem := range planMap {
		if stateItem, exists := stateMap[planKey]; exists {
			// Exists in both, check if different
			if !reflect.DeepEqual(planItem, stateItem) {
				// Update planItem but ensure ID comes from stateItem
				planItem.Id = stateItem.Id
				planMap[planKey] = planItem // Store back in planMap
				// Check if any field marked as requires_replace differs
				if planItem.NetworkDeviceId != stateItem.NetworkDeviceId {
					toReplace.FabricDevices = append(toReplace.FabricDevices, planItem)
					continue
				}
				if planItem.FabricId != stateItem.FabricId {
					toReplace.FabricDevices = append(toReplace.FabricDevices, planItem)
					continue
				}
				var planR []string
				var stateR []string
				stateItem.DeviceRoles.ElementsAs(ctx, &stateR, false)
				planItem.DeviceRoles.ElementsAs(ctx, &planR, false)
				sort.Strings(stateR)
				sort.Strings(planR)
				if !reflect.DeepEqual(stateR, planR) {
					toReplace.FabricDevices = append(toReplace.FabricDevices, planItem)
					continue
				}
				if planItem.LocalAutonomousSystemNumber != stateItem.LocalAutonomousSystemNumber {
					toReplace.FabricDevices = append(toReplace.FabricDevices, planItem)
					continue
				}
				if planItem.DefaultExit != stateItem.DefaultExit {
					toReplace.FabricDevices = append(toReplace.FabricDevices, planItem)
					continue
				}
				if planItem.ImportExternalRoutes != stateItem.ImportExternalRoutes {
					toReplace.FabricDevices = append(toReplace.FabricDevices, planItem)
					continue
				}
				toUpdate.FabricDevices = append(toUpdate.FabricDevices, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.FabricDevices = append(toCreate.FabricDevices, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.FabricDevices) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.FabricDevices)))
		// Clear IDs before recreating
		var toReplaceNoId = FabricDevices{
			FabricDevices: []FabricDevicesFabricDevices{},
		}
		for _, item := range toReplace.FabricDevices {
			item.Id = types.StringNull()
			toReplaceNoId.FabricDevices = append(toReplaceNoId.FabricDevices, item)
		}

		// Replace is done by delete + create
		toDelete.FabricDevices = append(toDelete.FabricDevices, toReplace.FabricDevices...)
		toCreate.FabricDevices = append(toCreate.FabricDevices, toReplaceNoId.FabricDevices...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.FabricDevices) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.FabricDevices)))
		for _, v := range toDelete.FabricDevices {
			if v.Id.IsNull() {
				continue // Skip if id is null
			}
			res, err := r.client.Delete(plan.getPath()+"/"+url.QueryEscape(v.Id.ValueString()), cc.UseMutex)
			if err != nil {
				errorCode := res.Get("response.errorCode").String()
				if errorCode == "NCDP10000" {
					// Log a warning and continue execution when device is unreachable
					failureReason := res.Get("response.failureReason").String()
					resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
				} else {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "DELETE", err, res.String()))
					return
				}
			}
		}
	}

	// CREATE
	// If there are objects marked for create
	if len(toCreate.FabricDevices) > 0 {

		maxElementsPerShard := 10
		var createList []FabricDevices
		for i := 0; i < len(toCreate.FabricDevices); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.FabricDevices))
			chunk := toCreate.FabricDevices[i:end]
			currentPlanForShard := plan
			currentPlanForShard.FabricDevices = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.FabricDevices)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, FabricDevices{}) // Convert to request body
			res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
			if err != nil {
				errorCode := res.Get("response.errorCode").String()
				if errorCode == "NCDP10000" {
					// Log a warning and continue execution when device is unreachable
					failureReason := res.Get("response.failureReason").String()
					resp.Diagnostics.AddWarning("Device Unreachability Warning", fmt.Sprintf("Device unreachability detected (error code: %s, reason %s).", errorCode, failureReason))
				} else {
					resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
					break
				}
			}
		}
		params += "?fabricId=" + url.QueryEscape(plan.FabricId.ValueString())
		res, err = r.client.Get(plan.getPath() + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		// Populate missing IDs using fromBodyUnknowns
		plan.fromBodyUnknowns(ctx, res)
	}

	// UPDATE
	// Update objects (objects that have different definition in plan and state)
	if len(toUpdate.FabricDevices) > 0 {

		maxElementsPerShard := 10
		var updateList []FabricDevices
		for i := 0; i < len(toUpdate.FabricDevices); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toUpdate.FabricDevices))
			chunk := toUpdate.FabricDevices[i:end]
			currentPlanForShard := plan
			currentPlanForShard.FabricDevices = chunk
			updateList = append(updateList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.FabricDevices)))
		planIndexMap := make(map[string]int)

		for _, pl := range updateList {
			for i, v := range plan.FabricDevices {
				planIndexMap[v.NetworkDeviceId.ValueString()] = i
			}
			for _, item := range pl.FabricDevices {
				toUpdateKey := item.NetworkDeviceId.ValueString()
				if updatedItem, exists := planMap[toUpdateKey]; exists {
					if index, found := planIndexMap[toUpdateKey]; found {
						plan.FabricDevices[index] = updatedItem
					}
				}
			}

			body := pl.toBody(ctx, FabricDevices{})
			params := ""
			res, err := r.client.Put(plan.getPath()+params, body, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				break
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *FabricDevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	params := "?fabricId=" + url.QueryEscape(state.FabricId.ValueString())
	res, err := r.client.Delete(state.getPath()+params, cc.UseMutex)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		errorCode := res.Get("response.errorCode").String()
		if errorCode == "NCDP10000" {
			// Log a warning and continue execution when device is unreachable
			failureReason := res.Get("response.failureReason").String()
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
func (r *FabricDevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <fabric_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fabric_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import
