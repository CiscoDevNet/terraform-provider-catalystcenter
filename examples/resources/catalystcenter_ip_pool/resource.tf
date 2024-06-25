resource "catalystcenter_ip_pool" "example" {
  name             = "MyPool1"
  ip_address_space = "IPv4"
  type             = "generic"
  ip_subnet        = "21.1.1.0/24"
  gateway          = ["21.1.1.1"]
  dhcp_server_ips  = ["1.2.3.4"]
  dns_server_ips   = ["2.3.4.5"]
}
