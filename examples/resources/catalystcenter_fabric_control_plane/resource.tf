resource "catalystcenter_fabric_control_plane" "example" {
  device_management_ip_address = "192.168.10.1"
  site_name_hierarchy          = "Global/Area1"
  route_distribution_protocol  = "LISP_BGP"
}
