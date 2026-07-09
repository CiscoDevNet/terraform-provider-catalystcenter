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
	"github.com/hashicorp/terraform-plugin-framework/path"
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
var _ resource.Resource = &WirelessCleanAirConfigurationResource{}
var _ resource.ResourceWithImportState = &WirelessCleanAirConfigurationResource{}

func NewWirelessCleanAirConfigurationResource() resource.Resource {
	return &WirelessCleanAirConfigurationResource{}
}

type WirelessCleanAirConfigurationResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *WirelessCleanAirConfigurationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_cleanair_configuration"
}

func (r *WirelessCleanAirConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Wireless CleanAir Configuration feature template. CleanAir feature templates define spectrum-intelligence settings (per radio band) that detect and classify RF interferers, and can be attached to wireless network profiles.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"design_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the CleanAir feature template. Forbidden characters are `% & < > [ ] ' /`.").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"radio_band": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The radio band the CleanAir configuration applies to.").AddStringEnumDescription("2_4GHZ", "5GHZ", "6GHZ").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2_4GHZ", "5GHZ", "6GHZ"),
				},
			},
			"clean_air": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CleanAir spectrum intelligence on the radio band.").String,
				Optional:            true,
			},
			"clean_air_device_reporting": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable reporting of CleanAir interferer devices.").String,
				Optional:            true,
			},
			"persistent_device_propagation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable persistent device propagation.").String,
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the CleanAir feature template.").String,
				Optional:            true,
			},
			"unlocked_attributes": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of feature attributes that are unlocked (editable) for this template.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ble_beacon": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect BLE Beacon interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"bluetooth_paging_inquiry": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Bluetooth Paging Inquiry interferers. Applicable to 2.4GHz.").String,
				Optional:            true,
			},
			"bluetooth_sco_acl": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Bluetooth SCO and ACL interferers. Applicable to 2.4GHz.").String,
				Optional:            true,
			},
			"continuous_transmitter": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Continuous Transmitter interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"generic_dect": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Generic DECT interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"generic_tdd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Generic TDD Transmitter interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"jammer": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Jammer interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"microwave_oven": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Microwave Oven interferers. Applicable to 2.4GHz.").String,
				Optional:            true,
			},
			"motorola_canopy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Motorola Canopy interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"si_fhss": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect SI FHSS interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"spectrum_80211_fh": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect 802.11 FH interferers. Applicable to 2.4GHz.").String,
				Optional:            true,
			},
			"spectrum_80211_non_standard_channel": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect 802.11 Non Standard Channel interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"spectrum_802154": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect 802.15.4 interferers. Applicable to 2.4GHz.").String,
				Optional:            true,
			},
			"spectrum_inverted": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Spectrum Inverted interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"super_ag": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Super AG interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"video_camera": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Video Camera interferers. Applicable to 2.4GHz and 5GHz.").String,
				Optional:            true,
			},
			"wimax_fixed": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect WiMax Fixed interferers. Applicable to 5GHz.").String,
				Optional:            true,
			},
			"wimax_mobile": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect WiMax Mobile interferers. Applicable to 5GHz.").String,
				Optional:            true,
			},
			"xbox": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Detect Xbox interferers. Applicable to 2.4GHz.").String,
				Optional:            true,
			},
		},
	}
}

func (r *WirelessCleanAirConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *WirelessCleanAirConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessCleanAirConfiguration

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessCleanAirConfiguration{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.data").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *WirelessCleanAirConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessCleanAirConfiguration

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(state.Id.ValueString())
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
func (r *WirelessCleanAirConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessCleanAirConfiguration

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body, cc.UseMutex)
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
func (r *WirelessCleanAirConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessCleanAirConfiguration

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath()+"/"+url.QueryEscape(state.Id.ValueString()), cc.UseMutex)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *WirelessCleanAirConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *WirelessCleanAirConfigurationResource) ReadCache(ctx context.Context, req resource.ReadRequest, state WirelessCleanAirConfiguration, params string) (cc.Res, error) {
	var err error
	cacheKey := "WirelessCleanAirConfiguration::"

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
