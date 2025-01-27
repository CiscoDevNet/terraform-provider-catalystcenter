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
type TelemetrySettings struct {
	Id                              types.String `tfsdk:"id"`
	SiteId                          types.String `tfsdk:"site_id"`
	EnableWiredDataCollection       types.Bool   `tfsdk:"enable_wired_data_collection"`
	EnableWirelessTelemetry         types.Bool   `tfsdk:"enable_wireless_telemetry"`
	UseBuiltinTrapServer            types.Bool   `tfsdk:"use_builtin_trap_server"`
	ExternalTrapServers             types.Set    `tfsdk:"external_trap_servers"`
	UseBuiltinSyslogServer          types.Bool   `tfsdk:"use_builtin_syslog_server"`
	ExternalSyslogServers           types.Set    `tfsdk:"external_syslog_servers"`
	NetflowCollector                types.String `tfsdk:"netflow_collector"`
	NetflowCollectorIpAddress       types.String `tfsdk:"netflow_collector_ip_address"`
	NetflowCollectorPort            types.Int64  `tfsdk:"netflow_collector_port"`
	EnableNetflowCollectorOnDevices types.Bool   `tfsdk:"enable_netflow_collector_on_devices"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TelemetrySettings) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/telemetrySettings", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TelemetrySettings) toBody(ctx context.Context, state TelemetrySettings) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.EnableWiredDataCollection.IsNull() {
		body, _ = sjson.Set(body, "wiredDataCollection.enableWiredDataCollection", data.EnableWiredDataCollection.ValueBool())
	}
	if !data.EnableWirelessTelemetry.IsNull() {
		body, _ = sjson.Set(body, "wirelessTelemetry.enableWirelessTelemetry", data.EnableWirelessTelemetry.ValueBool())
	}
	if !data.UseBuiltinTrapServer.IsNull() {
		body, _ = sjson.Set(body, "snmpTraps.useBuiltinTrapServer", data.UseBuiltinTrapServer.ValueBool())
	}
	if !data.ExternalTrapServers.IsNull() {
		var values []string
		data.ExternalTrapServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "snmpTraps.externalTrapServers", values)
	}
	if !data.UseBuiltinSyslogServer.IsNull() {
		body, _ = sjson.Set(body, "syslogs.useBuiltinSyslogServer", data.UseBuiltinSyslogServer.ValueBool())
	}
	if !data.ExternalSyslogServers.IsNull() {
		var values []string
		data.ExternalSyslogServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "syslogs.externalSyslogServers", values)
	}
	if !data.NetflowCollector.IsNull() {
		body, _ = sjson.Set(body, "applicationVisibility.collector.collectorType", data.NetflowCollector.ValueString())
	}
	if !data.NetflowCollectorIpAddress.IsNull() {
		body, _ = sjson.Set(body, "applicationVisibility.collector.address", data.NetflowCollectorIpAddress.ValueString())
	}
	if !data.NetflowCollectorPort.IsNull() {
		body, _ = sjson.Set(body, "applicationVisibility.collector.port", data.NetflowCollectorPort.ValueInt64())
	}
	if !data.EnableNetflowCollectorOnDevices.IsNull() {
		body, _ = sjson.Set(body, "applicationVisibility.enableOnWiredAccessDevices", data.EnableNetflowCollectorOnDevices.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TelemetrySettings) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get(""); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.wiredDataCollection.enableWiredDataCollection"); value.Exists() {
		data.EnableWiredDataCollection = types.BoolValue(value.Bool())
	} else {
		data.EnableWiredDataCollection = types.BoolNull()
	}
	if value := res.Get("response.wirelessTelemetry.enableWirelessTelemetry"); value.Exists() {
		data.EnableWirelessTelemetry = types.BoolValue(value.Bool())
	} else {
		data.EnableWirelessTelemetry = types.BoolNull()
	}
	if value := res.Get("response.snmpTraps.useBuiltinTrapServer"); value.Exists() {
		data.UseBuiltinTrapServer = types.BoolValue(value.Bool())
	} else {
		data.UseBuiltinTrapServer = types.BoolNull()
	}
	if value := res.Get("response.snmpTraps.externalTrapServers"); value.Exists() && len(value.Array()) > 0 {
		data.ExternalTrapServers = helpers.GetStringSet(value.Array())
	} else {
		data.ExternalTrapServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.syslogs.useBuiltinSyslogServer"); value.Exists() {
		data.UseBuiltinSyslogServer = types.BoolValue(value.Bool())
	} else {
		data.UseBuiltinSyslogServer = types.BoolNull()
	}
	if value := res.Get("response.syslogs.externalSyslogServers"); value.Exists() && len(value.Array()) > 0 {
		data.ExternalSyslogServers = helpers.GetStringSet(value.Array())
	} else {
		data.ExternalSyslogServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.applicationVisibility.collector.collectorType"); value.Exists() {
		data.NetflowCollector = types.StringValue(value.String())
	} else {
		data.NetflowCollector = types.StringNull()
	}
	if value := res.Get("response.applicationVisibility.collector.addres"); value.Exists() {
		data.NetflowCollectorIpAddress = types.StringValue(value.String())
	} else {
		data.NetflowCollectorIpAddress = types.StringNull()
	}
	if value := res.Get("response.applicationVisibility.collector.port"); value.Exists() {
		data.NetflowCollectorPort = types.Int64Value(value.Int())
	} else {
		data.NetflowCollectorPort = types.Int64Null()
	}
	if value := res.Get("response.applicationVisibility.enableOnWiredAccessDevices"); value.Exists() {
		data.EnableNetflowCollectorOnDevices = types.BoolValue(value.Bool())
	} else {
		data.EnableNetflowCollectorOnDevices = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TelemetrySettings) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.wiredDataCollection.enableWiredDataCollection"); value.Exists() && !data.EnableWiredDataCollection.IsNull() {
		data.EnableWiredDataCollection = types.BoolValue(value.Bool())
	} else {
		data.EnableWiredDataCollection = types.BoolNull()
	}
	if value := res.Get("response.wirelessTelemetry.enableWirelessTelemetry"); value.Exists() && !data.EnableWirelessTelemetry.IsNull() {
		data.EnableWirelessTelemetry = types.BoolValue(value.Bool())
	} else {
		data.EnableWirelessTelemetry = types.BoolNull()
	}
	if value := res.Get("response.snmpTraps.useBuiltinTrapServer"); value.Exists() && !data.UseBuiltinTrapServer.IsNull() {
		data.UseBuiltinTrapServer = types.BoolValue(value.Bool())
	} else {
		data.UseBuiltinTrapServer = types.BoolNull()
	}
	if value := res.Get("response.snmpTraps.externalTrapServers"); value.Exists() && !data.ExternalTrapServers.IsNull() {
		data.ExternalTrapServers = helpers.GetStringSet(value.Array())
	} else {
		data.ExternalTrapServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.syslogs.useBuiltinSyslogServer"); value.Exists() && !data.UseBuiltinSyslogServer.IsNull() {
		data.UseBuiltinSyslogServer = types.BoolValue(value.Bool())
	} else {
		data.UseBuiltinSyslogServer = types.BoolNull()
	}
	if value := res.Get("response.syslogs.externalSyslogServers"); value.Exists() && !data.ExternalSyslogServers.IsNull() {
		data.ExternalSyslogServers = helpers.GetStringSet(value.Array())
	} else {
		data.ExternalSyslogServers = types.SetNull(types.StringType)
	}
	if value := res.Get("response.applicationVisibility.collector.collectorType"); value.Exists() && !data.NetflowCollector.IsNull() {
		data.NetflowCollector = types.StringValue(value.String())
	} else {
		data.NetflowCollector = types.StringNull()
	}
	if value := res.Get("response.applicationVisibility.collector.addres"); value.Exists() && !data.NetflowCollectorIpAddress.IsNull() {
		data.NetflowCollectorIpAddress = types.StringValue(value.String())
	} else {
		data.NetflowCollectorIpAddress = types.StringNull()
	}
	if value := res.Get("response.applicationVisibility.collector.port"); value.Exists() && !data.NetflowCollectorPort.IsNull() {
		data.NetflowCollectorPort = types.Int64Value(value.Int())
	} else {
		data.NetflowCollectorPort = types.Int64Null()
	}
	if value := res.Get("response.applicationVisibility.enableOnWiredAccessDevices"); value.Exists() && !data.EnableNetflowCollectorOnDevices.IsNull() {
		data.EnableNetflowCollectorOnDevices = types.BoolValue(value.Bool())
	} else {
		data.EnableNetflowCollectorOnDevices = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *TelemetrySettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.EnableWiredDataCollection.IsUnknown() {
		if value := res.Get("response.wiredDataCollection.enableWiredDataCollection"); value.Exists() && !data.EnableWiredDataCollection.IsNull() {
			data.EnableWiredDataCollection = types.BoolValue(value.Bool())
		} else {
			data.EnableWiredDataCollection = types.BoolNull()
		}
	}
	if data.EnableWirelessTelemetry.IsUnknown() {
		if value := res.Get("response.wirelessTelemetry.enableWirelessTelemetry"); value.Exists() && !data.EnableWirelessTelemetry.IsNull() {
			data.EnableWirelessTelemetry = types.BoolValue(value.Bool())
		} else {
			data.EnableWirelessTelemetry = types.BoolNull()
		}
	}
	if data.UseBuiltinTrapServer.IsUnknown() {
		if value := res.Get("response.snmpTraps.useBuiltinTrapServer"); value.Exists() && !data.UseBuiltinTrapServer.IsNull() {
			data.UseBuiltinTrapServer = types.BoolValue(value.Bool())
		} else {
			data.UseBuiltinTrapServer = types.BoolNull()
		}
	}
	if value := res.Get("response.snmpTraps.externalTrapServers"); value.Exists() && !data.ExternalTrapServers.IsNull() {
		data.ExternalTrapServers = helpers.GetStringSet(value.Array())
	} else {
		data.ExternalTrapServers = types.SetNull(types.StringType)
	}
	if data.UseBuiltinSyslogServer.IsUnknown() {
		if value := res.Get("response.syslogs.useBuiltinSyslogServer"); value.Exists() && !data.UseBuiltinSyslogServer.IsNull() {
			data.UseBuiltinSyslogServer = types.BoolValue(value.Bool())
		} else {
			data.UseBuiltinSyslogServer = types.BoolNull()
		}
	}
	if value := res.Get("response.syslogs.externalSyslogServers"); value.Exists() && !data.ExternalSyslogServers.IsNull() {
		data.ExternalSyslogServers = helpers.GetStringSet(value.Array())
	} else {
		data.ExternalSyslogServers = types.SetNull(types.StringType)
	}
	if data.NetflowCollector.IsUnknown() {
		if value := res.Get("response.applicationVisibility.collector.collectorType"); value.Exists() && !data.NetflowCollector.IsNull() {
			data.NetflowCollector = types.StringValue(value.String())
		} else {
			data.NetflowCollector = types.StringNull()
		}
	}
	if data.NetflowCollectorIpAddress.IsUnknown() {
		if value := res.Get("response.applicationVisibility.collector.addres"); value.Exists() && !data.NetflowCollectorIpAddress.IsNull() {
			data.NetflowCollectorIpAddress = types.StringValue(value.String())
		} else {
			data.NetflowCollectorIpAddress = types.StringNull()
		}
	}
	if data.NetflowCollectorPort.IsUnknown() {
		if value := res.Get("response.applicationVisibility.collector.port"); value.Exists() && !data.NetflowCollectorPort.IsNull() {
			data.NetflowCollectorPort = types.Int64Value(value.Int())
		} else {
			data.NetflowCollectorPort = types.Int64Null()
		}
	}
	if data.EnableNetflowCollectorOnDevices.IsUnknown() {
		if value := res.Get("response.applicationVisibility.enableOnWiredAccessDevices"); value.Exists() && !data.EnableNetflowCollectorOnDevices.IsNull() {
			data.EnableNetflowCollectorOnDevices = types.BoolValue(value.Bool())
		} else {
			data.EnableNetflowCollectorOnDevices = types.BoolNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TelemetrySettings) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.EnableWiredDataCollection.IsNull() {
		return false
	}
	if !data.EnableWirelessTelemetry.IsNull() {
		return false
	}
	if !data.UseBuiltinTrapServer.IsNull() {
		return false
	}
	if !data.ExternalTrapServers.IsNull() {
		return false
	}
	if !data.UseBuiltinSyslogServer.IsNull() {
		return false
	}
	if !data.ExternalSyslogServers.IsNull() {
		return false
	}
	if !data.NetflowCollector.IsNull() {
		return false
	}
	if !data.NetflowCollectorIpAddress.IsNull() {
		return false
	}
	if !data.NetflowCollectorPort.IsNull() {
		return false
	}
	if !data.EnableNetflowCollectorOnDevices.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
