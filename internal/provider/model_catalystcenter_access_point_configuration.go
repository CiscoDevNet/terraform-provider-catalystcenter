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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AccessPointConfiguration struct {
	Id                          types.String                                  `tfsdk:"id"`
	ApList                      []AccessPointConfigurationApList              `tfsdk:"ap_list"`
	ConfigureAdminStatus        types.Bool                                    `tfsdk:"configure_admin_status"`
	AdminStatus                 types.Bool                                    `tfsdk:"admin_status"`
	ConfigureApMode             types.Bool                                    `tfsdk:"configure_ap_mode"`
	ApMode                      types.Int64                                   `tfsdk:"ap_mode"`
	ConfigureFailoverPriority   types.Bool                                    `tfsdk:"configure_failover_priority"`
	FailoverPriority            types.Int64                                   `tfsdk:"failover_priority"`
	ConfigureLedStatus          types.Bool                                    `tfsdk:"configure_led_status"`
	LedStatus                   types.Bool                                    `tfsdk:"led_status"`
	ConfigureLedBrightnessLevel types.Bool                                    `tfsdk:"configure_led_brightness_level"`
	LedBrightnessLevel          types.Int64                                   `tfsdk:"led_brightness_level"`
	ConfigureLocation           types.Bool                                    `tfsdk:"configure_location"`
	Location                    types.String                                  `tfsdk:"location"`
	ConfigureHaController       types.Bool                                    `tfsdk:"configure_ha_controller"`
	PrimaryControllerName       types.String                                  `tfsdk:"primary_controller_name"`
	PrimaryIpAddress            types.String                                  `tfsdk:"primary_ip_address"`
	SecondaryControllerName     types.String                                  `tfsdk:"secondary_controller_name"`
	SecondaryIpAddress          types.String                                  `tfsdk:"secondary_ip_address"`
	TertiaryControllerName      types.String                                  `tfsdk:"tertiary_controller_name"`
	TertiaryIpAddress           types.String                                  `tfsdk:"tertiary_ip_address"`
	IsAssignedSiteAsLocation    types.Bool                                    `tfsdk:"is_assigned_site_as_location"`
	RadioConfigurations         []AccessPointConfigurationRadioConfigurations `tfsdk:"radio_configurations"`
}

type AccessPointConfigurationApList struct {
	MacAddress types.String `tfsdk:"mac_address"`
	ApName     types.String `tfsdk:"ap_name"`
	ApNameNew  types.String `tfsdk:"ap_name_new"`
}

