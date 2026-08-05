resource "catalystcenter_wireless_pre_auth_acl" "example" {
  acl_list_name                = "PreAuth_ACL"
  acl_name                     = "PreAuth_ACL"
  ipv6_acl_enabled             = false
  include_auto_generated_rules = true
  ip_acl_rules = [
    {
      source_address                    = "250.162.252.170"
      source_subnet_mask_or_prefix      = 32
      destination_address               = "250.162.252.171"
      destination_subnet_mask_or_prefix = 32
      source_ports                      = "100"
      destination_ports                 = "100-200"
      protocol                          = "IP"
    }
  ]
}
