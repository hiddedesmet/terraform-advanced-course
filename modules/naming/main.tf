locals {
  # Resource abbreviations
  resource_type_abbreviations = {
    resource_group         = "rg"
    virtual_network        = "vnet"
    subnet                 = "snet"
    network_security_group = "nsg"
    storage_account        = "st"
    storage_container      = "stcont"
    app_service_plan       = "asp"
    web_app                = "app"
    key_vault              = "kv"
  }
  
  # Common tags that should be applied to all resources
  common_tags = merge(
    var.tags,
    {
      environment = var.environment
      project     = var.project_name
      managed_by  = "terraform"
    }
  )
}

# Resource naming function
resource "null_resource" "resource_name" {
  triggers = {
    for resource_type, abbreviation in local.resource_type_abbreviations :
    resource_type => "${var.prefix}-${abbreviation}-${var.environment}-${var.suffix}"
  }
}

# Special naming function for storage accounts (no dashes allowed)
resource "null_resource" "storage_account_name" {
  triggers = {
    name = "${var.prefix}st${var.environment}${var.suffix}"
  }
}
