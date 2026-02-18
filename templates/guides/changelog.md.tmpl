---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## 0.5.3

- Fix `catalystcenter_assign_credentials` resource with multiple improvements: automatically retry when encountering "Global Settings Save is in progress" error (NCND00010) by waiting 15 seconds and retrying once; include all credential types in request payload even when null to resolve API error when assigning credentials at global level
- Fix `catalystcenter_deploy_template` resource to properly handle updates to composite template `member_template_deployment_info`, including `redeploy` parameter changes and template variable modifications. Previously, changes to member template configurations were ignored during updates, causing composite templates to not redeploy when member attributes changed
- **CRITICAL FIX**: Prevent DELETE requests to base endpoint when device ID is empty or null in `catalystcenter_provision_device` and `catalystcenter_provision_devices` resources. Previously, empty IDs could result in DELETE requests to `/sda/provisionDevices/`, which the API could interpret as deleting all provisioned devices. The provider now skips delete operations for resources with empty IDs and fails provisioning if device ID cannot be retrieved
- BREAKING CHANGE: Update `auth_type` enum value from `OPEN_SECURED` to `OPEN-SECURED` in `catalystcenter_wireless_ssid` resource to match API specification

## 0.5.2

- Add `catalystcenter_dot11be_profile` resource and data source
- Add `catalystcenter_ap_profile` resource and data source

## 0.5.1

- Fix `catalystcenter_template_version` resource to properly replace the resource when `template_id` is changed
- Update `go-catalystcenter` SDK dependency to `v0.1.11` to support automatic reauthentication of expired tokens
- Add workaround for inconsistent Catalyst Center API behavior in `catalystcenter_ip_pools` data source. The provider retries with an alternative endpoint when the primary returns 404 or 500 errors

## 0.5.0

- Add workaround for inconsistent Catalyst Center API behavior in `catalystcenter_ip_pool` and `catalystcenter_ip_pool_reservation` resources. The provider retries with an alternative endpoint when the primary returns 404 or 500 errors
- BREAKING CHANGE: Modify `params` and `selection_values` attributes in `catalystcenter_deploy_template` and `catalystcenter_template` resources from `Map of String` to `Map of List of String` to support both single and multiple values. Single-element lists are automatically converted to strings when sent to the API for backward compatibility with existing templates
- Add `catalystcenter_global_credential_netconf` resource and data source for managing NETCONF device credentials
- Fix DELETE operation for `catalystcenter_aaa_settings`, `catalystcenter_banner_settings`, `catalystcenter_dhcp_settings`, `catalystcenter_dns_settings`, `catalystcenter_ntp_settings`, `catalystcenter_telemetry_settings`, and `catalystcenter_timezone_settings` resources to send properly structured empty bodies instead of `{}`, resolving NCND* empty input errors
- Fix issue where failed API task errors were silently ignored during resource creation, causing resources to be saved in state with null IDs and leading to idempotency issues on subsequent applies
- Fix DELETE operation for `catalystcenter_network_profile_for_sites_assignments` resource to use individual site unassignment endpoint instead of bulk endpoint, avoiding API limitations and improving reliability when managing large numbers of site assignments
- BREAKING CHANGE: Convert `catalystcenter_areas`, `catalystcenter_buildings`, and `catalystcenter_floors` resources from List to Map type for clearer resource identification using `parent_name_hierarchy/name` keys instead of position-based indexing

## 0.4.7

- Add `catalystcenter_areas`, `catalystcenter_buildings`, `catalystcenter_floors` bulk resources for efficiently managing multiple sites at scale with a single resource declaration
- Add read cache to `catalystcenter_area`, `catalystcenter_building`, `catalystcenter_floor`, `catalystcenter_ip_pool`, and `catalystcenter_ip_pool_reservation` resources to improve performance with large state files
- Add `data_source` and `import` to `catalystcenter_virtual_network_to_fabric_site` resource
- Fix `catalystcenter_discovery` resource to handle API issue where discovery type response is identical for Range and Multi Range types by making `discovery_type` write-only

