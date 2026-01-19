resource "catalystcenter_areas" "example" {
  areas = {
    "Global/Area1" = {
      parent_name_hierarchy = "Global"
      name                  = "Area1"
    }
  }
}
