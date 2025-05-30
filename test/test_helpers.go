package test

import (
	"math/rand"
	"os"
	"testing"
	"time"
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
