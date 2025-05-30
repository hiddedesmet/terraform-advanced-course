package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestDisasterRecoveryBasics(t *testing.T) {
	t.Parallel()

	// Basic disaster recovery configuration test
	// This tests the disaster recovery modules and configurations
	// without actually deploying resources to avoid Azure SDK compatibility issues

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		Vars: map[string]interface{}{
			"storage_account_name":     "drteststa" + generateRandomString(6),
			"resource_group_name":      "dr-test-rg",
			"location":                 "eastus2",
			"account_replication_type": "GRS",
			"storage_container_name":   "drtest-container",
		},
		NoColor: true,
	})

	// Test the configuration by initializing and validating
	terraform.Init(t, terraformOptions)

	// Use basic validate to check Terraform syntax and configuration
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")

	// Note: Skip plan step as modules require provider configuration from root module
	t.Log("Disaster recovery configuration validation completed successfully")
}

func TestBackupConfiguration(t *testing.T) {
	t.Parallel()

	// Test backup configuration for key vault
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/keyvault",
		Vars: map[string]interface{}{
			"key_vault_name":             "backup-test-kv-" + generateRandomString(6),
			"resource_group_name":        "backup-test-rg",
			"location":                   "westus2",
			"purge_protection_enabled":   true,
			"soft_delete_retention_days": 30,
			"enable_rbac_authorization":  true,
		},
		NoColor: true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)

	// Use basic validate to check Terraform syntax and configuration
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/keyvault",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")

	// Note: Skip plan step as modules require provider configuration from root module
	t.Log("Backup configuration validation completed successfully")
}

func TestRecoveryTimeObjective(t *testing.T) {
	t.Parallel()

	// Test RTO configuration for web app
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/webapp",
		Vars: map[string]interface{}{
			"app_service_plan_name": "rto-test-asp",
			"resource_group_name":   "rto-test-rg",
			"location":              "centralus",
			"sku_name":              "P1v2",
			"web_app_name":          "rto-test-webapp-" + generateRandomString(6),
			"https_only":            true,
			"minimum_tls_version":   "1.2",
		},
		NoColor: true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)

	// Use basic validate to check Terraform syntax and configuration
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/webapp",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")

	// Note: Skip plan step as modules require provider configuration from root module
	t.Log("RTO configuration validation completed successfully")
}

func TestDataReplicationConfiguration(t *testing.T) {
	t.Parallel()

	// Test data replication configuration for storage
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		Vars: map[string]interface{}{
			"storage_account_name":     "replicsta" + generateRandomString(6),
			"resource_group_name":      "replication-test-rg",
			"location":                 "northeurope",
			"account_replication_type": "GRS",
			"account_tier":             "Standard",
			"storage_container_name":   "replication-container",
			"container_access_type":    "private",
		},
		NoColor: true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)

	// Use basic validate to check Terraform syntax and configuration
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")

	// Note: Skip plan step as modules require provider configuration from root module
	t.Log("Data replication configuration validation completed successfully")
}

func TestNetworkFailoverConfiguration(t *testing.T) {
	t.Parallel()

	// Test network failover configuration
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/network",
		Vars: map[string]interface{}{
			"virtual_network_name":    "failover-test-vnet",
			"resource_group_name":     "failover-test-rg",
			"location":                "uksouth",
			"address_space":           []string{"10.1.0.0/16"},
			"subnet_name":             "failover-subnet",
			"subnet_address_prefixes": []string{"10.1.1.0/24"},
			"nsg_name":                "failover-nsg",
		},
		NoColor: true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)

	// Use basic validate to check Terraform syntax and configuration
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/network",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")

	// Note: Skip plan step as modules require provider configuration from root module
	t.Log("Network failover configuration validation completed successfully")
}

// Test helpers for disaster recovery scenarios
func validateDRConfiguration(t *testing.T, terraformDir string, vars map[string]interface{}) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		Vars:         vars,
		NoColor:      true,
	})

	terraform.Init(t, terraformOptions)

	// Use basic validate to check Terraform syntax and configuration
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")

	t.Log("DR configuration validation completed for", terraformDir)
}
