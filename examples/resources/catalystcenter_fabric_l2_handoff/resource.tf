resource "catalystcenter_fabric_l2_handoff" "example" {
  network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  fabric_id         = "c4b85bb2-ce3f-4db9-a32b-e439a388ac2f"
  interface_name    = "GigabitEthernet1/0/4"
  internal_vlan_id  = 300
  external_vlan_id  = 400
}
