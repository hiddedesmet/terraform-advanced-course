output "resource_group_id" {
  description = "ID of the resource group"
  value       = azurerm_resource_group.rg.id
}

output "virtual_network_id" {
  description = "ID of the virtual network"
  value       = module.network.virtual_network_id
}

output "subnet_id" {
  description = "ID of the subnet"
  value       = module.network.subnet_id
}

output "storage_account_name" {
  description = "Name of the storage account"
  value       = module.storage.storage_account_name
}

output "storage_account_primary_blob_endpoint" {
  description = "Primary blob endpoint of the storage account"
  value       = module.storage.storage_account_primary_blob_endpoint
}

output "web_app_url" {
  description = "URL of the web app"
  value       = module.webapp.web_app_url
}

output "web_app_name" {
  description = "Name of the web app"
  value       = module.webapp.web_app_name
}

output "key_vault_uri" {
  description = "URI of the Key Vault"
  value       = module.keyvault.key_vault_uri
}

output "network_security_group_id" {
  description = "ID of the network security group"
  value       = module.network.network_security_group_id
}
