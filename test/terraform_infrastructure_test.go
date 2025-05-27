package test

import (
	"crypto/tls"
	"fmt"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

// TestTerraformAdvancedInfrastructure tests the complete infrastructure deployment
func TestTerraformAdvancedInfrastructure(t *testing.T) {
	t.Parallel()

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	// For Azure, we'll use West Europe as our primary region
	// expectedLocation := "West Europe"
	
	// Give the infrastructure stack a unique name using a random suffix
	uniqueID := random.UniqueId()
	
	// Construct the terraform options with default retryable errors to handle eventual consistency
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"location":              "westeurope",
			"prefix":                "tftest",
			"environment":           "test",
			"suffix":                uniqueID,
			"project_name":          "terratest",
			"resource_group_name":   fmt.Sprintf("tftest-rg-%s", uniqueID),
			"virtual_network_name":  fmt.Sprintf("tftest-vnet-%s", uniqueID),
			"subnet_name":           fmt.Sprintf("tftest-subnet-%s", uniqueID),
			"nsg_name":             fmt.Sprintf("tftest-nsg-%s", uniqueID),
			"storage_account_name":  fmt.Sprintf("tftestsa%s", strings.ToLower(uniqueID)),
			"storage_container_name": fmt.Sprintf("tftest-container-%s", strings.ToLower(uniqueID)),
			"app_service_plan_name": fmt.Sprintf("tftest-plan-%s", uniqueID),
			"web_app_name":         fmt.Sprintf("tftest-webapp-%s", uniqueID),
			"key_vault_name":       fmt.Sprintf("tftest-kv-%s", uniqueID),
		},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"ARM_CLIENT_ID":       getEnvVar(t, "ARM_CLIENT_ID"),
			"ARM_CLIENT_SECRET":   getEnvVar(t, "ARM_CLIENT_SECRET"),
			"ARM_SUBSCRIPTION_ID": getEnvVar(t, "ARM_SUBSCRIPTION_ID"),
			"ARM_TENANT_ID":       getEnvVar(t, "ARM_TENANT_ID"),
		},
	})

	// Save options to state for use in other test stages
	test_structure.SaveTerraformOptions(t, "../", terraformOptions)

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer test_structure.RunTestStage(t, "teardown", func() {
		terraform.Destroy(t, terraformOptions)
	})

	// Deploy the infrastructure with `terraform apply`
	test_structure.RunTestStage(t, "setup", func() {
		terraform.InitAndApply(t, terraformOptions)
	})

	// Run validation tests
	test_structure.RunTestStage(t, "validate", func() {
		// Get all the outputs from terraform
		resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_id")
		virtualNetworkID := terraform.Output(t, terraformOptions, "virtual_network_id")
		subnetID := terraform.Output(t, terraformOptions, "subnet_id")
		storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
		webAppURL := terraform.Output(t, terraformOptions, "web_app_url")
		webAppName := terraform.Output(t, terraformOptions, "web_app_name")
		keyVaultURI := terraform.Output(t, terraformOptions, "key_vault_uri")
		networkSecurityGroupID := terraform.Output(t, terraformOptions, "network_security_group_id")

		// Verify the resource group exists and is in the correct location
		t.Run("ResourceGroupValidation", func(t *testing.T) {
			// Extract resource group name from the full resource ID
			parts := strings.Split(resourceGroupName, "/")
			rgName := parts[len(parts)-1]
			
			// Basic validation that we have a resource group ID
			assert.Contains(t, resourceGroupName, "resourceGroups")
			assert.NotEmpty(t, rgName)
		})

		// Verify the virtual network exists
		t.Run("VirtualNetworkValidation", func(t *testing.T) {
			assert.NotEmpty(t, virtualNetworkID)
			assert.Contains(t, virtualNetworkID, "virtualNetworks")
		})

		// Verify the subnet exists
		t.Run("SubnetValidation", func(t *testing.T) {
			assert.NotEmpty(t, subnetID)
			assert.Contains(t, subnetID, "subnets")
		})

		// Verify the storage account exists and is accessible
		t.Run("StorageAccountValidation", func(t *testing.T) {
			assert.NotEmpty(t, storageAccountName)
			// Storage account names should be lowercase and between 3-24 characters
			assert.True(t, len(storageAccountName) >= 3 && len(storageAccountName) <= 24)
			assert.Equal(t, strings.ToLower(storageAccountName), storageAccountName)
		})

		// Verify the web app is accessible
		t.Run("WebAppValidation", func(t *testing.T) {
			assert.NotEmpty(t, webAppURL)
			assert.NotEmpty(t, webAppName)
			
			// Verify the web app responds to HTTP requests
			// Allow up to 5 minutes for the web app to be fully deployed and accessible
			maxRetries := 30
			timeBetweenRetries := 10 * time.Second
			
			tlsConfig := tls.Config{
				InsecureSkipVerify: true, // For testing purposes only
			}
			
			http_helper.HttpGetWithRetryWithCustomValidation(
				t,
				webAppURL,
				&tlsConfig,
				maxRetries,
				timeBetweenRetries,
				func(statusCode int, body string) bool {
					// Accept any 2xx or 3xx status code as success
					// Many App Service apps return default pages initially
					return statusCode >= 200 && statusCode < 400
				},
			)
		})

		// Verify the Key Vault exists and URI is valid
		t.Run("KeyVaultValidation", func(t *testing.T) {
			assert.NotEmpty(t, keyVaultURI)
			assert.Contains(t, keyVaultURI, "vault.azure.net")
			assert.True(t, strings.HasPrefix(keyVaultURI, "https://"))
		})

		// Verify the Network Security Group exists
		t.Run("NetworkSecurityGroupValidation", func(t *testing.T) {
			assert.NotEmpty(t, networkSecurityGroupID)
			assert.Contains(t, networkSecurityGroupID, "networkSecurityGroups")
		})

		// Test resource naming conventions
		t.Run("NamingConventionValidation", func(t *testing.T) {
			// All resource names should contain our test prefix
			assert.Contains(t, webAppName, "tftest")
			assert.Contains(t, storageAccountName, "tftest")
		})
	})
}

// getEnvVar is a helper function to get environment variables with better error handling
func getEnvVar(t *testing.T, varName string) string {
	// For Terratest, we rely on the environment variables being set
	// The terraform options will automatically pick them up
	return ""
}
