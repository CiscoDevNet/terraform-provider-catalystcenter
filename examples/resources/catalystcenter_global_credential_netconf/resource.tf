resource "catalystcenter_global_credential_netconf" "example" {
  netconf_port    = "840"
  comments        = "My NETCONF Credential"
  credential_type = "GLOBAL"
  description     = "NETCONF credential for network devices"
}
