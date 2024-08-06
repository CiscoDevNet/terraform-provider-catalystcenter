resource "catalystcenter_wireless_device_provision" "example" {
  device_name          = "WLC_01"
  site                 = "Global/Area1"
  managed_ap_locations = ["Global/Area1"]
}
