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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type MapArchiveImport struct {
	Id                types.String `tfsdk:"id"`
	SourcePath        types.String `tfsdk:"source_path"`
	AcceptDataLoss    types.Bool   `tfsdk:"accept_data_loss"`
	ArchiveSha256     types.String `tfsdk:"archive_sha256"`
	ImportContextUuid types.String `tfsdk:"import_context_uuid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data MapArchiveImport) getPath() string {
	return "/dna/intent/api/v1/maps/import/start"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data MapArchiveImport) toBody(ctx context.Context, state MapArchiveImport) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.SourcePath.IsNull() {
		body, _ = sjson.Set(body, "sourcePath", data.SourcePath.ValueString())
	}
	if !data.AcceptDataLoss.IsNull() {
		body, _ = sjson.Set(body, "acceptDataLoss", data.AcceptDataLoss.ValueBool())
	}
	if !data.ArchiveSha256.IsNull() {
		body, _ = sjson.Set(body, "archiveSha256", data.ArchiveSha256.ValueString())
	}
	if data.ImportContextUuid.ValueString() != "" && !data.ImportContextUuid.IsNull() {
		body, _ = sjson.Set(body, "importContextUuid", data.ImportContextUuid.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *MapArchiveImport) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("acceptDataLoss"); value.Exists() {
		data.AcceptDataLoss = types.BoolValue(value.Bool())
	} else {
		data.AcceptDataLoss = types.BoolNull()
	}
	if value := res.Get("archiveSha256"); value.Exists() {
		data.ArchiveSha256 = types.StringValue(value.String())
	} else {
		data.ArchiveSha256 = types.StringNull()
	}
	if value := res.Get("importContextUuid"); value.Exists() {
		data.ImportContextUuid = types.StringValue(value.String())
	} else {
		data.ImportContextUuid = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *MapArchiveImport) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("acceptDataLoss"); value.Exists() && !data.AcceptDataLoss.IsNull() {
		data.AcceptDataLoss = types.BoolValue(value.Bool())
	} else {
		data.AcceptDataLoss = types.BoolNull()
	}
	if value := res.Get("archiveSha256"); value.Exists() && !data.ArchiveSha256.IsNull() {
		data.ArchiveSha256 = types.StringValue(value.String())
	} else {
		data.ArchiveSha256 = types.StringNull()
	}
	if value := res.Get("importContextUuid"); value.Exists() && !data.ImportContextUuid.IsNull() {
		data.ImportContextUuid = types.StringValue(value.String())
	} else {
		data.ImportContextUuid = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *MapArchiveImport) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.SourcePath.IsNull() {
		return false
	}
	if !data.AcceptDataLoss.IsNull() {
		return false
	}
	if !data.ArchiveSha256.IsNull() {
		return false
	}
	if !data.ImportContextUuid.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
