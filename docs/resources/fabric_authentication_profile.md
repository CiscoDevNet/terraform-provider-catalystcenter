---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_fabric_authentication_profile Resource - terraform-provider-catalystcenter"
subcategory: "SDA"
description: |-
  This resource can manage a Fabric Authentication Profile.
---

# catalystcenter_fabric_authentication_profile (Resource)

This resource can manage a Fabric Authentication Profile.

## Example Usage

```terraform
resource "catalystcenter_fabric_authentication_profile" "example" {
  site_name_hierarchy          = "Global/Site1"
  authentication_template_name = "No Authentication"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `authentication_template_name` (String) Authentication Template Name
  - Choices: `No Authentication`, `Open Authentication`, `Closed Authentication`, `Low Impact`
- `site_name_hierarchy` (String) Path of SDA Fabric Site

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_fabric_authentication_profile.example "<site_name_hierarchy>"
```
