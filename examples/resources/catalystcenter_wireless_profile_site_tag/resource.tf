resource "catalystcenter_wireless_profile_site_tag" "example" {
  wireless_profile_id = "12345678-1234-1234-1234-123456789012"
  site_tag_name       = "SiteTag1"
  ap_profile_name     = "default-ap-profile"
  site_ids            = ["12345678-1234-1234-1234-123456789000"]
}
