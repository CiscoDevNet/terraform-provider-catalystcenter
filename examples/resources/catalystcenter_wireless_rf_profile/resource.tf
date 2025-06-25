resource "catalystcenter_wireless_rf_profile" "example" {
  rf_profile_name                      = "RF_Profile_1"
  default_rf_profile                   = false
  enable_radio_type_a                  = true
  enable_radio_type_b                  = true
  enable_radio_type6_g_hz              = false
  radio_type_a_parent_profile          = "CUSTOM"
  radio_type_a_radio_channels          = "36,40,44,48,124,128,157,161"
  radio_type_a_data_rates              = "6,9,12,18,24,36,48,54"
  radio_type_a_mandatory_data_rates    = "6"
  radio_type_a_power_threshold_v1      = -60
  radio_type_a_rx_sop_threshold        = "CUSTOM"
  radio_type_a_min_power_level         = 8
  radio_type_a_max_power_level         = 20
  radio_type_a_channel_width           = "20"
  radio_type_a_custom_rx_sop_threshold = -70
  radio_type_b_parent_profile          = "HIGH"
  radio_type_b_radio_channels          = "1,6,11"
  radio_type_b_data_rates              = "1,2,5.5,6,9,11,12,18,24,36,48,54"
  radio_type_b_mandatory_data_rates    = "1,2"
  radio_type_b_power_threshold_v1      = -60
  radio_type_b_rx_sop_threshold        = "CUSTOM"
  radio_type_b_min_power_level         = 8
  radio_type_b_max_power_level         = 20
  radio_type_b_custom_rx_sop_threshold = -70
}
