resource "catalystcenter_ap_profile" "example" {
  ap_profile_name        = "AP_Profile_1"
  description            = "My AP Profile Description"
  remote_worker_enabled  = false
  awips_enabled          = false
  awips_forensic_enabled = false
  pmf_denial_enabled     = false
  mesh_enabled           = false
}
