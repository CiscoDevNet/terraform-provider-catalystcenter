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

// This file is manually maintained (not generated). It implements the
// provider_meta pathway that lets a Terraform *module* identify itself to
// Catalyst Center telemetry, distinguishing module-driven traffic from bare
// provider usage.

import (
	"context"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cc "github.com/netascode/go-catalystcenter"
)

var uaOnce sync.Once

type ccProviderMeta struct {
	ModuleName types.String `tfsdk:"module_name"`
}

func applyProviderMeta(client *cc.Client, ctx context.Context, meta tfsdk.Config) {
	// meta.Raw is null when the module declares no provider_meta block for us.
	if meta.Raw.IsNull() {
		return
	}

	var m ccProviderMeta
	if diags := meta.Get(ctx, &m); diags.HasError() {
		// Never fail an operation over telemetry; just skip the marker.
		return
	}

	module := m.ModuleName.ValueString()
	if module == "" {
		return
	}

	uaOnce.Do(func() {
		if client == nil {
			return
		}
		ua := client.UserAgent
		if ua == "" {
			client.UserAgent = module
			return
		}
		const vendor = " Cisco"
		if strings.HasSuffix(ua, vendor) {
			client.UserAgent = ua[:len(ua)-len(vendor)] + " " + module + vendor
		} else {
			client.UserAgent = ua + " " + module
		}
	})
}
