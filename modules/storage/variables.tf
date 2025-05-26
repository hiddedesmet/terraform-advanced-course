variable "storage_account_name" {
  description = "Name of the storage account (must be globally unique)"
  type        = string
}

variable "resource_group_name" {
  description = "Name of the resource group"
  type        = string
}

variable "location" {
  description = "Azure region for resources"
  type        = string
}

variable "account_tier" {
  description = "Tier of the storage account"
  type        = string
  default     = "Standard"
}

variable "account_replication_type" {
  description = "Replication type for the storage account"
  type        = string
  default     = "LRS"
}

variable "storage_container_name" {
  description = "Name of the storage container"
  type        = string
}

variable "container_access_type" {
  description = "Access type for the storage container"
  type        = string
  default     = "private"
}

variable "tags" {
  description = "Tags for resources"
  type        = map(string)
  default     = {}
}
