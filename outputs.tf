output "resource_group_id" {
  description = "ID of the resource group"
  value       = azurerm_resource_group.rg.id
}

output "virtual_network_id" {
  description = "ID of the virtual network"
  value       = azurerm_virtual_network.vnet.id
}

output "subnet_id" {
  description = "ID of the subnet"
  value       = azurerm_subnet.subnet.id
}

output "storage_account_name" {
  description = "Name of the storage account"
  value       = azurerm_storage_account.storage.name
}

output "storage_account_primary_blob_endpoint" {
  description = "Primary blob endpoint of the storage account"
  value       = azurerm_storage_account.storage.primary_blob_endpoint
}

output "web_app_url" {
  description = "URL of the web app"
  value       = "https://${azurerm_linux_web_app.web_app.default_hostname}"
}

output "web_app_name" {
  description = "Name of the web app"
  value       = azurerm_linux_web_app.web_app.name
}

output "key_vault_uri" {
  description = "URI of the Key Vault"
  value       = azurerm_key_vault.key_vault.vault_uri
}

output "network_security_group_id" {
  description = "ID of the network security group"
  value       = azurerm_network_security_group.nsg.id
}
