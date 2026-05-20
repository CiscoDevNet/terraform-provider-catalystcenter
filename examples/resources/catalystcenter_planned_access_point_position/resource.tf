resource "catalystcenter_planned_access_point_position" "example" {
  floor_id   = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  name       = "AP-1001"
  ap_type    = "AP9166I"
  position_x = 50.5
  position_y = 30.2
  position_z = 3.5
  radios = [
    {
      bands             = ["5"]
      channel           = 36
      tx_power          = 17
      antenna_name      = "Internal-CW9166I-x-5GHz"
      antenna_azimuth   = 0
      antenna_elevation = 0
    }
  ]
}
