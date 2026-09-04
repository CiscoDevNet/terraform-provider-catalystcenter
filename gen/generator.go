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

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath   = "./gen/definitions/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_catalystcenter_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_catalystcenter_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_catalystcenter_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_catalystcenter_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_catalystcenter_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/catalystcenter_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/catalystcenter_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/catalystcenter_",
		suffix: "/import.sh",
	},
}

type YamlConfig struct {
	Name                          string                `yaml:"name"`
	NoResource                    bool                  `yaml:"no_resource"`
	NoDataSource                  bool                  `yaml:"no_data_source"`
	RestEndpoint                  string                `yaml:"rest_endpoint"`
	GetRestEndpoint               string                `yaml:"get_rest_endpoint"`
	PostRestEndpoint              string                `yaml:"post_rest_endpoint"`
	PutRestEndpoint               string                `yaml:"put_rest_endpoint"`
	DeleteRestEndpoint            string                `yaml:"delete_rest_endpoint"`
	FallbackRestEndpoint          string                `yaml:"fallback_rest_endpoint"`
	GetNoId                       bool                  `yaml:"get_no_id"`
	GetFromAll                    bool                  `yaml:"get_from_all"`
	GetRequiresId                 bool                  `yaml:"get_requires_id"`
	GetExtraQueryParams           string                `yaml:"get_extra_query_params"`
	NoDelete                      bool                  `yaml:"no_delete"`
	DataSourceNoId                bool                  `yaml:"data_source_no_id"`
	DeleteNoId                    bool                  `yaml:"delete_no_id"`
	NoUpdate                      bool                  `yaml:"no_update"`
	NoRead                        bool                  `yaml:"no_read"`
	NoImport                      bool                  `yaml:"no_import"`
	ImportNoId                    bool                  `yaml:"import_no_id"`
	PostUpdate                    bool                  `yaml:"post_update"`
	PutCreate                     bool                  `yaml:"put_create"`
	PutDelete                     bool                  `yaml:"put_delete"`
	PostDelete                    bool                  `yaml:"post_delete"`
	UpdateComputed                bool                  `yaml:"update_computed"`
	DeleteIterative               bool                  `yaml:"delete_iterative"`
	RootList                      bool                  `yaml:"root_list"`
	NoReadPrefix                  bool                  `yaml:"no_read_prefix"`
	NoWait                        bool                  `yaml:"no_wait"`
	IdPath                        string                `yaml:"id_path"`
	IdFromQueryPath               string                `yaml:"id_from_query_path"`
	IdFromQueryPathAttribute      string                `yaml:"id_from_query_path_attribute"`
	IdQueryRestEndpoint           string                `yaml:"id_query_rest_endpoint"`
	IdQueryParam                  string                `yaml:"id_query_param"`
	IdFromAttribute               bool                  `yaml:"id_from_attribute"`
	DeviceUnreachabilityWarning   bool                  `yaml:"device_unreachability_warning"`
	AllowExistingOnCreate         bool                  `yaml:"allow_existing_on_create"`
	SkipUnmatchedRows             bool                  `yaml:"skip_unmatched_rows"`
	RemoveResourceOnEmptyResponse bool                  `yaml:"remove_resource_on_empty_response"`
	PutIdIncludePath              string                `yaml:"put_id_include_path"`
	PutIdQueryParam               string                `yaml:"put_id_query_param"`
	PutNoId                       bool                  `yaml:"put_no_id"`
	PutId                         bool                  `yaml:"put_id"`
	PutUpdateId                   bool                  `yaml:"put_update_id"`
	DeleteIdQueryParam            string                `yaml:"delete_id_query_param"`
	DeleteBody                    string                `yaml:"delete_body"`
	MinimumVersion                string                `yaml:"minimum_version"`
	DsDescription                 string                `yaml:"ds_description"`
	ResDescription                string                `yaml:"res_description"`
	DocCategory                   string                `yaml:"doc_category"`
	ExcludeTest                   bool                  `yaml:"exclude_test"`
	SkipMinimumTest               bool                  `yaml:"skip_minimum_test"`
	Attributes                    []YamlConfigAttribute `yaml:"attributes"`
	TestTags                      []string              `yaml:"test_tags"`
	TestPrerequisites             string                `yaml:"test_prerequisites"`
	MaxAsyncWaitTime              int64                 `yaml:"max_async_wait_time"`
	RetryOnErrorCodes             []string              `yaml:"retry_on_error_codes"`
	Mutex                         bool                  `yaml:"mutex"`
	UseCache                      bool                  `yaml:"use_cache"`
	CacheRestEndpoint             string                `yaml:"cache_rest_endpoint"`
	CacheFilterAttributes         []string              `yaml:"cache_filter_attributes"`
	SkipDeleteOnEmptyId           bool                  `yaml:"skip_delete_on_empty_id"`
	GetBeforeDelete               bool                  `yaml:"get_before_delete"`
}

