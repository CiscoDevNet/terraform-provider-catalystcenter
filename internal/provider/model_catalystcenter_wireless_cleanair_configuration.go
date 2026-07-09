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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type WirelessCleanAirConfiguration struct {
	Id                              types.String `tfsdk:"id"`
	DesignName                      types.String `tfsdk:"design_name"`
	RadioBand                       types.String `tfsdk:"radio_band"`
	CleanAir                        types.Bool   `tfsdk:"clean_air"`
	CleanAirDeviceReporting         types.Bool   `tfsdk:"clean_air_device_reporting"`
	PersistentDevicePropagation     types.Bool   `tfsdk:"persistent_device_propagation"`
	Description                     types.String `tfsdk:"description"`
	UnlockedAttributes              types.List   `tfsdk:"unlocked_attributes"`
	BleBeacon                       types.Bool   `tfsdk:"ble_beacon"`
	BluetoothPagingInquiry          types.Bool   `tfsdk:"bluetooth_paging_inquiry"`
	BluetoothScoAcl                 types.Bool   `tfsdk:"bluetooth_sco_acl"`
	ContinuousTransmitter           types.Bool   `tfsdk:"continuous_transmitter"`
	GenericDect                     types.Bool   `tfsdk:"generic_dect"`
	GenericTdd                      types.Bool   `tfsdk:"generic_tdd"`
	Jammer                          types.Bool   `tfsdk:"jammer"`
	MicrowaveOven                   types.Bool   `tfsdk:"microwave_oven"`
	MotorolaCanopy                  types.Bool   `tfsdk:"motorola_canopy"`
	SiFhss                          types.Bool   `tfsdk:"si_fhss"`
	Spectrum80211Fh                 types.Bool   `tfsdk:"spectrum_80211_fh"`
	Spectrum80211NonStandardChannel types.Bool   `tfsdk:"spectrum_80211_non_standard_channel"`
	Spectrum802154                  types.Bool   `tfsdk:"spectrum_802154"`
	SpectrumInverted                types.Bool   `tfsdk:"spectrum_inverted"`
	SuperAg                         types.Bool   `tfsdk:"super_ag"`
	VideoCamera                     types.Bool   `tfsdk:"video_camera"`
	WimaxFixed                      types.Bool   `tfsdk:"wimax_fixed"`
	WimaxMobile                     types.Bool   `tfsdk:"wimax_mobile"`
	Xbox                            types.Bool   `tfsdk:"xbox"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data WirelessCleanAirConfiguration) getPath() string {
	return "/dna/intent/api/v1/featureTemplates/wireless/cleanAirConfigurations"
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
func (data WirelessCleanAirConfiguration) toBody(ctx context.Context, state WirelessCleanAirConfiguration) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DesignName.IsNull() {
		body, _ = sjson.Set(body, "designName", data.DesignName.ValueString())
	}
	if !data.RadioBand.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.radioBand", data.RadioBand.ValueString())
	}
	if !data.CleanAir.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.cleanAir", data.CleanAir.ValueBool())
	}
	if !data.CleanAirDeviceReporting.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.cleanAirDeviceReporting", data.CleanAirDeviceReporting.ValueBool())
	}
	if !data.PersistentDevicePropagation.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.persistentDevicePropagation", data.PersistentDevicePropagation.ValueBool())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.description", data.Description.ValueString())
	}
	if !data.UnlockedAttributes.IsNull() {
		var values []string
		data.UnlockedAttributes.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "unlockedAttributes", values)
	}
	if !data.BleBeacon.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.bleBeacon", data.BleBeacon.ValueBool())
	}
	if !data.BluetoothPagingInquiry.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.bluetoothPagingInquiry", data.BluetoothPagingInquiry.ValueBool())
	}
	if !data.BluetoothScoAcl.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.bluetoothScoAcl", data.BluetoothScoAcl.ValueBool())
	}
	if !data.ContinuousTransmitter.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.continuousTransmitter", data.ContinuousTransmitter.ValueBool())
	}
	if !data.GenericDect.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.genericDect", data.GenericDect.ValueBool())
	}
	if !data.GenericTdd.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.genericTdd", data.GenericTdd.ValueBool())
	}
	if !data.Jammer.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.jammer", data.Jammer.ValueBool())
	}
	if !data.MicrowaveOven.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.microwaveOven", data.MicrowaveOven.ValueBool())
	}
	if !data.MotorolaCanopy.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.motorolaCanopy", data.MotorolaCanopy.ValueBool())
	}
	if !data.SiFhss.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.siFhss", data.SiFhss.ValueBool())
	}
	if !data.Spectrum80211Fh.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.spectrum80211Fh", data.Spectrum80211Fh.ValueBool())
	}
	if !data.Spectrum80211NonStandardChannel.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.spectrum80211NonStandardChannel", data.Spectrum80211NonStandardChannel.ValueBool())
	}
	if !data.Spectrum802154.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.spectrum802154", data.Spectrum802154.ValueBool())
	}
	if !data.SpectrumInverted.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.spectrumInverted", data.SpectrumInverted.ValueBool())
	}
	if !data.SuperAg.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.superAg", data.SuperAg.ValueBool())
	}
	if !data.VideoCamera.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.videoCamera", data.VideoCamera.ValueBool())
	}
	if !data.WimaxFixed.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.wimaxFixed", data.WimaxFixed.ValueBool())
	}
	if !data.WimaxMobile.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.wimaxMobile", data.WimaxMobile.ValueBool())
	}
	if !data.Xbox.IsNull() {
		body, _ = sjson.Set(body, "featureAttributes.interferersFeatures.xbox", data.Xbox.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *WirelessCleanAirConfiguration) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.designName"); value.Exists() {
		data.DesignName = types.StringValue(value.String())
	} else {
		data.DesignName = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.radioBand"); value.Exists() {
		data.RadioBand = types.StringValue(value.String())
	} else {
		data.RadioBand = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.cleanAir"); value.Exists() {
		data.CleanAir = types.BoolValue(value.Bool())
	} else {
		data.CleanAir = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.cleanAirDeviceReporting"); value.Exists() {
		data.CleanAirDeviceReporting = types.BoolValue(value.Bool())
	} else {
		data.CleanAirDeviceReporting = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.persistentDevicePropagation"); value.Exists() {
		data.PersistentDevicePropagation = types.BoolValue(value.Bool())
	} else {
		data.PersistentDevicePropagation = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.unlockedAttributes"); value.Exists() && len(value.Array()) > 0 {
		data.UnlockedAttributes = helpers.GetStringList(value.Array())
	} else {
		data.UnlockedAttributes = types.ListNull(types.StringType)
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.bleBeacon"); value.Exists() {
		data.BleBeacon = types.BoolValue(value.Bool())
	} else {
		data.BleBeacon = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.bluetoothPagingInquiry"); value.Exists() {
		data.BluetoothPagingInquiry = types.BoolValue(value.Bool())
	} else {
		data.BluetoothPagingInquiry = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.bluetoothScoAcl"); value.Exists() {
		data.BluetoothScoAcl = types.BoolValue(value.Bool())
	} else {
		data.BluetoothScoAcl = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.continuousTransmitter"); value.Exists() {
		data.ContinuousTransmitter = types.BoolValue(value.Bool())
	} else {
		data.ContinuousTransmitter = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.genericDect"); value.Exists() {
		data.GenericDect = types.BoolValue(value.Bool())
	} else {
		data.GenericDect = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.genericTdd"); value.Exists() {
		data.GenericTdd = types.BoolValue(value.Bool())
	} else {
		data.GenericTdd = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.jammer"); value.Exists() {
		data.Jammer = types.BoolValue(value.Bool())
	} else {
		data.Jammer = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.microwaveOven"); value.Exists() {
		data.MicrowaveOven = types.BoolValue(value.Bool())
	} else {
		data.MicrowaveOven = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.motorolaCanopy"); value.Exists() {
		data.MotorolaCanopy = types.BoolValue(value.Bool())
	} else {
		data.MotorolaCanopy = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.siFhss"); value.Exists() {
		data.SiFhss = types.BoolValue(value.Bool())
	} else {
		data.SiFhss = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrum80211Fh"); value.Exists() {
		data.Spectrum80211Fh = types.BoolValue(value.Bool())
	} else {
		data.Spectrum80211Fh = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrum80211NonStandardChannel"); value.Exists() {
		data.Spectrum80211NonStandardChannel = types.BoolValue(value.Bool())
	} else {
		data.Spectrum80211NonStandardChannel = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrum802154"); value.Exists() {
		data.Spectrum802154 = types.BoolValue(value.Bool())
	} else {
		data.Spectrum802154 = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrumInverted"); value.Exists() {
		data.SpectrumInverted = types.BoolValue(value.Bool())
	} else {
		data.SpectrumInverted = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.superAg"); value.Exists() {
		data.SuperAg = types.BoolValue(value.Bool())
	} else {
		data.SuperAg = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.videoCamera"); value.Exists() {
		data.VideoCamera = types.BoolValue(value.Bool())
	} else {
		data.VideoCamera = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.wimaxFixed"); value.Exists() {
		data.WimaxFixed = types.BoolValue(value.Bool())
	} else {
		data.WimaxFixed = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.wimaxMobile"); value.Exists() {
		data.WimaxMobile = types.BoolValue(value.Bool())
	} else {
		data.WimaxMobile = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.xbox"); value.Exists() {
		data.Xbox = types.BoolValue(value.Bool())
	} else {
		data.Xbox = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *WirelessCleanAirConfiguration) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.designName"); value.Exists() && !data.DesignName.IsNull() {
		data.DesignName = types.StringValue(value.String())
	} else {
		data.DesignName = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.radioBand"); value.Exists() && !data.RadioBand.IsNull() {
		data.RadioBand = types.StringValue(value.String())
	} else {
		data.RadioBand = types.StringNull()
	}
	if value := res.Get("response.featureAttributes.cleanAir"); value.Exists() && !data.CleanAir.IsNull() {
		data.CleanAir = types.BoolValue(value.Bool())
	} else {
		data.CleanAir = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.cleanAirDeviceReporting"); value.Exists() && !data.CleanAirDeviceReporting.IsNull() {
		data.CleanAirDeviceReporting = types.BoolValue(value.Bool())
	} else {
		data.CleanAirDeviceReporting = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.persistentDevicePropagation"); value.Exists() && !data.PersistentDevicePropagation.IsNull() {
		data.PersistentDevicePropagation = types.BoolValue(value.Bool())
	} else {
		data.PersistentDevicePropagation = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.unlockedAttributes"); value.Exists() && !data.UnlockedAttributes.IsNull() {
		data.UnlockedAttributes = helpers.GetStringList(value.Array())
	} else {
		data.UnlockedAttributes = types.ListNull(types.StringType)
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.bleBeacon"); value.Exists() && !data.BleBeacon.IsNull() {
		data.BleBeacon = types.BoolValue(value.Bool())
	} else {
		data.BleBeacon = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.bluetoothPagingInquiry"); value.Exists() && !data.BluetoothPagingInquiry.IsNull() {
		data.BluetoothPagingInquiry = types.BoolValue(value.Bool())
	} else {
		data.BluetoothPagingInquiry = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.bluetoothScoAcl"); value.Exists() && !data.BluetoothScoAcl.IsNull() {
		data.BluetoothScoAcl = types.BoolValue(value.Bool())
	} else {
		data.BluetoothScoAcl = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.continuousTransmitter"); value.Exists() && !data.ContinuousTransmitter.IsNull() {
		data.ContinuousTransmitter = types.BoolValue(value.Bool())
	} else {
		data.ContinuousTransmitter = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.genericDect"); value.Exists() && !data.GenericDect.IsNull() {
		data.GenericDect = types.BoolValue(value.Bool())
	} else {
		data.GenericDect = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.genericTdd"); value.Exists() && !data.GenericTdd.IsNull() {
		data.GenericTdd = types.BoolValue(value.Bool())
	} else {
		data.GenericTdd = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.jammer"); value.Exists() && !data.Jammer.IsNull() {
		data.Jammer = types.BoolValue(value.Bool())
	} else {
		data.Jammer = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.microwaveOven"); value.Exists() && !data.MicrowaveOven.IsNull() {
		data.MicrowaveOven = types.BoolValue(value.Bool())
	} else {
		data.MicrowaveOven = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.motorolaCanopy"); value.Exists() && !data.MotorolaCanopy.IsNull() {
		data.MotorolaCanopy = types.BoolValue(value.Bool())
	} else {
		data.MotorolaCanopy = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.siFhss"); value.Exists() && !data.SiFhss.IsNull() {
		data.SiFhss = types.BoolValue(value.Bool())
	} else {
		data.SiFhss = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrum80211Fh"); value.Exists() && !data.Spectrum80211Fh.IsNull() {
		data.Spectrum80211Fh = types.BoolValue(value.Bool())
	} else {
		data.Spectrum80211Fh = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrum80211NonStandardChannel"); value.Exists() && !data.Spectrum80211NonStandardChannel.IsNull() {
		data.Spectrum80211NonStandardChannel = types.BoolValue(value.Bool())
	} else {
		data.Spectrum80211NonStandardChannel = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrum802154"); value.Exists() && !data.Spectrum802154.IsNull() {
		data.Spectrum802154 = types.BoolValue(value.Bool())
	} else {
		data.Spectrum802154 = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.spectrumInverted"); value.Exists() && !data.SpectrumInverted.IsNull() {
		data.SpectrumInverted = types.BoolValue(value.Bool())
	} else {
		data.SpectrumInverted = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.superAg"); value.Exists() && !data.SuperAg.IsNull() {
		data.SuperAg = types.BoolValue(value.Bool())
	} else {
		data.SuperAg = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.videoCamera"); value.Exists() && !data.VideoCamera.IsNull() {
		data.VideoCamera = types.BoolValue(value.Bool())
	} else {
		data.VideoCamera = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.wimaxFixed"); value.Exists() && !data.WimaxFixed.IsNull() {
		data.WimaxFixed = types.BoolValue(value.Bool())
	} else {
		data.WimaxFixed = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.wimaxMobile"); value.Exists() && !data.WimaxMobile.IsNull() {
		data.WimaxMobile = types.BoolValue(value.Bool())
	} else {
		data.WimaxMobile = types.BoolNull()
	}
	if value := res.Get("response.featureAttributes.interferersFeatures.xbox"); value.Exists() && !data.Xbox.IsNull() {
		data.Xbox = types.BoolValue(value.Bool())
	} else {
		data.Xbox = types.BoolNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *WirelessCleanAirConfiguration) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DesignName.IsNull() {
		return false
	}
	if !data.RadioBand.IsNull() {
		return false
	}
	if !data.CleanAir.IsNull() {
		return false
	}
	if !data.CleanAirDeviceReporting.IsNull() {
		return false
	}
	if !data.PersistentDevicePropagation.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.UnlockedAttributes.IsNull() {
		return false
	}
	if !data.BleBeacon.IsNull() {
		return false
	}
	if !data.BluetoothPagingInquiry.IsNull() {
		return false
	}
	if !data.BluetoothScoAcl.IsNull() {
		return false
	}
	if !data.ContinuousTransmitter.IsNull() {
		return false
	}
	if !data.GenericDect.IsNull() {
		return false
	}
	if !data.GenericTdd.IsNull() {
		return false
	}
	if !data.Jammer.IsNull() {
		return false
	}
	if !data.MicrowaveOven.IsNull() {
		return false
	}
	if !data.MotorolaCanopy.IsNull() {
		return false
	}
	if !data.SiFhss.IsNull() {
		return false
	}
	if !data.Spectrum80211Fh.IsNull() {
		return false
	}
	if !data.Spectrum80211NonStandardChannel.IsNull() {
		return false
	}
	if !data.Spectrum802154.IsNull() {
		return false
	}
	if !data.SpectrumInverted.IsNull() {
		return false
	}
	if !data.SuperAg.IsNull() {
		return false
	}
	if !data.VideoCamera.IsNull() {
		return false
	}
	if !data.WimaxFixed.IsNull() {
		return false
	}
	if !data.WimaxMobile.IsNull() {
		return false
	}
	if !data.Xbox.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
