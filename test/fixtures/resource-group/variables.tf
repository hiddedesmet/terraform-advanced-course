variable "name" {
  description = "Name of the resource group"
  type        = string
}

variable "location" {
  description = "Location for the resource group"
  type        = string
}

variable "tags" {
  description = "Tags for the resource group"
  type        = map(string)
  default     = {}
}
