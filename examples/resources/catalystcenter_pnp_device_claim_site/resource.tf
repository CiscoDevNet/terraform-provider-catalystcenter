resource "catalystcenter_pnp_device_claim_site" "example" {
  device_id  = "12345678-1234-1234-1234-123456789012"
  site_id    = "12345678-1234-1234-1234-123456789012"
  type       = "Default"
  image_id   = ""
  image_skip = true
  config_id  = "template1"
  config_parameters = [
    {
      name  = "HOSTNAME"
      value = "switch1"
    }
  ]
}
