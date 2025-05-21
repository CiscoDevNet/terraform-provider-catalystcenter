resource "catalystcenter_deploy_template" "example" {
  template_id         = "12345678-1234-1234-1234-123456789012"
  force_push_template = false
  copying_config      = true
  is_composite        = false
  target_info = [
    {
      host_name             = "SW01"
      type                  = "MANAGED_DEVICE_HOSTNAME"
      versioned_template_id = "12345678-1234-1234-1234-123456789012"
    }
  ]
}
