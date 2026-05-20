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
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a planned (virtual) Access Point position on a floor map for RF planning in Cisco Catalyst Center. Planned APs are Day 0 design-time resources - no physical APs are needed in the inventory.").String,

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

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PlannedAccessPointPositionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Custom Read: template markers removed for import ID resolution
func (r *PlannedAccessPointPositionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *PlannedAccessPointPositionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *PlannedAccessPointPositionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PlannedAccessPointPosition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

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
