resource "catalystcenter_extranet_policy" "example" {
  extranet_policy_name             = "MyExtranetPolicy"
  fabric_ids                       = ["5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"]
  provider_virtual_network_name    = "SDA_VN1"
  subscriber_virtual_network_names = ["SDA_VN2"]
}