## 0.4.6

- Fix `catalystcenter_credentials_cli` resource update operation to use correct API endpoint and payload structure
- Fix `catalystcenter_provision_devices` resource to properly populate missing IDs during update operation

## 0.4.5

- Fix `catalystcenter_deploy_template` resource to properly handle `redeploy` attribute when set to `ALWAYS` by creating drift on every apply

## 0.4.4

- Fix `catalystcenter_discovery` resource to use discovery ID for read operations and avoid waiting for discovery completion during creation
- Update the `redeploy` attribute in the `catalystcenter_deploy_template` resource to controls when a template is redeployed. `ALWAYS` redeploys it on every Terraform apply, `ON_CHANGE` redeploys only when the template content changes, and `NEVER` prevents redeployment.
- Round `catalystcenter_floor` resource's length, width and height attributes to 3 decimal places
- Fix issue with `catalystcenter_fabric_multicast_virtual_networks` when `rp_device_location` is `EXTERNAL`
- Fix issue with `catalystcenter_virtual_network_to_fabric_site` resource destroy method
- Fix issue with the destroy operation in the `catalystcenter_provision_devices` resource
- Add `catalystcenter_integrate_ise` resource to accept Cisco ISE server certificate for Cisco ISE server integration

## 0.4.3

- Add `catalystcenter_fabric_devices` resource and data_source to manage multiple fabric devices under same fabric site within single resource 
- Add `catalystcenter_fabric_multicast_virtual_networks` resource and data_source to manage multiple multicast virtual networks under same fabric site within single resource
- Add `catalystcenter_fabric_multicast_virtual_network` resource and data_source
- Add `catalystcenter_extranet_policy` resource and data_source

## 0.4.2

- Add `catalystcenter_fabric_ewlc` resource to enable fabric embedded wireless capabilities on switch devices

## 0.4.1

- Add `catalystcenter_provision_access_points` resource to provision multiple access points within single resource
- Fix issue with `catalystcenter_ip_pool_reservation` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/300)

## 0.4.0

- Enhancement to `catalystcenter_provision_device` and `catalystcenter_provision_devices`: during resource creation, if a device is already provisioned, it will now be reprovisioned via a PUT request instead of returning an error
- Add `apply_pending_fabric_events` resource to apply all pending fabric events
- BREAKING CHANGE: Rename resource `catalystcenter_fabric_provision_device` to `catalystcenter_provision_device`
- BREAKING CHANGE: Modified `catalystcenter_assign_credentials` to use the `/dna/intent/api/v1/sites/%v/deviceCredentials` API endpoint instead of `/dna/intent/api/v2/credential-to-site`
- BREAKING CHANGE: Modified `catalystcenter_ip_pool_reservation` to use the `/dna/intent/api/v1/ipam/siteIpAddressPools` API endpoint instead of `/dna/intent/api/v1/reserve-ip-subpool`
- BREAKING CHANGE: Modified `catalystcenter_ip_pool` to use the `/dna/intent/api/v1/ipam/globalIpAddressPools` API endpoint instead of `/api/v2/ippool`
- Fix issue with `catalystcenter_fabric_l2_handoff` idempotency, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/296)
- Add `catalystcenter_provision_devices` resource and data source to provision multiple devices within single resource
- BREAKING CHANGE: Modified `catalystcenter_area` to use the `/dna/intent/api/v1/areas` API endpoint instead of `/dna/intent/api/v1/site`
- BREAKING CHANGE: Modified `catalystcenter_building` to use the `/dna/intent/api/v2/buildings` API endpoint instead of `/dna/intent/api/v1/site`
- Removed deprecated API endpoints from `catalystcenter_image` resource
- BREAKING CHANGE: Removed `catalystcenter_fabric_virtual_network` (deprecated in CC 2.3.7.9). Use `catalystcenter_fabric_l3_virtual_network` instead
- BREAKING CHANGE: Removed `catalystcenter_associate_site_to_network_profile` (deprecated in CC 2.3.7.6). Use `catalystcenter_network_profile_for_sites_assignments` instead
- BREAKING CHANGE: Removed `catalystcenter_network` (deprecated in CC 2.3.7.6). Use `catalystcenter_ntp_settings`, `catalystcenter_dhcp_settings`, `catalystcenter_dns_settings`, `catalystcenter_timezone_settings`, `catalystcenter_banner_settings`, `catalystcenter_telemetry_settings` or `catalystcenter_aaa_settings` instead
- BREAKING CHANGE: Removed `catalystcenter_wireless_enterprise_ssid` (deprecated in CC 2.3.7.6). Use `catalystcenter_wireless_ssid` instead
- BREAKING CHANGE: Removed `catalystcenter_fabric_authentication_profile`. Use `catalystcenter_update_authentication_profile` instead
- BREAKING CHANGE: Removed `catalystcenter_virtual_network_ip_pool` (deprecated in CC 2.3.7.9). Use `catalystcenter_anycast_gateway` or `catalystcenter_anycast_gateways` instead
- BREAKING CHANGE: Modified `catalystcenter_floor` to use the `/dna/intent/api/v2/floors` API endpoint instead of `/dna/intent/api/v1/site`

