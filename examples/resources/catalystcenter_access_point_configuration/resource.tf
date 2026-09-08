resource "catalystcenter_access_point_configuration" "example" {
  ap_list = [
    {
      mac_address = "00:11:22:33:44:55"
      ap_name     = "AP-old-name"
      ap_name_new = "AP-new-name"
    }
  ]
}
