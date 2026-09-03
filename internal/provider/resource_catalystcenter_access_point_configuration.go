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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AccessPointConfigurationResource{}

func NewAccessPointConfigurationResource() resource.Resource {
	return &AccessPointConfigurationResource{}
}

type AccessPointConfigurationResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *AccessPointConfigurationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_point_configuration"
}

func (r *AccessPointConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource configures one or more access points using the Catalyst Center \"Configure Access Points V1\" intent API. Access points are selected by their ethernet MAC address. <p/> The endpoint is an asynchronous action API: each create issues a POST and waits for the resulting task to complete. There is no corresponding GET, so this resource cannot be read back, refreshed or imported, and destroying it does not revert any configuration on the access points. <p/> Every value is applied only when its matching `configure_*` flag is set to `true`. For example `admin_status` is ignored unless `configure_admin_status` is `true`. <p/> Note: this API does not support configuration of CleanAir or Spectrum Intelligence for IOS-XE devices running version 17.9 or later.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ap_list": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of access points to configure.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ethernet MAC address of the access point.").String,
							Required:            true,
						},
						"ap_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The current host name of the access point.").String,
							Optional:            true,
						},
						"ap_name_new": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The modified hostname of the access point.").String,
							Optional:            true,
						},
					},
				},
			},
			"configure_admin_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's admin status, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"admin_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the access point's admin status. Set this parameter's value to `true` to enable it and `false` to disable it.").String,
				Optional:            true,
			},
			"configure_ap_mode": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's mode, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"ap_mode": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the access point's mode: for local/flexconnect mode, set `0`; for monitor mode, set `1`; for sniffer mode, set `4`; and for bridge/flex+bridge mode, set `5`.").AddStringEnumDescription("0", "1", "4", "5").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.OneOf(0, 1, 4, 5),
				},
			},
			"configure_failover_priority": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's failover priority, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"failover_priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the access point's failover priority: for low, set `1`; for medium, set `2`; for high, set `3`; and for critical, set `4`.").AddStringEnumDescription("1", "2", "3", "4").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.OneOf(1, 2, 3, 4),
				},
			},
			"configure_led_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's LED status, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"led_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the access point's LED status. Set `true` to enable its status and `false` to disable it.").String,
				Optional:            true,
			},
			"configure_led_brightness_level": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's LED brightness level, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"led_brightness_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the access point's LED brightness level by setting a value between 1 and 8.").AddIntegerRangeDescription(1, 8).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8),
				},
			},
			"configure_location": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's location, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the access point's location.").String,
				Optional:            true,
			},
			"configure_ha_controller": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("To change the access point's HA controller, set this parameter's value to `true`.").String,
				Optional:            true,
			},
			"primary_controller_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the hostname for an access point's primary controller.").String,
				Optional:            true,
			},
			"primary_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the IP address for an access point's primary controller.").String,
				Optional:            true,
			},
			"secondary_controller_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the hostname for an access point's secondary controller.").String,
				Optional:            true,
			},
			"secondary_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the IP address for an access point's secondary controller.").String,
				Optional:            true,
			},
			"tertiary_controller_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the hostname for an access point's tertiary controller.").String,
				Optional:            true,
			},
			"tertiary_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the IP address for an access point's tertiary controller.").String,
				Optional:            true,
			},
			"is_assigned_site_as_location": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If the access point is assigned to a site, set this parameter's value to `true` to assign the access point location as the site name.").String,
				Optional:            true,
			},
			"radio_configurations": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio parameters configuration for the selected access points. Note: the API schema marks this element as required alongside `ap_list`; some Catalyst Center versions may reject a request that omits it, in which case pass an empty list.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"radio_band": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the band on the specified radio for an access point: for 2.4 GHz, set `RADIO24`; for 5 GHz, set `RADIO5`; for 6 GHz, set `RADIO6`. Any other string is invalid, including empty string.").AddStringEnumDescription("RADIO24", "RADIO5", "RADIO6").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("RADIO24", "RADIO5", "RADIO6"),
							},
						},
						"radio_type": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure an access point's radio band: for 2.4 GHz, set `1`; for 5 GHz, set `2`; for XOR, set `3`; and for 6 GHz, set `6`.").AddStringEnumDescription("1", "2", "3", "6").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.OneOf(1, 2, 3, 6),
							},
						},
						"configure_radio_role_assignment": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the radio role on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"radio_role_assignment": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure only one of the following roles on the specified radio for an access point as `auto`, `serving`, or `monitor`. Any other string is invalid, including empty string.").AddStringEnumDescription("auto", "serving", "monitor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("auto", "serving", "monitor"),
							},
						},
						"configure_admin_status": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the admin status on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"admin_status": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the admin status on the specified radio for an access point. Set this parameter's value to `true` to enable it and `false` to disable it.").String,
							Optional:            true,
						},
						"configure_antenna_pattern_name": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the antenna gain on the specified radio for an access point, set the value for this parameter to `true`.").String,
							Optional:            true,
						},
						"antenna_pattern_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.").String,
							Optional:            true,
						},
						"antenna_gain": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure `antenna_gain`, set `antenna_pattern_name` value to `other`. The External Antenna Gain value is applied in 0.5 dBi increments on the controller, therefore the value entered is multiplied by 2 to configure the absolute gain value.").AddIntegerRangeDescription(0, 20).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 20),
							},
						},
						"configure_antenna_cable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the antenna cable name on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"antenna_cable_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to `other`.").String,
							Optional:            true,
						},
						"cable_loss": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).").String,
							Optional:            true,
						},
						"configure_channel": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the channel on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"channel_assignment_mode": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the channel assignment mode on the specified radio for an access point: for global mode, set `1`; and for custom mode, set `2`.").AddStringEnumDescription("1", "2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.OneOf(1, 2),
							},
						},
						"channel_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the channel number on the specified radio for an access point.").String,
							Optional:            true,
						},
						"configure_channel_width": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the channel width on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"channel_width": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the channel width on the specified radio for an access point: for 20 MHz, set `3`; for 40 MHz, set `4`; for 80 MHz, set `5`; for 160 MHz, set `6`; and for 320 MHz, set `7`.").AddStringEnumDescription("3", "4", "5", "6", "7").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.OneOf(3, 4, 5, 6, 7),
							},
						},
						"configure_power": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To change the power assignment mode on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"power_assignment_mode": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the power assignment mode on the specified radio for an access point: for global mode, set `1`; and for custom mode, set `2`.").AddStringEnumDescription("1", "2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.OneOf(1, 2),
							},
						},
						"power_level": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the power level on the specified radio for an access point by setting a value between 1 and 8.").AddIntegerRangeDescription(1, 8).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 8),
							},
						},
						"configure_clean_air_si": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("To enable or disable either CleanAir or Spectrum Intelligence on the specified radio for an access point, set this parameter's value to `true`.").String,
							Optional:            true,
						},
						"clean_air_si": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure CleanAir or Spectrum Intelligence on the specified radio for an access point. Set this parameter's value to `0` to disable the feature or `1` to enable it.").AddStringEnumDescription("0", "1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.OneOf(0, 1),
							},
						},
					},
				},
			},
		},
	}
}

func (r *AccessPointConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *AccessPointConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccessPointConfiguration

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AccessPointConfiguration{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body, cc.UseMutex)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *AccessPointConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AccessPointConfiguration

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
func (r *AccessPointConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AccessPointConfiguration

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
func (r *AccessPointConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AccessPointConfiguration

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
func (r *AccessPointConfigurationResource) ReadCache(ctx context.Context, req resource.ReadRequest, state AccessPointConfiguration, params string) (cc.Res, error) {
	var err error
	cacheKey := "AccessPointConfiguration::"

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
