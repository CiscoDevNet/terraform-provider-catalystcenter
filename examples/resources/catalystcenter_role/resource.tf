resource "catalystcenter_role" "example" {
  name        = "Role1"
  description = "Role1 description"
  resource_types = [
    {
      type       = "Platform"
      operations = ["gRead"]
    }
  ]
}
