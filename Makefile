# Makefile for Terratest

.PHONY: help test test-all test-validation test-modules test-infrastructure clean fmt lint setup

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Test targets
test: test-validation ## Run validation tests only (fastest)

test-all: test-validation test-modules test-infrastructure ## Run all tests

test-validation: ## Run Terraform validation and planning tests
	@echo "Running validation tests..."
	cd test && go test -v -run TestTerraformValidation -timeout 10m
	cd test && go test -v -run TestTerraformPlan -timeout 10m
	cd test && go test -v -run TestNamingConventions -timeout 10m
	cd test && go test -v -run TestValidationModule -timeout 10m

test-modules: ## Run individual module tests
	@echo "Running module tests..."
	cd test && go test -v -run TestNetworkModule -timeout 20m
	cd test && go test -v -run TestStorageModule -timeout 20m
	cd test && go test -v -run TestWebAppModule -timeout 20m
	cd test && go test -v -run TestKeyVaultModule -timeout 20m

test-infrastructure: ## Run full infrastructure end-to-end test
	@echo "Running infrastructure test..."
	@echo "WARNING: This will deploy real Azure resources and may incur costs!"
	@echo "Press Ctrl+C to cancel, or wait 10 seconds to continue..."
	@sleep 10
	cd test && go test -v -run TestTerraformAdvancedInfrastructure -timeout 45m

# Development targets
setup: ## Initialize Go module and download dependencies
	cd test && go mod download
	cd test && go mod tidy

fmt: ## Format Go code
	cd test && go fmt ./...

lint: ## Run golint on test code
	cd test && golint ./...

clean: ## Clean up test artifacts
	cd test && go clean -testcache
	find . -name "*.tfstate*" -delete
	find . -name "*.terraform*" -type d -exec rm -rf {} +
	find . -name ".terraform.lock.hcl" -delete

# CI/CD targets
ci-test: setup test-validation ## Run tests suitable for CI (no real deployments)

ci-test-full: setup test-all ## Run all tests in CI (WARNING: deploys real resources)

# Docker targets (if you want to run tests in containers)
docker-test: ## Run tests in Docker container
	docker run --rm -v $(PWD):/workspace -w /workspace/test \
		-e ARM_CLIENT_ID \
		-e ARM_CLIENT_SECRET \
		-e ARM_SUBSCRIPTION_ID \
		-e ARM_TENANT_ID \
		golang:1.21 \
		sh -c "go mod download && go test -v ./... -timeout 45m"

# Environment check
check-env: ## Check if required environment variables are set
	@echo "Checking environment variables..."
	@test -n "$(ARM_CLIENT_ID)" || (echo "ARM_CLIENT_ID is not set" && exit 1)
	@test -n "$(ARM_CLIENT_SECRET)" || (echo "ARM_CLIENT_SECRET is not set" && exit 1)
	@test -n "$(ARM_SUBSCRIPTION_ID)" || (echo "ARM_SUBSCRIPTION_ID is not set" && exit 1)
	@test -n "$(ARM_TENANT_ID)" || (echo "ARM_TENANT_ID is not set" && exit 1)
	@echo "All required environment variables are set!"

# Terraform targets
tf-init: ## Initialize Terraform
	terraform init

tf-plan-dev: ## Plan for development environment
	terraform workspace select dev || terraform workspace new dev
	terraform plan -var-file=environments/dev.tfvars

tf-plan-prod: ## Plan for production environment
	terraform workspace select prod || terraform workspace new prod
	terraform plan -var-file=environments/prod.tfvars

tf-fmt: ## Format Terraform code
	terraform fmt -recursive

tf-validate: ## Validate Terraform configuration
	terraform validate
