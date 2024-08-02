resource "catalystcenter_anycast_gateway" "example" {
  fabric_id                    = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  virtual_network_name         = "SDA_VN1"
  ip_pool_name                 = "MyRes1"
  tcp_mss_adjustment           = 1400
  vlan_name                    = "VLAN401"
  vlan_id                      = 401
  traffic_type                 = "DATA"
  critical_pool                = false
  l2_flooding_enabled          = false
  wireless_pool                = false
  ip_directed_broadcast        = false
  intra_subnet_routing_enabled = false
  multiple_ip_to_mac_addresses = false
}
