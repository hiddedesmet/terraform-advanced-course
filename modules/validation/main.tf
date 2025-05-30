locals {
  # Maximum length validation for Azure resources
  max_length = {
    resource_group_name    = 90
    storage_account_name   = 24
    key_vault_name         = 24
    web_app_name           = 60
    virtual_network_name   = 64
    subnet_name            = 80
    nsg_name               = 80
    storage_container_name = 63
    app_service_plan_name  = 40
  }

  # Resource name validation - return validation results instead of causing errors
  validate_resource_group_name    = length(var.resource_group_name) <= local.max_length.resource_group_name
  validate_storage_account_name   = length(var.storage_account_name) <= local.max_length.storage_account_name
  validate_key_vault_name         = length(var.key_vault_name) <= local.max_length.key_vault_name
  validate_web_app_name           = length(var.web_app_name) <= local.max_length.web_app_name
  validate_virtual_network_name   = length(var.virtual_network_name) <= local.max_length.virtual_network_name
  validate_subnet_name            = length(var.subnet_name) <= local.max_length.subnet_name
  validate_nsg_name               = length(var.nsg_name) <= local.max_length.nsg_name
  validate_storage_container_name = length(var.storage_container_name) <= local.max_length.storage_container_name
  validate_app_service_plan_name  = length(var.app_service_plan_name) <= local.max_length.app_service_plan_name

  # Additional storage account name validation (lowercase letters and numbers only)
  validate_storage_account_chars = can(regex("^[a-z0-9]+$", var.storage_account_name))

  # Overall validation result
  is_valid = local.validate_resource_group_name && local.validate_storage_account_name && local.validate_key_vault_name && local.validate_web_app_name && local.validate_virtual_network_name && local.validate_subnet_name && local.validate_nsg_name && local.validate_storage_container_name && local.validate_app_service_plan_name && local.validate_storage_account_chars
}
