resource "azurerm_resource_group" "sshstats-resource-group" {
  location = "westeurope"
  name     = "default"
}

resource "azurerm_storage_account" "sshstats-storage" {
  account_kind                     = "Storage"
  account_replication_type         = "LRS"
  account_tier                     = "Standard"
  allow_nested_items_to_be_public  = false
  cross_tenant_replication_enabled = false
  default_to_oauth_authentication  = true
  location                         = "westeurope"
  name                             = "sshstatsstore"
  resource_group_name              = azurerm_resource_group.sshstats-resource-group.name
}

resource "azurerm_service_plan" "sshstats-fn-serviceplan" {
  location            = "westeurope"
  name                = "sshstats-fn-service-plan"
  os_type             = "Linux"
  resource_group_name = azurerm_resource_group.sshstats-resource-group.name
  sku_name            = "Y1"
}

resource "azurerm_static_web_app" "sshstats-frontend" {
  location            = "westeurope"
  name                = "sshstats-fe-swa"
  resource_group_name = azurerm_resource_group.sshstats-resource-group.name
}

resource "azurerm_linux_function_app" "sshstats-backend" {
  app_settings = {
    GIN_MODE                  = "release"
    WEBSITE_RUN_FROM_PACKAGE  = "1"
    SSHSTATS_AUTH_HEADERKEY   = var.sshstats_fn_config_headerkey
    SSHSTATS_AUTH_HEADERVALUE = var.sshstats_fn_config_headervalue
    SSHSTATS_LOKI_ENDPOINT    = var.sshstats_fn_config_loki_endpoint
    SSHSTATS_LOKI_PASSWORD    = var.sshstats_fn_config_loki_pwd
    SSHSTATS_LOKI_USERNAME    = var.sshstats_fn_config_loki_usr
  }

  builtin_logging_enabled                  = false
  client_certificate_mode                  = "Required"
  ftp_publish_basic_authentication_enabled = false
  https_only                               = true
  location                                 = "westeurope"
  name                                     = "sshstats-be-fn"
  resource_group_name                      = azurerm_resource_group.sshstats-resource-group.name
  service_plan_id                          = azurerm_service_plan.sshstats-fn-serviceplan.id
  storage_account_access_key               = azurerm_storage_account.sshstats-storage.primary_access_key
  storage_account_name                     = azurerm_storage_account.sshstats-storage.name

  site_config {
    ftps_state = "FtpsOnly"
    application_stack {
      use_custom_runtime = true
    }
  }
}
