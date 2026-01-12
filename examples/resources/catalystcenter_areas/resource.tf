resource "catalystcenter_areas" "example" {
  areas = [
    {
      parent_name_hierarchy = "Global"
      name                  = "Area1"
    }
  ]
}
