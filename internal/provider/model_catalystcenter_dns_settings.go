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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type DNSSettings struct {
	Id         types.String `tfsdk:"id"`
	SiteId     types.String `tfsdk:"site_id"`
	DomainName types.String `tfsdk:"domain_name"`
	DnsServers types.Set    `tfsdk:"dns_servers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data DNSSettings) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/dnsSettings", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data DNSSettings) toBody(ctx context.Context, state DNSSettings) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DomainName.IsNull() {
		body, _ = sjson.Set(body, "dns.domainName", data.DomainName.ValueString())
	}
	if !data.DnsServers.IsNull() {
		var values []string
		data.DnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dns.dnsServers", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *DNSSettings) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get(""); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.dns.domainName"); value.Exists() {
		data.DomainName = types.StringValue(value.String())
	} else {
		data.DomainName = types.StringNull()
	}
	if value := res.Get("response.dns.dnsServers"); value.Exists() && len(value.Array()) > 0 {
		data.DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.DnsServers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *DNSSettings) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.dns.domainName"); value.Exists() && !data.DomainName.IsNull() {
		data.DomainName = types.StringValue(value.String())
	} else {
		data.DomainName = types.StringNull()
	}
	if value := res.Get("response.dns.dnsServers"); value.Exists() && !data.DnsServers.IsNull() {
		data.DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.DnsServers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DNSSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.DomainName.IsUnknown() {
		if value := res.Get("response.dns.domainName"); value.Exists() && !data.DomainName.IsNull() {
			data.DomainName = types.StringValue(value.String())
		} else {
			data.DomainName = types.StringNull()
		}
	}
	if value := res.Get("response.dns.dnsServers"); value.Exists() && !data.DnsServers.IsNull() {
		data.DnsServers = helpers.GetStringSet(value.Array())
	} else {
		data.DnsServers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *DNSSettings) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DomainName.IsNull() {
		return false
	}
	if !data.DnsServers.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
