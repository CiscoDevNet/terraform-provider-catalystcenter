resource "catalystcenter_wireless_device_provision" "example" {
  device_name          = "WLC_01"
  network_device_id    = "e2e6ae2f-d526-459f-bfdf-3281d74b6dea"
  site                 = "Global/Area1"
  managed_ap_locations = ["Global/Area1"]
}
