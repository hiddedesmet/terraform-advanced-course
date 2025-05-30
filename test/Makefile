.PHONY: help test test-validation test-modules test-security test-performance test-dr test-all test-suite clean setup

help:
	@echo "Available targets:"
	@echo "  test            - Run validation tests (no Azure credentials required)"
	@echo "  test-validation - Run Terraform validation tests"
	@echo "  test-modules    - Run module tests (tagging works without Azure, others skip)"
	@echo "  test-security   - Run security tests (requires Azure credentials)"
	@echo "  test-performance- Run performance tests (requires Azure credentials)"
	@echo "  test-dr         - Run disaster recovery tests"
	@echo "  test-all        - Run all tests"
	@echo "  test-suite      - Run comprehensive test suite"
	@echo "  setup           - Install Go dependencies"
	@echo "  clean           - Clean up test artifacts"
	@echo "  status          - Show environment status"

setup:
	@echo "Setting up test environment..."
	cd .. && go mod tidy
	cd .. && go mod download
	@echo "Setup complete!"

test: test-validation

test-validation:
	@echo "Running Terraform validation tests..."
	cd .. && go test -v test/terraform_validation_test.go test/test_helpers.go -timeout 30m

test-modules:
	@echo "Running module tests..."
	cd .. && go test -v test/terraform_modules_test.go test/test_helpers.go -timeout 30m

test-security:
	@echo "Running security tests..."
	cd .. && go test -v test/terraform_security_test.go test/test_helpers.go -timeout 30m

test-performance:
	@echo "Running performance tests..."
	cd .. && go test -v test/terraform_performance_test.go test/test_helpers.go -timeout 45m

test-dr:
	@echo "Running disaster recovery tests..."
	cd .. && go test -v test/terraform_disaster_recovery_test.go test/test_helpers.go -timeout 30m

test-all: test-validation test-modules test-dr test-security test-performance

test-suite:
	@echo "Running comprehensive test suite..."
	cd .. && go test -v test/*.go -timeout 60m

clean:
	@echo "Cleaning up test artifacts..."
	find .. -name "*.tfstate*" -delete
	find .. -name ".terraform" -type d -exec rm -rf {} + 2>/dev/null || true
	find .. -name ".terraform.lock.hcl" -delete
	cd .. && go clean -testcache
	@echo "Cleanup complete!"

status:
	@echo "=== Terratest Environment Status ==="
	@echo ""
	@echo "Go version:"
	@go version
	@echo ""
	@echo "Terraform version:"
	@terraform version
	@echo ""
	@echo "Test files present:"
	@ls -la *.go | wc -l | awk '{print $$1 " Go test files found"}'
	@echo ""
	@echo "Azure Environment Variables:"
	@if [ -n "$$ARM_SUBSCRIPTION_ID" ]; then echo "✓ ARM_SUBSCRIPTION_ID is set"; else echo "✗ ARM_SUBSCRIPTION_ID not set"; fi
	@if [ -n "$$ARM_CLIENT_ID" ]; then echo "✓ ARM_CLIENT_ID is set"; else echo "✗ ARM_CLIENT_ID not set"; fi
	@if [ -n "$$ARM_CLIENT_SECRET" ]; then echo "✓ ARM_CLIENT_SECRET is set"; else echo "✗ ARM_CLIENT_SECRET not set"; fi
	@if [ -n "$$ARM_TENANT_ID" ]; then echo "✓ ARM_TENANT_ID is set"; else echo "✗ ARM_TENANT_ID not set"; fi
	@echo ""
	@echo "Available test targets:"
	@echo "- make test            # Run validation tests"
	@echo "- make test-modules    # Run module tests"
	@echo "- make test-all        # Run all tests"
	@echo "- make setup           # Setup dependencies"
	@echo "- make clean           # Clean up artifacts"
