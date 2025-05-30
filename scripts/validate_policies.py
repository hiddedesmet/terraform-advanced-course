#!/usr/bin/env python3
"""
OPA Policy Validation Script

This script validates Terraform plans against Open Policy Agent (OPA) policies
to ensure compliance with organizational standards before deployment.
"""

import os
import sys
import json
import subprocess
import argparse
from datetime import datetime

def validate_with_opa(terraform_dir, policy_dir, output_dir):
    """Validate Terraform plans against OPA policies."""
    
    # Create output directory if it doesn't exist
    os.makedirs(output_dir, exist_ok=True)
    
    # Get current date for reporting
    current_date = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    report_date = datetime.now().strftime("%Y%m%d_%H%M%S")
    
    # Generate Terraform plan in JSON format
    plan_file = os.path.join(output_dir, f"tf_plan_{report_date}.json")
    
    print("Generating Terraform plan...")
    try:
        # Create plan
        subprocess.run(
            ["terraform", "-chdir=" + terraform_dir, "plan", "-out=tfplan"],
            check=True
        )
        
        # Convert to JSON
        subprocess.run(
            ["terraform", "-chdir=" + terraform_dir, "show", "-json", "tfplan"],
            check=True,
            stdout=open(plan_file, "w")
        )
        
        print(f"Terraform plan created: {plan_file}")
        
    except subprocess.CalledProcessError as e:
        print(f"Error generating Terraform plan: {e}")
        sys.exit(1)
    
    # Run OPA evaluation
    print("\nValidating against policies...")
    results_file = os.path.join(output_dir, f"policy_results_{report_date}.json")
    
    try:
        # Run OPA evaluation and capture results
        result = subprocess.run(
            ["opa", "eval", "--format", "json", "--data", policy_dir, "--input", plan_file, "data.terraform.deny"],
            check=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True
        )
        
        # Parse OPA results
        opa_results = json.loads(result.stdout)
        
        # Check for policy violations
        violations = []
        if "result" in opa_results and opa_results["result"]:
            for violation in opa_results["result"]:
                violations.append(violation)
        
        # Create validation report
        validation_report = {
            "report_date": current_date,
            "terraform_dir": terraform_dir,
            "policy_dir": policy_dir,
            "violations_found": len(violations) > 0,
            "violation_count": len(violations),
            "violations": violations
        }
        
        # Save report
        with open(results_file, "w") as f:
            json.dump(validation_report, f, indent=2)
        
        # Generate HTML report
        html_report_file = os.path.join(output_dir, f"policy_results_{report_date}.html")
        generate_html_report(validation_report, html_report_file)
        
        # Print summary
        print("\nPolicy Validation Summary:")
        print(f"- Date: {current_date}")
        print(f"- Terraform directory: {terraform_dir}")
        print(f"- Policy directory: {policy_dir}")
        print(f"- Violation count: {len(violations)}")
        
        if violations:
            print("\nPolicy Violations:")
            for i, violation in enumerate(violations, 1):
                print(f"\n{i}. {violation}")
            
            print(f"\nDetailed report: {results_file}")
            print(f"HTML report: {html_report_file}")
            
            # Exit with error if violations found
            sys.exit(1)
        else:
            print("\nSuccess! No policy violations found.")
            
        return validation_report
            
    except subprocess.CalledProcessError as e:
        print(f"Error during policy validation: {e}")
        print(f"stdout: {e.stdout}")
        print(f"stderr: {e.stderr}")
        sys.exit(1)
    except Exception as e:
        print(f"Unexpected error during policy validation: {e}")
        sys.exit(1)

def generate_html_report(validation_report, output_file):
    """Generate an HTML report from the policy validation results."""
    
    # Status color based on violations
    status_color = "#d13438" if validation_report["violations_found"] else "#107c10"
    status_text = "VIOLATIONS FOUND" if validation_report["violations_found"] else "PASSED"
    
    html_content = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>Terraform Policy Validation Report</title>
        <style>
            body {{ font-family: Arial, sans-serif; margin: 0; padding: 20px; }}
            h1, h2 {{ color: #0078d4; }}
            .summary {{ background-color: #f0f0f0; padding: 15px; border-radius: 5px; margin-bottom: 20px; }}
            .status {{ color: white; padding: 10px; border-radius: 5px; font-weight: bold; }}
            .violation {{ background-color: #fdf2f2; padding: 15px; margin-bottom: 15px; border-left: 4px solid #d13438; }}
            pre {{ background-color: #f8f8f8; padding: 10px; border-radius: 5px; overflow: auto; }}
            code {{ font-family: Consolas, Monaco, 'Andale Mono', monospace; }}
        </style>
    </head>
    <body>
        <h1>Terraform Policy Validation Report</h1>
        <div class="summary">
            <h2>Summary</h2>
            <p><strong>Report Date:</strong> {validation_report["report_date"]}</p>
            <p><strong>Project:</strong> {validation_report["terraform_dir"]}</p>
            <p><strong>Status:</strong> <span class="status" style="background-color: {status_color};">{status_text}</span></p>
            <p><strong>Violation Count:</strong> {validation_report["violation_count"]}</p>
        </div>
    """
    
    # Add violations section if there are any
    if validation_report["violations_found"]:
        html_content += """
        <div class="violations">
            <h2>Policy Violations</h2>
        """
        
        for i, violation in enumerate(validation_report["violations"], 1):
            html_content += f"""
            <div class="violation">
                <h3>Violation #{i}</h3>
                <pre><code>{json.dumps(violation, indent=2)}</code></pre>
            </div>
            """
            
        html_content += """
        </div>
        """
    else:
        html_content += """
        <div>
            <h2>No Policy Violations</h2>
            <p>All resources comply with the defined policies.</p>
        </div>
        """
    
    html_content += """
    </body>
    </html>
    """
    
    with open(output_file, 'w') as f:
        f.write(html_content)
    
    print(f"HTML report generated: {output_file}")

def main():
    parser = argparse.ArgumentParser(description="Validate Terraform plans against OPA policies")
    parser.add_argument("--terraform-dir", default=".", help="Directory containing Terraform configuration")
    parser.add_argument("--policy-dir", default="./policies", help="Directory containing OPA policy files")
    parser.add_argument("--output-dir", default="./docs/policies", help="Output directory for validation reports")
    
    args = parser.parse_args()
    validate_with_opa(args.terraform_dir, args.policy_dir, args.output_dir)

if __name__ == "__main__":
    main()
