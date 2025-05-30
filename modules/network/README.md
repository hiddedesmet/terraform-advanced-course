# Network Module

## Overview

This module provisions Azure networking resources including virtual networks, subnets, and network security groups. It provides a standardized approach to network configuration across different environments.

## Features

- Virtual Network creation with configurable address space
- Subnet provisioning with customizable address prefixes
- Network Security Group with predefined rules for common protocols
- Association between subnets and NSGs
- Consistent tagging for all network resources

## Usage

```hcl
module "network" {
  source = "./modules/network"
  
  # Required variables
  resource_group_name     = azurerm_resource_group.rg.name
  location                = azurerm_resource_group.rg.location
  virtual_network_name    = "my-vnet"
  subnet_name             = "my-subnet"
  
  # Optional variables
  address_space           = ["10.0.0.0/16"]
  subnet_address_prefixes = ["10.0.1.0/24"]
  nsg_name                = "my-nsg"
  tags                    = {
    Environment = "dev"
    Project     = "demo"
  }
}
```

## Advanced Example

```hcl
module "network" {
  source = "./modules/network"
  
  resource_group_name     = azurerm_resource_group.rg.name
  location                = azurerm_resource_group.rg.location
  virtual_network_name    = module.naming.virtual_network
  subnet_name             = module.naming.subnet
  nsg_name                = module.naming.network_security_group
  address_space           = ["10.0.0.0/16"]
  subnet_address_prefixes = ["10.0.1.0/24"]
  
  # Adding custom NSG rules
  nsg_rules = [
    {
      name                       = "SSH"
      priority                   = 100
      direction                  = "Inbound"
      access                     = "Allow"
      protocol                   = "Tcp"
      source_port_range          = "*"
      destination_port_range     = "22"
      source_address_prefix      = "10.0.0.0/24"
      destination_address_prefix = "*"
    },
    {
      name                       = "HTTPS"
      priority                   = 110
      direction                  = "Inbound"
      access                     = "Allow"
      protocol                   = "Tcp"
      source_port_range          = "*"
      destination_port_range     = "443"
      source_address_prefix      = "*"
      destination_address_prefix = "*"
    }
  ]
  
  tags = module.tagging.tags
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.1.0 |
| azurerm | ~> 4.0 |

## Resources

| Name | Type |
|------|------|
| [azurerm_virtual_network](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network) | resource |
| [azurerm_subnet](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet) | resource |
| [azurerm_network_security_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_group) | resource |
| [azurerm_subnet_network_security_group_association](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet_network_security_group_association) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| resource_group_name | Name of the resource group | `string` | n/a | yes |
| location | Azure region for resources | `string` | n/a | yes |
| virtual_network_name | Name of the virtual network | `string` | n/a | yes |
| subnet_name | Name of the subnet | `string` | n/a | yes |
| address_space | Address space for the virtual network | `list(string)` | `["10.0.0.0/16"]` | no |
| subnet_address_prefixes | Address prefixes for the subnet | `list(string)` | `["10.0.1.0/24"]` | no |
| nsg_name | Name of the network security group | `string` | n/a | yes |
| nsg_rules | List of NSG rules to create | `list(object)` | `[]` | no |
| tags | Tags to apply to resources | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| virtual_network_id | ID of the virtual network |
| virtual_network_name | Name of the virtual network |
| subnet_id | ID of the subnet |
| subnet_name | Name of the subnet |
| network_security_group_id | ID of the network security group |
| network_security_group_name | Name of the network security group |

## Notes

- Consider using the naming module to generate consistent resource names
- For production environments, consider implementing network peering and more restrictive NSG rules
- When using custom NSG rules, ensure they don't conflict with existing rules
