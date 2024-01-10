resource "catalystcenter_image" "example" {
  source_path = "../software.bin"
  name        = "basename(\"../software.bin\")"
}
