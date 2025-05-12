resource "catalystcenter_update_authentication_profile" "example" {
  authentication_profile_id     = "cc0b30bc-94e7-458a-80e2-c7bbecc829f5"
  authentication_profile_name   = "Closed Authentication"
  authentication_order          = "mac"
  dot1x_to_mab_fallback_timeout = 10
  wake_on_lan                   = true
  number_of_hosts               = "Unlimited"
}
