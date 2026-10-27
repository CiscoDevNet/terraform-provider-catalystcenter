---
subcategory: "Guides"
page_title: "Managing Secrets"
description: |-
  Configure and rotate write-only secrets without storing their values in Terraform plans or state.
---

# Managing Secrets

Use Terraform 1.11 or later for [write-only arguments](https://developer.hashicorp.com/terraform/language/manage-sensitive-data/write-only). The provider exposes these arguments with a `_wo` suffix and a corresponding version argument. Terraform keeps the version in state, but omits the write-only value from saved plans and state.

The older secret arguments, such as `password`, remain supported. They are marked sensitive to redact normal Terraform output, but their values are still stored in state. Use the write-only alternatives when the value should not be persisted.

## CLI credentials example

Declare secret input variables as both `sensitive` and `ephemeral`. The ephemeral setting also prevents Terraform from storing the input variable values in saved plans; using only `sensitive` does not provide that protection.

```terraform
variable "cli_password" {
  type      = string
  sensitive = true
  ephemeral = true
}

variable "cli_enable_password" {
  type      = string
  sensitive = true
  ephemeral = true
}

resource "catalystcenter_credentials_cli" "example" {
  description                = "Managed CLI credentials"
  username                   = "network-admin"
  password_wo                = var.cli_password
  password_wo_version        = 1
  enable_password_wo         = var.cli_enable_password
  enable_password_wo_version = 1
}
```

Supply these inputs at runtime, for example through `TF_VAR_cli_password` and `TF_VAR_cli_enable_password` environment variables. Supply them again when applying a saved plan, because Terraform does not retain ephemeral inputs in that plan. Write-only arguments do not protect secrets stored in configuration files or other secret sources.

## Rotating credentials

Each write-only secret has its own version. When changing `cli_password`, increment `password_wo_version`, for example from `1` to `2`. When changing the enable password, increment `enable_password_wo_version`. The version is an integer used by Terraform to request an update; it is not a secret and is not sent to Catalyst Center.

Changing only the write-only value produces no Terraform diff, because Terraform has no previous value to compare. Keep all configured write-only values available whenever applying an update to the resource, including an update to an ordinary attribute such as `username`. The provider reads those values from the current configuration when sending the update.

The provider cannot recover write-only values from state or use them to detect secret changes made outside Terraform.

## Other resources

The same argument and version pattern is available for these credentials:

| Resources | Write-only secret arguments |
| --- | --- |
| `catalystcenter_credentials_cli` | `password_wo`, `enable_password_wo` |
| `catalystcenter_credentials_https_read`, `catalystcenter_credentials_https_write` | `password_wo` |
| `catalystcenter_credentials_snmpv2_read` | `read_community_wo` |
| `catalystcenter_credentials_snmpv2_write` | `write_community_wo` |
| `catalystcenter_credentials_snmpv3` | `auth_password_wo`, `privacy_password_wo` |

For each argument in the table, its version argument adds `_version`, for example `auth_password_wo_version`. A version is required when its write-only argument is configured. Use either the older secret argument or its write-only alternative, never both for the same secret.

Write-only arguments are also available on `catalystcenter_aaa_settings`, `catalystcenter_authentication_policy_server`, `catalystcenter_ap_profile`, `catalystcenter_lan_automation`, `catalystcenter_user`, and `catalystcenter_wireless_ssid`. Consult each resource's schema for its supported secret arguments and requirements.

## Template deployment parameters

The [template deployment resource](https://registry.terraform.io/providers/CiscoDevNet/catalystcenter/latest/docs/resources/deploy_template) uses `secret_params` because template parameter names are user-defined. Each entry identifies an existing target using `target_id` or `target_host_name`, and contains:

- `params_wo`: a write-only map from parameter names to lists of secret values.
- `params_wo_versions`: a state-backed map from the same parameter names to positive integer rotation versions.
- `member_template_id`: the matching member template ID when deploying a composite template member.

Keep ordinary parameters in the target or member's `params` map. A parameter cannot appear in both its ordinary and secret maps. State retains the target and member identifiers and the version map; it does not retain the values in `params_wo`.

Increment a parameter's version when rotating its secret. Deployment still follows the resource's `redeploy` policy: `ON_CHANGE` deploys when tracked configuration changes, `ALWAYS` deploys on every apply, and `NEVER` suppresses deployment. Supply all current secrets whenever the target is redeployed, including after an ordinary parameter changes. Removing a secret parameter from Terraform does not erase a password previously configured on a device.

## Migrating existing secrets

An upgrade alone does not require switching existing secret arguments. To migrate a credential, keep its resource address, remove the older argument such as `password`, and configure `password_wo` with the current secret plus `password_wo_version = 1`. Review and apply the plan; Terraform will track the version and remove the old secret value from the resource's current state. Other resource-specific changes can still cause replacement, as described by that resource's schema.

For a template parameter, remove its name from the ordinary `params` map and add it to the matching `secret_params` entry's `params_wo` and `params_wo_versions` maps. Review the planned deployment and the `redeploy` policy before applying.

Migration does not remove values from older state snapshots, backups, or saved plans. Those artifacts still contain any secrets that were previously stored in ordinary or sensitive arguments.
