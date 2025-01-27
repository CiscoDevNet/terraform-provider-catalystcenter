## 0.1.20 (unreleased)

- Add `catalystcenter_banner_settings` resource and data source
- Add `catalystcenter_dhcp_settings` resource and data source
- Add `catalystcenter_dns_settings` resource and data source
- Add `catalystcenter_ntp_settings` resource and data source
- Add `catalystcenter_telemetry_settings` resource and data source
- Add `catalystcenter_timezone_settings` resource and data source

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
