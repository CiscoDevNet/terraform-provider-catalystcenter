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

// Section below is manually modified - generator markings removed for imports
import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Time conversion helpers for calendar power profiles
// API requires 12-hour AM/PM format, but we accept 24-hour format from users

// convertTo12Hour converts 24-hour format (e.g., "22:00") to 12-hour AM/PM format (e.g., "10:00 PM")
func convertTo12Hour(time24 string) string {
	if time24 == "" {
		return ""
	}
	// If already in 12-hour format, return as-is
	if strings.Contains(strings.ToUpper(time24), "AM") || strings.Contains(strings.ToUpper(time24), "PM") {
		return time24
	}
	parts := strings.Split(time24, ":")
	if len(parts) != 2 {
		return time24
	}
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return time24
	}
	minute := parts[1]

	suffix := "AM"
	if hour >= 12 {
		suffix = "PM"
	}
	if hour > 12 {
		hour -= 12
	}
	if hour == 0 {
		hour = 12
	}
	return fmt.Sprintf("%02d:%s %s", hour, minute, suffix)
}

// convertTo24Hour converts 12-hour AM/PM format (e.g., "10:00 PM") to 24-hour format (e.g., "22:00")
func convertTo24Hour(time12 string) string {
	if time12 == "" {
		return ""
	}
	// If already in 24-hour format (no AM/PM), return as-is
	upper := strings.ToUpper(time12)
	if !strings.Contains(upper, "AM") && !strings.Contains(upper, "PM") {
		return time12
	}

	isPM := strings.Contains(upper, "PM")
	// Remove AM/PM suffix
	timeOnly := strings.TrimSpace(strings.Replace(strings.Replace(upper, "AM", "", 1), "PM", "", 1))
	parts := strings.Split(timeOnly, ":")
	if len(parts) != 2 {
		return time12
	}
	hour, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return time12
	}
	minute := strings.TrimSpace(parts[1])

	if isPM && hour != 12 {
		hour += 12
	}
	if !isPM && hour == 12 {
		hour = 0
	}
	return fmt.Sprintf("%02d:%s", hour, minute)
}

// normalizeTimeZone maps API response (ALL-CAPS) to user's original value or Title Case
// API accepts Title Case but returns ALL-CAPS
func normalizeTimeZone(apiValue string, userValue string) string {
	// Mapping from ALL-CAPS (API response) to Title Case (API docs/user-friendly)
	timeZoneMap := map[string]string{
		"NOT CONFIGURED":        "Not Configured",
		"CONTROLLER":            "Controller",
		"DELTA FROM CONTROLLER": "Delta from Controller",
	}

	// If user specified a value that matches case-insensitively, preserve user's casing
	if userValue != "" && strings.EqualFold(apiValue, userValue) {
		return userValue
	}

	// Otherwise normalize to Title Case (matching API documentation)
	if normalized, ok := timeZoneMap[strings.ToUpper(apiValue)]; ok {
		return normalized
	}
	return apiValue
}

// normalizeCountryCode handles API returning "unconfigured" which is not in the enum
// Treats "unconfigured" as null/empty (no country code set)
func normalizeCountryCode(apiValue string) string {
	if strings.EqualFold(apiValue, "unconfigured") {
		return "" // Treat as null
	}
	return apiValue
}

// End of manually modified section

