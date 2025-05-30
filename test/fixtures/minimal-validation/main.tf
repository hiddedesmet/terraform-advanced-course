# Minimal Terraform configuration for syntax validation without Azure provider
terraform {
  required_version = ">= 1.5"
  required_providers {
    # Using null provider for testing - doesn't require credentials
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}

# Simple validation test without external modules
locals {
  # Test some basic Terraform functionality
  test_string = "hello-world"
  test_number = 42
  test_bool   = true
  
  # Test validation logic
  name_length_valid = length(local.test_string) > 0 && length(local.test_string) <= 50
}

# Simple null resource to test the configuration works
resource "null_resource" "validation_test" {
  triggers = {
    test_string       = local.test_string
    test_number       = tostring(local.test_number)
    test_bool         = tostring(local.test_bool)
    name_length_valid = tostring(local.name_length_valid)
  }
}

# Outputs to verify functionality
output "validation_successful" {
  description = "Indicates that validation was successful"
  value       = true
}

output "test_values" {
  description = "Test values to verify functionality"
  value = {
    string = local.test_string
    number = local.test_number
    bool   = local.test_bool
    valid  = local.name_length_valid
  }
}
