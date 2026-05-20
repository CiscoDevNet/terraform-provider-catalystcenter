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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PlannedAccessPointPosition struct {
	Id         types.String                       `tfsdk:"id"`
	FloorId    types.String                       `tfsdk:"floor_id"`
	Name       types.String                       `tfsdk:"name"`
	MacAddress types.String                       `tfsdk:"mac_address"`
	ApType     types.String                       `tfsdk:"ap_type"`
	PositionX  types.Float64                      `tfsdk:"position_x"`
	PositionY  types.Float64                      `tfsdk:"position_y"`
	PositionZ  types.Float64                      `tfsdk:"position_z"`
	Radios     []PlannedAccessPointPositionRadios `tfsdk:"radios"`
}

type PlannedAccessPointPositionRadios struct {
	RadioId          types.String `tfsdk:"radio_id"`
	Bands            types.Set    `tfsdk:"bands"`
	Channel          types.Int64  `tfsdk:"channel"`
	TxPower          types.Int64  `tfsdk:"tx_power"`
	AntennaName      types.String `tfsdk:"antenna_name"`
	AntennaAzimuth   types.Int64  `tfsdk:"antenna_azimuth"`
	AntennaElevation types.Int64  `tfsdk:"antenna_elevation"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PlannedAccessPointPosition) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v2/floors/%v/plannedAccessPointPositions", url.QueryEscape(data.FloorId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

func (data PlannedAccessPointPosition) getPathPost() string {
	return fmt.Sprintf("/dna/intent/api/v2/floors/%v/plannedAccessPointPositions/bulk", url.QueryEscape(data.FloorId.ValueString()))
}

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

func (data PlannedAccessPointPosition) getPathPut() string {
	return fmt.Sprintf("/dna/intent/api/v2/floors/%v/plannedAccessPointPositions/bulkChange", url.QueryEscape(data.FloorId.ValueString()))
}

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PlannedAccessPointPosition) toBody(ctx context.Context, state PlannedAccessPointPosition) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "0.id", state.Id.ValueString())
	}
	_ = put
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "0.name", data.Name.ValueString())
	}
	if !data.MacAddress.IsNull() {
		body, _ = sjson.Set(body, "0.macAddress", data.MacAddress.ValueString())
	}
	if !data.ApType.IsNull() {
		body, _ = sjson.Set(body, "0.type", data.ApType.ValueString())
	}
	if !data.PositionX.IsNull() {
		body, _ = sjson.Set(body, "0.position.x", data.PositionX.ValueFloat64())
	}
	if !data.PositionY.IsNull() {
		body, _ = sjson.Set(body, "0.position.y", data.PositionY.ValueFloat64())
	}
	if !data.PositionZ.IsNull() {
		body, _ = sjson.Set(body, "0.position.z", data.PositionZ.ValueFloat64())
	}
	if len(data.Radios) > 0 {
		body, _ = sjson.Set(body, "0.radios", []interface{}{})
		for _, item := range data.Radios {
			itemBody := ""
			if item.RadioId.ValueString() != "" && !item.RadioId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.RadioId.ValueString())
			}
			if !item.Bands.IsNull() {
				var values []string
				item.Bands.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "bands", values)
			}
			if !item.Channel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "channel", item.Channel.ValueInt64())
			}
			if !item.TxPower.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "txPower", item.TxPower.ValueInt64())
			}
			if !item.AntennaName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "antenna.name", item.AntennaName.ValueString())
			}
			if !item.AntennaAzimuth.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "antenna.azimuth", item.AntennaAzimuth.ValueInt64())
			}
			if !item.AntennaElevation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "antenna.elevation", item.AntennaElevation.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "0.radios.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PlannedAccessPointPosition) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.macAddress"); value.Exists() {
		data.MacAddress = types.StringValue(value.String())
	} else {
		data.MacAddress = types.StringNull()
	}
	if value := res.Get("response.0.type"); value.Exists() {
		data.ApType = types.StringValue(value.String())
	} else {
		data.ApType = types.StringNull()
	}
	if value := res.Get("response.0.position.x"); value.Exists() {
		data.PositionX = types.Float64Value(value.Float())
	} else {
		data.PositionX = types.Float64Null()
	}
	if value := res.Get("response.0.position.y"); value.Exists() {
		data.PositionY = types.Float64Value(value.Float())
	} else {
		data.PositionY = types.Float64Null()
	}
	if value := res.Get("response.0.position.z"); value.Exists() {
		data.PositionZ = types.Float64Value(value.Float())
	} else {
		data.PositionZ = types.Float64Null()
	}
	if value := res.Get("response.0.radios"); value.Exists() && len(value.Array()) > 0 {
		data.Radios = make([]PlannedAccessPointPositionRadios, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PlannedAccessPointPositionRadios{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.RadioId = types.StringValue(cValue.String())
			} else {
				item.RadioId = types.StringNull()
			}
			if cValue := v.Get("bands"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Bands = helpers.GetStringSet(cValue.Array())
			} else {
				item.Bands = types.SetNull(types.StringType)
			}
			if cValue := v.Get("channel"); cValue.Exists() {
				item.Channel = types.Int64Value(cValue.Int())
			} else {
				item.Channel = types.Int64Null()
			}
			if cValue := v.Get("txPower"); cValue.Exists() {
				item.TxPower = types.Int64Value(cValue.Int())
			} else {
				item.TxPower = types.Int64Null()
			}
			if cValue := v.Get("antenna.name"); cValue.Exists() {
				item.AntennaName = types.StringValue(cValue.String())
			} else {
				item.AntennaName = types.StringNull()
			}
			if cValue := v.Get("antenna.azimuth"); cValue.Exists() {
				item.AntennaAzimuth = types.Int64Value(cValue.Int())
			} else {
				item.AntennaAzimuth = types.Int64Null()
			}
			if cValue := v.Get("antenna.elevation"); cValue.Exists() {
				item.AntennaElevation = types.Int64Value(cValue.Int())
			} else {
				item.AntennaElevation = types.Int64Null()
			}
			data.Radios = append(data.Radios, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PlannedAccessPointPosition) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.macAddress"); value.Exists() && !data.MacAddress.IsNull() {
		data.MacAddress = types.StringValue(value.String())
	} else {
		data.MacAddress = types.StringNull()
	}
	if value := res.Get("response.0.type"); value.Exists() && !data.ApType.IsNull() {
		data.ApType = types.StringValue(value.String())
	} else {
		data.ApType = types.StringNull()
	}
	if value := res.Get("response.0.position.x"); value.Exists() && !data.PositionX.IsNull() {
		data.PositionX = types.Float64Value(value.Float())
	} else {
		data.PositionX = types.Float64Null()
	}
	if value := res.Get("response.0.position.y"); value.Exists() && !data.PositionY.IsNull() {
		data.PositionY = types.Float64Value(value.Float())
	} else {
		data.PositionY = types.Float64Null()
	}
	if value := res.Get("response.0.position.z"); value.Exists() && !data.PositionZ.IsNull() {
		data.PositionZ = types.Float64Value(value.Float())
	} else {
		data.PositionZ = types.Float64Null()
	}
	for i := range data.Radios {
		keys := [...]string{"id", "channel", "txPower", "antenna.name", "antenna.azimuth", "antenna.elevation"}
		keyValues := [...]string{data.Radios[i].RadioId.ValueString(), strconv.FormatInt(data.Radios[i].Channel.ValueInt64(), 10), strconv.FormatInt(data.Radios[i].TxPower.ValueInt64(), 10), data.Radios[i].AntennaName.ValueString(), strconv.FormatInt(data.Radios[i].AntennaAzimuth.ValueInt64(), 10), strconv.FormatInt(data.Radios[i].AntennaElevation.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("response.0.radios").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.Radios[i].RadioId.IsNull() {
			data.Radios[i].RadioId = types.StringValue(value.String())
		} else {
			data.Radios[i].RadioId = types.StringNull()
		}
		if value := r.Get("bands"); value.Exists() && !data.Radios[i].Bands.IsNull() {
			data.Radios[i].Bands = helpers.GetStringSet(value.Array())
		} else {
			data.Radios[i].Bands = types.SetNull(types.StringType)
		}
		if value := r.Get("channel"); value.Exists() && !data.Radios[i].Channel.IsNull() {
			data.Radios[i].Channel = types.Int64Value(value.Int())
		} else {
			data.Radios[i].Channel = types.Int64Null()
		}
		if value := r.Get("txPower"); value.Exists() && !data.Radios[i].TxPower.IsNull() {
			data.Radios[i].TxPower = types.Int64Value(value.Int())
		} else {
			data.Radios[i].TxPower = types.Int64Null()
		}
		if value := r.Get("antenna.name"); value.Exists() && !data.Radios[i].AntennaName.IsNull() {
			data.Radios[i].AntennaName = types.StringValue(value.String())
		} else {
			data.Radios[i].AntennaName = types.StringNull()
		}
		if value := r.Get("antenna.azimuth"); value.Exists() && !data.Radios[i].AntennaAzimuth.IsNull() {
			data.Radios[i].AntennaAzimuth = types.Int64Value(value.Int())
		} else {
			data.Radios[i].AntennaAzimuth = types.Int64Null()
		}
		if value := r.Get("antenna.elevation"); value.Exists() && !data.Radios[i].AntennaElevation.IsNull() {
			data.Radios[i].AntennaElevation = types.Int64Value(value.Int())
		} else {
			data.Radios[i].AntennaElevation = types.Int64Null()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PlannedAccessPointPosition) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Name.IsUnknown() {
		if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
	if data.MacAddress.IsUnknown() {
		if value := res.Get("response.0.macAddress"); value.Exists() && !data.MacAddress.IsNull() {
			data.MacAddress = types.StringValue(value.String())
		} else {
			data.MacAddress = types.StringNull()
		}
	}
	if data.ApType.IsUnknown() {
		if value := res.Get("response.0.type"); value.Exists() && !data.ApType.IsNull() {
			data.ApType = types.StringValue(value.String())
		} else {
			data.ApType = types.StringNull()
		}
	}
	if data.PositionX.IsUnknown() {
		if value := res.Get("response.0.position.x"); value.Exists() && !data.PositionX.IsNull() {
			data.PositionX = types.Float64Value(value.Float())
		} else {
			data.PositionX = types.Float64Null()
		}
	}
	if data.PositionY.IsUnknown() {
		if value := res.Get("response.0.position.y"); value.Exists() && !data.PositionY.IsNull() {
			data.PositionY = types.Float64Value(value.Float())
		} else {
			data.PositionY = types.Float64Null()
		}
	}
	if data.PositionZ.IsUnknown() {
		if value := res.Get("response.0.position.z"); value.Exists() && !data.PositionZ.IsNull() {
			data.PositionZ = types.Float64Value(value.Float())
		} else {
			data.PositionZ = types.Float64Null()
		}
	}
	for i := range data.Radios {
		keys := [...]string{"channel", "txPower", "antenna.name", "antenna.azimuth", "antenna.elevation"}
		keyValues := [...]string{strconv.FormatInt(data.Radios[i].Channel.ValueInt64(), 10), strconv.FormatInt(data.Radios[i].TxPower.ValueInt64(), 10), data.Radios[i].AntennaName.ValueString(), strconv.FormatInt(data.Radios[i].AntennaAzimuth.ValueInt64(), 10), strconv.FormatInt(data.Radios[i].AntennaElevation.ValueInt64(), 10)}

		var r gjson.Result
		res.Get("response.0.radios").ForEach(
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
		if data.Radios[i].RadioId.IsUnknown() {
			if value := r.Get("id"); value.Exists() && !data.Radios[i].RadioId.IsNull() {
				data.Radios[i].RadioId = types.StringValue(value.String())
			} else {
				data.Radios[i].RadioId = types.StringNull()
				continue
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PlannedAccessPointPosition) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.MacAddress.IsNull() {
		return false
	}
	if !data.ApType.IsNull() {
		return false
	}
	if !data.PositionX.IsNull() {
		return false
	}
	if !data.PositionY.IsNull() {
		return false
	}
	if !data.PositionZ.IsNull() {
		return false
	}
	if len(data.Radios) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
