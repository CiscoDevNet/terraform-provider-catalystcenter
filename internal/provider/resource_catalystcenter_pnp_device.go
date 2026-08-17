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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PnPDeviceResource{}
var _ resource.ResourceWithImportState = &PnPDeviceResource{}

func NewPnPDeviceResource() resource.Resource {
	return &PnPDeviceResource{}
}

type PnPDeviceResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *PnPDeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pnp_device"
}

func (r *PnPDeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a PnP Device.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device serial number").String,
				Required:            true,
			},
			"stack": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device is a stacked switch").String,
				Optional:            true,
			},
			"pid": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device product ID").String,
				Required:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device hostname").String,
				Required:            true,
			},
		},
	}
}

func (r *PnPDeviceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

func (r *PnPDeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PnPDevice

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, PnPDevice{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		errorCode := res.Get("response.errorCode").String()
		// if the error code is NCOB01019, the device already exists in PnP
		// Recover its id via GET, then PUT the desired config so attributes
		// such as hostname/pid are reconciled with the user's plan.
		if errorCode == "NCOB01019" {
			// Retrieve Id from the GET response
			params := ""
			params += "?serialNumber=" + url.QueryEscape(plan.SerialNumber.ValueString())
			res, err = r.client.Get(plan.getPath() + params)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve existing PnP device (GET), got error: %s, %s", err, res.String()))
				return
			}
			// Set ID from GET response
			plan.Id = types.StringValue(res.Get("0.id").String())
			if plan.Id.ValueString() == "" {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("NCOB01019 returned but no device found by serial number: %s", plan.SerialNumber.ValueString()))
				return
			}

			// Reconcile existing device with desired config (hostname, pid, ...)
			body = plan.toBody(ctx, plan)
			res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Existing PnP device found but failed to update (PUT), got error: %s, %s", err, res.String()))
				return
			}

			tflog.Debug(ctx, fmt.Sprintf("%s: Create (reconcile existing) finished successfully", plan.Id.ValueString()))

			// Save to state
			diags := resp.State.Set(ctx, &plan)
			resp.Diagnostics.Append(diags...)
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			return
		}
	}
	plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Custom read: hardened against the PnP empty-serial device-adoption bug.
// The generated template queries by the state serial number and blindly
// accepts element 0 of the response. When the state serial becomes empty
// (e.g. a prior empty HTTP 200 response nulled the identity fields), that
// produces a "?serialNumber=" query which Catalyst Center answers with an
// unrelated PnP device. The provider would then adopt that device (including
// its id) into this resource. We guard against that by refusing to query with
// an empty serial and by requiring the returned primary serial (or a stack
// member serial) to match, case-insensitively, before writing state.
func (r *PnPDeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PnPDevice

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	// Never issue an empty "?serialNumber=" query. Catalyst Center returns an
	// unrelated device for an empty serial, which would be silently adopted
	// into this resource. A missing serial means the resource is no longer
	// resolvable, so remove it from state and let Terraform plan a re-create.
	serial := state.SerialNumber.ValueString()
	if serial == "" {
		tflog.Warn(ctx, fmt.Sprintf("%s: state serial number is empty; removing resource from state instead of issuing an empty serialNumber query", state.Id.ValueString()))
		resp.State.RemoveResource(ctx)
		return
	}

	params := ""
	params += "?serialNumber=" + url.QueryEscape(serial)
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// Require the requested serial to match before writing state. Catalyst
	// Center may answer with an empty array (HTTP 200) when the device is
	// temporarily absent, or an unrelated device for a malformed query. Accept
	// a match against the returned primary serial or, for stacks, any stack
	// member serial. Comparison is case-insensitive to tolerate serial casing
	// drift. Otherwise remove the resource rather than adopting another device
	// or nulling this resource's identity fields.
	returnedSerial := res.Get("0.deviceInfo.serialNumber").String()
	matched := strings.EqualFold(returnedSerial, serial)
	if !matched {
		res.Get("0.deviceInfo.stackInfo.stackMemberList.#.serialNumber").ForEach(func(_, member gjson.Result) bool {
			if strings.EqualFold(member.String(), serial) {
				matched = true
				return false
			}
			return true
		})
	}
	if !matched {
		tflog.Warn(ctx, fmt.Sprintf("%s: no PnP device matching serial %q (API returned primary serial %q); removing resource from state", state.Id.ValueString(), serial, returnedSerial))
		resp.State.RemoveResource(ctx)
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

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *PnPDeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PnPDevice

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

func (r *PnPDeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PnPDevice

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") && !strings.Contains(err.Error(), "StatusCode 400") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *PnPDeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <serial_number>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial_number"), idParts[0])...)
}

// End of section. //template:end import
