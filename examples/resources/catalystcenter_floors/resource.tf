resource "catalystcenter_floors" "example" {
  floors = [
    {
      parent_name_hierarchy = "Global/USA/San Jose/Building1"
      name                  = "Floor1"
      floor_number          = 1
      rf_model              = "Drywall Office Only"
      width                 = 40
      length                = 50.5
      height                = 3.5
      units_of_measure      = "meters"
    }
  ]
}
