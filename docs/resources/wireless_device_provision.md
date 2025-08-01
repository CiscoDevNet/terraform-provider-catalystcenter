---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_device_provision Resource - terraform-provider-catalystcenter"
subcategory: "Wireless"
description: |-
  This resource is used to provision a wireless device. To reprovision a device, set the reprovision attribute to true. The resource will then trigger reprovisioning on every Terraform apply.
---

# catalystcenter_wireless_device_provision (Resource)

This resource is used to provision a wireless device. To reprovision a device, set the `reprovision` attribute to `true`. The resource will then trigger reprovisioning on every Terraform apply.

## Example Usage

```terraform
resource "catalystcenter_wireless_device_provision" "example" {
  network_device_id = "e2e6ae2f-d526-459f-bfdf-3281d74b6dea"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_device_id` (String) Network Device ID

### Optional

- `ap_authorization_list_name` (String) AP Authorization List Name
- `ap_reboot_percentage` (Number) AP Reboot Percentage
- `authorize_mesh_and_non_mesh_access_points` (Boolean) True if Mesh and Non-Mesh Access Points are authorized, else False
- `enable_rolling_ap_upgrade` (Boolean) True if Rolling AP Upgrade is enabled, else False
- `interfaces` (Attributes List) Dynamic Interface Details. The required attributes depend on the device type (see [below for nested schema](#nestedatt--interfaces))
- `reprovision` (Boolean) Flag to indicate whether the device should be reprovisioned. If set to `true`, reprovisioning will be triggered on every Terraform apply
- `skip_ap_provision` (Boolean) True if Skip AP Provision is enabled, else False

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Required:

- `interface_name` (String) Interface Name
- `vlan_id` (Number) VLAN ID
  - Range: `1`-`4094`

Optional:

- `interface_gateway` (String) Interface Gateway
- `interface_ip_address` (String) Interface IP Address
- `interface_netmask` (Number) Interface Netmask In CIDR
  - Range: `1`-`30`
- `lag_or_port_number` (Number) LAG or Port Number
