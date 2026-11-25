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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PlannedAccessPointPositionResource{}
var _ resource.ResourceWithImportState = &PlannedAccessPointPositionResource{}

func NewPlannedAccessPointPositionResource() resource.Resource {
	return &PlannedAccessPointPositionResource{}
}

type PlannedAccessPointPositionResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *PlannedAccessPointPositionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_planned_access_point_position"
}

func (r *PlannedAccessPointPositionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a planned (virtual) Access Point position on a floor map for RF planning in Cisco Catalyst Center. Planned APs are Day 0 design-time resources - no physical APs are needed in the inventory. <p/> Set `access_point_id`, also on an existing planned AP, to assign an operational Access Point to the planned AP once it is in the inventory. The operational AP takes the planned position and is assigned to the floor, and Catalyst Center removes the planned AP. After that, refresh checks that the operational AP is still positioned on the floor, any change, including to `access_point_id`, requires replacing the resource, and destroy unassigns the operational AP from the floor unless it has been provisioned. If the assignment fails, for example because the operational AP is assigned to another floor, the planned AP is deleted again.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"floor_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Floor ID where the planned AP is placed").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the planned Access Point").String,
				Required:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MAC address of the planned Access Point").String,
				Optional:            true,
			},
			"ap_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AP model type. Use /dna/intent/api/v1/maps/supported-access-points to find supported models").String,
				Required:            true,
			},
			"position_x": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("X coordinate of the planned AP position in feet").String,
				Required:            true,
			},
			"position_y": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Y coordinate of the planned AP position in feet").String,
				Required:            true,
			},
			"position_z": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Z coordinate (height) of the planned AP position in feet").String,
				Optional:            true,
			},
			"radios": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio configurations for the planned AP (1-4 radios). Any change to any radio attribute forces destroy+create because the Catalyst Center bulkChange API has limited update support.").String,
				Required:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"radio_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Radio ID (read-only, assigned by Catalyst Center)").String,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"bands": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Radio frequency bands in GHz. Valid values are 2.4, 5, and 6").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"channel": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Channel number for map visualization").String,
							Optional:            true,
						},
						"tx_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Transmit power in dBm for map visualization").String,
							Optional:            true,
						},
						"antenna_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Antenna type name. Use /dna/intent/api/v1/maps/supported-access-points to find supported antennas").String,
							Optional:            true,
						},
						"antenna_azimuth": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Antenna azimuth angle in degrees (0-360)").String,
							Optional:            true,
						},
						"antenna_elevation": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Antenna elevation angle in degrees (-90 to 90)").String,
							Optional:            true,
						},
					},
				},
			},
			"access_point_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network device ID of an operational Access Point to assign to this planned Access Point").String,
				Optional:            true,
			},
		},
	}
}

func (r *PlannedAccessPointPositionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Custom ModifyPlan: once access_point_id is assigned, the planned access point
// no longer exists, and updating it brings it back and moves the operational
// access point to a stale position.
func (r *PlannedAccessPointPositionResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var plan, state PlannedAccessPointPosition
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || state.AccessPointId.IsNull() {
		return
	}
	changed := map[string]bool{
		"name":            !plan.Name.Equal(state.Name),
		"mac_address":     !plan.MacAddress.Equal(state.MacAddress),
		"ap_type":         !plan.ApType.Equal(state.ApType),
		"position_x":      !plan.PositionX.Equal(state.PositionX),
		"position_y":      !plan.PositionY.Equal(state.PositionY),
		"position_z":      !plan.PositionZ.Equal(state.PositionZ),
		"access_point_id": !plan.AccessPointId.Equal(state.AccessPointId),
	}
	for attr, c := range changed {
		if c {
			resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attr))
		}
	}
}

// Custom create: assigns access_point_id to the planned access point once it exists.
func (r *PlannedAccessPointPositionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan PlannedAccessPointPosition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, PlannedAccessPointPosition{})

	params := ""
	res, err := r.client.Post(plan.getPathPost()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	params += "?name=" + url.QueryEscape(plan.Name.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.0.id").String())
	plan.fromBodyUnknowns(ctx, res)

	if !plan.AccessPointId.IsNull() {
		// A failed assignment renames the planned access point after the
		// operational access point, so it is deleted instead of being tracked.
		if err := r.assignAccessPoint(plan); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			res, err := r.client.Delete(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()))
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete planned access point %s after failed assignment (DELETE), got error: %s, %s", plan.Id.ValueString(), err, res.String()))
			}
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Custom Read: template markers removed for import ID resolution
func (r *PlannedAccessPointPositionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state PlannedAccessPointPosition

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?name=" + url.QueryEscape(state.Name.ValueString())
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// Catalyst Center removes a planned access point once it is assigned, so the
	// assigned access point being positioned on the floor means it is in sync.
	if !state.AccessPointId.IsNull() && !res.Get("response.0").Exists() {
		positioned, err := positionedAccessPoints(r.client, state.FloorId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
		if !positioned[state.AccessPointId.ValueString()] {
			resp.State.RemoveResource(ctx)
		}
		return
	}

	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	if state.Id.IsNull() || state.Id.ValueString() == "" {
		state.Id = types.StringValue(res.Get("response.0.id").String())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of custom Read

// Custom update: assigns access_point_id when it is added to an existing
// planned access point.
func (r *PlannedAccessPointPositionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan, state PlannedAccessPointPosition

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
	res, err := r.client.Post(plan.getPathPut()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	if state.AccessPointId.IsNull() && !plan.AccessPointId.IsNull() {
		if err := r.assignAccessPoint(plan); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			res, err := r.client.Delete(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()))
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete planned access point %s after failed assignment (DELETE), got error: %s, %s", plan.Id.ValueString(), err, res.String()))
				return
			}
			resp.State.RemoveResource(ctx)
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Custom delete: a planned access point that was assigned no longer exists and
// its delete task fails with NCMP00102, and the assigned access point is
// unassigned from the floor.
func (r *PlannedAccessPointPositionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state PlannedAccessPointPosition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "NCMP00102") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}
	if !state.AccessPointId.IsNull() {
		if err := unassignAccessPointsFromFloor(r.client, state.FloorId.ValueString(), []string{state.AccessPointId.ValueString()}); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *PlannedAccessPointPositionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <floor_id>,<name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("floor_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[1])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *PlannedAccessPointPositionResource) ReadCache(ctx context.Context, req resource.ReadRequest, state PlannedAccessPointPosition, params string) (cc.Res, error) {
	var err error
	cacheKey := "PlannedAccessPointPosition::"

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

func (r *PlannedAccessPointPositionResource) assignAccessPoint(data PlannedAccessPointPosition) error {
	floorId := data.FloorId.ValueString()
	body, _ := sjson.Set("[]", "0.plannedAccessPointId", data.Id.ValueString())
	body, _ = sjson.Set(body, "0.accessPointId", data.AccessPointId.ValueString())
	path := fmt.Sprintf("/dna/intent/api/v2/floors/%v/plannedAccessPointPositions/assignAccessPointPositions", url.QueryEscape(floorId))
	res, err := r.client.Post(path, body, cc.UseMutex)
	if err != nil {
		return fmt.Errorf("failed to assign access point %s (POST), got error: %s, %s", data.AccessPointId.ValueString(), err, res.String())
	}
	positioned, err := positionedAccessPoints(r.client, floorId)
	if err != nil {
		return err
	}
	if !positioned[data.AccessPointId.ValueString()] {
		return fmt.Errorf("access point %s was not positioned on floor %s. Catalyst Center does not move an access point that is assigned to another floor", data.AccessPointId.ValueString(), floorId)
	}
	return nil
}
