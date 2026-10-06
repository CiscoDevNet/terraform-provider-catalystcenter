resource "catalystcenter_credentials_snmpv2_write" "example" {
  description                = "My SNMPv2 write credentials"
  write_community_wo         = "community1"
  write_community_wo_version = 1
}
