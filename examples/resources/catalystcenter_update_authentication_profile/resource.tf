resource "catalystcenter_update_authentication_profile" "example" {
  fabric_id                     = "ae8a501f-2cdb-4d87-b535-1e568a091de1"
  authentication_profile_name   = "Open Authentication"
  authentication_order          = "mac"
  dot1x_to_mab_fallback_timeout = 10
  wake_on_lan                   = true
  number_of_hosts               = "Unlimited"
}