// APProfile types
type APProfile struct {
	Id                              types.String                     `tfsdk:"id"`
	ApProfileName                   types.String                     `tfsdk:"ap_profile_name"`
	Description                     types.String                     `tfsdk:"description"`
	RemoteWorkerEnabled             types.Bool                       `tfsdk:"remote_worker_enabled"`
	AuthType                        types.String                     `tfsdk:"auth_type"`
	Dot1xUsername                   types.String                     `tfsdk:"dot1x_username"`
	Dot1xPassword                   types.String                     `tfsdk:"dot1x_password"`
	SshEnabled                      types.Bool                       `tfsdk:"ssh_enabled"`
	TelnetEnabled                   types.Bool                       `tfsdk:"telnet_enabled"`
	ManagementUserName              types.String                     `tfsdk:"management_user_name"`
	ManagementPassword              types.String                     `tfsdk:"management_password"`
	ManagementEnablePassword        types.String                     `tfsdk:"management_enable_password"`
	CdpState                        types.Bool                       `tfsdk:"cdp_state"`
	AwipsEnabled                    types.Bool                       `tfsdk:"awips_enabled"`
	AwipsForensicEnabled            types.Bool                       `tfsdk:"awips_forensic_enabled"`
	RogueDetection                  types.Bool                       `tfsdk:"rogue_detection"`
	RogueDetectionMinRssi           types.Int64                      `tfsdk:"rogue_detection_min_rssi"`
	RogueDetectionTransientInterval types.Int64                      `tfsdk:"rogue_detection_transient_interval"`
	RogueDetectionReportInterval    types.Int64                      `tfsdk:"rogue_detection_report_interval"`
	PmfDenialEnabled                types.Bool                       `tfsdk:"pmf_denial_enabled"`
	MeshEnabled                     types.Bool                       `tfsdk:"mesh_enabled"`
	BridgeGroupName                 types.String                     `tfsdk:"bridge_group_name"`
	BackhaulClientAccess            types.Bool                       `tfsdk:"backhaul_client_access"`
	Range                           types.Int64                      `tfsdk:"range"`
	Ghz5BackhaulDataRates           types.String                     `tfsdk:"ghz5_backhaul_data_rates"`
	Ghz24BackhaulDataRates          types.String                     `tfsdk:"ghz24_backhaul_data_rates"`
	RapDownlinkBackhaul             types.String                     `tfsdk:"rap_downlink_backhaul"`
	ApPowerProfileName              types.String                     `tfsdk:"ap_power_profile_name"`
	CalendarPowerProfiles           []APProfileCalendarPowerProfiles `tfsdk:"calendar_power_profiles"`
	CountryCode                     types.String                     `tfsdk:"country_code"`
	TimeZone                        types.String                     `tfsdk:"time_zone"`
	TimeZoneOffsetHour              types.Int64                      `tfsdk:"time_zone_offset_hour"`
	TimeZoneOffsetMinutes           types.Int64                      `tfsdk:"time_zone_offset_minutes"`
	ClientLimit                     types.Int64                      `tfsdk:"client_limit"`
}

type APProfileCalendarPowerProfiles struct {
	CalendarProfileName types.String `tfsdk:"calendar_profile_name"`
	PowerProfileName    types.String `tfsdk:"power_profile_name"`
	SchedulerType       types.String `tfsdk:"scheduler_type"`
	SchedulerStartTime  types.String `tfsdk:"scheduler_start_time"`
	SchedulerEndTime    types.String `tfsdk:"scheduler_end_time"`
	SchedulerDay        types.Set    `tfsdk:"scheduler_day"`
	SchedulerDate       types.Set    `tfsdk:"scheduler_date"`
}

func (data APProfile) getPath() string {
	return "/dna/intent/api/v1/wirelessSettings/apProfiles"
}

