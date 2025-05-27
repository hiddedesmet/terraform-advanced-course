package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestNetworkModule tests just the network module in isolation
func TestNetworkModule(t *testing.T) {
	t.Parallel()

	// uniqueID := random.UniqueId()
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/network",
		
		Vars: map[string]interface{}{
			"location":              "West Europe",
			"resource_group_name":   "test-rg",
			"virtual_network_name":  "test-vnet",
			"subnet_name":          "test-subnet",
			"nsg_name":             "test-nsg",
			"virtual_network_address_space": []string{"10.0.0.0/16"},
			"subnet_address_prefixes": []string{"10.0.1.0/24"},
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Validate outputs
	vnetID := terraform.Output(t, terraformOptions, "virtual_network_id")
	subnetID := terraform.Output(t, terraformOptions, "subnet_id")
	nsgID := terraform.Output(t, terraformOptions, "network_security_group_id")

	assert.NotEmpty(t, vnetID)
	assert.NotEmpty(t, subnetID)
	assert.NotEmpty(t, nsgID)
	
	// Verify resource types in the IDs
	assert.Contains(t, vnetID, "virtualNetworks")
	assert.Contains(t, subnetID, "subnets")
	assert.Contains(t, nsgID, "networkSecurityGroups")
}

// TestStorageModule tests just the storage module in isolation
func TestStorageModule(t *testing.T) {
	t.Parallel()

	uniqueID := strings.ToLower(random.UniqueId())
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		
		Vars: map[string]interface{}{
			"location":               "West Europe",
			"resource_group_name":    "test-rg",
			"storage_account_name":   "testsa" + uniqueID,
			"storage_container_name": "test-container",
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Validate outputs
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
	storageAccountEndpoint := terraform.Output(t, terraformOptions, "storage_account_primary_blob_endpoint")

	assert.NotEmpty(t, storageAccountName)
	assert.NotEmpty(t, storageAccountEndpoint)
	
	// Verify naming conventions
	assert.True(t, len(storageAccountName) >= 3 && len(storageAccountName) <= 24)
	assert.Equal(t, strings.ToLower(storageAccountName), storageAccountName)
	assert.Contains(t, storageAccountEndpoint, "blob.core.windows.net")
}

// TestWebAppModule tests just the web app module in isolation
func TestWebAppModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/webapp",
		
		Vars: map[string]interface{}{
			"location":             "West Europe",
			"resource_group_name":  "test-rg",
			"app_service_plan_name": "test-plan-" + uniqueID,
			"web_app_name":         "test-webapp-" + uniqueID,
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Validate outputs
	webAppURL := terraform.Output(t, terraformOptions, "web_app_url")
	webAppName := terraform.Output(t, terraformOptions, "web_app_name")
	webAppID := terraform.Output(t, terraformOptions, "web_app_id")

	assert.NotEmpty(t, webAppURL)
	assert.NotEmpty(t, webAppName)
	assert.NotEmpty(t, webAppID)
	
	// Verify URL format
	assert.True(t, strings.HasPrefix(webAppURL, "https://"))
	assert.Contains(t, webAppURL, ".azurewebsites.net")
	assert.Contains(t, webAppID, "sites")
}

// TestKeyVaultModule tests just the key vault module in isolation
func TestKeyVaultModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/keyvault",
		
		Vars: map[string]interface{}{
			"location":          "West Europe",
			"resource_group_name": "test-rg",
			"key_vault_name":    "test-kv-" + uniqueID,
			"tenant_id":         "00000000-0000-0000-0000-000000000000", // Placeholder tenant ID for testing
			"tags": map[string]string{
				"Environment": "test",
				"Purpose":     "terratest",
			},
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Validate outputs
	keyVaultURI := terraform.Output(t, terraformOptions, "key_vault_uri")

	assert.NotEmpty(t, keyVaultURI)
	assert.True(t, strings.HasPrefix(keyVaultURI, "https://"))
	assert.Contains(t, keyVaultURI, ".vault.azure.net")
}
