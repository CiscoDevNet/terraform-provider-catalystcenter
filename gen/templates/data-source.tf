data "catalystcenter_{{snakeCase .Name}}" "example" {
  {{- if and (not .DataSourceNoId) (not .DataSourceNoInputs)}}
  id = "{{$id := false}}{{range .Attributes}}{{if .Id}}{{$id = true}}{{.Example}}{{end}}{{end}}{{if not $id}}76d24097-41c4-4558-a4d0-a8c07ac08470{{end}}"
  {{- end}}
  {{- range  .Attributes}}
  {{- if and (not $.DataSourceNoInputs) (or .Reference .GetQueryParam .QueryParam (and .Id .QueryParamNoBody $.IdFromAttribute $.DataSourceNoId))}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
  {{- end}}
  {{- end}}
}
