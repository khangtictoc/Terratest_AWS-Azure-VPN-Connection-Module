resource "azurerm_network_interface" "network_interface" {
  name                = var.name
  ip_configuration {
    name                          = var.network_interface_config.ip_configuration.name
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = var.network_interface_config.ip_configuration.private_ip_address_allocation
    public_ip_address_id          = azurerm_public_ip.public_ip[0].id == null ? null : azurerm_public_ip.public_ip[0].id
  }

  location            = var.resource_group.location
  resource_group_name = var.resource_group.name
}

resource "azurerm_public_ip" "public_ip" {
  count               = var.network_interface_config.public_ip != null ? 1 : 0

  name                = var.network_interface_config.public_ip.name
  allocation_method   = var.network_interface_config.public_ip.allocation_method

  location            = var.resource_group.location
  resource_group_name = var.resource_group.name
}

resource "azurerm_network_interface_security_group_association" "nsg_interface_attachment" {
  network_interface_id      = azurerm_network_interface.network_interface.id
  network_security_group_id = var.network_security_group_id
}