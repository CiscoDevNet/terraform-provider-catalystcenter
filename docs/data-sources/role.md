---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_role Data Source - terraform-provider-catalystcenter"
subcategory: "System"
description: |-
  This data source can read the Role.
---

# catalystcenter_role (Data Source)

This data source can read the Role.

## Example Usage

```terraform
data "catalystcenter_role" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the role

### Read-Only

- `description` (String) The description of the role
- `resource_types` (Attributes List) The resource types of the role (see [below for nested schema](#nestedatt--resource_types))

<a id="nestedatt--resource_types"></a>
### Nested Schema for `resource_types`

Read-Only:

- `operations` (Set of String) List of operations allowed for the application. Possible values are `gRead`, `gWrite`, `gUpdate`, `gDelete`, or some combination of these.
- `type` (String) Name of the application