## 0.3.5

- Change `device_types` attributes in `catalystcenter_template` resource from `List` to `Set`
- Add `data_source` and `import` to `catalystcenter_assign_device_to_site` resource
- Change `border_types` and `device_roles` attributes in `catalystcenter_fabric_device` resource from `List` to `Set`
- Replace `catalystcenter_deploy_template` resource when `params` attribute is changed

## 0.3.4

- Change in `catalystcenter_deploy_template` resource to throw a warning instead of an error when template deployment fails on a device
- Add support for handling more than the API’s maximum number of objects supported in a single request, allowing larger configurations under a single `catalystcenter_fabric_l3_handoff_ip_transits` resource
- Add support for handling more than the API’s maximum number of objects supported in a single request, allowing larger configurations under a single `catalystcenter_anycast_gateways` resource

## 0.3.3

- Change the attribute type of `mappings` in `catalystcenter_fabric_vlan_to_ssid` from `List` to `Set`
- Fix issue with `catalystcenter_deploy_template` where terraform did not properly report an error when a template failed to deploy
- Add `additional_interfaces` and `ap_zones` attributes to `catalystcenter_wireless_profile`
- `catalystcenter_wireless_ssid` change type of attribute `openSsid` to `string`

## 0.3.2

- Add `catalystcenter_template_versions` data source to retrieve a list of all versions of a specified template
- Modify `catalystcenter_deploy_template` to use `deployment_id` as resource `id`
- Add `catalystcenter_anycast_gateways` resource and data source
- BREAKING CHANGE: Rename `preferred_ip_method` attribute of `catalystcenter_discovery` resource to `preferred_mgmt_ip_method`

## 0.3.1

- Add workaround for API defect in `catalystcenter_aaa_settings`: using `TACACS` incorrectly returns `RADIUS` in API response
- Allow the destroy operation on `catalystcenter_assign_device_to_site` if the device has been removed from the inventory

## 0.3.0

- Change data_source and import to use `wireless_profile_name` instead of `id` in `catalystcenter_wireless_profile`
- Add `allow_existing_on_create` attribute to the provider configuration to allow existing objects in Catalyst Center to be managed. This is en experimental feature (use at your own risk)
- Fix issue with `pre_auth_acl_description` in `catalystcenter_update_authentication_profile` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/258)

## 0.2.16

- Add `catalystcenter_fabric_sites` data source to retrieve list of all fabric sites

## 0.2.15

- BREAKING CHANGE: Modified `catalystcenter_virtual_network_to_fabric_site` to use the `/dna/intent/api/v1/sda/layer3VirtualNetworks` API endpoint instead of `/dna/intent/api/v1/business/sda/virtual-network` and remove data_source and import
- Add a debug flag to enable starting a provider in debug mode
- Fix idempotency issue with `catalystcenter_telemetry_settings`

