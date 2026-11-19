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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
var _ resource.Resource = &MapArchiveImportResource{}

func NewMapArchiveImportResource() resource.Resource {
	return &MapArchiveImportResource{}
}

type MapArchiveImportResource struct {
	client                *cc.Client
	AllowExistingOnCreate bool
	cache                 *ThreadSafeCache
}

func (r *MapArchiveImportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_map_archive_import"
}

func (r *MapArchiveImportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource imports a Catalyst Center map archive (tar.gz) containing xmlDir/MapsImportExport.xml plus floor images under /images. It orchestrates start import, pre-import validation, and perform import. <p/> The archive must be Cisco's XML interchange format (for example an Ekahau Cloud Project Export via API). Raw Ekahau .esx files are not accepted by the Intent API. Sites (area, building, floor) named in the archive must already exist. Perform fully replaces existing map data (images, obstacles, planned APs) for those sites. <p/> Re-apply is a no-op while the archive file is unchanged (SHA-256). Changing the file or source_path replaces the resource and imports again. Destroy removes the resource from state only; Catalyst Center map data is left in place.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"source_path": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Local path to a Catalyst Center map archive tar.gz (xmlDir/MapsImportExport.xml plus images/). Ekahau .esx files are not supported.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"accept_data_loss": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Must be true. Acknowledges that perform import fully replaces existing map data for the site(s) defined in the archive.").String,
				Required:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"archive_sha256": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SHA-256 of the archive file. Replacement is triggered when this hash changes.").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"import_context_uuid": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Import context UUID returned by start import. Retained after a successful import for troubleshooting.").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *MapArchiveImportResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
	r.AllowExistingOnCreate = req.ProviderData.(*CcProviderData).AllowExistingOnCreate
	r.cache = req.ProviderData.(*CcProviderData).Cache
}

// End of section. //template:end model

var _ resource.ResourceWithModifyPlan = &MapArchiveImportResource{}

