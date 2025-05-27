# Contributing to Terraform Advanced Project

Thank you for your interest in contributing to our Terraform project! This document provides guidelines and instructions for contributing.

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [Getting Started](#getting-started)
3. [Development Workflow](#development-workflow)
4. [Testing Requirements](#testing-requirements)
5. [Coding Standards](#coding-standards)
6. [Commit Message Guidelines](#commit-message-guidelines)
7. [Pull Request Process](#pull-request-process)
8. [Release Process](#release-process)

## Code of Conduct

This project adheres to a Code of Conduct that promotes a welcoming and inclusive environment. By participating, you are expected to uphold this code.

## Testing Requirements

This project uses Terratest for comprehensive infrastructure testing. All contributions should maintain test compatibility and, where applicable, include appropriate tests.

### Running Tests

```bash
# Quick validation (recommended for development)
make test

# Complete test suite (use sparingly due to cost)
make test-all

# Specific test categories
make test-validation         # Fast syntax and plan validation
make test-modules           # Module-specific functionality tests
make test-infrastructure    # Full deployment tests (incurs Azure costs)

# Development helpers
make setup                  # Initialize testing environment
make fmt                    # Format test code
make lint                   # Quality checks
make clean                  # Clean test artifacts
```

### Test Guidelines

1. **Always run validation tests** before submitting PRs
2. **Module tests** should be run when modifying specific modules
3. **Infrastructure tests** should be used sparingly due to cost implications
4. **Format your test code** using `make fmt`
5. **Ensure tests pass** in CI/CD pipeline before merging

## Getting Started

### Prerequisites

- Terraform (v1.1.0 or newer)
- Azure CLI
- Git
- Go (v1.21 or newer) for running tests
- Make (for using standardized development workflows)

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-organization/terraform-advanced-course.git
   cd terraform-advanced-course
   ```

2. Initialize Terraform:
   ```bash
   terraform init
   ```

3. Set up Go for testing:
   ```bash
   cd test/
   go mod tidy
   ```

4. Configure Azure authentication:
   ```bash
   export ARM_CLIENT_ID="your-client-id"
   export ARM_CLIENT_SECRET="your-client-secret"
   export ARM_SUBSCRIPTION_ID="your-subscription-id"
   export ARM_TENANT_ID="your-tenant-id"
   ```

5. Create a workspace (if needed):
   ```bash
   terraform workspace new dev
   ```

## Development Workflow

1. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes, adhering to the coding standards.

3. Test your changes:
   ```bash
   # Terraform validation
   terraform validate
   terraform plan -var-file=environments/dev.tfvars
   
   # Run automated tests
   make test                    # Quick validation tests
   make test-modules            # Module-specific tests
   # Note: Avoid running infrastructure tests locally unless necessary
   # as they deploy real resources and incur costs
   ```

4. Update documentation (README.md, module READMEs) if necessary.

5. Update the CHANGELOG.md with your changes under the [Unreleased] section.

6. Commit and push your changes (see Commit Message Guidelines).

7. Create a pull request (see Pull Request Process).

## Coding Standards

### Module Structure

Each module should have:
- `main.tf` - Primary resource definitions
- `variables.tf` - Input variable declarations
- `outputs.tf` - Output declarations
- `README.md` - Documentation for the module

### Naming Conventions

- Resources: `azurerm_<resource_type>_<name>`
- Variables: Snake case (e.g., `resource_group_name`)
- Outputs: Snake case, descriptive names
- Modules: Snake case, reflecting functionality

### Style Guidelines

- Indent using 2 spaces
- Use lowercase for all resource names
- Use descriptive names for resources, variables, and outputs
- Use snake_case for naming
- Include comprehensive descriptions for all variables and outputs
- Group related resources together
- Use locals for values used multiple times

## Commit Message Guidelines

Follow the [Conventional Commits](https://www.conventionalcommits.org/) standard:

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

Types:
- feat: A new feature
- fix: A bug fix
- docs: Documentation changes
- style: Changes that don't affect code functionality
- refactor: Code refactoring without changing functionality
- perf: Performance improvements
- test: Adding or modifying tests
- chore: Changes to build process or auxiliary tools

Example:
```
feat(network): add support for network peering

Added the ability to peer virtual networks across different resource groups.
```

## Pull Request Process

1. Ensure your code follows our coding standards.
2. Update documentation as necessary.
3. Update the CHANGELOG.md with your changes.
4. Make sure all tests pass.
5. Get at least one code review approval.

## Release Process

1. Update the CHANGELOG.md by moving [Unreleased] changes to a new version section.
2. Tag the release in Git:
   ```bash
   git tag -a v0.x.0 -m "Release version 0.x.0"
   git push origin v0.x.0
   ```
3. Create a GitHub release with release notes.
