#!/usr/bin/env python3
"""
Terraform Drift Detection Tool

This script detects and reports on drift between the Terraform state and actual
infrastructure in Azure. It helps identify resources that have been changed outside
of Terraform control.
"""

import os
import json
import argparse
import subprocess
from datetime import datetime
from pathlib import Path
import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

def detect_drift(terraform_dir, output_dir):
    """Detect drift between Terraform state and actual infrastructure."""
    
    # Create output directory if it doesn't exist
    os.makedirs(output_dir, exist_ok=True)
    
    # Get current date for report
    current_date = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    report_date = datetime.now().strftime("%Y%m%d_%H%M%S")
    
    # Run terraform plan to detect drift
    plan_output_file = os.path.join(output_dir, f"drift_plan_{report_date}.txt")
    
    # Use -refresh-only to just check for drift without proposing changes to fix
    try:
        result = subprocess.run(
            ["terraform", "-chdir=" + terraform_dir, "plan", "-refresh-only", "-detailed-exitcode"],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True
        )
        
        # Write plan output to file
        with open(plan_output_file, 'w') as f:
            f.write(result.stdout)
            f.write(result.stderr)
        
        # Terraform returns:
        # 0 - No changes
        # 1 - Error
        # 2 - Changes present (drift detected)
        has_drift = result.returncode == 2
        error_occurred = result.returncode == 1
        
        # Create drift report
        drift_report = {
            "report_date": current_date,
            "terraform_dir": terraform_dir,
            "drift_detected": has_drift,
            "error_occurred": error_occurred,
            "plan_output_file": plan_output_file,
            "drifted_resources": []
        }
        
        # If drift detected, parse the plan output to identify drifted resources
        if has_drift:
            # In a real implementation, we would parse the plan output to identify specific drifted resources
            # For this example, adding placeholder drifted resources
            drift_report["drifted_resources"] = [
                {
                    "name": "azurerm_storage_account.example",
                    "type": "azurerm_storage_account",
                    "attribute": "tags",
                    "expected": {"Environment": "Production"},
                    "actual": {"Environment": "Production", "Owner": "DevOps"}
                },
                {
                    "name": "azurerm_network_security_group.example",
                    "type": "azurerm_network_security_group",
                    "attribute": "security_rule",
                    "details": "Security rule was added outside of Terraform"
                }
            ]
        
        # Write drift report to file
        report_file = os.path.join(output_dir, f"drift_report_{report_date}.json")
        with open(report_file, 'w') as f:
            json.dump(drift_report, f, indent=2)
        
        # Generate HTML report
        html_report_file = os.path.join(output_dir, f"drift_report_{report_date}.html")
        generate_html_report(drift_report, html_report_file)
        
        print(f"Drift detection completed. Report saved to {report_file} and {html_report_file}")
        
        return drift_report
        
    except Exception as e:
        print(f"Error running drift detection: {e}")
        return {"error": str(e), "report_date": current_date}

