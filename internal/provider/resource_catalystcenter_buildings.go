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

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &BuildingsResource{}
var _ resource.ResourceWithImportState = &BuildingsResource{}

func NewBuildingsResource() resource.Resource {
	return &BuildingsResource{}
}

type BuildingsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *BuildingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_buildings"
}

func (r *BuildingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages multiple buildings within a single resource, specifying a list of buildings as input").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"buildings": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of buildings").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The id of the building").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"parent_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The parent ID of the building").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"parent_name_hierarchy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Parent hierarchical name (e.g., Global/USA/San Jose)").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the building").String,
							Required:            true,
						},
						"country": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The country of the building").String,
							Required:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The address of the building").String,
							Optional:            true,
						},
						"latitude": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Latitude").String,
							Optional:            true,
						},
						"longitude": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Longitude").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *BuildingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

func (r *BuildingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Buildings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set resource-level ID for bulk resource at the start
	plan.Id = types.StringValue("buildings-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Buildings{})
	var planList []Buildings
	maxElementsPerShard := 1000
	originalList := plan.Buildings
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.Buildings = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, Buildings{})
		res, err = r.client.Post(plan.getPath()+params, body)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			break
		}
	}
	// Get all buildings to populate IDs
	params = "?type=building"
	res, err = r.client.Get("/dna/intent/api/v1/sites" + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	// Set resource-level ID for bulk resource (use constant since this manages multiple items)
	plan.Id = types.StringValue("buildings-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BuildingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Buildings

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	// Get all buildings with type=building query parameter
	params := "?type=building"
	res, err := r.client.Get("/dna/intent/api/v1/sites" + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	// Filter out any buildings without a name
	filteredBuildings := []BuildingsBuildings{}
	for _, building := range state.Buildings {
		if !building.Name.IsNull() && building.Name.ValueString() != "" {
			filteredBuildings = append(filteredBuildings, building)
		}
	}
	state.Buildings = filteredBuildings

	// Set resource-level ID for bulk resource
	state.Id = types.StringValue("buildings-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *BuildingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Buildings

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

	// Set resource-level ID for bulk resource at the start
	plan.Id = types.StringValue("buildings-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Initialize toDelete, toCreate, toReplace, and toUpdate with empty slices
	var toDelete = Buildings{
		Buildings: []BuildingsBuildings{},
	}
	var toCreate = Buildings{
		Buildings: []BuildingsBuildings{},
	}
	var toUpdate = Buildings{
		Buildings: []BuildingsBuildings{},
	}
	var toReplace = Buildings{
		Buildings: []BuildingsBuildings{},
	}

	planMap := make(map[string]BuildingsBuildings)
	stateMap := make(map[string]BuildingsBuildings)

	// Populate state map - use parentNameHierarchy + name as unique key
	for _, v := range state.Buildings {
		key := v.ParentNameHierarchy.ValueString() + "/" + v.Name.ValueString()
		stateMap[key] = v
	}

	// Populate plan map - use parentNameHierarchy + name as unique key
	for _, v := range plan.Buildings {
		key := v.ParentNameHierarchy.ValueString() + "/" + v.Name.ValueString()
		planMap[key] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.Buildings = append(toDelete.Buildings, stateItem)
		}
	}

	// Find items to create update and replace
	for planKey, planItem := range planMap {
		if stateItem, exists := stateMap[planKey]; exists {
			// Exists in both, check if different (excluding computed fields)
			// Compare only user-configurable fields
			planItemCopy := planItem
			planItemCopy.Id = stateItem.Id             // Use state ID
			planItemCopy.ParentId = stateItem.ParentId // Use state ParentId

			if !reflect.DeepEqual(planItemCopy, stateItem) {
				// Update planItem but ensure ID comes from stateItem
				planItem.Id = stateItem.Id
				planItem.ParentId = stateItem.ParentId
				planMap[planKey] = planItem // Store back in planMap
				// Check if any field marked as requires_replace differs
				toUpdate.Buildings = append(toUpdate.Buildings, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.Buildings = append(toCreate.Buildings, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.Buildings) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.Buildings)))
		// Clear IDs before recreating
		var toReplaceNoId = Buildings{
			Buildings: []BuildingsBuildings{},
		}
		for _, item := range toReplace.Buildings {
			item.Id = types.StringNull()
			toReplaceNoId.Buildings = append(toReplaceNoId.Buildings, item)
		}

		// Replace is done by delete + create
		toDelete.Buildings = append(toDelete.Buildings, toReplace.Buildings...)
		toCreate.Buildings = append(toCreate.Buildings, toReplaceNoId.Buildings...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.Buildings) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.Buildings)))
		for _, v := range toDelete.Buildings {
			if v.Id.IsNull() {
				continue // Skip if id is null
			}
			// Use individual site delete endpoint
			res, err := r.client.Delete("/dna/intent/api/v1/site/" + url.QueryEscape(v.Id.ValueString()))
			if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (%s), got error: %s, %s", "DELETE", err, res.String()))
				return
			}
		}
	}

	// CREATE
	// If there are objects marked for create
	if len(toCreate.Buildings) > 0 {

		maxElementsPerShard := 1000
		var createList []Buildings
		for i := 0; i < len(toCreate.Buildings); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.Buildings))
			chunk := toCreate.Buildings[i:end]
			currentPlanForShard := plan
			currentPlanForShard.Buildings = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.Buildings)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, Buildings{}) // Convert to request body
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
		// Get all buildings to populate IDs
		params = "?type=building"
		res, err = r.client.Get("/dna/intent/api/v1/sites" + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s, %s", err, res.String()))
			return
		}

		// Populate missing IDs using fromBodyUnknowns
		plan.fromBodyUnknowns(ctx, res)
	}

	// UPDATE
	// Update objects (objects that have different definition in plan and state)
	if len(toUpdate.Buildings) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.Buildings)))

		// Update each building individually using PUT endpoint
		for _, item := range toUpdate.Buildings {
			if item.Id.IsNull() {
				continue // Skip if id is null
			}
			// Create a temporary Buildings structure with single item for toBody
			tempBuildings := Buildings{Buildings: []BuildingsBuildings{item}}
			tempBuildings.Id = types.StringValue("dummy") // Set to trigger PUT mode in toBody
			body := tempBuildings.toBody(ctx, tempBuildings)

			// Extract the single item from the array for individual PUT
			itemBody := gjson.Parse(body).Array()[0].Raw

			// Use individual building PUT endpoint
			res, err := r.client.Put("/dna/intent/api/v2/buildings/"+url.QueryEscape(item.Id.ValueString()), itemBody, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to update object (PUT) for building %s, got error: %s, %s", item.Name.ValueString(), err, res.String()))
				// Continue with other updates instead of breaking
			}
		}
		// Get all buildings to populate IDs after updates
		params := "?type=building"
		res, err := r.client.Get("/dna/intent/api/v1/sites" + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s, %s", err, res.String()))
			return
		}
		plan.fromBodyUnknowns(ctx, res)
	}

	// Set resource-level ID for bulk resource
	plan.Id = types.StringValue("buildings-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BuildingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Buildings

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// Delete all buildings in the list
	for _, building := range state.Buildings {
		if building.Id.IsNull() {
			continue // Skip if id is null
		}
		res, err := r.client.Delete("/dna/intent/api/v1/site/" + url.QueryEscape(building.Id.ValueString()))
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete building %s (DELETE), got error: %s, %s", building.Name.ValueString(), err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *BuildingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// For bulk resources, use a constant ID
	// Import command: terraform import catalystcenter_buildings.bulk_buildings buildings-bulk
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "buildings-bulk")...)
}
