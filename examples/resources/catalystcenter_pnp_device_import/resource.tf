resource "catalystcenter_pnp_device_import" "example" {
  devices = [
    {
      serial_number = "FOC12345678"
      stack         = false
      pid           = "C9300-24P"
      hostname      = "switch1"
    }
  ]
}
