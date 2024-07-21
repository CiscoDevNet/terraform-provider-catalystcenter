resource "catalystcenter_fabric_device" "example" {
  network_device_id               = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  fabric_id                       = "c4b85bb2-ce3f-4db9-a32b-e439a388ac2f"
  device_roles                    = ["CONTROL_PLANE_NODE"]
  border_types                    = ["LAYER_3"]
  local_autonomous_system_number  = "65000"
  default_exit                    = true
  import_external_routes          = false
  border_priority                 = 5
  prepend_autonomous_system_count = 1
}
