resource "catalystcenter_buildings" "example" {
  scope = "Global/Poland"
  buildings = {
    "Global/USA/Building1" = {
      parent_name_hierarchy = "Global/USA"
      name                  = "Building1"
      country               = "United States"
      address               = "150 W Tasman Dr, San Jose"
      latitude              = 37.338
      longitude             = -121.832
    }
  }
}
