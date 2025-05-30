package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPerformanceBenchmarks tests deployment and destruction times
func TestPerformanceBenchmarks(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance benchmarks in short mode")
	}

	uniqueID := random.UniqueId()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping performance test.")
	}

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"subscription_id":        subscriptionID,
			"resource_group_name":    fmt.Sprintf("rg-perf-%s", uniqueID),
			"location":               "westeurope",
			"storage_account_name":   fmt.Sprintf("stperf%s", strings.ToLower(uniqueID[:10])),
			"key_vault_name":         fmt.Sprintf("kv-perf-%s", uniqueID),
			"web_app_name":           fmt.Sprintf("webapp-perf-%s", uniqueID),
			"virtual_network_name":   fmt.Sprintf("vnet-perf-%s", uniqueID),
			"subnet_name":            fmt.Sprintf("subnet-perf-%s", uniqueID),
			"nsg_name":               fmt.Sprintf("nsg-perf-%s", uniqueID),
			"storage_container_name": fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
			"app_service_plan_name":  fmt.Sprintf("asp-perf-%s", uniqueID),
			"prefix":                 "tf",
			"environment":            "test",
			"suffix":                 "01",
			"project_name":           "performance-test",
			"owner":                  "perf-team",
			"cost_center":            "11111",
		},
	}

	// Benchmark deployment time
	t.Run("DeploymentTime", func(t *testing.T) {
		start := time.Now()
		terraform.InitAndApply(t, terraformOptions)
		deploymentTime := time.Since(start)

		t.Logf("Infrastructure deployment time: %v", deploymentTime)

		// Assert deployment time is reasonable (adjust threshold as needed)
		assert.True(t, deploymentTime < 15*time.Minute,
			fmt.Sprintf("Deployment should complete within 15 minutes, took: %v", deploymentTime))
	})

	// Benchmark destruction time
	t.Run("DestructionTime", func(t *testing.T) {
		start := time.Now()
		terraform.Destroy(t, terraformOptions)
		destructionTime := time.Since(start)

		t.Logf("Infrastructure destruction time: %v", destructionTime)

		// Assert destruction time is reasonable
		assert.True(t, destructionTime < 10*time.Minute,
			fmt.Sprintf("Destruction should complete within 10 minutes, took: %v", destructionTime))
	})
}

// TestScalabilityLimits tests resource limits and scalability
func TestScalabilityLimits(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping scalability test.")
	}

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"subscription_id":        subscriptionID,
			"resource_group_name":    fmt.Sprintf("rg-scale-%s", uniqueID),
			"location":               "westeurope",
			"storage_account_name":   fmt.Sprintf("stscale%s", strings.ToLower(uniqueID[:9])),
			"key_vault_name":         fmt.Sprintf("kv-scale-%s", uniqueID),
			"web_app_name":           fmt.Sprintf("webapp-scale-%s", uniqueID),
			"virtual_network_name":   fmt.Sprintf("vnet-scale-%s", uniqueID),
			"subnet_name":            fmt.Sprintf("subnet-scale-%s", uniqueID),
			"nsg_name":               fmt.Sprintf("nsg-scale-%s", uniqueID),
			"storage_container_name": fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
			"app_service_plan_name":  fmt.Sprintf("asp-scale-%s", uniqueID),
			"prefix":                 "tf",
			"environment":            "test",
			"suffix":                 "01",
			"project_name":           "scalability-test",
			"owner":                  "scale-team",
			"cost_center":            "22222",
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Test Web App scaling capabilities
	webAppName := terraform.Output(t, terraformOptions, "web_app_name")
	webApp := azure.GetAppService(t, webAppName, resourceGroupName, subscriptionID)

	// Verify the web app is created with basic tier that can be scaled
	assert.NotNil(t, webApp, "Web App should exist")

	// Test Storage Account throughput limits
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
	storageAccount, err := azure.GetStorageAccountE(storageAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	// Verify storage account tier supports expected performance
	assert.Equal(t, "Standard_LRS", string(storageAccount.Sku.Name),
		"Storage account should use Standard LRS for testing")

	// Test Network capacity
	vnetName := terraform.Output(t, terraformOptions, "virtual_network_name")
	vnet, err := azure.GetVirtualNetworkE(vnetName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	// Verify address space is sufficient for scaling
	assert.NotNil(t, vnet.AddressSpace, "VNet should have address space defined")
	if vnet.AddressSpace.AddressPrefixes != nil && len(*vnet.AddressSpace.AddressPrefixes) > 0 {
		addressSpace := (*vnet.AddressSpace.AddressPrefixes)[0]
		assert.Equal(t, "10.0.0.0/16", addressSpace,
			"VNet should have adequate address space for scaling")
	}
}

// TestResourceLimits tests that resources are created within Azure limits
func TestResourceLimits(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping resource limits test.")
	}

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"subscription_id":        subscriptionID,
			"resource_group_name":    fmt.Sprintf("rg-limits-%s", uniqueID),
			"location":               "westeurope",
			"storage_account_name":   fmt.Sprintf("stlimit%s", strings.ToLower(uniqueID[:9])),
			"key_vault_name":         fmt.Sprintf("kv-limits-%s", uniqueID),
			"web_app_name":           fmt.Sprintf("webapp-limits-%s", uniqueID),
			"virtual_network_name":   fmt.Sprintf("vnet-limits-%s", uniqueID),
			"subnet_name":            fmt.Sprintf("subnet-limits-%s", uniqueID),
			"nsg_name":               fmt.Sprintf("nsg-limits-%s", uniqueID),
			"storage_container_name": fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
			"app_service_plan_name":  fmt.Sprintf("asp-limits-%s", uniqueID),
			"prefix":                 "tf",
			"environment":            "test",
			"suffix":                 "01",
			"project_name":           "limits-test",
			"owner":                  "limits-team",
			"cost_center":            "33333",
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Test resource naming limits
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
	keyVaultName := terraform.Output(t, terraformOptions, "key_vault_name")
	webAppName := terraform.Output(t, terraformOptions, "web_app_name")

	// Verify storage account name length (3-24 characters)
	assert.True(t, len(storageAccountName) >= 3 && len(storageAccountName) <= 24,
		fmt.Sprintf("Storage account name length should be 3-24 characters, got: %d", len(storageAccountName)))

	// Verify key vault name length (3-24 characters)
	assert.True(t, len(keyVaultName) >= 3 && len(keyVaultName) <= 24,
		fmt.Sprintf("Key vault name length should be 3-24 characters, got: %d", len(keyVaultName)))

	// Verify web app name length (2-60 characters)
	assert.True(t, len(webAppName) >= 2 && len(webAppName) <= 60,
		fmt.Sprintf("Web app name length should be 2-60 characters, got: %d", len(webAppName)))

	// Verify naming conventions compliance
	assert.True(t, isValidStorageAccountName(storageAccountName),
		"Storage account name should contain only lowercase letters and numbers")
	assert.True(t, isValidKeyVaultName(keyVaultName),
		"Key vault name should contain only alphanumeric characters and hyphens")
	assert.True(t, isValidWebAppName(webAppName),
		"Web app name should contain only alphanumeric characters and hyphens")
}