## 0.2.14

- Add data source and import support to `catalystcenter_virtual_network_to_fabric_site`
- Fix create and destroy operation for `catalystcenter_fabric_virtual_network` when using `INFRA_VN` or `DEFAULT_VN`
- Add `catalystcenter_virtual_network_to_fabric_site` data source and import

## 0.2.13

- Fix destroy operation for `catalystcenter_fabric_l3_virtual_network` when using `INFRA_VN` or `DEFAULT_VN`

## 0.2.12

- BREAKING CHANGE: Modified `catalystcenter_wireless_device_provision` resource to use the `/dna/intent/api/v1/wirelessControllers/%v/provision` API endpoint instead of `/dna/intent/api/v1/wireless/provision`
- Add `assign_managed_ap_locations` resource
- Add `reprovision` attribute to `catalystcenter_wireless_device_provision`
- Add `catalystcenter_ip_pools` data source to retrieve list of all global pools

## 0.2.11

- Add mutex to `catalystcenter_assign_device_to_site` resource to prevent multiple concurrent operations
- Add `catalystcenter_sites` data source to retrieve list of all sites

## 0.2.10

- Add enhancement to `catalystcenter_pnp_device` resource, to skip error during DELETE operation, if pnp device was provisioned
- BREAKING CHANGE: Modified `catalystcenter_wireless_rf_profile` resource and data_source to use the `/dna/intent/api/v1/wirelessSettings/rfProfiles` API endpoint instead of `/dna/intent/api/v1/wireless/rf-profile`
- Fix issue with inconsistent state in `catalystcenter_fabric_port_assignments` if static port assignment is cleared out-of-band, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/234)
- Add `catalystcenter_network_profile_for_sites_assignments` resource and data source
- Fix issue with `id` not computed with import in `catalystcenter_xxx_settings` resources, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/232)

## 0.2.9

- Modify `catalystcenter_pnp_device` resource and skip error if device already exists and add to terraform state
- Add 400 406 and 500 status codes for handling manually deleted objects 
- Fix issue with list order of `ssid_details` attribute in `catalystcenter_wireless_profile` resource

## 0.2.8

- Remove `authentication_profile_id` attribute from `catalystcenter_update_authentication_profile` and allow modifications to global authentication profiles
- Add support for `INFRA_VN` and `DEFAULT_VN` under `catalystcenter_fabric_l3_virtual_network` resource
- Add `copying_config` attribute to `catalystcenter_deploy_template` resource

## 0.2.7

- Add import and data source to `catalystcenter_update_authentication_profile` resource
- Remove default value from `interface_name` in the `catalystcenter_wireless_profile` resource 
- Add import to `catalystcenter_assign_templates_to_tag` resource
- Add `catalystcenter_assign_devices_to_tag` resource and data source

## 0.2.6

- BREAKING CHANGE: Add `network_device_id` attribute to `catalystcenter_wireless_device_provision` resource to fix issue with terraform destroy operation
- Add `catalystcenter_fabric_zone` resource and data source
- Fix issue with updating `catalystcenter_anycast_gateway` while using autogenerated vlan name
- Add `catalystcenter_update_authentication_profile` resource and data_source
- Add `catalystcenter_wireless_interface` resource and data source. This resource only works with Catalyst Center version 2.3.7.9+
- Fix the issue with ID assignment when using multiple interfaces with the same `internal_vlan_id` and `external_vlan_id` in the `catalystcenter_fabric_l2_handoff` resource
- Add the `serial_number` query parameter to the `catalystcenter_pnp_device` data source to fix the issue with result pagination

## 0.2.5

- Add `catalystcenter_fabric_l3_handoff_sda_transit` resource and data source
- Modify `catalystcenter_lan_automation` to use V2 LAN Automation Start API, which supports optional auto-stop processing feature based on the provided timeout or a specific device list, or both.
- Add `catalystcenter_assign_device_to_site` resource. This resource only works with Catalyst Center version 2.3.7.9+
- Improve delete operations to treat `404 Not Found` responses as successful, ensuring idempotent behavior when resources are already removed

