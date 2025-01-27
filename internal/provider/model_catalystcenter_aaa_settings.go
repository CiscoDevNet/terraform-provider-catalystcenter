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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AAASettings struct {
	Id                          types.String `tfsdk:"id"`
	SiteId                      types.String `tfsdk:"site_id"`
	NetworkAaaServerType        types.String `tfsdk:"network_aaa_server_type"`
	NetworkAaaProtocol          types.String `tfsdk:"network_aaa_protocol"`
	NetworkAaaPan               types.String `tfsdk:"network_aaa_pan"`
	NetworkAaaPrimaryServerIp   types.String `tfsdk:"network_aaa_primary_server_ip"`
	NetworkAaaSecondaryServerIp types.String `tfsdk:"network_aaa_secondary_server_ip"`
	NetworkAaaSharedSecret      types.String `tfsdk:"network_aaa_shared_secret"`
	ClientAaaServerType         types.String `tfsdk:"client_aaa_server_type"`
	ClientAaaProtocol           types.String `tfsdk:"client_aaa_protocol"`
	ClientAaaPan                types.String `tfsdk:"client_aaa_pan"`
	ClientAaaPrimaryServerIp    types.String `tfsdk:"client_aaa_primary_server_ip"`
	ClientAaaSecondaryServerIp  types.String `tfsdk:"client_aaa_secondary_server_ip"`
	ClientAaaSharedSecret       types.String `tfsdk:"client_aaa_shared_secret"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AAASettings) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/aaaSettings", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AAASettings) toBody(ctx context.Context, state AAASettings) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.NetworkAaaServerType.IsNull() {
		body, _ = sjson.Set(body, "aaaNetwork.serverType", data.NetworkAaaServerType.ValueString())
	}
	if !data.NetworkAaaProtocol.IsNull() {
		body, _ = sjson.Set(body, "aaaNetwork.protocol", data.NetworkAaaProtocol.ValueString())
	}
	if !data.NetworkAaaPan.IsNull() {
		body, _ = sjson.Set(body, "aaaNetwork.pan", data.NetworkAaaPan.ValueString())
	}
	if !data.NetworkAaaPrimaryServerIp.IsNull() {
		body, _ = sjson.Set(body, "aaaNetwork.primaryServerIp", data.NetworkAaaPrimaryServerIp.ValueString())
	}
	if !data.NetworkAaaSecondaryServerIp.IsNull() {
		body, _ = sjson.Set(body, "aaaNetwork.secondaryServerIp", data.NetworkAaaSecondaryServerIp.ValueString())
	}
	if !data.NetworkAaaSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "aaaNetwork.sharedSecret", data.NetworkAaaSharedSecret.ValueString())
	}
	if !data.ClientAaaServerType.IsNull() {
		body, _ = sjson.Set(body, "aaaClient.serverType", data.ClientAaaServerType.ValueString())
	}
	if !data.ClientAaaProtocol.IsNull() {
		body, _ = sjson.Set(body, "aaaClient.protocol", data.ClientAaaProtocol.ValueString())
	}
	if !data.ClientAaaPan.IsNull() {
		body, _ = sjson.Set(body, "aaaClient.pan", data.ClientAaaPan.ValueString())
	}
	if !data.ClientAaaPrimaryServerIp.IsNull() {
		body, _ = sjson.Set(body, "aaaClient.primaryServerIp", data.ClientAaaPrimaryServerIp.ValueString())
	}
	if !data.ClientAaaSecondaryServerIp.IsNull() {
		body, _ = sjson.Set(body, "aaaClient.secondaryServerIp", data.ClientAaaSecondaryServerIp.ValueString())
	}
	if !data.ClientAaaSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "aaaClient.sharedSecret", data.ClientAaaSharedSecret.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AAASettings) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get(""); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.serverType"); value.Exists() {
		data.NetworkAaaServerType = types.StringValue(value.String())
	} else {
		data.NetworkAaaServerType = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.protocol"); value.Exists() {
		data.NetworkAaaProtocol = types.StringValue(value.String())
	} else {
		data.NetworkAaaProtocol = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.pan"); value.Exists() {
		data.NetworkAaaPan = types.StringValue(value.String())
	} else {
		data.NetworkAaaPan = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.primaryServerIp"); value.Exists() {
		data.NetworkAaaPrimaryServerIp = types.StringValue(value.String())
	} else {
		data.NetworkAaaPrimaryServerIp = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.secondaryServerIp"); value.Exists() {
		data.NetworkAaaSecondaryServerIp = types.StringValue(value.String())
	} else {
		data.NetworkAaaSecondaryServerIp = types.StringNull()
	}
	if value := res.Get("response.aaaClient.serverType"); value.Exists() {
		data.ClientAaaServerType = types.StringValue(value.String())
	} else {
		data.ClientAaaServerType = types.StringNull()
	}
	if value := res.Get("response.aaaClient.protocol"); value.Exists() {
		data.ClientAaaProtocol = types.StringValue(value.String())
	} else {
		data.ClientAaaProtocol = types.StringNull()
	}
	if value := res.Get("response.aaaClient.pan"); value.Exists() {
		data.ClientAaaPan = types.StringValue(value.String())
	} else {
		data.ClientAaaPan = types.StringNull()
	}
	if value := res.Get("response.aaaClient.primaryServerIp"); value.Exists() {
		data.ClientAaaPrimaryServerIp = types.StringValue(value.String())
	} else {
		data.ClientAaaPrimaryServerIp = types.StringNull()
	}
	if value := res.Get("response.aaaClient.secondaryServerIp"); value.Exists() {
		data.ClientAaaSecondaryServerIp = types.StringValue(value.String())
	} else {
		data.ClientAaaSecondaryServerIp = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AAASettings) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.aaaNetwork.serverType"); value.Exists() && !data.NetworkAaaServerType.IsNull() {
		data.NetworkAaaServerType = types.StringValue(value.String())
	} else {
		data.NetworkAaaServerType = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.protocol"); value.Exists() && !data.NetworkAaaProtocol.IsNull() {
		data.NetworkAaaProtocol = types.StringValue(value.String())
	} else {
		data.NetworkAaaProtocol = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.pan"); value.Exists() && !data.NetworkAaaPan.IsNull() {
		data.NetworkAaaPan = types.StringValue(value.String())
	} else {
		data.NetworkAaaPan = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.primaryServerIp"); value.Exists() && !data.NetworkAaaPrimaryServerIp.IsNull() {
		data.NetworkAaaPrimaryServerIp = types.StringValue(value.String())
	} else {
		data.NetworkAaaPrimaryServerIp = types.StringNull()
	}
	if value := res.Get("response.aaaNetwork.secondaryServerIp"); value.Exists() && !data.NetworkAaaSecondaryServerIp.IsNull() {
		data.NetworkAaaSecondaryServerIp = types.StringValue(value.String())
	} else {
		data.NetworkAaaSecondaryServerIp = types.StringNull()
	}
	if value := res.Get("response.aaaClient.serverType"); value.Exists() && !data.ClientAaaServerType.IsNull() {
		data.ClientAaaServerType = types.StringValue(value.String())
	} else {
		data.ClientAaaServerType = types.StringNull()
	}
	if value := res.Get("response.aaaClient.protocol"); value.Exists() && !data.ClientAaaProtocol.IsNull() {
		data.ClientAaaProtocol = types.StringValue(value.String())
	} else {
		data.ClientAaaProtocol = types.StringNull()
	}
	if value := res.Get("response.aaaClient.pan"); value.Exists() && !data.ClientAaaPan.IsNull() {
		data.ClientAaaPan = types.StringValue(value.String())
	} else {
		data.ClientAaaPan = types.StringNull()
	}
	if value := res.Get("response.aaaClient.primaryServerIp"); value.Exists() && !data.ClientAaaPrimaryServerIp.IsNull() {
		data.ClientAaaPrimaryServerIp = types.StringValue(value.String())
	} else {
		data.ClientAaaPrimaryServerIp = types.StringNull()
	}
	if value := res.Get("response.aaaClient.secondaryServerIp"); value.Exists() && !data.ClientAaaSecondaryServerIp.IsNull() {
		data.ClientAaaSecondaryServerIp = types.StringValue(value.String())
	} else {
		data.ClientAaaSecondaryServerIp = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AAASettings) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.NetworkAaaServerType.IsNull() {
		return false
	}
	if !data.NetworkAaaProtocol.IsNull() {
		return false
	}
	if !data.NetworkAaaPan.IsNull() {
		return false
	}
	if !data.NetworkAaaPrimaryServerIp.IsNull() {
		return false
	}
	if !data.NetworkAaaSecondaryServerIp.IsNull() {
		return false
	}
	if !data.NetworkAaaSharedSecret.IsNull() {
		return false
	}
	if !data.ClientAaaServerType.IsNull() {
		return false
	}
	if !data.ClientAaaProtocol.IsNull() {
		return false
	}
	if !data.ClientAaaPan.IsNull() {
		return false
	}
	if !data.ClientAaaPrimaryServerIp.IsNull() {
		return false
	}
	if !data.ClientAaaSecondaryServerIp.IsNull() {
		return false
	}
	if !data.ClientAaaSharedSecret.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
