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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &APProfileResource{}
var _ resource.ResourceWithImportState = &APProfileResource{}

func NewAPProfileResource() resource.Resource {
	return &APProfileResource{}
}

type APProfileResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *APProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ap_profile"
}

func (r *APProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages an Access Point (AP) Profile in wireless settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ap_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Access Point profile. Max length is 32 characters.").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description of the AP profile. Max length is 241 characters").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 241),
				},
			},
			"remote_worker_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS, Forensic Capture Enablement, Rogue Detection and Rogue Containment.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auth_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication type used in the AP profile. These settings are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.").AddStringEnumDescription("NO-AUTH", "EAP-TLS", "EAP-PEAP", "EAP-FAST").AddDefaultValueDescription("NO-AUTH").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NO-AUTH", "EAP-TLS", "EAP-PEAP", "EAP-FAST"),
				},
				Default: stringdefault.StaticString("NO-AUTH"),
			},
			"dot1x_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"dot1x_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for 802.1X authentication. AP dot1x password length should not exceed 120.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 120),
				},
			},
			"ssh_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if SSH is enabled on the AP. Enable SSH to add credentials for device management.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"telnet_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"management_user_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management username must have a minimum of 1 character and a maximum of 32 characters.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"management_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Management password for the AP. Length must be 8-120 characters.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(8, 120),
				},
			},
			"management_enable_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable password for managing the AP. Length must be 8-120 characters.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(8, 120),
				},
			},
			"cdp_state": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"awips_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if AWIPS is enabled on the AP.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"awips_forensic_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rogue_detection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rogue_detection_min_rssi": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts").AddIntegerRangeDescription(-128, -70).AddDefaultValueDescription("-90").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(-128, -70),
				},
				Default: int64default.StaticInt64(-90),
			},
			"rogue_detection_transient_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Transient interval for rogue detection. Value should be 0 or from 120 to 1800.").AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"rogue_detection_report_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Report interval for rogue detection. Value should be in range 10 and 300.").AddIntegerRangeDescription(10, 300).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 300),
				},
				Default: int64default.StaticInt64(10),
			},
			"pmf_denial_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mesh_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"bridge_group_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.").String,
				Optional:            true,
			},
			"backhaul_client_access": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if backhaul client access is enabled on the AP.").String,
				Optional:            true,
			},
			"range": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Range of the mesh network. Value should be between 150-132000").AddIntegerRangeDescription(150, 132000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(150, 132000),
				},
			},
			"ghz5_backhaul_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("5GHz backhaul data rates.").AddStringEnumDescription("auto", "802.11abg", "802.12ac", "802.11ax", "802.11n").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "802.11abg", "802.12ac", "802.11ax", "802.11n"),
				},
			},
			"ghz24_backhaul_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("2.4GHz backhaul data rates.").AddStringEnumDescription("auto", "802.11abg", "802.11ax", "802.11n").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "802.11abg", "802.11ax", "802.11n"),
				},
			},
			"rap_downlink_backhaul": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of downlink backhaul used.").AddStringEnumDescription("5 GHz", "2.4 GHz").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("5 GHz", "2.4 GHz"),
				},
			},
			"ap_power_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the existing AP power profile.").String,
				Optional:            true,
			},
			"power_profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.").String,
				Optional:            true,
			},
			"scheduler_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of the scheduler.").AddStringEnumDescription("DAILY", "WEEKLY", "MONTHLY").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DAILY", "WEEKLY", "MONTHLY"),
				},
			},
			"scheduler_start_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Start time of the duration setting.").String,
				Optional:            true,
			},
			"scheduler_end_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("End time of the duration setting.").String,
				Optional:            true,
			},
			"scheduler_day": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Applies every week on the selected days").String,
				Optional:            true,
			},
			"scheduler_date": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Start and End date of the duration setting, applicable for MONTHLY schedulers.").String,
				Optional:            true,
			},
			"country_code": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Country code for the AP profile.").AddStringEnumDescription("AF", "AE", "AL", "AR", "AT", "AO", "AU", "BD", "BA", "BB", "BE", "BG", "BH", "BM", "BN", "BO", "BR", "BT", "BY", "CA", "CD", "CH", "CI", "CL", "CM", "CN", "CO", "CR", "CU", "CY", "CZ", "DE", "DK", "DO", "DZ", "EC", "EE", "EG", "EL", "ES", "ET", "FI", "FJ", "FR", "GB", "GH", "GI", "GE", "GR", "GT", "HK", "HN", "HR", "HU", "ID", "IE", "IL", "IN", "IQ", "IS", "IT", "J2", "J4", "JM", "JO", "KE", "KH", "KN", "KW", "KZ", "LA", "LB", "LI", "LK", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MK", "MN", "MM", "MO", "MT", "MX", "MY", "NG", "NI", "NL", "NO", "NP", "NZ", "OM", "PA", "PE", "PH", "PK", "PL", "PR", "PT", "PY", "QA", "RO", "RS", "RU", "SA", "SD", "SE", "SG", "SI", "SK", "SM", "TH", "TI", "TN", "TR", "TW", "TZ", "UA", "US", "UY", "VA", "VE", "VN", "XK", "YE", "ZA", "ZW", "MU", "ZM", "BI", "NA", "BW", "GA", "UG", "UZ").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AF", "AE", "AL", "AR", "AT", "AO", "AU", "BD", "BA", "BB", "BE", "BG", "BH", "BM", "BN", "BO", "BR", "BT", "BY", "CA", "CD", "CH", "CI", "CL", "CM", "CN", "CO", "CR", "CU", "CY", "CZ", "DE", "DK", "DO", "DZ", "EC", "EE", "EG", "EL", "ES", "ET", "FI", "FJ", "FR", "GB", "GH", "GI", "GE", "GR", "GT", "HK", "HN", "HR", "HU", "ID", "IE", "IL", "IN", "IQ", "IS", "IT", "J2", "J4", "JM", "JO", "KE", "KH", "KN", "KW", "KZ", "LA", "LB", "LI", "LK", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MK", "MN", "MM", "MO", "MT", "MX", "MY", "NG", "NI", "NL", "NO", "NP", "NZ", "OM", "PA", "PE", "PH", "PK", "PL", "PR", "PT", "PY", "QA", "RO", "RS", "RU", "SA", "SD", "SE", "SG", "SI", "SK", "SM", "TH", "TI", "TN", "TR", "TW", "TZ", "UA", "US", "UY", "VA", "VE", "VN", "XK", "YE", "ZA", "ZW", "MU", "ZM", "BI", "NA", "BW", "GA", "UG", "UZ"),
				},
			},
			"time_zone": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("In the Time Zone area, choose one of the following options. Not Configured - APs operate in the UTC time zone. Controller - APs operate in the Cisco Wireless Controller time zone. Delta from Controller - APs operate in the offset time from the wireless controller time zone.").AddStringEnumDescription("Not Configured", "Controller", "Delta from Controller").AddDefaultValueDescription("NOT CONFIGURED").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Not Configured", "Controller", "Delta from Controller"),
				},
				Default: stringdefault.StaticString("NOT CONFIGURED"),
			},
			"time_zone_offset_hour": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter the hour value (HH). The valid range is from -12 through 14.").AddIntegerRangeDescription(-12, 14).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(-12, 14),
				},
				Default: int64default.StaticInt64(0),
			},
			"time_zone_offset_minutes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enter the minute value (MM). The valid range is from 0 through 59.").AddIntegerRangeDescription(0, 59).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 59),
				},
				Default: int64default.StaticInt64(0),
			},
			"client_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of clients. Value should be between 0-1200.").AddIntegerRangeDescription(0, 1200).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1200),
				},
				Default: int64default.StaticInt64(0),
			},
		},
	}
}

func (r *APProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *APProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan APProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, APProfile{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	params = ""
	params += "?apProfileName=" + url.QueryEscape(plan.ApProfileName.ValueString())
	res, err = r.client.Get(plan.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.0.id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *APProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state APProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?apProfileName=" + url.QueryEscape(state.ApProfileName.ValueString())
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
func (r *APProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state APProfile

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString())+params, body)
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
func (r *APProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state APProfile

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
func (r *APProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <ap_profile_name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ap_profile_name"), idParts[0])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *APProfileResource) ReadCache(ctx context.Context, req resource.ReadRequest, state APProfile, params string) (cc.Res, error) {
	var err error
	cacheKey := "APProfile::"

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
