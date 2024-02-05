resource "catalystcenter_discovery" "example" {
  discovery_type  = "Range"
  ip_address_list = "192.168.0.1-192.168.0.99"
  name            = "disco42"
  netconf_port    = "830"
  protocol_order  = "SSH"
  retry           = 3
  timeout_seconds = 5
}