type AccessPointConfigurationRadioConfigurations struct {
	RadioBand                    types.String  `tfsdk:"radio_band"`
	RadioType                    types.Int64   `tfsdk:"radio_type"`
	ConfigureRadioRoleAssignment types.Bool    `tfsdk:"configure_radio_role_assignment"`
	RadioRoleAssignment          types.String  `tfsdk:"radio_role_assignment"`
	ConfigureAdminStatus         types.Bool    `tfsdk:"configure_admin_status"`
	AdminStatus                  types.Bool    `tfsdk:"admin_status"`
	ConfigureAntennaPatternName  types.Bool    `tfsdk:"configure_antenna_pattern_name"`
	AntennaPatternName           types.String  `tfsdk:"antenna_pattern_name"`
	AntennaGain                  types.Int64   `tfsdk:"antenna_gain"`
	ConfigureAntennaCable        types.Bool    `tfsdk:"configure_antenna_cable"`
	AntennaCableName             types.String  `tfsdk:"antenna_cable_name"`
	CableLoss                    types.Float64 `tfsdk:"cable_loss"`
	ConfigureChannel             types.Bool    `tfsdk:"configure_channel"`
	ChannelAssignmentMode        types.Int64   `tfsdk:"channel_assignment_mode"`
	ChannelNumber                types.Int64   `tfsdk:"channel_number"`
	ConfigureChannelWidth        types.Bool    `tfsdk:"configure_channel_width"`
	ChannelWidth                 types.Int64   `tfsdk:"channel_width"`
	ConfigurePower               types.Bool    `tfsdk:"configure_power"`
	PowerAssignmentMode          types.Int64   `tfsdk:"power_assignment_mode"`
	PowerLevel                   types.Int64   `tfsdk:"power_level"`
	ConfigureCleanAirSi          types.Bool    `tfsdk:"configure_clean_air_si"`
	CleanAirSi                   types.Int64   `tfsdk:"clean_air_si"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AccessPointConfiguration) getPath() string {
	return "/dna/intent/api/v1/wireless/accesspoint-configuration"
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
func (data AccessPointConfiguration) toBody(ctx context.Context, state AccessPointConfiguration) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if len(data.ApList) > 0 {
		body, _ = sjson.Set(body, "apList", []interface{}{})
		for _, item := range data.ApList {
			itemBody := ""
			if !item.MacAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "macAddress", item.MacAddress.ValueString())
			}
			if !item.ApName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "apName", item.ApName.ValueString())
			}
			if !item.ApNameNew.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "apNameNew", item.ApNameNew.ValueString())
			}
			body, _ = sjson.SetRaw(body, "apList.-1", itemBody)
		}
	}
	if !data.ConfigureAdminStatus.IsNull() {
		body, _ = sjson.Set(body, "configureAdminStatus", data.ConfigureAdminStatus.ValueBool())
	}
	if !data.AdminStatus.IsNull() {
		body, _ = sjson.Set(body, "adminStatus", data.AdminStatus.ValueBool())
	}
	if !data.ConfigureApMode.IsNull() {
		body, _ = sjson.Set(body, "configureApMode", data.ConfigureApMode.ValueBool())
	}
	if !data.ApMode.IsNull() {
		body, _ = sjson.Set(body, "apMode", data.ApMode.ValueInt64())
	}
	if !data.ConfigureFailoverPriority.IsNull() {
		body, _ = sjson.Set(body, "configureFailoverPriority", data.ConfigureFailoverPriority.ValueBool())
	}
	if !data.FailoverPriority.IsNull() {
		body, _ = sjson.Set(body, "failoverPriority", data.FailoverPriority.ValueInt64())
	}
	if !data.ConfigureLedStatus.IsNull() {
		body, _ = sjson.Set(body, "configureLedStatus", data.ConfigureLedStatus.ValueBool())
	}
	if !data.LedStatus.IsNull() {
		body, _ = sjson.Set(body, "ledStatus", data.LedStatus.ValueBool())
	}
	if !data.ConfigureLedBrightnessLevel.IsNull() {
		body, _ = sjson.Set(body, "configureLedBrightnessLevel", data.ConfigureLedBrightnessLevel.ValueBool())
	}
	if !data.LedBrightnessLevel.IsNull() {
		body, _ = sjson.Set(body, "ledBrightnessLevel", data.LedBrightnessLevel.ValueInt64())
	}
	if !data.ConfigureLocation.IsNull() {
		body, _ = sjson.Set(body, "configureLocation", data.ConfigureLocation.ValueBool())
	}
	if !data.Location.IsNull() {
		body, _ = sjson.Set(body, "location", data.Location.ValueString())
	}
	if !data.ConfigureHaController.IsNull() {
		body, _ = sjson.Set(body, "configureHAController", data.ConfigureHaController.ValueBool())
	}
	if !data.PrimaryControllerName.IsNull() {
		body, _ = sjson.Set(body, "primaryControllerName", data.PrimaryControllerName.ValueString())
	}
	if !data.PrimaryIpAddress.IsNull() {
		body, _ = sjson.Set(body, "primaryIpAddress.address", data.PrimaryIpAddress.ValueString())
	}
	if !data.SecondaryControllerName.IsNull() {
		body, _ = sjson.Set(body, "secondaryControllerName", data.SecondaryControllerName.ValueString())
	}
	if !data.SecondaryIpAddress.IsNull() {
		body, _ = sjson.Set(body, "secondaryIpAddress.address", data.SecondaryIpAddress.ValueString())
	}
	if !data.TertiaryControllerName.IsNull() {
		body, _ = sjson.Set(body, "tertiaryControllerName", data.TertiaryControllerName.ValueString())
	}
	if !data.TertiaryIpAddress.IsNull() {
		body, _ = sjson.Set(body, "tertiaryIpAddress.address", data.TertiaryIpAddress.ValueString())
	}
	if !data.IsAssignedSiteAsLocation.IsNull() {
		body, _ = sjson.Set(body, "isAssignedSiteAsLocation", data.IsAssignedSiteAsLocation.ValueBool())
	}
	if len(data.RadioConfigurations) > 0 {
		body, _ = sjson.Set(body, "radioConfigurations", []interface{}{})
		for _, item := range data.RadioConfigurations {
			itemBody := ""
			if !item.RadioBand.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "radioBand", item.RadioBand.ValueString())
			}
			if !item.RadioType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "radioType", item.RadioType.ValueInt64())
			}
			if !item.ConfigureRadioRoleAssignment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureRadioRoleAssignment", item.ConfigureRadioRoleAssignment.ValueBool())
			}
			if !item.RadioRoleAssignment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "radioRoleAssignment", item.RadioRoleAssignment.ValueString())
			}
			if !item.ConfigureAdminStatus.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureAdminStatus", item.ConfigureAdminStatus.ValueBool())
			}
			if !item.AdminStatus.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "adminStatus", item.AdminStatus.ValueBool())
			}
			if !item.ConfigureAntennaPatternName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureAntennaPatternName", item.ConfigureAntennaPatternName.ValueBool())
			}
			if !item.AntennaPatternName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "antennaPatternName", item.AntennaPatternName.ValueString())
			}
			if !item.AntennaGain.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "antennaGain", item.AntennaGain.ValueInt64())
			}
			if !item.ConfigureAntennaCable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureAntennaCable", item.ConfigureAntennaCable.ValueBool())
			}
			if !item.AntennaCableName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "antennaCableName", item.AntennaCableName.ValueString())
			}
			if !item.CableLoss.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "cableLoss", item.CableLoss.ValueFloat64())
			}
			if !item.ConfigureChannel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureChannel", item.ConfigureChannel.ValueBool())
			}
			if !item.ChannelAssignmentMode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "channelAssignmentMode", item.ChannelAssignmentMode.ValueInt64())
			}
			if !item.ChannelNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "channelNumber", item.ChannelNumber.ValueInt64())
			}
			if !item.ConfigureChannelWidth.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureChannelWidth", item.ConfigureChannelWidth.ValueBool())
			}
			if !item.ChannelWidth.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "channelWidth", item.ChannelWidth.ValueInt64())
			}
			if !item.ConfigurePower.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configurePower", item.ConfigurePower.ValueBool())
			}
			if !item.PowerAssignmentMode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "powerAssignmentMode", item.PowerAssignmentMode.ValueInt64())
			}
			if !item.PowerLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "powerlevel", item.PowerLevel.ValueInt64())
			}
			if !item.ConfigureCleanAirSi.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "configureCleanAirSI", item.ConfigureCleanAirSi.ValueBool())
			}
			if !item.CleanAirSi.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "cleanAirSI", item.CleanAirSi.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "radioConfigurations.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AccessPointConfiguration) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("apList"); value.Exists() && len(value.Array()) > 0 {
		data.ApList = make([]AccessPointConfigurationApList, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessPointConfigurationApList{}
			if cValue := v.Get("macAddress"); cValue.Exists() {
				item.MacAddress = types.StringValue(cValue.String())
			} else {
				item.MacAddress = types.StringNull()
			}
			if cValue := v.Get("apName"); cValue.Exists() {
				item.ApName = types.StringValue(cValue.String())
			} else {
				item.ApName = types.StringNull()
			}
			if cValue := v.Get("apNameNew"); cValue.Exists() {
				item.ApNameNew = types.StringValue(cValue.String())
			} else {
				item.ApNameNew = types.StringNull()
			}
			data.ApList = append(data.ApList, item)
			return true
		})
	}
	if value := res.Get("configureAdminStatus"); value.Exists() {
		data.ConfigureAdminStatus = types.BoolValue(value.Bool())
	} else {
		data.ConfigureAdminStatus = types.BoolNull()
	}
	if value := res.Get("adminStatus"); value.Exists() {
		data.AdminStatus = types.BoolValue(value.Bool())
	} else {
		data.AdminStatus = types.BoolNull()
	}
	if value := res.Get("configureApMode"); value.Exists() {
		data.ConfigureApMode = types.BoolValue(value.Bool())
	} else {
		data.ConfigureApMode = types.BoolNull()
	}
	if value := res.Get("apMode"); value.Exists() {
		data.ApMode = types.Int64Value(value.Int())
	} else {
		data.ApMode = types.Int64Null()
	}
	if value := res.Get("configureFailoverPriority"); value.Exists() {
		data.ConfigureFailoverPriority = types.BoolValue(value.Bool())
	} else {
		data.ConfigureFailoverPriority = types.BoolNull()
	}
	if value := res.Get("failoverPriority"); value.Exists() {
		data.FailoverPriority = types.Int64Value(value.Int())
	} else {
		data.FailoverPriority = types.Int64Null()
	}
	if value := res.Get("configureLedStatus"); value.Exists() {
		data.ConfigureLedStatus = types.BoolValue(value.Bool())
	} else {
		data.ConfigureLedStatus = types.BoolNull()
	}
	if value := res.Get("ledStatus"); value.Exists() {
		data.LedStatus = types.BoolValue(value.Bool())
	} else {
		data.LedStatus = types.BoolNull()
	}
	if value := res.Get("configureLedBrightnessLevel"); value.Exists() {
		data.ConfigureLedBrightnessLevel = types.BoolValue(value.Bool())
	} else {
		data.ConfigureLedBrightnessLevel = types.BoolNull()
	}
	if value := res.Get("ledBrightnessLevel"); value.Exists() {
		data.LedBrightnessLevel = types.Int64Value(value.Int())
	} else {
		data.LedBrightnessLevel = types.Int64Null()
	}
	if value := res.Get("configureLocation"); value.Exists() {
		data.ConfigureLocation = types.BoolValue(value.Bool())
	} else {
		data.ConfigureLocation = types.BoolNull()
	}
	if value := res.Get("location"); value.Exists() {
		data.Location = types.StringValue(value.String())
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get("configureHAController"); value.Exists() {
		data.ConfigureHaController = types.BoolValue(value.Bool())
	} else {
		data.ConfigureHaController = types.BoolNull()
	}
	if value := res.Get("primaryControllerName"); value.Exists() {
		data.PrimaryControllerName = types.StringValue(value.String())
	} else {
		data.PrimaryControllerName = types.StringNull()
	}
	if value := res.Get("primaryIpAddress.address"); value.Exists() {
		data.PrimaryIpAddress = types.StringValue(value.String())
	} else {
		data.PrimaryIpAddress = types.StringNull()
	}
	if value := res.Get("secondaryControllerName"); value.Exists() {
		data.SecondaryControllerName = types.StringValue(value.String())
	} else {
		data.SecondaryControllerName = types.StringNull()
	}
	if value := res.Get("secondaryIpAddress.address"); value.Exists() {
		data.SecondaryIpAddress = types.StringValue(value.String())
	} else {
		data.SecondaryIpAddress = types.StringNull()
	}
	if value := res.Get("tertiaryControllerName"); value.Exists() {
		data.TertiaryControllerName = types.StringValue(value.String())
	} else {
		data.TertiaryControllerName = types.StringNull()
	}
	if value := res.Get("tertiaryIpAddress.address"); value.Exists() {
		data.TertiaryIpAddress = types.StringValue(value.String())
	} else {
		data.TertiaryIpAddress = types.StringNull()
	}
	if value := res.Get("isAssignedSiteAsLocation"); value.Exists() {
		data.IsAssignedSiteAsLocation = types.BoolValue(value.Bool())
	} else {
		data.IsAssignedSiteAsLocation = types.BoolNull()
	}
	if value := res.Get("radioConfigurations"); value.Exists() && len(value.Array()) > 0 {
		data.RadioConfigurations = make([]AccessPointConfigurationRadioConfigurations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessPointConfigurationRadioConfigurations{}
			if cValue := v.Get("radioBand"); cValue.Exists() {
				item.RadioBand = types.StringValue(cValue.String())
			} else {
				item.RadioBand = types.StringNull()
			}
			if cValue := v.Get("radioType"); cValue.Exists() {
				item.RadioType = types.Int64Value(cValue.Int())
			} else {
				item.RadioType = types.Int64Null()
			}
			if cValue := v.Get("configureRadioRoleAssignment"); cValue.Exists() {
				item.ConfigureRadioRoleAssignment = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureRadioRoleAssignment = types.BoolNull()
			}
			if cValue := v.Get("radioRoleAssignment"); cValue.Exists() {
				item.RadioRoleAssignment = types.StringValue(cValue.String())
			} else {
				item.RadioRoleAssignment = types.StringNull()
			}
			if cValue := v.Get("configureAdminStatus"); cValue.Exists() {
				item.ConfigureAdminStatus = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureAdminStatus = types.BoolNull()
			}
			if cValue := v.Get("adminStatus"); cValue.Exists() {
				item.AdminStatus = types.BoolValue(cValue.Bool())
			} else {
				item.AdminStatus = types.BoolNull()
			}
			if cValue := v.Get("configureAntennaPatternName"); cValue.Exists() {
				item.ConfigureAntennaPatternName = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureAntennaPatternName = types.BoolNull()
			}
			if cValue := v.Get("antennaPatternName"); cValue.Exists() {
				item.AntennaPatternName = types.StringValue(cValue.String())
			} else {
				item.AntennaPatternName = types.StringNull()
			}
			if cValue := v.Get("antennaGain"); cValue.Exists() {
				item.AntennaGain = types.Int64Value(cValue.Int())
			} else {
				item.AntennaGain = types.Int64Null()
			}
			if cValue := v.Get("configureAntennaCable"); cValue.Exists() {
				item.ConfigureAntennaCable = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureAntennaCable = types.BoolNull()
			}
			if cValue := v.Get("antennaCableName"); cValue.Exists() {
				item.AntennaCableName = types.StringValue(cValue.String())
			} else {
				item.AntennaCableName = types.StringNull()
			}
			if cValue := v.Get("cableLoss"); cValue.Exists() {
				item.CableLoss = types.Float64Value(cValue.Float())
			} else {
				item.CableLoss = types.Float64Null()
			}
			if cValue := v.Get("configureChannel"); cValue.Exists() {
				item.ConfigureChannel = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureChannel = types.BoolNull()
			}
			if cValue := v.Get("channelAssignmentMode"); cValue.Exists() {
				item.ChannelAssignmentMode = types.Int64Value(cValue.Int())
			} else {
				item.ChannelAssignmentMode = types.Int64Null()
			}
			if cValue := v.Get("channelNumber"); cValue.Exists() {
				item.ChannelNumber = types.Int64Value(cValue.Int())
			} else {
				item.ChannelNumber = types.Int64Null()
			}
			if cValue := v.Get("configureChannelWidth"); cValue.Exists() {
				item.ConfigureChannelWidth = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureChannelWidth = types.BoolNull()
			}
			if cValue := v.Get("channelWidth"); cValue.Exists() {
				item.ChannelWidth = types.Int64Value(cValue.Int())
			} else {
				item.ChannelWidth = types.Int64Null()
			}
			if cValue := v.Get("configurePower"); cValue.Exists() {
				item.ConfigurePower = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigurePower = types.BoolNull()
			}
			if cValue := v.Get("powerAssignmentMode"); cValue.Exists() {
				item.PowerAssignmentMode = types.Int64Value(cValue.Int())
			} else {
				item.PowerAssignmentMode = types.Int64Null()
			}
			if cValue := v.Get("powerlevel"); cValue.Exists() {
				item.PowerLevel = types.Int64Value(cValue.Int())
			} else {
				item.PowerLevel = types.Int64Null()
			}
			if cValue := v.Get("configureCleanAirSI"); cValue.Exists() {
				item.ConfigureCleanAirSi = types.BoolValue(cValue.Bool())
			} else {
				item.ConfigureCleanAirSi = types.BoolNull()
			}
			if cValue := v.Get("cleanAirSI"); cValue.Exists() {
				item.CleanAirSi = types.Int64Value(cValue.Int())
			} else {
				item.CleanAirSi = types.Int64Null()
			}
			data.RadioConfigurations = append(data.RadioConfigurations, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AccessPointConfiguration) updateFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.ApList {
		keys := [...]string{"macAddress", "apName", "apNameNew"}
		keyValues := [...]string{data.ApList[i].MacAddress.ValueString(), data.ApList[i].ApName.ValueString(), data.ApList[i].ApNameNew.ValueString()}

		var r gjson.Result
		res.Get("apList").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("macAddress"); value.Exists() && !data.ApList[i].MacAddress.IsNull() {
			data.ApList[i].MacAddress = types.StringValue(value.String())
		} else {
			data.ApList[i].MacAddress = types.StringNull()
		}
		if value := r.Get("apName"); value.Exists() && !data.ApList[i].ApName.IsNull() {
			data.ApList[i].ApName = types.StringValue(value.String())
		} else {
			data.ApList[i].ApName = types.StringNull()
		}
		if value := r.Get("apNameNew"); value.Exists() && !data.ApList[i].ApNameNew.IsNull() {
			data.ApList[i].ApNameNew = types.StringValue(value.String())
		} else {
			data.ApList[i].ApNameNew = types.StringNull()
		}
	}
	if value := res.Get("configureAdminStatus"); value.Exists() && !data.ConfigureAdminStatus.IsNull() {
		data.ConfigureAdminStatus = types.BoolValue(value.Bool())
	} else {
		data.ConfigureAdminStatus = types.BoolNull()
	}
	if value := res.Get("adminStatus"); value.Exists() && !data.AdminStatus.IsNull() {
		data.AdminStatus = types.BoolValue(value.Bool())
	} else {
		data.AdminStatus = types.BoolNull()
	}
	if value := res.Get("configureApMode"); value.Exists() && !data.ConfigureApMode.IsNull() {
		data.ConfigureApMode = types.BoolValue(value.Bool())
	} else {
		data.ConfigureApMode = types.BoolNull()
	}
	if value := res.Get("apMode"); value.Exists() && !data.ApMode.IsNull() {
		data.ApMode = types.Int64Value(value.Int())
	} else {
		data.ApMode = types.Int64Null()
	}
	if value := res.Get("configureFailoverPriority"); value.Exists() && !data.ConfigureFailoverPriority.IsNull() {
		data.ConfigureFailoverPriority = types.BoolValue(value.Bool())
	} else {
		data.ConfigureFailoverPriority = types.BoolNull()
	}
	if value := res.Get("failoverPriority"); value.Exists() && !data.FailoverPriority.IsNull() {
		data.FailoverPriority = types.Int64Value(value.Int())
	} else {
		data.FailoverPriority = types.Int64Null()
	}
	if value := res.Get("configureLedStatus"); value.Exists() && !data.ConfigureLedStatus.IsNull() {
		data.ConfigureLedStatus = types.BoolValue(value.Bool())
	} else {
		data.ConfigureLedStatus = types.BoolNull()
	}
	if value := res.Get("ledStatus"); value.Exists() && !data.LedStatus.IsNull() {
		data.LedStatus = types.BoolValue(value.Bool())
	} else {
		data.LedStatus = types.BoolNull()
	}
	if value := res.Get("configureLedBrightnessLevel"); value.Exists() && !data.ConfigureLedBrightnessLevel.IsNull() {
		data.ConfigureLedBrightnessLevel = types.BoolValue(value.Bool())
	} else {
		data.ConfigureLedBrightnessLevel = types.BoolNull()
	}
	if value := res.Get("ledBrightnessLevel"); value.Exists() && !data.LedBrightnessLevel.IsNull() {
		data.LedBrightnessLevel = types.Int64Value(value.Int())
	} else {
		data.LedBrightnessLevel = types.Int64Null()
	}
	if value := res.Get("configureLocation"); value.Exists() && !data.ConfigureLocation.IsNull() {
		data.ConfigureLocation = types.BoolValue(value.Bool())
	} else {
		data.ConfigureLocation = types.BoolNull()
	}
	if value := res.Get("location"); value.Exists() && !data.Location.IsNull() {
		data.Location = types.StringValue(value.String())
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get("configureHAController"); value.Exists() && !data.ConfigureHaController.IsNull() {
		data.ConfigureHaController = types.BoolValue(value.Bool())
	} else {
		data.ConfigureHaController = types.BoolNull()
	}
	if value := res.Get("primaryControllerName"); value.Exists() && !data.PrimaryControllerName.IsNull() {
		data.PrimaryControllerName = types.StringValue(value.String())
	} else {
		data.PrimaryControllerName = types.StringNull()
	}
	if value := res.Get("primaryIpAddress.address"); value.Exists() && !data.PrimaryIpAddress.IsNull() {
		data.PrimaryIpAddress = types.StringValue(value.String())
	} else {
		data.PrimaryIpAddress = types.StringNull()
	}
	if value := res.Get("secondaryControllerName"); value.Exists() && !data.SecondaryControllerName.IsNull() {
		data.SecondaryControllerName = types.StringValue(value.String())
	} else {
		data.SecondaryControllerName = types.StringNull()
	}
	if value := res.Get("secondaryIpAddress.address"); value.Exists() && !data.SecondaryIpAddress.IsNull() {
		data.SecondaryIpAddress = types.StringValue(value.String())
	} else {
		data.SecondaryIpAddress = types.StringNull()
	}
	if value := res.Get("tertiaryControllerName"); value.Exists() && !data.TertiaryControllerName.IsNull() {
		data.TertiaryControllerName = types.StringValue(value.String())
	} else {
		data.TertiaryControllerName = types.StringNull()
	}
	if value := res.Get("tertiaryIpAddress.address"); value.Exists() && !data.TertiaryIpAddress.IsNull() {
		data.TertiaryIpAddress = types.StringValue(value.String())
	} else {
		data.TertiaryIpAddress = types.StringNull()
	}
	if value := res.Get("isAssignedSiteAsLocation"); value.Exists() && !data.IsAssignedSiteAsLocation.IsNull() {
		data.IsAssignedSiteAsLocation = types.BoolValue(value.Bool())
	} else {
		data.IsAssignedSiteAsLocation = types.BoolNull()
	}
	for i := range data.RadioConfigurations {
		keys := [...]string{"radioBand", "radioType", "configureRadioRoleAssignment", "radioRoleAssignment", "configureAdminStatus", "adminStatus", "configureAntennaPatternName", "antennaPatternName", "antennaGain", "configureAntennaCable", "antennaCableName", "configureChannel", "channelAssignmentMode", "channelNumber", "configureChannelWidth", "channelWidth", "configurePower", "powerAssignmentMode", "powerlevel", "configureCleanAirSI", "cleanAirSI"}
		keyValues := [...]string{data.RadioConfigurations[i].RadioBand.ValueString(), strconv.FormatInt(data.RadioConfigurations[i].RadioType.ValueInt64(), 10), strconv.FormatBool(data.RadioConfigurations[i].ConfigureRadioRoleAssignment.ValueBool()), data.RadioConfigurations[i].RadioRoleAssignment.ValueString(), strconv.FormatBool(data.RadioConfigurations[i].ConfigureAdminStatus.ValueBool()), strconv.FormatBool(data.RadioConfigurations[i].AdminStatus.ValueBool()), strconv.FormatBool(data.RadioConfigurations[i].ConfigureAntennaPatternName.ValueBool()), data.RadioConfigurations[i].AntennaPatternName.ValueString(), strconv.FormatInt(data.RadioConfigurations[i].AntennaGain.ValueInt64(), 10), strconv.FormatBool(data.RadioConfigurations[i].ConfigureAntennaCable.ValueBool()), data.RadioConfigurations[i].AntennaCableName.ValueString(), strconv.FormatBool(data.RadioConfigurations[i].ConfigureChannel.ValueBool()), strconv.FormatInt(data.RadioConfigurations[i].ChannelAssignmentMode.ValueInt64(), 10), strconv.FormatInt(data.RadioConfigurations[i].ChannelNumber.ValueInt64(), 10), strconv.FormatBool(data.RadioConfigurations[i].ConfigureChannelWidth.ValueBool()), strconv.FormatInt(data.RadioConfigurations[i].ChannelWidth.ValueInt64(), 10), strconv.FormatBool(data.RadioConfigurations[i].ConfigurePower.ValueBool()), strconv.FormatInt(data.RadioConfigurations[i].PowerAssignmentMode.ValueInt64(), 10), strconv.FormatInt(data.RadioConfigurations[i].PowerLevel.ValueInt64(), 10), strconv.FormatBool(data.RadioConfigurations[i].ConfigureCleanAirSi.ValueBool()), strconv.FormatInt(data.RadioConfigurations[i].CleanAirSi.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("radioConfigurations").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("radioBand"); value.Exists() && !data.RadioConfigurations[i].RadioBand.IsNull() {
			data.RadioConfigurations[i].RadioBand = types.StringValue(value.String())
		} else {
			data.RadioConfigurations[i].RadioBand = types.StringNull()
		}
		if value := r.Get("radioType"); value.Exists() && !data.RadioConfigurations[i].RadioType.IsNull() {
			data.RadioConfigurations[i].RadioType = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].RadioType = types.Int64Null()
		}
		if value := r.Get("configureRadioRoleAssignment"); value.Exists() && !data.RadioConfigurations[i].ConfigureRadioRoleAssignment.IsNull() {
			data.RadioConfigurations[i].ConfigureRadioRoleAssignment = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureRadioRoleAssignment = types.BoolNull()
		}
		if value := r.Get("radioRoleAssignment"); value.Exists() && !data.RadioConfigurations[i].RadioRoleAssignment.IsNull() {
			data.RadioConfigurations[i].RadioRoleAssignment = types.StringValue(value.String())
		} else {
			data.RadioConfigurations[i].RadioRoleAssignment = types.StringNull()
		}
		if value := r.Get("configureAdminStatus"); value.Exists() && !data.RadioConfigurations[i].ConfigureAdminStatus.IsNull() {
			data.RadioConfigurations[i].ConfigureAdminStatus = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureAdminStatus = types.BoolNull()
		}
		if value := r.Get("adminStatus"); value.Exists() && !data.RadioConfigurations[i].AdminStatus.IsNull() {
			data.RadioConfigurations[i].AdminStatus = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].AdminStatus = types.BoolNull()
		}
		if value := r.Get("configureAntennaPatternName"); value.Exists() && !data.RadioConfigurations[i].ConfigureAntennaPatternName.IsNull() {
			data.RadioConfigurations[i].ConfigureAntennaPatternName = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureAntennaPatternName = types.BoolNull()
		}
		if value := r.Get("antennaPatternName"); value.Exists() && !data.RadioConfigurations[i].AntennaPatternName.IsNull() {
			data.RadioConfigurations[i].AntennaPatternName = types.StringValue(value.String())
		} else {
			data.RadioConfigurations[i].AntennaPatternName = types.StringNull()
		}
		if value := r.Get("antennaGain"); value.Exists() && !data.RadioConfigurations[i].AntennaGain.IsNull() {
			data.RadioConfigurations[i].AntennaGain = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].AntennaGain = types.Int64Null()
		}
		if value := r.Get("configureAntennaCable"); value.Exists() && !data.RadioConfigurations[i].ConfigureAntennaCable.IsNull() {
			data.RadioConfigurations[i].ConfigureAntennaCable = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureAntennaCable = types.BoolNull()
		}
		if value := r.Get("antennaCableName"); value.Exists() && !data.RadioConfigurations[i].AntennaCableName.IsNull() {
			data.RadioConfigurations[i].AntennaCableName = types.StringValue(value.String())
		} else {
			data.RadioConfigurations[i].AntennaCableName = types.StringNull()
		}
		if value := r.Get("cableLoss"); value.Exists() && !data.RadioConfigurations[i].CableLoss.IsNull() {
			data.RadioConfigurations[i].CableLoss = types.Float64Value(value.Float())
		} else {
			data.RadioConfigurations[i].CableLoss = types.Float64Null()
		}
		if value := r.Get("configureChannel"); value.Exists() && !data.RadioConfigurations[i].ConfigureChannel.IsNull() {
			data.RadioConfigurations[i].ConfigureChannel = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureChannel = types.BoolNull()
		}
		if value := r.Get("channelAssignmentMode"); value.Exists() && !data.RadioConfigurations[i].ChannelAssignmentMode.IsNull() {
			data.RadioConfigurations[i].ChannelAssignmentMode = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].ChannelAssignmentMode = types.Int64Null()
		}
		if value := r.Get("channelNumber"); value.Exists() && !data.RadioConfigurations[i].ChannelNumber.IsNull() {
			data.RadioConfigurations[i].ChannelNumber = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].ChannelNumber = types.Int64Null()
		}
		if value := r.Get("configureChannelWidth"); value.Exists() && !data.RadioConfigurations[i].ConfigureChannelWidth.IsNull() {
			data.RadioConfigurations[i].ConfigureChannelWidth = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureChannelWidth = types.BoolNull()
		}
		if value := r.Get("channelWidth"); value.Exists() && !data.RadioConfigurations[i].ChannelWidth.IsNull() {
			data.RadioConfigurations[i].ChannelWidth = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].ChannelWidth = types.Int64Null()
		}
		if value := r.Get("configurePower"); value.Exists() && !data.RadioConfigurations[i].ConfigurePower.IsNull() {
			data.RadioConfigurations[i].ConfigurePower = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigurePower = types.BoolNull()
		}
		if value := r.Get("powerAssignmentMode"); value.Exists() && !data.RadioConfigurations[i].PowerAssignmentMode.IsNull() {
			data.RadioConfigurations[i].PowerAssignmentMode = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].PowerAssignmentMode = types.Int64Null()
		}
		if value := r.Get("powerlevel"); value.Exists() && !data.RadioConfigurations[i].PowerLevel.IsNull() {
			data.RadioConfigurations[i].PowerLevel = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].PowerLevel = types.Int64Null()
		}
		if value := r.Get("configureCleanAirSI"); value.Exists() && !data.RadioConfigurations[i].ConfigureCleanAirSi.IsNull() {
			data.RadioConfigurations[i].ConfigureCleanAirSi = types.BoolValue(value.Bool())
		} else {
			data.RadioConfigurations[i].ConfigureCleanAirSi = types.BoolNull()
		}
		if value := r.Get("cleanAirSI"); value.Exists() && !data.RadioConfigurations[i].CleanAirSi.IsNull() {
			data.RadioConfigurations[i].CleanAirSi = types.Int64Value(value.Int())
		} else {
			data.RadioConfigurations[i].CleanAirSi = types.Int64Null()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AccessPointConfiguration) isNull(ctx context.Context, res gjson.Result) bool {
	if len(data.ApList) > 0 {
		return false
	}
	if !data.ConfigureAdminStatus.IsNull() {
		return false
	}
	if !data.AdminStatus.IsNull() {
		return false
	}
	if !data.ConfigureApMode.IsNull() {
		return false
	}
	if !data.ApMode.IsNull() {
		return false
	}
	if !data.ConfigureFailoverPriority.IsNull() {
		return false
	}
	if !data.FailoverPriority.IsNull() {
		return false
	}
	if !data.ConfigureLedStatus.IsNull() {
		return false
	}
	if !data.LedStatus.IsNull() {
		return false
	}
	if !data.ConfigureLedBrightnessLevel.IsNull() {
		return false
	}
	if !data.LedBrightnessLevel.IsNull() {
		return false
	}
	if !data.ConfigureLocation.IsNull() {
		return false
	}
	if !data.Location.IsNull() {
		return false
	}
	if !data.ConfigureHaController.IsNull() {
		return false
	}
	if !data.PrimaryControllerName.IsNull() {
		return false
	}
	if !data.PrimaryIpAddress.IsNull() {
		return false
	}
	if !data.SecondaryControllerName.IsNull() {
		return false
	}
	if !data.SecondaryIpAddress.IsNull() {
		return false
	}
	if !data.TertiaryControllerName.IsNull() {
		return false
	}
	if !data.TertiaryIpAddress.IsNull() {
		return false
	}
	if !data.IsAssignedSiteAsLocation.IsNull() {
		return false
	}
	if len(data.RadioConfigurations) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
