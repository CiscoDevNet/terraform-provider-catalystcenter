terraform import catalystcenter_{{snakeCase .Name}}.example "{{range $index, $attr := importAttributes .}}{{if $index}},{{end}}<{{$attr.TfName}}>{{else}}4b0b7a80-44c0-4bf2-bab5-fc24b4e0a17e{{end}}"
