resource "catalystcenter_credentials_https_read" "example" {
  description = "My HTTPS read credentials"
  username    = "user1"
  password    = "password1"
  port        = 444
}
