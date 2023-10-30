resource "catalystcenter_ip_pool_reservation" "example" {
  site_id            = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  name               = "MyRes1"
  type               = "Generic"
  ipv6_address_space = false
  ipv4_global_pool   = "172.32.0.0/16"
  ipv4_prefix        = true
  ipv4_prefix_length = 24
  ipv4_subnet        = "172.32.1.0"
  ipv4_gateway       = "172.32.1.1"
  ipv4_dhcp_servers  = ["1.2.3.4"]
  ipv4_dns_servers   = ["2.3.4.5"]
}
