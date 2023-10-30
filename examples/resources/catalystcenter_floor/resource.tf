resource "catalystcenter_floor" "example" {
  name        = "Floor1"
  parent_name = "Global/Building1"
  rf_model    = "Drywall Office Only"
  width       = 30.5
  length      = 50.5
  height      = 3.5
}
