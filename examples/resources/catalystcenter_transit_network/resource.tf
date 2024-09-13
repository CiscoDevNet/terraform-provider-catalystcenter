resource "catalystcenter_transit_network" "example" {
  name                     = "TRANSIT_1"
  type                     = "IP_BASED_TRANSIT"
  routing_protocol_name    = "BGP"
  autonomous_system_number = "65010"
}
