output "resource_group_name" {
  description = "Name of the resource group"
  value       = data.azurerm_resource_group.rg.name
}

output "storage_account_name" {
  description = "Name of the storage account"
  value       = module.storage.storage_account_name
}

output "key_vault_name" {
  description = "Name of the Key Vault"
  value       = module.keyvault.key_vault_name
}

output "web_app_name" {
  description = "Name of the Web App"
  value       = module.webapp.web_app_name
}

output "nsg_name" {
  description = "Name of the Network Security Group"
  value       = module.network.nsg_name
}
