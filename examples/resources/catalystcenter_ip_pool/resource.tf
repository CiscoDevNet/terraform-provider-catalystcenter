resource "catalystcenter_ip_pool" "example" {
  name                        = "MyPool1"
  pool_type                   = "Generic"
  address_space_subnet        = "192.168.0.0"
  address_space_prefix_length = 16
  address_space_gateway       = "192.168.0.1"
  address_space_dhcp_servers  = ["192.168.0.10"]
  address_space_dns_servers   = ["192.168.0.53"]
}
