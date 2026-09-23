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

// Custom imports: this resource hand-writes read-modify-write CRUD (see below), so the import
// block is maintained manually and intentionally has no template markers. `types` is required by
// the custom CRUD; `net/url` is still used by the generated ReadCache section.
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WirelessProfileSiteTagSiteAssignmentResource{}
var _ resource.ResourceWithImportState = &WirelessProfileSiteTagSiteAssignmentResource{}

func NewWirelessProfileSiteTagSiteAssignmentResource() resource.Resource {
	return &WirelessProfileSiteTagSiteAssignmentResource{}
}

type WirelessProfileSiteTagSiteAssignmentResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *WirelessProfileSiteTagSiteAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_profile_site_tag_site_assignment"
}

func (r *WirelessProfileSiteTagSiteAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages the membership of a single site in a Wireless Profile Site Tag using a read-modify-write approach, so that the same Site Tag can be shared across independent Terraform states (for example, multistate deployments where each site is managed separately). Each instance of this resource owns exactly one `site_id` inside the Site Tag identified by `site_tag_name`.<p/> On create the resource looks up the Site Tag by name under the given Wireless Profile: if the tag does not exist yet it is created (first-touch) using `ap_profile_name` and `flex_profile_name`; if it already exists the `site_id` is appended to the existing `site_ids` without disturbing the sites contributed by other states. On destroy the resource removes only its own `site_id`; when it removes the last remaining site the Site Tag itself is deleted. Because the underlying Catalyst Center API replaces the whole `siteIds` list on every write, concurrent applies against the same Site Tag are serialized with a mutex; run the owning states sequentially to avoid lost updates.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"wireless_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the Wireless Profile").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_tag_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Site Tag. Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space. Cannot be modified after creation.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The single Site ID (Area, Building, or Floor - not Global) to attach to the Site Tag. Assigning a parent site also applies the tag to child sites.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ap_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the AP Profile to associate with this Site Tag. Only used when this resource has to create the Site Tag (first-touch); ignored when the tag already exists.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"flex_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Flex Profile. Only used when this resource has to create the Site Tag (first-touch); ignored when the tag already exists.").AddDefaultValueDescription("default-flex-profile").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-flex-profile"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_tag_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the Site Tag (resolved by name after create).").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *WirelessProfileSiteTagSiteAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Custom section: the whole CRUD is hand-written to implement read-modify-write membership of a
// single site inside a shared Site Tag (see resource description). Template markers are removed on
// the create/read/update/delete/import sections so `go generate` leaves them untouched.

// siteTagInfo holds the fields read back from GET .../siteTags?siteTagName=<name>.
type siteTagInfo struct {
	found           bool
	siteTagId       string
	apProfileName   string
	flexProfileName string
	siteIds         []string
}

// siteTagSiteIdPresent reports whether siteId is already a member of the tag.
func siteTagSiteIdPresent(siteIds []string, siteId string) bool {
	for _, s := range siteIds {
		if s == siteId {
			return true
		}
	}
	return false
}

// findTagByName resolves the Site Tag by name under the given Wireless Profile. The list endpoint
// supports server-side filtering by siteTagName; the name is still re-checked defensively.
func (r *WirelessProfileSiteTagSiteAssignmentResource) findTagByName(data WirelessProfileSiteTagSiteAssignment) (siteTagInfo, error) {
	var info siteTagInfo
	res, err := r.client.Get(data.getPathListByName())
	if err != nil {
		return info, err
	}
	tag := res.Get("response.0")
	if tag.Exists() && tag.Get("siteTagName").String() == data.SiteTagName.ValueString() {
		info.found = true
		info.siteTagId = tag.Get("siteTagId").String()
		info.apProfileName = tag.Get("apProfileName").String()
		info.flexProfileName = tag.Get("flexProfileName").String()
		for _, s := range tag.Get("siteIds").Array() {
			info.siteIds = append(info.siteIds, s.String())
		}
	}
	return info, nil
}

