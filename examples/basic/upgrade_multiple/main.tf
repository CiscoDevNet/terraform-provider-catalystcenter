resource "catalystcenter_image_activation" "this" {
  for_each = {
    for k, v in local.devices : k => v
    if
    v.role == "DISTRIBUTION" &&
    v.platform_id == "C9300-24P" &&
    v.software_type == "IOS-XE" &&
    !startswith(v.location, "Global/USA/SAN JOSE/") &&
    !strcontains(v.hostname, "prod") &&
    !startswith(v.management_ip_address, "10.10.10.100")
  }

  device_uuid                  = each.key
  distribute_if_needed         = true
  activate_lower_image_version = false
  image_uuid_list = [
    catalystcenter_image.this["cat9k_iosxe.17.12.02prd4.SPA.bin"].id,
    catalystcenter_image.this["C9800-SW-iosxe-wlc.17.12.02prd4.SPA.bin"].id,
  ]
}