func (data APProfile) toBody(ctx context.Context, state APProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.ApProfileName.IsNull() {
		body, _ = sjson.Set(body, "apProfileName", data.ApProfileName.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.RemoteWorkerEnabled.IsNull() {
		body, _ = sjson.Set(body, "remoteWorkerEnabled", data.RemoteWorkerEnabled.ValueBool())
	}
	if !data.AuthType.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.authType", data.AuthType.ValueString())
	}
	if !data.Dot1xUsername.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.dot1xUsername", data.Dot1xUsername.ValueString())
	}
	if !data.Dot1xPassword.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.dot1xPassword", data.Dot1xPassword.ValueString())
	}
	if !data.SshEnabled.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.sshEnabled", data.SshEnabled.ValueBool())
	}
	if !data.TelnetEnabled.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.telnetEnabled", data.TelnetEnabled.ValueBool())
	}
	if !data.ManagementUserName.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.managementUserName", data.ManagementUserName.ValueString())
	}
	if !data.ManagementPassword.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.managementPassword", data.ManagementPassword.ValueString())
	}
	if !data.ManagementEnablePassword.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.managementEnablePassword", data.ManagementEnablePassword.ValueString())
	}
	if !data.CdpState.IsNull() {
		body, _ = sjson.Set(body, "managementSetting.cdpState", data.CdpState.ValueBool())
	}
	if !data.AwipsEnabled.IsNull() {
		body, _ = sjson.Set(body, "awipsEnabled", data.AwipsEnabled.ValueBool())
	}
	if !data.AwipsForensicEnabled.IsNull() {
		body, _ = sjson.Set(body, "awipsForensicEnabled", data.AwipsForensicEnabled.ValueBool())
	}
	if !data.RogueDetection.IsNull() {
		body, _ = sjson.Set(body, "rogueDetectionSetting.rogueDetection", data.RogueDetection.ValueBool())
	}
	if !data.RogueDetectionMinRssi.IsNull() {
		body, _ = sjson.Set(body, "rogueDetectionSetting.rogueDetectionMinRssi", data.RogueDetectionMinRssi.ValueInt64())
	}
	if !data.RogueDetectionTransientInterval.IsNull() {
		body, _ = sjson.Set(body, "rogueDetectionSetting.rogueDetectionTransientInterval", data.RogueDetectionTransientInterval.ValueInt64())
	}
	if !data.RogueDetectionReportInterval.IsNull() {
		body, _ = sjson.Set(body, "rogueDetectionSetting.rogueDetectionReportInterval", data.RogueDetectionReportInterval.ValueInt64())
	}
	if !data.PmfDenialEnabled.IsNull() {
		body, _ = sjson.Set(body, "pmfDenialEnabled", data.PmfDenialEnabled.ValueBool())
	}
	if !data.MeshEnabled.IsNull() {
		body, _ = sjson.Set(body, "meshEnabled", data.MeshEnabled.ValueBool())
	}
	if !data.BridgeGroupName.IsNull() {
		body, _ = sjson.Set(body, "meshSetting.bridgeGroupName", data.BridgeGroupName.ValueString())
	}
	if !data.BackhaulClientAccess.IsNull() {
		body, _ = sjson.Set(body, "meshSetting.backhaulClientAccess", data.BackhaulClientAccess.ValueBool())
	}
	if !data.Range.IsNull() {
		body, _ = sjson.Set(body, "meshSetting.range", data.Range.ValueInt64())
	}
	if !data.Ghz5BackhaulDataRates.IsNull() {
		body, _ = sjson.Set(body, "meshSetting.ghz5BackhaulDataRates", data.Ghz5BackhaulDataRates.ValueString())
	}
	if !data.Ghz24BackhaulDataRates.IsNull() {
		body, _ = sjson.Set(body, "meshSetting.ghz24BackhaulDataRates", data.Ghz24BackhaulDataRates.ValueString())
	}
	if !data.RapDownlinkBackhaul.IsNull() {
		body, _ = sjson.Set(body, "meshSetting.rapDownlinkBackhaul", data.RapDownlinkBackhaul.ValueString())
	}
	if !data.ApPowerProfileName.IsNull() {
		body, _ = sjson.Set(body, "apPowerProfileName", data.ApPowerProfileName.ValueString())
	}
	if len(data.CalendarPowerProfiles) > 0 {
		body, _ = sjson.Set(body, "calendarPowerProfiles", []interface{}{})
		for _, item := range data.CalendarPowerProfiles {
			itemBody := ""
			if !item.PowerProfileName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "powerProfileName", item.PowerProfileName.ValueString())
			}
			if !item.SchedulerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "schedulerType", item.SchedulerType.ValueString())
			}
			if !item.SchedulerStartTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "duration.schedulerStartTime", convertTo12Hour(item.SchedulerStartTime.ValueString()))
			}
			if !item.SchedulerEndTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "duration.schedulerEndTime", convertTo12Hour(item.SchedulerEndTime.ValueString()))
			}
			if !item.SchedulerDay.IsNull() {
				var days []string
				item.SchedulerDay.ElementsAs(ctx, &days, false)
				// Convert to lowercase for API
				for i := range days {
					days[i] = strings.ToLower(days[i])
				}
				itemBody, _ = sjson.Set(itemBody, "duration.schedulerDay", days)
			}
			if !item.SchedulerDate.IsNull() {
				var dates []string
				item.SchedulerDate.ElementsAs(ctx, &dates, false)
				itemBody, _ = sjson.Set(itemBody, "duration.schedulerDate", dates)
			}
			body, _ = sjson.SetRaw(body, "calendarPowerProfiles.-1", itemBody)
		}
	}
	if !data.CountryCode.IsNull() {
		body, _ = sjson.Set(body, "countryCode", data.CountryCode.ValueString())
	}
	if !data.TimeZone.IsNull() {
		body, _ = sjson.Set(body, "timeZone", data.TimeZone.ValueString())
	}
	if !data.TimeZoneOffsetHour.IsNull() {
		body, _ = sjson.Set(body, "timeZoneOffsetHour", data.TimeZoneOffsetHour.ValueInt64())
	}
	if !data.TimeZoneOffsetMinutes.IsNull() {
		body, _ = sjson.Set(body, "timeZoneOffsetMinutes", data.TimeZoneOffsetMinutes.ValueInt64())
	}
	if !data.ClientLimit.IsNull() {
		body, _ = sjson.Set(body, "clientLimit", data.ClientLimit.ValueInt64())
	}
	return body
}

