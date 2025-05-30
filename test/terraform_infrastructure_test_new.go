package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTerraformAdvancedInfrastructure validates the full infrastructure deployment
func TestTerraformAdvancedInfrastructure(t *testing.T) {
	t.Parallel()

	subscriptionID := getAzureSubscriptionID(t)
	uniqueID := strings.ToLower(random.UniqueId())
	resourceGroupName := fmt.Sprintf("rg-terratest-%s", uniqueID)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{"environments/dev.tfvars"},
		Vars: map[string]interface{}{
			"resource_group_name":     resourceGroupName,
			"storage_account_name":    fmt.Sprintf("st%s", strings.ToLower(uniqueID)),
			"key_vault_name":          fmt.Sprintf("kv-%s", uniqueID),
			"web_app_name":            fmt.Sprintf("webapp-%s", uniqueID),
			"virtual_network_name":    fmt.Sprintf("vnet-%s", uniqueID),
			"subnet_name":             fmt.Sprintf("subnet-%s", uniqueID),
			"nsg_name":                fmt.Sprintf("nsg-%s", uniqueID),
			"storage_container_name":  fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
			"app_service_plan_name":   fmt.Sprintf("asp-%s", uniqueID),
			"prefix":                  "tf",
			"environment":             "dev",
			"suffix":                  "01",
			"project_name":            "terratest",
			"owner":                   "test-team",
		},
		RetryableTerraformErrors: map[string]string{
			".*": "Will try again",
		},
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	resourceGroupName = terraform.Output(t, terraformOptions, "resource_group_name")
	storageAccountName := terraform.Output(t, terraformOptions, "storage_account_name")
	keyVaultName := terraform.Output(t, terraformOptions, "key_vault_name")
	webAppName := terraform.Output(t, terraformOptions, "web_app_name")
	vnetName := terraform.Output(t, terraformOptions, "virtual_network_name")

	// Verify resources exist in Azure
	assert.True(t, azure.ResourceGroupExists(t, resourceGroupName, subscriptionID))
	assert.True(t, azure.StorageAccountExists(t, storageAccountName, resourceGroupName, subscriptionID))
	assert.True(t, azure.AppExists(t, webAppName, resourceGroupName, subscriptionID))

	// Verify network resources
	actualVnet, err := azure.GetVirtualNetworkE(t, vnetName, resourceGroupName, subscriptionID)
	assert.NoError(t, err)
	assert.Equal(t, vnetName, *actualVnet.Name)

	// Verify storage account properties
	actualStorageAccount, err := azure.GetStorageAccountE(t, storageAccountName, resourceGroupName, subscriptionID)
	assert.NoError(t, err)
	assert.Equal(t, "Standard_LRS", string(actualStorageAccount.Sku.Name))

	// Verify web app properties
	actualWebApp := azure.GetAppService(t, webAppName, resourceGroupName, subscriptionID)
	assert.Equal(t, webAppName, *actualWebApp.Name)

	// Verify tags
	expectedTags := map[string]string{
		"Environment": "dev",
		"Project":     "terratest",
		"Owner":       "test-team",
	}
	
	for key, expectedValue := range expectedTags {
		if actualValue, exists := actualWebApp.Tags[key]; exists {
			assert.Equal(t, expectedValue, *actualValue)
		}
	}
}

