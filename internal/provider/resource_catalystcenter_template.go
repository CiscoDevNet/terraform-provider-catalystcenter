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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &TemplateResource{}
var _ resource.ResourceWithImportState = &TemplateResource{}

func NewTemplateResource() resource.Resource {
	return &TemplateResource{}
}

type TemplateResource struct {
	client *cc.Client
}

func (r *TemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_template"
}

func (r *TemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Template.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the project").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the template").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"device_types": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of device types").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"product_family": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Product family").String,
							Required:            true,
						},
						"product_series": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Product series").String,
							Optional:            true,
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Product type").String,
							Optional:            true,
						},
					},
				},
			},
			"language": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Language of the template").AddStringEnumDescription("JINJA", "VELOCITY").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("JINJA", "VELOCITY"),
				},
			},
			"software_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Software type").String,
				Required:            true,
			},
			"software_variant": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Software variant").String,
				Optional:            true,
			},
			"software_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Software version").String,
				Optional:            true,
			},
			"template_content": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Template content").String,
				Optional:            true,
			},
			"template_params": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of template parameters").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"binding": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bind to source").String,
							Optional:            true,
						},
						"data_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Datatype of template parameter").AddStringEnumDescription("STRING", "INTEGER", "IPADDRESS", "MACADDRESS", "SECTIONDIVIDER").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("STRING", "INTEGER", "IPADDRESS", "MACADDRESS", "SECTIONDIVIDER"),
							},
						},
						"default_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default value of template parameter").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of template parameter").String,
							Optional:            true,
						},
						"display_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Display name of template parameter").String,
							Optional:            true,
						},
						"instruction_text": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Instruction text").String,
							Optional:            true,
						},
						"not_param": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Is it not a variable").String,
							Optional:            true,
						},
						"param_array": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Is it an array").String,
							Optional:            true,
						},
						"parameter_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the template parameter").String,
							Optional:            true,
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of ranges").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"max_value": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Max value of range").String,
										Optional:            true,
									},
									"min_value": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Min value of range").String,
										Optional:            true,
									},
								},
							},
						},
						"required": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Is parameter required").String,
							Optional:            true,
						},
						"default_selected_values": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default selection values").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"selection_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of selection").AddStringEnumDescription("SINGLE_SELECT", "MULTI_SELECT").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("SINGLE_SELECT", "MULTI_SELECT"),
							},
						},
						"selection_values": schema.MapAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Selection values").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
			"composite": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Is it composite template").String,
				Optional:            true,
			},
			"containing_templates": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Containing templates for composite template").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the template").String,
							Required:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the template").String,
							Required:            true,
						},
						"project_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Project name").String,
							Required:            true,
						},
						"language": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Language of the template").AddStringEnumDescription("JINJA", "VELOCITY").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("JINJA", "VELOCITY"),
							},
						},
						"composite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Is it composite template").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *TemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (r *TemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Template

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Template{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	data := res.Get("response.data").String()
	// In case of validation errors/warnings the template ID is encoded in data
	if gjson.Get(data, "templateId").Exists() {
		plan.Id = types.StringValue(gjson.Get(data, "templateId").String())
		// Check for validation errors
		gjson.Get(data, "templateErrors").ForEach(func(k, v gjson.Result) bool {
			resp.Diagnostics.AddWarning("Template Validation Warning", fmt.Sprintf("Template Validation Warning: line number: %s, type: %s, message: %s", v.Get("lineNumber").String(), v.Get("type").String(), v.Get("message").String()))
			return true
		})
	} else {
		plan.Id = types.StringValue(res.Get("response.data").String())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *TemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Template

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(state.Id.ValueString())
	params += "?unCommitted=true"
	res, err := r.client.Get("/dna/intent/api/v1/template-programmer/template" + params)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406")) {
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
func (r *TemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Template

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
	res, err := r.client.Put("/dna/intent/api/v1/template-programmer/template"+params, body)
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
func (r *TemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Template

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPathDelete() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *TemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <project_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

// End of section. //template:end import
