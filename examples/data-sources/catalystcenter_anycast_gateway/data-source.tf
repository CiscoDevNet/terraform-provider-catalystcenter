data "catalystcenter_anycast_gateway" "example" {
  fabric_id            = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  virtual_network_name = "SDA_VN1"
  ip_pool_name         = "MyRes1"
}
