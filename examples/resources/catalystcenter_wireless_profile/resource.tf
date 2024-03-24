resource "catalystcenter_wireless_profile" "example" {
  name = "Wireless_Profile_1"
  ssid_details = [
    {
      name                = "mySSID1"
      enable_fabric       = true
      enable_flex_connect = false
    }
  ]
}
