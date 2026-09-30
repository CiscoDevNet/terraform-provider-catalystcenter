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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PnPNetworkDeviceClaimResource{}

func NewPnPNetworkDeviceClaimResource() resource.Resource {
	return &PnPNetworkDeviceClaimResource{}
}

type PnPNetworkDeviceClaimResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *PnPNetworkDeviceClaimResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pnp_network_device_claim"
}

func (r *PnPNetworkDeviceClaimResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a PnP Network Device Claim.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The PnP device ID to claim.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of device being claimed.").String,
				Required:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unique identifier of the site to claim the device to.").String,
				Optional:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hostname to be pushed to the device.").String,
				Optional:            true,
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unique identifier of the image to be installed on the device.").String,
				Optional:            true,
			},
			"remove_inactive": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Delete unused image .bin and .pkg files on the device.").String,
				Optional:            true,
			},
			"template_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unique identifier of the CLI template. Required if the device is claimed without a site.").String,
				Optional:            true,
			},
			"template_parameters": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI template parameters.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The key of the configuration parameter.").String,
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The value of the configuration parameter.").String,
							Optional:            true,
						},
					},
				},
			},
			"domain": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("SVL domain number.").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"svl_members": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of member switches in the SVL.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"serial_number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of the member switch.").String,
							Required:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Role of a member in an SVL.").AddStringEnumDescription("ACTIVE", "STANDBY").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ACTIVE", "STANDBY"),
							},
						},
						"svl_links": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of links between member switches in the SVL.").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"local_interface": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Local interface name.").String,
										Required:            true,
									},
									"remote_interface": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Remote interface name.").String,
										Required:            true,
									},
								},
							},
						},
						"local_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Local interface name of the Dual-Active-Detection (DAD) link.").String,
							Optional:            true,
						},
						"remote_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remote interface name of the Dual-Active-Detection (DAD) link.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *PnPNetworkDeviceClaimResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PnPNetworkDeviceClaimResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PnPNetworkDeviceClaim

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, PnPNetworkDeviceClaim{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.DeviceId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *PnPNetworkDeviceClaimResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PnPNetworkDeviceClaim

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	res, err := r.client.Get(state.getPathGet() + params)
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
func (r *PnPNetworkDeviceClaimResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PnPNetworkDeviceClaim

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
func (r *PnPNetworkDeviceClaimResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PnPNetworkDeviceClaim

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPathDelete())
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *PnPNetworkDeviceClaimResource) ReadCache(ctx context.Context, req resource.ReadRequest, state PnPNetworkDeviceClaim, params string) (cc.Res, error) {
	var err error
	cacheKey := "PnPNetworkDeviceClaim::"

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
	res, err := r.client.Get("/dna/intent/api/v1/pnpNetworkDevices/%v" + params)
	singleRes := res
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache
