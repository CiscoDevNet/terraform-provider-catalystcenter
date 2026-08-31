resource "catalystcenter_credentials_snmpv3" "example" {
  description                 = "My SNMPv3 credentials"
  username                    = "user1"
  privacy_type                = "AES128"
  privacy_password_wo         = "password1"
  privacy_password_wo_version = 1
  auth_type                   = "SHA"
  auth_password_wo            = "password1"
  auth_password_wo_version    = 1
  snmp_mode                   = "AUTHPRIV"
}
