package test

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

var (
	// Global variables for shared resource group
	sharedResourceGroupName string
	sharedResourceGroupLock sync.Once
	sharedLocation          = "westeurope" // Default location
)

// generateRandomString generates a random string of given length for unique resource naming
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// getEnvVar retrieves an environment variable or returns a default value
func getEnvVar(t *testing.T, key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// skipIfNoAzureCredentials skips the test if Azure credentials are not configured
func skipIfNoAzureCredentials(t *testing.T) {
	if os.Getenv("ARM_SUBSCRIPTION_ID") == "" {
		t.Skip("Azure credentials not configured. Set ARM_SUBSCRIPTION_ID environment variable.")
	}
	if os.Getenv("ARM_CLIENT_ID") == "" {
		t.Skip("Azure credentials not configured. Set ARM_CLIENT_ID environment variable.")
	}
	if os.Getenv("ARM_CLIENT_SECRET") == "" {
		t.Skip("Azure credentials not configured. Set ARM_CLIENT_SECRET environment variable.")
	}
	if os.Getenv("ARM_TENANT_ID") == "" {
		t.Skip("Azure credentials not configured. Set ARM_TENANT_ID environment variable.")
	}
}

// GetSharedResourceGroup ensures a shared resource group exists and returns its name
func GetSharedResourceGroup(t *testing.T) string {
	sharedResourceGroupLock.Do(func() {
		uniqueID := generateRandomString(8)
		sharedResourceGroupName = fmt.Sprintf("rg-terratest-shared-%s", uniqueID)

		subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
		if subscriptionID == "" {
			t.Fatal("AZURE_SUBSCRIPTION_ID environment variable not set")
		}

		// Create a shared resource group using the resource-group fixture
		rgOptions := &terraform.Options{
			TerraformDir: "./fixtures/resource-group",
			Vars: map[string]interface{}{
				"name":     sharedResourceGroupName,
				"location": sharedLocation,
				"tags": map[string]string{
					"Environment": "test",
					"Purpose":     "terratest-shared",
				},
			},
		}

		_, err := terraform.InitAndApplyE(t, rgOptions)
		if err != nil {
			t.Fatalf("Failed to create shared resource group: %v", err)
		}

		// Register cleanup to destroy the resource group after all tests complete
		t.Cleanup(func() {
			terraform.Destroy(t, rgOptions)
		})
	})

	return sharedResourceGroupName
}
