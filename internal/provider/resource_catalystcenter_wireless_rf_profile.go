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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WirelessRFProfileResource{}
var _ resource.ResourceWithImportState = &WirelessRFProfileResource{}

func NewWirelessRFProfileResource() resource.Resource {
	return &WirelessRFProfileResource{}
}

type WirelessRFProfileResource struct {
	client *cc.Client
}

func (r *WirelessRFProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profile"
}

func (r *WirelessRFProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource manages a Wireless RF Profile. <p/> Updating or re-creating this resource might lead to subsequent failures when modifying the resource due to a known API issue.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RF Profile Name").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"default_rf_profile": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("is Default Rf Profile").String,
				Required:            true,
			},
			"enable_radio_type_a": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Radio Type A").String,
				Required:            true,
			},
			"enable_radio_type_b": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Radio Type B").String,
				Required:            true,
			},
			"enable_radio_type_c": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Radio Type C (6GHz)").String,
				Optional:            true,
			},
			"channel_width": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Channel Width").String,
				Required:            true,
			},
			"enable_custom": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Custom").String,
				Required:            true,
			},
			"enable_brown_field": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Brown Field").String,
				Required:            true,
			},
			"radio_type_a_parent_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Parent Profile").String,
				Optional:            true,
			},
			"radio_type_a_radio_channels": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Radio Channels").String,
				Optional:            true,
			},
			"radio_type_a_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Data Rates").String,
				Optional:            true,
			},
			"radio_type_a_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Mandatory Data Rates").String,
				Optional:            true,
			},
			"radio_type_a_power_treshold_v1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Power Threshold V1").String,
				Optional:            true,
			},
			"radio_type_a_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Rx Sop Threshold").String,
				Optional:            true,
			},
			"radio_type_a_min_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Min Power Level").String,
				Optional:            true,
			},
			"radio_type_a_max_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeA Properties - Max Power Level").String,
				Optional:            true,
			},
			"radio_type_b_parent_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Parent Profile").String,
				Optional:            true,
			},
			"radio_type_b_radio_channels": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Radio Channels").String,
				Optional:            true,
			},
			"radio_type_b_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Data Rates").String,
				Optional:            true,
			},
			"radio_type_b_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Mandatory Data Rates").String,
				Optional:            true,
			},
			"radio_type_b_power_treshold_v1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Power Threshold V1").String,
				Optional:            true,
			},
			"radio_type_b_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Rx Sop Threshold").String,
				Optional:            true,
			},
			"radio_type_b_min_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Min Power Level").String,
				Optional:            true,
			},
			"radio_type_b_max_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeB Properties - Max Power Level").String,
				Optional:            true,
			},
			"radio_type_c_parent_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Parent Profile").String,
				Optional:            true,
			},
			"radio_type_c_radio_channels": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Radio Channels").String,
				Optional:            true,
			},
			"radio_type_c_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Data Rates").String,
				Optional:            true,
			},
			"radio_type_c_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Mandatory Data Rates").String,
				Optional:            true,
			},
			"radio_type_c_power_treshold_v1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Power Threshold V1").String,
				Optional:            true,
			},
			"radio_type_c_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Rx Sop Threshold").String,
				Optional:            true,
			},
			"radio_type_c_min_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Min Power Level").String,
				Optional:            true,
			},
			"radio_type_c_max_power_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Radio TypeC Properties - Max Power Level").String,
				Optional:            true,
			},
		},
	}
}

func (r *WirelessRFProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *WirelessRFProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessRFProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessRFProfile{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.Name.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *WirelessRFProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessRFProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?rf-profile-name=" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
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
func (r *WirelessRFProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessRFProfile

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
	res, err := r.client.Post(plan.getPath()+params, body)
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
func (r *WirelessRFProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessRFProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *WirelessRFProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <name>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[0])...)
}

// End of section. //template:end import
