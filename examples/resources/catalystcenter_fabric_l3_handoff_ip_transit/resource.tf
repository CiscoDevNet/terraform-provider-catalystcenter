resource "catalystcenter_fabric_l3_handoff_ip_transit" "example" {
  network_device_id    = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  fabric_id            = "c4b85bb2-ce3f-4db9-a32b-e439a388ac2f"
  transit_network_id   = "CONTROL_PLANE_NODE"
  interface_name       = "TenGigabitEthernet1/0/2"
  virtual_network_name = "SDA_VN1"
  vlan_id              = 205
  tcp_mss_adjustment   = 1400
  local_ip_address     = "10.0.0.1/24"
  remote_ip_address    = "10.0.0.2/24"
}
