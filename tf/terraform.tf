terraform {
  cloud {
    organization = "my-own-stuff"

    workspaces {
      name = "azure"
    }
  }

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.99.0"
    }
  }
}
