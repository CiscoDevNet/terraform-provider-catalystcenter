resource "catalystcenter_device" "example" {
  serial_number = "FOC12345678"
  stack         = false
  pid           = "C9300-24P"
  hostname      = "switch1"
}
