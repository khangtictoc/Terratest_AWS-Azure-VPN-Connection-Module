# output vpn_s2s_aws_azure_debug {
#     value = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_connection_info
#     sensitive = true
# }

output aws_vpn_connection_unittest {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_connection_info
    sensitive = true
}

output network_unittest {
    value = module.vpn_s2s_aws_azure.vpn_connection_output.aws_vpn_network_info
}