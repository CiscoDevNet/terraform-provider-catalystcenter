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

import (
	"context"
	"strconv"
	"strings"

	cc "github.com/netascode/go-catalystcenter"
)

// paginatePageLimit is the page size requested from paginated "get all" list
// endpoints. Catalyst Center list endpoints accept a limit of up to 500, which
// also matches the go-catalystcenter client's internal page-full threshold.
const paginatePageLimit = 500

// getAllPagesByID issues a "get all" request with an explicit limit so the
// go-catalystcenter client's built-in offset pagination retrieves and merges
// every page.
//
// The client (cc.Client.Get) already walks pages: it appends an increasing
// "offset" query parameter and concatenates the "response" arrays, continuing
// only while a page comes back completely full (exactly 500 items, the client's
// maxItems). Several list endpoints (notably the /dna/intent/api/v1/global-pool
// fallback) return a small default page (25 items) when no "limit" is supplied,
// so the client sees a non-full first page, stops, and objects beyond it are
// invisible — which caused reads/imports to fail with
// "Cannot import non-existent remote object" once the object count exceeded the
// default page size.
//
// Requesting limit=500 makes those endpoints return full pages so the client's
// own pagination kicks in and walks them all. We intentionally do NOT add an
// explicit "offset": the client owns the offset parameter, and supplying our own
// would collide with the one it appends for page 2+ (producing a duplicate
// "offset" and an HTTP 400). This helper therefore only sets the limit and
// delegates the actual page-walking to the client.
func getAllPagesByID(_ context.Context, client *cc.Client, path string) (cc.Res, error) {
	sep := "?"
	if strings.Contains(path, "?") {
		sep = "&"
	}
	return client.Get(path + sep + "limit=" + strconv.Itoa(paginatePageLimit))
}
