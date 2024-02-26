resource "catalystcenter_fabric_virtual_network" "example" {
  virtual_network_name = "SDA_VN1"
  is_guest             = false
  sg_names             = ["Employees"]
}
