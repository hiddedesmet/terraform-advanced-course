terraform {
  required_version = ">= 1.0"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
}

# Use data source for existing resource group
data "azurerm_resource_group" "rg" {
  name = var.resource_group_name
}

# Use the same modules as the main configuration
module "naming" {
  source = "../../../modules/naming"

  prefix         = var.prefix
  environment    = var.environment
  location       = var.location
  suffix         = var.suffix
  resource_group = var.resource_group_name
}

module "tagging" {
  source = "../../../modules/tagging"

  environment  = var.environment
  project_name = var.project_name
  owner        = var.owner
  cost_center  = var.cost_center
}

module "network" {
  source = "../../../modules/network"

  resource_group_name  = data.azurerm_resource_group.rg.name
  location             = data.azurerm_resource_group.rg.location
  virtual_network_name = var.virtual_network_name
  subnet_name          = var.subnet_name
  nsg_name             = var.nsg_name

  tags = module.tagging.tags
}

module "storage" {
  source = "../../../modules/storage"

  resource_group_name  = data.azurerm_resource_group.rg.name
  location             = data.azurerm_resource_group.rg.location
  storage_account_name = var.storage_account_name
  container_name       = var.storage_container_name

  tags = module.tagging.tags
}

module "keyvault" {
  source = "../../../modules/keyvault"

  resource_group_name = data.azurerm_resource_group.rg.name
  location            = data.azurerm_resource_group.rg.location
  key_vault_name      = var.key_vault_name

  tags = module.tagging.tags
}

module "webapp" {
  source = "../../../modules/webapp"

  resource_group_name   = data.azurerm_resource_group.rg.name
  location              = data.azurerm_resource_group.rg.location
  app_service_plan_name = var.app_service_plan_name
  web_app_name          = var.web_app_name

  tags = module.tagging.tags
}
