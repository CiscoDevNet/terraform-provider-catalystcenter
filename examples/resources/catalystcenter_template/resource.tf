resource "catalystcenter_template" "example" {
  project_id   = "12345678-1234-1234-1234-123456789012"
  name         = "Template1"
  project_name = "ProjectName"
  description  = "My description"
  device_types = [
    {
      product_family = "Switches and Hubs"
      product_series = "Cisco Catalyst 9300 Series Switches"
      product_type   = "Cisco Catalyst 9300 Switch"
    }
  ]
  language         = "JINJA"
  software_type    = "IOS-XE"
  software_variant = "XE"
  software_version = "16.12.1a"
  template_content = "hostname {{hostname}}"
  template_params = [
    {
      data_type        = "STRING"
      default_value    = "ABC"
      description      = "My parameter"
      display_name     = "Custom hostname"
      instruction_text = "My instructions"
      not_param        = false
      param_array      = false
      parameter_name   = "hostname"
      required         = false
      selection_type   = "SINGLE_SELECT"
      selection_values = { host1 = "host1" }
    }
  ]
  composite = false
}