## 0.2.4

- Add `redeploy` attribute to `catalystcenter_deploy_template`
- Add `catalystcenter_fabric_l3_handoff_ip_transits` resource and data source
- Add `catalystcenter_update_device_management_address` resource
- Fix issue with terraform not detecting delta in `catalystcenter_device` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/200)
- Make all attributes optional in `catalystcenter_device` resource to match API documentation
- Fix an issue in the `catalystcenter_fabric_l3_virtual_network` resource where modifying the `fabric_ids` attribute would incorrectly recreate L3 Virtual Network
- Add mutex to resources to prevent multiple concurrent operations (such as Create, Update, or Delete) on the same resource instance or API endpoint, in order to bypass the sequential API execution limitation

## 0.2.3

- BREAKING CHANGE: Rename resource `catalystcenter_fabric_port_assignment` to `catalystcenter_fabric_port_assignments` and fix issue with adding/removing ports to/from `catalystcenter_fabric_port_assignments` resource

## 0.2.2

- Fix issue with isGroupBasedPolicyEnforcementEnabled flag missing in `catalystcenter_anycast_gateway` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/190)
- Remove query_param from anchoredSiteId and fabricIds in `catalystcenter_fabric_l3_virtual_network`

## 0.2.1

- Add device unreachability warning to 'update' method in `catalystcenter_fabric_l3_handoff_ip_transit`, `catalystcenter_anycast_gateway`, `catalystcenter_fabric_l2_handoff`, `catalystcenter_fabric_device`, `catalystcenter_fabric_port_assignment` and `catalystcenter_provision_device`
- Add `reprovision` attribute to `catalystcenter_fabric_provision_device`

## 0.2.0

- Fix issue with security group name in `catalystcenter_anycast_gateway` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/179)
- Add `catalystcenter_banner_settings` resource and data source
- Add `catalystcenter_dhcp_settings` resource and data source
- Add `catalystcenter_dns_settings` resource and data source
- Add `catalystcenter_ntp_settings` resource and data source
- Add `catalystcenter_telemetry_settings` resource and data source
- Add `catalystcenter_timezone_settings` resource and data source
- Changed default timeout of asynchronous tasks to 180 seconds

## 0.1.19

- Add update support for individual port assignments with `catalystcenter_fabric_port_assignment` resource
- Add support to auto-generate VLAN name in `catalystcenter_anycast_gateway` resource
- Fix issue with tags removal from templates in `catalystcenter_assign_templates_to_tag` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/165)

## 0.1.18

- Add `catalystcenter_fabric_l3_virtual_network` resource and data source
- Add `Free Space` allowed value to `rf_model` attribute of `catalystcenter_floor` resource
- Add device unreachability warning to 'delete' implementation of `catalystcenter_fabric_l3_handoff_ip_transit`, `catalystcenter_anycast_gateway`, `catalystcenter_fabric_l2_handoff`, `catalystcenter_fabric_device`, `catalystcenter_fabric_port_assignment` and `catalystcenter_provision_device` resources
- Fix issue with `catalystcenter_ip_pool_reservation` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/156)

## 0.1.17

- BREAKING CHANGE: Rename `auth_key_sae_ext_plus_tf` attribute of `catalystcenter_wireless_ssid` resource to `auth_key_sae_ext_plus_ft`
- BREAKING CHANGE: Rename `auth_key8021x_plus_tf` attribute of `catalystcenter_wireless_ssid` resource to `auth_key8021x_plus_ft`
- Add device unreachability warning to `catalystcenter_anycast_gateway`, `catalystcenter_fabric_l2_handoff`, `catalystcenter_fabric_device`, `catalystcenter_fabric_port_assignment` and `catalystcenter_provision_device` resources
- Add `auth_key_psk` and `auth_key_psk_plus_ft` attributes to `catalystcenter_wireless_ssid` resource and data source

