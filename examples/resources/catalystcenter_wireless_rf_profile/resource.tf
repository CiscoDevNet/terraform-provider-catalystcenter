resource "catalystcenter_wireless_rf_profile" "example" {
  rf_profile_name         = "RF_Profile_1"
  default_rf_profile      = false
  enable_radio_type_a     = true
  enable_radio_type_b     = false
  enable_radio_type6_g_hz = false
}
