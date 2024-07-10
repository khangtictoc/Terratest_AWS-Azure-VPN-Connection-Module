output id {
  value = azurerm_network_interface.network_interface.id
}

output private_ip_address {
  value = azurerm_network_interface.network_interface.private_ip_address
}

output public_ip_address {
  value = azurerm_public_ip.public_ip[0].ip_address
}