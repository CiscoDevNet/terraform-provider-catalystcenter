resource "catalystcenter_map_archive_import" "example" {
  source_path      = "testdata/map_archives/nac_map_import.tar.gz"
  accept_data_loss = true
  archive_sha256   = "11052e417e1a8b0f763bc363c4cb994be3c0330f33afba927a6a402c6e2a0364"
}
