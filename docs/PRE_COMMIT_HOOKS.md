# Pre-commit Hooks Guide

This project uses pre-commit hooks to enforce code quality and consistency. These hooks run automatically before each commit to ensure that your changes meet the project's standards.

## Setup

1. Install pre-commit:

```bash
# Using pip
pip install pre-commit

# Or using Homebrew (macOS)
brew install pre-commit
```

2. Install the hooks:

```bash
pre-commit install
```

## Available Hooks

The following pre-commit hooks are configured for this project:

### Basic File Hygiene
- **trailing-whitespace**: Removes trailing whitespace
- **end-of-file-fixer**: Ensures files end with a newline
- **check-yaml**: Validates YAML files
- **check-added-large-files**: Prevents giant files from being committed
- **check-merge-conflict**: Checks for files containing merge conflict strings

### Terraform-specific Hooks
- **terraform_fmt**: Formats Terraform files according to canonical style
- **terraform_docs**: Updates README.md files in all modules with generated documentation
- **terraform_tflint**: Runs tflint to find possible errors and enforce best practices
- **terraform_validate**: Validates Terraform configuration files
- **terraform_checkov**: Scans for security and compliance issues
- **terrascan**: Detects security vulnerabilities and compliance issues
- **tfupdate**: Ensures Azure provider versions are up to date

### Security Hooks
- **gitleaks**: Scans for sensitive information and secrets

## Running Hooks Manually

Run all hooks against all files:

```bash
pre-commit run --all-files
```

Run a specific hook:

```bash
pre-commit run terraform_fmt --all-files
```

## Skipping Hooks

In rare cases when you need to bypass hooks:

```bash
git commit --no-verify -m "Your commit message"
```

⚠️ **Warning**: Skipping hooks should be done only in exceptional circumstances.

## Updating Hooks

To update hooks to their latest versions:

```bash
pre-commit autoupdate
```
