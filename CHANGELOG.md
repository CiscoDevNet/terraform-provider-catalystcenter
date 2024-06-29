## 0.1.10 (unreleased)

- Add `catalystcenter_anycast_gateway` resource and data source
- BREAKING CHANGE: Modified `fabric_site` resource to use `/dna/intent/api/v1/sda/fabricSites` API endpoint, this resource now only works with Catalyst Center version 2.3.7.5+
- BREAKING CHANGE: Fix `ip_pool` update if more than 25 pools are registered
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
