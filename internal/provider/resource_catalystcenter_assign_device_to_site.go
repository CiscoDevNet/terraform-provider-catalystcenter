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
	var deviceIds []string
	plan.DeviceIds.ElementsAs(ctx, &deviceIds, false)

	if err := r.assignToSite(ctx, plan, deviceIds); err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}

	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// assignedSite returns the id and the hierarchy name of the site a device is
// currently assigned to. Both are empty when the device is unassigned. The id
// is what can be compared against site_id, the name is what is worth putting in
// front of an operator.
func (r *AssignDeviceToSiteResource) assignedSite(deviceId string) (string, string, error) {
	res, err := r.client.Get("/dna/intent/api/v1/networkDevices/" + url.QueryEscape(deviceId) + "/assignedToSite")
	if err != nil {
		return "", "", err
	}
	return res.Get("response.siteId").String(), res.Get("response.siteNameHierarchy").String(), nil
}

// isProvisionedDeviceError reports whether Catalyst Center refused a request
// because it covers a device that has already been provisioned, which it does
// not allow to be moved between sites.
func isProvisionedDeviceError(res cc.Res) bool {
	return strings.Contains(res.Get("response.detail").String(), "NCIM80106")
}

// isUnassignableDevice reports whether Catalyst Center refused to unassign a
// device because it has been provisioned. "Input data is invalid" is an older
// phrasing of the same refusal.
func isUnassignableDevice(res cc.Res) bool {
	return isProvisionedDeviceError(res) || res.Get("response.detail").String() == "Input data is invalid"
}

func assignToSiteBody(siteId string, deviceIds []string) string {
	body := ""
	body, _ = sjson.Set(body, "deviceIds", deviceIds)
	body, _ = sjson.Set(body, "siteId", siteId)
	return body
}

func unassignFromSiteBody(deviceIds []string) string {
	body := ""
	body, _ = sjson.Set(body, "deviceIds", deviceIds)
	return body
}

// unassignFromSite removes devices from their site. A provisioned device cannot
// be unassigned, and Catalyst Center rejects the whole request because of it,
// so ignoring that error for a batch would silently leave every other device in
// the batch assigned while Terraform reported them removed. The batch is
// therefore retried one device at a time, which skips only the devices that
// genuinely cannot be unassigned and needs the extra requests solely in the
// mixed case.
func (r *AssignDeviceToSiteResource) unassignFromSite(ctx context.Context, data AssignDeviceToSite, deviceIds []string) error {
	if len(deviceIds) == 0 {
		return nil
	}

	res, err := r.client.Post(data.getPathDelete(), unassignFromSiteBody(deviceIds), cc.UseMutex)
	if err == nil {
		return nil
	}
	if !isUnassignableDevice(res) {
		return fmt.Errorf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String())
	}

	if len(deviceIds) > 1 {
		for _, deviceId := range deviceIds {
			res, err := r.client.Post(data.getPathDelete(), unassignFromSiteBody([]string{deviceId}), cc.UseMutex)
			if err == nil {
				continue
			}
			if !isUnassignableDevice(res) {
				return fmt.Errorf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String())
			}
			tflog.Debug(ctx, fmt.Sprintf("%s: Device %s has been provisioned and cannot be unassigned, leaving it on its site", data.Id.ValueString(), deviceId))
		}
		return nil
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Device %s has been provisioned and cannot be unassigned, leaving it on its site", data.Id.ValueString(), deviceIds[0]))
	return nil
}

// assignToSite assigns devices to a site, tolerating the NCIM80106 Catalyst
// Center returns when the request covers a device it has already provisioned.
// Such a device cannot be moved, but needs no action when it already sits on
// the requested site, which is the normal state after a Plug and Play
// site-claim because the claim both assigns and provisions. The endpoint
// rejects the whole batch rather than applying it partially, and the error does
// not identify the offending device, so the devices that still need assigning
// are established and, if there are any, assigned on their own. A batch mixing
// a provisioned controller with devices Terraform has not assigned yet
// therefore still completes. Only a device that has to move and cannot is
// reported as an error.
func (r *AssignDeviceToSiteResource) assignToSite(ctx context.Context, data AssignDeviceToSite, deviceIds []string) error {
	if len(deviceIds) == 0 {
		return nil
	}

	siteId := data.SiteId.ValueString()

	res, err := r.client.Post(data.getPath(), assignToSiteBody(siteId, deviceIds), cc.UseMutex)
	if err == nil {
		return nil
	}
	if !isProvisionedDeviceError(res) {
		return fmt.Errorf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String())
	}

	var pending, pendingDetail []string
	for _, deviceId := range deviceIds {
		currentSiteId, currentSiteName, siteErr := r.assignedSite(deviceId)
		if siteErr != nil {
			return fmt.Errorf("Failed to retrieve the current site of device %s (GET), got error: %s", deviceId, siteErr)
		}
		if currentSiteId != siteId {
			where := "currently unassigned"
			if currentSiteName != "" {
				where = fmt.Sprintf("currently on site %s", currentSiteName)
			}
			pending = append(pending, deviceId)
			pendingDetail = append(pendingDetail, fmt.Sprintf("%s (%s)", deviceId, where))
		}
	}

	if len(pending) == 0 {
		tflog.Debug(ctx, fmt.Sprintf("%s: Device(s) already assigned to the requested site, adopting the existing assignment", siteId))
		return nil
	}

	res, err = r.client.Post(data.getPath(), assignToSiteBody(siteId, pending), cc.UseMutex)
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("%s: Adopted the existing assignment of the provisioned device(s) and assigned the remaining %d device(s)", siteId, len(pending)))
		return nil
	}
	if !isProvisionedDeviceError(res) {
		return fmt.Errorf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String())
	}
	return fmt.Errorf("Cannot assign device(s) to site %s: %s. At least one of them has been provisioned, and Catalyst Center does not allow moving a provisioned device between sites; unprovision it first", siteId, strings.Join(pendingDetail, ", "))
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
		if err := r.assignToSite(ctx, plan, added); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("Device added: %s", added))
	}

	// Unassign devices from site based on removed devices
	if len(removed) > 0 {
		if err := r.unassignFromSite(ctx, plan, removed); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
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

	var deviceIds []string
	state.DeviceIds.ElementsAs(ctx, &deviceIds, false)

	if err := r.unassignFromSite(ctx, state, deviceIds); err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
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
