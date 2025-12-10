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
	"encoding/json"
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &FloorResource{}
var _ resource.ResourceWithImportState = &FloorResource{}

func NewFloorResource() resource.Resource {
	return &FloorResource{}
}

type FloorResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *FloorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_floor"
}

func (r *FloorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Floor.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the parent building").String,
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Floor name").String,
				Required:            true,
			},
			"floor_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Floor number").String,
				Required:            true,
			},
			"rf_model": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The RF model").AddStringEnumDescription("Cubes And Walled Offices", "Drywall Office Only", "Indoor High Ceiling", "Outdoor Open Space", "Free Space").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Cubes And Walled Offices", "Drywall Office Only", "Indoor High Ceiling", "Outdoor Open Space", "Free Space"),
				},
			},
			"width": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Width").String,
				Required:            true,
			},
			"length": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Length").String,
				Required:            true,
			},
			"height": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Height").String,
				Required:            true,
			},
			"units_of_measure": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The unit of measurement").AddStringEnumDescription("feet", "meters").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("feet", "meters"),
				},
			},
		},
	}
}

func (r *FloorResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *FloorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Floor
	cacheKey := "Floor"
	r.cache.DeletePattern(cacheKey)

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Floor{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		if r.AllowExistingOnCreate {
			tflog.Info(ctx, fmt.Sprintf("Failed to configure object (%s), got error: %s, %s. allow_existing_on_create is true, beginning update", "POST", err, res.String()))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "POST", err, res.String()))
			return
		}
	}
	plan.Id = types.StringValue(res.Get("response.data").String())
	if !r.AllowExistingOnCreate {
		tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))
	} else {
		params = "?name=" + url.QueryEscape(plan.Name.ValueString())
		res, err = r.client.Get("/dna/intent/api/v1/sites" + params)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		plan.Id = types.StringValue(res.Get("response.#(parentId==\"" + plan.ParentId.ValueString() + "\").id").String())
		body = plan.toBody(ctx, Floor{Id: plan.Id})
		res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (%s), got error: %s, %s", "PUT", err, res.String()))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Fallback to update existing resource finished successfully", plan.Id.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *FloorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Floor

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(state.Id.ValueString())
	params += "?_unitsOfMeasure=" + url.QueryEscape(state.UnitsOfMeasure.ValueString())
	res, err := r.ReadCache(ctx, req, state, params)
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
func (r *FloorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Floor
	cacheKey := "Floor"
	r.cache.DeletePattern(cacheKey)

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
func (r *FloorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Floor
	cacheKey := "Floor"
	r.cache.DeletePattern(cacheKey)

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
func (r *FloorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <units_of_measure>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("units_of_measure"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin readcache
func (r *FloorResource) ReadCache(ctx context.Context, req resource.ReadRequest, state Floor, params string) (cc.Res, error) {
	var err error
	cacheKey := "Floor"

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
		filteredValue := cachedValue.(cc.Res).Get("response.#(id==\"" + state.Id.ValueString() + "\")")
		jsonBytes, _ := json.Marshal(map[string]interface{}{"response": filteredValue.Value()})
		wrappedRes := gjson.ParseBytes(jsonBytes)
		return wrappedRes, nil
	}
	res, err := r.client.Get("/dna/intent/api/v1/sites?type=floor" + strings.Replace(cacheSuffix, "?", "&", 1))
	foundRes := res.Get("response.#(id==\"" + state.Id.ValueString() + "\")")
	jsonBytes, _ := json.Marshal(map[string]interface{}{"response": foundRes.Value()})
	singleRes := gjson.ParseBytes(jsonBytes)
	if err == nil {
		tflog.Debug(ctx, fmt.Sprintf("set cache for %s", cacheKey))
		r.cache.Set(cacheKey, res)
	}
	return singleRes, err
}

// End of section. //template:end readcache
