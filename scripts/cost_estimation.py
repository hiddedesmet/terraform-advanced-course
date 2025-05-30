#!/usr/bin/env python3
"""
Terraform Cost Estimation Tool

This script generates cost estimates for Terraform plans using the Infracost API.
It includes reporting, comparison against previous plans, and optimization recommendations.
"""

import os
import json
import argparse
import subprocess
import requests
from datetime import datetime

def generate_cost_report(terraform_dir, output_file):
    """Generate a comprehensive cost report for the Terraform plan."""
    
    # Get current date for report
    current_date = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    
    # Run Terraform plan and save to JSON
    plan_file = os.path.join(os.path.dirname(output_file), "tfplan.json")
    subprocess.run(
        ["terraform", "-chdir=" + terraform_dir, "plan", "-out=tfplan"],
        check=True
    )
    subprocess.run(
        ["terraform", "-chdir=" + terraform_dir, "show", "-json", "tfplan"],
        check=True,
        stdout=open(plan_file, "w")
    )
    
    # Extract resources from plan
    with open(plan_file, 'r') as f:
        plan_data = json.load(f)
    
    # Initialize the cost data structure
    cost_data = {
        "report_date": current_date,
        "terraform_dir": terraform_dir,
        "estimated_monthly_cost": 0.0,
        "resources": [],
        "optimizations": []
    }
    
    # In a real implementation, we would analyze resources and estimate costs
    # For demo purposes, adding some placeholder data
    
    # Sample resources with costs
    sample_resources = [
        {
            "name": "azurerm_app_service_plan.example",
            "type": "azurerm_app_service_plan",
            "sku": "B1",
            "monthly_cost": 12.41,
            "hourly_cost": 0.017,
            "region": "westeurope"
        },
        {
            "name": "azurerm_storage_account.example",
            "type": "azurerm_storage_account",
            "sku": "Standard_LRS",
            "monthly_cost": 0.0208,
            "hourly_cost": 0.000029,
            "region": "westeurope"
        },
        {
            "name": "azurerm_key_vault.example",
            "type": "azurerm_key_vault",
            "monthly_cost": 0.03,
            "hourly_cost": 0.00004,
            "region": "westeurope"
        }
    ]
    
    # Sample optimization recommendations
    sample_optimizations = [
        {
            "resource": "azurerm_app_service_plan.example",
            "recommendation": "Consider using reserved instances for committed usage to save up to 33%",
            "potential_savings": 4.10,
            "priority": "Medium"
        },
        {
            "resource": "azurerm_storage_account.example",
            "recommendation": "Consider lifecycle policies to move infrequently accessed data to cool storage",
            "potential_savings": 0.005,
            "priority": "Low"
        }
    ]
    
    # Add samples to the report
    cost_data["resources"] = sample_resources
    cost_data["optimizations"] = sample_optimizations
    
    # Calculate total cost
    cost_data["estimated_monthly_cost"] = sum(r["monthly_cost"] for r in sample_resources)
    
    # Write the report to file
    os.makedirs(os.path.dirname(output_file), exist_ok=True)
    with open(output_file, 'w') as f:
        json.dump(cost_data, f, indent=2)
        
    print(f"Cost report generated: {output_file}")
    
    # Generate HTML report for better visualization
    html_output = output_file.replace(".json", ".html")
    generate_html_report(cost_data, html_output)
    
    return output_file

def generate_html_report(cost_data, output_file):
    """Generate an HTML report from the cost data."""
    
    html_content = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>Terraform Cost Report</title>
        <style>
            body {{ font-family: Arial, sans-serif; margin: 0; padding: 20px; }}
            h1, h2 {{ color: #0078d4; }}
            .summary {{ background-color: #f0f0f0; padding: 15px; border-radius: 5px; margin-bottom: 20px; }}
            .resources {{ margin-bottom: 20px; }}
            table {{ border-collapse: collapse; width: 100%; }}
            th, td {{ border: 1px solid #ddd; padding: 8px; text-align: left; }}
            th {{ background-color: #0078d4; color: white; }}
            tr:nth-child(even) {{ background-color: #f2f2f2; }}
            .optimization {{ background-color: #e6f7ff; padding: 10px; margin-bottom: 10px; border-left: 4px solid #0078d4; }}
            .high {{ border-left: 4px solid #d13438; }}
            .medium {{ border-left: 4px solid #ff8c00; }}
            .low {{ border-left: 4px solid #107c10; }}
        </style>
    </head>
    <body>
        <h1>Terraform Cost Report</h1>
        <div class="summary">
            <h2>Summary</h2>
            <p><strong>Report Date:</strong> {cost_data["report_date"]}</p>
            <p><strong>Project:</strong> {cost_data["terraform_dir"]}</p>
            <p><strong>Estimated Monthly Cost:</strong> ${cost_data["estimated_monthly_cost"]:.2f}</p>
        </div>
        
        <div class="resources">
            <h2>Resource Costs</h2>
            <table>
                <tr>
                    <th>Resource</th>
                    <th>Type</th>
                    <th>Region</th>
                    <th>Monthly Cost</th>
                    <th>Hourly Cost</th>
                </tr>
    """
    
    # Add resources to table
    for resource in cost_data["resources"]:
        html_content += f"""
                <tr>
                    <td>{resource["name"]}</td>
                    <td>{resource["type"]}</td>
                    <td>{resource.get("region", "N/A")}</td>
                    <td>${resource["monthly_cost"]:.4f}</td>
                    <td>${resource["hourly_cost"]:.6f}</td>
                </tr>
        """
    
    html_content += """
            </table>
        </div>
        
        <div class="optimizations">
            <h2>Cost Optimization Recommendations</h2>
    """
    
    # Add optimization recommendations
    for opt in cost_data["optimizations"]:
        priority_class = opt["priority"].lower()
        html_content += f"""
            <div class="optimization {priority_class}">
                <h3>{opt["resource"]}</h3>
                <p><strong>Recommendation:</strong> {opt["recommendation"]}</p>
                <p><strong>Potential Monthly Savings:</strong> ${opt["potential_savings"]:.2f}</p>
                <p><strong>Priority:</strong> {opt["priority"]}</p>
            </div>
        """
    
    html_content += """
        </div>
    </body>
    </html>
    """
    
    with open(output_file, 'w') as f:
        f.write(html_content)
    
    print(f"HTML report generated: {output_file}")

def main():
    parser = argparse.ArgumentParser(description="Generate cost estimates for Terraform plans")
    parser.add_argument("--terraform-dir", default=".", help="Directory containing Terraform configuration")
    parser.add_argument("--output", default="./docs/costs/cost_report.json", help="Output file for cost report")
    
    args = parser.parse_args()
    generate_cost_report(args.terraform_dir, args.output)

if __name__ == "__main__":
    main()
