resource "catalystcenter_routing_profile" "example" {
  name = "Profile1"
  profile_attributes = [
    {
      key   = "nfv.device.name"
      value = "ISR"
      attributes = [
        {
          key   = "nfv.device.deviceType"
          value = "Cisco Catalyst 8000V Edge Software"
        }
      ]
    }
  ]
}
