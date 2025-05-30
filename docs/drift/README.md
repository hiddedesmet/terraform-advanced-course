# Infrastructure Drift Detection

This directory contains drift detection reports that identify differences between the Terraform state and the actual infrastructure in Azure.

## What is Infrastructure Drift?

Infrastructure drift occurs when the actual deployed resources in your cloud environment differ from what is defined in your Terraform code. This can happen due to:

1. Manual changes made directly in the Azure portal
2. Changes made by other automation tools
3. Changes made by other team members outside of the Terraform workflow
4. Service provider updates or automatic changes

## Drift Detection Process

Our project implements automated drift detection using the custom script at `../../scripts/drift_detection.py`. This tool:

1. Runs `terraform plan -refresh-only` to detect changes
2. Analyzes the plan output to identify specific drifted resources
3. Generates detailed reports with information about the drift
4. Notifies the team when drift is detected

## Generated Reports

- **drift_report_[timestamp].json**: Raw drift data in JSON format
- **drift_report_[timestamp].html**: Formatted HTML report with visualizations
- **drift_plan_[timestamp].txt**: Raw output from Terraform plan

## Running Drift Detection

To check for infrastructure drift:

```bash
# Run basic drift detection
python scripts/drift_detection.py

# Specify custom directories and notification settings
python scripts/drift_detection.py --terraform-dir=./environments/prod --output-dir=./docs/drift/prod --notify admin@example.com devops@example.com
```

## Integration with CI/CD

Our system performs automatic drift detection:

1. Daily scheduled checks for all environments
2. Before any deployment to detect potential conflicts
3. After deployments to verify successful application

## Remediation Process

When drift is detected, follow this process:

1. Review the drift report to understand the changes
2. Determine if the changes should be:
   - Imported into Terraform state
   - Reverted to match Terraform configuration
   - Updated in the Terraform code to match reality
3. Document the decision and actions taken
4. Verify remediation with another drift detection run