def generate_html_report(drift_report, output_file):
    """Generate an HTML report from the drift detection results."""
    
    # Status color based on drift
    status_color = "#d13438" if drift_report["drift_detected"] else "#107c10"
    status_text = "DRIFT DETECTED" if drift_report["drift_detected"] else "NO DRIFT"
    
    if drift_report.get("error_occurred", False):
        status_color = "#ff8c00"
        status_text = "ERROR OCCURRED"
    
    html_content = f"""
    <!DOCTYPE html>
    <html>
    <head>
        <title>Terraform Drift Report</title>
        <style>
            body {{ font-family: Arial, sans-serif; margin: 0; padding: 20px; }}
            h1, h2 {{ color: #0078d4; }}
            .summary {{ background-color: #f0f0f0; padding: 15px; border-radius: 5px; margin-bottom: 20px; }}
            .status {{ color: white; padding: 10px; border-radius: 5px; font-weight: bold; }}
            table {{ border-collapse: collapse; width: 100%; margin-top: 20px; }}
            th, td {{ border: 1px solid #ddd; padding: 8px; text-align: left; }}
            th {{ background-color: #0078d4; color: white; }}
            tr:nth-child(even) {{ background-color: #f2f2f2; }}
            .drift-resource {{ background-color: #fdf2f2; padding: 10px; margin-bottom: 10px; border-left: 4px solid #d13438; }}
        </style>
    </head>
    <body>
        <h1>Terraform Drift Report</h1>
        <div class="summary">
            <h2>Summary</h2>
            <p><strong>Report Date:</strong> {drift_report["report_date"]}</p>
            <p><strong>Project:</strong> {drift_report["terraform_dir"]}</p>
            <p><strong>Status:</strong> <span class="status" style="background-color: {status_color};">{status_text}</span></p>
        </div>
    """
    
    # Add drifted resources section if drift was detected
    if drift_report["drift_detected"] and drift_report.get("drifted_resources"):
        html_content += """
        <div class="resources">
            <h2>Drifted Resources</h2>
        """
        
        for resource in drift_report["drifted_resources"]:
            html_content += f"""
            <div class="drift-resource">
                <h3>{resource["name"]}</h3>
                <p><strong>Type:</strong> {resource["type"]}</p>
                <p><strong>Changed Attribute:</strong> {resource["attribute"]}</p>
            """
            
            if "expected" in resource and "actual" in resource:
                html_content += f"""
                <table>
                    <tr>
                        <th>Expected</th>
                        <th>Actual</th>
                    </tr>
                    <tr>
                        <td>{json.dumps(resource["expected"], indent=2)}</td>
                        <td>{json.dumps(resource["actual"], indent=2)}</td>
                    </tr>
                </table>
                """
            elif "details" in resource:
                html_content += f"""
                <p><strong>Details:</strong> {resource["details"]}</p>
                """
                
            html_content += """
            </div>
            """
            
        html_content += """
        </div>
        """
    elif drift_report.get("error_occurred", False):
        html_content += """
        <div>
            <h2>Error Information</h2>
            <p>An error occurred while checking for drift. Please check the plan output file for details.</p>
        </div>
        """
    else:
        html_content += """
        <div>
            <h2>No Drift Detected</h2>
            <p>All resources match the Terraform state. No infrastructure drift detected.</p>
        </div>
        """
    
    html_content += """
    </body>
    </html>
    """
    
    with open(output_file, 'w') as f:
        f.write(html_content)
    
    print(f"HTML report generated: {output_file}")

def notify_team(drift_report, recipients, smtp_server="localhost"):
    """Send email notification about drift to the team."""
    
    if not drift_report.get("drift_detected") and not drift_report.get("error_occurred"):
        print("No drift detected, no notification sent.")
        return
    
    # Create email
    msg = MIMEMultipart()
    msg['Subject'] = f"[ALERT] Terraform Drift Detected - {drift_report['terraform_dir']}" if drift_report.get("drift_detected") else f"[ERROR] Terraform Drift Check Failed - {drift_report['terraform_dir']}"
    msg['From'] = "terraform-monitor@example.com"
    msg['To'] = ", ".join(recipients)
    
    # Email body
    body = f"""
    Dear Team,
    
    {'Drift has been detected' if drift_report.get('drift_detected') else 'An error occurred while checking for drift'} in the Terraform project at {drift_report['terraform_dir']} on {drift_report['report_date']}.
    
    {'The following resources have drifted from their Terraform state:' if drift_report.get('drift_detected') else 'Please investigate the error:'}
    """
    
    if drift_report.get("drift_detected") and drift_report.get("drifted_resources"):
        for resource in drift_report["drifted_resources"]:
            body += f"\n- {resource['name']} ({resource['type']}): {resource.get('attribute', 'unknown attribute')} changed"
    
    body += """
    
    Please review the full report and take appropriate action to reconcile the differences.
    
    Regards,
    Terraform Monitoring System
    """
    
    msg.attach(MIMEText(body, 'plain'))
    
    # In a real implementation, you would send the email here
    # For this example, we'll just print the email content
    print("\n--- Email Notification (Not Actually Sent) ---")
    print(f"Subject: {msg['Subject']}")
    print(f"To: {msg['To']}")
    print(body)
    print("--- End Email ---\n")

def main():
    parser = argparse.ArgumentParser(description="Detect drift between Terraform state and actual infrastructure")
    parser.add_argument("--terraform-dir", default=".", help="Directory containing Terraform configuration")
    parser.add_argument("--output-dir", default="./docs/drift", help="Output directory for drift reports")
    parser.add_argument("--notify", nargs='+', help="Email addresses to notify if drift is detected")
    
    args = parser.parse_args()
    
    drift_report = detect_drift(args.terraform_dir, args.output_dir)
    
    if args.notify and (drift_report.get("drift_detected") or drift_report.get("error_occurred")):
        notify_team(drift_report, args.notify)

if __name__ == "__main__":
    main()
