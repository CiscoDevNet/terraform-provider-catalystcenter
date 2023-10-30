resource "catalystcenter_credentials_https_write" "example" {
  description = "My HTTPS write credentials"
  username    = "user1"
  password    = "password1"
  port        = 444
}
