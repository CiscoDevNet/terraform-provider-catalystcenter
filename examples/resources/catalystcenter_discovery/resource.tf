resource "catalystcenter_discovery" "example" {
  discovery_type            = "RANGE"
  global_credential_id_list = [""]
  ip_address_list           = "192.168.0.1-192.168.0.99"
  name                      = "disco42"
  netconf_port              = "830"
  preferred_ip_method       = "UseLoopBack"
  protocol_order            = "ssh"
  retry                     = 3
  timeout                   = 5
}