type YamlConfigAttribute struct {
	ModelName                 string                `yaml:"model_name"`
	ResponseModelName         string                `yaml:"response_model_name"`
	FallbackResponseModelName string                `yaml:"fallback_response_model_name"`
	TfName                    string                `yaml:"tf_name"`
	Type                      string                `yaml:"type"`
	ElementType               string                `yaml:"element_type"`
	DataPath                  string                `yaml:"data_path"`
	ResponseDataPath          string                `yaml:"response_data_path"`
	FallbackResponseDataPath  string                `yaml:"fallback_response_data_path"`
	PutDataPath               string                `yaml:"put_data_path"`
	Id                        bool                  `yaml:"id"`
	MatchId                   bool                  `yaml:"match_id"`
	Reference                 bool                  `yaml:"reference"`
	RequiresReplace           bool                  `yaml:"requires_replace"`
	QueryParam                bool                  `yaml:"query_param"`
	DeleteQueryParam          bool                  `yaml:"delete_query_param"`
	GetQueryParam             bool                  `yaml:"get_query_param"`
	PutQueryParam             bool                  `yaml:"put_query_param"`
	PostQueryParam            bool                  `yaml:"post_query_param"`
	QueryParamName            string                `yaml:"query_param_name"`
	DeleteQueryParamName      string                `yaml:"delete_query_param_name"`
	GetQueryParamName         string                `yaml:"get_query_param_name"`
	PutQueryParamName         string                `yaml:"put_query_param_name"`
	PostQueryParamName        string                `yaml:"post_query_param_name"`
	CreateQueryPath           bool                  `yaml:"create_query_path"`
	DataSourceQuery           bool                  `yaml:"data_source_query"`
	QueryParamNoBody          bool                  `yaml:"query_param_no_body"`
	Mandatory                 bool                  `yaml:"mandatory"`
	Computed                  bool                  `yaml:"computed"`
	ComputedRefreshValue      bool                  `yaml:"computed_refresh_value"`
	NoUseStateForUnknown      bool                  `yaml:"no_use_state_for_unknown"`
	WriteOnly                 bool                  `yaml:"write_only"`
	WriteOnlyTF               bool                  `yaml:"write_only_tf"`
	WoVersion                 bool                  `yaml:"-"` // Internal: marks a generated "<attr>_wo_version" companion attribute (state-only rotation trigger)
	CoexistingSecret          bool                  `yaml:"-"` // Internal: marks the legacy state-storing twin of a "<attr>_wo" write-only attribute
	WoBaseName                string                `yaml:"-"` // Internal: on a "<attr>_wo" attribute, the name of its legacy twin
	MutualExclusivityNote     string                `yaml:"-"` // Internal: documentation note carried by both halves of a write_only_tf pair
	WoPairMandatory           bool                  `yaml:"-"` // Internal: on a legacy twin, whether the pair must supply the secret through one of its halves
	WoPairHasVersion          bool                  `yaml:"-"` // Internal: on a legacy twin, whether a "_wo_version" companion was generated
	CoexistenceNote           string                `yaml:"coexistence_note"`
	ExcludeFromPut            bool                  `yaml:"exclude_from_put"`
	ExcludeTest               bool                  `yaml:"exclude_test"`
	ExcludeExample            bool                  `yaml:"exclude_example"`
	Description               string                `yaml:"description"`
	Example                   string                `yaml:"example"`
	EnumValues                []string              `yaml:"enum_values"`
	MinList                   int64                 `yaml:"min_list"`
	MaxList                   int64                 `yaml:"max_list"`
	MinInt                    int64                 `yaml:"min_int"`
	MaxInt                    int64                 `yaml:"max_int"`
	MinFloat                  float64               `yaml:"min_float"`
	MaxFloat                  float64               `yaml:"max_float"`
	DecimalPlaces             int64                 `yaml:"decimal_places"`
	StringPatterns            []string              `yaml:"string_patterns"`
	StringMinLength           int64                 `yaml:"string_min_length"`
	StringMaxLength           int64                 `yaml:"string_max_length"`
	DefaultValue              string                `yaml:"default_value"`
	Value                     string                `yaml:"value"`
	ValueCondition            string                `yaml:"value_condition"`
	TestValue                 string                `yaml:"test_value"`
	MinimumTestValue          string                `yaml:"minimum_test_value"`
	TestTags                  []string              `yaml:"test_tags"`
	Attributes                []YamlConfigAttribute `yaml:"attributes"`
	MaxElementsInRootList     int64                 `yaml:"max_elements_in_root_list"`
	MapKeyExample             string                `yaml:"map_key_example"`
	KeyPart                   bool                  `yaml:"key_part"`
	NoPut                     bool                  `yaml:"no_put"`
	NullOnEmpty               bool                  `yaml:"null_on_empty"`
	CustomModifier            string                `yaml:"custom_modifier"`
	AlwaysInclude             bool                  `yaml:"always_include"`
	GetIfUnsetOnUpdate        bool                  `yaml:"get_if_unset_on_update"`
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to build a SJSON path
func BuildPath(s []string) string {
	return strings.Join(s, ".")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id {
			return true
		}
	}
	return false
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return true if query parameters included in attributes
func HasQueryParam(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.QueryParam {
			return true
		}
	}
	return false
}