## 0.1.16

- Add device unreachability warning to `catalystcenter_fabric_l3_handoff_ip_transit` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/150)
- Add `catalystcenter_fabric_l2_virtual_network` resource and data source
- Remove resource specific asynchronous timeout configuration and use timeout from provider configuration
- Modify `catalystcenter_deploy_template` resource to support deploying composite templates
- Modify `catalystcenter_template_version` resource to use versioned template id as `id` and remove data source

## 0.1.15

- Fix issue in `catalystcenter_fabric_l3_handoff_ip_transit` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/146)

## 0.1.14

- Modify `catalystcenter_deploy_template` resource to support `resource_params`
- Add `catalystcenter_aaa_settings` resource and data source

## 0.1.13

- Modify `anycast_gateway` resource to support pool_type `FABRIC_AP`
- Add composite templates support to `catalystcenter_template` resource
- Fix issue with catalystcenter_ip_pool forces replacement on `catalystcenter_ip_pool` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/125)

## 0.1.12

- Change default timeout for asynchronous operations to 60 seconds and introduce provider attribute `max_timeout` to set a custom timeout

## 0.1.11

- Add `catalystcenter_ fabric_vlan_to_ssid` resource and data source
- Add `catalystcenter_wireless_ssid` resource and data source
- Add `catalystcenter_site` data source
- Add `catalystcenter_fabric_port_assignment` resource and data source
- BREAKING CHANGE: Replace `catalystcenter_peer_transit_network` with `catalystcenter_transit_network` resource and data source to use `/dna/intent/api/v1/sda/transitNetworks` API endpoint, this resource now only works with Catalyst Center version 2.3.7.6+
- Add `catalystcenter_authentication_policy_server` resource and data source
- Fix issue with import of `catalystcenter_ip_pool_reservation` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/122)

## 0.1.10

- Add `catalystcenter_fabric_l2_handoff` resource and data source
- BREAKING CHANGE: Modified `catalystcenter_wireless_profile` resource to use `/intent/api/v1/wirelessProfiles` API endpoint, this resource now only works with Catalyst Center version 2.3.7.6+
- Add `catalystcenter_wireless_device_provision` resource
- Add `catalystcenter_fabric_provision_device` resource and data source
- Add `catalystcenter_assign_templates_to_tag` resource and data source
- Add `catalystcenter_tag` resource and data source
- Add `catalystcenter_pnp_import_devices` resource
- Add `catalystcenter_fabric_device` resource and data source, this resource now only works with Catalyst Center version 2.3.7.5+
- Add `catalystcenter_fabric_l3_handoff_ip_transit` resource and data source
- Add `transitPeerNetworkId` as `id` to `catalystcenter_transit_peer_network` resource
- Add `catalystcenter_anycast_gateway` resource and data source, this resource now only works with Catalyst Center version 2.3.7.5+
- BREAKING CHANGE: Modified `catalystcenter_fabric_site` resource to use `/dna/intent/api/v1/sda/fabricSites` API endpoint, this resource now only works with Catalyst Center version 2.3.7.5+
- Fix issue with mandatory attributes in `catalystcenter_transit_peer_network` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/92)
- BREAKING CHANGE: Fix `catalystcenter_ip_pool` update if more than 25 pools are registered
- BREAKING CHANGE: Rename `radio_type_a_power_treshold_v1` attribute of `catalystcenter_wireless_rf_profile` resource to `radio_type_a_power_threshold_v1`
- BREAKING CHANGE: Rename `radio_type_b_power_treshold_v1` attribute of `catalystcenter_wireless_rf_profile` resource to `radio_type_b_power_threshold_v1`
- BREAKING CHANGE: Rename `radio_type_c_power_treshold_v1` attribute of `catalystcenter_wireless_rf_profile` resource to `radio_type_c_power_threshold_v1`

## 0.1.9

