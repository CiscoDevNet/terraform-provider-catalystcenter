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
	Name                     string                `yaml:"name"`
	NoResource               bool                  `yaml:"no_resource"`
	NoDataSource             bool                  `yaml:"no_data_source"`
	RestEndpoint             string                `yaml:"rest_endpoint"`
	GetRestEndpoint          string                `yaml:"get_rest_endpoint"`
	PutRestEndpoint          string                `yaml:"put_rest_endpoint"`
	DeleteRestEndpoint       string                `yaml:"delete_rest_endpoint"`
	GetNoId                  bool                  `yaml:"get_no_id"`
	GetFromAll               bool                  `yaml:"get_from_all"`
	GetRequiresId            bool                  `yaml:"get_requires_id"`
	GetExtraQueryParams      string                `yaml:"get_extra_query_params"`
	NoDelete                 bool                  `yaml:"no_delete"`
	NoRead                   bool                  `yaml:"no_read"`
	NoImport                 bool                  `yaml:"no_import"`
	PostUpdate               bool                  `yaml:"post_update"`
	RootList                 bool                  `yaml:"root_list"`
	NoReadPrefix             bool                  `yaml:"no_read_prefix"`
	IdPath                   string                `yaml:"id_path"`
	IdFromQueryPath          string                `yaml:"id_from_query_path"`
	IdFromQueryPathAttribute string                `yaml:"id_from_query_path_attribute"`
	IdQueryParam             string                `yaml:"id_query_param"`
	IdFromAttribute          string                `yaml:"id_from_attribute"`
	PutIdIncludePath         string                `yaml:"put_id_include_path"`
	PutIdQueryParam          string                `yaml:"put_id_query_param"`
	PutNoId                  bool                  `yaml:"put_no_id"`
	MinimumVersion           string                `yaml:"minimum_version"`
	DsDescription            string                `yaml:"ds_description"`
	ResDescription           string                `yaml:"res_description"`
	DocCategory              string                `yaml:"doc_category"`
	ExcludeTest              bool                  `yaml:"exclude_test"`
	SkipMinimumTest          bool                  `yaml:"skip_minimum_test"`
	Attributes               []YamlConfigAttribute `yaml:"attributes"`
	TestTags                 []string              `yaml:"test_tags"`
	TestPrerequisites        string                `yaml:"test_prerequisites"`
}

type YamlConfigAttribute struct {
	ModelName         string                `yaml:"model_name"`
	ResponseModelName string                `yaml:"response_model_name"`
	TfName            string                `yaml:"tf_name"`
	Type              string                `yaml:"type"`
	DataPath          string                `yaml:"data_path"`
	ResponseDataPath  string                `yaml:"response_data_path"`
	Id                bool                  `yaml:"id"`
	Reference         bool                  `yaml:"reference"`
	QueryParam        bool                  `yaml:"query_param"`
	DataSourceQuery   bool                  `yaml:"data_source_query"`
	Mandatory         bool                  `yaml:"mandatory"`
	WriteOnly         bool                  `yaml:"write_only"`
	WriteChangesOnly  bool                  `yaml:"write_changes_only"`
	ExcludeTest       bool                  `yaml:"exclude_test"`
	ExcludeExample    bool                  `yaml:"exclude_example"`
	Description       string                `yaml:"description"`
	Example           string                `yaml:"example"`
	EnumValues        []string              `yaml:"enum_values"`
	MinList           int64                 `yaml:"min_list"`
	MaxList           int64                 `yaml:"max_list"`
	MinInt            int64                 `yaml:"min_int"`
	MaxInt            int64                 `yaml:"max_int"`
	MinFloat          float64               `yaml:"min_float"`
	MaxFloat          float64               `yaml:"max_float"`
	StringPatterns    []string              `yaml:"string_patterns"`
	StringMinLength   int64                 `yaml:"string_min_length"`
	StringMaxLength   int64                 `yaml:"string_max_length"`
	DefaultValue      string                `yaml:"default_value"`
	Value             string                `yaml:"value"`
	ValueCondition    string                `yaml:"value_condition"`
	TestValue         string                `yaml:"test_value"`
	MinimumTestValue  string                `yaml:"minimum_test_value"`
	TestTags          []string              `yaml:"test_tags"`
	Attributes        []YamlConfigAttribute `yaml:"attributes"`
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

// Templating helper function to return the ID attribute
func GetId(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.Id {
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

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":              ToGoName,
	"camelCase":             CamelCase,
	"snakeCase":             SnakeCase,
	"sprintf":               fmt.Sprintf,
	"toLower":               strings.ToLower,
	"path":                  BuildPath,
	"hasId":                 HasId,
	"hasReference":          HasReference,
	"hasQueryParam":         HasQueryParam,
	"getId":                 GetId,
	"getQueryParam":         GetQueryParam,
	"hasDataSourceQuery":    HasDataSourceQuery,
	"firstPathElement":      FirstPathElement,
	"remainingPathElements": RemainingPathElements,
	"getFromAllPath":        GetFromAllPath,
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
