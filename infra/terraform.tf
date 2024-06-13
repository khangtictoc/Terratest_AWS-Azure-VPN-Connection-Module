terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.24.0"
    } 
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.84.0"
    }
  }
}