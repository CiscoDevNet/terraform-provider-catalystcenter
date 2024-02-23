terraform {
  required_providers {
    catalystcenter = {
      source = "CiscoDevNet/catalystcenter"
    }
  }
}

provider "catalystcenter" {
  // By default uses env $CC_USERNAME $CC_PASSWORD $CC_URL
}
