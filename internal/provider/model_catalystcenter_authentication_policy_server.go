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
type AuthenticationPolicyServer struct {
	Id                         types.String                                           `tfsdk:"id"`
	AuthenticationPort         types.Int64                                            `tfsdk:"authentication_port"`
	AccountingPort             types.Int64                                            `tfsdk:"accounting_port"`
	CiscoIseDtos               []AuthenticationPolicyServerCiscoIseDtos               `tfsdk:"cisco_ise_dtos"`
	IpAddress                  types.String                                           `tfsdk:"ip_address"`
	PxgridEnabled              types.Bool                                             `tfsdk:"pxgrid_enabled"`
	UseDnacCertForPxgrid       types.Bool                                             `tfsdk:"use_dnac_cert_for_pxgrid"`
	IsIseEnabled               types.Bool                                             `tfsdk:"is_ise_enabled"`
	Port                       types.Int64                                            `tfsdk:"port"`
	Protocol                   types.String                                           `tfsdk:"protocol"`
	Retries                    types.Int64                                            `tfsdk:"retries"`
	Role                       types.String                                           `tfsdk:"role"`
	SharedSecret               types.String                                           `tfsdk:"shared_secret"`
	TimeoutSeconds             types.Int64                                            `tfsdk:"timeout_seconds"`
	EncryptionScheme           types.String                                           `tfsdk:"encryption_scheme"`
	MessageKey                 types.String                                           `tfsdk:"message_key"`
	EncryptionKey              types.String                                           `tfsdk:"encryption_key"`
	ExternalCiscoIseIpAddrDtos []AuthenticationPolicyServerExternalCiscoIseIpAddrDtos `tfsdk:"external_cisco_ise_ip_addr_dtos"`
}

type AuthenticationPolicyServerCiscoIseDtos struct {
	Description    types.String `tfsdk:"description"`
	Fqdn           types.String `tfsdk:"fqdn"`
	Password       types.String `tfsdk:"password"`
	Sshkey         types.String `tfsdk:"sshkey"`
	IpAddress      types.String `tfsdk:"ip_address"`
	SubscriberName types.String `tfsdk:"subscriber_name"`
	UserName       types.String `tfsdk:"user_name"`
}

type AuthenticationPolicyServerExternalCiscoIseIpAddrDtos struct {
	ExternalCiscoIseIpAddresses []AuthenticationPolicyServerExternalCiscoIseIpAddrDtosExternalCiscoIseIpAddresses `tfsdk:"external_cisco_ise_ip_addresses"`
	Type                        types.String                                                                      `tfsdk:"type"`
}

