resource "catalystcenter_application_policy" "example" {
  policy_scope = "NOWEGO"
  items = [
    {
      name                       = "NOWEGO_collaboration-apps"
      policy_scope               = "NOWEGO"
      priority                   = "100"
      delete_policy_status       = "NONE"
      advanced_policy_scope_name = "NOWEGO"
      site_ids                   = ["12345678-1234-1234-1234-123456789012"]
      ssids                      = ["corp-ssid"]
      clause_type                = "BUSINESS_RELEVANCE"
      relevance_level            = "BUSINESS_RELEVANT"
      application_set_id         = "12345678-1234-1234-1234-123456789012"
    }
  ]
}
