resource "catalystcenter_sda_virtual_network_ip_pool" "example" {
  virtual_network_name  = "SDA_VN1"
  site_name_hierarchy   = "Global/Area1"
  layer2_only           = false
  ip_pool_name          = "MyPool1"
  vlan_id               = "401"
  vlan_name             = "VLAN401"
  traffic_type          = "DATA"
  l2_flooding_enabled   = false
  critical_pool         = false
  wireless_pool         = false
  ip_directed_broadcast = false
  common_pool           = false
  bridge_mode_vm        = false
}
