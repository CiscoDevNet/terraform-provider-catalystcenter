resource "catalystcenter_application" "example" {
  name               = "my-custom-app"
  application_set_id = "12345678-1234-1234-1234-123456789012"
  category_id        = "f378fea2-3b99-4172-b73b-2fd17a8b1df1"
  traffic_class      = "TRANSACTIONAL_DATA"
  help_string        = "Custom internal collaboration app"
  dscp               = "18"
  rank               = 1
  app_protocol       = "TCP"
  server_type        = "_servername"
  server_name        = "app.example.com"
  url                = "example.com/path"
  network_identity = [
    {
      protocol    = "TCP"
      ports       = "8080"
      ipv4_subnet = ["10.0.0.0/24"]
    }
  ]
}
