resource "catalystcenter_credentials_https_write" "example" {
  description         = "My HTTPS write credentials"
  username            = "user1"
  password_wo         = "password1"
  password_wo_version = 1
  port                = 444
}
