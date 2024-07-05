output aws_vpn_connection_tunnel_status {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_connection_info.*.vgw_telemetry
    sensitive = true
}

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