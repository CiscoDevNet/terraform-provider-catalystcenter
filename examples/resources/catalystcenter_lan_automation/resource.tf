resource "catalystcenter_lan_automation" "example" {
  discovered_device_site_name_hierarchy = "Global/Area1/Area2"
  primary_device_management_ip_address  = "1.2.3.4"
  peer_device_management_ip_address     = "1.2.3.5"
  primary_device_interface_names        = ["HundredGigE1/0/1"]
  ip_pools = [
    {
      ip_pool_name = "POOL1"
      ip_pool_role = "MAIN_POOL"
    }
  ]
  multicast_enabled        = true
  host_name_prefix         = "TEST"
  host_name_file_id        = "1234"
  isis_domain_password     = "cisco123"
  redistribute_isis_to_bgp = true
}
