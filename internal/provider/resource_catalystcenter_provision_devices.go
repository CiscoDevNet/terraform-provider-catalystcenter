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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
var _ resource.Resource = &ProvisionDevicesResource{}
var _ resource.ResourceWithImportState = &ProvisionDevicesResource{}

func NewProvisionDevicesResource() resource.Resource {
	return &ProvisionDevicesResource{}
}

type ProvisionDevicesResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *ProvisionDevicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_provision_devices"
}

func (r *ProvisionDevicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages the provisioning of devices within a single resource, specifying a list of devices to be provisioned in selected site.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Site Id").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"provision_devices": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of devices to be provisioned").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the provisioned device").String,
							Computed:            true,
						},
						"site_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the site this network device needs to be provisioned").String,
							Required:            true,
						},
						"network_device_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of network device to be provisioned").String,
							Required:            true,
						},
						"reprovision": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Flag to indicate whether the device should be reprovisioned. If set to `true`, reprovisioning will be triggered on every Terraform apply").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ProvisionDevicesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ProvisionDevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProvisionDevices

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ProvisionDevices{})
	var planList []ProvisionDevices
	maxElementsPerShard := 10
	originalList := plan.ProvisionDevices
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.ProvisionDevices = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, ProvisionDevices{})
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
	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))
	params = ""
	params += "?siteId=" + url.QueryEscape(plan.SiteId.ValueString())
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
func (r *ProvisionDevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProvisionDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?siteId=" + url.QueryEscape(state.SiteId.ValueString())
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
func (r *ProvisionDevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ProvisionDevices

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
	var toDelete = ProvisionDevices{
		ProvisionDevices: []ProvisionDevicesProvisionDevices{},
	}
	var toCreate = ProvisionDevices{
		ProvisionDevices: []ProvisionDevicesProvisionDevices{},
	}
	var toUpdate = ProvisionDevices{
		ProvisionDevices: []ProvisionDevicesProvisionDevices{},
	}
	var toReplace = ProvisionDevices{
		ProvisionDevices: []ProvisionDevicesProvisionDevices{},
	}

	planMap := make(map[string]ProvisionDevicesProvisionDevices)
	stateMap := make(map[string]ProvisionDevicesProvisionDevices)

	// Populate state map
	for _, v := range state.ProvisionDevices {
		stateMap[v.NetworkDeviceId.ValueString()] = v
	}

	// Populate plan map
	for _, v := range plan.ProvisionDevices {
		planMap[v.NetworkDeviceId.ValueString()] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.ProvisionDevices = append(toDelete.ProvisionDevices, stateItem)
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
				if planItem.SiteId != stateItem.SiteId {
					toReplace.ProvisionDevices = append(toReplace.ProvisionDevices, planItem)
					continue
				}
				if planItem.NetworkDeviceId != stateItem.NetworkDeviceId {
					toReplace.ProvisionDevices = append(toReplace.ProvisionDevices, planItem)
					continue
				}
				toUpdate.ProvisionDevices = append(toUpdate.ProvisionDevices, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.ProvisionDevices = append(toCreate.ProvisionDevices, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.ProvisionDevices) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.ProvisionDevices)))
		// Clear IDs before recreating
		var toReplaceNoId = ProvisionDevices{
			ProvisionDevices: []ProvisionDevicesProvisionDevices{},
		}
		for _, item := range toReplace.ProvisionDevices {
			item.Id = types.StringNull()
			toReplaceNoId.ProvisionDevices = append(toReplaceNoId.ProvisionDevices, item)
		}

		// Replace is done by delete + create
		toDelete.ProvisionDevices = append(toDelete.ProvisionDevices, toReplace.ProvisionDevices...)
		toCreate.ProvisionDevices = append(toCreate.ProvisionDevices, toReplaceNoId.ProvisionDevices...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.ProvisionDevices) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.ProvisionDevices)))
		for _, v := range toDelete.ProvisionDevices {
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
	if len(toCreate.ProvisionDevices) > 0 {

		maxElementsPerShard := 10
		var createList []ProvisionDevices
		for i := 0; i < len(toCreate.ProvisionDevices); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.ProvisionDevices))
			chunk := toCreate.ProvisionDevices[i:end]
			currentPlanForShard := plan
			currentPlanForShard.ProvisionDevices = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.ProvisionDevices)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, ProvisionDevices{}) // Convert to request body
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
		params += "?siteId=" + url.QueryEscape(plan.SiteId.ValueString())
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
	if len(toUpdate.ProvisionDevices) > 0 {

		maxElementsPerShard := 10
		var updateList []ProvisionDevices
		for i := 0; i < len(toUpdate.ProvisionDevices); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toUpdate.ProvisionDevices))
			chunk := toUpdate.ProvisionDevices[i:end]
			currentPlanForShard := plan
			currentPlanForShard.ProvisionDevices = chunk
			updateList = append(updateList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.ProvisionDevices)))
		planIndexMap := make(map[string]int)

		for _, pl := range updateList {
			for i, v := range plan.ProvisionDevices {
				planIndexMap[v.NetworkDeviceId.ValueString()] = i
			}
			for _, item := range pl.ProvisionDevices {
				toUpdateKey := item.NetworkDeviceId.ValueString()
				if updatedItem, exists := planMap[toUpdateKey]; exists {
					if index, found := planIndexMap[toUpdateKey]; found {
						plan.ProvisionDevices[index] = updatedItem
					}
				}
			}

			body := pl.toBody(ctx, ProvisionDevices{})
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
func (r *ProvisionDevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProvisionDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), cc.UseMutex)
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
func (r *ProvisionDevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <site_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import
