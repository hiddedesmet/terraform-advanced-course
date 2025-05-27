# Te## Architecture

This project deploys the following resources:
- Resource Group
- Virtual Network with Subnet
- Network Security Group with rules for SSH and HTTP
- Storage Account with Container
- App Service Plan and Linux Web App
- Key Vault

## Modules

The code is organized into the following modules:

- **Network**: Virtual Network, Subnet, and Network Security Group
- **Storage**: Storage Account and Container
- **WebApp**: App Service Plan and Linux Web App
- **KeyVault**: Key Vault
- **Naming**: Resource naming convention

## Directory Structure

```
terraform-advanced-course/
├── backend.tf               # Backend configuration for Terraform state
├── main.tf                  # Main Terraform configuration
├── outputs.tf               # Output variables
├── variables.tf             # Input variables
├── CHANGELOG.md             # Project change history
├── UPGRADE_GUIDE.md         # Version upgrade instructions
├── environments/            # Environment-specific configurations
│   ├── dev.tfvars           # Development environment variables
│   └── prod.tfvars          # Production environment variables
├── .github/                 # GitHub Actions workflows
│   └── workflows/           # CI/CD automation
│       ├── terraform-deploy.yml           # Main deployment workflow
│       ├── terraform-pr-validation.yml    # PR validation workflow
│       └── README.md        # Workflow documentation
├── test/                    # Terratest infrastructure testing
│   ├── go.mod               # Go module dependencies
│   ├── go.sum               # Dependency checksums
│   ├── terraform_validation_test.go       # Validation and plan tests
│   ├── terraform_modules_test.go          # Module-specific tests
│   └── terraform_infrastructure_test.go   # End-to-end infrastructure tests
└── modules/                 # Terraform modules
    ├── keyvault/            # Key Vault module
    │   ├── main.tf
    │   ├── variables.tf
    │   └── outputs.tf
    ├── naming/              # Naming convention module
    │   ├── main.tf
    │   ├── variables.tf
    │   └── outputs.tf
    ├── network/             # Network module
    │   ├── main.tf
    │   ├── variables.tf
    │   └── outputs.tf
    ├── storage/             # Storage module
    │   ├── main.tf
    │   ├── variables.tf
    │   └── outputs.tf
    ├── tagging/             # Resource tagging module
    │   ├── main.tf
    │   ├── variables.tf
    │   └── outputs.tf
    ├── validation/          # Resource validation module
    │   ├── main.tf
    │   ├── variables.tf
    │   └── outputs.tf
    └── webapp/              # Web App module
        ├── main.tf
        ├── variables.tf
        └── outputs.tf
```ed Project

This repository contains Terraform code to deploy Azure infrastructure in a modular, maintainable way.

## Architecture

This project creates the following Azure resources:

- **Resource Group**: Contains all other resources
- **Virtual Network**: Network infrastructure with 10.0.0.0/16 address space
- **Subnet**: Subnet within the VNet with 10.0.1.0/24 address space
- **Network Security Group**: Security rules for SSH and HTTP access
- **Storage Account**: Blob storage with a demo container
- **App Service Plan**: Linux-based hosting plan (Free tier)
- **Web App**: Node.js web application
- **Key Vault**: Secure storage for secrets and keys

## Deployment Options

This project supports both automated CI/CD deployment via GitHub Actions and manual command-line deployment.

### Automated Deployment (Recommended)

The project includes GitHub Actions workflows for automated deployment:

#### Setup Requirements
1. **Azure Service Principal**: Create a service principal with Contributor access
2. **GitHub Secrets**: Configure Azure authentication secrets
3. **GitHub Environments**: Set up `development` and `production` environments

#### Branch Strategy
- Push to `develop` → Deploys to development environment
- Push to `main` → Deploys to production environment
- Pull requests → Automated validation and planning

#### Quick Start
```bash
# Deploy to development
git checkout develop
git commit -m "Update infrastructure"
git push origin develop

# Deploy to production (via PR)
git checkout main
git merge develop
git push origin main
```

See [`.github/workflows/README.md`](.github/workflows/README.md) for detailed setup instructions.

### Manual Deployment

For local development or troubleshooting, you can still deploy manually:

#### Initialize Terraform

```bash
terraform init
```

#### Select Workspace (Environment)

```bash
# Create and switch to a new workspace
terraform workspace new dev
# Or switch to an existing workspace
terraform workspace select dev
```

#### Plan and Apply

```bash
# For development environment
terraform plan -var-file=environments/dev.tfvars
terraform apply -var-file=environments/dev.tfvars

# For production environment
terraform workspace select prod
terraform plan -var-file=environments/prod.tfvars
terraform apply -var-file=environments/prod.tfvars
```

#### Destroy