type AuthenticationPolicyServerExternalCiscoIseIpAddrDtosExternalCiscoIseIpAddresses struct {
	ExternalIpAddress types.String `tfsdk:"external_ip_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AuthenticationPolicyServer) getPath() string {
	return "/dna/intent/api/v1/authentication-policy-servers"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AuthenticationPolicyServer) toBody(ctx context.Context, state AuthenticationPolicyServer) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.AuthenticationPort.IsNull() {
		body, _ = sjson.Set(body, "authenticationPort", data.AuthenticationPort.ValueInt64())
	}
	if !data.AccountingPort.IsNull() {
		body, _ = sjson.Set(body, "accountingPort", data.AccountingPort.ValueInt64())
	}
	if len(data.CiscoIseDtos) > 0 {
		body, _ = sjson.Set(body, "ciscoIseDtos", []interface{}{})
		for _, item := range data.CiscoIseDtos {
			itemBody := ""
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Fqdn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fqdn", item.Fqdn.ValueString())
			}
			if !item.Password.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "password", item.Password.ValueString())
			}
			if !item.Sshkey.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sshkey", item.Sshkey.ValueString())
			}
			if !item.IpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress", item.IpAddress.ValueString())
			}
			if !item.SubscriberName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "subscriberName", item.SubscriberName.ValueString())
			}
			if !item.UserName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "userName", item.UserName.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ciscoIseDtos.-1", itemBody)
		}
	}
	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, "ipAddress", data.IpAddress.ValueString())
	}
	if !data.PxgridEnabled.IsNull() {
		body, _ = sjson.Set(body, "pxgridEnabled", data.PxgridEnabled.ValueBool())
	}
	if !data.UseDnacCertForPxgrid.IsNull() {
		body, _ = sjson.Set(body, "useDnacCertForPxgrid", data.UseDnacCertForPxgrid.ValueBool())
	}
	if !data.IsIseEnabled.IsNull() {
		body, _ = sjson.Set(body, "isIseEnabled", data.IsIseEnabled.ValueBool())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, "port", data.Port.ValueInt64())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "protocol", data.Protocol.ValueString())
	}
	if !data.Retries.IsNull() {
		body, _ = sjson.Set(body, "retries", data.Retries.ValueInt64())
	}
	if !data.Role.IsNull() {
		body, _ = sjson.Set(body, "role", data.Role.ValueString())
	}
	if !data.SharedSecret.IsNull() {
		body, _ = sjson.Set(body, "sharedSecret", data.SharedSecret.ValueString())
	}
	if !data.TimeoutSeconds.IsNull() {
		body, _ = sjson.Set(body, "timeoutSeconds", data.TimeoutSeconds.ValueInt64())
	}
	if !data.EncryptionScheme.IsNull() {
		body, _ = sjson.Set(body, "encryptionScheme", data.EncryptionScheme.ValueString())
	}
	if !data.MessageKey.IsNull() {
		body, _ = sjson.Set(body, "messageKey", data.MessageKey.ValueString())
	}
	if !data.EncryptionKey.IsNull() {
		body, _ = sjson.Set(body, "encryptionKey", data.EncryptionKey.ValueString())
	}
	if len(data.ExternalCiscoIseIpAddrDtos) > 0 {
		body, _ = sjson.Set(body, "externalCiscoIseIpAddrDtos", []interface{}{})
		for _, item := range data.ExternalCiscoIseIpAddrDtos {
			itemBody := ""
			if len(item.ExternalCiscoIseIpAddresses) > 0 {
				itemBody, _ = sjson.Set(itemBody, "externalCiscoIseIpAddresses", []interface{}{})
				for _, childItem := range item.ExternalCiscoIseIpAddresses {
					itemChildBody := ""
					if !childItem.ExternalIpAddress.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "externalIpAddress", childItem.ExternalIpAddress.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "externalCiscoIseIpAddresses.-1", itemChildBody)
				}
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "externalCiscoIseIpAddrDtos.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AuthenticationPolicyServer) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("authenticationPort"); value.Exists() {
		data.AuthenticationPort = types.Int64Value(value.Int())
	} else {
		data.AuthenticationPort = types.Int64Null()
	}
	if value := res.Get("accountingPort"); value.Exists() {
		data.AccountingPort = types.Int64Value(value.Int())
	} else {
		data.AccountingPort = types.Int64Null()
	}
	if value := res.Get("ciscoIseDtos"); value.Exists() && len(value.Array()) > 0 {
		data.CiscoIseDtos = make([]AuthenticationPolicyServerCiscoIseDtos, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AuthenticationPolicyServerCiscoIseDtos{}
			if cValue := v.Get("description"); cValue.Exists() {
				item.Description = types.StringValue(cValue.String())
			} else {
				item.Description = types.StringNull()
			}
			if cValue := v.Get("fqdn"); cValue.Exists() {
				item.Fqdn = types.StringValue(cValue.String())
			} else {
				item.Fqdn = types.StringNull()
			}
			if cValue := v.Get("password"); cValue.Exists() {
				item.Password = types.StringValue(cValue.String())
			} else {
				item.Password = types.StringNull()
			}
			if cValue := v.Get("sshkey"); cValue.Exists() {
				item.Sshkey = types.StringValue(cValue.String())
			} else {
				item.Sshkey = types.StringNull()
			}
			if cValue := v.Get("ipAddress"); cValue.Exists() {
				item.IpAddress = types.StringValue(cValue.String())
			} else {
				item.IpAddress = types.StringNull()
			}
			if cValue := v.Get("subscriberName"); cValue.Exists() {
				item.SubscriberName = types.StringValue(cValue.String())
			} else {
				item.SubscriberName = types.StringNull()
			}
			if cValue := v.Get("userName"); cValue.Exists() {
				item.UserName = types.StringValue(cValue.String())
			} else {
				item.UserName = types.StringNull()
			}
			data.CiscoIseDtos = append(data.CiscoIseDtos, item)
			return true
		})
	}
	if value := res.Get("ipAddress"); value.Exists() {
		data.IpAddress = types.StringValue(value.String())
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get("pxgridEnabled"); value.Exists() {
		data.PxgridEnabled = types.BoolValue(value.Bool())
	} else {
		data.PxgridEnabled = types.BoolNull()
	}
	if value := res.Get("useDnacCertForPxgrid"); value.Exists() {
		data.UseDnacCertForPxgrid = types.BoolValue(value.Bool())
	} else {
		data.UseDnacCertForPxgrid = types.BoolNull()
	}
	if value := res.Get("isIseEnabled"); value.Exists() {
		data.IsIseEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsIseEnabled = types.BoolNull()
	}
	if value := res.Get("port"); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get("protocol"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("retries"); value.Exists() {
		data.Retries = types.Int64Value(value.Int())
	} else {
		data.Retries = types.Int64Null()
	}
	if value := res.Get("role"); value.Exists() {
		data.Role = types.StringValue(value.String())
	} else {
		data.Role = types.StringNull()
	}
	if value := res.Get("timeoutSeconds"); value.Exists() {
		data.TimeoutSeconds = types.Int64Value(value.Int())
	} else {
		data.TimeoutSeconds = types.Int64Null()
	}
	if value := res.Get("encryptionScheme"); value.Exists() {
		data.EncryptionScheme = types.StringValue(value.String())
	} else {
		data.EncryptionScheme = types.StringNull()
	}
	if value := res.Get("externalCiscoIseIpAddrDtos"); value.Exists() && len(value.Array()) > 0 {
		data.ExternalCiscoIseIpAddrDtos = make([]AuthenticationPolicyServerExternalCiscoIseIpAddrDtos, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AuthenticationPolicyServerExternalCiscoIseIpAddrDtos{}
			if cValue := v.Get("externalCiscoIseIpAddresses"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ExternalCiscoIseIpAddresses = make([]AuthenticationPolicyServerExternalCiscoIseIpAddrDtosExternalCiscoIseIpAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AuthenticationPolicyServerExternalCiscoIseIpAddrDtosExternalCiscoIseIpAddresses{}
					if ccValue := cv.Get("externalIpAddress"); ccValue.Exists() {
						cItem.ExternalIpAddress = types.StringValue(ccValue.String())
					} else {
						cItem.ExternalIpAddress = types.StringNull()
					}
					item.ExternalCiscoIseIpAddresses = append(item.ExternalCiscoIseIpAddresses, cItem)
					return true
				})
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			data.ExternalCiscoIseIpAddrDtos = append(data.ExternalCiscoIseIpAddrDtos, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AuthenticationPolicyServer) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("authenticationPort"); value.Exists() && !data.AuthenticationPort.IsNull() {
		data.AuthenticationPort = types.Int64Value(value.Int())
	} else {
		data.AuthenticationPort = types.Int64Null()
	}
	if value := res.Get("accountingPort"); value.Exists() && !data.AccountingPort.IsNull() {
		data.AccountingPort = types.Int64Value(value.Int())
	} else {
		data.AccountingPort = types.Int64Null()
	}
	for i := range data.CiscoIseDtos {
		keys := [...]string{"description", "fqdn", "password", "sshkey", "ipAddress", "subscriberName", "userName"}
		keyValues := [...]string{data.CiscoIseDtos[i].Description.ValueString(), data.CiscoIseDtos[i].Fqdn.ValueString(), data.CiscoIseDtos[i].Password.ValueString(), data.CiscoIseDtos[i].Sshkey.ValueString(), data.CiscoIseDtos[i].IpAddress.ValueString(), data.CiscoIseDtos[i].SubscriberName.ValueString(), data.CiscoIseDtos[i].UserName.ValueString()}

		var r gjson.Result
		res.Get("ciscoIseDtos").ForEach(
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
		if value := r.Get("description"); value.Exists() && !data.CiscoIseDtos[i].Description.IsNull() {
			data.CiscoIseDtos[i].Description = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].Description = types.StringNull()
		}
		if value := r.Get("fqdn"); value.Exists() && !data.CiscoIseDtos[i].Fqdn.IsNull() {
			data.CiscoIseDtos[i].Fqdn = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].Fqdn = types.StringNull()
		}
		if value := r.Get("password"); value.Exists() && !data.CiscoIseDtos[i].Password.IsNull() {
			data.CiscoIseDtos[i].Password = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].Password = types.StringNull()
		}
		if value := r.Get("sshkey"); value.Exists() && !data.CiscoIseDtos[i].Sshkey.IsNull() {
			data.CiscoIseDtos[i].Sshkey = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].Sshkey = types.StringNull()
		}
		if value := r.Get("ipAddress"); value.Exists() && !data.CiscoIseDtos[i].IpAddress.IsNull() {
			data.CiscoIseDtos[i].IpAddress = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].IpAddress = types.StringNull()
		}
		if value := r.Get("subscriberName"); value.Exists() && !data.CiscoIseDtos[i].SubscriberName.IsNull() {
			data.CiscoIseDtos[i].SubscriberName = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].SubscriberName = types.StringNull()
		}
		if value := r.Get("userName"); value.Exists() && !data.CiscoIseDtos[i].UserName.IsNull() {
			data.CiscoIseDtos[i].UserName = types.StringValue(value.String())
		} else {
			data.CiscoIseDtos[i].UserName = types.StringNull()
		}
	}
	if value := res.Get("ipAddress"); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(value.String())
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get("pxgridEnabled"); value.Exists() && !data.PxgridEnabled.IsNull() {
		data.PxgridEnabled = types.BoolValue(value.Bool())
	} else {
		data.PxgridEnabled = types.BoolNull()
	}
	if value := res.Get("useDnacCertForPxgrid"); value.Exists() && !data.UseDnacCertForPxgrid.IsNull() {
		data.UseDnacCertForPxgrid = types.BoolValue(value.Bool())
	} else {
		data.UseDnacCertForPxgrid = types.BoolNull()
	}
	if value := res.Get("isIseEnabled"); value.Exists() && !data.IsIseEnabled.IsNull() {
		data.IsIseEnabled = types.BoolValue(value.Bool())
	} else {
		data.IsIseEnabled = types.BoolNull()
	}
	if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("retries"); value.Exists() && !data.Retries.IsNull() {
		data.Retries = types.Int64Value(value.Int())
	} else {
		data.Retries = types.Int64Null()
	}
	if value := res.Get("role"); value.Exists() && !data.Role.IsNull() {
		data.Role = types.StringValue(value.String())
	} else {
		data.Role = types.StringNull()
	}
	if value := res.Get("timeoutSeconds"); value.Exists() && !data.TimeoutSeconds.IsNull() {
		data.TimeoutSeconds = types.Int64Value(value.Int())
	} else {
		data.TimeoutSeconds = types.Int64Null()
	}
	if value := res.Get("encryptionScheme"); value.Exists() && !data.EncryptionScheme.IsNull() {
		data.EncryptionScheme = types.StringValue(value.String())
	} else {
		data.EncryptionScheme = types.StringNull()
	}
	for i := range data.ExternalCiscoIseIpAddrDtos {
		keys := [...]string{"type"}
		keyValues := [...]string{data.ExternalCiscoIseIpAddrDtos[i].Type.ValueString()}

		var r gjson.Result
		res.Get("externalCiscoIseIpAddrDtos").ForEach(
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
		for ci := range data.ExternalCiscoIseIpAddrDtos[i].ExternalCiscoIseIpAddresses {
			keys := [...]string{"externalIpAddress"}
			keyValues := [...]string{data.ExternalCiscoIseIpAddrDtos[i].ExternalCiscoIseIpAddresses[ci].ExternalIpAddress.ValueString()}

			var cr gjson.Result
			r.Get("externalCiscoIseIpAddresses").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("externalIpAddress"); value.Exists() && !data.ExternalCiscoIseIpAddrDtos[i].ExternalCiscoIseIpAddresses[ci].ExternalIpAddress.IsNull() {
				data.ExternalCiscoIseIpAddrDtos[i].ExternalCiscoIseIpAddresses[ci].ExternalIpAddress = types.StringValue(value.String())
			} else {
				data.ExternalCiscoIseIpAddrDtos[i].ExternalCiscoIseIpAddresses[ci].ExternalIpAddress = types.StringNull()
			}
		}
		if value := r.Get("type"); value.Exists() && !data.ExternalCiscoIseIpAddrDtos[i].Type.IsNull() {
			data.ExternalCiscoIseIpAddrDtos[i].Type = types.StringValue(value.String())
		} else {
			data.ExternalCiscoIseIpAddrDtos[i].Type = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AuthenticationPolicyServer) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.AuthenticationPort.IsNull() {
		return false
	}
	if !data.AccountingPort.IsNull() {
		return false
	}
	if len(data.CiscoIseDtos) > 0 {
		return false
	}
	if !data.IpAddress.IsNull() {
		return false
	}
	if !data.PxgridEnabled.IsNull() {
		return false
	}
	if !data.UseDnacCertForPxgrid.IsNull() {
		return false
	}
	if !data.IsIseEnabled.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Protocol.IsNull() {
		return false
	}
	if !data.Retries.IsNull() {
		return false
	}
	if !data.Role.IsNull() {
		return false
	}
	if !data.SharedSecret.IsNull() {
		return false
	}
	if !data.TimeoutSeconds.IsNull() {
		return false
	}
	if !data.EncryptionScheme.IsNull() {
		return false
	}
	if !data.MessageKey.IsNull() {
		return false
	}
	if !data.EncryptionKey.IsNull() {
		return false
	}
	if len(data.ExternalCiscoIseIpAddrDtos) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
