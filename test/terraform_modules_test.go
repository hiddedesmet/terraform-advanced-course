package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNetworkModule tests the network module in isolation
func TestNetworkModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping network module test.")
	}

	// Create a test resource group first
	resourceGroupName := fmt.Sprintf("rg-network-test-%s", uniqueID)
	location := "westeurope"

	// Create resource group for testing
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/network",
		Vars: map[string]interface{}{
			"subscription_id":         subscriptionID,
			"resource_group_name":     resourceGroupName,
			"location":                location,
			"virtual_network_name":    fmt.Sprintf("vnet-test-%s", uniqueID),
			"address_space":           []string{"10.0.0.0/16"},
			"subnet_name":             fmt.Sprintf("subnet-test-%s", uniqueID),
			"subnet_address_prefixes": []string{"10.0.1.0/24"},
			"nsg_name":                fmt.Sprintf("nsg-test-%s", uniqueID),
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Resource group would normally be created by main infrastructure
	// For this test, we assume it exists or will be created by the module

	// Deploy the network module
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	vnetName := terraform.Output(t, terraformOptions, "virtual_network_name")
	subnetName := terraform.Output(t, terraformOptions, "subnet_name")
	nsgName := terraform.Output(t, terraformOptions, "nsg_name")

	// Verify outputs are not empty
	assert.NotEmpty(t, vnetName)
	assert.NotEmpty(t, subnetName)
	assert.NotEmpty(t, nsgName)

	// Verify network configuration
	actualVnet, err := azure.GetVirtualNetworkE(vnetName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	assert.Equal(t, vnetName, *actualVnet.Name)
	assert.Equal(t, "10.0.0.0/16", (*actualVnet.AddressSpace.AddressPrefixes)[0])
}

// TestStorageModule tests the storage module in isolation
func TestStorageModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	// Ensure storage account name is valid (max 24 chars, must be lowercase)
	storageAccountSuffix := strings.ToLower(uniqueID)
	if len(storageAccountSuffix) > 10 {
		storageAccountSuffix = storageAccountSuffix[:10]
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping storage module test.")
	}

	resourceGroupName := fmt.Sprintf("rg-storage-test-%s", uniqueID)
	location := "westeurope"

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/storage",
		Vars: map[string]interface{}{
			"subscription_id":          subscriptionID,
			"resource_group_name":      resourceGroupName,
			"location":                 location,
			"storage_account_name":     fmt.Sprintf("sttest%s", storageAccountSuffix),
			"storage_container_name":   fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
			"account_tier":             "Standard",
			"account_replication_type": "LRS",
			"container_access_type":    "private",
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
	containerName := terraform.Output(t, terraformOptions, "storage_container_name")

	// Verify storage account exists and has correct properties
	actualStorageAccount, err := azure.GetStorageAccountE(storageAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	assert.Equal(t, storageAccountName, *actualStorageAccount.Name)
	assert.Equal(t, "Standard_LRS", string(actualStorageAccount.Sku.Name))

	// Verify container name is returned
	assert.NotEmpty(t, containerName)
}

// TestWebAppModule tests the web app module in isolation
func TestWebAppModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping webapp module test.")
	}

	resourceGroupName := fmt.Sprintf("rg-webapp-test-%s", uniqueID)
	location := "westeurope"

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/webapp",
		Vars: map[string]interface{}{
			"subscription_id":       subscriptionID,
			"resource_group_name":   resourceGroupName,
			"location":              location,
			"app_service_plan_name": fmt.Sprintf("asp-test-%s", uniqueID),
			"web_app_name":          fmt.Sprintf("webapp-test-%s", uniqueID),
			"os_type":               "Linux",
			"sku_name":              "B1",
			"https_only":            true,
			"minimum_tls_version":   "1.2",
			"php_version":           "8.0",
			"app_settings": map[string]string{
				"WEBSITE_RUN_FROM_PACKAGE": "1",
			},
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	webAppName := terraform.Output(t, terraformOptions, "web_app_name")
	appServicePlanName := terraform.Output(t, terraformOptions, "app_service_plan_name")

	// Verify web app exists and has correct properties
	actualWebApp := azure.GetAppService(t, webAppName, resourceGroupName, subscriptionID)
	assert.Equal(t, webAppName, *actualWebApp.Name)
	assert.True(t, *actualWebApp.HTTPSOnly)

	// Verify app service plan
	assert.NotEmpty(t, appServicePlanName)
}

// TestKeyVaultModule tests the key vault module in isolation
func TestKeyVaultModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping keyvault module test.")
	}

	resourceGroupName := fmt.Sprintf("rg-kv-test-%s", uniqueID)
	location := "westeurope"

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/keyvault",
		Vars: map[string]interface{}{
			"subscription_id":            subscriptionID,
			"resource_group_name":        resourceGroupName,
			"location":                   location,
			"key_vault_name":             fmt.Sprintf("kv-test-%s", uniqueID),
			"sku_name":                   "standard",
			"purge_protection_enabled":   false,
			"soft_delete_retention_days": 7,
			"enable_rbac_authorization":  true,
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	keyVaultName := terraform.Output(t, terraformOptions, "key_vault_name")

	// Verify key vault exists
	actualKeyVault, err := azure.GetKeyVaultE(t, keyVaultName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	assert.NotNil(t, actualKeyVault)
	assert.Equal(t, keyVaultName, *actualKeyVault.Name)
}

// TestTaggingModule tests the tagging module
func TestTaggingModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/tagging",
		Vars: map[string]interface{}{
			"environment":       "test",
			"project_name":      "terratest-project",
			"owner":             "test-team",
			"cost_center":       "12345",
			"terraform_version": "v1.12",
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs and verify tags
	tags := terraform.OutputMap(t, terraformOptions, "tags")

	assert.Equal(t, "test", tags["Environment"])
	assert.Equal(t, "terratest-project", tags["Project"])
	assert.Equal(t, "test-team", tags["Owner"])
	assert.Equal(t, "12345", tags["CostCenter"])
	assert.Equal(t, "v1.12", tags["TerraformVersion"])
	assert.Contains(t, tags, "CreationDateTime")
}

// TestModulesIntegration tests multiple modules working together
func TestModulesIntegration(t *testing.T) {
	t.Parallel()

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping integration test.")
	}

	// Test naming module first
	namingOptions := &terraform.Options{
		TerraformDir: "../modules/naming",
		Vars: map[string]interface{}{
			"prefix":       "tf",
			"environment":  "test",
			"suffix":       "01",
			"project_name": "integration-test",
			"tags": map[string]string{
				"Team": "DevOps",
			},
		},
	}

	defer terraform.Destroy(t, namingOptions)
	terraform.InitAndApply(t, namingOptions)

	// Get naming outputs
	resourceGroupName := terraform.Output(t, namingOptions, "resource_group")
	storageAccountName := terraform.Output(t, namingOptions, "storage_account")

	// Test tagging module
	taggingOptions := &terraform.Options{
		TerraformDir: "../modules/tagging",
		Vars: map[string]interface{}{
			"environment":       "test",
			"project_name":      "integration-test",
			"owner":             "test-team",
			"cost_center":       "12345",
			"terraform_version": "v1.12",
		},
	}

	defer terraform.Destroy(t, taggingOptions)
	terraform.InitAndApply(t, taggingOptions)

	// Verify naming conventions
	assert.Contains(t, resourceGroupName, "tf")
	assert.Contains(t, resourceGroupName, "test")
	assert.Contains(t, resourceGroupName, "01")

	assert.Contains(t, storageAccountName, "tf")
	assert.Contains(t, storageAccountName, "test")

	// Verify storage account name compliance (lowercase, no special chars, 3-24 chars)
	assert.Equal(t, strings.ToLower(storageAccountName), storageAccountName)
	assert.True(t, len(storageAccountName) >= 3 && len(storageAccountName) <= 24)
	assert.True(t, isAlphanumericModules(storageAccountName))
}

// Helper function to check if string contains only alphanumeric characters
func isAlphanumericModules(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}
