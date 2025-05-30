output "is_valid" {
  description = "Returns true if all resource names are valid"
  value       = local.is_valid
}

output "validation_details" {
  description = "Detailed validation results for each resource"
  value = {
    resource_group_name    = local.validate_resource_group_name
    storage_account_name   = local.validate_storage_account_name
    key_vault_name         = local.validate_key_vault_name
    web_app_name           = local.validate_web_app_name
    virtual_network_name   = local.validate_virtual_network_name
    subnet_name            = local.validate_subnet_name
    nsg_name               = local.validate_nsg_name
    storage_container_name = local.validate_storage_container_name
    app_service_plan_name  = local.validate_app_service_plan_name
    storage_account_chars  = local.validate_storage_account_chars
  }
}
