resource "catalystcenter_dot11be_profile" "example" {
  profile_name      = "dot11be_profile_1"
  ofdma_down_link   = true
  ofdma_up_link     = true
  mu_mimo_down_link = false
  mu_mimo_up_link   = false
  ofdma_multi_ru    = false
}
