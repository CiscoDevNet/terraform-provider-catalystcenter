resource "catalystcenter_fabric_vlan_to_ssid" "example" {
  fabric_id = "5e8e3e3e-1b6b-4b6b-8b6b-1b6b4b6b8b6b"
  mappings = [
    {
      vlan_name = "VLAN_1"
      ssid_details = [
        {
          name               = "mySSID1"
          security_group_tag = "Auditors"
        }
      ]
    }
  ]
}
