# Changelog

All notable changes to this Terraform project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.0] - 2025-05-26

### Added
- Validation module for Azure resource name compliance checks
- Comprehensive tagging module with cost center and owner information
- Changelog file to track project evolution
- Additional documentation in README

### Changed
- Updated main.tf to use the validation and tagging modules
- Improved README with detailed usage instructions and best practices

## [0.3.0] - 2025-05-26

### Added
- Naming convention module for consistent resource naming
- Environment-specific variable files for dev and prod
- Support for Terraform workspaces
- Enhanced README with architecture overview

### Changed
- Updated main.tf to use the naming module
- Standardized tags across all resources
- Refactored variables.tf with additional naming parameters

## [0.2.0] - 2025-05-26

### Added
- Modularized structure with network, storage, webapp, and keyvault modules
- Module-specific variables and outputs
- Enhanced output variables for all created resources

### Changed
- Refactored main.tf to use modules instead of direct resource definitions
- Updated outputs.tf to reference module outputs

## [0.1.0] - 2025-05-26

### Added
- Initial Terraform configuration with Azure resources:
  - Resource Group
  - Virtual Network and Subnet
  - Network Security Group
  - Storage Account and Container
  - App Service Plan and Linux Web App
  - Key Vault
- Backend configuration for state management in Azure Storage
- Basic variables and outputs
