# Terratest Documentation

This document provides comprehensive information about the Terratest implementation in our Terraform project.

## Overview

[Terratest](https://github.com/gruntwork-io/terratest) is a Go library that provides patterns and helper functions for testing infrastructure. Our implementation includes:

- **End-to-end infrastructure testing**
- **Individual module testing**
- **Validation and compliance testing**
- **Automated CI/CD integration**

## Test Structure

```
test/
├── go.mod                           # Go module dependencies
├── terraform_infrastructure_test.go # End-to-end infrastructure tests
├── terraform_modules_test.go        # Individual module tests
└── terraform_validation_test.go     # Validation and planning tests
```

## Test Categories

### 1. Validation Tests (`terraform_validation_test.go`)
- **Purpose**: Fast feedback on configuration validity
- **Runtime**: ~2-5 minutes
- **Azure Cost**: $0 (no resources deployed)

#### Tests Included:
- `TestTerraformValidation`: Validates Terraform syntax and configuration
- `TestTerraformPlan`: Generates and validates Terraform execution plans
- `TestNamingConventions`: Tests naming convention module
- `TestValidationModule`: Tests resource name validation rules

### 2. Module Tests (`terraform_modules_test.go`)
- **Purpose**: Test individual modules in isolation
- **Runtime**: ~5-15 minutes per module
- **Azure Cost**: Low (small test resources)

#### Tests Included:
- `TestNetworkModule`: Tests VNet, subnet, and NSG creation
- `TestStorageModule`: Tests storage account and container creation
- `TestWebAppModule`: Tests App Service Plan and Web App creation
- `TestKeyVaultModule`: Tests Key Vault creation

### 3. Infrastructure Tests (`terraform_infrastructure_test.go`)
- **Purpose**: Full end-to-end infrastructure deployment
- **Runtime**: ~20-30 minutes
- **Azure Cost**: Moderate (complete infrastructure stack)

#### Tests Included:
- `TestTerraformAdvancedInfrastructure`: Complete infrastructure deployment and validation

## Running Tests

### Prerequisites

1. **Go 1.21+** installed
2. **Terraform 1.5+** installed
3. **Azure CLI** configured with authentication
4. **Environment variables** set:
   ```bash
   export ARM_CLIENT_ID="your-service-principal-id"
   export ARM_CLIENT_SECRET="your-service-principal-secret"
   export ARM_SUBSCRIPTION_ID="your-subscription-id"
   export ARM_TENANT_ID="your-tenant-id"
   ```

### Local Testing

#### Using Make (Recommended)
```bash
# Check environment
make check-env

# Setup dependencies
make setup

# Run validation tests only (fastest)
make test

# Run all tests
make test-all

# Run specific test categories
make test-validation
make test-modules
make test-infrastructure
```

#### Direct Go Commands
```bash
# Navigate to test directory
cd test

# Download dependencies
go mod download

# Run specific tests
go test -v -run TestTerraformValidation -timeout 10m
go test -v -run TestNetworkModule -timeout 20m
go test -v -run TestTerraformAdvancedInfrastructure -timeout 45m

# Run all tests
go test -v ./... -timeout 45m
```

### CI/CD Testing

Tests are automatically triggered by:
- **Push to `main`/`develop`**: Validation tests + applicable deployment tests
- **Pull Requests**: Validation tests only
- **Manual Workflow Dispatch**: Choose test type

#### GitHub Environments
- **testing**: Required for infrastructure tests (can include approval gates)

## Test Features

### Parallel Execution
Most tests run in parallel (`t.Parallel()`) to reduce total execution time.

### Unique Resource Naming
Tests use `random.UniqueId()` to generate unique resource names, preventing conflicts.

### Automatic Cleanup
All tests include `defer terraform.Destroy()` to ensure resources are cleaned up.

### Staged Testing
Infrastructure tests use staged execution:
1. **Setup**: Deploy infrastructure
2. **Validate**: Run validation checks
3. **Teardown**: Clean up resources

### Comprehensive Validation
Tests validate:
- Resource existence and properties
- Naming conventions compliance
- Network connectivity
- HTTP endpoint accessibility
- Azure-specific constraints

## Test Configuration

### Custom Variables
Tests use environment-specific variables to avoid conflicts with existing resources:

```go
Vars: map[string]interface{}{
    "location":              "westeurope",
    "prefix":                "tftest",
    "environment":           "test",
    "suffix":                uniqueID,
    "resource_group_name":   fmt.Sprintf("tftest-rg-%s", uniqueID),
    // ... more variables
}
```

### Retry Logic
Tests include retry mechanisms for eventually consistent operations:
- HTTP endpoint checks with exponential backoff
- Azure resource state validation
- Network connectivity tests

## Best Practices

### 1. Test Isolation
- Each test uses unique resource names
- Tests can run in parallel without conflicts
- Proper cleanup prevents resource leaks

### 2. Fast Feedback
- Validation tests run first (fastest feedback)
- Module tests provide middle-tier validation
- Infrastructure tests provide comprehensive validation

### 3. Cost Management
- Tests use minimal resource sizes (Free/Basic tiers)
- Automatic cleanup prevents resource accumulation
- Clear cost warnings for expensive tests

### 4. Realistic Testing
- Tests use actual Azure services
- Network and security configurations are validated
- Real-world scenarios are tested

## Troubleshooting

### Common Issues

#### Authentication Failures
```bash
# Verify Azure CLI login
az account show

# Check service principal permissions
az role assignment list --assignee $ARM_CLIENT_ID

# Test authentication
az login --service-principal -u $ARM_CLIENT_ID -p $ARM_CLIENT_SECRET --tenant $ARM_TENANT_ID
```

#### Resource Quota Issues
- Tests may fail if subscription has insufficient quota
- Use different regions if quota is exhausted
- Consider using smaller resource sizes

#### Test Timeouts
- Increase timeout values for slow operations
- Check Azure service health status
- Verify network connectivity

### Debug Tips

1. **Enable Terraform Debug Output**:
   ```bash
   export TF_LOG=DEBUG
   go test -v -run TestName
   ```

2. **Run Tests Individually**:
   ```bash
   go test -v -run TestSpecificTest -timeout 30m
   ```

3. **Check Azure Portal**:
   - Monitor resource creation progress
   - Verify resource configurations
   - Check activity logs for errors

## Performance Optimization

### Parallel Test Execution
Tests are designed to run in parallel where possible:
- Validation tests: All parallel
- Module tests: Parallel by module type
- Infrastructure tests: Sequential (resource intensive)

### Resource Reuse
Future enhancements could include:
- Test environment pooling
- Resource reuse across test runs
- Optimized cleanup strategies

## Cost Considerations

| Test Type | Estimated Cost | Duration | Resources |
|-----------|---------------|----------|-----------|
| Validation | $0 | 2-5 min | None (plan only) |
| Module Tests | $1-5 | 10-20 min | Minimal resources |
| Infrastructure | $5-15 | 20-30 min | Full stack |

### Cost Optimization Tips
- Run validation tests frequently
- Run module tests on feature branches
- Run infrastructure tests on main branch only
- Use scheduled cleanup jobs for orphaned resources

## Future Enhancements

1. **Performance Testing**: Load testing for web applications
2. **Security Testing**: Vulnerability scanning and compliance checks
3. **Disaster Recovery**: Backup and restore testing
4. **Multi-Region**: Cross-region deployment testing
5. **Cost Analysis**: Automated cost estimation and tracking
6. **Integration Testing**: API and database connectivity tests

## Resources

- [Terratest Documentation](https://terratest.gruntwork.io/)
- [Azure Testing Best Practices](https://learn.microsoft.com/en-us/azure/developer/terraform/best-practices-end-to-end-testing)
- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Terraform Testing Guide](https://developer.hashicorp.com/terraform/tutorials/configuration-language/test)
