# Te## Architecture

This project deploys the following resources:
- Resource Group
- Virtual Network with Subnet
- Network Security Group with rules for SSH and HTTP
- Storage Account with Container
- App Service Plan and Linux Web App
- Key Vault

For detailed architecture diagrams, see the [diagrams documentation](./docs/diagrams/README.md).

## Modules

The code is organized into the following modules:

- **Network**: Virtual Network, Subnet, and Network Security Group
- **Storage**: Storage Account and Container
- **WebApp**: App Service Plan and Linux Web App
- **KeyVault**: Key Vault
- **Naming**: Resource naming convention
- **Tagging**: Consistent resource tagging
- **Validation**: Input validation rules

Each module has its own detailed README with usage examples.

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
├── docs/                    # Project documentation
│   ├── adr/                 # Architecture Decision Records
│   ├── diagrams/            # Architecture diagrams
│   ├── MODULE_README_TEMPLATE.md  # Template for module documentation
│   └── PRE_COMMIT_HOOKS.md  # Guide for pre-commit hooks
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

## CI/CD and Testing

This project includes a comprehensive CI/CD pipeline with multiple types of tests:

### Automated Testing Levels

1. **Validation Tests** (always run):
   - Terraform syntax validation
   - Go code formatting checks  
   - Basic validation logic tests
   - No Azure credentials required

2. **Azure Module Tests** (conditional):
   - Individual module testing against Azure
   - Requires Azure credentials
   - Only runs when explicitly requested

3. **Integration Tests** (scheduled):
   - Full infrastructure deployment tests
   - End-to-end validation
   - Runs on scheduled builds

### Triggering Azure Tests

Azure tests only run in the following scenarios:

- **Scheduled builds**: Nightly at 2 AM UTC
- **Pull requests**: Add the `test-azure` label
- **Commits**: Include `[test-azure]` in commit message

To run Azure tests locally, ensure you have:
```bash
# Azure CLI authentication
az login

# Required environment variables
export ARM_CLIENT_ID="your-client-id"
export ARM_CLIENT_SECRET="your-client-secret"
export ARM_TENANT_ID="your-tenant-id" 
export AZURE_SUBSCRIPTION_ID="your-subscription-id"

# Run tests
go test -v ./test/ -timeout 30m
```

For more testing details, see [test/README.md](./test/README.md).

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

## Documentation

This project includes comprehensive documentation:

- **[Architecture Decision Records](./docs/adr/README.md)**: Documents significant architectural decisions
- **[Architecture Diagrams](./docs/diagrams/README.md)**: Visual representations of the infrastructure
- **[Pre-commit Hooks Guide](./docs/PRE_COMMIT_HOOKS.md)**: How to use the pre-commit hooks
- **[Module Documentation](./docs/MODULE_README_TEMPLATE.md)**: Template for module documentation
- **[Upgrade Guide](./UPGRADE_GUIDE.md)**: Instructions for upgrading between versions

## Best Practices

This project follows these best practices:

1. **Modular Structure**: Resources are organized into logical, reusable modules
2. **Standardized Naming**: Consistent naming conventions across all resources
3. **Input Validation**: Validation of input variables to catch errors early
4. **Documentation**: Comprehensive documentation for all modules and components
5. **Code Quality**: Enforced through pre-commit hooks and automated tests
6. **Environment Separation**: Different configurations for development and production
7. **State Management**: Remote state with proper locking and encryption
8. **CI/CD Integration**: Automated testing and deployment workflows

For more information on contributing to this project, see [CONTRIBUTING.md](./CONTRIBUTING.md).

## Advanced Features

This project implements several advanced features that make it stand out:

### 1. Automated Infrastructure Visualization

Our project automatically generates architecture diagrams from the Terraform code, ensuring documentation is always up-to-date with the actual infrastructure. See [Architecture Diagrams](./docs/diagrams/README.md).

### 2. Cost Estimation and Optimization

We've implemented automated cost estimation that provides:
- Monthly cost projections for all environments
- Cost optimization recommendations
- Historical cost tracking

See [Cost Management Documentation](./docs/costs/README.md) for details.

### 3. Infrastructure Drift Detection

Our advanced drift detection system automatically identifies differences between Terraform state and actual infrastructure, helping prevent configuration drift and ensure infrastructure integrity. See [Drift Detection](./docs/drift/README.md).

### 4. Policy as Code

We enforce organizational standards and security best practices using Open Policy Agent (OPA) policies:
- Security requirements
- Resource tagging standards
- Naming conventions
- Cost controls

See [Policy Documentation](./policies/README.md) for details.

### 5. Comprehensive Documentation

We maintain extensive documentation:
- Architecture Decision Records
- Module usage examples
- Architecture diagrams
- Cost and optimization guidance

### 6. DevOps Integration

The project is fully integrated with modern DevOps practices:
- Pre-commit hooks for code quality
- Comprehensive CI/CD pipelines
- Automated testing
- Drift detection in CI/CD