// Templating helper function to return true if get query parameters included in attributes
func HasGetQueryParam(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.GetQueryParam {
			return true
		}
	}
	return false
}

// Templating helper function to return true if delete query parameters included in attributes
func HasDeleteQueryParam(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.DeleteQueryParam {
			return true
		}
	}
	return false
}

// Templating helper function to return true if any attribute is set as compute_refresh_value
func HasComputedRefreshValue(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.ComputedRefreshValue {
			return true
		}
		if len(attr.Attributes) > 0 {
			if HasComputedRefreshValue(attr.Attributes) {
				return true
			}
		}
	}
	return false
}

// Templating helper function to return true if any computed_refresh_value attribute participates in key matching
func HasComputedRefreshValueInKeys(attributes []YamlConfigAttribute) bool {
	noId := !HasId(attributes)
	for _, attr := range attributes {
		if attr.ComputedRefreshValue && (attr.Id || noId) {
			return true
		}
	}
	return false
}

// Templating helper function to return true if any child attribute is handled by fromBodyUnknowns.
// A nested list/set/map whose children are all plain configurable attributes contributes no
// statements there, so its lookup loop must be skipped entirely - otherwise the generated code
// declares the `r` gjson.Result without ever using it and does not compile.
func HasUnknownsChildren(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Value == "" && !attr.WriteOnly && !attr.Reference && attr.Computed {
			return true
		}
	}
	return false
}

// GetIfUnsetOnUpdateAttributes returns the top-level attributes flagged get_if_unset_on_update.
// These are values that Catalyst Center assigns out-of-band (for example an SDA anycast
// gateway) and that are not part of the data model, so they are null in the plan. The Update
// template uses this list to fetch each such value from the controller and include it in the
// PUT body whenever the plan leaves it empty, so a full-object replace PUT does not strip it.
func GetIfUnsetOnUpdateAttributes(config YamlConfig) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.GetIfUnsetOnUpdate {
			r = append(r, attr)
		}
	}
	return r
}

