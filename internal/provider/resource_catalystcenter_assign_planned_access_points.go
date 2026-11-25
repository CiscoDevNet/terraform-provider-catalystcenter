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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AssignPlannedAccessPointsResource{}

func NewAssignPlannedAccessPointsResource() resource.Resource {
	return &AssignPlannedAccessPointsResource{}
}

type AssignPlannedAccessPointsResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AssignPlannedAccessPointsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assign_planned_access_points"
}

func (r *AssignPlannedAccessPointsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource assigns planned Access Points on a floor map to operational Access Points from the inventory. Planned APs are referenced by name, for example planned APs from a map archive import. For planned APs managed with `catalystcenter_planned_access_point_position`, use its `access_point_id` attribute instead. Each operational AP takes the position, height and antenna orientation of its planned AP and is assigned to the floor; Catalyst Center then removes the planned AP. <p/> Because the planned AP no longer exists after assignment, refresh only checks that each operational AP is still positioned on the floor. Destroy unassigns the operational APs from the floor, except provisioned APs, which Catalyst Center does not allow to be unassigned; the planned APs are not recreated.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"floor_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Floor ID where the planned Access Points are placed").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"assignments": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Planned Access Point to operational Access Point assignments").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"planned_access_point_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the planned Access Point on the floor").String,
							Required:            true,
						},
						"access_point_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network device ID of the operational Access Point").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AssignPlannedAccessPointsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Custom create: planned Access Points are referenced by name and resolved to
