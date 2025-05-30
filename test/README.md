# Test Configuration

This directory contains comprehensive Terratest tests for the Terraform Advanced Course infrastructure.

## Test Structure

### Test Files

1. **terraform_validation_test.go** - Basic validation tests that don't require Azure deployment
2. **terraform_modules_test.go** - Individual module testing
3. **terraform_infrastructure_test_new.go** - End-to-end infrastructure tests
4. **terraform_security_test.go** - Security and compliance tests
5. **terraform_performance_test.go** - Performance and scalability tests
6. **terraform_disaster_recovery_test.go** - Disaster recovery and backup tests
7. **terraform_test_suite.go** - Main test suite coordinator

### Test Categories

#### Validation Tests (No Azure Required)
- Terraform syntax validation
- Naming convention validation
- Variable validation
- Module interface validation

#### Module Tests (Azure Required)
- Network module functionality
- Storage module functionality
- Web App module functionality
- Key Vault module functionality
- Tagging module functionality
- Integration between modules

#### Infrastructure Tests (Azure Required)
- Full infrastructure deployment
- Multi-environment deployment
- Performance benchmarking
- Resource verification

#### Security Tests (Azure Required)
- Security compliance checks
- Encryption verification
- Access control validation
- Compliance tag verification

#### Performance Tests (Azure Required)
- Deployment time benchmarks
- Scalability testing
- Resource limit testing
- Concurrent deployment testing

#### Disaster Recovery Tests (Azure Required)
- Multi-region deployment
- Backup and restore scenarios
- Data retention testing
- Cross-region replication

## Prerequisites

### Required Tools
- Go 1.21+
- Terraform 1.8.0+
- Azure CLI (authenticated)

### Required Environment Variables

For tests that deploy to Azure:
```bash
export AZURE_SUBSCRIPTION_ID="your-subscription-id"
export ARM_CLIENT_ID="your-client-id"
export ARM_CLIENT_SECRET="your-client-secret"
export ARM_TENANT_ID="your-tenant-id"
```

### Optional Environment Variables
```bash
export AZURE_LOCATION="East US"  # Default test location
export TEST_TIMEOUT="30m"        # Test timeout
```

## Running Tests

### Run All Tests
```bash
go test -v ./test/...
```

### Run Specific Test Categories

#### Validation Tests Only (No Azure Required)
```bash
go test -v ./test/ -run TestTerraformValidation
go test -v ./test/ -run TestNamingConventions
go test -v ./test/ -run TestValidationModule
```

#### Module Tests
```bash
go test -v ./test/ -run TestNetworkModule
go test -v ./test/ -run TestStorageModule
go test -v ./test/ -run TestWebAppModule
go test -v ./test/ -run TestKeyVaultModule
```

#### Infrastructure Tests
```bash
go test -v ./test/ -run TestTerraformAdvancedInfrastructure
go test -v ./test/ -run TestTerraformInfrastructureWithDifferentEnvironments
```

#### Security Tests
```bash
go test -v ./test/ -run TestSecurityCompliance
go test -v ./test/ -run TestDataEncryption
go test -v ./test/ -run TestAccessControl
```

#### Performance Tests
```bash
go test -v ./test/ -run TestPerformanceBenchmarks
go test -v ./test/ -run TestScalabilityLimits
```

#### Disaster Recovery Tests
```bash
go test -v ./test/ -run TestDisasterRecovery
go test -v ./test/ -run TestBackupAndRestore
```

### Run Tests with Specific Flags

#### Skip Long-Running Tests
```bash
go test -v -short ./test/...
```

#### Run Tests with Timeout
```bash
go test -v -timeout 30m ./test/...
```

#### Run Tests in Parallel
```bash
go test -v -parallel 4 ./test/...
```

#### Run Specific Test Suite
```bash
go test -v ./test/ -run TestTerraformTestSuite
```

## Test Configuration

### Resource Naming
All test resources use unique identifiers to avoid conflicts:
- Format: `{resource-type}-{test-name}-{unique-id}`
- Example: `rg-network-test-abc123`

### Resource Cleanup
All tests include proper cleanup using `defer terraform.Destroy()` to ensure resources are cleaned up even if tests fail.

### Test Isolation
Tests are designed to run in parallel without conflicts by using unique resource names and separate resource groups.

## Cost Considerations

**Warning**: Tests that deploy to Azure will incur costs. Consider:

1. Use Azure free tier or development subscription
2. Run tests during off-peak hours
3. Monitor resource usage
4. Clean up resources promptly
5. Use cheapest SKUs (B1 for App Service, Standard_LRS for Storage)

## Troubleshooting

### Common Issues

#### Authentication Errors
```bash
# Ensure Azure CLI is authenticated
az login
az account show

# Verify environment variables
echo $AZURE_SUBSCRIPTION_ID
echo $ARM_CLIENT_ID
```

#### Resource Quota Limits
- Check Azure quotas in your subscription
- Use different regions if quota is exceeded
- Clean up existing resources

#### Test Timeouts
- Increase timeout with `-timeout` flag
- Check Azure service health
- Verify network connectivity

#### Resource Conflicts
- Ensure unique resource names
- Check for orphaned resources from previous test runs
- Clean up manually if necessary

### Debug Mode
```bash
# Enable Terraform debug logging
export TF_LOG=DEBUG
go test -v ./test/...
```

### Cleanup Orphaned Resources
```bash
# List all resource groups with test prefix
az group list --query "[?starts_with(name, 'rg-test-') || starts_with(name, 'rg-perf-') || starts_with(name, 'rg-security-')].name" -o table

# Delete specific test resource group
az group delete --name "rg-test-abc123" --yes --no-wait
```

## Best Practices

1. **Always use unique identifiers** for resource names
2. **Include proper cleanup** in all tests
3. **Use appropriate test categories** (unit, integration, e2e)
4. **Mock external dependencies** when possible
5. **Test both success and failure scenarios**
6. **Use meaningful test names** and descriptions
7. **Keep tests independent** and isolated
8. **Monitor test execution time** and optimize as needed
9. **Use test fixtures** for common configurations
10. **Document test assumptions** and requirements
