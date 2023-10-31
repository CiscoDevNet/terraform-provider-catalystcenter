resource "catalystcenter_network" "example" {
  site_id                          = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  dhcp_servers                     = ["1.2.3.4"]
  domain_name                      = "cisco.com"
  primary_dns_server               = "1.2.3.4"
  secondary_dns_server             = "1.2.3.5"
  syslog_servers                   = ["1.2.3.4"]
  catalyst_center_as_syslog_server = false
  snmp_servers                     = ["1.2.3.4"]
  catalyst_center_as_snmp_server   = false
  netflow_collector                = "1.2.3.4"
  netflow_collector_port           = 1234
  ntp_servers                      = ["1.2.3.4"]
  timezone                         = "Europe/Vienna"
}
