---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_transit_peer_network Data Source - terraform-provider-catalystcenter"
subcategory: "SDA"
description: |-
  This data source can read the Transit Peer Network.
---

# catalystcenter_transit_peer_network (Data Source)

This data source can read the Transit Peer Network.

## Example Usage

```terraform
data "catalystcenter_transit_peer_network" "example" {
  transit_peer_network_name = "TRANSIT_1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `transit_peer_network_name` (String) Transit Peer Network Name

### Read-Only

- `autonomous_system_number` (String) Autonomous System Number
- `id` (String) The id of the object
- `routing_protocol_name` (String) Routing Protocol Name
- `transit_control_plane_settings` (Attributes List) Transit Control Plane Settings info (see [below for nested schema](#nestedatt--transit_control_plane_settings))
- `transit_peer_network_type` (String) Transit Peer Network Type

<a id="nestedatt--transit_control_plane_settings"></a>
### Nested Schema for `transit_control_plane_settings`

Read-Only:

- `device_management_ip_address` (String) Device Management Ip Address of provisioned device
- `site_name_hierarchy` (String) Site Name Hierarchy where device is provisioned
