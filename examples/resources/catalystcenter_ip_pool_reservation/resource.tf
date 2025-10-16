resource "catalystcenter_ip_pool_reservation" "example" {
  name                = "MyRes1"
  pool_type           = "Generic"
  site_id             = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  ipv4_subnet         = "172.32.1.0"
  ipv4_prefix_length  = 24
  ipv4_gateway        = "172.32.1.1"
  ipv4_dhcp_servers   = ["1.2.3.4"]
  ipv4_dns_servers    = ["2.3.4.5"]
  ipv4_global_pool_id = "1b2c3d4e-1111-2222-3333-abcdefabcdef"
  ipv6_subnet         = "2001:db8:85a3:0:100::"
  ipv6_prefix_length  = 64
  ipv6_gateway        = "2001:db8:85a3:0:100::1"
  ipv6_dhcp_servers   = ["2001:db8::1234"]
  ipv6_dns_servers    = ["2001:db8::1234"]
  ipv6_slaac_support  = true
  ipv6_global_pool_id = "9f8e7d6c-aaaa-bbbb-cccc-1234567890ab"
}