- Fix `mfp_client_protection` attribute validation of `catalystcenter_wireless_enterprise_ssid` resource, [link](https://github.com/CiscoDevNet/terraform-provider-catalystcenter/issues/74)
- BREAKING CHANGE: Modified `wireless_profile` resource to use `/dna/intent/api/v2/wireless/profile` API endpoint, this resource now only works with Catalyst Center version 2.3.7.5+
- Remove default value from ipv6_address_space attribute in `catalystcenter_ip_pool_reservation` resource

## 0.1.8

- Fix import functionality for resources with multipart keys, e.g. `catalystcenter_ip_pool_reservation`

## 0.1.7

- BREAKING CHANGE: Add support for multiple CLI and Day0 templates to `catalystcenter_network_profile` resource and data source

## 0.1.6

- Fix issue with unencoded characters ending up in parts of the URL
- Use `set` type for list attributes with primitive values
- Add `catalystcenter_pnp_config_preview` resource
- Add `catalystcenter_virtual_network_to_fabric_site` resource
- Add `catalystcenter_wireless_profile` resource
- Add `catalystcenter_virtual_network_ip_pool` resource
- Add `floor_number` attribute to `catalystcenter_floor` resource and data source

## 0.1.5

- Add `catalystcenter_fabric_site` resource and data source
- Add `catalystcenter_template_version` resource and data source
- Add `catalystcenter_fabric_authentication_profile` resource and data source
- Fix handling of validation warnings when creating templates
- Make `country`, `address`, `latitude` and `longitude` attributes of `catalystcenter_building` resource optional
- Make `selection_type` attribute of `catalystcenter_template` resource optional
- Add `catalystcenter_deploy_template` resource
- Add `catalystcenter_lan_automation` resource and data source
- Add `catalystcenter_wireless_enterprise_ssid` resource and data source
- Add `catalystcenter_wireless_rf_profile` resource and data source
- Add `catalystcenter_transit_peer_network` resource and data source
- Add `catalystcenter_fabric_virtual_network` resource and data source

## 0.1.4

- Fix import operation of resources
- Add `catalystcenter_device_detail` data source
- Add `catalystcenter_device_discovery` resource and data source
- Add `catalystcenter_device_network_devices` data source
- Add `catalystcenter_device_image` resource
- Add `catalystcenter_device_image_activation` resource
- Add `catalystcenter_device_image_distribution` resource
- Add `catalystcenter_device_associate_site_to_network_profile` resource

## 0.1.3

- Add `catalystcenter_device_role` resource
- BREAKING CHANGE: Rename `catalystcenter_device` resource and data source to `catalystcenter_pnp_device`
- BREAKING CHANGE: Rename `catalystcenter_device_claim_site` resource and data source to `catalystcenter_pnp_device_claim_site`
- Add `catalystcenter_device` resource
- Add `catalystcenter_role` resource and data source
- Add `catalystcenter_user` resource and data source

## 0.1.2

- Fix issue with `catalystcenter_ip_pool_reservation` resource and multiple IP pools under a single site
- Add `catalystcenter_device_claim_site` resource
- Add `catalystcenter_device` resource and data source
- Add `name` query option to `catalystcenter_ip_pool` data source
- Add `description` query option to `catalystcenter_credentials_cli`, `catalystcenter_credentials_https_write`, `catalystcenter_credentials_https_read`, `catalystcenter_credentials_snmpv2_read`, `catalystcenter_credentials_snmpv2_write` and `catalystcenter_credentials_snmpv3` data source
- Add `name` query option to `catalystcenter_project` data source
- Add `name` query option to `catalystcenter_template` data source
- Add `name` query option to `catalystcenter_area` and `catalystcenter_building` data source
- Add `name` query option to `catalystcenter_network_profile` data source

## 0.1.1

- BREAKING CHANGE: Rename `https_read` attribute of `catalystcenter_assign_credentials` to `https_read_id`
- BREAKING CHANGE: Rename `https_write` attribute of `catalystcenter_assign_credentials` to `https_write_id`

## 0.1.0

- Initial release

