resource "catalystcenter_wireless_profile" "example" {
  wireless_profile_name = "Wireless_Profile_1"
  ssid_details = [
    {
      ssid_name           = "mySSID1"
      enable_fabric       = false
      enable_flex_connect = false
    }
  ]
  ap_zones = [
    {
    }
  ]
}
