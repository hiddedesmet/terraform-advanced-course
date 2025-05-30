# Test Manual Workflow Dispatch

This file demonstrates how to trigger the complete test flow.

To test the manual workflow dispatch feature:

## Quick Test (Validation Only)
1. Go to GitHub Actions
2. Click "Terratest CI/CD" 
3. Click "Run workflow"
4. Select "validation-only"
5. Click "Run workflow" button

Expected result: Only validation tests should run (~2-5 minutes)

## Full Test Flow
1. Go to GitHub Actions
2. Click "Terratest CI/CD"
3. Click "Run workflow" 
4. Select "all-tests"
5. Click "Run workflow" button

Expected execution order:
- ✅ Validation Tests (2-5 min)
- ✅ Module Tests (10-15 min) - requires Azure creds
- ✅ Infrastructure Tests (20-45 min) - requires Azure creds  
- ✅ Security Tests (15-30 min) - requires Azure creds
- ✅ Performance Tests (30-60 min) - requires Azure creds
- ✅ Complete Test Suite (45-90 min) - requires Azure creds

Total time: ~2-3 hours for complete flow

## Alternative: Commit Trigger
Commit with message containing `[test-azure]` to trigger module tests:

```bash
git commit -m "test: Demonstrate complete workflow [test-azure]"
git push origin main
```

This will run validation + module tests, and if they pass, continue with the full flow.
