resource "catalystcenter_areas" "example" {
  scope = "Global/Poland"
  areas = {
    "Global/Area1" = {
      parent_name_hierarchy = "Global"
      name                  = "Area1"
    }
  }
}
