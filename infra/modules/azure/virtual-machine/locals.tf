locals {
    virtual_machine = {
        public_ip = module.network_interface.public_ip_address
        private_ip = module.network_interface.private_ip_address
    }
}