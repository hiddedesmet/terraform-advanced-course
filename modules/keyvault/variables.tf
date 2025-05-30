variable "subscription_id" {
  description = "Azure subscription ID"
  type        = string
  validation {
    condition     = can(regex("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$", var.subscription_id))
    error_message = "Subscription ID must be a valid UUID format."
  }
}

variable "key_vault_name" {
  description = "Name of the Key Vault (must be globally unique)"
  type        = string
}

variable "location" {
  description = "Azure region for resources"
  type        = string
}

variable "resource_group_name" {
  description = "Name of the resource group"
  type        = string
}

variable "sku_name" {
  description = "SKU name for the Key Vault"
  type        = string
  default     = "standard"
}

variable "purge_protection_enabled" {
  description = "Whether purge protection is enabled"
  type        = bool
  default     = false
}

variable "soft_delete_retention_days" {
  description = "Soft delete retention days"
  type        = number
  default     = 7
}

variable "enable_rbac_authorization" {
  description = "Whether RBAC authorization is enabled"
  type        = bool
  default     = true
}

variable "tags" {
  description = "Tags for resources"
  type        = map(string)
  default     = {}
}
