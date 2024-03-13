resource "catalystcenter_virtual_network_to_fabric_site" "example" {
  virtual_network_name = "SDA_VN1"
  site_name_hierarchy  = "Global/Area1"
}
