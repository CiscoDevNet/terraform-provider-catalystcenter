resource "catalystcenter_virtual_network_to_fabric_site" "example" {
  virtual_network_name = "SDA_VN1"
  virtual_network_id   = "c782eff3-b743-4da2-a6ea-a27582a71287"
  fabric_site_id       = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
}
