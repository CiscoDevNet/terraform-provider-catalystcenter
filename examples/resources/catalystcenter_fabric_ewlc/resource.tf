resource "catalystcenter_fabric_ewlc" "example" {
  fabric_id                 = "8e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  network_device_id         = "12345678-1234-1234-1234-123456789012"
  enable_wireless           = true
  enable_rolling_ap_upgrade = false
  ap_reboot_percentage      = 25
}
