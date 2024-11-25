provider "azurerm" {
  use_msi  = false
  use_oidc = false
  use_cli  = var.sshstats_az_use_cli

  subscription_id = var.sshstats_subscription
  tenant_id       = var.sshstats_az_tenant_id

  client_id     = var.sshstats_az_client_id
  client_secret = var.sshstats_az_client_secret

  features {}
}
