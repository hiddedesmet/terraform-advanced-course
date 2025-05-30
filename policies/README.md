# Policy as Code with Open Policy Agent (OPA)

This directory contains policy-as-code definitions using the Rego language for OPA (Open Policy Agent). These policies enforce organizational standards, security requirements, and best practices across our infrastructure.

## Policy Categories

Our policies are organized into the following categories:

1. **Security**: Enforcing security best practices and compliance requirements
2. **Cost**: Controlling and optimizing resource costs
3. **Tagging**: Ensuring proper resource tagging for organization and cost attribution
4. **Naming**: Enforcing naming conventions
5. **Compliance**: Meeting regulatory and organizational compliance requirements

## How Policies are Enforced

Our policies are enforced at multiple stages:

1. **Pre-Commit**: Local validation during development
2. **Pull Request**: Automated validation in CI/CD pipeline
3. **Runtime**: Continuous validation of deployed resources

## Running Policy Validation

To validate your Terraform code against our policies:

```bash
# Install OPA
brew install opa

# Run policy validation
opa eval --format pretty --data policies/ --input terraform.json "data.terraform.deny"

# Validate using the helper script
./scripts/validate_policies.sh
```

## Policy Development Guidelines

When developing new policies:

1. Place policies in the appropriate category directory
2. Include detailed documentation explaining the policy's purpose
3. Add test cases covering both compliant and non-compliant scenarios
4. Ensure policies are deterministic and performant

## Available Policies

### Security Policies

- **secure_storage**: Ensures storage accounts have network rules configured
- **secure_key_vault**: Validates key vault access policies and network configuration
- **secure_webapp**: Enforces HTTPS-only and minimum TLS version for web apps

### Cost Policies

- **cost_limits**: Prevents creation of high-cost resource SKUs without approval
- **unused_resources**: Identifies potentially unused resources

### Tagging Policies

- **required_tags**: Enforces mandatory tags on all resources
- **tag_validation**: Validates tag values follow specified formats

### Naming Policies

- **naming_convention**: Enforces organizational naming standards
- **resource_prefixes**: Validates resource type-specific prefix requirements

### Compliance Policies

- **data_sovereignty**: Enforces data residency requirements
- **pii_protection**: Protects personally identifiable information
