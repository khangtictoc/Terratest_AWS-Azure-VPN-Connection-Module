locals {
  ec2_instance = {
    public_ip = module.ec2_instance.public_ip
    private_ip = module.ec2_instance.private_ip
  }
}