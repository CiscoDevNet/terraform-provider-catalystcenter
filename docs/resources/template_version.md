---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_template_version Resource - terraform-provider-catalystcenter"
subcategory: "Templates"
description: |-
  This resource can manage a Template Version.
---

# catalystcenter_template_version (Resource)

This resource can manage a Template Version.

## Example Usage

```terraform
resource "catalystcenter_template_version" "example" {
  template_id = "12345678-1234-1234-1234-123456789012"
  comments    = "New Version"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `comments` (String) Template version comments
- `template_id` (String) UUID of template

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_template_version.example "<template_id>"
```
