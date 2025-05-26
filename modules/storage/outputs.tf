output "storage_account_name" {
  description = "Name of the storage account"
  value       = azurerm_storage_account.storage.name
}

output "storage_account_primary_blob_endpoint" {
  description = "Primary blob endpoint of the storage account"
  value       = azurerm_storage_account.storage.primary_blob_endpoint
}

output "storage_account_id" {
  description = "ID of the storage account"
  value       = azurerm_storage_account.storage.id
}

output "storage_container_name" {
  description = "Name of the storage container"
  value       = azurerm_storage_container.container.name
}
