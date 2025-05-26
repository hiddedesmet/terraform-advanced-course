variable "resource_group_name" {
  description = "Name of the resource group"
  type        = string
  default     = "myTFResourceGroup2"
}

variable "storage_account_name" {
  description = "Name of the storage account (must be globally unique)"
  type        = string
  default     = "mytfstoragedemohidde"
}

variable "web_app_name" {
  description = "Name of the web app (must be globally unique)"
  type        = string
  default     = "mytfwebappdemohidde"
}

variable "key_vault_name" {
  description = "Name of the Key Vault (must be globally unique)"
  type        = string
  default     = "mytfkeyvaultdemohidde"
}

variable "location" {
  description = "Azure region for resources"
  type        = string
  default     = "westeurope"
}

variable "virtual_network_name" {
  description = "Name of the virtual network"
  type        = string
  default     = "myTFVnet2"
}

variable "subnet_name" {
  description = "Name of the subnet"
  type        = string
  default     = "myTFSubnet"
}

variable "nsg_name" {
  description = "Name of the network security group"
  type        = string
  default     = "myTFNetworkSecurityGroup"
}

variable "storage_container_name" {
  description = "Name of the storage container"
  type        = string
  default     = "demo-container"
}

variable "app_service_plan_name" {
  description = "Name of the App Service Plan"
  type        = string
  default     = "myTFAppServicePlan"
}
