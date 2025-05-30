variable "prefix" {
  description = "Prefix to be used in resource names"
  type        = string
  default     = "tf"
}

variable "subscription_id" {
  description = "Azure subscription ID"
  type        = string
  validation {
    condition     = can(regex("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$", var.subscription_id))
    error_message = "Subscription ID must be a valid UUID format."
  }
}

variable "environment" {
  description = "Environment name (dev, test, prod, etc.)"
  type        = string
  default     = "dev"
}

variable "suffix" {
  description = "Suffix to be used in resource names"
  type        = string
  default     = "01"
}

variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "demo"
}

variable "owner" {
  description = "Owner of the resources"
  type        = string
  default     = "DevOps Team"
}

variable "cost_center" {
  description = "Cost center for billing"
  type        = string
  default     = "IT-12345"
}

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
