## 0.1.2 (unreleased)

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
