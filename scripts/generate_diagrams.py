#!/usr/bin/env python3
"""
This script generates architecture diagrams from Terraform configuration.
It uses the terraform-visual and graphviz libraries.
"""

import os
import subprocess
import json
import argparse
from pathlib import Path

def generate_diagram(terraform_dir, output_dir, format="png"):
    """Generate architecture diagram from Terraform configuration."""
    # Create output directory if it doesn't exist
    os.makedirs(output_dir, exist_ok=True)
    
    # Get terraform plan in JSON format
    plan_file = os.path.join(output_dir, "terraform_plan.json")
    subprocess.run(
        ["terraform", "-chdir=" + terraform_dir, "plan", "-out=tfplan"],
        check=True
    )
    subprocess.run(
        ["terraform", "-chdir=" + terraform_dir, "show", "-json", "tfplan"],
        check=True,
        stdout=open(plan_file, "w")
    )
    
    # Generate diagram using terraform-visual
    output_file = os.path.join(output_dir, f"architecture_diagram.{format}")
    subprocess.run(
        ["terraform-visual", "--plan", plan_file, "--output", output_file],
        check=True
    )
    
    print(f"Diagram generated: {output_file}")
    return output_file

def main():
    parser = argparse.ArgumentParser(description="Generate architecture diagrams from Terraform configuration.")
    parser.add_argument("--terraform-dir", default=".", help="Directory containing Terraform configuration")
    parser.add_argument("--output-dir", default="./docs/diagrams", help="Output directory for diagrams")
    parser.add_argument("--format", default="png", choices=["png", "svg", "pdf"], help="Output format")
    
    args = parser.parse_args()
    generate_diagram(args.terraform_dir, args.output_dir, args.format)

if __name__ == "__main__":
    main()
