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
	Name                        string                `yaml:"name"`
	NoResource                  bool                  `yaml:"no_resource"`
	NoDataSource                bool                  `yaml:"no_data_source"`
	RestEndpoint                string                `yaml:"rest_endpoint"`
	GetRestEndpoint             string                `yaml:"get_rest_endpoint"`
	PutRestEndpoint             string                `yaml:"put_rest_endpoint"`
	DeleteRestEndpoint          string                `yaml:"delete_rest_endpoint"`
	GetNoId                     bool                  `yaml:"get_no_id"`
	GetFromAll                  bool                  `yaml:"get_from_all"`
	GetRequiresId               bool                  `yaml:"get_requires_id"`
	GetExtraQueryParams         string                `yaml:"get_extra_query_params"`
	NoDelete                    bool                  `yaml:"no_delete"`
	DataSourceNoId              bool                  `yaml:"data_source_no_id"`
	DeleteNoId                  bool                  `yaml:"delete_no_id"`
	NoUpdate                    bool                  `yaml:"no_update"`
	NoRead                      bool                  `yaml:"no_read"`
	NoImport                    bool                  `yaml:"no_import"`
	ImportNoId                  bool                  `yaml:"import_no_id"`
	PostUpdate                  bool                  `yaml:"post_update"`
	PutCreate                   bool                  `yaml:"put_create"`
	PutDelete                   bool                  `yaml:"put_delete"`
	UpdateComputed              bool                  `yaml:"update_computed"`
	RootList                    bool                  `yaml:"root_list"`
	NoReadPrefix                bool                  `yaml:"no_read_prefix"`
	IdPath                      string                `yaml:"id_path"`
	IdFromQueryPath             string                `yaml:"id_from_query_path"`
	IdFromQueryPathAttribute    string                `yaml:"id_from_query_path_attribute"`
	IdQueryParam                string                `yaml:"id_query_param"`
	IdFromAttribute             bool                  `yaml:"id_from_attribute"`
	DeviceUnreachabilityWarning bool                  `yaml:"device_unreachability_warning"`
	PutIdIncludePath            string                `yaml:"put_id_include_path"`
	PutIdQueryParam             string                `yaml:"put_id_query_param"`
	PutNoId                     bool                  `yaml:"put_no_id"`
	PutUpdateId                 bool                  `yaml:"put_update_id"`
	DeleteIdQueryParam          string                `yaml:"delete_id_query_param"`
	MinimumVersion              string                `yaml:"minimum_version"`
	DsDescription               string                `yaml:"ds_description"`
	ResDescription              string                `yaml:"res_description"`
	DocCategory                 string                `yaml:"doc_category"`
	ExcludeTest                 bool                  `yaml:"exclude_test"`
	SkipMinimumTest             bool                  `yaml:"skip_minimum_test"`
	Attributes                  []YamlConfigAttribute `yaml:"attributes"`
	TestTags                    []string              `yaml:"test_tags"`
	TestPrerequisites           string                `yaml:"test_prerequisites"`
	MaxAsyncWaitTime            int64                 `yaml:"max_async_wait_time"`
}

