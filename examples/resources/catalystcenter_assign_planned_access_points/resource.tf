resource "catalystcenter_assign_planned_access_points" "example" {
  floor_id = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  assignments = [
    {
      planned_access_point_name = "NAC-PLANNED-AP-01"
      access_point_id           = "b054deab-f9fc-48c9-88d0-ab82f029e6aa"
    }
  ]
}
