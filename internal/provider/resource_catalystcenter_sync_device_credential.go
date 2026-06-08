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
	"time"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SyncDeviceCredentialResource{}

func NewSyncDeviceCredentialResource() resource.Resource {
	return &SyncDeviceCredentialResource{}
}

type SyncDeviceCredentialResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *SyncDeviceCredentialResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sync_device_credential"
}

func (r *SyncDeviceCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource applies (syncs) a device credential that is already assigned to a site to the site's network devices. It is the equivalent of the Apply action in Design > Network Settings > Credentials > Manage Credentials. The credential must already be assigned to the site (see `catalystcenter_assign_credentials`). Applying at a site also affects child sites that inherit the credential. The endpoint accepts a single credential ID per call (CLI, SNMPv2c Read, SNMPv2c Write or SNMPv3). <p/> This resource only triggers the apply operation on create and waits for the task to complete. Changing any attribute forces a new apply, and when the resource is destroyed or refreshed no actions are performed on Catalyst Center.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the site whose assigned credential should be applied to its devices.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"device_credential_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the credential to apply. Must be a CLI, SNMPv2c Read, SNMPv2c Write or SNMPv3 credential ID.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"configure_device": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only applicable when `device_credential_id` is a CLI credential; ignored by Catalyst Center for SNMPv2c/SNMPv3 credential IDs. If `true`, the applied CLI credential is configured on locally-authenticated devices and authentication is performed before the device is managed in inventory; AAA-authenticated devices are only managed, not reconfigured. If `false`, devices are updated in inventory with the applied credential without any authentication. When omitted, the parameter is not sent and the Catalyst Center default applies.").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *SyncDeviceCredentialResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SyncDeviceCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SyncDeviceCredential

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, SyncDeviceCredential{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		retryErrorCodes := []string{"NCND00010"}
		errorCode := res.Get("response.errorCode").String()

		shouldRetry := false
		for _, code := range retryErrorCodes {
			if errorCode == code {
				shouldRetry = true
				break
			}
		}

		if shouldRetry {
			maxWaitTime := time.Duration(r.client.DefaultMaxAsyncWaitTime) * time.Second
			startTime := time.Now()
			retryInterval := 15 * time.Second

			for shouldRetry && time.Since(startTime) < maxWaitTime {
				tflog.Warn(ctx, fmt.Sprintf("%s: Error code %s encountered, waiting %v before retry (elapsed: %v, max: %v)",
					plan.Id.ValueString(), errorCode, retryInterval, time.Since(startTime), maxWaitTime))
				time.Sleep(retryInterval)
				res, err = r.client.Post(plan.getPath()+params, body)

				if err == nil {
					shouldRetry = false
				} else {
					errorCode = res.Get("response.errorCode").String()
					shouldRetry = false
					for _, code := range retryErrorCodes {
						if errorCode == code {
							shouldRetry = true
							break
						}
					}
				}
			}
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	plan.Id = types.StringValue(fmt.Sprint(plan.SiteId.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *SyncDeviceCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SyncDeviceCredential

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
func (r *SyncDeviceCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SyncDeviceCredential

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
func (r *SyncDeviceCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SyncDeviceCredential

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
func (r *SyncDeviceCredentialResource) ReadCache(ctx context.Context, req resource.ReadRequest, state SyncDeviceCredential, params string) (cc.Res, error) {
	var err error
	cacheKey := "SyncDeviceCredential::"

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
