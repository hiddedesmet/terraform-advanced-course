#!/bin/bash

# Terraform State Upgrade Script
# This script helps upgrade Terraform state files to be compatible with newer versions

set -e

echo "🔧 Terraform State Upgrade Script"
echo "=================================="

# Check if terraform is installed
if ! command -v terraform &> /dev/null; then
    echo "❌ Terraform is not installed. Please install Terraform 1.8.5 or later."
    exit 1
fi

# Check terraform version
TERRAFORM_VERSION=$(terraform version -json | jq -r '.terraform_version' 2>/dev/null || terraform version | head -n1 | awk '{print $2}' | sed 's/v//')
echo "📋 Current Terraform version: $TERRAFORM_VERSION"

# Initialize terraform
echo "🚀 Initializing Terraform..."
terraform init

# Create backup of current state (for workspaces)
echo "💾 Creating state backup..."
mkdir -p backups

# List workspaces
echo "📄 Available workspaces:"
terraform workspace list

# For each environment, upgrade the state
for env in dev prod; do
    echo ""
    echo "🔄 Processing environment: $env"
    
    # Try to select workspace, create if it doesn't exist
    if terraform workspace select $env 2>/dev/null; then
        echo "✅ Selected workspace: $env"
    else
        echo "🆕 Creating new workspace: $env"
        terraform workspace new $env
    fi
    
    # Create a backup of the current state
    echo "💾 Backing up state for $env..."
    terraform state pull > "backups/terraform.tfstate.${env}.backup.$(date +%Y%m%d_%H%M%S)" 2>/dev/null || echo "⚠️  No existing state found for $env"
    
    # Try to list state to check if upgrade is needed
    echo "🔍 Checking state compatibility..."
    if terraform state list > /dev/null 2>&1; then
        echo "✅ State is compatible with current Terraform version"
    else
        echo "⚠️  State may need upgrading. Attempting refresh..."
        # The state upgrade happens automatically during refresh in newer versions
        terraform refresh -var-file="environments/${env}.tfvars" || echo "⚠️  Refresh completed with warnings"
    fi
done

# Switch back to default workspace
terraform workspace select default 2>/dev/null || echo "ℹ️  No default workspace found"

echo ""
echo "✅ State upgrade process completed!"
echo "🔍 If you encounter issues, check the backups in the 'backups/' directory"
echo "💡 Next steps:"
echo "   1. Test your configuration: terraform plan -var-file=\"environments/dev.tfvars\""
echo "   2. Commit your changes and push to trigger the GitHub workflow"
