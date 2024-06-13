resource "azurerm_linux_virtual_machine" "virtual_machine" {
  name                  = var.virtual_machine_config.name
  network_interface_ids = [var.nic_id]
  size                  = var.virtual_machine_config.size
  computer_name  = var.virtual_machine_config.computer_name
  admin_username = var.virtual_machine_config.admin_username

  os_disk {
    name                 = var.virtual_machine_config.os_disk.name
    caching              = var.virtual_machine_config.os_disk.caching
    storage_account_type = var.virtual_machine_config.os_disk.storage_account_type
  }

  source_image_reference {
    publisher = var.virtual_machine_config.source_image_reference.publisher
    offer     = var.virtual_machine_config.source_image_reference.offer
    sku       = var.virtual_machine_config.source_image_reference.sku
    version   = var.virtual_machine_config.source_image_reference.version
  }

  admin_ssh_key {
    username   = var.virtual_machine_config.admin_username
    public_key = file(var.virtual_machine_config.public_key_path)
  }

  location            = var.resource_group.location
  resource_group_name = var.resource_group.name
}