// IDs before the POST, and the result is verified against the floor because the
// task reports success even when Catalyst Center skips an assignment.
func (r *AssignPlannedAccessPointsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan AssignPlannedAccessPoints

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	if err := r.assign(ctx, plan, plan.Assignments); err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.FloorId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AssignPlannedAccessPointsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state AssignPlannedAccessPoints

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	res, err := r.client.Get(state.getPathGet() + params)
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

// Custom update: only new pairs are posted because the planned Access Points of
// existing pairs no longer exist, and Access Points removed from the
// configuration are unassigned from the floor.
func (r *AssignPlannedAccessPointsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan, state AssignPlannedAccessPoints

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

	existing := make(map[AssignPlannedAccessPointsAssignments]bool)
	for _, item := range state.Assignments {
		existing[item] = true
	}
	var added []AssignPlannedAccessPointsAssignments
	planned := make(map[string]bool)
	for _, item := range plan.Assignments {
		planned[item.AccessPointId.ValueString()] = true
		if !existing[item] {
			added = append(added, item)
		}
	}
	var removed []string
	for _, item := range state.Assignments {
		if !item.AccessPointId.IsNull() && !planned[item.AccessPointId.ValueString()] {
			removed = append(removed, item.AccessPointId.ValueString())
		}
	}

	if len(added) > 0 {
		if err := r.assign(ctx, plan, added); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
	}
	if len(removed) > 0 {
		if err := r.unassign(state.FloorId.ValueString(), removed); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Custom delete: there is no API to undo an assignment, so the Access Points
// are unassigned from the floor instead.
func (r *AssignPlannedAccessPointsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state AssignPlannedAccessPoints

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	var ids []string
	for _, item := range state.Assignments {
		if !item.AccessPointId.IsNull() {
			ids = append(ids, item.AccessPointId.ValueString())
		}
	}
	if len(ids) > 0 {
		if err := r.unassign(state.FloorId.ValueString(), ids); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *AssignPlannedAccessPointsResource) ReadCache(ctx context.Context, req resource.ReadRequest, state AssignPlannedAccessPoints, params string) (cc.Res, error) {
	var err error
	cacheKey := "AssignPlannedAccessPoints::"

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
	res, err := r.client.Get("/dna/intent/api/v2/floors/%v/accessPointPositions" + params)
	singleRes := res
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache

// The assign API accepts at most 100 pairs per request.
const assignPlannedAccessPointsMaxItems = 100

func (r *AssignPlannedAccessPointsResource) assign(ctx context.Context, data AssignPlannedAccessPoints, items []AssignPlannedAccessPointsAssignments) error {
	floorId := data.FloorId.ValueString()
	res, err := r.client.Get(fmt.Sprintf("/dna/intent/api/v2/floors/%v/plannedAccessPointPositions", url.QueryEscape(floorId)))
	if err != nil {
		return fmt.Errorf("failed to retrieve planned access points (GET), got error: %s, %s", err, res.String())
	}
	plannedIds := make(map[string]string)
	duplicates := make(map[string]bool)
	res.Get("response").ForEach(func(_, v cc.Res) bool {
		name := v.Get("name").String()
		if _, ok := plannedIds[name]; ok {
			duplicates[name] = true
		}
		plannedIds[name] = v.Get("id").String()
		return true
	})

	var missing, ambiguous []string
	for _, item := range items {
		name := item.PlannedAccessPointName.ValueString()
		if _, ok := plannedIds[name]; !ok {
			missing = append(missing, name)
		} else if duplicates[name] {
			ambiguous = append(ambiguous, name)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("planned access points not found on floor %s: %s", floorId, strings.Join(missing, ", "))
	}
	if len(ambiguous) > 0 {
		return fmt.Errorf("planned access point names are not unique on floor %s: %s", floorId, strings.Join(ambiguous, ", "))
	}

	for start := 0; start < len(items); start += assignPlannedAccessPointsMaxItems {
		end := min(start+assignPlannedAccessPointsMaxItems, len(items))
		body := assignPlannedAccessPointsBody(items[start:end], plannedIds)
		tflog.Debug(ctx, fmt.Sprintf("%s: Assigning planned access points %d-%d of %d", floorId, start+1, end, len(items)))
		res, err := r.client.Post(data.getPath(), body, cc.UseMutex)
		if err != nil {
			return fmt.Errorf("failed to assign planned access points (POST), got error: %s, %s", err, res.String())
		}
	}

	positioned, err := positionedAccessPoints(r.client, floorId)
	if err != nil {
		return err
	}
	var failed []string
	for _, item := range items {
		if !positioned[item.AccessPointId.ValueString()] {
			failed = append(failed, fmt.Sprintf("%s (planned access point %s)", item.AccessPointId.ValueString(), item.PlannedAccessPointName.ValueString()))
		}
	}
	if len(failed) > 0 {
		return fmt.Errorf("access points were not positioned on floor %s: %s. Catalyst Center does not move an access point that is assigned to another floor", floorId, strings.Join(failed, ", "))
	}
	return nil
}

func (r *AssignPlannedAccessPointsResource) unassign(floorId string, ids []string) error {
	return unassignAccessPointsFromFloor(r.client, floorId, ids)
}

// positionedAccessPoints returns the IDs of all access points assigned to the
// floor, mapped to whether they are placed on the floor map. Catalyst Center
// reports x and y as -1 for an access point that has no position.
func positionedAccessPoints(client *cc.Client, floorId string) (map[string]bool, error) {
	res, err := client.Get(fmt.Sprintf("/dna/intent/api/v2/floors/%v/accessPointPositions", url.QueryEscape(floorId)))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve access point positions (GET), got error: %s, %s", err, res.String())
	}
	positioned := make(map[string]bool)
	res.Get("response").ForEach(func(_, v cc.Res) bool {
		positioned[v.Get("id").String()] = v.Get("position.x").Float() != -1 || v.Get("position.y").Float() != -1
		return true
	})
	return positioned, nil
}

// unassignAccessPointsFromFloor skips access points that are not assigned to
// the floor, such as an access point that failed to move from another floor,
// and leaves provisioned access points on the floor because Catalyst Center
// refuses to unassign them and rejects the whole batch if it contains one.
func unassignAccessPointsFromFloor(client *cc.Client, floorId string, ids []string) error {
	onFloor, err := positionedAccessPoints(client, floorId)
	if err != nil {
		return err
	}
	var deviceIds []string
	for _, id := range ids {
		if _, ok := onFloor[id]; ok {
			deviceIds = append(deviceIds, id)
		}
	}
	if len(deviceIds) == 0 {
		return nil
	}
	path := "/dna/intent/api/v1/networkDevices/unassignFromSite/apply"
	res, err := client.Post(path, unassignFromSiteBody(deviceIds), cc.UseMutex)
	if err == nil {
		return nil
	}
	if !isUnassignableDevice(res) {
		return fmt.Errorf("failed to unassign access points from floor (POST), got error: %s, %s", err, res.String())
	}
	if len(deviceIds) > 1 {
		for _, id := range deviceIds {
			res, err := client.Post(path, unassignFromSiteBody([]string{id}), cc.UseMutex)
			if err != nil && !isUnassignableDevice(res) {
				return fmt.Errorf("failed to unassign access point %s from floor (POST), got error: %s, %s", id, err, res.String())
			}
		}
	}
	return nil
}