func (data *APProfile) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.apProfileName"); value.Exists() {
		data.ApProfileName = types.StringValue(value.String())
	} else {
		data.ApProfileName = types.StringNull()
	}
	if value := res.Get("response.0.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.0.remoteWorkerEnabled"); value.Exists() {
		data.RemoteWorkerEnabled = types.BoolValue(value.Bool())
	} else {
		data.RemoteWorkerEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.managementSetting.authType"); value.Exists() {
		data.AuthType = types.StringValue(value.String())
	} else {
		data.AuthType = types.StringValue("NO-AUTH")
	}
	if value := res.Get("response.0.managementSetting.dot1xUsername"); value.Exists() && value.Type != gjson.Null {
		data.Dot1xUsername = types.StringValue(value.String())
	} else {
		data.Dot1xUsername = types.StringNull()
	}
	if value := res.Get("response.0.managementSetting.sshEnabled"); value.Exists() {
		data.SshEnabled = types.BoolValue(value.Bool())
	} else {
		data.SshEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.managementSetting.telnetEnabled"); value.Exists() {
		data.TelnetEnabled = types.BoolValue(value.Bool())
	} else {
		data.TelnetEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.managementSetting.managementUserName"); value.Exists() && value.Type != gjson.Null {
		data.ManagementUserName = types.StringValue(value.String())
	} else {
		data.ManagementUserName = types.StringNull()
	}
	if value := res.Get("response.0.managementSetting.cdpState"); value.Exists() {
		data.CdpState = types.BoolValue(value.Bool())
	} else {
		data.CdpState = types.BoolValue(false)
	}
	if value := res.Get("response.0.awipsEnabled"); value.Exists() {
		data.AwipsEnabled = types.BoolValue(value.Bool())
	} else {
		data.AwipsEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.awipsForensicEnabled"); value.Exists() {
		data.AwipsForensicEnabled = types.BoolValue(value.Bool())
	} else {
		data.AwipsForensicEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetection"); value.Exists() {
		data.RogueDetection = types.BoolValue(value.Bool())
	} else {
		data.RogueDetection = types.BoolValue(false)
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetectionMinRssi"); value.Exists() {
		data.RogueDetectionMinRssi = types.Int64Value(value.Int())
	} else {
		data.RogueDetectionMinRssi = types.Int64Value(-90)
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetectionTransientInterval"); value.Exists() {
		data.RogueDetectionTransientInterval = types.Int64Value(value.Int())
	} else {
		data.RogueDetectionTransientInterval = types.Int64Value(0)
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetectionReportInterval"); value.Exists() {
		data.RogueDetectionReportInterval = types.Int64Value(value.Int())
	} else {
		data.RogueDetectionReportInterval = types.Int64Value(10)
	}
	if value := res.Get("response.0.pmfDenialEnabled"); value.Exists() {
		data.PmfDenialEnabled = types.BoolValue(value.Bool())
	} else {
		data.PmfDenialEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.meshEnabled"); value.Exists() {
		data.MeshEnabled = types.BoolValue(value.Bool())
	} else {
		data.MeshEnabled = types.BoolValue(false)
	}
	if value := res.Get("response.0.meshSetting.bridgeGroupName"); value.Exists() {
		data.BridgeGroupName = types.StringValue(value.String())
	} else {
		data.BridgeGroupName = types.StringNull()
	}
	if value := res.Get("response.0.meshSetting.backhaulClientAccess"); value.Exists() {
		data.BackhaulClientAccess = types.BoolValue(value.Bool())
	} else {
		data.BackhaulClientAccess = types.BoolNull()
	}
	if value := res.Get("response.0.meshSetting.range"); value.Exists() {
		data.Range = types.Int64Value(value.Int())
	} else {
		data.Range = types.Int64Null()
	}
	if value := res.Get("response.0.meshSetting.ghz5BackhaulDataRates"); value.Exists() {
		data.Ghz5BackhaulDataRates = types.StringValue(value.String())
	} else {
		data.Ghz5BackhaulDataRates = types.StringNull()
	}
	if value := res.Get("response.0.meshSetting.ghz24BackhaulDataRates"); value.Exists() {
		data.Ghz24BackhaulDataRates = types.StringValue(value.String())
	} else {
		data.Ghz24BackhaulDataRates = types.StringNull()
	}
	if value := res.Get("response.0.meshSetting.rapDownlinkBackhaul"); value.Exists() {
		data.RapDownlinkBackhaul = types.StringValue(value.String())
	} else {
		data.RapDownlinkBackhaul = types.StringNull()
	}
	if value := res.Get("response.0.apPowerProfileName"); value.Exists() && value.Type != gjson.Null {
		data.ApPowerProfileName = types.StringValue(value.String())
	} else {
		data.ApPowerProfileName = types.StringNull()
	}
	if value := res.Get("response.0.calendarPowerProfiles"); value.Exists() && len(value.Array()) > 0 {
		data.CalendarPowerProfiles = make([]APProfileCalendarPowerProfiles, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := APProfileCalendarPowerProfiles{}
			if cValue := v.Get("calendarProfileName"); cValue.Exists() && cValue.Type != gjson.Null {
				item.CalendarProfileName = types.StringValue(cValue.String())
			} else {
				item.CalendarProfileName = types.StringNull()
			}
			if cValue := v.Get("powerProfileName"); cValue.Exists() {
				item.PowerProfileName = types.StringValue(cValue.String())
			} else {
				item.PowerProfileName = types.StringNull()
			}
			if cValue := v.Get("schedulerType"); cValue.Exists() {
				item.SchedulerType = types.StringValue(cValue.String())
			} else {
				item.SchedulerType = types.StringNull()
			}
			if cValue := v.Get("duration.schedulerStartTime"); cValue.Exists() {
				item.SchedulerStartTime = types.StringValue(convertTo24Hour(cValue.String()))
			} else {
				item.SchedulerStartTime = types.StringNull()
			}
			if cValue := v.Get("duration.schedulerEndTime"); cValue.Exists() {
				item.SchedulerEndTime = types.StringValue(convertTo24Hour(cValue.String()))
			} else {
				item.SchedulerEndTime = types.StringNull()
			}
			if cValue := v.Get("duration.schedulerDay"); cValue.Exists() && cValue.Type != gjson.Null && cValue.IsArray() {
				dayValues := []attr.Value{}
				cValue.ForEach(func(_, dayVal gjson.Result) bool {
					dayValues = append(dayValues, types.StringValue(strings.ToLower(dayVal.String())))
					return true
				})
				item.SchedulerDay, _ = types.SetValue(types.StringType, dayValues)
			} else {
				item.SchedulerDay = types.SetNull(types.StringType)
			}
			if cValue := v.Get("duration.schedulerDate"); cValue.Exists() && cValue.Type != gjson.Null && cValue.IsArray() {
				dateValues := []attr.Value{}
				cValue.ForEach(func(_, dateVal gjson.Result) bool {
					dateValues = append(dateValues, types.StringValue(dateVal.String()))
					return true
				})
				item.SchedulerDate, _ = types.SetValue(types.StringType, dateValues)
			} else {
				item.SchedulerDate = types.SetNull(types.StringType)
			}
			data.CalendarPowerProfiles = append(data.CalendarPowerProfiles, item)
			return true
		})
	}
	if value := res.Get("response.0.countryCode"); value.Exists() {
		normalized := normalizeCountryCode(value.String())
		if normalized == "" {
			data.CountryCode = types.StringNull()
		} else {
			data.CountryCode = types.StringValue(normalized)
		}
	} else {
		data.CountryCode = types.StringNull()
	}
	if value := res.Get("response.0.timeZone"); value.Exists() {
		data.TimeZone = types.StringValue(normalizeTimeZone(value.String(), ""))
	} else {
		data.TimeZone = types.StringValue("Not Configured")
	}
	if value := res.Get("response.0.timeZoneOffsetHour"); value.Exists() {
		data.TimeZoneOffsetHour = types.Int64Value(value.Int())
	} else {
		data.TimeZoneOffsetHour = types.Int64Value(0)
	}
	if value := res.Get("response.0.timeZoneOffsetMinutes"); value.Exists() {
		data.TimeZoneOffsetMinutes = types.Int64Value(value.Int())
	} else {
		data.TimeZoneOffsetMinutes = types.Int64Value(0)
	}
	if value := res.Get("response.0.clientLimit"); value.Exists() {
		data.ClientLimit = types.Int64Value(value.Int())
	} else {
		data.ClientLimit = types.Int64Value(0)
	}
}

