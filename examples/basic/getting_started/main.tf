data "catalystcenter_site" "global" {
  name_hierarchy = "Global"
}

resource "catalystcenter_area" "san_jose" {
  name      = "San Jose"
  parent_id = data.catalystcenter_site.global.id
}

resource "catalystcenter_building" "main_building" {
  name      = "Main Building"
  parent_id = catalystcenter_area.san_jose.id
  country   = "United States"
  latitude  = 37.338
  longitude = -121.832
}

resource "catalystcenter_floor" "ground_floor" {
  name             = "Ground Floor"
  parent_id        = catalystcenter_building.main_building.id
  rf_model         = "Drywall Office Only"
  floor_number     = 1
  width            = 30.5
  length           = 50.5
  height           = 3.5
  units_of_measure = "feet"
}

resource "catalystcenter_ip_pool" "mgmt_pool_1" {
  name                        = "MgmtPool1"
  pool_type                   = "Generic"
  address_space_subnet        = "10.100.0.0"
  address_space_prefix_length = 16
  address_space_dhcp_servers  = ["10.101.1.1"]
  address_space_dns_servers   = ["10.101.1.101", "10.101.1.102"]
}

resource "catalystcenter_ip_pool_reservation" "san_jose_main_building_mgmt_1" {
  site_id             = catalystcenter_building.main_building.id
  name                = "SanJoseMainBuildingMgmt1"
  pool_type           = "Generic"
  ipv4_global_pool_id = catalystcenter_ip_pool.mgmt_pool_1.id
  ipv4_prefix_length  = 24
  ipv4_subnet         = "10.100.0.0"
  ipv4_gateway        = "10.100.0.1"
}