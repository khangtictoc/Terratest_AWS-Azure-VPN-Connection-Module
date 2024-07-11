# ------------------ [TEST 1] TUNNEL STATUS TESTING ------------------
output aws_vpn_connection_tunnel_status {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_connection_info.*.vgw_telemetry
    sensitive = true
}

# ------------------ [TEST 2] VPN CONNECTION INFO TESTING ------------------
output aws_vpn_connection_info {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_connection_info.*
    sensitive = true
}

output azure_vpn_connection_info {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.azure_vpn_connection_info.*
    sensitive = true
}

output azure_vnet_gateway_info {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.azure_vpn_network_info.*.virtual_network_gateway
    sensitive = true
}

output azure_vnet_info {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.azure_vpn_network_info.*.virtual_network
    sensitive = true
}

# ------------------ [TEST 3] TEST PING BETWEEN AWS INSTANCE & AZURE VM ------------------

output aws_instance_public_ip_list {
    value = module.aws_virtual_machine_demo.*.ec2_instance.public_ip
}

output azure_vm_private_ip_list {
    value = module.azure_virtual_machine_demo.*.virtual_machine.private_ip
}