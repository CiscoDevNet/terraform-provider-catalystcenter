resource "catalystcenter_credentials_snmpv2_read" "example" {
  description               = "My SNMPv2 read credentials"
  read_community_wo         = "community1"
  read_community_wo_version = 1
}
