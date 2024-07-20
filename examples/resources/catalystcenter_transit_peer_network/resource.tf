resource "catalystcenter_transit_peer_network" "example" {
  transit_peer_network_name = "TRANSIT_1"
  transit_peer_network_type = "ip_transit"
  routing_protocol_name     = "BGP"
  autonomous_system_number  = "65010"
}
