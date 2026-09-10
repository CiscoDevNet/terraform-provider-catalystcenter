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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &LANAutomationLinkResource{}

func NewLANAutomationLinkResource() resource.Resource {
	return &LANAutomationLinkResource{}
}

type LANAutomationLinkResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *LANAutomationLinkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lan_automation_link"
}

func (r *LANAutomationLinkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource adds or deletes a Layer 3 link between two LAN-automated devices using the DAY-N updateDevice API. Use action ADD_LINK to create a link or DELETE_LINK to remove one. When the resource is destroyed or refreshed, no actions are performed on Catalyst Center.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Link operation to perform. ADD_LINK creates a Layer 3 link; DELETE_LINK removes an existing link between the specified interfaces.").AddStringEnumDescription("ADD_LINK", "DELETE_LINK").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ADD_LINK", "DELETE_LINK"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"primary_device_management_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary (source) device management IP address.").String,
				Required:            true,
			},
			"primary_device_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary (source) device interface name.").String,
				Required:            true,
			},
			"peer_device_management_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer (destination) device management IP address.").String,
				Required:            true,
			},
			"peer_device_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer (destination) device interface name.").String,
				Required:            true,
			},
			"ip_pool_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the IP pool used for link add. Required when action is ADD_LINK; omit for DELETE_LINK.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *LANAutomationLinkResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

func lanAutomationLinkFeature(action string) (string, error) {
	switch action {
	case "ADD_LINK":
		return "LINK_ADD", nil
	case "DELETE_LINK":
		return "LINK_DELETE", nil
	default:
		return "", fmt.Errorf("unsupported action %q, expected ADD_LINK or DELETE_LINK", action)
	}
}

// Custom Create: map action to the feature query param and build a composite resource ID.
func (r *LANAutomationLinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LANAutomationLink

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	feature, err := lanAutomationLinkFeature(plan.Action.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	body := plan.toBody(ctx, LANAutomationLink{})
	params := "?feature=" + url.QueryEscape(feature)

	res, err := r.client.Put(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
		return
	}

	plan.Id = types.StringValue(strings.Join([]string{
		plan.Action.ValueString(),
		plan.PrimaryDeviceManagementIpAddress.ValueString(),
		plan.PrimaryDeviceInterfaceName.ValueString(),
		plan.PeerDeviceManagementIpAddress.ValueString(),
		plan.PeerDeviceInterfaceName.ValueString(),
	}, "/"))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *LANAutomationLinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state LANAutomationLink

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *LANAutomationLinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan, state LANAutomationLink

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *LANAutomationLinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state LANAutomationLink

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *LANAutomationLinkResource) ReadCache(ctx context.Context, req resource.ReadRequest, state LANAutomationLink, params string) (cc.Res, error) {
	var err error
	cacheKey := "LANAutomationLink::"

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