```bash
terraform destroy -var-file=environments/dev.tfvars
```

## Testing

This project includes comprehensive testing using Terratest, a Go-based testing framework for infrastructure code.

### Test Categories

1. **Validation Tests** (`test/terraform_validation_test.go`)
   - Fast syntax and configuration validation
   - Terraform plan generation testing
   - Naming convention compliance testing
   - **Runtime**: < 1 second

2. **Module Tests** (`test/terraform_modules_test.go`)
   - Individual module isolation testing
   - Module input/output validation
   - Module functionality verification
   - **Runtime**: < 1 minute

3. **Infrastructure Tests** (`test/terraform_infrastructure_test.go`)
   - Full end-to-end deployment testing
   - Real Azure resource creation and validation
   - Resource connectivity and functionality testing
   - **Runtime**: < 30 minutes

### Running Tests

#### Prerequisites
1. **Go Installation** (version 1.21 or later):
   ```bash
   # macOS with Homebrew
   brew install go
   
   # Verify installation
   go version
   ```

2. **Azure Authentication**:
   ```bash
   export ARM_CLIENT_ID="your-client-id"
   export ARM_CLIENT_SECRET="your-client-secret"
   export ARM_SUBSCRIPTION_ID="your-subscription-id"
   export ARM_TENANT_ID="your-tenant-id"
   ```

#### Test Execution

```bash
# Navigate to test directory
cd test/

# Initialize Go modules (first time only)
go mod tidy

# Using Makefile (Recommended)
make help                   # Show all available targets
make test                   # Run validation tests only (fastest)
make test-all               # Run all test categories
make test-validation        # Run validation and planning tests
make test-modules           # Run individual module tests
make test-infrastructure    # Run full infrastructure tests (includes cost warning)

# Direct Go test execution
go test -v -timeout 30m

# Run specific test categories
go test -v -run TestTerraformValidation    # Fast validation tests
go test -v -run TestTerraformPlan         # Plan generation tests
go test -v -run TestNamingConventions     # Naming convention tests
go test -v -run TestValidationModule      # Module validation tests
go test -v -run TestTerraformAdvanced     # Infrastructure tests

# Development helpers
make setup                  # Initialize Go module and dependencies
make fmt                    # Format Go test code
make lint                   # Run linting on test code
make clean                  # Clean test cache
```

#### Test Structure
```
test/
├── go.mod                           # Go module dependencies
├── go.sum                           # Dependency checksums
├── terraform_validation_test.go     # Validation and plan tests
├── terraform_modules_test.go        # Module-specific tests
└── terraform_infrastructure_test.go # End-to-end infrastructure tests
```

#### Makefile Integration
The project includes a comprehensive Makefile for standardized development workflows:

```bash
make help                   # Show all available commands
make test                   # Quick validation tests
make test-all               # Complete test suite
make test-validation        # Syntax and plan validation
make test-modules           # Individual module testing
make test-infrastructure    # Full deployment testing (with cost warning)
make setup                  # Initialize development environment
make fmt                    # Format Go code
make lint                   # Code quality checks
make clean                  # Clean test artifacts
```

### Continuous Integration

Tests are automatically executed in GitHub Actions:
- **Pull Request Validation**: Runs validation and module tests on every PR
- **Branch Deployment**: Runs full test suite before deployment
- **Scheduled Testing**: Daily infrastructure validation tests
- **Makefile Integration**: Uses standardized Makefile targets for consistent execution across local and CI environments

> **Note**: Current GitHub Actions workflow focuses on Terraform validation and planning. To add Terratest execution to CI/CD, consider enhancing the workflow with additional test steps that call `make test-validation` and `make test-modules` for cost-effective automated testing.

## Best Practices Implemented

1. **Modular Structure**: Code is organized into reusable modules
2. **Environment Separation**: Using workspaces and environment-specific variable files
3. **Naming Conventions**: Consistent resource naming through the naming module
4. **Resource Tagging**: All resources are tagged with environment, project, and management information
5. **State Management**: Remote state stored in Azure Storage Account
6. **Variable Defaults**: Sensible defaults for variables with option to override
7. **CI/CD Integration**: Automated deployment via GitHub Actions
8. **Security**: Service Principal authentication and environment protection
9. **Code Review**: Pull request workflow with automated validation
10. **Comprehensive Testing**: Automated testing with Terratest for infrastructure validation

## Security Considerations

1. HTTPS-only web app
2. NSG rules limiting access
3. Key Vault with RBAC authorization
4. Private storage container

## Considerations for Production Use

1. Add more granular RBAC permissions
2. Implement network isolation with private endpoints
3. Add monitoring and diagnostics settings
4. Implement backup policies
5. Consider using Azure Policy for governance