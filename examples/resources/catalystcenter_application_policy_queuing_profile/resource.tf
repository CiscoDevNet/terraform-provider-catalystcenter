resource "catalystcenter_application_policy_queuing_profile" "example" {
  name        = "branch-queuing"
  description = "Branch WAN queuing profile"
  clauses = [
    {
      type                                   = "BANDWIDTH"
      is_common_between_all_interface_speeds = true
      interface_speed_bandwidth_clauses = [
        {
          interface_speed = "ALL"
          tc_bandwidth_settings = [
            {
              traffic_class        = "VOIP_TELEPHONY"
              bandwidth_percentage = 10
            }
          ]
        }
      ]
      tc_dscp_settings = [
        {
          traffic_class = "VOIP_TELEPHONY"
          dscp          = "46"
        }
      ]
    }
  ]
}
