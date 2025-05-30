package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDisasterRecoveryBasics(t *testing.T) {
	t.Parallel()

	// Basic disaster recovery configuration test
	// This tests the disaster recovery modules and configurations
	// without actually deploying resources to avoid Azure SDK compatibility issues

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		VarFiles:     []string{"../environments/dev.tfvars"},
		Vars: map[string]interface{}{
			"environment":                  "dr-test",
			"location":                     "East US 2",
			"backup_retention_days":        30,
			"geo_redundant_backup_enabled": true,
		},
		NoColor:   true,
		PlanOnly:  true,
	})

	// Test the configuration by initializing and planning
	terraform.Init(t, terraformOptions)
	
	// Use basic validate without variables first
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")
	
	// Plan to ensure DR configuration is valid
	planOutput := terraform.Plan(t, terraformOptions)
	assert.Contains(t, planOutput, "azurerm_storage_account")
	
	t.Log("Disaster recovery configuration validation completed successfully")
}

func TestBackupConfiguration(t *testing.T) {
	t.Parallel()

	// Test backup configuration for key vault
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/keyvault",
		Vars: map[string]interface{}{
			"environment":           "backup-test",
			"location":             "West US 2",
			"soft_delete_enabled":  true,
			"purge_protection":     true,
			"backup_vault_enabled": true,
		},
		NoColor:   true,
		PlanOnly:  true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)
	
	// Use basic validate without variables first
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/keyvault",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")
	
	planOutput := terraform.Plan(t, terraformOptions)
	assert.Contains(t, planOutput, "azurerm_key_vault")
	
	t.Log("Backup configuration validation completed successfully")
}

func TestRecoveryTimeObjective(t *testing.T) {
	t.Parallel()

	// Test RTO configuration for web app
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/webapp",
		Vars: map[string]interface{}{
			"environment":           "rto-test",
			"location":             "Central US",
			"app_service_plan_sku": "P1v2",
			"auto_scaling_enabled": true,
			"health_check_enabled": true,
			"deployment_slots":     2,
		},
		NoColor:   true,
		PlanOnly:  true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)
	
	// Use basic validate without variables first
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/webapp",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")
	
	planOutput := terraform.Plan(t, terraformOptions)
	assert.Contains(t, planOutput, "azurerm_service_plan")
	
	t.Log("RTO configuration validation completed successfully")
}

func TestDataReplicationConfiguration(t *testing.T) {
	t.Parallel()

	// Test data replication configuration for storage
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		Vars: map[string]interface{}{
			"environment":                   "replication-test",
			"location":                     "North Europe",
			"account_replication_type":     "GRS",
			"cross_tenant_replication":     false,
			"versioning_enabled":           true,
			"change_feed_enabled":          true,
			"point_in_time_restore_enabled": true,
		},
		NoColor:   true,
		PlanOnly:  true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)
	
	// Use basic validate without variables first
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/storage",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")
	
	planOutput := terraform.Plan(t, terraformOptions)
	assert.Contains(t, planOutput, "azurerm_storage_account")
	
	t.Log("Data replication configuration validation completed successfully")
}

func TestNetworkFailoverConfiguration(t *testing.T) {
	t.Parallel()

	// Test network failover configuration
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/network",
		Vars: map[string]interface{}{
			"environment":             "failover-test",
			"location":               "UK South",
			"address_space":          []string{"10.1.0.0/16"},
			"availability_zones":     []string{"1", "2", "3"},
			"enable_ddos_protection": true,
			"enable_vm_protection":   true,
		},
		NoColor:   true,
		PlanOnly:  true,
	})

	// Test the configuration
	terraform.Init(t, terraformOptions)
	
	// Use basic validate without variables first
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../modules/network",
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")
	
	planOutput := terraform.Plan(t, terraformOptions)
	assert.Contains(t, planOutput, "azurerm_virtual_network")
	
	t.Log("Network failover configuration validation completed successfully")
}

// Test helpers for disaster recovery scenarios
func validateDRConfiguration(t *testing.T, terraformDir string, vars map[string]interface{}) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		Vars:         vars,
		NoColor:      true,
		PlanOnly:     true,
	})

	terraform.Init(t, terraformOptions)
	
	// Use basic validate without variables first
	emptyTerraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		NoColor:      true,
	})
	terraform.RunTerraformCommand(t, emptyTerraformOptions, "validate")
	
	planOutput := terraform.Plan(t, terraformOptions)
	require.NotEmpty(t, planOutput, "Plan output should not be empty")
	
	t.Log("DR configuration validation completed for", terraformDir)
}
