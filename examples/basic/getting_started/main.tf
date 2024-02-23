resource "catalystcenter_area" "san_jose" {
  name        = "San Jose"
  parent_name = "Global"
}

resource "catalystcenter_building" "main_building" {
  name        = "Main Building"
  parent_name = "Global/${catalystcenter_area.san_jose.name}"
  latitude    = 37.338
  longitude   = -121.832
}

resource "catalystcenter_floor" "ground_floor" {
  name        = "Ground Floor"
  parent_name = "${catalystcenter_building.main_building.parent_name}/${catalystcenter_building.main_building.name}"
  rf_model    = "Drywall Office Only"
  width       = 30.5
  length      = 50.5
  height      = 3.5
}

resource "catalystcenter_ip_pool" "mgmt_pool_1" {
  name             = "MgmtPool1"
  ip_address_space = "IPv4"
  type             = "Generic"
  ip_subnet        = "10.100.0.0/16"
  dhcp_server_ips  = ["10.101.1.1"]
  dns_server_ips   = ["10.101.1.101", "10.101.1.102"]
}

resource "catalystcenter_ip_pool_reservation" "san_jose_main_building_mgmt_1" {
  site_id            = catalystcenter_building.main_building.id
  name               = "SanJoseMainBuildingMgmt1"
  type               = "Generic"
  ipv6_address_space = false
  ipv4_global_pool   = catalystcenter_ip_pool.mgmt_pool_1.ip_subnet
  ipv4_prefix        = true
  ipv4_prefix_length = 24
  ipv4_subnet        = "10.100.0.0"
  ipv4_gateway       = "10.100.0.1"
}