// Templating helper function to return the ID attribute
func GetId(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.Id {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return the "match_id" attribute
func GetMatchId(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.MatchId {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to build a GJSON filter expression for cache lookups.
// Given cache_filter_attributes [siteId, networkDeviceId] it produces Go code like:
//
//	"response.#(siteId==\"" + state.SiteId.ValueString() + "\")#|#(networkDeviceId==\"" + state.NetworkDeviceId.ValueString() + "\")"
//
// For a single attribute it produces:
//
//	"response.#(siteId==\"" + state.SiteId.ValueString() + "\")"
func CacheFilterGjsonExpr(attributes []YamlConfigAttribute, filterNames []string) string {
	if len(filterNames) == 0 {
		return ""
	}
	// Build a lookup of model_name -> tf_name from attributes
	tfNameByModel := make(map[string]string)
	for _, attr := range attributes {
		tfNameByModel[attr.ModelName] = attr.TfName
	}

	// Each part produces: modelName==\"" + state.GoName.ValueString() + "\"
	parts := make([]string, 0, len(filterNames))
	for _, modelName := range filterNames {
		tfName := tfNameByModel[modelName]
		goName := ToGoName(tfName)
		parts = append(parts, modelName+`==\"" + state.`+goName+`.ValueString() + "\"`)
	}

	if len(parts) == 1 {
		return `"response.#(` + parts[0] + `)"`
	}
	// Chain: response.#(first=="X")#|#(second=="Y")
	expr := `"response.#(` + parts[0] + `)#`
	for _, p := range parts[1:] {
		expr += `|#(` + p + `)#`
	}
	// Remove trailing # from last filter to get first match
	expr = expr[:len(expr)-1]
	expr += `"`
	return expr
}

// Templating helper function to return the query parameter attribute
func GetQueryParam(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.QueryParam {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return the delete query parameter attribute
func GetDeleteQueryParam(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.DeleteQueryParam {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return true if data source query attribute(s) are present
func HasDataSourceQuery(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.DataSourceQuery {
			return true
		}
	}
	return false
}

// Templating helper function to return the first path element
func FirstPathElement(path, getFromAllPath string) string {
	if getFromAllPath != "" {
		return getFromAllPath
	} else if strings.HasPrefix(path, "response.") {
		return strings.Split(path, ".")[0]
	}
	return ""
}

// Templating helper function to return the second and subsequent path elements
func RemainingPathElements(path, getFromAllPath string) string {
	if getFromAllPath != "" {
		return path
	} else if strings.HasPrefix(path, "response.0.") {
		return strings.Join(strings.Split(path, ".")[2:], ".")
	} else if strings.HasPrefix(path, "response.") {
		return strings.Join(strings.Split(path, ".")[1:], ".")
	}
	return path
}

// Templating helper function to return the query path in case of "get_from_all" being enabled
func GetFromAllPath(config YamlConfig) string {
	if config.GetFromAll {
		if config.IdFromQueryPath != "" {
			return config.IdFromQueryPath
		}
	}
	return ""
}

// Templating helper function to return true if type is a list or set without nested elements
func IsListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list without nested elements
func IsList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set without nested elements
func IsSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of strings without nested elements
func IsStringListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "String" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of integers without nested elements
func IsInt64ListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "Int64" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set with nested elements
func IsNestedListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list with nested elements
func IsNestedList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set with nested elements
func IsNestedSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a map with nested elements
func IsNestedMap(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Map" && attribute.ElementType == "" && len(attribute.Attributes) > 0 {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list, set or map with nested elements
func IsNestedListSetMap(attribute YamlConfigAttribute) bool {
	return IsNestedListSet(attribute) || IsNestedMap(attribute)
}

// Templating helper function to check if any child attribute has key_part set
func HasKeyPart(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.KeyPart {
			return true
		}
	}
	return false
}

// Templating helper function to return child attributes with key_part set
func GetKeyPartAttributes(attributes []YamlConfigAttribute) []YamlConfigAttribute {
	var result []YamlConfigAttribute
	for _, attr := range attributes {
		if attr.KeyPart {
			result = append(result, attr)
		}
	}
	return result
}

// Templating helper function to return true if create query path included in attributes
func HasCreateQueryPath(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.CreateQueryPath {
			return true
		}
	}
	return false
}

// Templating helper function to return the create query path attribute
func GetCreateQueryPath(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.CreateQueryPath {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return a query parameter string based on the HTTP method input source (plan, state) and provided attributes.
// By default, it uses attr.QueryParam if specified, and for method-specific parameters like DeleteQueryParamName, GetQueryParamName, etc.,
// it uses those if available for the respective HTTP method. If no specific query parameter is provided for a method, it defaults to attr.ModelName.
// Returns the constructed query parameter string.
func GenerateQueryParamString(method string, inputSource string, attributes []YamlConfigAttribute) string {
	var params []string
	first := true

	for _, attr := range attributes {
		var queryParamName string
		includeParam := false

		// Determine the appropriate query parameter name based on the method
		switch method {
		case "DELETE":
			if attr.DeleteQueryParam {
				queryParamName = attr.DeleteQueryParamName
				includeParam = true
			}
		case "GET":
			if attr.GetQueryParam {
				queryParamName = attr.GetQueryParamName
				includeParam = true
			}
		case "POST":
			if attr.PostQueryParam {
				queryParamName = attr.PostQueryParamName
				includeParam = true
			}
		case "PUT":
			if attr.PutQueryParam {
				queryParamName = attr.PutQueryParamName
				includeParam = true
			}
		}

		// If no method-specific query parameter is set, fall back to default query parameter
		if !includeParam && attr.QueryParam {
			queryParamName = attr.QueryParamName
			includeParam = true
		}

		// Use model name if queryParamName is still empty
		if queryParamName == "" {
			queryParamName = attr.ModelName
		}

		// Construct the query parameter string if includeParam is true
		if includeParam {
			var valueExpr string
			switch attr.Type {
			case "Int64":
				valueExpr = `strconv.FormatInt(` + inputSource + `.` + ToGoName(attr.TfName) + `.Value` + attr.Type + `(), 10)`
			case "Bool":
				valueExpr = `strconv.FormatBool(` + inputSource + `.` + ToGoName(attr.TfName) + `.Value` + attr.Type + `())`
			default:
				valueExpr = inputSource + `.` + ToGoName(attr.TfName) + `.Value` + attr.Type + `()`
			}

			if first {
				params = append(params, `"?`+queryParamName+`=" + url.QueryEscape(`+valueExpr+`)`)
				first = false
			} else {
				params = append(params, `"&`+queryParamName+`=" + url.QueryEscape(`+valueExpr+`)`)
			}
		}
	}

	// Return the appropriate string based on whether params is empty or not
	if len(params) == 0 {
		return ""
	} else {
		return strings.Join(params, "+")
	}
}

func GenerateDeleteOnlyQueryParamString(inputSource string, attributes []YamlConfigAttribute) string {
	var params []string
	first := true

	for _, attr := range attributes {
		if !attr.DeleteQueryParam {
			continue
		}

		queryParamName := attr.DeleteQueryParamName
		if queryParamName == "" {
			queryParamName = attr.ModelName
		}

		var valueExpr string
		switch attr.Type {
		case "Int64":
			valueExpr = `strconv.FormatInt(` + inputSource + `.` + ToGoName(attr.TfName) + `.Value` + attr.Type + `(), 10)`
		case "Bool":
			valueExpr = `strconv.FormatBool(` + inputSource + `.` + ToGoName(attr.TfName) + `.Value` + attr.Type + `())`
		default:
			valueExpr = inputSource + `.` + ToGoName(attr.TfName) + `.Value` + attr.Type + `()`
		}

		if first {
			params = append(params, `"?`+queryParamName+`=" + url.QueryEscape(`+valueExpr+`)`)
			first = false
		} else {
			params = append(params, `"&`+queryParamName+`=" + url.QueryEscape(`+valueExpr+`)`)
		}
	}

	if len(params) == 0 {
		return ""
	}
	return strings.Join(params, "+")
}

// Templating helper function to return a list of import attributes
func ImportAttributes(config YamlConfig) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.Reference || attr.QueryParam || attr.GetQueryParam || attr.Id {
			r = append(r, attr)
		}
	}
	if !config.IdFromAttribute && !config.ImportNoId {
		attr := YamlConfigAttribute{}
		attr.ModelName = "id"
		attr.TfName = "id"
		attr.Type = "String"
		attr.Example = "4b0b7a80-44c0-4bf2-bab5-fc24b4e0a17e"
		r = append(r, attr)
	}
	return r
}

// Templating helper function to build the Float64 read expression for a gjson value.
// varName is the gjson result variable at the render site (e.g. "value", "cValue", "fv").
// When the attribute sets decimal_places, the value is rounded to that many decimals to
// prevent precision drift (e.g. CatC returning 5-decimal lat/long against a 6-decimal
// config). Otherwise it returns the plain .Float() call.
func FloatReadExpr(varName string, attr YamlConfigAttribute) string {
	if attr.DecimalPlaces > 0 {
		factor := "1"
		for i := int64(0); i < attr.DecimalPlaces; i++ {
			factor += "0"
		}
		return fmt.Sprintf("math.Round(%s.Float()*%s) / %s", varName, factor, factor)
	}
	return varName + ".Float()"
}

// Templating helper function to subtract one number from another
func Subtract(a, b int) int {
	return a - b
}

// HasWriteOnlyTFChildren reports whether a list/set attribute has at least one direct
// child flagged write_only_tf (after augmentation, these children are the renamed "_wo"
// leaves — detectable via the still-set WriteOnlyTF flag, which augmentWriteOnlyTF
// preserves on the renamed attribute). Used by resource.go to decide whether to emit the
// "read the whole parent list from config" block for nested write-only secrets.
func HasWriteOnlyTFChildren(attr YamlConfigAttribute) bool {
	for _, child := range attr.Attributes {
		if child.WriteOnlyTF {
			return true
		}
	}
	return false
}

// WriteOnlyTFChildren returns the direct children of a list/set attribute that are flagged
// write_only_tf (the renamed "_wo" leaves). Used to copy each write-only leaf from the
// config-read temp back into the plan element in the generated nested config-read block.
func WriteOnlyTFChildren(attr YamlConfigAttribute) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, child := range attr.Attributes {
		if child.WriteOnlyTF {
			r = append(r, child)
		}
	}
	return r
}

// CoexistingSecretAttributes returns the top-level legacy twins of write-only
// secrets, i.e. the attributes kept for backwards compatibility when their "_wo" variant
// was generated. Nested secrets are not included: they would need a per-element walk of
// the enclosing list, and this coexistence note is currently only surfaced in the
// documentation.
func CoexistingSecretAttributes(config YamlConfig) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.CoexistingSecret {
			r = append(r, attr)
		}
	}
	return r
}

// CoexistingSecretChildren returns the legacy twins of write-only secrets nested
// directly inside a list/set attribute. ValidateConfig applies the same pair checks to
// them, once per list element.
func CoexistingSecretChildren(attr YamlConfigAttribute) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, child := range attr.Attributes {
		if child.CoexistingSecret {
			r = append(r, child)
		}
	}
	return r
}

// CoexistingSecretParentLists returns the top-level list/set attributes holding at least
// one legacy twin of a write-only secret. ValidateConfig reads each such list from the
// configuration once and walks its elements.
func CoexistingSecretParentLists(config YamlConfig) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if len(CoexistingSecretChildren(attr)) > 0 {
			r = append(r, attr)
		}
	}
	return r
}

