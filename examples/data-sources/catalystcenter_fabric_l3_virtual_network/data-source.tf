data "catalystcenter_fabric_l3_virtual_network" "example" {
  virtual_network_name = "MyL3VN"
  fabric_ids           = ["5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"]
  anchored_site_id     = ""
}
