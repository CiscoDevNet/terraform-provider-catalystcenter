resource "catalystcenter_map_archive_import" "example" {
  source_path      = "testdata/map_archives/nac_map_import.tar.gz"
  accept_data_loss = true
}
