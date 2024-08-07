---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_profile Data Source - terraform-provider-catalystcenter"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless Profile.
---

# catalystcenter_wireless_profile (Data Source)

This data source can read the Wireless Profile.

## Example Usage

```terraform
data "catalystcenter_wireless_profile" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `ssid_details` (Attributes List) SSID Details (see [below for nested schema](#nestedatt--ssid_details))
- `wireless_profile_name` (String) Wireless Network Profile Name

<a id="nestedatt--ssid_details"></a>
### Nested Schema for `ssid_details`

Read-Only:

- `dot11be_profile_id` (String) 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured
- `enable_fabric` (Boolean) True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
- `enable_flex_connect` (Boolean) True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
- `interface_name` (String) Interface Name
- `local_to_vlan` (Number) Local To Vlan Id
- `ssid_name` (String) SSID Name
- `wlan_profile_name` (String) WLAN Profile Name