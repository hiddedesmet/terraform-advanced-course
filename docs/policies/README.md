# Policy as Code Documentation

This directory contains policy validation reports generated when infrastructure code is checked against our organizational policies.

## Policy Validation

Our project implements Policy as Code using Open Policy Agent (OPA) to enforce organizational standards, security requirements, and best practices across our infrastructure. The validation is performed using the custom script at `../../scripts/validate_policies.py`.

## Generated Reports

When policy validation is run, the following reports are generated:

- **tf_plan_[timestamp].json**: The Terraform plan in JSON format
- **policy_results_[timestamp].json**: Raw validation results in JSON format
- **policy_results_[timestamp].html**: Formatted HTML report of validation results

## Policy Categories

Our policies are organized into the following categories:

1. **Security**: Enforcing security best practices and compliance requirements
2. **Cost**: Controlling and optimizing resource costs
3. **Tagging**: Ensuring proper resource tagging for organization and cost attribution
4. **Naming**: Enforcing naming conventions
5. **Compliance**: Meeting regulatory and organizational compliance requirements

## Running Policy Validation

To validate your Terraform code against our policies:

```bash
# Install OPA
brew install opa

# Run validation script
python scripts/validate_policies.py

# Run with custom directories
python scripts/validate_policies.py --terraform-dir=./environments/prod --policy-dir=./policies --output-dir=./docs/policies/prod
```

## Integration with Development Workflow

Our policies are enforced at multiple stages:

1. **Pre-Commit**: Local validation during development using pre-commit hooks
2. **Pull Request**: Automated validation in CI/CD pipeline
3. **Deployment**: Pre-deployment validation to prevent non-compliant resources

## Handling Policy Exceptions

Sometimes exceptions to policies are necessary. To handle exceptions:

1. Document the reason for the exception in an ADR
2. Get approval from the security team
3. Apply the exception in one of two ways:
   - Update the policy with a specific exception case
   - Add an approved exception tag to the resource

## Policy Development

See the main [Policies README](../../policies/README.md) for information on policy development guidelines.
