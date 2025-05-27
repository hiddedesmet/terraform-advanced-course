output "resource_group" {
  description = "Name for resource group"
  value       = null_resource.resource_name.triggers.resource_group
}

output "virtual_network" {
  description = "Name for virtual network"
  value       = null_resource.resource_name.triggers.virtual_network
}

output "subnet" {
  description = "Name for subnet"
  value       = null_resource.resource_name.triggers.subnet
}

output "network_security_group" {
  description = "Name for network security group"
  value       = null_resource.resource_name.triggers.network_security_group
}

output "storage_account" {
  description = "Name for storage account"
  value       = null_resource.storage_account_name.triggers.name
}

output "storage_container" {
  description = "Name for storage container"
  value       = null_resource.storage_container_name.triggers.name
}

output "app_service_plan" {
  description = "Name for app service plan"
  value       = null_resource.resource_name.triggers.app_service_plan
}

output "web_app" {
  description = "Name for web app"
  value       = null_resource.resource_name.triggers.web_app
}

output "key_vault" {
  description = "Name for key vault"
  value       = null_resource.resource_name.triggers.key_vault
}

output "common_tags" {
  description = "Common tags to apply to all resources"
  value       = local.common_tags
}
