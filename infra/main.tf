resource "azurerm_resource_group" "vpn_resource_group" {
    count = length(var.resource_group)

    name     = var.resource_group[count.index].name
    location = var.resource_group[count.index].location

    tags = var.tags
}


module "vpn_s2s_aws_azure" {
    source = "./modules/s2s-aws-azure-vpn"

    s2s_vpn_connection_config = var.s2s_vpn_connection_config
    resource_group = slice(var.resource_group, 0, 2)
    tags = var.tags

    depends_on = [azurerm_resource_group.vpn_resource_group]
}


module "aws_virtual_machine_demo" {
    source = "./modules/aws/virtual-machine"
    count = length(var.aws_vm_config_demo)

    aws_vm_config = var.aws_vm_config_demo[count.index]
    vpc_id = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_network_info.vpc.id
    subnet_id = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_network_info.subnet[0].id

    tags = var.tags
}

module "azure_virtual_machine_demo" {
    source = "./modules/azure/virtual-machine"
    count = length(var.azure_vm_config_demo)

    azure_vm_config = var.azure_vm_config_demo[count.index]
    subnet_id = module.vpn_s2s_aws_azure.vpn_connection_output.azure_vpn_network_info[count.index].subnet[1].id

    tags = var.tags
    resource_group = var.resource_group[0]

    depends_on = [azurerm_resource_group.vpn_resource_group]
}