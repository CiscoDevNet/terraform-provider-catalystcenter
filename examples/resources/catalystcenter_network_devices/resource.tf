resource "catalystcenter_network_devices" "example" {
  response = [
    {
      id                    = "e0ba1a00-b69b-45aa-8c13-4cdfb59afe65"
      hostname              = "sw2.ciscotest.com"
      management_ip_address = "10.10.10.1"
      platform_id           = "C9KV-UADP-8P"
      role                  = "CORE"
      software_type         = "IOS-XE"
    }
  ]
}
