variable "storage_account_name" {
  description = "Name of the storage account (must be globally unique)"
  type        = string
  validation {
    condition     = can(regex("^[a-z0-9]{3,24}$", var.storage_account_name))
    error_message = "Storage account name must be between 3-24 characters and contain only lowercase letters and numbers."
  }
}

variable "resource_group_name" {
  description = "Name of the resource group"
  type        = string
  validation {
    condition     = length(var.resource_group_name) >= 1 && length(var.resource_group_name) <= 90
    error_message = "Resource group name must be between 1-90 characters."
  }
}

variable "location" {
  description = "Azure region for resources"
  type        = string
  validation {
    condition = contains([
      # North America regions
      "northcentralus", "eastus", "eastus2", "westus", "westus2", "westus3", 
      "southcentralus", "centralus", "westcentralus", "canadacentral", "canadaeast",
      
      # Europe regions
      "northeurope", "westeurope", "uksouth", "ukwest", "francecentral", "francesouth",
      "switzerlandnorth", "switzerlandwest", "germanywestcentral", "norwayeast", "norwaywest",
      "swedencentral", "swedensouth",
      
      # Asia Pacific regions
      "eastasia", "southeastasia",
      "japaneast", "japanwest",
      "australiaeast", "australiasoutheast", "australiacentral", "australiacentral2",
      "centralindia", "southindia", "westindia",
      "koreacentral", "koreasouth",

      # South American regions
      "brazilsouth", "brazilsoutheast",

      # Middle East and Africa regions
      "southafricanorth", "southafricawest",
      "uaenorth", "uaecentral"
    ], var.location)
    error_message = "Provided Azure location is not valid or not supported."
  }
}

variable "account_tier" {
  description = "Tier of the storage account"
  type        = string
  default     = "Standard"

  validation {
    condition     = contains(["Standard", "Premium"], var.account_tier)
    error_message = "Account tier must be either 'Standard' or 'Premium'."
  }
}

variable "account_replication_type" {
  description = "Replication type for the storage account"
  type        = string
  default     = "LRS"
  validation {
    condition     = contains(["LRS", "ZRS", "GRS", "GZRS", "RA-GRS", "RA-GZRS"], var.account_replication_type)
    error_message = "Replication type must be one of: LRS, ZRS, GRS, GZRS, RA-GRS, RA-GZRS."
  }
}

variable "storage_container_name" {
  description = "Name of the storage container"
  type        = string
  validation {
    condition     = can(regex("^[a-z0-9](?!.*--)[a-z0-9-]{1,61}[a-z0-9]$|^\\$root$|^\\$web$", var.storage_container_name))
    error_message = "Storage container name must be between 3-63 characters, start and end with a letter or number, and contain only lowercase letters, numbers, and hyphens."
  }
}

variable "container_access_type" {
  description = "Access type for the storage container"
  type        = string
  default     = "private"
  validation {
    condition     = contains(["private", "blob", "container"], var.container_access_type)
    error_message = "Container access type must be one of: private, blob, container."
  }
}

variable "tags" {
  description = "Tags for resources"
  type        = map(string)
  default     = {}
}
