---
subcategory: "Guides"
page_title: "Image Management"
description: |-
    Image Management
---

# Image Management

This example demonstrates how the provider can be used to upload, distribute and activate software images on devices managed by Catalyst Center. The full example can be found here: [https://github.com/CiscoDevNet/terraform-provider-catalystcenter/tree/main/examples/basic/upgrade_multiple](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/tree/main/examples/basic/upgrade_multiple)

First of all we need to add the necessary provider configuration to the Terraform configuration file:

```hcl
terraform {
  required_providers {
    catalystcenter = {
      source = "CiscoDevNet/catalystcenter"
    }
  }
}

provider "catalystcenter" {
  username = "admin"
  password = "password"
  url      = "https://10.1.1.1"
}
```

Next we would need to upload the image(s) to Catalyst Center:

```hcl
locals {
  images_uploaded = toset([
    "cat9k_iosxe.17.12.02.SPA.bin",
    "cat9k_iosxe.17.12.02prd4.SPA.bin",
    "C9800-SW-iosxe-wlc.17.12.02prd4.SPA.bin",
  ])
}

resource "catalystcenter_image" "this" {
  for_each = local.images_uploaded

  name        = basename(each.key)
  source_path = each.key
}
```

After uploading the images, we need to build a list of devices we want to upgrade:

```hcl
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
```

And finally, we can activate the image on the devices:

```hcl
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
```
