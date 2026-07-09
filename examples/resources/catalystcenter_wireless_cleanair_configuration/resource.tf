resource "catalystcenter_wireless_cleanair_configuration" "example" {
  design_name                         = "cleanAir_802_11b_custom"
  radio_band                          = "2_4GHZ"
  clean_air                           = true
  clean_air_device_reporting          = true
  persistent_device_propagation       = true
  ble_beacon                          = true
  bluetooth_paging_inquiry            = true
  bluetooth_sco_acl                   = true
  continuous_transmitter              = true
  generic_dect                        = true
  generic_tdd                         = true
  jammer                              = true
  microwave_oven                      = true
  motorola_canopy                     = true
  si_fhss                             = true
  spectrum_80211_fh                   = true
  spectrum_80211_non_standard_channel = true
  spectrum_802154                     = true
  spectrum_inverted                   = true
  super_ag                            = true
  video_camera                        = true
  wimax_fixed                         = true
  wimax_mobile                        = true
  xbox                                = true
}
