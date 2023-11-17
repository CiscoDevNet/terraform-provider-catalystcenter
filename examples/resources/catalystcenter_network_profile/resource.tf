resource "catalystcenter_network_profile" "example" {
  name = "Profile1"
  type = "switching"
  templates = [
    {
      type        = "cli.templates"
      template_id = "f8297e86-35b0-486c-8752-6169aa5eb43c"
    }
  ]
}
