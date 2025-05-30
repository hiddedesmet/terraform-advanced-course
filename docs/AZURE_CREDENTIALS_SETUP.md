# Azure Credentials Setup for CI/CD

This document explains how to set up Azure credentials for running the full test suite in GitHub Actions.

## Required GitHub Secrets

To run Azure-dependent tests, configure the following secrets in your GitHub repository:

1. Go to **Settings** → **Secrets and variables** → **Actions**
2. Add the following secrets:

| Secret Name | Description | Example |
|-------------|-------------|---------|
| `AZURE_CREDENTIALS` | Service principal credentials in JSON format | `{"clientId":"xxx","clientSecret":"xxx","subscriptionId":"xxx","tenantId":"xxx"}` |
| `AZURE_SUBSCRIPTION_ID` | Azure subscription ID | `12345678-1234-1234-1234-123456789012` |
| `ARM_CLIENT_ID` | Service principal client ID | `12345678-1234-1234-1234-123456789012` |
| `ARM_CLIENT_SECRET` | Service principal client secret | `your-secret-value` |
| `ARM_TENANT_ID` | Azure tenant ID | `12345678-1234-1234-1234-123456789012` |

## Creating a Service Principal

1. **Create the service principal**:
   ```bash
   az ad sp create-for-rbac --name "terraform-ci-cd" \
     --role contributor \
     --scopes /subscriptions/{subscription-id} \
     --sdk-auth
   ```

2. **Copy the JSON output** and use it as the `AZURE_CREDENTIALS` secret

3. **Extract individual values** for the other secrets from the JSON

## Testing Without Azure Credentials

The CI/CD pipeline is designed to work without Azure credentials by default:

- **Validation tests** always run and test syntax, formatting, and basic logic
- **Azure module tests** only run when explicitly requested
- **Integration tests** are reserved for scheduled builds and special conditions

This ensures that:
- Contributors can submit PRs without Azure access
- Basic validation catches most issues early
- Full Azure testing is available when needed

## Local Development

For local development with Azure tests:

```bash
# Authenticate with Azure CLI
az login

# Set environment variables (optional, CLI auth is usually sufficient)
export ARM_CLIENT_ID="your-client-id"
export ARM_CLIENT_SECRET="your-client-secret"  
export ARM_TENANT_ID="your-tenant-id"
export AZURE_SUBSCRIPTION_ID="your-subscription-id"

# Run all tests
go test -v ./test/ -timeout 30m

# Run only validation tests (no Azure required)
go test -v ./test/ -run "TestTerraformValidation|TestNamingConventions" -timeout 10m
```

## Security Notes

- Use separate service principals for different environments
- Follow the principle of least privilege
- Regularly rotate secrets
- Monitor service principal usage in Azure
- Consider using federated credentials for enhanced security

## Troubleshooting

### Common Issues

1. **"Login failed"**: Check that all required secrets are set correctly
2. **"No subscription"**: Verify `AZURE_SUBSCRIPTION_ID` is correct
3. **"Insufficient permissions"**: Ensure service principal has contributor role
4. **"Tenant mismatch"**: Verify `ARM_TENANT_ID` matches your Azure tenant

### Debug Steps

1. Check GitHub Actions logs for specific error messages
2. Verify service principal exists: `az ad sp list --display-name "terraform-ci-cd"`
3. Test locally with the same credentials
4. Ensure subscription is active and accessible
