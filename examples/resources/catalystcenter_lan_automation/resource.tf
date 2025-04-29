resource "catalystcenter_lan_automation" "example" {
  discovered_device_site_name_hierarchy = "Global/Area1/Area2"
  primary_device_management_ip_address  = "1.2.3.4"
  primary_device_interface_names        = ["HundredGigE1/0/1"]
  ip_pools = [
    {
      ip_pool_name = "POOL1"
      ip_pool_role = "MAIN_POOL"
    }
  ]
  multicast_enabled        = true
  host_name_prefix         = "TEST"
  isis_domain_password     = "cisco123"
  redistribute_isis_to_bgp = true
  discovery_level          = 2
  discovery_timeout        = 20
  discovery_devices = [
    {
      device_serial_number         = "FOC12345678"
      device_host_name             = "EDGE01"
      device_site_name_hierarchy   = "Global/Area1/Area2"
      device_management_ip_address = "10.0.0.1"
    }
  ]
}
