resource "catalystcenter_credentials_https_read" "example" {
  description         = "My HTTPS read credentials"
  username            = "user1"
  password_wo         = "password1"
  password_wo_version = 1
  port                = 444
}
