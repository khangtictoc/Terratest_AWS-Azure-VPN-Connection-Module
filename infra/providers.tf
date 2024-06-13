provider "aws" {
  region     = var.region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

provider "azurerm" {
  features {
     resource_group {
       prevent_deletion_if_contains_resources = false
     }
   }

  subscription_id   = var.az_subscription_id
  tenant_id         = var.az_tenant_id
  client_id         = var.az_client_id
  client_secret     = var.az_client_secret
}