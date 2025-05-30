package terraform.tagging

import input.planned_values as tfplan

# List of resource types that require tags
resource_types = [
    "azurerm_resource_group",
    "azurerm_virtual_network",
    "azurerm_subnet",
    "azurerm_network_security_group",
    "azurerm_storage_account",
    "azurerm_app_service",
    "azurerm_app_service_plan",
    "azurerm_key_vault"
]

# Required tags for all resources
required_tags = [
    "Environment",
    "Project",
    "Owner",
    "ManagedBy",
    "CostCenter"
]

# Check if resource is of a type that requires tags
is_taggable(resource) {
    resource.type == resource_types[_]
}

# Get missing tags for a resource
missing_tags(resource) = tags {
    # Get the tags that should be defined
    required := {t | t := required_tags[_]}
    
    # Get the tags that are defined
    tags_resource := {t | resource.values.tags[t]}
    
    # Calculate missing tags
    tags := required - tags_resource
}

# Deny resources with missing tags
deny[msg] {
    resource := tfplan.root_module.resources[_]
    is_taggable(resource)
    missing := missing_tags(resource)
    count(missing) > 0
    
    msg := sprintf(
        "%s (%s) is missing required tags: %v",
        [resource.address, resource.type, concat(", ", missing)]
    )
}

# Validate Environment tag values
valid_environments = [
    "dev",
    "test",
    "staging",
    "prod"
]

deny[msg] {
    resource := tfplan.root_module.resources[_]
    is_taggable(resource)
    env := resource.values.tags.Environment
    count([e | e := valid_environments[_]; e == lower(env)]) == 0
    
    msg := sprintf(
        "%s has invalid Environment tag value: %s. Must be one of: %v",
        [resource.address, env, concat(", ", valid_environments)]
    )
}

# Ensure ManagedBy tag indicates Terraform
deny[msg] {
    resource := tfplan.root_module.resources[_]
    is_taggable(resource)
    managedBy := resource.values.tags.ManagedBy
    lower(managedBy) != "terraform"
    
    msg := sprintf(
        "%s has incorrect ManagedBy tag. Should be 'terraform' but is '%s'",
        [resource.address, managedBy]
    )
}

# Validate CostCenter tag format (should be department-id format)
deny[msg] {
    resource := tfplan.root_module.resources[_]
    is_taggable(resource)
    costCenter := resource.values.tags.CostCenter
    not regex.match(`^[A-Z]+-[0-9]+$`, costCenter)
    
    msg := sprintf(
        "%s has invalid CostCenter tag format: %s. Should match pattern 'DEPT-123'",
        [resource.address, costCenter]
    )
}
