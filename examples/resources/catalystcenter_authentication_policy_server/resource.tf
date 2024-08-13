resource "catalystcenter_authentication_policy_server" "example" {
  authentication_port      = 1812
  accounting_port          = 1813
  ip_address               = "10.0.0.1"
  pxgrid_enabled           = true
  use_dnac_cert_for_pxgrid = false
  is_ise_enabled           = false
  port                     = 49
  protocol                 = "RADIUS"
  retries                  = 2
  role                     = "secondary"
  shared_secret            = "Cisco123"
  timeout_seconds          = 2
}
