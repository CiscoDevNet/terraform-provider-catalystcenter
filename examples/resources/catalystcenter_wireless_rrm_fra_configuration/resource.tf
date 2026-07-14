resource "catalystcenter_wireless_rrm_fra_configuration" "example" {
  design_name     = "Test"
  radio_band      = "2_4GHZ_5GHZ"
  fra_freeze      = true
  fra_status      = true
  fra_interval    = 3
  fra_sensitivity = "MEDIUM"
}
