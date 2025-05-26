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
├── environments/            # Environment-specific configurations
│   ├── dev.tfvars           # Development environment variables
│   └── prod.tfvars          # Production environment variables
└── modules/                 # Terraform modules
    ├── keyvault/            # Key Vault module
    ├── naming/              # Naming convention module
    ├── network/             # Network module
    ├── storage/             # Storage module
    └── webapp/              # Web App module
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

## Usage

### Initialize Terraform

```bash
terraform init
```

### Select Workspace (Environment)

```bash
# Create and switch to a new workspace
terraform workspace new dev
# Or switch to an existing workspace
terraform workspace select dev
```

### Plan and Apply

```bash
# For development environment
terraform plan -var-file=environments/dev.tfvars
terraform apply -var-file=environments/dev.tfvars

# For production environment
terraform workspace select prod
terraform plan -var-file=environments/prod.tfvars
terraform apply -var-file=environments/prod.tfvars
```

### Destroy

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