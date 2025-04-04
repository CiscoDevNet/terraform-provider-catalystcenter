resource "catalystcenter_fabric_port_assignments" "example" {
  fabric_id         = "e02d9911-b0a7-435b-bb46-079d877d7b3e"
  network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b"
  port_assignments = [
    {
      fabric_id                  = "c4b85bb2-ce3f-4db9-a32b-e439a388ac2f"
      network_device_id          = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
      interface_name             = "GigabitEthernet1/0/2"
      connected_device_type      = "USER_DEVICE"
      data_vlan_name             = "DATA_VLAN"
      voice_vlan_name            = "VOICE_VLAN"
      authenticate_template_name = "No Authentication"
    }
  ]
}
