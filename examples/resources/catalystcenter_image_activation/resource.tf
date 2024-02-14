resource "catalystcenter_image_activation" "example" {
  device_uuid                  = "138b3181-f9c5-4271-9292-cf3152ab4d3e"
  image_uuid_list              = [""]
  activate_lower_image_version = true
  distribute_if_needed         = true
}
