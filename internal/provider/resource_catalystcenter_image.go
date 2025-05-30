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
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
var _ resource.Resource = &ImageResource{}

func NewImageResource() resource.Resource {
	return &ImageResource{}
}

type ImageResource struct {
	client *cc.Client
}

func (r *ImageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image"
}

func (r *ImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can fetch a software image from a local file and upload it on the Catalyst Center. It can be further used by `catalystcenter_image_distribution`, `catalystcenter_image_activation`, etc.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"third_party_application_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Third party application type").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"family": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Third party image family").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"source_path": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Local path where the software image file resides. Supported file extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz, qcow2.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("File name that uniquely identifies the software image. It should not contain any path. Usually this can be specified as `basename(source_path)`").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"third_party_vendor": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Third party Vendor").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"is_third_party": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the software image is from a third party.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *ImageResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

// End of section. //template:end model

func (r *ImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Image

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	params := ""
	params += "?isThirdParty=" + plan.IsThirdParty.String()
	params += "&thirdPartyImageFamily=" + plan.Family.ValueString()
	params += "&thirdPartyApplicationType=" + plan.ThirdPartyApplicationType.ValueString()
	params += "&thirdPartyVendor=" + plan.ThirdPartyVendor.ValueString()

	// the body of POST will be a stream (io.Reader), as body might be larger than available memory
	body, contentType, err := pipeMultiPartUpload(ctx, plan.SourcePath.ValueString(), plan.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to upload local file: %v", err))
		return
	}

	client := r.clientWithHttpTimeout(15 * time.Minute)
	request := client.NewReq("POST", plan.getPath()+params, body, func(r *cc.Req) { r.MaxAsyncWaitTime = 600 })

	// Set the content type header to multipart/form-data
	request.HttpReq.Header.Set("Content-Type", contentType)

	// No logging. The tflog can only log from memory, but the body might be larger than available memory.
	request.LogPayload = false

	err = client.Authenticate()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s", err))
		return
	}

	res, err := client.Do(request)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	params = ""
	params += "?name=" + plan.Name.ValueString()
	res, err = r.client.Get("/dna/intent/api/v1/image/importation" + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.#(name==\"" + plan.Name.ValueString() + "\").imageUuid").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// pipeMultiPartUpload opens a file located at the sourcePath, formats the contents as a MIME multipart/form-data
// body of a HTTP POST request (so, a file upload) and returns them as a stream.
// The second return is the corresponding Content-Type header value, and the third an error.
//
// The basename is how the file would be identied internally within the upload.
func pipeMultiPartUpload(ctx context.Context, sourcePath, basename string) (io.Reader, string, error) {
	const dummyContentType = "application/octet-stream"

	source, err := os.Open(sourcePath)
	if err != nil {
		return nil, dummyContentType, err
	}

	tflog.Info(ctx, "swim file opened", map[string]interface{}{"sourcePath": sourcePath})

	readPipe, writePipe := io.Pipe()

	mw := multipart.NewWriter(writePipe)

	// data is copied in parallel, as soon as remote accepts it
	go func() {
		defer source.Close()

		// Create a form file writer for the file field
		partWriter, err := mw.CreateFormFile("file", basename)
		if err != nil {
			panic(err)
		}

		tflog.Info(ctx, "swim mw created")

		// Copy the file data to the form file writer
		var count int64
		if count, err = io.Copy(partWriter, source); err != nil {
			panic(err)
		}

		tflog.Info(ctx, "swim copied through pipe", map[string]interface{}{"countBytes": count})

		// Close the multipart writer to get the terminating boundary

		if err := mw.Close(); err != nil {
			return
		}

		if err := writePipe.Close(); err != nil {
			return
		}

		tflog.Info(ctx, "swim: pipe closed for writing")
	}()

	return readPipe, mw.FormDataContentType(), nil
}

// clientWithHttpTimeout returns a shallow copy of cc.Client with a modified HttpClient.Timeout value.
// The HttpClient.Timeout cannot be overriden by Request.Context() timeout, always the lower of these applies.
// The intent is to increase Timeout without impacting other requests on the client.
func (r *ImageResource) clientWithHttpTimeout(timeout time.Duration) cc.Client {
	res := *r.client
	res.HttpClient = &http.Client{}
	*res.HttpClient = *r.client.HttpClient
	res.HttpClient.Timeout = timeout

	return res
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ImageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Image

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "?imageUuid=" + url.QueryEscape(state.Id.ValueString())
	res, err := r.client.Get("/dna/intent/api/v1/image/importation" + params)
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
func (r *ImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Image

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
func (r *ImageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Image

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
// End of section. //template:end import