// WriteOnlyTFParentLists returns the top-level list/set attributes that contain at least
// one write_only_tf child. Emitted once per parent list in the Create/Update config-read
// codegen.
func WriteOnlyTFParentLists(config YamlConfig) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if HasWriteOnlyTFChildren(attr) {
			r = append(r, attr)
		}
	}
	return r
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                           ToGoName,
	"camelCase":                          CamelCase,
	"strContains":                        strings.Contains,
	"snakeCase":                          SnakeCase,
	"sprintf":                            fmt.Sprintf,
	"toLower":                            strings.ToLower,
	"path":                               BuildPath,
	"hasId":                              HasId,
	"hasReference":                       HasReference,
	"hasQueryParam":                      HasQueryParam,
	"hasGetQueryParam":                   HasGetQueryParam,
	"hasDeleteQueryParam":                HasDeleteQueryParam,
	"generateQueryParamString":           GenerateQueryParamString,
	"generateDeleteOnlyQueryParamString": GenerateDeleteOnlyQueryParamString,
	"getId":                              GetId,
	"getMatchId":                         GetMatchId,
	"cacheFilterGjsonExpr":               CacheFilterGjsonExpr,
	"hasCreateQueryPath":                 HasCreateQueryPath,
	"getCreateQueryPath":                 GetCreateQueryPath,
	"getQueryParam":                      GetQueryParam,
	"getDeleteQueryParam":                GetDeleteQueryParam,
	"hasDataSourceQuery":                 HasDataSourceQuery,
	"hasComputedRefreshValue":            HasComputedRefreshValue,
	"hasComputedRefreshValueInKeys":      HasComputedRefreshValueInKeys,
	"hasUnknownsChildren":                HasUnknownsChildren,
	"firstPathElement":                   FirstPathElement,
	"remainingPathElements":              RemainingPathElements,
	"getFromAllPath":                     GetFromAllPath,
	"isListSet":                          IsListSet,
	"isList":                             IsList,
	"isSet":                              IsSet,
	"isStringListSet":                    IsStringListSet,
	"isInt64ListSet":                     IsInt64ListSet,
	"isNestedListSet":                    IsNestedListSet,
	"isNestedList":                       IsNestedList,
	"isNestedSet":                        IsNestedSet,
	"isNestedMap":                        IsNestedMap,
	"isNestedListSetMap":                 IsNestedListSetMap,
	"hasKeyPart":                         HasKeyPart,
	"getKeyPartAttributes":               GetKeyPartAttributes,
	"importAttributes":                   ImportAttributes,
	"floatReadExpr":                      FloatReadExpr,
	"subtract":                           Subtract,
	"hasWriteOnlyTFChildren":             HasWriteOnlyTFChildren,
	"writeOnlyTFChildren":                WriteOnlyTFChildren,
	"writeOnlyTFParentLists":             WriteOnlyTFParentLists,
	"coexistingSecretAttributes":         CoexistingSecretAttributes,
	"coexistingSecretChildren":           CoexistingSecretChildren,
	"coexistingSecretParentLists":        CoexistingSecretParentLists,
	"getIfUnsetOnUpdateAttributes":       GetIfUnsetOnUpdateAttributes,
}

