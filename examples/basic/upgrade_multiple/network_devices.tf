data "catalystcenter_network_devices" "this" {
}

data "catalystcenter_device_detail" "this" {
  for_each = local.devices_raw

  id = each.key
}

locals {
  devices_raw = {
    for v in data.catalystcenter_network_devices.this.devices : v.id => {
      management_ip_address = v.management_ip_address
      management_state      = v.management_state
      platform_id           = v.platform_id
      role                  = v.role
      software_type         = v.software_type
      hostname              = v.hostname
    }
    if v.management_state == "Managed"
  }

  devices = {
    for k, v in local.devices_raw : k => {
      location              = data.catalystcenter_device_detail.this[k].location
      management_ip_address = v.management_ip_address
      platform_id           = v.platform_id
      role                  = v.role
      software_type         = v.software_type
      hostname              = v.hostname
    }
  }
}
