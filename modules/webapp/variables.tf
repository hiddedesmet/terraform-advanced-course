variable "app_service_plan_name" {
  description = "Name of the App Service Plan"
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

variable "os_type" {
  description = "OS type for the App Service Plan"
  type        = string
  default     = "Linux"
}

variable "sku_name" {
  description = "SKU name for the App Service Plan"
  type        = string
  default     = "B1"
}

variable "web_app_name" {
  description = "Name of the web app"
  type        = string
}

variable "https_only" {
  description = "Whether HTTPS only is enabled"
  type        = bool
  default     = true
}

variable "minimum_tls_version" {
  description = "Minimum TLS version"
  type        = string
  default     = "1.2"
}

variable "php_version" {
  description = "PHP version for the web app"
  type        = string
  default     = "8.0"
}

variable "app_settings" {
  description = "App settings for the web app"
  type        = map(string)
  default     = {
    "WEBSITE_RUN_FROM_PACKAGE" = "1"
  }
}

variable "tags" {
  description = "Tags for resources"
  type        = map(string)
  default     = {}
}
