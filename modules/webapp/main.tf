terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>4.0"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
}

resource "azurerm_service_plan" "app_service_plan" {
  name                = var.app_service_plan_name
  resource_group_name = var.resource_group_name
  location            = var.location
  os_type             = var.os_type
  sku_name            = var.sku_name

  tags = var.tags
}

resource "azurerm_linux_web_app" "web_app" {
  name                = var.web_app_name
  resource_group_name = var.resource_group_name
  location            = azurerm_service_plan.app_service_plan.location
  service_plan_id     = azurerm_service_plan.app_service_plan.id
  https_only          = var.https_only

  site_config {
    minimum_tls_version = var.minimum_tls_version

    application_stack {
      php_version = var.php_version
    }
  }

  app_settings = var.app_settings

  tags = var.tags

  depends_on = [
    azurerm_service_plan.app_service_plan
  ]
}
