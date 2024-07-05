module "security_group" {
    source = "./security-group"

    security_group_config = var.aws_vm_config.security_group
    vpc_id = var.vpc_id

    tags = var.tags
}

module "keypair" {
    source = "./keypair"

    keypair_config = var.aws_vm_config.keypair
}

module "ec2_instance" {
    source = "./instance"

    instance_config = var.aws_vm_config.ec2
    security_group_ids = module.security_group.id
    subnet_id = var.subnet_id

    keypair_name = module.keypair.keypair_name
    tags          = var.tags
}

resource "local_file" "aws_private_ssh_key" {
    content  = module.keypair.ssh_private_key
    filename = "credentials/aws/private_key"
    file_permission = "0400"
}

resource "null_resource" "add_host_key" {
  triggers = {
    always_run = module.ec2_instance.instance_id
  }

  provisioner "local-exec" {
    command = "ssh-keyscan -t rsa ${module.ec2_instance.public_ip} >> C:\\Users\\tranh\\.ssh\\known_hosts"
  }
}
