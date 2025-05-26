variable "environment" {
  description = "Environment name (dev, test, prod, etc.)"
  type        = string
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "owner" {
  description = "Owner of the resource"
  type        = string
  default     = "DevOps Team"
}

variable "cost_center" {
  description = "Cost center for billing"
  type        = string
  default     = "IT-12345"
}

variable "terraform_version" {
  description = "Version of Terraform used to create the resources"
  type        = string
  default     = "1.1.0+"
}

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}
