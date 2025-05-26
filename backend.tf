terraform {
  backend "azurerm" {
    resource_group_name   = "myTFResourceGroup"
    storage_account_name  = "terraformstatehidde"
    container_name        = "tfstate2"
    key                   = "prod.terraform.tfstate"
  }
}
