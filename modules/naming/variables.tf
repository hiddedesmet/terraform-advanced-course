variable "prefix" {
  description = "Prefix to be used in resource names"
  type        = string
  default     = "tf"
}

variable "environment" {
  description = "Environment name (dev, test, prod, etc.)"
  type        = string
}

variable "suffix" {
  description = "Suffix to be used in resource names"
  type        = string
  default     = ""
}

variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "demo"
}

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}
