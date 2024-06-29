resource "catalystcenter_wireless_rf_profile" "example" {
  name                              = "RF_Profile_1"
  default_rf_profile                = false
  enable_radio_type_a               = true
  enable_radio_type_b               = true
  enable_radio_type_c               = false
  channel_width                     = "20"
  enable_custom                     = true
  enable_brown_field                = false
  radio_type_a_parent_profile       = "CUSTOM"
  radio_type_a_radio_channels       = "36,40,44,48,52,56,60,64,144,149,153,157,161,165,169,173"
  radio_type_a_data_rates           = "6,9,12,18,24,36,48,54"
  radio_type_a_mandatory_data_rates = "12,24"
  radio_type_a_power_threshold_v1   = -60
  radio_type_a_rx_sop_threshold     = "LOW"
  radio_type_a_min_power_level      = 8
  radio_type_a_max_power_level      = 20
  radio_type_b_parent_profile       = "CUSTOM"
  radio_type_b_radio_channels       = "1,6,11"
  radio_type_b_data_rates           = "9,11,12,18,24,36,48,54"
  radio_type_b_mandatory_data_rates = "12"
  radio_type_b_power_threshold_v1   = -60
  radio_type_b_rx_sop_threshold     = "LOW"
  radio_type_b_min_power_level      = 8
  radio_type_b_max_power_level      = 20
}
