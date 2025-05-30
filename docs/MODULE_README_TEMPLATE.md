# Module Documentation Template

To improve consistency and quality of module documentation, use this template for all module README files.

## Template

```markdown
# [Module Name] Module

## Overview

Brief description of what this module does and why it exists.

## Features

- Feature 1
- Feature 2
- Feature 3

## Usage

```hcl
module "example" {
  source = "./modules/[module_name]"
  
  # Required variables
  required_variable = "value"
  
  # Optional variables with defaults
  optional_variable = "custom_value"
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
| [resource_name](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/resource_type) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| variable_name | Description of variable | `string` | `"default"` | no |

## Outputs

| Name | Description |
|------|-------------|
| output_name | Description of output |

## Notes

Any additional information or considerations when using this module.
```

## Implementation

1. Create a README.md file in each module directory
2. Fill in the template with module-specific information
3. Include working usage examples
4. Keep documentation updated when module changes

## Example

See the `network` module README for a complete example.
