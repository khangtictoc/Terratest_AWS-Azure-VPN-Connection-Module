############# AZURE #############

variable "tags" {}

variable "az_subscription_id" {}
variable "az_tenant_id" {}
variable "az_client_id" {}
variable "az_client_secret" {}

variable "resource_group" {}

############# AWS #############

# General
variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "region" {}



########### VPN Site-to-Site Configuration ###########
variable s2s_vpn_connection_config {}
variable aws_vm_config_demo {}
variable azure_vm_config_demo {}