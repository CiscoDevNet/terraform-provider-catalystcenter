---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_pnp_config_preview Resource - terraform-provider-catalystcenter"
subcategory: "Plug and Play"
description: |-
  This resource triggers a preview for site-based Day 0 Configuration. When this resource is destroyed, updated or refreshed, no actions are done either on Catalyst Center or on devices. After creating this resource the config can be previewed in the GUI of Catalyst Center.
---

# catalystcenter_pnp_config_preview (Resource)

This resource triggers a preview for site-based Day 0 Configuration. When this resource is destroyed, updated or refreshed, no actions are done either on Catalyst Center or on devices. After creating this resource the config can be previewed in the GUI of Catalyst Center.

## Example Usage

```terraform
resource "catalystcenter_pnp_config_preview" "example" {
  device_id = "65e422375b4b6e40ef3423f8"
  site_id   = "12345678-1234-1234-1234-123456789012"
  type      = "Default"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device_id` (String) The device ID
- `site_id` (String) The site ID
- `type` (String) The device type
  - Choices: `Default`, `StackSwitch`, `AccessPoint`, `Sensor`, `MobilityExpress`

### Read-Only

- `id` (String) The id of the object
