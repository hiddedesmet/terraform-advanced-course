# Configure the Azure provider
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0.2"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.0"
    }
  }

  required_version = ">= 1.1.0"
}

provider "azurerm" {
  features {}
}

# Use environment variable to determine environment or default to "dev"
locals {
  environment = terraform.workspace == "default" ? var.environment : terraform.workspace
}

module "validation" {
  source = "./modules/validation"

  resource_group_name    = var.resource_group_name
  storage_account_name   = var.storage_account_name
  key_vault_name         = var.key_vault_name
  web_app_name           = var.web_app_name
  virtual_network_name   = var.virtual_network_name
  subnet_name            = var.subnet_name
  nsg_name               = var.nsg_name
  storage_container_name = var.storage_container_name
  app_service_plan_name  = var.app_service_plan_name
}

module "naming" {
  source = "./modules/naming"

  prefix       = var.prefix
  environment  = local.environment
  suffix       = var.suffix
  project_name = var.project_name
  tags         = {
    Team = "DevOps"
  }
}

module "tagging" {
  source = "./modules/tagging"

  environment       = local.environment
  project_name      = var.project_name
  owner             = var.owner
  cost_center       = var.cost_center
  terraform_version = "v${terraform.version}"
}

resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name  # You can switch to module.naming.resource_group when ready
  location = var.location

  tags = module.tagging.tags
}

module "network" {
  source = "./modules/network"

  resource_group_name     = azurerm_resource_group.rg.name
  location                = azurerm_resource_group.rg.location
  virtual_network_name    = var.virtual_network_name
  address_space           = ["10.0.0.0/16"]
  subnet_name             = var.subnet_name
  subnet_address_prefixes = ["10.0.1.0/24"]
  nsg_name                = var.nsg_name
  tags                    = module.tagging.tags
}

module "storage" {
  source = "./modules/storage"

  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  storage_account_name     = var.storage_account_name
  storage_container_name   = var.storage_container_name
  account_tier             = "Standard"
  account_replication_type = "LRS"
  container_access_type    = "private"
  tags                     = module.tagging.tags
}

module "webapp" {
  source = "./modules/webapp"

  resource_group_name     = azurerm_resource_group.rg.name
  location                = azurerm_resource_group.rg.location
  app_service_plan_name   = var.app_service_plan_name
  web_app_name            = var.web_app_name
  os_type                 = "Linux"
  sku_name                = "B1"
  https_only              = true
  minimum_tls_version     = "1.2"
  php_version             = "8.0"
  app_settings            = {
    "WEBSITE_RUN_FROM_PACKAGE" = "1"
  }
  tags                    = module.tagging.tags
}

module "keyvault" {
  source = "./modules/keyvault"

  resource_group_name         = azurerm_resource_group.rg.name
  location                    = azurerm_resource_group.rg.location
  key_vault_name              = var.key_vault_name
  sku_name                    = "standard"
  purge_protection_enabled    = false
  soft_delete_retention_days  = 7
  enable_rbac_authorization   = true
  tags                        = module.tagging.tags
}
