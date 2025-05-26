output "virtual_network_id" {
  description = "ID of the virtual network"
  value       = azurerm_virtual_network.vnet.id
}

output "subnet_id" {
  description = "ID of the subnet"
  value       = azurerm_subnet.subnet.id
}

output "network_security_group_id" {
  description = "ID of the network security group"
  value       = azurerm_network_security_group.nsg.id
}
