# Test Execution Guide

This guide explains how to run the complete Terraform test flow from validation through to full infrastructure testing.

## Test Flow Overview

The testing pipeline follows this sequential flow:
1. **Validation Tests** - Terraform syntax, formatting, and basic validation (no Azure required)
2. **Module Tests** - Individual module testing with Azure resources
3. **Infrastructure Tests** - Full infrastructure deployment testing
4. **Security Tests** - Security compliance and policy validation
5. **Performance Tests** - Performance benchmarks and resource limits
6. **Complete Test Suite** - Comprehensive end-to-end testing

## How to Run the Complete Test Flow

### Option 1: Manual Workflow Dispatch (Recommended)

1. **Go to GitHub Actions**:
   - Navigate to your repository on GitHub
   - Click on "Actions" tab
   - Find "Terratest CI/CD" workflow

2. **Run Workflow Manually**:
   - Click "Run workflow" button
   - Select the branch (usually `main`)
   - Choose test suite from dropdown:

#### Test Suite Options:

| Option | What Runs | Use Case |
|--------|-----------|----------|
| `validation-only` | Only validation tests | Quick syntax/format checks |
| `module-tests` | Validation + Module tests | Test individual modules |
| `infrastructure-tests` | + Infrastructure tests | Test full deployments |
| `security-tests` | + Security tests | Add security compliance checks |
| `performance-tests` | + Performance tests | Add performance benchmarks |
| `complete-suite` | + Complete test suite | Full end-to-end testing |
| `all-tests` | Everything | Complete test coverage |

**For the full flow, select either `complete-suite` or `all-tests`**

### Option 2: Automatic Triggers

#### Scheduled Runs
- **Nightly at 2 AM UTC**: Runs complete test suite automatically
- **Only on main branch**: Full suite runs only on production branch

#### Commit Message Triggers
- Add `[test-azure]` to commit message to trigger module tests
- Example: `git commit -m "fix: Update storage module [test-azure]"`

#### Pull Request Labels
- Add labels to PR to trigger specific test suites:
  - `test-azure`: Module tests
  - `test-infrastructure`: Infrastructure tests
  - `test-security`: Security tests

### Option 3: Local Testing

For development and debugging, you can run tests locally:

```bash
# Run validation tests (no Azure required)
go test -v ./test/ -run "TestTerraformValidation|TestNamingConventions" -timeout 10m

# Run module tests (requires Azure credentials)
export ARM_CLIENT_ID="your-client-id"
export ARM_CLIENT_SECRET="your-client-secret"
export ARM_TENANT_ID="your-tenant-id"
export ARM_SUBSCRIPTION_ID="your-subscription-id"

go test -v ./test/ -run "TestValidationModule" -timeout 30m

# Run specific infrastructure tests
go test -v ./test/ -run "TestTerraformAdvancedInfrastructure" -timeout 45m
```

## Test Dependencies and Flow

```
┌─────────────────┐
│ Validation      │ ← Always runs first
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Module Tests    │ ← Requires: Validation success
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Infrastructure  │ ← Requires: Module Tests success
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Security Tests  │ ← Can run in parallel with Infrastructure
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Performance     │ ← Requires: Module Tests success
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Complete Suite  │ ← Requires: Validation + Module Tests
└─────────────────┘
```

## Required Credentials

### GitHub Secrets Required
For Azure tests to work, these secrets must be configured in GitHub:

```
AZURE_CLIENT_ID          # Service Principal Client ID
AZURE_CLIENT_SECRET      # Service Principal Client Secret
AZURE_TENANT_ID          # Azure AD Tenant ID
AZURE_SUBSCRIPTION_ID    # Target Azure Subscription ID
```

### Setup Instructions
See [AZURE_CREDENTIALS_SETUP.md](./AZURE_CREDENTIALS_SETUP.md) for detailed credential setup.

## Monitoring Test Execution

### Real-time Monitoring
1. Go to Actions tab in GitHub
2. Click on the running workflow
3. Monitor each job's progress in real-time
4. View logs for debugging failures

### Test Results
- ✅ **Green**: All tests passed
- ❌ **Red**: Tests failed (check logs)
- ⚪ **Skipped**: Job conditions not met
- 🟡 **In Progress**: Currently running

### Typical Execution Times
- **Validation**: ~2-5 minutes
- **Module Tests**: ~10-15 minutes
- **Infrastructure**: ~20-45 minutes
- **Security**: ~15-30 minutes
- **Performance**: ~30-60 minutes
- **Complete Suite**: ~45-90 minutes

## Troubleshooting

### Common Issues

1. **Azure Authentication Failures**
   - Verify GitHub secrets are properly configured
   - Check service principal permissions
   - Ensure subscription ID is correct

2. **Terraform State Conflicts**
   - Tests use isolated resource groups
   - Each test run creates unique resources
   - Cleanup happens automatically

3. **Timeout Issues**
   - Infrastructure tests can take 45+ minutes
   - Performance tests can take 60+ minutes
   - Increase timeout if needed

4. **Module Dependencies**
   - Some tests require specific modules to pass first
   - Check job dependencies in workflow file

### Debug Mode
Add debug output to workflow runs by setting debug logging:
- Go to repository Settings → Secrets → Variables
- Add `ACTIONS_STEP_DEBUG` = `true`
- Add `ACTIONS_RUNNER_DEBUG` = `true`

## Best Practices

1. **Development Workflow**:
   - Start with `validation-only` for quick feedback
   - Use `module-tests` when working on modules
   - Run `all-tests` before merging to main

2. **CI/CD Integration**:
   - Validation runs on every PR
   - Module tests trigger with `[test-azure]` tag
   - Complete suite runs nightly on main

3. **Resource Management**:
   - Tests create temporary resources
   - Automatic cleanup after test completion
   - Monitor Azure costs during development

4. **Performance**:
   - Run tests in parallel when possible
   - Use appropriate timeouts
   - Cache Go modules and Terraform providers

## Example: Running Complete Flow

To run the entire test flow manually:

1. **GitHub UI Method**:
   ```
   Actions → Terratest CI/CD → Run workflow → Select "all-tests"
   ```

2. **Commit Message Method**:
   ```bash
   git commit -m "feat: Major infrastructure update [test-azure]"
   git push origin main
   ```

3. **Expected Execution**:
   ```
   ✅ Validation Tests (2-5 min)
   ✅ Module Tests (10-15 min) 
   ✅ Infrastructure Tests (20-45 min)
   ✅ Security Tests (15-30 min)
   ✅ Performance Tests (30-60 min)
   ✅ Complete Test Suite (45-90 min)
   ```

The complete flow ensures your Terraform infrastructure is properly validated, secure, performant, and ready for production deployment.
