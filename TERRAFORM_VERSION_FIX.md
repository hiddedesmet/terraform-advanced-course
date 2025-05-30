# Terraform Version Compatibility Fix

## Problem
Your GitHub workflows were encountering the error:
```
Error loading state: 7 problems:
- unsupported checkable object kind "var"
```

This error occurs when there's a version incompatibility between the Terraform state file and the Terraform version being used in the workflow.

## Root Cause
- Your workflows were using Terraform v1.5.0
- Your configuration specifies Azure provider v4.x (`~> 4.0`)
- Azure provider v4.x requires Terraform v1.8.0 or higher
- The state file was likely created with a newer version, causing compatibility issues

## Changes Made

### 1. Updated Terraform Version Requirements
- **main.tf**: Updated `required_version` from `>= 1.1.0` to `>= 1.8.0`
- **All workflow files**: Updated Terraform version from `1.5.0` to `1.8.5`

### 2. Enhanced Workflows with State Compatibility Checks
Added state compatibility checking steps to all workflows:
- `terraform-deploy.yml`
- `terraform-pr-validation.yml` 
- `terraform-local-test.yml`

### 3. Created State Upgrade Script
- **scripts/upgrade_terraform_state.sh**: A comprehensive script to safely upgrade your state files locally

## Files Modified
```
├── main.tf (required_version updated)
├── .github/workflows/
│   ├── terraform-deploy.yml (TF version + state checks)
│   ├── terraform-pr-validation.yml (TF version + state checks)
│   └── terraform-local-test.yml (TF version updated)
└── scripts/
    └── upgrade_terraform_state.sh (new script)
```

## Next Steps

### 1. Local State Upgrade (Recommended)
Run the state upgrade script locally first:
```bash
cd /path/to/your/terraform/project
./scripts/upgrade_terraform_state.sh
```

### 2. Test Locally
```bash
# Install Terraform 1.8.5 locally
terraform version

# Test with dev environment
terraform init
terraform workspace select dev || terraform workspace new dev
terraform plan -var-file="environments/dev.tfvars"
```

### 3. Commit and Push
Once local testing is successful:
```bash
git add .
git commit -m "fix: upgrade Terraform version compatibility and add state upgrade handling"
git push
```

### 4. Monitor Workflow
- The workflows now include better error handling for state compatibility
- If issues persist, check the workflow logs for the new state compatibility check steps

## Key Improvements
1. **Version Alignment**: All components now use compatible versions
2. **Better Error Handling**: Workflows now detect and handle state compatibility issues
3. **Safety Measures**: Local upgrade script with automatic backups
4. **Future-Proofing**: Improved workflow structure for handling version upgrades

## Troubleshooting
If you still encounter issues:
1. Check that all secrets are properly configured in GitHub
2. Verify Azure authentication is working
3. Review the state compatibility check logs in the workflow
4. Consider manually recreating workspaces if the state is severely corrupted

The workflows should now run successfully with proper state handling!
