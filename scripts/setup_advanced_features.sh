#!/bin/bash

# Setup script for advanced Terraform features

echo "Setting up advanced Terraform features..."

# Create required directories if they don't exist
mkdir -p docs/diagrams docs/costs docs/drift docs/policies

# Install required Python packages
pip install -q diagrams terraform-visual graphviz

# Install OPA
if ! command -v opa &> /dev/null
then
    echo "Installing Open Policy Agent..."
    curl -L -o opa https://openpolicyagent.org/downloads/latest/opa_darwin_amd64
    chmod +x opa
    mv opa /usr/local/bin/
fi

# Install pre-commit hooks
pip install -q pre-commit
pre-commit install

# Setup environment for cost estimation
python -c "import os; os.makedirs('docs/costs/history', exist_ok=True)"

# Generate initial architecture diagram
echo "Generating initial architecture diagram..."
python scripts/generate_diagrams.py || echo "Error generating diagrams. You may need to run terraform init first."

# Run policy validation
echo "Running initial policy validation..."
python scripts/validate_policies.py || echo "Policy validation requires terraform init and opa installation."

# Create demo reports
echo "Creating sample cost and drift reports..."
python scripts/cost_estimation.py || echo "Error generating cost report."
python scripts/drift_detection.py || echo "Error running drift detection."

echo "Setup complete! The following features are now available:"
echo "- Infrastructure Visualization: scripts/generate_diagrams.py"
echo "- Cost Estimation: scripts/cost_estimation.py"
echo "- Drift Detection: scripts/drift_detection.py"
echo "- Policy Validation: scripts/validate_policies.py"
echo ""
echo "Documentation is available in the docs directory."
