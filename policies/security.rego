package terraform.security

import input.planned_values as tfplan

# Ensure storage accounts use the secure transfer required setting
deny[msg] {
    resource := tfplan.root_module.resources[_]
    resource.type == "azurerm_storage_account"
    not resource.values.enable_https_traffic_only
    
    msg := sprintf(
        "Storage account %s must have enable_https_traffic_only set to true",
        [resource.address]
    )
}

# Ensure storage accounts have network rules configured
deny[msg] {
    resource := tfplan.root_module.resources[_]
    resource.type == "azurerm_storage_account"
    not resource.values.network_rules
    
    msg := sprintf(
        "Storage account %s must have network rules configured",
        [resource.address]
    )
}

# Ensure key vaults have purge protection enabled
deny[msg] {
    resource := tfplan.root_module.resources[_]
    resource.type == "azurerm_key_vault"
    not resource.values.purge_protection_enabled
    
    msg := sprintf(
        "Key vault %s must have purge_protection_enabled set to true",
        [resource.address]
    )
}

# Ensure web apps enforce HTTPS
deny[msg] {
    resource := tfplan.root_module.resources[_]
    resource.type == "azurerm_app_service"
    not resource.values.https_only
    
    msg := sprintf(
        "Web app %s must have https_only set to true",
        [resource.address]
    )
}

# Ensure web apps use a minimum TLS version of 1.2
deny[msg] {
    resource := tfplan.root_module.resources[_]
    resource.type == "azurerm_app_service"
    
    # Check if minimum_tls_version is less than 1.2
    version := resource.values.site_config.minimum_tls_version
    version == "1.0" or version == "1.1"
    
    msg := sprintf(
        "Web app %s must use minimum TLS version 1.2",
        [resource.address]
    )
}

# Ensure virtual networks don't use overlapping address spaces
deny[msg] {
    resource1 := tfplan.root_module.resources[i]
    resource2 := tfplan.root_module.resources[j]
    
    resource1.type == "azurerm_virtual_network"
    resource2.type == "azurerm_virtual_network"
    
    # Don't compare a resource with itself
    i < j
    
    # Check for address space overlap
    cidr1 := resource1.values.address_space[_]
    cidr2 := resource2.values.address_space[_]
    cidr_overlap(cidr1, cidr2)
    
    msg := sprintf(
        "Virtual networks %s and %s have overlapping address spaces: %s and %s",
        [resource1.address, resource2.address, cidr1, cidr2]
    )
}

# Helper function to check if two CIDR blocks overlap
# Note: In a real implementation, this would perform actual CIDR overlap checking
cidr_overlap(cidr1, cidr2) {
    cidr1 == cidr2
}
