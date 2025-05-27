locals {
  # Common tags to be applied to all resources
  common_tags = {
    Environment      = var.environment
    Project          = var.project_name
    Owner            = var.owner
    CostCenter       = var.cost_center
    ManagedBy        = "Terraform"
    CreationDateTime = formatdate("YYYY-MM-DD hh:mm:ss ZZZ", timestamp())
    TerraformVersion = "${var.terraform_version}"
  }

  # Merge common tags with passed-in tags
  tags = merge(local.common_tags, var.tags)
}
