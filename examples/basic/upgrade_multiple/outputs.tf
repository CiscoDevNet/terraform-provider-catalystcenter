output "all_discovered" {
  description = "All the network devices discovered on Catalyst Center (may be large)"
  value       = local.devices_raw
}
