# Cost Management Documentation

This directory contains cost estimates, reports, and optimization recommendations for our Terraform infrastructure.

## Cost Estimation

The cost estimation reports are automatically generated using our custom script at `../../scripts/cost_estimation.py`. This tool provides:

1. **Monthly cost estimates** for all resources
2. **Optimization recommendations** with potential savings
3. **Historical cost trends** to track infrastructure spending over time

## Generated Reports

- **cost_report.json**: Raw cost data in JSON format
- **cost_report.html**: Formatted HTML report with visualizations
- **cost_history/**: Historical reports for trend analysis

## Running Cost Estimation

To generate a new cost estimate:

```bash
# Install dependencies
pip install requests

# Generate cost report
python scripts/cost_estimation.py

# Generate report with custom directory
python scripts/cost_estimation.py --terraform-dir=./environments/prod --output=./docs/costs/prod_cost_report.json
```

## Cost Optimization Best Practices

1. **Right-sizing resources** - Ensure resource SKUs match actual workload requirements
2. **Reserved Instances** - Use reserved instances for predictable workloads
3. **Autoscaling** - Implement autoscaling for variable workloads
4. **Storage Lifecycle Management** - Move infrequently accessed data to cooler storage tiers
5. **Dev/Test Subscriptions** - Use dev/test subscription pricing for non-production environments

## Cost Management Process

Our team follows this process for infrastructure cost management:

1. **Weekly cost review** during sprint planning
2. **Monthly comprehensive cost analysis**
3. **Quarterly optimization initiatives**
4. **Cost consideration** for all new features and infrastructure changes

## Integration with CI/CD

Cost estimates are automatically generated during pull request validation to ensure team members understand the cost impact of their infrastructure changes.