type YamlConfigAttribute struct {
	ModelName            string                `yaml:"model_name"`
	ResponseModelName    string                `yaml:"response_model_name"`
	TfName               string                `yaml:"tf_name"`
	Type                 string                `yaml:"type"`
	ElementType          string                `yaml:"element_type"`
	DataPath             string                `yaml:"data_path"`
	ResponseDataPath     string                `yaml:"response_data_path"`
	Id                   bool                  `yaml:"id"`
	MatchId              bool                  `yaml:"match_id"`
	Reference            bool                  `yaml:"reference"`
	RequiresReplace      bool                  `yaml:"requires_replace"`
	QueryParam           bool                  `yaml:"query_param"`
	DeleteQueryParam     bool                  `yaml:"delete_query_param"`
	GetQueryParam        bool                  `yaml:"get_query_param"`
	PutQueryParam        bool                  `yaml:"put_query_param"`
	PostQueryParam       bool                  `yaml:"post_query_param"`
	QueryParamName       string                `yaml:"query_param_name"`
	DeleteQueryParamName string                `yaml:"delete_query_param_name"`
	GetQueryParamName    string                `yaml:"get_query_param_name"`
	PutQueryParamName    string                `yaml:"put_query_param_name"`
	PostQueryParamName   string                `yaml:"post_query_param_name"`
	CreateQueryPath      bool                  `yaml:"create_query_path"`
	DataSourceQuery      bool                  `yaml:"data_source_query"`
	QueryParamNoBody     bool                  `yaml:"query_param_no_body"`
	Mandatory            bool                  `yaml:"mandatory"`
	Computed             bool                  `yaml:"computed"`
	WriteOnly            bool                  `yaml:"write_only"`
	ExcludeFromPut       bool                  `yaml:"exclude_from_put"`
	ExcludeTest          bool                  `yaml:"exclude_test"`
	ExcludeExample       bool                  `yaml:"exclude_example"`
	Description          string                `yaml:"description"`
	Example              string                `yaml:"example"`
	EnumValues           []string              `yaml:"enum_values"`
	MinList              int64                 `yaml:"min_list"`
	MaxList              int64                 `yaml:"max_list"`
	MinInt               int64                 `yaml:"min_int"`
	MaxInt               int64                 `yaml:"max_int"`
	MinFloat             float64               `yaml:"min_float"`
	MaxFloat             float64               `yaml:"max_float"`
	StringPatterns       []string              `yaml:"string_patterns"`
	StringMinLength      int64                 `yaml:"string_min_length"`
	StringMaxLength      int64                 `yaml:"string_max_length"`
	DefaultValue         string                `yaml:"default_value"`
	Value                string                `yaml:"value"`
	ValueCondition       string                `yaml:"value_condition"`
	TestValue            string                `yaml:"test_value"`
	MinimumTestValue     string                `yaml:"minimum_test_value"`
	TestTags             []string              `yaml:"test_tags"`
	Attributes           []YamlConfigAttribute `yaml:"attributes"`
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

// Templating helper function to return true if delete query parameters included in attributes
func HasDeleteQueryParam(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.DeleteQueryParam {
			return true
		}
	}
	return false
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
			if first {
				if attr.Type == "Int64" {
					params = append(params, `"?`+queryParamName+`=" + url.QueryEscape(strconv.FormatInt(`+inputSource+`.`+ToGoName(attr.TfName)+`.Value`+attr.Type+`(), 10))`)
				} else {
					params = append(params, `"?`+queryParamName+`=" + url.QueryEscape(`+inputSource+`.`+ToGoName(attr.TfName)+`.Value`+attr.Type+`())`)
				}
				first = false
			} else {
				if attr.Type == "Int64" {
					params = append(params, `"&`+queryParamName+`=" + url.QueryEscape(strconv.FormatInt(`+inputSource+`.`+ToGoName(attr.TfName)+`.Value`+attr.Type+`(), 10))`)
				} else {
					params = append(params, `"&`+queryParamName+`=" + url.QueryEscape(`+inputSource+`.`+ToGoName(attr.TfName)+`.Value`+attr.Type+`())`)
				}
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

// Templating helper function to return a list of import attributes
func ImportAttributes(config YamlConfig) []YamlConfigAttribute {
	r := []YamlConfigAttribute{}
	for _, attr := range config.Attributes {
		if attr.Reference || attr.QueryParam || attr.Id {
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

// Templating helper function to subtract one number from another
func Subtract(a, b int) int {
	return a - b
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                 ToGoName,
	"camelCase":                CamelCase,
	"strContains":              strings.Contains,
	"snakeCase":                SnakeCase,
	"sprintf":                  fmt.Sprintf,
	"toLower":                  strings.ToLower,
	"path":                     BuildPath,
	"hasId":                    HasId,
	"hasReference":             HasReference,
	"hasQueryParam":            HasQueryParam,
	"hasDeleteQueryParam":      HasDeleteQueryParam,
	"generateQueryParamString": GenerateQueryParamString,
	"getId":                    GetId,
	"getMatchId":               GetMatchId,
	"hasCreateQueryPath":       HasCreateQueryPath,
	"getCreateQueryPath":       GetCreateQueryPath,
	"getQueryParam":            GetQueryParam,
	"getDeleteQueryParam":      GetDeleteQueryParam,
	"hasDataSourceQuery":       HasDataSourceQuery,
	"firstPathElement":         FirstPathElement,
	"remainingPathElements":    RemainingPathElements,
	"getFromAllPath":           GetFromAllPath,
	"isListSet":                IsListSet,
	"isList":                   IsList,
	"isSet":                    IsSet,
	"isStringListSet":          IsStringListSet,
	"isInt64ListSet":           IsInt64ListSet,
	"isNestedListSet":          IsNestedListSet,
	"isNestedList":             IsNestedList,
	"isNestedSet":              IsNestedSet,
	"importAttributes":         ImportAttributes,
	"subtract":                 Subtract,
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
	if attr.Type == "List" || attr.Type == "Set" {
		for a := range attr.Attributes {
			augmentAttribute(&attr.Attributes[a])
		}
	}
}

func augmentConfig(config *YamlConfig) {
	for ia := range config.Attributes {
		augmentAttribute(&config.Attributes[ia])
	}
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
