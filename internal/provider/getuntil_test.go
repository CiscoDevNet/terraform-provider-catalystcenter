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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	cc "github.com/netascode/go-catalystcenter"
)

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func testClient(responses []string) *cc.Client {
	resp := make(chan string, len(responses)+1)
	resp <- `{"Token":"dummy"}`
	for _, v := range responses {
		resp <- v
	}
	close(resp)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, ok := <-resp
		if !ok {
			w.WriteHeader(400)
		}

		fmt.Fprintln(w, v)
	}))
	// defer ts.Close()

	res := must(cc.NewClient(ts.URL, "", ""))

	return &res
}

func TestGetUntil(t *testing.T) {
	type args struct {
		path       string
		jsonSearch string
		mods       []func(*cc.Req)
	}
	tests := []struct {
		name      string
		responses []string
		args      args
		want      string
		wantErr   bool
	}{{
		name:      "empty",
		responses: []string{`{"response":["x1"]}`},
		args: args{
			jsonSearch: "response",
		},
		want:    `["x1"]`,
		wantErr: false,
	}, {
		name:      "basic",
		responses: []string{`{"response":["x1","x2"]}`},
		args: args{
			jsonSearch: "response",
		},
		want:    `["x1","x2"]`,
		wantErr: false,
	}, {
		name:      "errNotArray",
		responses: []string{`{"response":"stringInsteadOfArray"}`},
		args: args{
			jsonSearch: "response.#(a==3)",
		},
		wantErr: true,
	}, {
		name:      "onePage",
		responses: []string{`{"response":[{"a":1}, {"a":2}, {"a":3}, {"a":4}]}`},
		args: args{
			jsonSearch: "response.#(a==3)",
		},
		want:    `{"a":3}`,
		wantErr: false,
	}, {
		name: "manyPages",
		responses: []string{
			`{"response":[
							{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":10},
							{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":20},
							{}, {}, {}, {}, {"a":25}
						]}`,
			`{"response":[
							{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":10},
							{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":20},
							{}, {}, {}, {}, {"a":25}
						]}`,
			`{"response":[{"a":1}, {"a":2}, {"a":3}, {"a":4}]}`,
		},
		args: args{
			jsonSearch: "response.#(a==3)",
		},
		want:    `{"a":3}`,
		wantErr: false,
	}, {
		name: "onePageNotMore",
		responses: []string{
			`{"response":[
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":10},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":20},
				{}, {}, {}, {"a":24}
			]}`,
			`{"response":[{"a":1}]}`,
		},
		args: args{
			jsonSearch: "response.#(a==1)",
		},
		want:    ``, // Second page should not be read when the first has 24 items.
		wantErr: false,
	}, {
		name: "twoPagesNotMore",
		responses: []string{
			`{"response":[
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":10},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":20},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":30},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":40},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":50},
			]}`,
			`{"response":[
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":10},
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {"a":20},
				{}, {}, {}, {}, {"a":25}
			]}`,
			`{"response":[{"a":1}]}`,
		},
		args: args{
			jsonSearch: "response.#(a==1)",
		},
		want:    ``, // Third page should not be read when the first has 50 items and second less than that.
		wantErr: false,
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotR, err := GetUntil(testClient(tt.responses), tt.args.path, tt.args.jsonSearch, tt.args.mods...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUntil() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := gotR.String()
			if got != tt.want {
				t.Errorf("GetUntil() = %v, want %v", got, tt.want)
			}
		})
	}
}
