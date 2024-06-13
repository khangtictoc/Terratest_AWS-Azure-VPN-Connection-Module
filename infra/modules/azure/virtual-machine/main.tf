module "network_security_group" {
    source = "./network-security-group"

    name = var.azure_vm_config.network_security_group.name
    security_rule  = var.azure_vm_config.network_security_group.security_rule

    resource_group = var.resource_group
}

module "network_interface" {
    source = "./network-interface"

    name = var.azure_vm_config.network_interface.name
    network_interface_config = var.azure_vm_config.network_interface
    subnet_id = var.subnet_id
    network_security_group_id = module.network_security_group.id

    resource_group = var.resource_group
}

module "virtual_machine" {
    source = "./virtual-machine"

    virtual_machine_config = var.azure_vm_config.virtual_machine
    nic_id = module.network_interface.id

    resource_group = var.resource_group

}