// TestTerraformInfrastructureWithDifferentEnvironments tests infrastructure across environments
func TestTerraformInfrastructureWithDifferentEnvironments(t *testing.T) {
	t.Parallel()

	subscriptionID := getAzureSubscriptionID(t)
	environments := []string{"dev", "prod"}

	for _, env := range environments {
		env := env // capture range variable
		t.Run(env, func(t *testing.T) {
			t.Parallel()

			uniqueID := strings.ToLower(random.UniqueId())
			resourceGroupName := fmt.Sprintf("rg-terratest-%s-%s", env, uniqueID)

			terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: "../",
				VarFiles:     []string{fmt.Sprintf("environments/%s.tfvars", env)},
				Vars: map[string]interface{}{
					"resource_group_name":     resourceGroupName,
					"storage_account_name":    fmt.Sprintf("st%s%s", env, strings.ToLower(uniqueID)),
					"key_vault_name":          fmt.Sprintf("kv-%s-%s", env, uniqueID),
					"web_app_name":            fmt.Sprintf("webapp-%s-%s", env, uniqueID),
					"virtual_network_name":    fmt.Sprintf("vnet-%s-%s", env, uniqueID),
					"subnet_name":             fmt.Sprintf("subnet-%s-%s", env, uniqueID),
					"nsg_name":                fmt.Sprintf("nsg-%s-%s", env, uniqueID),
					"storage_container_name":  fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
					"app_service_plan_name":   fmt.Sprintf("asp-%s-%s", env, uniqueID),
					"prefix":                  "tf",
					"environment":             env,
					"suffix":                  "01",
					"project_name":            "terratest",
					"owner":                   "test-team",
				},
				RetryableTerraformErrors: map[string]string{
					".*": "Will try again",
				},
				MaxRetries:         3,
				TimeBetweenRetries: 5 * time.Second,
			})

			defer terraform.Destroy(t, terraformOptions)
			terraform.InitAndApply(t, terraformOptions)

			// Verify environment-specific tags
			resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
			rg, err := azure.GetResourceGroupE(t, resourceGroupName, subscriptionID)
			assert.NoError(t, err)
			
			// Verify environment tag
			if envTag, exists := rg.Tags["Environment"]; exists {
				assert.Equal(t, env, *envTag)
			}
		})
	}
}

// TestInfrastructurePerformance tests infrastructure performance characteristics
func TestInfrastructurePerformance(t *testing.T) {
	t.Parallel()

	subscriptionID := getAzureSubscriptionID(t)
	uniqueID := strings.ToLower(random.UniqueId())
	resourceGroupName := fmt.Sprintf("rg-terratest-perf-%s", uniqueID)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{"environments/dev.tfvars"},
		Vars: map[string]interface{}{
			"resource_group_name":     resourceGroupName,
			"storage_account_name":    fmt.Sprintf("stperf%s", strings.ToLower(uniqueID)),
			"key_vault_name":          fmt.Sprintf("kv-perf-%s", uniqueID),
			"web_app_name":            fmt.Sprintf("webapp-perf-%s", uniqueID),
			"virtual_network_name":    fmt.Sprintf("vnet-perf-%s", uniqueID),
			"subnet_name":             fmt.Sprintf("subnet-perf-%s", uniqueID),
			"nsg_name":                fmt.Sprintf("nsg-perf-%s", uniqueID),
			"storage_container_name":  fmt.Sprintf("container%s", strings.ToLower(uniqueID)),
			"app_service_plan_name":   fmt.Sprintf("asp-perf-%s", uniqueID),
			"prefix":                  "tf",
			"environment":             "performance",
			"suffix":                  "01",
			"project_name":            "terratest",
			"owner":                   "test-team",
		},
		RetryableTerraformErrors: map[string]string{
			".*": "Will try again",
		},
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})

	// Measure deployment time
	start := time.Now()
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	deploymentTime := time.Since(start)

	// Assert deployment completed within reasonable time (adjust as needed)
	assert.Less(t, deploymentTime, 15*time.Minute, "Deployment took too long")

	// Verify resources are accessible
	resourceGroupName = terraform.Output(t, terraformOptions, "resource_group_name")
	assert.True(t, azure.ResourceGroupExists(t, resourceGroupName, subscriptionID))

	t.Logf("Infrastructure deployment completed in %v", deploymentTime)
}

// getAzureSubscriptionID retrieves the Azure subscription ID from environment or skips test
func getAzureSubscriptionID(t *testing.T) string {
	subscriptionID := getEnvVar(t, "ARM_SUBSCRIPTION_ID", "")
	if subscriptionID == "" {
		t.Skip("ARM_SUBSCRIPTION_ID environment variable not set")
	}
	return subscriptionID
}
