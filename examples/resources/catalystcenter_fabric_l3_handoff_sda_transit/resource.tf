resource "catalystcenter_fabric_l3_handoff_sda_transit" "example" {
  network_device_id                 = "a144a086-750c-4af1-ac57-feab33a69851"
  fabric_id                         = "c4b85bb2-ce3f-4db9-a32b-e439a388ac2f"
  transit_network_id                = "d71c847b-e9c2-4f13-928c-223372b72b06"
  affinity_id_prime                 = 100
  affinity_id_decider               = 100
  connected_to_internet             = false
  is_multicast_over_transit_enabled = false
}
