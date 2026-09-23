resource "catalystcenter_wireless_profile_site_tag_site_assignment" "example" {
  wireless_profile_id = "12345678-1234-1234-1234-123456789012"
  site_tag_name       = "SiteTag1"
  site_id             = "12345678-1234-1234-1234-123456789000"
  ap_profile_name     = "default-ap-profile"
}
