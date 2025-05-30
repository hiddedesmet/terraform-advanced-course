# ADR 0002: Naming Conventions

## Status

Accepted

## Context

Azure resources need to follow specific naming constraints and we wanted to establish consistent naming across all resources for better organization and management.

## Decision

We have chosen to implement a dedicated naming module that:

1. Defines abbreviations for all resource types (e.g., rg for resource group)
2. Enforces a standard naming pattern: `{prefix}-{resource_type}-{environment}-{suffix}`
3. Handles special cases like storage accounts that have restricted character sets
4. Provides a centralized means to generate resource names

## Consequences

### Positive

- **Consistency**: All resources follow the same naming pattern
- **Compliance**: Names conform to Azure-specific constraints
- **Readability**: Clear indication of resource purpose, environment, and ownership
- **Maintainability**: Naming patterns can be updated in one place

### Negative

- **Transition**: Existing resources may need to be renamed
- **Complexity**: Additional module dependency for all resources

## Implementation Notes

Special considerations were made for:
- Storage accounts: No hyphens, only alphanumeric
- Global uniqueness: Including unique identifiers where required
- Length constraints: Ensuring generated names fit within Azure limits
