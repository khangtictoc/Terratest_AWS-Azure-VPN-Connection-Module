resource "azurerm_network_security_group" "network_security_group" {
  name                = var.name
  
  location            = var.resource_group.location
  resource_group_name = var.resource_group.name
}

resource "azurerm_network_security_rule" "security_rule" {
  count                      = length(var.security_rule)

  name                       = var.security_rule[count.index].name
  priority                   = var.security_rule[count.index].priority
  direction                  = var.security_rule[count.index].direction
  access                     = var.security_rule[count.index].access
  protocol                   = var.security_rule[count.index].protocol
  source_port_range          = var.security_rule[count.index].source_port_range
  destination_port_range     = var.security_rule[count.index].destination_port_range
  source_address_prefix      = var.security_rule[count.index].source_address_prefix
  destination_address_prefix = var.security_rule[count.index].destination_address_prefix

  resource_group_name         = var.resource_group.name
  network_security_group_name = azurerm_network_security_group.network_security_group.name
}