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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &FloorsResource{}
var _ resource.ResourceWithImportState = &FloorsResource{}

func NewFloorsResource() resource.Resource {
	return &FloorsResource{}
}

type FloorsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
}

func (r *FloorsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_floors"
}

func (r *FloorsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages multiple floors within a single resource, specifying a list of floors as input").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"floors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of floors").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The id of the floor").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"parent_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The parent id of the floor").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"parent_name_hierarchy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Parent hierarchical name (e.g., Global/USA/San Jose/Building1)").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Floor name").String,
							Required:            true,
						},
						"floor_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Floor number").String,
							Required:            true,
						},
						"rf_model": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RF model").AddStringEnumDescription("Cubes And Walled Offices", "Drywall Office Only", "Indoor High Ceiling", "Outdoor Open Space", "Free Space").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Cubes And Walled Offices", "Drywall Office Only", "Indoor High Ceiling", "Outdoor Open Space", "Free Space"),
							},
						},
						"width": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Width").String,
							Required:            true,
						},
						"length": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Length").String,
							Required:            true,
						},
						"height": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Height").String,
							Required:            true,
						},
						"units_of_measure": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The unit of measurement").AddStringEnumDescription("feet", "meters").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("feet", "meters"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *FloorsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
}

// End of section. //template:end model

func (r *FloorsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Floors

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set resource-level ID for bulk resource at the start
	plan.Id = types.StringValue("floors-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Floors{})
	var planList []Floors
	maxElementsPerShard := 1000
	originalList := plan.Floors
	for i := 0; i < len(originalList); i += maxElementsPerShard {
		end := i + maxElementsPerShard
		if end > len(originalList) {
			end = len(originalList)
		}
		chunk := originalList[i:end]
		currentPlanForShard := plan
		currentPlanForShard.Floors = chunk
		planList = append(planList, currentPlanForShard)

	}

	params := ""
	var err error
	var res gjson.Result
	for _, pl := range planList {
		body = pl.toBody(ctx, Floors{})
		res, err = r.client.Post(plan.getPath()+params, body)
		if err != nil {
			resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			break
		}
	}
	// Get all floors to populate IDs
	params = "?type=floor"
	res, err = r.client.Get("/dna/intent/api/v1/sites" + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	// Set resource-level ID for bulk resource (use constant since this manages multiple items)
	plan.Id = types.StringValue("floors-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FloorsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Floors

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	// Get all floors with type=floor query parameter
	params := "?type=floor"
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

	// Filter out any floors without a name
	filteredFloors := []FloorsFloors{}
	for _, floor := range state.Floors {
		if !floor.Name.IsNull() && floor.Name.ValueString() != "" {
			filteredFloors = append(filteredFloors, floor)
		}
	}
	state.Floors = filteredFloors

	// Set resource-level ID for bulk resource
	state.Id = types.StringValue("floors-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *FloorsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Floors

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
	plan.Id = types.StringValue("floors-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Initialize toDelete, toCreate, toReplace, and toUpdate with empty slices
	var toDelete = Floors{
		Floors: []FloorsFloors{},
	}
	var toCreate = Floors{
		Floors: []FloorsFloors{},
	}
	var toUpdate = Floors{
		Floors: []FloorsFloors{},
	}
	var toReplace = Floors{
		Floors: []FloorsFloors{},
	}

	planMap := make(map[string]FloorsFloors)
	stateMap := make(map[string]FloorsFloors)

	// Populate state map - use parentNameHierarchy + name as unique key
	for _, v := range state.Floors {
		key := v.ParentNameHierarchy.ValueString() + "/" + v.Name.ValueString()
		stateMap[key] = v
	}

	// Populate plan map - use parentNameHierarchy + name as unique key
	for _, v := range plan.Floors {
		key := v.ParentNameHierarchy.ValueString() + "/" + v.Name.ValueString()
		planMap[key] = v
	}

	// Find items to delete (exist in state but not in plan)
	for stateKey, stateItem := range stateMap {
		if _, exists := planMap[stateKey]; !exists {
			// Exists only in state → Needs to be deleted
			toDelete.Floors = append(toDelete.Floors, stateItem)
		}
	}

	// Find items to create update and replace
	for planKey, planItem := range planMap {
		if stateItem, exists := stateMap[planKey]; exists {
			// Exists in both, check if different (excluding computed fields)
			// Compare only user-configurable fields
			planItemCopy := planItem
			planItemCopy.Id = stateItem.Id // Use state ID
			planItemCopy.ParentId = stateItem.ParentId // Use state ParentId

			if !reflect.DeepEqual(planItemCopy, stateItem) {
				// Update planItem but ensure ID comes from stateItem
				planItem.Id = stateItem.Id
				planItem.ParentId = stateItem.ParentId
				planMap[planKey] = planItem // Store back in planMap
				// Check if any field marked as requires_replace differs
				toUpdate.Floors = append(toUpdate.Floors, planItem)
			}
		} else {
			// Exists only in plan → New item
			toCreate.Floors = append(toCreate.Floors, planItem)
		}
	}

	// REPLACE
	// If there are objects marked to be replaced
	if len(toReplace.Floors) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to replace: %d", state.Id.ValueString(), len(toReplace.Floors)))
		// Clear IDs before recreating
		var toReplaceNoId = Floors{
			Floors: []FloorsFloors{},
		}
		for _, item := range toReplace.Floors {
			item.Id = types.StringNull()
			toReplaceNoId.Floors = append(toReplaceNoId.Floors, item)
		}

		// Replace is done by delete + create
		toDelete.Floors = append(toDelete.Floors, toReplace.Floors...)
		toCreate.Floors = append(toCreate.Floors, toReplaceNoId.Floors...)
	}

	// DELETE
	// If there are objects marked to be deleted
	if len(toDelete.Floors) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to delete: %d", state.Id.ValueString(), len(toDelete.Floors)))
		for _, v := range toDelete.Floors {
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
	if len(toCreate.Floors) > 0 {

		maxElementsPerShard := 1000
		var createList []Floors
		for i := 0; i < len(toCreate.Floors); i += maxElementsPerShard {
			end := min(i+maxElementsPerShard, len(toCreate.Floors))
			chunk := toCreate.Floors[i:end]
			currentPlanForShard := plan
			currentPlanForShard.Floors = chunk
			createList = append(createList, currentPlanForShard)

		}

		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to create: %d", state.Id.ValueString(), len(toCreate.Floors)))
		var err error
		var res gjson.Result
		params := ""
		for _, pl := range createList {
			body := pl.toBody(ctx, Floors{}) // Convert to request body
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
		// Get all floors to populate IDs
		params = "?type=floor"
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
	if len(toUpdate.Floors) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Number of items to update: %d", state.Id.ValueString(), len(toUpdate.Floors)))

		// Update each floor individually using PUT endpoint
		for _, item := range toUpdate.Floors {
			if item.Id.IsNull() {
				continue // Skip if id is null
			}
			// Create a temporary Floors structure with single item for toBody
			tempFloors := Floors{Floors: []FloorsFloors{item}}
			tempFloors.Id = types.StringValue("dummy") // Set to trigger PUT mode in toBody
			body := tempFloors.toBody(ctx, tempFloors)

			// Extract the single item from the array for individual PUT
			itemBody := gjson.Parse(body).Array()[0].Raw

			// Use individual floor PUT endpoint
			res, err := r.client.Put("/dna/intent/api/v2/floors/"+url.QueryEscape(item.Id.ValueString()), itemBody, cc.UseMutex)
			if err != nil {
				resp.Diagnostics.AddWarning("Client Error", fmt.Sprintf("Failed to update object (PUT) for floor %s, got error: %s, %s", item.Name.ValueString(), err, res.String()))
				// Continue with other updates instead of breaking
			}
		}
		// Get all floors to populate IDs after updates
		params := "?type=floor"
		res, err := r.client.Get("/dna/intent/api/v1/sites" + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects (GET), got error: %s, %s", err, res.String()))
			return
		}
		plan.fromBodyUnknowns(ctx, res)
	}
	// Set resource-level ID for bulk resource
	plan.Id = types.StringValue("floors-bulk")

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FloorsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Floors

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// Delete all floors in the list
	for _, floor := range state.Floors {
		if floor.Id.IsNull() {
			continue // Skip if id is null
		}
		res, err := r.client.Delete("/dna/intent/api/v1/site/" + url.QueryEscape(floor.Id.ValueString()))
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete floor %s (DELETE), got error: %s, %s", floor.Name.ValueString(), err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *FloorsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// For bulk resources, use a constant ID
	// Import command: terraform import catalystcenter_floors.bulk_floors floors-bulk
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "floors-bulk")...)
}

// End of section. //template:end import