func (data *APProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.apProfileName"); value.Exists() && !data.ApProfileName.IsNull() {
		data.ApProfileName = types.StringValue(value.String())
	} else {
		data.ApProfileName = types.StringNull()
	}
	if value := res.Get("response.0.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.0.remoteWorkerEnabled"); value.Exists() && !data.RemoteWorkerEnabled.IsNull() {
		data.RemoteWorkerEnabled = types.BoolValue(value.Bool())
	} else if data.RemoteWorkerEnabled.ValueBool() != false {
		data.RemoteWorkerEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.managementSetting.authType"); value.Exists() && !data.AuthType.IsNull() {
		data.AuthType = types.StringValue(value.String())
	} else if data.AuthType.ValueString() != "NO-AUTH" {
		data.AuthType = types.StringNull()
	}
	if value := res.Get("response.0.managementSetting.dot1xUsername"); value.Exists() && value.Type != gjson.Null && !data.Dot1xUsername.IsNull() {
		data.Dot1xUsername = types.StringValue(value.String())
	} else {
		data.Dot1xUsername = types.StringNull()
	}
	if value := res.Get("response.0.managementSetting.sshEnabled"); value.Exists() && !data.SshEnabled.IsNull() {
		data.SshEnabled = types.BoolValue(value.Bool())
	} else if data.SshEnabled.ValueBool() != false {
		data.SshEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.managementSetting.telnetEnabled"); value.Exists() && !data.TelnetEnabled.IsNull() {
		data.TelnetEnabled = types.BoolValue(value.Bool())
	} else if data.TelnetEnabled.ValueBool() != false {
		data.TelnetEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.managementSetting.managementUserName"); value.Exists() && value.Type != gjson.Null && !data.ManagementUserName.IsNull() {
		data.ManagementUserName = types.StringValue(value.String())
	} else {
		data.ManagementUserName = types.StringNull()
	}
	if value := res.Get("response.0.managementSetting.cdpState"); value.Exists() && !data.CdpState.IsNull() {
		data.CdpState = types.BoolValue(value.Bool())
	} else if data.CdpState.ValueBool() != false {
		data.CdpState = types.BoolNull()
	}
	if value := res.Get("response.0.awipsEnabled"); value.Exists() && !data.AwipsEnabled.IsNull() {
		data.AwipsEnabled = types.BoolValue(value.Bool())
	} else if data.AwipsEnabled.ValueBool() != false {
		data.AwipsEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.awipsForensicEnabled"); value.Exists() && !data.AwipsForensicEnabled.IsNull() {
		data.AwipsForensicEnabled = types.BoolValue(value.Bool())
	} else if data.AwipsForensicEnabled.ValueBool() != false {
		data.AwipsForensicEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetection"); value.Exists() && !data.RogueDetection.IsNull() {
		data.RogueDetection = types.BoolValue(value.Bool())
	} else if data.RogueDetection.ValueBool() != false {
		data.RogueDetection = types.BoolNull()
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetectionMinRssi"); value.Exists() && !data.RogueDetectionMinRssi.IsNull() {
		data.RogueDetectionMinRssi = types.Int64Value(value.Int())
	} else if data.RogueDetectionMinRssi.ValueInt64() != -90 {
		data.RogueDetectionMinRssi = types.Int64Null()
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetectionTransientInterval"); value.Exists() && !data.RogueDetectionTransientInterval.IsNull() {
		data.RogueDetectionTransientInterval = types.Int64Value(value.Int())
	} else if data.RogueDetectionTransientInterval.ValueInt64() != 0 {
		data.RogueDetectionTransientInterval = types.Int64Null()
	}
	if value := res.Get("response.0.rogueDetectionSetting.rogueDetectionReportInterval"); value.Exists() && !data.RogueDetectionReportInterval.IsNull() {
		data.RogueDetectionReportInterval = types.Int64Value(value.Int())
	} else if data.RogueDetectionReportInterval.ValueInt64() != 10 {
		data.RogueDetectionReportInterval = types.Int64Null()
	}
	if value := res.Get("response.0.pmfDenialEnabled"); value.Exists() && !data.PmfDenialEnabled.IsNull() {
		data.PmfDenialEnabled = types.BoolValue(value.Bool())
	} else if data.PmfDenialEnabled.ValueBool() != false {
		data.PmfDenialEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.meshEnabled"); value.Exists() && !data.MeshEnabled.IsNull() {
		data.MeshEnabled = types.BoolValue(value.Bool())
	} else if data.MeshEnabled.ValueBool() != false {
		data.MeshEnabled = types.BoolNull()
	}
	if value := res.Get("response.0.meshSetting.bridgeGroupName"); value.Exists() && !data.BridgeGroupName.IsNull() {
		data.BridgeGroupName = types.StringValue(value.String())
	} else {
		data.BridgeGroupName = types.StringNull()
	}
	if value := res.Get("response.0.meshSetting.backhaulClientAccess"); value.Exists() && !data.BackhaulClientAccess.IsNull() {
		data.BackhaulClientAccess = types.BoolValue(value.Bool())
	} else {
		data.BackhaulClientAccess = types.BoolNull()
	}
	if value := res.Get("response.0.meshSetting.range"); value.Exists() && !data.Range.IsNull() {
		data.Range = types.Int64Value(value.Int())
	} else {
		data.Range = types.Int64Null()
	}
	if value := res.Get("response.0.meshSetting.ghz5BackhaulDataRates"); value.Exists() && !data.Ghz5BackhaulDataRates.IsNull() {
		data.Ghz5BackhaulDataRates = types.StringValue(value.String())
	} else {
		data.Ghz5BackhaulDataRates = types.StringNull()
	}
	if value := res.Get("response.0.meshSetting.ghz24BackhaulDataRates"); value.Exists() && !data.Ghz24BackhaulDataRates.IsNull() {
		data.Ghz24BackhaulDataRates = types.StringValue(value.String())
	} else {
		data.Ghz24BackhaulDataRates = types.StringNull()
	}
	if value := res.Get("response.0.meshSetting.rapDownlinkBackhaul"); value.Exists() && !data.RapDownlinkBackhaul.IsNull() {
		data.RapDownlinkBackhaul = types.StringValue(value.String())
	} else {
		data.RapDownlinkBackhaul = types.StringNull()
	}
	if value := res.Get("response.0.apPowerProfileName"); value.Exists() && value.Type != gjson.Null && !data.ApPowerProfileName.IsNull() {
		data.ApPowerProfileName = types.StringValue(value.String())
	} else {
		data.ApPowerProfileName = types.StringNull()
	}
	for i := range data.CalendarPowerProfiles {
		// Keys for matching - use only the fixed fields (not arrays) for matching
		// Match based on powerProfileName and schedulerType which should uniquely identify a calendar profile
		var r gjson.Result
		res.Get("response.0.calendarPowerProfiles").ForEach(
			func(_, v gjson.Result) bool {
				// Match by powerProfileName and schedulerType
				if v.Get("powerProfileName").String() == data.CalendarPowerProfiles[i].PowerProfileName.ValueString() &&
					v.Get("schedulerType").String() == data.CalendarPowerProfiles[i].SchedulerType.ValueString() {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("calendarProfileName"); value.Exists() && value.Type != gjson.Null {
			data.CalendarPowerProfiles[i].CalendarProfileName = types.StringValue(value.String())
		} else {
			data.CalendarPowerProfiles[i].CalendarProfileName = types.StringNull()
		}
		if value := r.Get("powerProfileName"); value.Exists() && !data.CalendarPowerProfiles[i].PowerProfileName.IsNull() {
			data.CalendarPowerProfiles[i].PowerProfileName = types.StringValue(value.String())
		} else {
			data.CalendarPowerProfiles[i].PowerProfileName = types.StringNull()
		}
		if value := r.Get("schedulerType"); value.Exists() && !data.CalendarPowerProfiles[i].SchedulerType.IsNull() {
			data.CalendarPowerProfiles[i].SchedulerType = types.StringValue(value.String())
		} else {
			data.CalendarPowerProfiles[i].SchedulerType = types.StringNull()
		}
		if value := r.Get("duration.schedulerStartTime"); value.Exists() && !data.CalendarPowerProfiles[i].SchedulerStartTime.IsNull() {
			data.CalendarPowerProfiles[i].SchedulerStartTime = types.StringValue(convertTo24Hour(value.String()))
		} else {
			data.CalendarPowerProfiles[i].SchedulerStartTime = types.StringNull()
		}
		if value := r.Get("duration.schedulerEndTime"); value.Exists() && !data.CalendarPowerProfiles[i].SchedulerEndTime.IsNull() {
			data.CalendarPowerProfiles[i].SchedulerEndTime = types.StringValue(convertTo24Hour(value.String()))
		} else {
			data.CalendarPowerProfiles[i].SchedulerEndTime = types.StringNull()
		}
		if value := r.Get("duration.schedulerDay"); value.Exists() && value.Type != gjson.Null && value.IsArray() && !data.CalendarPowerProfiles[i].SchedulerDay.IsNull() {
			dayValues := []attr.Value{}
			value.ForEach(func(_, dayVal gjson.Result) bool {
				dayValues = append(dayValues, types.StringValue(strings.ToLower(dayVal.String())))
				return true
			})
			data.CalendarPowerProfiles[i].SchedulerDay, _ = types.SetValue(types.StringType, dayValues)
		} else {
			data.CalendarPowerProfiles[i].SchedulerDay = types.SetNull(types.StringType)
		}
		if value := r.Get("duration.schedulerDate"); value.Exists() && value.Type != gjson.Null && value.IsArray() && !data.CalendarPowerProfiles[i].SchedulerDate.IsNull() {
			dateValues := []attr.Value{}
			value.ForEach(func(_, dateVal gjson.Result) bool {
				dateValues = append(dateValues, types.StringValue(dateVal.String()))
				return true
			})
			data.CalendarPowerProfiles[i].SchedulerDate, _ = types.SetValue(types.StringType, dateValues)
		} else {
			data.CalendarPowerProfiles[i].SchedulerDate = types.SetNull(types.StringType)
		}
	}
	if value := res.Get("response.0.countryCode"); value.Exists() && !data.CountryCode.IsNull() {
		normalized := normalizeCountryCode(value.String())
		if normalized == "" {
			data.CountryCode = types.StringNull()
		} else {
			data.CountryCode = types.StringValue(normalized)
		}
	} else {
		data.CountryCode = types.StringNull()
	}
	if value := res.Get("response.0.timeZone"); value.Exists() && !data.TimeZone.IsNull() {
		data.TimeZone = types.StringValue(normalizeTimeZone(value.String(), data.TimeZone.ValueString()))
	} else if data.TimeZone.ValueString() != "Not Configured" {
		data.TimeZone = types.StringNull()
	}
	if value := res.Get("response.0.timeZoneOffsetHour"); value.Exists() && !data.TimeZoneOffsetHour.IsNull() {
		data.TimeZoneOffsetHour = types.Int64Value(value.Int())
	} else if data.TimeZoneOffsetHour.ValueInt64() != 0 {
		data.TimeZoneOffsetHour = types.Int64Null()
	}
	if value := res.Get("response.0.timeZoneOffsetMinutes"); value.Exists() && !data.TimeZoneOffsetMinutes.IsNull() {
		data.TimeZoneOffsetMinutes = types.Int64Value(value.Int())
	} else if data.TimeZoneOffsetMinutes.ValueInt64() != 0 {
		data.TimeZoneOffsetMinutes = types.Int64Null()
	}
	if value := res.Get("response.0.clientLimit"); value.Exists() && !data.ClientLimit.IsNull() {
		data.ClientLimit = types.Int64Value(value.Int())
	} else if data.ClientLimit.ValueInt64() != 0 {
		data.ClientLimit = types.Int64Null()
	}
}

func (data *APProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Description.IsNull() {
		return false
	}
	if !data.RemoteWorkerEnabled.IsNull() {
		return false
	}
	if !data.AuthType.IsNull() {
		return false
	}
	if !data.Dot1xUsername.IsNull() {
		return false
	}
	if !data.Dot1xPassword.IsNull() {
		return false
	}
	if !data.SshEnabled.IsNull() {
		return false
	}
	if !data.TelnetEnabled.IsNull() {
		return false
	}
	if !data.ManagementUserName.IsNull() {
		return false
	}
	if !data.ManagementPassword.IsNull() {
		return false
	}
	if !data.ManagementEnablePassword.IsNull() {
		return false
	}
	if !data.CdpState.IsNull() {
		return false
	}
	if !data.AwipsEnabled.IsNull() {
		return false
	}
	if !data.AwipsForensicEnabled.IsNull() {
		return false
	}
	if !data.RogueDetection.IsNull() {
		return false
	}
	if !data.RogueDetectionMinRssi.IsNull() {
		return false
	}
	if !data.RogueDetectionTransientInterval.IsNull() {
		return false
	}
	if !data.RogueDetectionReportInterval.IsNull() {
		return false
	}
	if !data.PmfDenialEnabled.IsNull() {
		return false
	}
	if !data.MeshEnabled.IsNull() {
		return false
	}
	if !data.BridgeGroupName.IsNull() {
		return false
	}
	if !data.BackhaulClientAccess.IsNull() {
		return false
	}
	if !data.Range.IsNull() {
		return false
	}
	if !data.Ghz5BackhaulDataRates.IsNull() {
		return false
	}
	if !data.Ghz24BackhaulDataRates.IsNull() {
		return false
	}
	if !data.RapDownlinkBackhaul.IsNull() {
		return false
	}
	if !data.ApPowerProfileName.IsNull() {
		return false
	}
	if len(data.CalendarPowerProfiles) > 0 {
		return false
	}
	if !data.CountryCode.IsNull() {
		return false
	}
	if !data.TimeZone.IsNull() {
		return false
	}
	if !data.TimeZoneOffsetHour.IsNull() {
		return false
	}
	if !data.TimeZoneOffsetMinutes.IsNull() {
		return false
	}
	if !data.ClientLimit.IsNull() {
		return false
	}
	return true
}
