# GitHub Workflows Documentation

This directory contains GitHub Actions workflows for automated Terraform infrastructure deployment.

## Workflows

### 1. terraform-deploy.yml
Main deployment workflow that handles:
- **Terraform validation and formatting checks**
- **Environment-specific planning** (dev/prod)
- **Automated deployment** based on branch triggers
- **Manual deployment** via workflow dispatch
- **Infrastructure destruction** (manual only)

#### Triggers
- **Push to `develop` branch**: Deploys to development environment
- **Push to `main` branch**: Deploys to production environment
- **Manual dispatch**: Allows choosing environment and action (plan/apply/destroy)
- **Pull requests**: Runs validation only

#### Environments
- **Development**: Triggered by `develop` branch
- **Production**: Triggered by `main` branch

### 2. terraform-pr-validation.yml
Pull request validation workflow that:
- **Validates Terraform configuration**
- **Checks formatting**
- **Generates plans for both environments**
- **Posts results as PR comments**

## Setup Requirements

### 1. Azure Service Principal
Create an Azure Service Principal with appropriate permissions:

```bash
# Create service principal
az ad sp create-for-rbac --name "terraform-github-actions" \
  --role="Contributor" \
  --scopes="/subscriptions/YOUR_SUBSCRIPTION_ID"
```

### 2. GitHub Secrets
Add the following secrets to your GitHub repository:

| Secret Name | Description |
|-------------|-------------|
| `AZURE_CLIENT_ID` | Service Principal Application ID |
| `AZURE_CLIENT_SECRET` | Service Principal Password |
| `AZURE_SUBSCRIPTION_ID` | Azure Subscription ID |
| `AZURE_TENANT_ID` | Azure Tenant ID |

### 3. GitHub Environments
Create the following environments in your GitHub repository:
- `development`
- `production`

Optionally, configure environment protection rules for production.

## Branch Strategy

```
main (production)
├── develop (development)
├── feature/xyz
└── hotfix/xyz
```

- **main**: Production deployments
- **develop**: Development deployments
- **feature/***: Development testing via PR
- **hotfix/***: Emergency fixes

## Workflow Features

### Security
- Uses GitHub environments for deployment protection
- Separates dev and prod deployments
- Requires manual approval for production (configurable)
- Uses least-privilege Azure service principal

### Validation
- Terraform format checking
- Configuration validation
- Plan generation and review
- PR comment integration

### Flexibility
- Manual workflow dispatch for ad-hoc deployments
- Support for both environments
- Plan, apply, and destroy operations
- Artifact storage for plans

## Usage Examples

### Deploying to Development
```bash
# Push to develop branch
git checkout develop
git commit -m "Update infrastructure"
git push origin develop
```

### Deploying to Production
```bash
# Push to main branch (typically via PR)
git checkout main
git merge develop
git push origin main
```

### Manual Deployment
1. Go to GitHub Actions
2. Select "Terraform Deploy" workflow
3. Click "Run workflow"
4. Choose environment and action
5. Click "Run workflow"

### Emergency Destruction
1. Go to GitHub Actions
2. Select "Terraform Deploy" workflow
3. Click "Run workflow"
4. Choose environment and "destroy" action
5. Confirm and run

## Monitoring

### Workflow Status
- All workflow runs are visible in the Actions tab
- Failed deployments send notifications
- Plan outputs are stored as artifacts

### Terraform State
- State is stored in Azure Storage (configured in backend.tf)
- Workspaces separate dev and prod state
- State locking prevents concurrent modifications

## Troubleshooting

### Common Issues

1. **Authentication Failures**
   - Verify Azure secrets are correct
   - Check service principal permissions
   - Ensure subscription ID is correct

2. **State Lock Issues**
   - Wait for concurrent operations to complete
   - Manually break locks if necessary (with caution)

3. **Plan Failures**
   - Check Terraform configuration syntax
   - Verify resource quotas in Azure
   - Review variable files for correct values

### Debug Steps
1. Check workflow logs in GitHub Actions
2. Verify Terraform configuration locally
3. Test Azure authentication with Azure CLI
4. Review Terraform state in Azure Storage

## Best Practices

1. **Always use pull requests** for production changes
2. **Review plans carefully** before applying
3. **Test in development** before promoting to production
4. **Use semantic commit messages** for better tracking
5. **Keep secrets secure** and rotate regularly
6. **Monitor costs** and resource usage
7. **Use tags consistently** for resource management
