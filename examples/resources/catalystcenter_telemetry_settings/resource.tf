resource "catalystcenter_telemetry_settings" "example" {
  site_id                             = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  enable_wired_data_collection        = true
  enable_wireless_telemetry           = true
  use_builtin_trap_server             = true
  external_trap_servers               = ["10.0.0.1"]
  use_builtin_syslog_server           = true
  external_syslog_servers             = ["10.0.0.1"]
  netflow_collector                   = "Builtin"
  enable_netflow_collector_on_devices = false
}
