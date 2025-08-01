---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_fabric_port_assignments Resource - terraform-provider-catalystcenter"
subcategory: "SDA"
description: |-
  Manages port assignments in SD-Access fabric.
---

# catalystcenter_fabric_port_assignments (Resource)

Manages port assignments in SD-Access fabric.

## Example Usage

```terraform
resource "catalystcenter_fabric_port_assignments" "example" {
  fabric_id         = "e02d9911-b0a7-435b-bb46-079d877d7b3e"
  network_device_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b"
  port_assignments = [
    {
      fabric_id                  = "c4b85bb2-ce3f-4db9-a32b-e439a388ac2f"
      network_device_id          = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
      interface_name             = "GigabitEthernet1/0/2"
      connected_device_type      = "USER_DEVICE"
      data_vlan_name             = "DATA_VLAN"
      voice_vlan_name            = "VOICE_VLAN"
      authenticate_template_name = "No Authentication"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fabric_id` (String) ID of the fabric the device is assigned to
- `network_device_id` (String) Network device ID of the port assignment
- `port_assignments` (Attributes Set) List of port assignments in SD-Access fabric (see [below for nested schema](#nestedatt--port_assignments))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--port_assignments"></a>
### Nested Schema for `port_assignments`

Required:

- `connected_device_type` (String) Connected device type of the port assignment
  - Choices: `USER_DEVICE`, `ACCESS_POINT`, `TRUNKING_DEVICE`, `AUTHENTICATOR_SWITCH`, `SUPPLICANT_BASED_EXTENDED_NODE`
- `fabric_id` (String) ID of the fabric the device is assigned to
- `interface_name` (String) Interface name of the port assignment
- `network_device_id` (String) Network device ID of the port assignment

Optional:

- `authenticate_template_name` (String) Authenticate template name of the port assignment
  - Choices: `No Authentication`, `Open Authentication`, `Closed Authentication`, `Low Impact`
- `data_vlan_name` (String) Data VLAN name of the port assignment
- `interface_description` (String) Interface description of the port assignment
- `security_group_name` (String) Security group name of the port assignment
- `voice_vlan_name` (String) Voice VLAN name of the port assignment

Read-Only:

- `id` (String) ID of the port assignment

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_fabric_port_assignments.example "<fabric_id>,<network_device_id>"
```
