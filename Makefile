# Makefile for Terraform Advanced Course

.PHONY: help init plan-dev plan-prod apply-dev apply-prod fmt validate clean

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Terraform targets
init: ## Initialize Terraform
	terraform init

plan-dev: ## Plan for development environment
	terraform workspace select dev || terraform workspace new dev
	terraform plan -var-file=environments/dev.tfvars

plan-prod: ## Plan for production environment
	terraform workspace select prod || terraform workspace new prod
	terraform plan -var-file=environments/prod.tfvars

apply-dev: ## Apply changes for development environment
	terraform workspace select dev || terraform workspace new dev
	terraform apply -var-file=environments/dev.tfvars

apply-prod: ## Apply changes for production environment
	terraform workspace select prod || terraform workspace new prod
	terraform apply -var-file=environments/prod.tfvars

fmt: ## Format Terraform code
	terraform fmt -recursive

validate: ## Validate Terraform configuration
	terraform validate

clean: ## Clean up Terraform artifacts
	find . -name "*.tfstate*" -delete
	find . -name "*.terraform*" -type d -exec rm -rf {} +
	find . -name ".terraform.lock.hcl" -delete

# Quick development workflow
dev-workflow: init fmt validate plan-dev ## Run complete development workflow

prod-workflow: init fmt validate plan-prod ## Run complete production workflow
