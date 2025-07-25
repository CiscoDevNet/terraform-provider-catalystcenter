---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_fabric_site Resource - terraform-provider-catalystcenter"
subcategory: "SDA"
description: |-
  Manages Fabric Sites
---

# catalystcenter_fabric_site (Resource)

Manages Fabric Sites

## Example Usage

```terraform
resource "catalystcenter_fabric_site" "example" {
  site_id                     = "8e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  authentication_profile_name = "No Authentication"
  pub_sub_enabled             = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `authentication_profile_name` (String) Authentication profile used for this fabric
  - Choices: `Closed Authentication`, `Low Impact`, `No Authentication`, `Open Authentication`
- `pub_sub_enabled` (Boolean) Specifies whether this fabric site will use pub/sub for control nodes
- `site_id` (String) ID of the network hierarchy

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_fabric_site.example "<id>"
```
