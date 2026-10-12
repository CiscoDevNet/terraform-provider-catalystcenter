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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AssignDeviceToSiteResource{}
var _ resource.ResourceWithImportState = &AssignDeviceToSiteResource{}

func NewAssignDeviceToSiteResource() resource.Resource {
	return &AssignDeviceToSiteResource{}
}

type AssignDeviceToSiteResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AssignDeviceToSiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assign_device_to_site"
}

func (r *AssignDeviceToSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Assign unprovisioned network devices to a site.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unassigned network device ID").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("This must be building Id or floor Id. Access points, Sensors are assigned to floor. Remaining network devices are assigned to building.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *AssignDeviceToSiteResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

func (r *AssignDeviceToSiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan AssignDeviceToSite

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AssignDeviceToSite{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		// NCIM80106 means Catalyst Center refuses to move an already provisioned
		// device. A device onboarded through PnP hits this on its first apply,
		// because the site-claim both assigned and provisioned it, and so does a
		// brownfield device Terraform has not assigned itself. That is only
		// benign when the device already sits on the requested site — adopt it
		// then, and report a move that cannot happen otherwise. Delete ignores
		// the same error unconditionally.
		if !strings.Contains(res.Get("response.detail").String(), "NCIM80106") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			return
		}
		var misplaced []string
		for _, deviceId := range plan.DeviceIds.Elements() {
			id := deviceId.(types.String).ValueString()
			currentSite, siteErr := r.assignedSite(id)
			if siteErr != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve the current site of device %s (GET), got error: %s", id, siteErr))
				return
			}
			if currentSite != plan.SiteId.ValueString() {
				misplaced = append(misplaced, fmt.Sprintf("%s (currently on site %q)", id, currentSite))
			}
		}
		if len(misplaced) > 0 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Cannot assign provisioned device(s) to site %s: %s. Catalyst Center does not allow moving a provisioned device between sites; unprovision it first.", plan.SiteId.ValueString(), strings.Join(misplaced, ", ")))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Device(s) already assigned to the requested site, adopting the existing assignment", plan.SiteId.ValueString()))
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// assignedSite returns the id of the site a device is currently assigned to, or
// an empty string when it is unassigned.
func (r *AssignDeviceToSiteResource) assignedSite(deviceId string) (string, error) {
	res, err := r.client.Get("/dna/intent/api/v1/networkDevices/" + url.QueryEscape(deviceId) + "/assignedToSite")
	if err != nil {
		return "", err
	}
	return res.Get("response.siteId").String(), nil
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AssignDeviceToSiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state AssignDeviceToSite

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?siteId=" + url.QueryEscape(state.SiteId.ValueString())
	res, err := r.client.Get("/dna/intent/api/v1/networkDevices/assignedToSite" + params)
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

func (r *AssignDeviceToSiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan, state AssignDeviceToSite

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

	var planDeviceIds []string
	plan.DeviceIds.ElementsAs(ctx, &planDeviceIds, false)

	var stateDeviceIds []string
	state.DeviceIds.ElementsAs(ctx, &stateDeviceIds, false)

	planMap := make(map[string]struct{})
	for _, id := range planDeviceIds {
		planMap[id] = struct{}{}
	}

	stateMap := make(map[string]struct{})
	for _, id := range stateDeviceIds {
		stateMap[id] = struct{}{}
	}

	// Find added (in plan, not in state)
	var added []string
	for id := range planMap {
		if _, exists := stateMap[id]; !exists {
			added = append(added, id)
		}
	}

	// Find removed (in state, not in plan)
	var removed []string
	for id := range stateMap {
		if _, exists := planMap[id]; !exists {
			removed = append(removed, id)
		}
	}

	// Assign devices to site based on added devices
	if len(added) > 0 {
		body := ""
		for i, id := range added {
			body, _ = sjson.Set(body, fmt.Sprintf("deviceIds.%d", i), id)
		}
		body, _ = sjson.Set(body, "siteId", plan.SiteId.ValueString())

		res, err := r.client.Post(plan.getPath(), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("Device added: %s", added))
	}

	// Unassign devices from site based on removed devices
	if len(removed) > 0 {
		body := ""
		for i, id := range removed {
			body, _ = sjson.Set(body, fmt.Sprintf("deviceIds.%d", i), id)
		}

		res, err := r.client.Post(plan.getPathDelete(), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST2), got error: %s, %s", err, res.String()))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("Device removed: %s", removed))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AssignDeviceToSiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state AssignDeviceToSite

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	body := state.toBody(ctx, state)
	params := ""
	res, err := r.client.Post(state.getPathDelete()+params, body, cc.UseMutex)
	if err != nil {
		errorDetail := res.Get("response.detail").String()
		if strings.HasPrefix(errorDetail, "NCIM80106") || errorDetail == "Input data is invalid" {
			// This error means that the device was provisioned, and cannot be moved, so we need to ignore it.
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *AssignDeviceToSiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <site_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
}

// End of section. //template:end import
