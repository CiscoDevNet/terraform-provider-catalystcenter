resource "catalystcenter_floor" "example" {
  parent_id        = "34af13e4-da94-4c5f-8da2-aa51a0f0655e"
  name             = "Floor1"
  floor_number     = 1
  rf_model         = "Drywall Office Only"
  width            = 40
  length           = 50.5
  height           = 3.5
  units_of_measure = "meters"
}
