locals {
  images_uploaded = toset([
    "cat9k_iosxe.17.12.02.SPA.bin",
    "cat9k_iosxe.17.12.02prd4.SPA.bin",
    "C9800-SW-iosxe-wlc.17.12.02prd4.SPA.bin",
  ])
}

resource "catalystcenter_image" "this" {
  for_each = local.images_uploaded

  name        = basename(each.key)
  source_path = each.key
}
