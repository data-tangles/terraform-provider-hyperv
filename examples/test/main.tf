terraform {
  required_providers {
    hyperv = {
      version = "1.0.0"
      source  = "registry.terraform.io/qman-being/hyperv"
    }
  }
}

provider "hyperv" {

}

data "hyperv_network_switch" "dmz_network_switch" {
  name = "dmz"
}
