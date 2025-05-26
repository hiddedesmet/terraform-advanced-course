# Terraform Advanced Course Demo

This Terraform project demonstrates Azure infrastructure deployment with multiple resources.

## Resources Created

This project creates the following Azure resources:

- **Resource Group**: Contains all other resources
- **Virtual Network**: Network infrastructure with 10.0.0.0/16 address space
- **Subnet**: Subnet within the VNet with 10.0.1.0/24 address space
- **Network Security Group**: Security rules for SSH and HTTP access
- **Storage Account**: Blob storage with a demo container
- **App Service Plan**: Linux-based hosting plan (Free tier)
- **Web App**: Node.js web application
- **Key Vault**: Secure storage for secrets and keys

## Variables

All resource names are parameterized using variables for consistency and flexibility. You can customize the deployment by modifying the variables in `variables.tf`:

### Core Variables:
- `resource_group_name`: Name of the resource group
- `location`: Azure region for deployment (default: westeurope)

### Globally Unique Resource Names:
- `storage_account_name`: Storage account name (must be globally unique)
- `web_app_name`: Web app name (must be globally unique)
- `key_vault_name`: Key Vault name (must be globally unique)

### Other Resource Names:
- `virtual_network_name`: Name of the virtual network
- `subnet_name`: Name of the subnet
- `nsg_name`: Name of the network security group
- `storage_container_name`: Name of the storage container
- `app_service_plan_name`: Name of the App Service Plan

## Usage

1. Initialize Terraform:
   ```bash
   terraform init
   ```

2. Plan the deployment:
   ```bash
   terraform plan
   ```

3. Apply the configuration:
   ```bash
   terraform apply
   ```

4. To destroy resources when done:
   ```bash
   terraform destroy
   ```

## Outputs

After deployment, the following information will be displayed:

- Resource Group ID
- Virtual Network ID
- Subnet ID
- Storage Account details
- Web App URL
- Key Vault URI
- Network Security Group ID

## Backend Configuration

This project uses Azure Storage as the backend for state management. The backend configuration is in `backend.tf`.