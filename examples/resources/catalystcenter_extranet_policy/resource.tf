resource "catalystcenter_extranet_policy" "example" {
  extranet_policy_name             = "MyExtranetPolicy"
  fabric_ids                       = ["5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"]
  provider_virtual_network_name    = "8b6830b4-bf36-43d3-a97c-d94dc7d4715d"
  subscriber_virtual_network_names = ["afc4ed30-3c57-4992-a41a-97d3f48bf071"]
}