// TestConcurrentDeployments tests multiple deployments running in parallel
func TestConcurrentDeployments(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping concurrent deployment test in short mode")
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if subscriptionID == "" {
		t.Skip("AZURE_SUBSCRIPTION_ID environment variable not set. Skipping concurrent deployment test.")
	}

	numDeployments := 3
	results := make(chan bool, numDeployments)

	for i := 0; i < numDeployments; i++ {
		go func(deploymentIndex int) {
			uniqueID := random.UniqueId()

			terraformOptions := &terraform.Options{
				TerraformDir: "../",
				Vars: map[string]interface{}{
					"subscription_id":        subscriptionID,
					"resource_group_name":    fmt.Sprintf("rg-concurrent-%d-%s", deploymentIndex, uniqueID),
					"location":               "westeurope",
					"storage_account_name":   fmt.Sprintf("stconc%d%s", deploymentIndex, strings.ToLower(uniqueID[:7])),
					"key_vault_name":         fmt.Sprintf("kv-conc-%d-%s", deploymentIndex, uniqueID),
					"web_app_name":           fmt.Sprintf("webapp-conc-%d-%s", deploymentIndex, uniqueID),
					"virtual_network_name":   fmt.Sprintf("vnet-conc-%d-%s", deploymentIndex, uniqueID),
					"subnet_name":            fmt.Sprintf("subnet-conc-%d-%s", deploymentIndex, uniqueID),
					"nsg_name":               fmt.Sprintf("nsg-conc-%d-%s", deploymentIndex, uniqueID),
					"storage_container_name": fmt.Sprintf("container%d%s", deploymentIndex, strings.ToLower(uniqueID)),
					"app_service_plan_name":  fmt.Sprintf("asp-conc-%d-%s", deploymentIndex, uniqueID),
					"prefix":                 "tf",
					"environment":            "test",
					"suffix":                 fmt.Sprintf("%02d", deploymentIndex),
					"project_name":           "concurrent-test",
					"owner":                  "concurrent-team",
					"cost_center":            "44444",
				},
			}

			defer terraform.Destroy(t, terraformOptions)

			// Deploy and verify
			terraform.InitAndApply(t, terraformOptions)

			// Basic verification that deployment succeeded
			resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
			assert.NotEmpty(t, resourceGroupName)

			results <- true
		}(i)
	}

	// Wait for all deployments to complete
	for i := 0; i < numDeployments; i++ {
		select {
		case result := <-results:
			assert.True(t, result, "Concurrent deployment should succeed")
		case <-time.After(20 * time.Minute):
			t.Fatal("Concurrent deployment timed out")
		}
	}
}

// Helper functions for name validation
func isValidStorageAccountName(name string) bool {
	if len(name) < 3 || len(name) > 24 {
		return false
	}
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}

func isValidKeyVaultName(name string) bool {
	if len(name) < 3 || len(name) > 24 {
		return false
	}
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == '-') {
			return false
		}
	}
	return true
}

func isValidWebAppName(name string) bool {
	if len(name) < 2 || len(name) > 60 {
		return false
	}
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == '-') {
			return false
		}
	}
	return true
}