// Custom create: map archive import is a multipart upload plus a two-step
// async workflow (start → poll validation → perform → poll completion). The
// generator cannot express this; floor_image is the closest upload precedent.
func (r *MapArchiveImportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan MapArchiveImport

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	if !plan.AcceptDataLoss.ValueBool() {
		resp.Diagnostics.AddError("Invalid Attribute",
			"`accept_data_loss` must be true. Perform import fully replaces existing map data for the site(s) in the archive.")
		return
	}

	sourcePath := plan.SourcePath.ValueString()
	if strings.HasSuffix(strings.ToLower(sourcePath), ".esx") {
		resp.Diagnostics.AddError("Invalid Attribute",
			"source_path points to an Ekahau .esx file. The Intent API only accepts a Catalyst Center map archive tar.gz (xmlDir/MapsImportExport.xml plus images/). Export the project from Ekahau Cloud as Catalyst Center import files first.")
		return
	}

	hash, err := hashFileSHA256(sourcePath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to read map archive: %s", err))
		return
	}
	plan.ArchiveSha256 = types.StringValue(hash)
	plan.Id = types.StringValue(hash)

	importUUID, err := r.startMapArchiveImport(ctx, sourcePath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	plan.ImportContextUuid = types.StringValue(importUUID)
	tflog.Debug(ctx, fmt.Sprintf("Map archive start import context UUID: %s", importUUID))

	status, err := r.waitForMapImportStatus(ctx, importUUID, []string{"DEFINED", "COMPLETED"})
	if err != nil {
		r.cancelMapArchiveImport(ctx, importUUID)
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	if msgs := mapImportAuditMessages(status, "errors"); len(msgs) > 0 {
		r.cancelMapArchiveImport(ctx, importUUID)
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Map archive failed pre-import validation:\n%s", strings.Join(msgs, "\n")))
		return
	}

	if mapImportStatus(status) != "COMPLETED" {
		if err := r.performMapArchiveImport(ctx, importUUID); err != nil {
			r.cancelMapArchiveImport(ctx, importUUID)
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
		status, err = r.waitForMapImportStatus(ctx, importUUID, []string{"COMPLETED"})
		if err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
		if msgs := mapImportAuditMessages(status, "errors"); len(msgs) > 0 {
			resp.Diagnostics.AddError("Client Error",
				fmt.Sprintf("Map archive import completed with errors:\n%s", strings.Join(msgs, "\n")))
			return
		}
	}

	if msgs := mapImportAuditMessages(status, "warnings"); len(msgs) > 0 {
		resp.Diagnostics.AddWarning("Map archive import warnings", strings.Join(msgs, "\n"))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *MapArchiveImportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state MapArchiveImport

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
func (r *MapArchiveImportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var plan, state MapArchiveImport

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
func (r *MapArchiveImportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	applyProviderMeta(r.client, ctx, req.ProviderMeta)
	var state MapArchiveImport

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
func (r *MapArchiveImportResource) ReadCache(ctx context.Context, req resource.ReadRequest, state MapArchiveImport, params string) (cc.Res, error) {
	var err error
	cacheKey := "MapArchiveImport::"

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

func (r *MapArchiveImportResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan MapArchiveImport
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.SourcePath.IsNull() || plan.SourcePath.IsUnknown() {
		return
	}

	hash, err := hashFileSHA256(plan.SourcePath.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("source_path"), "Cannot read map archive", err.Error())
		return
	}
	plan.ArchiveSha256 = types.StringValue(hash)

	if !req.State.Raw.IsNull() {
		var state MapArchiveImport
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !state.ArchiveSha256.IsNull() && state.ArchiveSha256.ValueString() != hash {
			resp.RequiresReplace = append(resp.RequiresReplace, path.Root("source_path"))
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (r *MapArchiveImportResource) startMapArchiveImport(ctx context.Context, sourcePath string) (string, error) {
	body, contentType, err := pipeMapArchiveUpload(ctx, sourcePath)
	if err != nil {
		return "", fmt.Errorf("failed to open map archive: %w", err)
	}

	if err := r.client.Authenticate(); err != nil {
		return "", fmt.Errorf("failed to authenticate: %w", err)
	}

	// Use HttpClient directly: go-catalystcenter Do() always adds
	// Content-Type: application/json, which this multipart endpoint rejects.
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.client.Url+"/dna/intent/api/v1/maps/import/start", body)
	if err != nil {
		return "", fmt.Errorf("failed to build start import request: %w", err)
	}
	httpReq.Header.Set("Content-Type", contentType)
	httpReq.Header.Set("X-Auth-Token", r.client.Token)
	if r.client.UserAgent != "" {
		httpReq.Header.Set("User-Agent", r.client.UserAgent)
	}

	httpRes, err := r.client.HttpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to start map archive import (POST /maps/import/start): %w", err)
	}
	defer httpRes.Body.Close()

	respBody, _ := io.ReadAll(httpRes.Body)
	res := cc.Res(gjson.ParseBytes(respBody))
	if httpRes.StatusCode < 200 || httpRes.StatusCode > 299 {
		msg := strings.TrimSpace(string(respBody))
		if msg == "" {
			msg = httpRes.Status
		}
		return "", fmt.Errorf("failed to start map archive import (POST /maps/import/start): HTTP %d: %s", httpRes.StatusCode, msg)
	}

	uuid := parseImportContextUUID(res)
	if uuid == "" {
		return "", fmt.Errorf("start import did not return an import context UUID; response: %s", res.Raw)
	}
	return uuid, nil
}

// pipeMapArchiveUpload builds a multipart/form-data body with the archive field
// required by POST /dna/intent/api/v1/maps/import/start.
func pipeMapArchiveUpload(ctx context.Context, sourcePath string) (io.Reader, string, error) {
	source, err := os.Open(sourcePath)
	if err != nil {
		return nil, "", err
	}

	readPipe, writePipe := io.Pipe()
	mw := multipart.NewWriter(writePipe)
	contentType := mw.FormDataContentType()

	go func() {
		defer source.Close()
		partWriter, err := mw.CreateFormFile("archive", filepath.Base(sourcePath))
		if err != nil {
			writePipe.CloseWithError(err)
			return
		}
		if _, err = io.Copy(partWriter, source); err != nil {
			writePipe.CloseWithError(err)
			return
		}
		if err := mw.Close(); err != nil {
			writePipe.CloseWithError(err)
			return
		}
		writePipe.Close()
		tflog.Debug(ctx, "Map archive multipart body closed")
	}()

	return readPipe, contentType, nil
}

func (r *MapArchiveImportResource) performMapArchiveImport(ctx context.Context, importUUID string) error {
	path := fmt.Sprintf("/dna/intent/api/v1/maps/import/%s/perform", url.PathEscape(importUUID))
	res, err := r.client.Post(path, "", cc.NoWait, cc.UseMutex)
	if err != nil {
		return fmt.Errorf("failed to perform map archive import: %s, %s", err, res.String())
	}
	return nil
}

func (r *MapArchiveImportResource) cancelMapArchiveImport(ctx context.Context, importUUID string) {
	if importUUID == "" {
		return
	}
	path := fmt.Sprintf("/dna/intent/api/v1/maps/import/%s", url.PathEscape(importUUID))
	if _, err := r.client.Delete(path, cc.NoWait); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to cancel map archive import %s: %s", importUUID, err))
	}
}

func (r *MapArchiveImportResource) waitForMapImportStatus(ctx context.Context, importUUID string, want []string) (cc.Res, error) {
	wanted := make(map[string]struct{}, len(want))
	for _, s := range want {
		wanted[s] = struct{}{}
	}

	maxWait := time.Duration(r.client.DefaultMaxAsyncWaitTime) * time.Second
	if maxWait <= 0 {
		maxWait = 10 * time.Minute
	}
	deadline := time.Now().Add(maxWait)
	statusPath := fmt.Sprintf("/dna/intent/api/v1/maps/import/%s/status", url.PathEscape(importUUID))

	var res cc.Res
	for {
		var err error
		res, err = r.client.Get(statusPath)
		if err != nil {
			return res, fmt.Errorf("failed to read map archive import status: %s, %s", err, res.String())
		}

		status := mapImportStatus(res)
		tflog.Debug(ctx, fmt.Sprintf("Map archive import %s status: %s", importUUID, status))
		// Observed statuses: VALIDATING, DEFINED, UNDERWAY, COMPLETED, CANCELLED.

		if _, ok := wanted[status]; ok {
			return res, nil
		}
		if status == "CANCELLED" {
			msgs := mapImportAuditMessages(res, "errors")
			if len(msgs) == 0 {
				msgs = []string{"import was cancelled"}
			}
			return res, fmt.Errorf("map archive import cancelled: %s", strings.Join(msgs, "\n"))
		}

		if time.Now().After(deadline) {
			return res, fmt.Errorf("timed out waiting for map archive import status %s (last status %q)", strings.Join(want, "/"), status)
		}

		select {
		case <-ctx.Done():
			return res, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func mapImportStatus(res cc.Res) string {
	for _, p := range []string{"status", "response.status"} {
		if v := res.Get(p).String(); v != "" {
			return v
		}
	}
	return ""
}

// mapImportAuditMessages collects errors or warnings from the whole auditLog
// tree; Catalyst Center reports them per site, building and floor under
// nested children, not only at the root.
func mapImportAuditMessages(res cc.Res, kind string) []string {
	root := res.Get("auditLog")
	if !root.Exists() {
		root = res.Get("response.auditLog")
	}

	var msgs []string
	var walk func(node cc.Res)
	walk = func(node cc.Res) {
		entity := strings.TrimSpace(node.Get("entityType").String() + " " + node.Get("entityName").String())
		node.Get(kind).ForEach(func(_, value cc.Res) bool {
			m := value.Get("message").String()
			if m == "" {
				m = value.String()
			}
			if m == "" || m == "{}" {
				return true
			}
			if node.Get("entityName").String() != "" {
				m = fmt.Sprintf("%s: %s", entity, m)
			}
			msgs = append(msgs, m)
			return true
		})
		node.Get("children").ForEach(func(_, child cc.Res) bool {
			walk(child)
			return true
		})
	}
	walk(root)
	return msgs
}

var uuidRE = regexp.MustCompile(`(?i)[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)

func parseImportContextUUID(res cc.Res) string {
	for _, p := range []string{
		"importContextUuid",
		"response.importContextUuid",
		"response",
		"uuid",
	} {
		if v := res.Get(p).String(); uuidRE.MatchString(v) {
			return uuidRE.FindString(v)
		}
	}
	if v := uuidRE.FindString(res.Raw); v != "" {
		return v
	}
	return ""
}

func hashFileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