func augmentAttribute(attr *YamlConfigAttribute) {
	if attr.TfName == "" {
		var words []string
		l := 0
		for s := attr.ModelName; s != ""; s = s[l:] {
			l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
			if l <= 0 {
				l = len(s)
			}
			words = append(words, strings.ToLower(s[:l]))
		}
		attr.TfName = strings.Join(words, "_")
	}
	if attr.Type == "List" || attr.Type == "Set" || attr.Type == "Map" {
		for a := range attr.Attributes {
			augmentAttribute(&attr.Attributes[a])
		}
	}
	// Auto-populate MapKeyExample for nested maps from key_part child Example values
	if attr.Type == "Map" && attr.MapKeyExample == "" && len(attr.Attributes) > 0 {
		var keyParts []string
		for _, child := range attr.Attributes {
			if child.KeyPart && child.Example != "" {
				keyParts = append(keyParts, child.Example)
			}
		}
		if len(keyParts) > 0 {
			attr.MapKeyExample = strings.Join(keyParts, "_")
		}
	}
}

func augmentConfig(config *YamlConfig) {
	for ia := range config.Attributes {
		augmentAttribute(&config.Attributes[ia])
	}
	// For each top-level attribute marked write_only_tf, rename it to "<tf_name>_wo"
	// (the Terraform-core write-only attribute) and inject a companion
	// "<tf_name>_wo_version" (Int64, Optional) that IS stored in state and drives
	// rotation. The base name (before "_wo") is derived once, since augmentAttribute
	// has already populated TfName.
	augmentWriteOnlyTF(config)
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s.", config.Name)
	}
	if config.ResDescription == "" {
		name := strings.ToLower(config.Name)
		if strings.HasPrefix(name, "a") || strings.HasPrefix(name, "e") || strings.HasPrefix(name, "i") || strings.HasPrefix(name, "o") || strings.HasPrefix(name, "u") {
			config.ResDescription = fmt.Sprintf("This resource can manage an %s.", config.Name)
		} else {
			config.ResDescription = fmt.Sprintf("This resource can manage a %s.", config.Name)
		}
	}
}

