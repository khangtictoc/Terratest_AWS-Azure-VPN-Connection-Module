resource "aws_instance" "instance" {
  ami           = var.instance_config.ami_id
  instance_type = var.instance_config.instance_type
  associate_public_ip_address = var.instance_config.associate_public_ip_address

  vpc_security_group_ids = [var.security_group_ids]
  subnet_id              = var.subnet_id
  key_name               = var.keypair_name
  user_data = var.instance_config.user_data_file

  ebs_block_device {
    device_name = "/dev/sda1"
    volume_size = 30
  }
  tags = merge(
    var.tags,
    {
      "Name" = var.instance_config.name
    }
  )
}

