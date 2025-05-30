package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestTerraformValidation tests that our Terraform configuration is valid
func TestTerraformValidation(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// Use minimal validation config that doesn't require Azure credentials
		TerraformDir: "./fixtures/minimal-validation",
		// No variables needed for validation
	}

	// Run terraform init first to install modules
	terraform.Init(t, terraformOptions)

	// Then run terraform validate
	terraform.Validate(t, terraformOptions)
}

// TestNamingConventions tests our naming convention module
func TestNamingConventions(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/naming",
		Vars: map[string]interface{}{
			"prefix":       "tf",
			"environment":  "dev",
			"suffix":       "01",
			"project_name": "naming-test",
			"tags": map[string]string{
				"Team": "DevOps",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs from naming module
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group")
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account")
	webAppName := terraform.Output(t, terraformOptions, "web_app")
	keyVaultName := terraform.Output(t, terraformOptions, "key_vault")

	// Test naming conventions
	assert.Contains(t, resourceGroupName, "tf")
	assert.Contains(t, resourceGroupName, "dev")
	assert.Contains(t, resourceGroupName, "01")
	// Note: Project name is included in tags, not in resource names

	assert.Contains(t, storageAccountName, "tf")
	assert.Contains(t, storageAccountName, "dev")

	assert.Contains(t, webAppName, "tf")
	assert.Contains(t, webAppName, "dev")

	assert.Contains(t, keyVaultName, "tf")
	assert.Contains(t, keyVaultName, "dev")

	// Test storage account name compliance (lowercase, no special chars, 3-24 chars)
	assert.Equal(t, strings.ToLower(storageAccountName), storageAccountName)
	assert.True(t, len(storageAccountName) >= 3 && len(storageAccountName) <= 24)
	assert.True(t, isAlphanumeric(storageAccountName))

	// Test key vault name compliance (3-24 chars, alphanumeric and hyphens only)
	assert.True(t, len(keyVaultName) >= 3 && len(keyVaultName) <= 24)
}

// TestValidationModule tests the validation module
// NOTE: This test is excluded from CI because it requires Azure provider initialization
// Run locally with: go test -v ./test -run "TestValidationModule"
func TestValidationModule(t *testing.T) {
	t.Parallel()

	// Test with valid names
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/validation",
		Vars: map[string]interface{}{
			"resource_group_name":    "valid-rg-name",
			"storage_account_name":   "validstorageaccount",
			"key_vault_name":         "valid-kv-name",
			"web_app_name":           "valid-webapp-name",
			"virtual_network_name":   "valid-vnet-name",
			"subnet_name":            "valid-subnet-name",
			"nsg_name":               "valid-nsg-name",
			"storage_container_name": "valid-container",
			"app_service_plan_name":  "valid-plan-name",
		},
	}

	// This should succeed with valid names
	terraform.InitAndPlan(t, terraformOptions)

	// Now test with invalid storage account name (too long)
	invalidOptions := &terraform.Options{
		TerraformDir: "../modules/validation",
		Vars: map[string]interface{}{
			"resource_group_name":    "valid-rg-name",
			"storage_account_name":   "thisstorageaccountnameistoolongandwillfail",
			"key_vault_name":         "valid-kv-name",
			"web_app_name":           "valid-webapp-name",
			"virtual_network_name":   "valid-vnet-name",
			"subnet_name":            "valid-subnet-name",
			"nsg_name":               "valid-nsg-name",
			"storage_container_name": "valid-container",
			"app_service_plan_name":  "valid-plan-name",
		},
	}

	// Plan with invalid data should succeed but output should show validation failed
	terraform.InitAndPlan(t, invalidOptions)
	
	// Apply to get outputs and check validation result
	terraform.Apply(t, invalidOptions)
	defer terraform.Destroy(t, invalidOptions)
	
	// Check that validation correctly identified the invalid storage account name
	isValid := terraform.Output(t, invalidOptions, "is_valid")
	assert.Equal(t, "false", isValid, "Expected validation to fail for invalid storage account name")
}

// Helper function to check if string contains only alphanumeric characters
func isAlphanumeric(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}
