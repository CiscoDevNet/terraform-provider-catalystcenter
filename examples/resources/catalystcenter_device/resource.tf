resource "catalystcenter_device" "example" {
  cli_transport           = "ssh"
  compute_device          = false
  enable_password         = "cisco123"
  extended_discovery_info = "DISCOVER_WITH_CANNED_DATA"
  http_password           = "cisco123"
  http_port               = "80"
  http_secure             = true
  http_user_name          = "admin"
  ip_address              = "1.2.3.4"
  meraki_org_ids          = ["12345678901234567890"]
  netconf_port            = "830"
  password                = "cisco123"
  serial_number           = "FOC12345678"
  snmp_auth_passphrase    = "cisco123"
  snmp_auth_protocol      = "sha"
  snmp_mode               = "authPriv"
  snmp_priv_passphrase    = "cisco123"
  snmp_priv_protocol      = "AES128"
  snmp_ro_community       = "com1"
  snmp_rw_community       = "com2"
  snmp_retry              = 3
  snmp_timeout            = 10
  snmp_user_name          = "admin"
  snmp_version            = "v3"
  type                    = "NETWORK_DEVICE"
  user_name               = "admin"
}
