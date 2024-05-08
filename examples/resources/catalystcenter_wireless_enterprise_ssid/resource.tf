resource "catalystcenter_wireless_enterprise_ssid" "example" {
  name                                  = "mySSID1"
  security_level                        = "wpa3_enterprise"
  passphrase                            = "Cisco123"
  enable_fast_lane                      = false
  enable_mac_filtering                  = false
  traffic_type                          = "data"
  radio_policy                          = "Triple band operation(2.4GHz, 5GHz and 6GHz)"
  enable_broadcast_ssid                 = true
  fast_transition                       = "Adaptive"
  enable_session_time_out               = true
  session_time_out                      = 1800
  enable_client_exclusion               = true
  client_exclusion_timeout              = 180
  enable_basic_service_set_max_idle     = true
  basic_service_set_client_idle_timeout = 300
  enable_directed_multicast_service     = true
  enable_neighbor_list                  = true
  mfp_client_protection                 = "Optional"
  protected_management_frame            = "Required"
}