// augmentWriteOnlyTF expands every attribute flagged write_only_tf into three
// coexisting attributes, so that adding write-only support is backwards compatible:
//
//	<tf_name>              the original attribute, kept as-is. Still Optional
//	                       and still written to the same API path, so existing
//	                       configurations keep working. The secret remains in state.
//	<tf_name>_wo           the Terraform-core write-only variant. Same ModelName /
//	                       DataPath / PutDataPath, so the secret reaches the same API
//	                       field, but it is never persisted to plan or state.
//	<tf_name>_wo_version   Int64, Optional, stored in state, never sent to the API.
//	                       The rotation trigger for the write-only variant.
//
// The legacy attribute and the "_wo" variant are mutually exclusive; the schema
// template emits ExactlyOneOf when the secret is mandatory and ConflictsWith when it
// is optional. Because both carry the same ModelName, toBody writes whichever of the
// two is non-null to the same JSON path, and the validators guarantee they are never
// both set.
func augmentWriteOnlyTF(config *YamlConfig) {
	config.Attributes = rewriteWriteOnlyTF(config.Attributes)
}

// rewriteWriteOnlyTF walks a single attribute list. For each attribute it first
// recurses into any nested children (list-of-objects), so a write-only secret nested
// arbitrarily deep is expanded too. Then, if the attribute itself is flagged
// write_only_tf, it is replaced by the legacy/"_wo"/"_wo_version" triple described on
// augmentWriteOnlyTF. Nested secrets get a per-element version companion, matching the
// per-element rotation granularity of the enclosing list.
func rewriteWriteOnlyTF(attrs []YamlConfigAttribute) []YamlConfigAttribute {
	newAttrs := make([]YamlConfigAttribute, 0, len(attrs))
	for _, attr := range attrs {
		if len(attr.Attributes) > 0 {
			attr.Attributes = rewriteWriteOnlyTF(attr.Attributes)
		}
		if !attr.WriteOnlyTF {
			newAttrs = append(newAttrs, attr)
			continue
		}
		// The generated "_wo" schema entry emits []validator.String, so a non-String
		// secret would not compile. Fail here rather than at build time, with a message
		// that names the offending attribute.
		if attr.Type != "String" {
			panic(fmt.Sprintf("write_only_tf is only supported on String attributes, but %q has type %q", attr.TfName, attr.Type))
		}
		// write_only_tf relies on write_only to keep the secret out of the read path.
		// fromBody and updateFromBody skip attributes on .WriteOnly alone, so a secret
		// flagged write_only_tf without it would have the API response written back into
		// its "_wo" field on every read. The framework nulls write-only attributes before
		// they reach state, so that is currently dead work rather than a leak, but it
		// leaves the guarantee resting entirely on framework behaviour. Require the flags
		// together instead.
		if !attr.WriteOnly {
			panic(fmt.Sprintf("write_only_tf requires write_only to be set as well, but %q sets only write_only_tf", attr.TfName))
		}
		baseName := attr.TfName

		// Both halves of the pair document the constraint the validators enforce.
		// tfplugindocs derives Required/Optional from the schema booleans alone, so a
		// mandatory secret reachable through either spelling would otherwise be listed as
		// merely Optional twice, with nothing saying one of them has to be supplied.
		// Computed from the original attribute, before Mandatory is cleared below.
		// A mandatory secret leads with "**Required**" because the pair is listed under
		// the documentation's Optional heading: neither half can carry the schema's
		// Required flag without making that spelling the only usable one, so the word has
		// to come from the description instead.
		exclusivity := fmt.Sprintf("Only one of `%s` and `%s_wo` can be set.", baseName, baseName)
		if attr.Mandatory {
			exclusivity = fmt.Sprintf("**Required**: exactly one of `%s` and `%s_wo` must be set.", baseName, baseName)
		}

		// The legacy attribute keeps its name and its place in the schema so existing
		// configurations are untouched. It is no longer Mandatory on its own, because the
		// "_wo" variant is an equally valid way to supply the secret; ExactlyOneOf carries
		// that requirement instead. It is excluded from the generated acceptance tests so
		// the test configurations exercise the write-only path and do not trip the
		// mutual-exclusion validator by setting both spellings.
		legacy := attr
		legacy.WriteOnlyTF = false
		legacy.CoexistingSecret = true
		legacy.Mandatory = false
		legacy.ExcludeTest = true
		legacy.ExcludeExample = true
		// MinimumTestValue has to be cleared as well, not just ExcludeTest: the minimum
		// test configuration selects attributes on MinimumTestValue without consulting
		// ExcludeTest, so a secret carrying one would otherwise be emitted alongside its
		// "_wo" twin and trip the mutual-exclusion validator. No currently converted
		// secret carries a minimum_test_value, but catalystcenter_device's do, so this
		// matters as soon as that resource is converted.
		legacy.MinimumTestValue = ""
		legacy.CoexistenceNote = fmt.Sprintf("This attribute stores the secret in Terraform state. Prefer `%s_wo` together with `%s_wo_version`, which keeps it out of state.", baseName, baseName)
		legacy.MutualExclusivityNote = exclusivity
		legacy.WoPairMandatory = attr.Mandatory
		legacy.WoPairHasVersion = !attr.RequiresReplace
		newAttrs = append(newAttrs, legacy)

		// The write-only variant is always Optional in the schema, because the legacy
		// attribute is an equally valid way to supply the secret. Mandatory is kept so the
		// template can pick ExactlyOneOf (the secret must come from one of the two) over
		// ConflictsWith (at most one of the two) for the mutual-exclusion validator.
		wo := attr
		wo.TfName = baseName + "_wo"
		wo.WoBaseName = baseName
		wo.MutualExclusivityNote = exclusivity
		newAttrs = append(newAttrs, wo)

		// requires_replace secrets rotate by recreate, so a state-stored version int
		// (which can only drive an in-place update) is meaningless. Convert to write-only
		// but drop the version companion.
		if attr.RequiresReplace {
			continue
		}

		version := YamlConfigAttribute{
			TfName:      baseName + "_wo_version",
			Type:        "Int64",
			WoVersion:   true,
			Description: fmt.Sprintf("Rotation trigger for `%s_wo`. Increment this integer whenever the write-only value changes so Terraform sends the new secret. The value is stored in state; the secret is not.", baseName),
			Example:     "1",
			ExcludeTest: attr.ExcludeTest,
		}
		// The version belongs in the minimum test configuration exactly when its "_wo"
		// secret does, since the schema requires the pair together. The minimum template
		// selects on Mandatory/MinimumTestValue and does not consult ExcludeTest, so both
		// conditions have to be mirrored here or the version would be emitted alone.
		if !attr.ExcludeTest && (attr.Mandatory || attr.MinimumTestValue != "") {
			version.MinimumTestValue = "1"
		}
		newAttrs = append(newAttrs, version)
	}
	return newAttrs
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		output = bytes.NewBufferString(newContent)
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	files, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(files))

	// Load configs
	for i, filename := range files {
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	for i := range configs {
		// Augment config
		augmentConfig(&configs[i])

		// Iterate over templates and render files
		for _, t := range templates {
			if (configs[i].NoImport && t.path == "./gen/templates/import.sh") ||
				(configs[i].NoDataSource && t.path == "./gen/templates/data_source.go") ||
				(configs[i].NoDataSource && t.path == "./gen/templates/data_source_test.go") ||
				(configs[i].NoDataSource && t.path == "./gen/templates/data-source.tf") ||
				(configs[i].NoResource && t.path == "./gen/templates/resource.go") ||
				(configs[i].NoResource && t.path == "./gen/templates/resource_test.go") ||
				(configs[i].NoResource && t.path == "./gen/templates/resource.tf") ||
				(configs[i].NoResource && t.path == "./gen/templates/import.sh") {
				continue
			}
			renderTemplate(t.path, t.prefix+SnakeCase(configs[i].Name)+t.suffix, configs[i])
		}
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, configs)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
