---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## 0.1.4 (unreleased)

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

