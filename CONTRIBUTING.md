# Contributing to Terraform Advanced Project

Thank you for your interest in contributing to our Terraform project! This document provides guidelines and instructions for contributing.

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [Getting Started](#getting-started)
3. [Development Workflow](#development-workflow)
4. [Coding Standards](#coding-standards)
5. [Commit Message Guidelines](#commit-message-guidelines)
6. [Pull Request Process](#pull-request-process)
7. [Release Process](#release-process)

## Code of Conduct

This project adheres to a Code of Conduct that promotes a welcoming and inclusive environment. By participating, you are expected to uphold this code.

## Getting Started

### Prerequisites

- Terraform (v1.1.0 or newer)
- Azure CLI
- Git

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

3. Create a workspace (if needed):
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
   terraform validate
   terraform plan -var-file=environments/dev.tfvars
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
