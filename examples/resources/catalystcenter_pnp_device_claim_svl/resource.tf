resource "catalystcenter_pnp_device_claim_svl" "example" {
  device_id   = "683413ec67f7d77edb0fd605"
  template_id = "1e3b9f1b-acd9-462a-8a2e-13b527934549"
  template_parameters = [
    {
      name  = "param5"
      value = "value5"
    }
  ]
  domain = 2
  svl_members = [
    {
      serial_number = "FJC28041KKA"
      role          = "ACTIVE"
      svl_links = [
        {
          local_interface  = "TwentyFiveGigE1/0/47"
          remote_interface = "TwentyFiveGigE2/0/47"
        }
      ]
    }
  ]
}
