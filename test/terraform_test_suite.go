package test

import (
	"testing"
)

// TestAllValidation runs all validation tests without requiring Azure credentials
func TestAllValidation(t *testing.T) {
	t.Run("TerraformValidation", TestTerraformValidation)
	t.Run("NamingConventions", TestNamingConventions)
	t.Run("ValidationModule", TestValidationModule)
}

// TestAllModules runs all module tests (requires Azure credentials)
func TestAllModules(t *testing.T) {
	subscriptionID := getEnvVar(t, "ARM_SUBSCRIPTION_ID", "")
	if subscriptionID == "" {
		t.Skip("Azure subscription ID not available")
	}
	
	t.Run("NetworkModule", TestNetworkModule)
	t.Run("StorageModule", TestStorageModule)
	t.Run("WebAppModule", TestWebAppModule)
	t.Run("KeyVaultModule", TestKeyVaultModule)
	t.Run("TaggingModule", TestTaggingModule)
	t.Run("ModulesIntegration", TestModulesIntegration)
}

// TestAllInfrastructure runs all infrastructure tests (requires Azure credentials)
func TestAllInfrastructure(t *testing.T) {
	subscriptionID := getEnvVar(t, "ARM_SUBSCRIPTION_ID", "")
	if subscriptionID == "" {
		t.Skip("Azure subscription ID not available")
	}
	
	t.Run("AdvancedInfrastructure", TestTerraformAdvancedInfrastructure)
}

// TestAllSecurity runs all security tests (requires Azure credentials)
func TestAllSecurity(t *testing.T) {
	subscriptionID := getEnvVar(t, "ARM_SUBSCRIPTION_ID", "")
	if subscriptionID == "" {
		t.Skip("Azure subscription ID not available")
	}
	
	t.Run("SecurityCompliance", TestSecurityCompliance)
	t.Run("DataEncryption", TestDataEncryption)
	t.Run("AccessControl", TestAccessControl)
	t.Run("ComplianceTags", TestComplianceTags)
}

// TestAllPerformance runs all performance tests (requires Azure credentials)
func TestAllPerformance(t *testing.T) {
	subscriptionID := getEnvVar(t, "ARM_SUBSCRIPTION_ID", "")
	if subscriptionID == "" {
		t.Skip("Azure subscription ID not available")
	}
	
	t.Run("PerformanceBenchmarks", TestPerformanceBenchmarks)
	t.Run("ScalabilityLimits", TestScalabilityLimits)
	t.Run("ResourceLimits", TestResourceLimits)
	t.Run("ConcurrentDeployments", TestConcurrentDeployments)
}

// TestAllDisasterRecovery runs all disaster recovery tests
func TestAllDisasterRecovery(t *testing.T) {
	t.Run("DisasterRecoveryBasics", TestDisasterRecoveryBasics)
	t.Run("BackupConfiguration", TestBackupConfiguration)
	t.Run("RecoveryTimeObjective", TestRecoveryTimeObjective)
	t.Run("DataReplicationConfiguration", TestDataReplicationConfiguration)
	t.Run("NetworkFailoverConfiguration", TestNetworkFailoverConfiguration)
}
