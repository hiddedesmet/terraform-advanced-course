output "name" {
  description = "Name of the resource group"
  value       = azurerm_resource_group.test.name
}

output "location" {
  description = "Location of the resource group"
  value       = azurerm_resource_group.test.location
}

output "id" {
  description = "ID of the resource group"
  value       = azurerm_resource_group.test.id
}
