resource "tls_private_key" "pk" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "my_keypair" {
  key_name   = var.keypair_config.name    # Create a "myKey" to AWS!!
  public_key = tls_private_key.pk.public_key_openssh
}


