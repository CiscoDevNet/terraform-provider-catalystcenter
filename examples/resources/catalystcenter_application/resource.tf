resource "catalystcenter_application" "example" {
  name = "my-custom-app"
  network_applications = [
    {
      help_string      = "This is a custom application"
      application_type = "CUSTOM"
      server_type      = "_server-ip"
      dscp             = "18"
      app_protocol     = "TCP"
      server_name      = "app.example.com"
      url              = "example.com/path"
      traffic_class    = "TRANSACTIONAL_DATA"
      category_id      = "de3f7e1e-bdda-4a2c-87f8-0ff49b463997"
      ignore_conflict  = false
      rank             = 1
      engine_id        = 6
    }
  ]
  application_set_id = "12345678-1234-1234-1234-123456789012"
  network_identity = [
    {
      protocol    = "TCP"
      ports       = "65529"
      lower_port  = 0
      upper_port  = 0
      ipv4_subnet = ["10.0.0.0/24"]
    }
  ]
}
