resource "catalystcenter_wireless_ssid" "example" {
  site_id                               = "5e8e3e3e-1b6b-4b6b-8b6b-1b6b4b6b8b6b"
  ssid                                  = "mySSID1"
  auth_type                             = "WPA3_PERSONAL"
  passphrase                            = "Cisco123"
  fast_lane                             = false
  mac_filtering                         = false
  ssid_radio_type                       = "Triple band operation(2.4GHz, 5GHz and 6GHz)"
  broadcast_ssid                        = true
  fast_transition                       = "ADAPTIVE"
  session_timeout_enable                = true
  session_timeout                       = 1800
  client_exclusion                      = true
  client_exclusion_timeout              = 1800
  basic_service_set_max_idle            = true
  basic_service_set_client_idle_timeout = 300
  directed_multicast_service            = true
  neighbor_list                         = true
  mft_client_protection                 = "OPTIONAL"
  aaa_override                          = false
  protected_management_frame            = "REQUIRED"
  rsn_cipher_suite_ccmp128              = true
  wlan_type                             = "Enterprise"
  auth_key_sae_ext                      = true
  ghz24_policy                          = "dot11-g-only"
  hex                                   = false
  random_mac_filter                     = false
}
