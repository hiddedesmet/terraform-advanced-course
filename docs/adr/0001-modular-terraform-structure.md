# ADR 0001: Modular Terraform Structure

## Status

Accepted

## Context

When creating infrastructure as code with Terraform, we needed to decide between a monolithic approach (all resources defined in a single directory) versus a modular approach (resources organized into reusable modules).

## Decision

We have chosen to adopt a modular structure for our Terraform code. This means:

1. Creating separate modules for logical components (network, storage, webapp, keyvault)
2. Implementing naming conventions as a dedicated module
3. Using validation modules to enforce rules and constraints
4. Creating tagging conventions as a separate module

## Consequences

### Positive

- **Reusability**: Modules can be reused across different environments and projects
- **Maintainability**: Smaller, focused modules are easier to understand and maintain
- **Testing**: Individual modules can be tested in isolation
- **Versioning**: Modules can be versioned independently
- **Separation of concerns**: Clear boundaries between different infrastructure components

### Negative

- **Learning curve**: Higher initial complexity compared to monolithic approach
- **Module interfaces**: Need to define clear inputs/outputs between modules
- **Dependency management**: Must manage dependencies between modules

## Implementation Notes

Each module follows a standard structure:
- `main.tf`: Primary resource definitions
- `variables.tf`: Input variables with descriptions and validation
- `outputs.tf`: Output values that can be consumed by root module