// ensureMembership makes sure the plan's site_id is a member of the Site Tag, creating the tag on
// first touch, and returns the resolved site_tag_id. All writes are serialized with the mutex.
func (r *WirelessProfileSiteTagSiteAssignmentResource) ensureMembership(plan WirelessProfileSiteTagSiteAssignment) (string, error) {
	info, err := r.findTagByName(plan)
	if err != nil {
		return "", err
	}

	if !info.found {
		// First touch: create the Site Tag with just this site.
		body := plan.createTagBody([]string{plan.SiteId.ValueString()})
		if res, err := r.client.Post(plan.getPathBulk(), body, cc.UseMutex); err != nil {
			return "", fmt.Errorf("%s, %s", err, res.String())
		}
		// Resolve the newly created site_tag_id.
		created, err := r.findTagByName(plan)
		if err != nil {
			return "", err
		}
		if !created.found {
			return "", fmt.Errorf("site tag %q was created but could not be found on read-back", plan.SiteTagName.ValueString())
		}
		return created.siteTagId, nil
	}

	// Tag exists: append our site if it is not already a member, preserving the tag's existing
	// profile attributes and the sites contributed by other states.
	if !siteTagSiteIdPresent(info.siteIds, plan.SiteId.ValueString()) {
		newSiteIds := append(append([]string{}, info.siteIds...), plan.SiteId.ValueString())
		body := putTagBody(plan.SiteTagName.ValueString(), info.apProfileName, info.flexProfileName, newSiteIds)
		if res, err := r.client.Put(plan.getPathTag(info.siteTagId), body, cc.UseMutex); err != nil {
			return "", fmt.Errorf("%s, %s", err, res.String())
		}
	}
	return info.siteTagId, nil
}

func (r *WirelessProfileSiteTagSiteAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessProfileSiteTagSiteAssignment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getAssignmentId()))

	siteTagId, err := r.ensureMembership(plan)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to assign site to site tag, got error: %s", err))
		return
	}

	plan.SiteTagId = types.StringValue(siteTagId)
	plan.Id = types.StringValue(plan.getAssignmentId())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WirelessProfileSiteTagSiteAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessProfileSiteTagSiteAssignment

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	info, err := r.findTagByName(state)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	// The assignment no longer exists if the tag is gone or our site was removed from it out of band.
	if !info.found || !siteTagSiteIdPresent(info.siteIds, state.SiteId.ValueString()) {
		resp.State.RemoveResource(ctx)
		return
	}

	state.SiteTagId = types.StringValue(info.siteTagId)
	// ap_profile_name / flex_profile_name are creation hints. Populate them from the controller only
	// when unset (import), so a normal refresh never fights the state that owns the tag's profiles.
	if state.ApProfileName.IsNull() {
		state.ApProfileName = types.StringValue(info.apProfileName)
	}
	if state.FlexProfileName.IsNull() {
		if info.flexProfileName != "" {
			state.FlexProfileName = types.StringValue(info.flexProfileName)
		} else {
			state.FlexProfileName = types.StringValue("default-flex-profile")
		}
	}
	if state.Id.IsNull() || state.Id.ValueString() == "" {
		state.Id = types.StringValue(state.getAssignmentId())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update is effectively only reached for computed/plan-only churn: wireless_profile_id,
// site_tag_name, site_id, ap_profile_name and flex_profile_name all force replacement. It re-runs
// the idempotent ensure so membership self-heals if the tag drifted.
func (r *WirelessProfileSiteTagSiteAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WirelessProfileSiteTagSiteAssignment

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getAssignmentId()))

	siteTagId, err := r.ensureMembership(plan)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to assign site to site tag, got error: %s", err))
		return
	}

	plan.SiteTagId = types.StringValue(siteTagId)
	plan.Id = types.StringValue(plan.getAssignmentId())

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Delete removes only this resource's own site_id from the Site Tag. If it was the last member the
// whole Site Tag is deleted; otherwise the remaining membership (contributed by other states) is
// rewritten via PUT. Writes are serialized with the mutex.
func (r *WirelessProfileSiteTagSiteAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessProfileSiteTagSiteAssignment

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	info, err := r.findTagByName(state)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	if !info.found {
		// Tag already gone.
		resp.State.RemoveResource(ctx)
		return
	}

	remainder := make([]string, 0, len(info.siteIds))
	for _, s := range info.siteIds {
		if s != state.SiteId.ValueString() {
			remainder = append(remainder, s)
		}
	}

	if len(remainder) == 0 {
		// We owned the last site: delete the whole Site Tag.
		res, err := r.client.Delete(state.getPathTag(info.siteTagId), cc.UseMutex)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		// Rewrite the tag without our site, preserving the other states' membership and profiles.
		body := putTagBody(state.SiteTagName.ValueString(), info.apProfileName, info.flexProfileName, remainder)
		res, err := r.client.Put(state.getPathTag(info.siteTagId), body, cc.UseMutex)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove site from site tag (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// ImportState takes a composite <wireless_profile_id>,<site_tag_name>,<site_id> identifier; Read
// then resolves site_tag_id, ap_profile_name and flex_profile_name from the controller.
func (r *WirelessProfileSiteTagSiteAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <wireless_profile_id>,<site_tag_name>,<site_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("wireless_profile_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_tag_name"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0]+"/"+idParts[1]+"/"+idParts[2])...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *WirelessProfileSiteTagSiteAssignmentResource) ReadCache(ctx context.Context, req resource.ReadRequest, state WirelessProfileSiteTagSiteAssignment, params string) (cc.Res, error) {
	var err error
	cacheKey := "WirelessProfileSiteTagSiteAssignment::"

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
